package utils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"regexp"
	"strings"
	"time"

	"github.com/nebula-aac/kubernexus-mono/pkg/errors"
)

// GetHome returns the home path
func GetHome() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

func DownloadFile(filepath string, url string) error {
	// Create an HTTP client with custom settings if needed
	client := &http.Client{}

	// Create a request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			// Handle error from closing the response body
			err = fmt.Errorf("failed to close response body: %w", closeErr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file, received status code %d for URL: %s", resp.StatusCode, url)
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// Copy the response body to the file
	_, err = io.CopyBuffer(out, resp.Body, make([]byte, 8192))
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func safeClose(co io.Closer) {
	if cerr := co.Close(); cerr != nil {
		errors.Errorf("error with closing: %v", errors.Alert, cerr.Error())
	}
}

func GetLatestReleaseTagsSorted(org string, repo string) ([]string, error) {
	ctx := context.Background()
	client := &http.Client{
		Timeout: 10 * time.Second, // Set a timeout for requests
	}

	url := "https://github.com/" + org + "/" + repo + "/releases"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer safeClose(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received status code %d for %s", resp.StatusCode, url)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile("/releases/tag/(.*?)\"")
	releases := re.FindAllString(string(body), -1)
	if len(releases) == 0 {
		return nil, fmt.Errorf("no releases found")
	}
	var versions []string
	for _, rel := range releases {
		latest := strings.ReplaceAll(rel, "/releases/tag/", "")
		latest = strings.ReplaceAll(latest, "\"", "")
		versions = append(versions, latest)
	}
	versions = SortDottedStringsByDigits(versions)
	return versions, nil
}

// ReadFileSource supports "http", "https", and "file" protocols.
// It takes in the location as a URI and returns the contents of
// the file as a string.
func ReadFileSource(uri string) (string, error) {
	if strings.HasPrefix(uri, "http") {
		return ReadRemoteFile(uri)
	}
	if strings.HasPrefix(uri, "file") {
		return ReadLocalFile(uri)
	}

	return "", fmt.Errorf("invalid protocol: %s", uri)
}

// ReadRemoteFile takes in the location of a remote file
// in the format 'http://location/of/file' or 'https://location/file'
// and returns the content of the file if the location is valid and
// no error occurs.
func ReadRemoteFile(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to read remote file: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return "", fmt.Errorf("remote file not found: %s", url)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read remote file body: %w", err)
	}

	return string(body), nil
}

// ReadLocalFile takes in the location of a local file
// in the format `file://location/of/file` and returns
// the content of the file if the path is valid and no
// error occurs.
func ReadLocalFile(location string) (string, error) {
	// Remove the protocol prefix
	location = strings.TrimPrefix(location, "file://")

	// Need to support variable file locations, hence #nosec
	data, err := os.ReadFile(location)
	if err != nil {
		return "", fmt.Errorf("failed to read local file: %w", err)
	}

	return string(data), nil
}

func IsClosed(ch chan struct{}) bool {
	if ch == nil {
		return true
	}
	select {
	case <-ch:
		return true
	default:
	}
	return false
}
