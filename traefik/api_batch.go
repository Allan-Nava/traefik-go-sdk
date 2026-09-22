package traefik

import "github.com/go-resty/resty/v2"

// ApplyConfiguration applies a complete configuration (HTTP + TCP + UDP routers, services, middlewares)
// This is a batch operation that replaces the current configuration
func (o *traefikSdk) ApplyConfiguration(config interface{}) (*resty.Response, error) {
	fullURL := o.BaseUrl + "/api/config"
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Post(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetConfiguration retrieves the complete current configuration
// Returns all routers, services, and middlewares for HTTP, TCP, and UDP
func (o *traefikSdk) GetConfiguration() (*resty.Response, error) {
	url := "/api/rawdata"
	return o.restyGet(url, nil)
}

// ValidateConfiguration validates a configuration before applying it
// Checks for errors and inconsistencies without actually applying the config
func (o *traefikSdk) ValidateConfiguration(config interface{}) (*resty.Response, error) {
	fullURL := o.BaseUrl + "/api/config/validate"
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Post(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ResetConfiguration resets the configuration to the default state
// Removes all user-defined routers, services, and middlewares
func (o *traefikSdk) ResetConfiguration() (*resty.Response, error) {
	fullURL := o.BaseUrl + "/api/config"
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		Delete(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
