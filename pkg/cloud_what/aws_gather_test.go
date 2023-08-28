package cloud_what

import (
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func TestAWSInstance_GetTokenFromServer(t *testing.T) {
	tests := []struct {
		description string
		call        func(client *http.Client, request *http.Request) (*http.Response, error)
		wants       awsToken
		wantsErr    error
	}{
		{
			description: "Valid run",
			call: func(client *http.Client, request *http.Request) (*http.Response, error) {
				return &http.Response{}, nil
			},
			wants:    awsToken{},
			wantsErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			IdentityService.Call = test.call

			instance := NewAWSInstance()
			token, err := instance.getTokenFromServer()

			if test.wantsErr == nil {
				// We expect valid results
				if !cmp.Equal(test.wantsErr, err) {
					t.Errorf("CachedData expected as %s, got %s", test.wants.cachedData, token.cachedData)
				}
			} else {
				// We expect failed results
				if !cmp.Equal(test.wantsErr, err) {
					t.Errorf("Error expected as %s, got %s", test.wantsErr, err)
				}
			}
		})
	}
}
