# AWS ARN Parser

A simple Go HTTP handler that parses AWS ARNs (Amazon Resource Names) and extracts their components including region information.

## App

https://aws-arn-parser.vercel.app/

## Usage

### As a Serverless Function (Vercel)

Deploy the `api/parse-arn.go` handler directly to Vercel or similar platforms.

### As a Standalone Server

```bash
go run main.go
```

Then make requests to:

```bash
http://localhost:8080/parse-arn?arn=arn:aws:s3:::my-bucket/folder/file.txt
```

### API Endpoint

**GET** `/parse-arn?arn={ARN}`

**Parameters:**

- `arn` (required): The AWS ARN to parse

**Example Request:**

```http
GET /parse-arn?arn=arn:aws:lambda:us-east-1:123456789012:function:my-function
```

**Example Response:**

```json
{
  "partition": "aws",
  "service": "lambda",
  "region": "us-east-1",
  "accountId": "123456789012",
  "resource": "function:my-function"
}
```

**Error Response:**

```json
{
  "error": "invalid ARN format"
}
```
