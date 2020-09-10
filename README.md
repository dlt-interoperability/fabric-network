[//]: # "SPDX-License-Identifier: CC-BY-4.0"

# Hyperledger Fabric Samples

This is a highly modified version of the Hyperledger Fabric Samples repo. It
brings up a network with the following components:

- Org1 peer.
- Org1 ca.
- Org2 peer.
- Org2 ca.
- Orderer.
- Orderer ca.

It then deploys a very simple chaincode that stores a value against a key.
It also uses a javascript application to initialise and invoke the chaincode.
This Fabric network and chaincode is designed to be used in conjunction with the
[Fabric commitment
agent](https://github.com/dlt-interoperability/commitment-agent).

## Quickstart guide

```
make start
make deploy-cc
make invoke-cc
```

Note that the application is set up to make a `CreateAsset` transaction every
10 seconds, to help with the development of the Fabric commitment agent.

## Stopping

```
make stop
```

## TODO

- Change the keys added to the ledger to be random numbers so there is no key
  clash in the accumulator.

## License <a name="license"></a>

Hyperledger Project source code files are made available under the Apache
License, Version 2.0 (Apache-2.0), located in the [LICENSE](LICENSE) file.
Hyperledger Project documentation files are made available under the Creative
Commons Attribution 4.0 International License (CC-BY-4.0), available at
http://creativecommons.org/licenses/by/4.0/.
