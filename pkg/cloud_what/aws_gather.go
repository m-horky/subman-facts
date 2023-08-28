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
	cachedData map[string]any
	cachedAt   uint
	cacheTTL   uint
}

// getToken obtains the token from a cache or from the TokenURL and follows the scheme described
// in https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-service.html.
// When the token is received from the server, it is cached locally in a file.
func (i AWSInstance) getToken() (awsToken, error) {
	token, err := i.getTokenFromFileCache()
	if err == nil {
		log.Debugf("using token from the cache file")
		return token, nil
	}

	return i.getTokenFromServer()
}

// getTokenFromFileCache tries to read the cache file and return cached token.
// An error is returned when no cache file exists or when the token is not valid anymore.
func (i AWSInstance) getTokenFromFileCache() (awsToken, error) {
	// TODO Read from `i.TokenCacheFile`
	log.Warn("Cache file is not implemented.")
	return awsToken{}, fmt.Errorf("getTokenFromFileCache(): Not implemented")
}

// saveTokenToFileCache saves the token into a file. It can be retrieved later using
// getTokenFromFileCache.
func (i AWSInstance) saveTokenToFileCache(token awsToken) error {
	// TODO Save to `i.TokenCacheFile`
	log.Warn("Cache file is not implemented.")
	return nil
}

// getTokenFromServer fetches the token from the server and saves it into a cache file.
func (i AWSInstance) getTokenFromServer() (awsToken, error) {
	log.Debugf("requesting AWS token from %s", i.TokenURL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(i.ServerTimeout)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, i.TokenURL, nil)
	if err != nil {
		return awsToken{}, err
	}

	// TODO Proxy settings need to be configured through the Transport here
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	for k, v := range i.CustomHTTPHeaders {
		req.Header.Add(k, v)
	}

	resp, err := IdentityService.Call(client, req)
	defer resp.Body.Close()

	if err != nil {
		log.Errorf("unable to receive the token from AWS: %s", err)
		return awsToken{}, err
	}

	if resp.StatusCode != 200 {
		log.Errorf("unable to receive the token from AWS, got status code %d", resp.StatusCode)
		return awsToken{}, fmt.Errorf("server responded with status code %d", resp.StatusCode)
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("unable to read the token from AWS: %s", err)
		return awsToken{}, err
	}

	var token awsToken
	err = json.Unmarshal(response, &token)
	if err != nil {
		log.Errorf("could not decode AWS token content: %s", err)
		return awsToken{}, err
	}

	_ = i.saveTokenToFileCache(token)
	return token, nil
}

// GetMetadata returns the metadata obtained from IMDS server.
// A cache may be used, if still valid.
func (i AWSInstance) GetMetadata() {

}

// GetSignature returns the metadata signature obtained from IMDS server.
// A cache may be used, if still valid.
func (i AWSInstance) GetSignature() {

}
