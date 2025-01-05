# Indonesia Cities & Villages API

## Overview
RESTful API service providing information about Indonesian cities and villages, built with Go Fiber.

## Prerequisites
- Go 1.16+
- Git
- Text editor (VS Code recommended)
- Terminal/Command Prompt

## Project Structure
```
kota-api
|   go.mod
|   go.sum
|   main.go
|   readme.md
|
+---controller
|       controller.go
|
+---data
|       desakel.csv
|       kabkota.csv
|
+---middleware
|       recovery.go
|
\---routes
        routes.go
```

## Installation Steps

1. Clone repository
```bash
git clone https://github.com/rrraf1/city-api.git
cd kota-api
```

2. Install dependencies
```bash
go mod init your-url
go mod tidy
```

3. Prepare data files
- Create 

data

 folder in root directory
- Add required CSV files:
  - 

kabkota.csv

 (Cities data)
  - 

desakel.csv

 (Villages data)

4. Configure environment
```bash
echo PORT=3000 > .env
```

5. Run application
```bash
go run main.go
```


## API Documentation

### Search Cities/Villages
- **Endpoint**: GET `/news/:id`
- **Rate Limit**: 15 requests/minute
- **Parameters**: 
  - `id`: Search term (URL encoded)

### Example Requests

1. Simple search:
```bash
curl http://localhost:3000/news/bandung
```

2. Search with spaces:
```bash
curl http://localhost:3000/news/margahayu%20utara
```

### Response Format
```json
{
  "message": "News detail found",
  "data": [
    {
      "kota": "City Name",
      "type": "Kabupaten/Kota",
      "villages": [
        {
          "ID": "village_id",
          "Nama": "Village Name",
          "Zip": "postal_code"
        }
      ]
    }
  ]
}
```

## Error Handling

### Common HTTP Status Codes
- 200: Success
- 400: Bad Request
- 404: Not Found
- 429: Too Many Requests
- 500: Internal Server Error