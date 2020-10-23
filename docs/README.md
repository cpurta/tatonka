# Quick Start

## Requirements

* macOS / Linux / Windows
* Docker
* Docker-compose
* Golang (1.15+)

## Install Tatanka

Clone from github into your `$GOPATH`:

```
$ git clone https://github.com/cpurta/tatanka
```

Build the linux tatanka binary:

```
$ make tatanka
```

## Configure tatanka

There is currently a sample config in the project root called `sample-config.yaml`.
You will need to create a new `config.yaml` with your coinbase api key, secret and
passphrase.

```
$ cp ./sample-config.yaml ./config.yaml
```

Fill in the following fields with your credentials in the `config.yaml` file:

 - `api_key`
 - `api_secret`
 - `api_passphrase`

```yaml
cassandra:
  cluster:
    - "cassandra"
  keyspace: "tatanka"

gdax:
  api_key: "your_gdax_api_key"
  api_secret: "your_gdax_api_secret"
  api_passphrase: "your_gdax_api_passphrase"
```

Once you have entered your credentials tatanka should be able to successfully connect
to the Coinbase API to then pull account balance information and be able to live
trade if allowed.

## Running in Docker

Currently tatanka is configured to run in docker and a docker compose environment.
Although you can set up tatanka to run locally if will require you to run a Cassandra
cluster on your machine and point tatanka to connect to that cluster. For now it
is recommended you use the docker-compose environment.

### Build

You can build the necessary docker images by running the following command:

```
$ docker-compose -f ./docker/docker-compose.yml build
```

Or alternately you can use the `make build` command:

```
$ make build
```

### Run

Since tatanka is running in docker-compose you can run or execute tatanka commands
by running those through docker-compose commands.

#### Examples

**Balance:**

```
$ docker-compose -f ./docker/docker-compose.yml run tatanka balance gdax.BTC-USD
```

**List Strategies**

```
$ docker-compose -f ./docker/docker-compose.yml run tatanka list-strategies
```

**List Selectors**

```
$ docker-compose -f ./docker/docker-compose.yml run tatanka list-selectors
```

**Simulate Strategy**

```
$ docker-compose -f ./docker/docker-compose.yml run tatanka sim gdax.BTC-USD
```

**Backfill Historical Data**

```
$ docker-compose -f ./docker/docker-compose.yml run tatanka backfill gdax.BTC-USD
```

**Live/Paper Trade on Live Market Data**

```
$ docker-compose -f ./docker/docker-compose.yml run tatanka trade gdax.BTC-USD
```
