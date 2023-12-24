package externalfile

import (
	"app/pkg/logger"
	"net/http"
	"net/url"
	"time"
)

const storageName = "external file"

type Storage struct {
	token string

	scheme       string
	hostWithPort string

	client *http.Client

	logger *logger.Logger
}

func New(token string, scheme string, hostWithPort string, logger *logger.Logger) *Storage {
	return &Storage{
		token:        token,
		scheme:       scheme,
		hostWithPort: hostWithPort,
		client: &http.Client{
			Timeout: time.Minute,
		},
		logger: logger,
	}
}

func (s *Storage) url(path string) string {
	u := url.URL{
		Scheme: s.scheme,
		Host:   s.hostWithPort,
		Path:   path,
	}

	return u.String()
}