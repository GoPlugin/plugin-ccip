<div style="text-align:center" align="center">
    <a href="https://goplugin.co" target="_blank">
        <img src="https://raw.githubusercontent.com/goplugin/pluginv3.0/develop/docs/logo-plugin-blue.svg" width="225" alt="plugin logo">
    </a>

[![License](https://img.shields.io/static/v1?label=license&message=BUSL%201.1&color=green)](https://github.com/goplugin/plugin-ccip/blob/master/LICENSE)
[![Code Documentation](https://img.shields.io/static/v1?label=code-docs&message=latest&color=blue)](docs/ccip_protocol.md)
</div>

# plugin-ccip

This repo contains [OCR3 plugins][ocr3] for CCIP. See the [documentation](docs/ccip_protocol.md) for more.

## Getting Started

### Go Version

The version of go is specified in the project's [go.mod](go.mod) file.
You can install Go from their [installation page](https://go.dev/doc/install).

### Running the Linter

We use golangci-lint as our linting tool. Run the linter like this:

```sh
make lint
```

### Running the Unit Tests

```sh
make test
```

### Generating the Mocks

We use mockery to generate mocks and they're organized in the [mockery.yaml](./.mockery.yaml) file.

```sh
make generate
```

## Development Cycle

In order to keep the `ccip-develop` branch in working condition, we need to make sure the integration test
[written in the CCIP repo](https://github.com/goplugin/ccip/blob/03ae3bbed0e6020be5fa9be26d03af21f152d7dc/core/capabilities/ccip/ccip_integration_tests/ocr3_node_test.go#L37)
will pass.

As such, part of CI will run this integration test combined with your latest pushed change.

Follow the steps below to ensure that we don't run into any unexpected breakages.

1. Create a PR on plugin-ccip with the changes you want to make.
2. CI will run the integration test in the CCIP repo after applying your changes.
3. If the integration test fails, make sure to fix it first before merging your changes into
the `ccip-develop` branch of plugin-ccip. You can do this by:
    - Creating a branch in the CCIP repo and running `go get github.com/goplugin/plugin-ccip@<your-branch-commit-sha>`.
    - Fixing the build/tests.
4. Once your ccip PR is approved, merge it.
5. Go back to your plugin-ccip PR and re-run the integration test workflow.
6. Once the integration test passes, merge your plugin-ccip PR into `ccip-develop`, however do not delete the branch on the remote.
7. Create a new PR in ccip that points to the newly merged commit in the `ccip-develop` tree and merge that.

[ocr3]: https://github.com/goplugin/plugin-libocr/blob/master/offchainreporting2plus/ocr3types/plugin.go#L108
