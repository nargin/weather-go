# weather-go
For sure a weather app ☁️

## Setup

1. Copy the environment example file:
   ```bash
   cp env.example .env
   ```

2. Add your PandaScore API key to `.env`:
   ```
   PANDASCORE_API_KEY=your_api_key_here
   ```

## Launch the Fetcher

```bash
go run .
```

This will fetch Worlds 2025 matches from the PandaScore API and save them to `worlds_2025_matches.json`.
