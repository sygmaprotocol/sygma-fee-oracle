# Fee estimate message

Different properties of a fee estimate message may have varying meanings based on their context. A description of each
context is provided below.

### Message properties

<details>
<summary>ber</summary>
<br>

Base Effective Rate (BER) is the rate between the base currencies of source and destination networks (domains).

> Example: If we are requesting a fee estimate to bridge from Polygon to Ethereum, BER would indicate the current ratio
> of MATIC/ETH

*This is a floating-point number represented by `uint256` and has a fixed precision of 18.*

</details>

<details>
<summary>ter</summary>
<br>

Token Effective Rate is the rate between the provided token (specified as the resourceID) and the base currency on the
destination network (domain).

> Example: If we are requesting a fee estimate to pay the fee in UNI tokens when bridging from Polygon to Ethereum, TER
> would indicate the current ratio of UNI/ETH

*This is a floating-point number represented by `uint256` and has a fixed precision of 18.*

</details>

<details>
<summary>expiresAt</summary>
<br>

This property is a Unix timestamp that specifies the point in time until which the fee estimate is considered valid.

</details>

<details>
<summary>fromDomainID / toDomainID</summary>
<br>

These properties define source and destination networks (domains) for the issued fee estimate.

</details>

<details>
<summary>resourceID</summary>
<br>

This property defines resource being bridged for the issued fee estimate.

</details>

<details>
<summary>msgGasLimit</summary>
<br>

This property defines the maximum amount of gas units that should be spent on execution.

> msgGasLimit parameter is only being used for generic message requests

</details>

<details>
<summary>sig</summary>
<br>

Signature of entire fee estimate message - signed by Fee Oracle identity key.

</details>

### EVM context

When the destination network (domain) is an EVM-based chain, the following properties apply:

```
ber * 10^18:  uint256                   //
ter * 10^18:  uint256                   // 
dstGasPrice:  uint256                   // ** 
expiresAt:    uint256                   // 
fromDomainID: uint8 encoded as uint256  // 
toDomainID:   uint8 encoded as uint256  // 
resourceID:   bytes32                   // 
msgGasLimit:  uint256                   // 
sig:          bytes(65 bytes)           // 
```

#### dstGasPrice

In this context, the `dstGasPrice` property is interpreted as the gas price on the destination network (chain).

> Example: On Ethereum, this property represents the number of wei per unit of gas that needs to be paid for the
> transaction to likely be executed on the destination chain.

### Substrate context

When the destination network (domain) is an Substrate-based chain, the following properties apply:

```
ber * 10^18:  uint256                   //
ter * 10^18:  uint256                   //
inclusionFee: uint256                   // **
expiresAt:    uint256                   // 
fromDomainID: uint8 encoded as uint256  // 
toDomainID:   uint8 encoded as uint256  //
resourceID:   bytes32                   // 
msgGasLimit:  uint256                   // 
sig:          bytes(65 bytes)           // 
```

#### inclusionFee

In this context, the `inclusionFee` property is interpreted as an fee that is required for extrinsic to be
invoked.

For more information on what inclusion fee is in context of Substrate check out
their [documentation](https://docs.substrate.io/build/tx-weights-fees/).

### Generic messaging context

When the user is bridging a generic message, the `msgGasLimit` property is used and interpreted. In all other cases,
this property is ignored.

#### msgGasLimit

In this context, the `msgGasLimit` property is the maximum amount of gas units that should be spent on the execution.
