# nsgflowlogsbeat

Welcome to nsgflowlogsbeat, an Azure NSG Flow Logs shipper for Logstash and Elasticsearch.

## Usage

Configuration:

```
############################# {Beat} ######################################

nsgflowlogsbeat:
  # Defines how often storage account is scanned for changes
  #scan_frequency: 30s

  # Storage account name where NSG logs are stored
  storage_account_name: '<storage account name>'
  # Storage account key
  storage_account_key: '<storage account key>'
  
  # Name of the storage account container  
  #container_name: 'insights-logs-networksecuritygroupflowevent'
  
  # Checkpoints table name
  #checkpoints_table_name: 'nsgflowlogsbeatcheckpoints'

  # Storage table operations timeout in seconds
  #checkpoints_table_timeout: 15

  # Ignores NSG logs older than specified time offset
  #ignore_older: 10m
  
  # Number of workers
  #workers: 4
```

Running the beat:

```
./nsgflowlogsbeat -c nsgflowlogsbeat.yml -e
```

## Development

To set up a working development environment follow the official [beat developer guide](https://www.elastic.co/guide/en/beats/devguide/7.6/newbeat-getting-ready.html) and specifically [Setting Up Your Dev Environment](https://www.elastic.co/guide/en/beats/devguide/7.6/beats-contributing.html#setting-up-dev-environment).

Once set up, and you can successfully build the official beats repo, clone this repo to your `${GOPATH}`.

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/lstyles/nsgflowlogsbeat`

### Requirements

Nsgflowlogsbeat builds on top of the libbeat framework. Please follow the official [Contributing to beats](https://www.elastic.co/guide/en/beats/devguide/7.6/beats-contributing.html#beats-contributing) guide

### Build

To build the binary for nsgflowlogsbeat run the command below. This will generate a binary in the same directory with the name nsgflowlogsbeat.

```
mage build
```

If you have problems running mage, install it as described [here](https://github.com/magefile/mage), then add `${GOPATH}/bin` to your path.

### Run

To run {Beat} with debugging output enabled, run:

```
./nsgflowlogsbeat -c nsgflowlogsbeat.yml -e -d "*"
```


### Test

To test nsgflowlogsbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  nsgflowlogsbeat source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone nsgflowlogsbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/lstyles/nsgflowlogsbeat
git clone https://github.com/lstyles/nsgflowlogsbeat ${GOPATH}/src/github.com/lstyles/nsgflowlogsbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
