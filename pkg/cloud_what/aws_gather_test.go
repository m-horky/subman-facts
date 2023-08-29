package cloud_what

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAWSInstance_GetTokenFromServer(t *testing.T) {
	tests := []struct {
		description         string
		mockResponseStatus  int
		mockResponseContent string
		wants               string
		wantsErr            error
	}{
		{
			description:         "Valid run",
			mockResponseStatus:  http.StatusOK,
			mockResponseContent: "testtoken",
			wants:               "testtoken",
			wantsErr:            nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(test.mockResponseStatus)
				_, err := io.WriteString(w, test.mockResponseContent)
				if err != nil {
					t.Errorf("Mock server could not write a response: %s", err)
				}
			}))
			defer server.Close()

			instance := NewAWSInstance()
			instance.Configuration.TokenURL = server.URL
			token, err := instance.getTokenFromServer()

			if test.wantsErr == nil {
				// We expect valid results
				if !cmp.Equal(test.wants, token) {
					t.Errorf("Token expected as %s, got %s", test.wants, token)
				}
			} else {
				// We expect failed results
				if !cmp.Equal(test.wantsErr.Error(), err.Error()) {
					t.Errorf("Error expected as %s, got %s", test.wantsErr, err)
				}
			}
		})
	}
}

func TestAWSInstance_GetMetadata(t *testing.T) {
	tests := []struct {
		description         string
		mockResponseStatus  int
		mockResponseContent string
		wants               AWSMetadata
		wantsErr            error
	}{
		{
			description:         "eu-central-1",
			mockResponseStatus:  http.StatusOK,
			mockResponseContent: readTestFile("aws-metadata-eu.json"),
			wants: AWSMetadata{
				AccountId:        "012345678900",
				Architecture:     "x86_64",
				AvailabilityZone: "eu-central-1b",
				BillingProducts:  []string{"bp-0124abcd", "bp-63a5400a"},
				ImageId:          "ami-0123456789abcdeff",
				InstanceId:       "i-abcdef01234567890",
				InstanceType:     "m5.large",
				PendingTime:      "2020-02-02T02:02:02Z",
				PrivateIp:        "12.34.56.78",
				Region:           "eu-central-1",
				Version:          "2017-09-30",
			},
			wantsErr: nil,
		},
		{
			description:         "bad token",
			mockResponseStatus:  http.StatusUnauthorized,
			mockResponseContent: "",
			wants:               AWSMetadata{},
			wantsErr:            errors.New("server responded with status code 401"),
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(test.mockResponseStatus)
				_, err := io.WriteString(w, test.mockResponseContent)
				if err != nil {
					t.Errorf("Mock server could not write a response: %s", err)
				}
			}))
			defer server.Close()

			instance := NewAWSInstance()
			instance.Configuration.MetadataURL = server.URL
			metadata, err := instance.GetMetadata()

			if test.wantsErr == nil {
				// We expect valid results
				if !cmp.Equal(test.wants, metadata) {
					t.Errorf("Metadata expected as %#v, got %#v", test.wants, metadata)
				}
			} else {
				// We expect failed results
				if !cmp.Equal(test.wantsErr.Error(), err.Error()) {
					t.Errorf("Error expected as %s, got %s", test.wantsErr, err)
				}
			}
		})
	}
}
