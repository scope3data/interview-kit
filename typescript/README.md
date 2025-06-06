# TypeScript Measure CLI

This is a TypeScript implementation of the Scope3 Measure CLI.

## Prerequisites

- Node 20 or higher

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/scope3data/interview-kit
cd interview-kit/typescript
```

### 2. Install Dependencies

```bash
npm install
```

## Configuration

The CLI requires a Scope3 API key to function. Set your API key as an environment variable:

1. Copy the environment template and set up your environment variables:
```bash
cp ../.env.template ./.env
```

2. Open the `.env` file and set your API key:
```bash
API_KEY="your-api-key-here"
LOG_LEVEL="debug"  # Optional: Change log level if needed (default: debug)
```

## Usage

You can run the CLI by executing the following command:

```bash
npm run cli $command
```

### Test API Connection

Test that the API is reachable and your API key is properly configured:

```bash
npm run cli probe
```

You should see this output:
```
RequestID: b3f8c247-34f3-4bc2-bffd-e64219bd6c7e (can be anything)
Total Emissions: 0.2183
Breakdown:
  - Ad Selection: 0.1676
  - Creative Delivery: 0.0009
  - Media Distribution: 0.0498
```
