# mqtt_client

## Install
```
    go get github.com/tinobt/mqtt_client
```

## Import
```go
    improt mqtt_client "github.com/tinobt/mqtt_client"
```

## USE

## Step1
Create config file config/example.json

```
{
    "mqtt_client": {
        "password": "your_password",
        "username": "your_username",
        "topic": "your_topic"
    },
    "mqtt_broker": {
        "connect_type": "tcp|ws",
        "port": "your_port_mqtt_broker",
        "host": "your_host_mqtt_broker"
    }
}

```
