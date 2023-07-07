package cloud_what

import (
	"git.sr.ht/~spc/go-log"
	"github.com/m-horky/subman-facts/pkg/facts"
)

type Cloud struct {
	// Unique cloud identifier: 'aws', 'azure', 'gcp'
	ID string `json:"id"`
	// Server providing metadata: 'http://192.0.2.1/path/to/document'
	MetadataURL string `json:"metadata_url"`
	// Type of metadata: 'application/json', 'text/xml'
	MetadataType string `json:"metadata_type"`
	// Filesystem path: '/var/lib/cloud-what/cache/mycloud-metadata.json'
	MetadataCacheFile string `json:"metadata_cache_file"`
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
	TokenTTL int `json:"token_ttl"`
	// Custom HTTP headers, like an User Agent
	CustomHTTPHeaders map[string]string `json:"custom_http_headers"`
	// Memory cache validity, in seconds
	MemoryCacheTTL int `json:"memory_cache_ttl"`
	// Connection timeout, in seconds
	ServerTimeout int `json:"server_timeout"`
}

// RunsOnCloud inspects the instance of the cloud object using collected
// system facts and decides whether we are currently running on one of the
// supported clouds.
func RunsOnCloud(c Cloud, f facts.CollectedFacts) bool {
	switch c.ID {
	case "aws":
		return runsOnCloud_AWS(c, f)
	default:
		log.Errorf("RunsOnCloud is not implemented for %s.", c.ID)
		return false
	}
}
