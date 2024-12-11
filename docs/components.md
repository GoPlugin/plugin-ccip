# Plugin Components

* shared types and interfaces - [plugin-common, ccipocr3](https://github.com/goplugin/plugin-common/tree/main/pkg/types/ccipocr3)
* [OCR Plugins](https://github.com/goplugin/plugin-ccip)
  * [commit](https://github.com/goplugin/plugin-ccip/tree/main/commit)
  * [execute](https://github.com/goplugin/plugin-ccip/tree/main/execute)
  * [CCIPReader](https://github.com/goplugin/plugin-ccip/blob/main/pkg/reader/ccip.go) - contract reader wrapper interface for core protocol data access.
  * [Home Chain Reader](https://github.com/goplugin/plugin-ccip/blob/main/pkg/reader/home_chain.go) - contract reader wrapper for home chain data access.
* core node integration ([CCIP Capability](https://github.com/goplugin/pluginv3.0/tree/develop/core/capabilities/ccip))
  * EVM
    * [providers (hashing, encoding, etc)](https://github.com/goplugin/pluginv3.0/tree/develop/core/capabilities/ccip/ccipevm)
    * [contract reader & writer configuration](https://github.com/goplugin/pluginv3.0/tree/develop/core/capabilities/ccip/configs/evm)
* integration tests
    * [initial deploy test](https://github.com/goplugin/pluginv3.0/blob/develop/integration-tests/deployment/ccip/changeset/initial_deploy_test.go)
