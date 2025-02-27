# Pub-Sub Implementation in Go

This project is a simple implementation of a Publish-Subscribe (Pub-Sub) system in Go, featuring both direct and topic-based message brokers.

## Features

- **Direct Broker**: Subscribers receive messages sent to a specific queue.
- **Topic Broker**: Subscribers receive messages based on topics they subscribe to.
- **Concurrency Support**: Uses Go channels and goroutines for efficient messaging.
- **Graceful Shutdown**: Ensures proper closing of channels to prevent resource leaks.

## Project Structure

```
├── main.go          # Entry point of the application
├── brokers/
│   ├── direct.go    # Direct broker implementation
│   ├── topics.go    # Topic broker implementation
├── go.mod           # Go module file
├── go.sum           # Dependency file
└── README.md        # Project documentation
```

## Installation

Ensure you have [Go installed](https://golang.org/doc/install) on your system.

1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/pub-sub.git
   cd pub-sub
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Usage

### Running the Direct Broker Example

```sh
 go run main.go
```

By default, the `directBrokerExample()` function is executed.

### Running the Topic Broker Example

Modify `main.go` to call `topicBrokkerExample()` instead of `directBrokerExample()` and run:

```sh
 go run main.go
```

## How It Works

### Direct Broker

1. Subscribers register for specific queues.
2. The publisher sends messages to queues.
3. Messages are received only by subscribers of that queue.

### Topic Broker

1. Subscribers register to topics with unique identifiers.
2. Publishers send messages to topics.
3. All subscribers of a topic receive the messages.

## Example Output

**Direct Broker Example:**

```
New song of Lynkin Park is just released
New song of Justin Biber is just released
Bangladesh won the match
```

**Topic Broker Example:**

```
Listener 1: Going to release new music
Listener 2: Going to release new music
Listener 1: Released new music
Listener 2: Released new music
```

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

## License

This project is open-source and available under the [MIT License](LICENSE).
