package traefik

import "github.com/go-resty/resty/v2"

type traefikSdk struct {
	BaseUrl        string
	restClient *resty.Client
	debug      bool
}

type ITraefikClient interface {
	//
	HealthCheck() error
	IsDebug() bool
	//
}

func (o *traefikSdk) HealthCheck() error {
	_, err := o.restyGet(o.BaseUrl, nil)
	if err != nil {
		return nil
	}
	return nil
}

// Resty Methods

func (o *traefikSdk) restyPost(url string, body interface{}) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(body).
		Post(url)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// get request
func (o *traefikSdk) restyGet(url string, queryParams map[string]string) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetQueryParams(queryParams).
		Get(url)
	//
	if err != nil {
		return nil, err
	}
	return resp, nil
}
