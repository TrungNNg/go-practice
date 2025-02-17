# OG Tag URL Preview Service

## Overview
This project is an API that extracts Open Graph (OG) metadata from a given URL and returns it as JSON. The service is designed with performance, reliability, and monitoring in mind.

## Features
- Fetches OG metadata from web pages
- Returns structured JSON response
- Caches popular URLs using Redis for faster retrieval
- Implements failure handling and circuit breakers to improve resilience
- Includes test coverage for reliability
- Tracks metrics for monitoring API usage

## Installation
### Prerequisites
- Go 1.20+
- Redis (optional, for caching)

### Steps
1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/og-preview-service.git
   cd og-preview-service
   ```
2. Initialize the Go module:
   ```sh
   go mod tidy
   ```
3. Run the service:
   ```sh
   go run main.go
   ```

## API Usage
### Endpoint
```
GET /api/preview?url={target_url}
```

### Example Request
```
GET /api/preview?url=https://example.com
```

### Example Response
```json
{
  "title": "Example Title",
  "description": "Example description of the page.",
  "image": "https://example.com/image.jpg",
  "url": "https://example.com"
}
```

## Configuration
Environment variables:
| Variable       | Default           | Description                          |
|---------------|-------------------|--------------------------------------|
| PORT          | 8080              | Port number for the API server      |
| REDIS_ADDR    | localhost:6379    | Redis server address                 |
| CACHE_TTL     | 3600              | Time-to-live for cached entries (s)  |

## Testing
Run unit tests using:
```sh
  go test ./...
```

## Future Improvements
- Support for additional metadata extraction
- Rate limiting to prevent abuse
- More detailed logging and monitoring

## License
This project is licensed under the MIT License.

