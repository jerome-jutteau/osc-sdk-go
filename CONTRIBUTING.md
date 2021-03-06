# Hacking Outscale SDK

SDK itself is generated from Outscale's [OpenAPI description](https://github.com/outscale/osc-api) in [osc](osc/) folder.

Other contributions like examples and tests are welcome!

# Generate SDK

1. have some tools ready: GNU make, git, docker
2. edit `version` file and to the latest version
4. launch sdk generation by running `make gen`
5. new sdk is now generated in `osc` folder

Under the hood:
- we get official Outscale yaml
- apply some patch if needed (.patch folder)
- run openapi-generator through docker to build osc folder

# Sending a Merge Request

Content in `osc` folder is generated at each release.
If you plan to make some change here, consider making a pull request in [openapi-generator project](https://github.com/OpenAPITools/openapi-generator/).

Otherwise:
- your merge request must be rebased on `next-release` branch
- be sure that tests still pass by running `make test`

# How to release

1. edit `version` file
2. `make gen` to update the sdk
3. `make test` and `make run-examples` and fix any issue
4. update `changelog.md` file