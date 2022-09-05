package configuration

import (
	"errors"
	"github.com/kelseyhightower/envconfig"
	"net/url"
)

type UrlAddress string

type Configuration struct {
	Port        int        `envconfig:"port" default:"8080" required:"true"`
	DbUrl       UrlAddress `envconfig:"db_url" default:"https://github.com/kelseyhightower/envconfig"  required:"true"`
	JaegerUrl   UrlAddress `envconfig:"jaeger_url"  required:"true"`
	SentryUrl   UrlAddress `envconfig:"sentry_url" required:"true"`
	KafkaBroker UrlAddress `envconfig:"kafka_broker" required:"true"`
	SomeAppId   string     `envconfig:"some_app_id" required:"true"`
	SomeAppKey  string     `envconfig:"some_app_key" required:"true"`
}

// Тут уже можно реализовать все необходимые проверки по парсингу адреса
func (ua *UrlAddress) Decode(value string) error {
	urlData, _ := url.Parse(value)
	if urlData.Scheme != "https" {
		return errors.New("Протокол должен быть Https!")
	}
	*ua = UrlAddress(urlData.String())

	return nil
}

func GetConfiguration() (Configuration, error) {
	conf := Configuration{}
	err := envconfig.Process("envconfig", &conf)
	if err != nil {
		return Configuration{}, err
	}

	return conf, nil
}
