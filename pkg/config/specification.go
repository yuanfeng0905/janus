package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Specification for basic configurations
type Specification struct {
	Port                int    `envconfig:"PORT" default:"8080"`
	APIPort             int    `envconfig:"API_PORT" default:"8081"`
	Debug               bool   `envconfig:"DEBUG" description:"Enable debug mode"`
	LogLevel            string `envconfig:"LOG_LEVEL" default:"info" description:"Log level"`
	GraceTimeOut        int64  `envconfig:"GRACE_TIMEOUT" description:"Duration to give active requests a chance to finish during hot-reload"`
	MaxIdleConnsPerHost int    `envconfig:"MAX_IDLE_CONNS_PER_HOST" description:"If non-zero, controls the maximum idle (keep-alive) to keep per-host."`
	InsecureSkipVerify  bool   `envconfig:"INSECURE_SKIP_VERIFY" description:"Disable SSL certificate verification"`
	// The Storage DSN, this could be `memory` or `redis`
	StorageDSN string `envconfig:"STORAGE_DSN" default:"memory://localhost"`

	// Path of certificate when using TLS
	CertPathTLS string `envconfig:"CERT_PATH"`

	// Path of key when using TLS
	KeyPathTLS string `envconfig:"KEY_PATH"`

	// Flush interval for upgraded Proxy connections
	BackendFlushInterval time.Duration `envconfig:"BACKEND_FLUSH_INTERVAL" default:"20ms"`

	// Defines the time period of how often the idle connections maintained
	// by the proxy are closed.
	CloseIdleConnsPeriod time.Duration `envconfig:"CLOSE_IDLE_CONNS_PERIOD"`

	Database    Database
	Statsd      Statsd
	Credentials Credentials
}

// IsHTTPS checks if you have https enabled
func (s *Specification) IsHTTPS() bool {
	return s.CertPathTLS != "" && s.KeyPathTLS != ""
}

// Database holds the configuration for a database
type Database struct {
	DSN string `envconfig:"DATABASE_DSN" default:"file:///etc/janus"`
}

// Statsd holds the configuration for statsd
type Statsd struct {
	DSN    string `envconfig:"STATSD_DSN"`
	Prefix string `envconfig:"STATSD_PREFIX"`
	IDs    string `envconfig:"STATSD_IDS"`
}

// Credentials represents the credentials that are going to be
// used by JWT configuration
type Credentials struct {
	Secret   string `envconfig:"SECRET" required:"true"`
	Username string `envconfig:"ADMIN_USERNAME" default:"admin"`
	Password string `envconfig:"ADMIN_PASSWORD" default:"admin"`
}

//LoadEnv loads environment variables
func LoadEnv() (*Specification, error) {
	var config Specification
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
