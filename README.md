# Nightscaler

## Running locally

First you need to start the Pub/Sub emulator, to do so have a look at the [.hack](./.hack/README.md) folder.

```shell
export PUBSUB_PROJECT_ID=es-standalone-cb21
export PUBSUB_EMULATOR_HOST=localhost:8085

go run .
```
