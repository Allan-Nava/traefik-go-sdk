package traefik

import "github.com/go-resty/resty/v2"

// HTTP Middleware Write Operations

func (o *traefikSdk) CreateHttpMiddleware(name string, config interface{}) (*resty.Response, error) {
	fullURL := o.BaseUrl + HTTP_MIDDLEWARES + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Post(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *traefikSdk) UpdateHttpMiddleware(name string, config interface{}) (*resty.Response, error) {
	fullURL := o.BaseUrl + HTTP_MIDDLEWARES + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Put(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *traefikSdk) DeleteHttpMiddleware(name string) (*resty.Response, error) {
	fullURL := o.BaseUrl + HTTP_MIDDLEWARES + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		Delete(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TCP Middleware Write Operations

func (o *traefikSdk) CreateTcpMiddleware(name string, config interface{}) (*resty.Response, error) {
	fullURL := o.BaseUrl + TCP_MIDDLEWARES + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Post(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *traefikSdk) UpdateTcpMiddleware(name string, config interface{}) (*resty.Response, error) {
	fullURL := o.BaseUrl + TCP_MIDDLEWARES + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Put(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *traefikSdk) DeleteTcpMiddleware(name string) (*resty.Response, error) {
	fullURL := o.BaseUrl + TCP_MIDDLEWARES + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		Delete(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
