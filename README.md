# QR image converter Go

Just a QR image converter from url in golang.

## Run

Change Get env PORT by the port in development mode and run:

```go
go run .
```

## Endpoints

- GET  /api/codes                            For see all codes created in 24h
- GET  /api/code/id                          For see specific qr code
- POST /api/create   {"url": "url_content"}  For create a QR image
