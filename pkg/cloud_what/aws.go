package cloud_what

type AWSInstance struct {
	Configuration CloudConfiguration
	// Instance of `key` region will be redirected to `value`'s region for content
	RegionRedirects map[string]string
}

func NewAWSInstance() AWSInstance {
	return AWSInstance{
		Configuration: CloudConfiguration{
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
			MetadataCacheTTL:   0,
			ServerTimeout:      1,
		},
		RegionRedirects: map[string]string{
			"us-gov-east-1": "us-east-2",
			"us-gov-west-1": "us-west-2",
		},
	}
}
