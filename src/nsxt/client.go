package nsxt

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"go-vrf/src/configs"
)

const (
	maxRetries     = 5
	backoffInitial = 500 * time.Millisecond
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

func RequestNSXTApi(ctx context.Context, url, edge string) (*http.Response, error) {
	var (
		user = configs.GetEnvKeys(fmt.Sprintf("%s_USERNAME", edge))
		pass = configs.GetEnvKeys(fmt.Sprintf("%s_PASSWORD", edge))
	)

	d, err := getClient()
	if err != nil {
		return nil, err
	}

	var lastErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}

		req.SetBasicAuth(user, pass)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")

		res, err := d.Do(req)
		if err != nil {
			// Don't retry if the caller cancelled / timed out.
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			lastErr = err
			log.Printf("request failed on attempt %d: %v", attempt, err)
		} else if res.StatusCode >= 500 {
			// Server-side error: drain and close before retrying so the
			// connection can be reused.
			lastErr = fmt.Errorf("server returned status %d", res.StatusCode)
			_, _ = io.Copy(io.Discard, res.Body)
			_ = res.Body.Close()
			log.Printf("request failed on attempt %d: status %d", attempt, res.StatusCode)
		} else {
			// 2xx, 3xx, or 4xx: hand back to the caller (fetch checks status).
			return res, nil
		}

		if attempt < maxRetries {
			if err := sleepBackoff(ctx, attempt); err != nil {
				return nil, err
			}
		}
	}

	return nil, fmt.Errorf("failed to connect to NSX-T after %d attempts: %w", maxRetries, lastErr)
}

// sleepBackoff waits backoffInitial * 2^(attempt-1), returning early with
// the context error if the context is cancelled while waiting.
func sleepBackoff(ctx context.Context, attempt int) error {
	delay := backoffInitial * (1 << (attempt - 1))
	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
