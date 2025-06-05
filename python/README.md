# Python Measure CLI

This is a Python implementation of the Scope3 Measure CLI tool using Click. It provides functionality to interact with the Scope3 API and parse measurement data.

## Prerequisites

- Python 3.8 or higher
- pip (Python package installer)
- virtualenv or venv (Python virtual environment manager)

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/scope3data/interview-kit
cd interview-kit/python
```

### 2. Create a Virtual Environment

First, create and activate a virtual environment to isolate the project dependencies:

```bash
# Create a virtual environment
python -m venv venv

# Activate the virtual environment
# On Windows:
venv\Scripts\activate
# On macOS/Linux:
source venv/bin/activate
```

### 3. Install Dependencies

With the virtual environment activated, install the required packages:

```bash
pip install -r requirements.txt
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

The CLI provides several commands to interact with the Scope3 API:

### Test API Connection

Test that the API is reachable and your API key is properly configured:

```bash
python cli.py probe
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

## Troubleshooting

1. If you see import errors, ensure:
   - Your virtual environment is activated
   - You're running the commands from the project root directory
   - All dependencies are installed

2. If you get API errors:
   - Verify your API key is set correctly
   - Check your internet connection
   - Ensure the API endpoint is reachabl
