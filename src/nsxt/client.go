package nsxt

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"go-vrf/src/configs"
)

func RequestNSXTApi(url, edge string) (r *http.Response, err error) {
	var (
		user = configs.GetEnvKeys(fmt.Sprintf("%s_USERNAME", edge))
		pass = configs.GetEnvKeys(fmt.Sprintf("%s_PASSWORD", edge))
	)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	d := http.Client{
		Timeout:   time.Second * 10,
		Transport: tr,
	}

	// Define retry attempts
	maxRetries := 5

	for attempt := 1; attempt <= maxRetries; attempt++ {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}

		req.SetBasicAuth(user, pass)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")

		res, err := d.Do(req)

		if err == nil {
			return res, nil // Successful response, return
		}

		// Handle error on attempt. Consider logging or specific checks here
		log.Printf("Request failed on attempt %d: %v\n", attempt, err)

		// Implement backoff strategy (optional)
		time.Sleep(time.Second * 2) // Exponential backoff
	}

	// All retries failed, return final error
	return nil, fmt.Errorf("failed to connect to NSX-T after %d attempts", maxRetries)
}
