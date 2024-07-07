### README.md

# NATS Practice Project

This project demonstrates the use of NATS and its JetStream capabilities along with Go routines, channels, and NATS Pub-Sub, Queue Subscribe, and Request-Reply patterns.

## Table of Contents

- [Project Structure](#project-structure)
- [Requirements](#requirements)
- [Setup](#setup)
- [Running the Project](#running-the-project)
- [Code Overview](#code-overview)
  - [main.go](#maingo)
  - [goroutines.go](#goroutinesgo)
  - [nats_jetstream.go](#nats_jetstreamgo)
  - [nats_pub_sub.go](#nats_pub_subgo)
  - [nats_queue_subscribe.go](#nats_queue_subscribego)
  - [nats_request_reply.go](#nats_request_replygo)
- [Docker Compose](#docker-compose)
- [License](#license)

## Project Structure

```
.Nats_Practice_Project
├── docker-compose.yml
├── main.go
├── go.sum
├── go.mod
├── ReadMe
├── goroutines
│   └── goroutines.go
├── nats_basic
│   ├── nats_jetstream.go
│   ├── nats_pub_sub.go
│   ├── nats_queue_subscribe.go
│   └── nats_request_reply.go
```

## Requirements

- [Go](https://golang.org/doc/install) 1.22 or higher
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Setup

1. Clone this repository:
    ```sh
    git clone https://github.com/vyacheslavmeyerzon/nats_practice_project
    cd nats_practice_project
    ```

2. Install Go dependencies (if any):
    ```sh
    go mod tidy
    ```

## Running the Project

### Using Docker Compose

1. Ensure Docker is running.

2. Start the services using Docker Compose:
    ```sh
    docker-compose up
    ```

### Running the Go application

1. Ensure NATS server is running.

2. Run the Go application:
    ```sh
    go run main.go
    ```

## Code Overview

### main.go

This is the entry point of the application, which runs various examples related to goroutines and NATS functionalities.

### goroutines.go

- **LaunchGoroutines**: Launches 10 goroutines, each printing numbers from 1 to 10.
- **ChannelExample**: Demonstrates sending and receiving data via a channel.
- **BufferedChannelExample**: Demonstrates the usage of buffered channels.
- **SelectExample**: Demonstrates using the select statement with multiple channels.

### nats_jetstream.go

Demonstrates setting up and using NATS JetStream for:
- Stream creation and message publishing
- Object store operations
- Key-value store operations
- Configuring consumers with filtering, ack wait, and max delivery settings

### nats_pub_sub.go

Provides a basic example of the Pub-Sub pattern with NATS.

### nats_queue_subscribe.go

Demonstrates setting up queue subscribers to distribute tasks among workers.

### nats_request_reply.go

Illustrates the request-reply pattern with NATS, where a subscriber responds to requests.

## Docker Compose

The `docker-compose.yml` file defines a NATS service with JetStream enabled. It includes volume and port configurations.

```yaml
version: '3'

services:
  nats:
    image: nats:latest
    ports:
      - "4222:4222"
      - "8222:8222"
    command: ["-js"]
    volumes:
      - nats-data:/data
    restart: unless-stopped

volumes:
  nats-data:
```

## License

MIT License

```
MIT License

Copyright (c) 2024 - Meyerzon Vyacheslav

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```