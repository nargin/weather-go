# Dofus API Scanner

A high-performance concurrent API scanner for the Dofus API with built-in rate limiting and retry logic.

## Features

- **Concurrent Workers**: Multiple goroutines working in parallel
- **Rate Limiting**: Configurable requests per second to avoid being blocked
- **Retry with Exponential Backoff**: Automatically retries failed requests with increasing delays
- **Progress Tracking**: Real-time feedback on scan progress
- **Result Storage**: Saves both all results and successful IDs to JSON files

## Installation

```bash
cd dofus-scanner
go mod download
go build
```

## Usage

### Basic Usage

```bash
./dofus-scanner
```

### Advanced Configuration

```bash
./dofus-scanner -start 1 -end 5000 -workers 20 -rps 15
```

### Command-Line Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-start` | 1 | Starting ID to scan |
| `-end` | 1000 | Ending ID to scan |
| `-workers` | 10 | Number of concurrent workers |
| `-rps` | 10 | Maximum requests per second |
| `-retries` | 3 | Maximum retry attempts for failed requests |
| `-timeout` | 10 | Request timeout in seconds |
| `-url` | https://dofocus.fr/api/items | Base API URL |

## Examples

Scan IDs 1-10000 with 20 workers at 15 requests/second:
```bash
./dofus-scanner -start 1 -end 10000 -workers 20 -rps 15
```

Scan with more aggressive retry logic:
```bash
./dofus-scanner -retries 5 -timeout 15
```

## Output

The scanner generates two JSON files:

1. **all_results.json**: Contains all scan results with status codes and response data
2. **successful_ids.json**: Simple array of IDs that returned valid responses

## Performance Tuning

- **Workers**: More workers = faster scanning, but may trigger rate limits
- **RPS (Requests Per Second)**: Start conservative (5-10) and increase if not rate-limited
- **Retries**: Higher retry count helps with intermittent failures but slows overall scan

## How It Works

1. Creates a worker pool with specified number of goroutines
2. Each worker pulls IDs from a job queue
3. Rate limiter ensures requests stay within specified RPS
4. Failed requests are retried with exponential backoff (500ms → 1s → 2s → 4s...)
5. 429 (rate limit) responses trigger longer backoff
6. Results are collected and saved to JSON files

## Avoiding Rate Limits

The scanner has multiple strategies to avoid rate limits:

- Token bucket rate limiter (golang.org/x/time/rate)
- Exponential backoff on failures
- Special handling for 429 status codes
- Configurable request timeout

Start with conservative settings and increase gradually:
```bash
# Start conservative
./dofus-scanner -workers 5 -rps 5

# If no rate limiting, increase gradually
./dofus-scanner -workers 10 -rps 10
./dofus-scanner -workers 20 -rps 15
```
