# blackbird
Template for Chainlink External Adapters written in Go. Get an adapter deployed in under 10 minutes.

## Introduction

Blackbird is intended as a lightweight template for building [Chainlink External Adapters](https://docs.chain.link/docs/external-adapters/) (EA) in Go. It uses Gin as the underlying HTTP framework to handle the requests and provides a set of helpful convenience features for deploying the EA into a cloud infrastructure.

Features included:
* /health response URL for health and uptime checkers
* /metrics response URL to provide metrics and statistics to Prometheus
* response format compatible with Chainlink node bridges
* JSON-logging to STDOUT for production container environments
* Docker support - ship the External Adapter as a small 20Mb Docker container image.

Developers should customise the contents of [handler.go](handler.go) which contains the guts of the external adapter. The out-of-the-box example shows a ETHUSD pricefeed being accessed, JSON data being parsed, reformatted and returned.

## Getting started

1. Clone the Blackbird repo and download the repo.
2. Install the go-dependencies

        go mod tidy

3. Run the server locally

        bin/run.sh

You should be able to access the following URLs:
- http://localhost:8080 - Main endpoint
- http://localhost:8080/health - Simple health=OK endpoint
- http://localhost:8080/metrics - Prometheus-based metrics endpoint

## Building & Shipping

Please ensure you're using Go version 1.18.

To dockerize and upload the Docker container image run the build script:

    bin/build.sh 0.0.1

The above command will generate a Linux/AMD64-compatible Docker container image and upload it to Docker.io's Hub (it assumes you're logged in with `docker login`). The `0.0.1` parameter above is used to tag the container image.

NOTE: You should customise the `translucentlink/blackbird` references in the `build.sh` to make them work with your Docker credentials, e.g. replace them with `my-acme-corp/price-feed`.

## Deployment

How you deploy the container is up to you (Docker, AWS, Kubernetes, etc.) but nothing is quite as fast & convenient as using [Fly.io](https://fly.io/)

    flyctl launch --image=translucentlink/blackbird:0.0.1

If you haven't got the `flyctl` command installed, check out their [2-minute intro](https://fly.io/docs/getting-started/installing-flyctl/) on installing and logging in.

To access your deployed external adapter

    flyctl open

To deploy an update to your external adapter

    bin/build.sh 0.0.2
    flyctl deploy --image=translucentlink/blackbird:0.0.2

The example container is deployed at https://ethusd-example.fly.dev/.

## Go-based External Adapters

Below is a list of Chainlink External Adapters crafted using Go:
* [Stonechat](https://github.com/translucent-link/stonechat) - connecting Google Sheets to EVM blockchains

## Support & Help

Feel free open a [Github Issue](https://github.com/translucent-link/blackbird/issues) or come find us in the [Translucent Discord](https://discord.gg/JxKT6R9Xpz).
