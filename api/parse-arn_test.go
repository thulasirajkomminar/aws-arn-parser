package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name           string
		arn            string
		expectedStatus int
		expectedRegion string
		expectedError  bool
	}{
		{
			name:           "Valid S3 ARN",
			arn:            "arn:aws:s3:::my-bucket/folder/file.txt",
			expectedStatus: http.StatusOK,
			expectedRegion: "",
			expectedError:  false,
		},
		{
			name:           "Valid EC2 ARN with region",
			arn:            "arn:aws:ec2:us-west-2:123456789012:instance/i-1234567890abcdef0",
			expectedStatus: http.StatusOK,
			expectedRegion: "us-west-2",
			expectedError:  false,
		},
		{
			name:           "Invalid ARN format",
			arn:            "invalid-arn",
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name:           "Missing ARN parameter",
			arn:            "",
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name:           "ARN with wrong prefix",
			arn:            "notarn:aws:s3:::my-bucket",
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/?arn="+tt.arn, nil)
			w := httptest.NewRecorder()

			Handler(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if !tt.expectedError && w.Code == http.StatusOK {
				var parsed ParsedARN
				if err := json.Unmarshal(w.Body.Bytes(), &parsed); err != nil {
					t.Errorf("failed to parse JSON response: %v", err)
				}

				if parsed.Region != tt.expectedRegion {
					t.Errorf("expected region '%s', got '%s'", tt.expectedRegion, parsed.Region)
				}
			}
		})
	}
}

func TestParseARNComponents(t *testing.T) {
	req := httptest.NewRequest("GET", "/?arn=arn:aws:lambda:us-east-1:123456789012:function:my-function", nil)
	w := httptest.NewRecorder()

	Handler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var parsed ParsedARN
	if err := json.Unmarshal(w.Body.Bytes(), &parsed); err != nil {
		t.Fatalf("failed to parse JSON response: %v", err)
	}

	expected := ParsedARN{
		Partition: "aws",
		Service:   "lambda",
		Region:    "us-east-1",
		AccountID: "123456789012",
		Resource:  "function:my-function",
	}

	if parsed != expected {
		t.Errorf("parsed ARN doesn't match expected. Got: %+v, Expected: %+v", parsed, expected)
	}
}
