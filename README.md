# Covidify a very stupid and insecure restaurant registration service
This app is meant as a end-to-end example on DC/OS using multiple stateful packages, edgelb, marathon and metronome.

Its currently in an early development state and is not yet ready to use.


## Goals
Main goal of this project is to simulate an end-user app which receives data and shares and stores the data in multiple layers of states ( Cassandra, Kafka ). Also Edgelb should redirect the traffic path based to multiple instances of the app to separate the load of the application.

## User simulation
To simulate the user behaviour a visit generator based on locust and boom is generating requests.

## Monitoring
The app and the generator should expose their metrics via Prometheus Endpoint config so dcos-monitoring will be able to receive the metrics


## API
See [API Docs](./doc)



## DC/OS

Run a 3 node cassandra cluster


Get node0 task ID

```
dcos task list testing.covidify.cassandra__node-0 --json | jq -r '.[0].id'
```

Get shell on node0

```
dcos task exec --tty --interactive testing.covidify.cassandra__node-2-server__40d24407-7d53-4acc-bc6c-82d70cda14a8 "/bin/sh"
```

Exec cql

```
dcos cassandra --name "testing/covidify/cassandra" endpoints native-client | jq -r '.dns | map(split(":")[0])[0]'
```


```
apache-cassandra-*/bin/cqlsh <nodelist>
```


Now place the content of [model.csql](./model.csql)


### Marathon app
Get the list of DNS names:

```
dcos cassandra --name "testing/covidify/cassandra" endpoints native-client | jq -r '.dns | map(split(":")[0]) | join(",")'
```

```json
{
  "env": {
    "COVIDIFY_CASSANDRA": "node-0-server.testingcovidifycassandra.autoip.dcos.thisdcos.directory,node-1-server.testingcovidifycassandra.autoip.dcos.thisdcos.directory,node-2-server.testingcovidifycassandra.autoip.dcos.thisdcos.directory"
  },
  "id": "/testing/covidify/api",
  "instances": 1,
  "portDefinitions": [],
  "container": {
    "type": "MESOS",
    "volumes": [],
    "docker": {
      "image": "fatz/covidify",
      "forcePullImage": false,
      "parameters": []
    }
  },
  "cpus": 0.1,
  "mem": 256,
  "requirePorts": false,
  "networks": [],
  "healthChecks": [],
  "fetch": [],
  "constraints": []
}
```

### EdgeLB pool

`dcos edgelb create pool.json`

```json
{
  "apiVersion": "V2",
  "name": "http",
  "count": 1,
  "haproxy": {
    "frontends": [{
      "bindPort": 80,
      "protocol": "HTTP",
      "linkBackend": {
        "map": [{
          "hostEq": "covidify.dcos.d2iq.com",
          "backend": "covidifyapi"
        }]
      }
    }],
    "backends": [{
      "name": "covidifyapi",
      "protocol": "HTTP",
      "services": [{
        "marathon": {
          "serviceID": "/testing/covidify/api"
        },
        "endpoint": {
          "portName": "http"
        }
      }]
    }]
  }
}

```


curl -H "Host: covidify.dcos.d2iq.com" -X POST "http://<yourclusteraddress>/visit" -H  "accept: */*" -H  "Content-Type: application/json" -d "{\"id\":\"d290f1ee-6c54-4b01-90e6-d701748f0851\",\"table_number\":\"outside-1\",\"visitors\":[{\"name\":\"John Doe\",\"email\":\"john.doe@googlemail.com\",\"phone\":\"+49-30-123456789\",\"country\":\"DEU\",\"city\":\"Berlin\",\"zip_code\":\"11011\",\"street\":\"Platz der Republik 1\"}],\"risk\":{\"risk\":\"low\",\"description\":\"string\"}}"
