package cloud_what

import "github.com/m-horky/subman-facts/pkg/facts"

var AWS = Cloud{
	ID:                 "aws",
	MetadataURL:        "http://169.254.169.254/latest/dynamic/instance-identity/document",
	MetadataCacheFile:  "",
	MetadataType:       "application/json",
	SignatureUrl:       "http://169.254.169.254/latest/dynamic/instance-identity/rsa2048",
	SignatureType:      "text/plain",
	SignatureCacheFile: "",
	TokenURL:           "http://169.254.169.254/latest/api/token",
	TokenCacheFile:     "/var/cache/cloud-what/aws_token.json",
	TokenTTL:           3600,
	CustomHTTPHeaders:  map[string]string{"User-Agent": "cloud-what/2.0"},
	MemoryCacheTTL:     0,
	ServerTimeout:      0,
}

func runsOnCloud_AWS(c Cloud, f facts.CollectedFacts) bool {
	return false
}
