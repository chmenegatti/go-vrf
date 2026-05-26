package nsxt

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"go-vrf/src/configs"
)

var (
	clientOnce sync.Once
	httpClient *http.Client
	clientErr  error
)

func buildClient() (*http.Client, error) {
	skip := strings.EqualFold(configs.GetEnvKeys("NSXT_SKIP_TLS_VERIFY"), "true")

	tlsCfg := &tls.Config{InsecureSkipVerify: skip}

	if caPath := configs.GetEnvKeys("NSXT_CA_BUNDLE"); caPath != "" {
		pem, err := os.ReadFile(caPath)
		if err != nil {
			return nil, fmt.Errorf("read NSXT_CA_BUNDLE %q: %w", caPath, err)
		}
		pool, _ := x509.SystemCertPool()
		if pool == nil {
			pool = x509.NewCertPool()
		}
		if !pool.AppendCertsFromPEM(pem) {
			return nil, fmt.Errorf("NSXT_CA_BUNDLE %q has no valid certificates", caPath)
		}
		tlsCfg.RootCAs = pool
	}

	if skip {
		log.Println("WARNING: NSX-T TLS certificate verification is disabled (NSXT_SKIP_TLS_VERIFY=true)")
	}

	return &http.Client{
		Timeout:   10 * time.Second,
		Transport: &http.Transport{TLSClientConfig: tlsCfg},
	}, nil
}

func getClient() (*http.Client, error) {
	clientOnce.Do(func() {
		httpClient, clientErr = buildClient()
	})
	return httpClient, clientErr
}

func RequestNSXTApi(url, edge string) (r *http.Response, err error) {
	var (
		user = configs.GetEnvKeys(fmt.Sprintf("%s_USERNAME", edge))
		pass = configs.GetEnvKeys(fmt.Sprintf("%s_PASSWORD", edge))
	)

	d, err := getClient()
	if err != nil {
		return nil, err
	}

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
			return res, nil
		}

		log.Printf("Request failed on attempt %d: %v\n", attempt, err)

		time.Sleep(time.Second * 2)
	}

	return nil, fmt.Errorf("failed to connect to NSX-T after %d attempts", maxRetries)
}
