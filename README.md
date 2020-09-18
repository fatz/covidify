# Covidify a very stupid and insecure restaurant registration service
This app is meant as a end-to-end example on DC/OS using multiple stateful packages, edgelb, marathon and metronome.

Its currently in an early development state and is not yet ready to use.


## Goals
Main goal of this project is to simulate an end-user app which receives data and shares and stores the data in multiple layers of states ( Cassandra, Kafka ). Also Edgelb should redirect the traffic path based to multiple instances of the app to separate the load of the application.

## User simulation
To simulate the user behaviour a visit generator based on locust and boom is generating requests.

## Monitoring
The app and the generator should expose their metrics via Prometheus Endpoint config so dcos-monitoring will be able to receive the metrics


... TBD
