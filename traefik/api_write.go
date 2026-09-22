package traefik

import "github.com/go-resty/resty/v2"

// CreateHttpRouter creates a new HTTP router in Traefik
func (o *traefikSdk) CreateHttpRouter(name string, config interface{}) (*resty.Response, error) {
	url := HTTP_ROUTERS + "/" + name
	return o.restyPost(url, config)
}

// UpdateHttpRouter updates an existing HTTP router in Traefik
func (o *traefikSdk) UpdateHttpRouter(name string, config interface{}) (*resty.Response, error) {
	url := HTTP_ROUTERS + "/" + name
	fullURL := o.BaseUrl + url
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Put(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteHttpRouter deletes an HTTP router from Traefik
func (o *traefikSdk) DeleteHttpRouter(name string) (*resty.Response, error) {
	url := HTTP_ROUTERS + "/" + name
	fullURL := o.BaseUrl + url
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		Delete(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateHttpService creates a new HTTP service in Traefik
func (o *traefikSdk) CreateHttpService(name string, config interface{}) (*resty.Response, error) {
	url := HTTP_SERVICES + "/" + name
	return o.restyPost(url, config)
}

// UpdateHttpService updates an existing HTTP service in Traefik
func (o *traefikSdk) UpdateHttpService(name string, config interface{}) (*resty.Response, error) {
	url := HTTP_SERVICES + "/" + name
	fullURL := o.BaseUrl + url
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Put(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteHttpService deletes an HTTP service from Traefik
func (o *traefikSdk) DeleteHttpService(name string) (*resty.Response, error) {
	url := HTTP_SERVICES + "/" + name
	fullURL := o.BaseUrl + url
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		Delete(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
