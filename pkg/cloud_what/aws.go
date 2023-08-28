package cloud_what

type AWSInstance struct {
	ID                 string            `json:"id"`
	MetadataURL        string            `json:"metadata_url"`
	MetadataType       string            `json:"metadata_type"`
	MetadataCacheFile  string            `json:"metadata_cache_file"`
	SignatureUrl       string            `json:"signature_url"`
	SignatureType      string            `json:"signature_type"`
	SignatureCacheFile string            `json:"signature_cache_file"`
	TokenURL           string            `json:"token_url"`
	TokenCacheFile     string            `json:"token_cache_file"`
	TokenTTL           int               `json:"token_ttl"`
	CustomHTTPHeaders  map[string]string `json:"custom_http_headers"`
	MemoryCacheTTL     int               `json:"memory_cache_ttl"`
	ServerTimeout      int               `json:"server_timeout"`
	// Instance of `key` region will be redirected to `value`'s region for content
	RegionRedirects map[string]string
}

func NewAWSInstance() AWSInstance {
	return AWSInstance{
		ID:                 "aws",
		CustomHTTPHeaders:  map[string]string{"User-Agent": "cloud-what/2.0"},
		TokenURL:           "http://169.254.169.254/latest/api/token",
		TokenCacheFile:     "/var/cache/cloud-what/aws_token.json",
		TokenTTL:           3600,
		MetadataURL:        "http://169.254.169.254/latest/dynamic/instance-identity/document",
		MetadataCacheFile:  "",
		MetadataType:       "application/json",
		SignatureUrl:       "http://169.254.169.254/latest/dynamic/instance-identity/rsa2048",
		SignatureType:      "text/plain",
		SignatureCacheFile: "",
		MemoryCacheTTL:     0,
		ServerTimeout:      0,
		RegionRedirects: map[string]string{
			"us-gov-east-1": "us-east-2",
			"us-gov-west-1": "us-west-2",
		},
	}
}
