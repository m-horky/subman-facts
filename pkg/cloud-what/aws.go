package cloud_what

import (
	"encoding/json"
	"fmt"
	"git.sr.ht/~spc/go-log"
	"github.com/m-horky/subman-facts/pkg/facts"
	"github.com/nqd/flat"
	"strings"
)

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
		RegionRedirects: map[string]string{
			"us-gov-east-1": "us-east-2",
			"us-gov-west-1": "us-west-2",
		},
	}
}

// SystemRunsOn employs a strong detection of host system.
// It is the equivalent of cloud-what's `is_running_on_cloud()`.
func (_ AWSInstance) SystemRunsOn() bool {
	if systemIsVM() {
		return false
	}

	dmiCollector := facts.NewDmidecodeCollector()
	dmiFacts, err := dmiCollector.GetData(false)
	if err != nil {
		log.Info("could not get dmidecode facts, playing it safe: not running on AWS")
		return false
	}
	if strings.Contains(dmiFacts.BIOS.Version, "amazon") {
		return true
	}
	if strings.Contains(dmiFacts.BIOS.Vendor, "Amazon EC2") {
		return true
	}

	virtCollector := facts.NewVirtCollector()
	virtFacts, err := virtCollector.GetData(false)
	if err != nil {
		log.Info("could not get virt facts, playing it safe: not running on AWS")
		return false
	}
	if strings.Contains(virtFacts.HostType, "aws") {
		return true
	}

	return false
}

// SystemMaybeRunsOn employs a strong detection of host system.
// It is the equivalent of cloud-what's `is_likely_running_on_cloud()`.
// Returns the probability of the system running on this cloud.
func (i AWSInstance) SystemMaybeRunsOn() float64 {
	if systemIsVM() {
		return 0.0
	}

	virtCollector := facts.NewVirtCollector()
	virtFacts, err := virtCollector.GetData(false)
	if err != nil {
		log.Info("could not get virt facts, playing it safe: not running on AWS")
		return 0.0
	}
	dmiCollector := facts.NewDmidecodeCollector()
	dmiFacts, err := dmiCollector.GetData(false)
	if err != nil {
		log.Info("could not get dmidecode facts, playing it safe: not running on AWS")
		return 0.0
	}

	var probability = 0.0

	// We know AWS mostly uses KVM and in some cases Xen
	if strings.Contains(virtFacts.HostType, "kvm") {
		probability += 0.3
	} else if strings.Contains(virtFacts.HostType, "xen") {
		probability += 0.2
	}

	// Every UUID of a system running on AWS EC2 starts with the EC2 string
	if strings.HasPrefix(strings.ToLower(dmiFacts.System.UUID), "ec2") {
		probability += 0.1
	}

	// Strings "Amazon EC2", "Amazon" and "AWS" may be keywords in output of dmidecode
	var foundAmazonEC2 bool
	var foundAmazon bool
	var foundAWS bool
	flatennedDMIFacts, err := i.flattenDmiFacts(dmiFacts)
	if err == nil {
		for _, rawValue := range flatennedDMIFacts {
			value := strings.ToLower(fmt.Sprintf("%v", rawValue))
			if strings.Contains(value, "amazon ec2") {
				foundAmazonEC2 = true
			} else if strings.Contains(value, "amazon") {
				foundAmazon = true
			} else if strings.Contains(value, "aws") {
				foundAWS = true
			}
		}
	}
	if foundAmazonEC2 {
		probability += 0.3
	}
	if foundAmazon {
		probability += 0.2
	}
	if foundAWS {
		probability += 0.1
	}

	return probability
}

func (_ AWSInstance) flattenDmiFacts(facts facts.DmidecodeFacts) (map[string]any, error) {
	// Convert the nested structure into generic nested map
	temp, err := json.Marshal(&facts)
	if err != nil {
		log.Errorf("could not marshal DMI facts: %s", err)
		return map[string]any{}, err
	}
	var converted map[string]any
	err = json.Unmarshal(temp, &converted)
	if err != nil {
		log.Errorf("could not unmarshal marshalled DMI facts: %s", err)
		return map[string]any{}, err
	}

	flattened, err := flat.Flatten(converted, nil)
	if err != nil {
		log.Errorf("could not flatten DMI facts: %s", err)
		return map[string]any{}, err
	}
	return flattened, err
}

// getToken obtains the token from TokenURL and follows the scheme described
// in https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-service.html.
// When the token is received from the server, it is cached locally in a file.
func (i AWSInstance) getToken() {
	log.Debugf("trying to get AWS token from %s", i.TokenURL)
}

func (i AWSInstance) getTokenFromServer() {

}

func (i AWSInstance) getTokenFromCache() {

}

// GetMetadata returns the metadata obtained from IMDS server.
// A cache may be used, if still valid.
func (i AWSInstance) GetMetadata() {

}

// GetSignature returns the metadata signature obtained from IMDS server.
// A cache may be used, if still valid.
func (i AWSInstance) GetSignature() {

}
