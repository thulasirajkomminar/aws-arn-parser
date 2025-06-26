package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
)

type ParsedARN struct {
	Partition string `json:"partition"`
	Service   string `json:"service"`
	Region    string `json:"region"`
	AccountID string `json:"account_id"`
	Resource  string `json:"resource"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	arnStr := r.URL.Query().Get("arn")
	if arnStr == "" {
		http.Error(w, `{"error":"arn parameter is required"}`, http.StatusBadRequest)
		return
	}

	parts, err := arn.Parse(arnStr)
	if err != nil {
		http.Error(w, `{"error":"invalid ARN format"}`, http.StatusBadRequest)
		return
	}

	parsed := ParsedARN{
		Partition: parts.Partition,
		Service:   parts.Service,
		Region:    parts.Region,
		AccountID: parts.AccountID,
		Resource:  parts.Resource,
	}

	json.NewEncoder(w).Encode(parsed)
}
