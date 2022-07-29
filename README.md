# mqtt_client

## Install
```
$ go get github.com/tinobt/mqtt_client
```

## Import
```go
improt mqtt_client "github.com/tinobt/mqtt_client"
```

## Usage

### Step1
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
### Step 2
```go
mqtt_client.MQTTInit("your_path_config", "your_file_config")
```
