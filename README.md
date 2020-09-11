[//]: # "SPDX-License-Identifier: CC-BY-4.0"

# Hyperledger Fabric Samples

This is a highly modified version of the Hyperledger Fabric Samples repo. It
contains the configuration and scripts for a two peer network with certificate
authorities, a simple KV chaincode, and a node.js application. It is designed to
be used for doing proof of state with the [Fabric
agent](https://github.com/dlt-interoperability/commitment-agent).

It brings up a network with the following components:

- Org1 peer.
- Org1 ca.
- Org2 peer.
- Org2 ca.
- Orderer.
- Orderer ca.

## Quickstart guide

Start the Fabric network and deploy and invoke the chaincode with:

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
