microq uses a multi stage docker build for hermetic builds, while creating a minimal image. Hence, please ensure you use
Docker v17.05 or newer.

```shell
git clone https://github.com/c16a/microq.git
cd microq
docker build -t microq .
```

### Running the image

```shell
docker run -p 8080:8080 -v $pwd/config.json:/app/config.json microq
```

The above example assumes that the TCP server has been configured to listen on port 8080. In case that is configured to
another port, please configure the docker exposed port accordingly.

#### SELinux policies

When using Docker on a host with SELinux enabled, the container is denied access to certain parts of host file system
unless it is run in privileged mode. To resolve this, you can use a named volume

```shell
# Create a docker volume and map it to /tmp/microq on the host
docker volume create --driver local --opt type=none --opt device=/tmp/microq --opt o=bind microq_volume

# Ensure /tmp/microq/config.json has the required broker configuration
# Use the above created microq_volume to mount the config file into the container
docker run -p 8080:8080 -e CONFIG_FILE_PATH=/tmp/microq/config.json --mount source=microq_volume,target=/tmp/microq microq
```

Please note that however, you place your `config.json` in the `/tmp` directory, SELinux does not restrict you access
when you use a direct volume mapping.

```shell
# This won't work with SELinux enabled
docker run -p 8080:8080 -e CONFIG_FILE_PATH=/tmp/microq/config.json -v /home/user/config.json:/tmp/microq/config.json microq

# This will work
docker run -p 8080:8080 -e CONFIG_FILE_PATH=/tmp/microq/config.json -v /tmp/microq/config.json:/tmp/microq/config.json microq
```

The [Configuration](configuration.md) section has more details on which attributes of the broker can be configured.

### Running in Compose mode

Create the named volume `microq_volume`.

```shell
# Create a docker volume and map it to /tmp/microq on the host
docker volume create --driver local --opt type=none --opt device=/tmp/microq --opt o=bind microq_volume
```

Reference the named volume for the service

```yaml
version: "3.9"
services:
  broker:
    build:
      context: .
    environment:
      CONFIG_FILE_PATH: "/tmp/microq/config.json"
    volumes:
      - microq_volume:/tmp/microq
    ports:
      - 8080:8080
volumes:
  microq_volume:
    external: true
```