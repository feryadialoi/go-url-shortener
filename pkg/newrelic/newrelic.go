package newrelic

import (
	"github.com/feryadialoi/go-url-shortener/config"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

func New(conf config.Config) (*newrelic.Application, error) {
	logrus.Info("New New Relic")

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(conf.Name),
		newrelic.ConfigLicense(conf.NewRelicLicenseKey),
	)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func MustNew(conf config.Config) *newrelic.Application {
	app, err := New(conf)
	if err != nil {
		logrus.Fatalf("Failed to create new relic agent instance, err: %v", err)
	}

	return app
}
