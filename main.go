package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/thulasirajkomminar/aws-arn-parser/api"
)

func main() {
	http.HandleFunc("/parse-arn", handler.Handler)

	fmt.Println("Server starting on :8080")
	fmt.Println("Try: http://localhost:8080/parse-arn?arn=arn:aws:s3:::my-bucket/folder/file.txt")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
