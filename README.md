# golang-interview-kit

A Go CLI tool for interacting with the Measure API. This tool provides commands for measuring and analyzing data through a simple command-line interface.

This CLI uses the Scope3 Measure API, you can find the documentation here: https://docs.scope3.com/reference/measure-1

## Setup

1. Clone the repository:
```bash
git clone https://github.com/scope3data/golang-interview-kit
cd golang-interview-kit
```

2. Copy the environment template and set up your environment variables:
```bash
cp .env.template .env
```

3. Open the `.env` file and set your API key:
```bash
API_KEY="your-api-key-here"
LOG_LEVEL="debug"  # Optional: Change log level if needed (default: debug)
```

## Building and Running

There are two ways to run the CLI:

### 1. Using `go run`

Run the CLI directly using Go:

```bash
go run main.go [command]
```

### 2. Building and Running the Binary

Build the binary:

```bash
go build -o measure-cli
```

Run the compiled binary:

```bash
./measure-cli [command]
```

## Available Commands

### probe

The `probe` command tests that your API connection is working properly and verifies your API key is set correctly:

```bash
go run main.go probe
# or if you built the binary:
./measure-cli probe
```

If successful, you'll see a response from the API. If there's an error, check that your API_KEY is set correctly in the `.env` file.

## Dependencies

The project uses Go modules for dependency management. The main dependencies will be automatically installed when you run `go build` or `go run`. You can also explicitly download dependencies with:

```bash
go mod download
```
