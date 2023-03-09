# Sygma Fee Oracle

Sygma fee oracle is a go implemented, centralized service that provides the endpoints to Sygma UI
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
$ git clone https://github.com/ChainSafe/Sygma-fee-oracle.git
$ make install
```

# Configuration
Fee oracle needs three config files in the `./` dir of the codebase:
1. `config.yaml`: this is the fee oracle application config file.
2. `domain.json`: this is the domain config file.

### Application config file `config.yaml`
Template of the config.yaml can be found in `./config/config.template.yaml`.

### Domain config
For domain configuration, it is posible to use local file, or the [shared configuration]file that is uploaded to some remote service (eg. ipfs). Depending on what value is provided to the flag `domain_config_path`, the aplication will load the domain configuration from local file or from the remote service.
This file indicates all the domains and resources the fee oracle needs to fetch data for. Details need to be matched with
Sygma core configuration, such as `id`. For this reason, the prefered 

Sygma currently does not support bridging the native currency, such as Ether on Ethereum, Matic on Polygon, however, the `id` is constructed with zero address and its native `domainId` and is used in baseRate calculation internally.

Example:
```json
{
  "domains": [
    {
      "id": 0,
      "name": "ethereum",
      "nativeTokenDecimals": 18,
      "nativeTokenSymbol": "eth",
      "resources": [
        {
          "resourceId": "0x0000000000000000000000000000000000000000000000000000000000000001",
          "decimals": 18,
          "symbol": "usdt"
        }
      ]
    },
    {
      "id": 1,
      "name": "polygon",
      "nativeTokenDecimals": 18,
      "nativeTokenSymbol": "matic",
      "resources": [
        {
          "resourceId": "0x0000000000000000000000000000000000000000000000000000000000000001",
          "decimals": 18,
          "symbol": "usdt"
        }
      ]
    },
    {
      "id": 2,
      "name": "moonbeam",
      "nativeTokenDecimals": 18,
      "nativeTokenSymbol": "glmr",
      "resources": [
        {
          "resourceId": "0x0000000000000000000000000000000000000000000000000000000000000001",
          "decimals": 18,
          "symbol": "usdt"
        }
      ]
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

For general help:`$ sygma-fee-oracle -h`

#### `$ sygma-fee-oracle server`
```
Start Sygma fee oracle main service

Usage:
  sygma-fee-oracle server [flags]

Flags:
  -c, --config_path string
  -d, --domain_config_path string
  -h, --help                          help for server
  -t, --key_type string               Support: secp256k1
  -k, --keyfile_path string
  -r, --resource_config_path string
```

#### <a id="keycli"></a>`$ sygma-fee-oracle key-generate`
```
Start Sygma fee oracle identity key generation

Usage:
  sygma-fee-oracle key-generate [flags]

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
```text
APP_MODE=release                                             // app mode: debug or release. app mode is used for internal testing only.
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
MOONSCAN_API_KEY=                                            // api key of moonscan
DATA_VALID_INTERVAL=3600                                     // Time of valid fee oracle response data in seconds
CONVERSION_RATE_PAIRS=eth,usdt,matic,usdt                    // conversion rate pairs that enabled for price fetching. Must be paired
```

# API Documentation
[Swagger API Doc](https://app.swaggerhub.com/apis-docs/cb-fee-oracle/fee-oracle/1.0.0)

# License
_Business Source License 1.1_
