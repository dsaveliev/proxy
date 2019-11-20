package client

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/hellofreshdevtests/dsaveliev-golang-test/internal/config"
)

var once sync.Once
var client *http.Client

// Setup reasonable timeouts and configure TLS
func initialize() {
	cfg := config.GetConfig()

	var tlsConfig *tls.Config
	if caCert, err := ioutil.ReadFile(cfg.CertPath); err == nil {
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConfig = &tls.Config{RootCAs: caCertPool}
	} else {
		tlsConfig = &tls.Config{InsecureSkipVerify: true}
	}

	client = &http.Client{
		Timeout: time.Duration(cfg.ClientTimeout) * time.Millisecond,
		Transport: &http.Transport{
			MaxIdleConns:    cfg.MaxIdleConns,
			IdleConnTimeout: time.Duration(cfg.IdleConnTimeout) * time.Millisecond,
			TLSClientConfig: tlsConfig,
		},
	}
}

// SetClient initialize client for the purposes of testing
func SetClient(c *http.Client) {
	client = c
}

// GetClient method initializes the http client
func GetClient() *http.Client {
	once.Do(initialize)
	return client
}
