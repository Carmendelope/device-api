# device-api

This component is in charge of receiving and processing device requests such as registration, login, etc.

Notice that the device API supports REST and gRPC request by means of the [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) that is launched when the server starts.

## Getting Started

### Prerequisites

* [device-manager](https://github.com/nalej/device-manager)
* [authx](https://github.com/nalej/authx)
* [application-manager](https://github.com/nalej/application-manager)

### Build and compile

In order to build and compile this repository use the provided Makefile:

```shell script
make all
```

This operation generates the binaries for this repo, download dependencies,
run existing tests and generate ready-to-deploy Kubernetes files.

### Run tests

Tests are executed using Ginkgo. To run all the available tests:

```shell script
make test
```

### Update dependencies

Dependencies are managed using Godep. For an automatic dependencies download use:

```shell script
make dep
```

In order to have all dependencies up-to-date run:

```shell script
dep ensure -update -v
```

## Known Issues

## Contributing

Please read [contributing.md](contributing.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/nalej/device-api/tags). 

## Authors

See also the list of [contributors](https://github.com/nalej/device-api/contributors) who participated in this project.

## License
This project is licensed under the Apache 2.0 License - see the [LICENSE-2.0.txt](LICENSE-2.0.txt) file for details.