package cloud_what

import (
	"git.sr.ht/~spc/go-log"
	"github.com/m-horky/subman-facts/pkg/facts"
	"strings"
)

type AWSInstance struct {
	Cloud
	RedirectMap map[string]string
}

func newAWSInstance() AWSInstance {
	return AWSInstance{
		Cloud: Cloud{
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
		},
		RedirectMap: map[string]string{
			"us-gov-east-1": "us-east-2",
			"us-gov-west-1": "us-west-2",
		},
	}
}

func (i AWSInstance) GetID() string {
	return i.ID
}

// systemRunsOn employs a strong detection of host system.
// It is the equivalent of cloud-what's `is_running_on_cloud()`.
func (_ AWSInstance) systemRunsOn() bool {
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

// systemMaybeRunsOnAWS employs a strong detection of host system.
// It is the equivalent of cloud-what's `is_likely_running_on_cloud()`.
// Returns the probability of the system running on this cloud.
func (_ AWSInstance) systemMaybeRunsOn() float64 {
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

	// TODO How can we get the values from arbitrarily nested keys?
	//  Don't we have some specific list of keys to check instead?

	return probability
}
