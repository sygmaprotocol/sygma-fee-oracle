# Chainbridge Fee Oracle

Chainbridge fee oracle is a go implemented, centralized service that provides the endpoints to Chainbridge UI
for all necessary data related to bridging fee.

# Architecture Overview
There are three main parts in fee oracle codebase in an abstract way: `App base`, `Data fetcher`, `Data provider`.

1. App base: This is the combination of `base`, `app`, `config` packages. App base loads the config, preforms health check and starts the server when fee oracle launches. It maintains the cleanup process when the app gets the termination signal.
2. Data fetcher: This is the combination of `oracle`, `store`, `cronjob` packages. Data fetcher follows the scheduled jobs to fetch data from registered external oracles and store data into store module.
3. Data provider: This is the combination of `store`, `api`, `consensus`, `identity` packages. Data provider queries the store based on the request from the endpoint, refines the data based on the configured `strategy`, then returns the data along with the signature of the fee oracle identity key.

# Installation & Build
**Make sure `Go1.17` has been installed.**  

This will clone the main branch of fee oracle codebase to your `workplace` dir and compile the binary into your
`$GOPATH/bin`
```
$ mkdir workplace && cd workplace  
$ git clone https://github.com/ChainSafe/chainbridge-fee-oracle.git
$ make install
```

# Configuration
Fee oracle needs three config files in the `./` dir of the codebase:
1. `config.yaml`: this is the fee oracle application config file.
2. `domain.json`: this is the domain config file.
3. `resource.json`: this is the resource config file.

### Application config file `config.yaml`
Template of the config.yaml can be found in `./config/config.template.yaml`.

### Domain config file `domain.json`
This file indicates all the domains the fee oracle needs to fetch data for. Details need to be matched with 
Chainbridge core configuration, such as `id`.

Example:
```json
{
  "domains": [
    {
      "id": 1,
      "name": "ethereum",
      "baseCurrencyFullName": "ether",
      "baseCurrencySymbol": "eth",
      "addressPrefix": "0x"
    },
    {
      "id": 2,
      "name": "polygon",
      "baseCurrencyFullName": "matic",
      "baseCurrencySymbol": "matic",
      "addressPrefix": "0x"
    }
  ]
}
```

### Resource config file `resource.json`
This file indicates all the resources(assets) the fee oracle needs to fetch data for. Details need to be matched with the asset contract that 
deployed on the target chain. For example, USDT token address on ethereum(domainId is 1 in domain.json) is `0xdAC17F958D2ee523a2206206994597C13D831ec7`.
For the native current address, such as Ether on Ethereum, Matic on Polygon, the `tokenAddress` is always `0x0000000000000000000000000000000000000000`

```json
{
  "resources": [
    {
      "symbol": "eth",
      "decimal": 18,
      "tokenAddress": "0x0000000000000000000000000000000000000000",
      "domainId": 1
    },
    {
      "symbol": "usdt",
      "decimal": 6,
      "tokenAddress": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
      "domainId": 1
    },
    {
      "symbol": "matic",
      "decimal": 18,
      "tokenAddress": "0x0000000000000000000000000000000000000000",
      "domainId": 2
    },
    {
      "symbol": "doge",
      "decimal": 18,
      "tokenAddress": "0x8A953CfE442c5E8855cc6c61b1293FA648BAE472",
      "domainId": 2
    }
  ]
}
```

# Fee Oracle Identity
Each fee oracle server associates with a private key, which is used to sign the endpoint response data.
There should be a `keyfile.priv` keyfile in the root dir of the fee oracle codebase, or you can specify which keyfile to use in CLI. 

**Fee oracle provides [key generation CLI](#keycli), keyfile needs to be generated separately.**

# Quick Start
To quickly start from makefile, make sure `config.yaml`, `domain.json`, `resource.json` and `keyfile.priv` are ready in the root dir of the codebase, then execute:
  
`$ make start`

# Command Line
Fee oracle provides CLI.  

For general help:`$ chainbridge-fee-oracle -h`  

#### `$ chainbridge-fee-oracle server`
```
Start ChainBridge fee oracle main service

Usage:
  chainbridge-fee-oracle server [flags]

Flags:
  -c, --config_path string            
  -d, --domain_config_path string     
  -h, --help                          help for server
  -t, --key_type string               Support: secp256k1
  -k, --keyfile_path string           
  -r, --resource_config_path string
```

#### <a id="keycli"></a>`$ chainbridge-fee-oracle key-generate`
```
Start ChainBridge fee oracle identity key generation

Usage:
  chainbridge-fee-oracle key-generate [flags]

Flags:
  -h, --help                  help for key-generate
  -t, --key_type string       Support: secp256k1 (default "secp256k1")
  -k, --keyfile_path string   Output dir for generated key file, filename is required with .priv as file extension (default "keyfile.priv")
```

# Unit Test
`$ make test`

# Gosec Checking
`$ make check`

# Lint Checking
`$ make lint`

# Using Docker
Fee oracle provides a Dockerfile to containerize the codebase.  
To build docker image:
```
$ docker build -t fee_oracle .
```
To run docker container:
```
$ docker run -p 8091:8091 -it fee_oracle
```
`8091` will be the exposed part for the endpoint access.

**Note**: fee oracle requires a private key file when starting, this key file must be a `secp256k1` type and named as `keyfile.priv` and put in the `./` dir of the codebase when building docker image,
if no keyfile exists in `./` dir, fee oracle will auto generate a `secp256k1` keyfile to use.

# End to End Test
This will start fee oracle service, ganache-cli locally, install `solcjs`, `abigen` and generate contract go binding code, deploy fee handler contracts to local ganache.  
`$ make e2e-test`

# EVN Params
Fee oracle loads important configs and prikey from files in CLI flags; however, the following EVN params will suppress CLI flags if provided.  
Note: if `REMOTE_PARAM_OPERATOR_ENABLE` is set to `true`, valid credentials of the remote service must be setup.
```text
IDENTITY_KEY=                                                // fee oracle prikey in hex, without 0x prefix 
IDENTITY_KEY_TYPE=secp256k1                                  // fee oracle prikey type, only support secp256k1 for now
WORKING_ENV=production                                       // fee oracle app running mode: dev or production
LOG_LEVEL=4                                                  // log level, 4 is info, 5 is debug, using 4 on production
HTTP_SERVER_MODE=release                                     // fee oracle http server running mode: debug or release
HTTP_SERVER_PORT=8091                                        // fee oracle http server exposed port
CONVERSION_RATE_JOB_FREQUENCY="* * * * *"                    // conversion rate job frequency, using cron schedule expressions(https://crontab.guru)
GAS_PRICE_JOB_FREQUENCY="* * * * *"                          // gas price job frequency, using cron schedule expressions(https://crontab.guru)
ETHERSCAN_API_KEY=                                           // api key of etherscan
POLYGONSCAN_API_KEY=                                         // api key of polygonscan
COINMARKETCAP_API_KEY=                                       // api key of coinmarketcap
DATA_VALID_INTERVAL=3600                                     // Time of valid fee oracle response data in seconds
CONVERSION_RATE_PAIRS=eth,usdt,matic,usdt                    // conversion rate pairs that enabled for price fetching. Must be paired
REMOTE_PARAM_OPERATOR_ENABLE=true                            // enable remote param operator, only enable this when deploying to remote environment like staging or prod
```

# API Documentation
[Swagger API Doc](https://app.swaggerhub.com/apis-docs/cb-fee-oracle/fee-oracle/1.0.0)

# License
_GNU Lesser General Public License v3.0_