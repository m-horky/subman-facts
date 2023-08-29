package cloud_what

import (
	"git.sr.ht/~spc/go-log"
	"net/http"
)

type CloudInstance interface {
	GetID() string
	SystemRunsOn() bool
	SystemMaybeRunsOn() float64
	getToken() string
	GetMetadata() any
	GetSignature() any
}

type CloudConfiguration struct {
	// Unique cloud identifier: 'aws', 'azure', 'gcp'
	ID string `json:"id"`
	// Server providing metadata: 'http://192.0.2.1/path/to/document'
	MetadataURL string `json:"metadata_url"`
	// Type of metadata: 'application/json', 'text/xml'
	MetadataType string `json:"metadata_type"`
	// Filesystem path: '/var/lib/cloud-what/cache/mycloud-metadata.json'
	MetadataCacheFile string `json:"metadata_cache_file"`
	// Memory cache validity, in seconds
	MetadataCacheTTL uint `json:"memory_cache_ttl"`
	// Server providing metadata signature: 'http://192.0.2.1/path/to/signature'
	SignatureUrl string `json:"signature_url"`
	// Type of signature: 'application/json', 'text/xml', 'text/pem'
	SignatureType string `json:"signature_type"`
	// Filesystem path: '/var/lib/cloud-what/cache/mycloud-signature.json'
	SignatureCacheFile string `json:"signature_cache_file"`
	// Server providing the token: 'http://192.0.2.1/path/to/token'
	TokenURL string `json:"token_url"`
	// Filesystem path: '/var/lib/cloud-what/cache/mycloud-token.json'
	TokenCacheFile string `json:"token_cache_file"`
	// Token validity, in seconds
	TokenTTL uint `json:"token_ttl"`
	// Custom HTTP headers, like an User Agent
	CustomHTTPHeaders map[string]string `json:"custom_http_headers"`
	// Connection timeout, in seconds
	ServerTimeout uint `json:"server_timeout"`
}

// SystemRunsOnCloud inspects the instance of the cloud object using collected
// system facts and decides whether we are currently running on one of the
// supported cloud providers.
func SystemRunsOnCloud(c CloudInstance) bool {
	switch c.GetID() {
	case "aws":
		return c.SystemRunsOn()
	default:
		log.Errorf("RunsOnCloud is not implemented for %s.", c.GetID())
		return false
	}
}

// SystemMaybeRunsOnCloud inspects the instance of the cloud object using collected
// system facts and decides whether we are likely running on one of the
// supported cloud providers.
func SystemMaybeRunsOnCloud(c CloudInstance) float64 {
	switch c.GetID() {
	case "aws":
		return c.SystemMaybeRunsOn()
	default:
		log.Errorf("SystemMaybeRunsOnCloud is not implemented for %s.", c.GetID())
		return 0.0
	}
}

// RemoteService is a wrapper around network calls.
type RemoteService struct {
	// Call performs an HTTP request using supplied 'client' to the remove specified in 'request'.
	Call func(client *http.Client, request *http.Request) (*http.Response, error)
}

// IdentityService is an instance of RemoteService.
var IdentityService = RemoteService{Call: call}

// call performs a network request using method specified in request.
func call(client *http.Client, request *http.Request) (*http.Response, error) {
	// TODO Proxy settings can to be configured through client.Transport here
	// TODO Logging of requests and replies can be done here
	return client.Do(request)
}
