# Fee Oracle technical documentation

Fee Oracle service
facilitates [Sygmas Dynamic fee strategy](https://github.com/sygmaprotocol/sygma-relayer/blob/main/docs/general/Fees.md),
allowing users to request a **signed fee estimate** through its API.

To request a fee estimate, users must provide the following parameters:

- **fromDomainID** - the identifier of the source domain (chain)
- **toDomainID** - the identifier of the destination domain (chain)
- **resourceID** - the identifier of the cross-chain resource

When requesting a fee estimate for a generic message, users must also provide the following parameter:

- **msgGasLimit** - the maximum amount of gas units that can be used for execution

For more information on the FeeOracle API, please refer to
the [Swagger documentation](https://app.swaggerhub.com/apis-docs/cb-fee-oracle/fee-oracle).

## Fee estimate message

Fee estimate message is signed fee estimate provided by fee oracle for specified bridging route. This message needs to
be provided on executing deposit on source chain. This message is required to execute a deposit on the source chain. To
learn more about the format of the message and its
different contexts, please refer to the [technical documentation on fee estimate messages](/docs/Message.md).
