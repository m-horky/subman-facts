package cloud_what

import (
	"context"
	"encoding/json"
	"fmt"
	"git.sr.ht/~spc/go-log"
	"io"
	"net/http"
	"time"
)

type awsToken struct {
	data     string
	cachedAt int64
	cacheTTL uint
}

func newAWSToken(token string, cachedAt int64, cacheTTL uint) awsToken {
	return awsToken{data: token, cachedAt: cachedAt, cacheTTL: cacheTTL}
}

func (t awsToken) IsValid() bool {
	now := time.Now().Unix()
	validUntil := t.cachedAt + int64(t.cacheTTL)
	return now <= validUntil
}

// Get returns the token if it is still valid.
// An error is returned when no token is cached or if it has expired.
func (t awsToken) Get() (string, error) {
	// TODO
	return "", fmt.Errorf("not implemented")
}

// Save caches the token into a file.
// An error is returned when the token could not be saved.
func (t awsToken) Save() error {
	path := NewAWSInstance().Configuration.TokenCacheFile
	return fmt.Errorf("write to path %s not implemented", path)
}

type AWSMetadata struct {
	AccountId               string   `json:"accountId"`
	Architecture            string   `json:"architecture"`
	AvailabilityZone        string   `json:"availabilityZone"`
	BillingProducts         []string `json:"billingProducts"`
	DevpayProductCodes      []string `json:"devpayProductCodes"`
	MarketplaceProductCodes []string `json:"marketplaceProductCodes"`
	ImageId                 string   `json:"imageId"`
	InstanceId              string   `json:"instanceId"`
	InstanceType            string   `json:"instanceType"`
	KernelId                string   `json:"kernelId"`
	PendingTime             string   `json:"pendingTime"`
	PrivateIp               string   `json:"privateIp"`
	RamdiskId               string   `json:"ramdiskId"`
	Region                  string   `json:"region"`
	Version                 string   `json:"version"`
}

// getToken obtains the token from a cache or from the TokenURL and follows the scheme described
// in https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-service.html.
// When the token is received from the server, it is cached locally in a file.
func (i AWSInstance) getToken() (string, error) {
	token, err := awsToken{}.Get()
	if err == nil {
		log.Debugf("using token from the cache file")
		return token, nil
	}

	return i.getTokenFromServer()
}

// getTokenFromFileCache tries to read the cache file and return cached token.
// An error is returned when no cache file exists or when the token is not valid anymore.
func (i AWSInstance) getTokenFromFileCache() (string, error) {
	// TODO Read from `i.TokenCacheFile`
	log.Warn("Cache file is not implemented.")
	return "", fmt.Errorf("getTokenFromFileCache(): Not implemented")
}

// saveTokenToFileCache saves the token into a file. It can be retrieved later using
// getTokenFromFileCache.
func (i AWSInstance) saveTokenToFileCache(token string) error {
	// TODO Save to `i.TokenCacheFile`
	log.Warn("Cache file is not implemented.")
	return nil
}

// getTokenFromServer fetches the token from the server and saves it into a cache file.
func (i AWSInstance) getTokenFromServer() (string, error) {
	log.Debugf("requesting AWS token from %s", i.Configuration.TokenURL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(i.Configuration.ServerTimeout)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, i.Configuration.TokenURL, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	for k, v := range i.Configuration.CustomHTTPHeaders {
		req.Header.Add(k, v)
	}

	resp, err := IdentityService.Call(client, req)
	if err != nil {
		log.Errorf("unable to receive the token from AWS: %s", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Errorf("unable to receive the token from AWS, got status code %d", resp.StatusCode)
		return "", fmt.Errorf("server responded with status code %d", resp.StatusCode)
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("unable to read the token from AWS: %s", err)
		return "", err
	}

	token := string(response)
	if err != nil {
		log.Errorf("could not decode AWS token content: %s", err)
		return "", err
	}

	cache := newAWSToken(token, time.Now().Unix(), i.Configuration.TokenTTL)
	err = cache.Save()
	if err != nil {
		log.Errorf("could not cache AWS token: %s", err)
	}

	return token, nil
}

// GetMetadata returns the metadata obtained from IMDS server.
// A cache may be used, if still valid.
func (i AWSInstance) GetMetadata() (AWSMetadata, error) {
	log.Debugf("requesting AWS metadata from %s", i.Configuration.MetadataURL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(i.Configuration.ServerTimeout)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, i.Configuration.MetadataURL, nil)
	if err != nil {
		return AWSMetadata{}, err
	}

	client := &http.Client{}
	for k, v := range i.Configuration.CustomHTTPHeaders {
		req.Header.Add(k, v)
	}

	resp, err := IdentityService.Call(client, req)
	if err != nil {
		log.Errorf("unable to receive the metadata from AWS: %s", err)
		return AWSMetadata{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Errorf("unable to receive the token from AWS, got status code %d", resp.StatusCode)
		return AWSMetadata{}, fmt.Errorf("server responded with status code %d", resp.StatusCode)
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("unable to read the metadata from AWS: %s", err)
		return AWSMetadata{}, err
	}

	var metadata AWSMetadata
	err = json.Unmarshal(response, &metadata)
	if err != nil {
		log.Errorf("could not decode AWS metadata: %s", err)
		return AWSMetadata{}, err
	}

	return metadata, nil
}

// GetSignature returns the metadata signature obtained from IMDS server.
// A cache may be used, if still valid.
func (i AWSInstance) GetSignature() {

}
