package traefik

import "github.com/go-resty/resty/v2"

// TCP Router Write Operations

func (o *traefikSdk) CreateTcpRouter(name string, config interface{}) (*resty.Response, error) {
	fullURL := o.BaseUrl + TCP_ROUTERS + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Post(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *traefikSdk) UpdateTcpRouter(name string, config interface{}) (*resty.Response, error) {
	fullURL := o.BaseUrl + TCP_ROUTERS + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Put(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *traefikSdk) DeleteTcpRouter(name string) (*resty.Response, error) {
	fullURL := o.BaseUrl + TCP_ROUTERS + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		Delete(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TCP Service Write Operations

func (o *traefikSdk) CreateTcpService(name string, config interface{}) (*resty.Response, error) {
	fullURL := o.BaseUrl + TCP_SERVICES + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Post(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *traefikSdk) UpdateTcpService(name string, config interface{}) (*resty.Response, error) {
	fullURL := o.BaseUrl + TCP_SERVICES + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Put(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *traefikSdk) DeleteTcpService(name string) (*resty.Response, error) {
	fullURL := o.BaseUrl + TCP_SERVICES + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		Delete(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UDP Router Write Operations

func (o *traefikSdk) CreateUdpRouter(name string, config interface{}) (*resty.Response, error) {
	fullURL := o.BaseUrl + UDP_ROUTERS + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Post(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *traefikSdk) UpdateUdpRouter(name string, config interface{}) (*resty.Response, error) {
	fullURL := o.BaseUrl + UDP_ROUTERS + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Put(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *traefikSdk) DeleteUdpRouter(name string) (*resty.Response, error) {
	fullURL := o.BaseUrl + UDP_ROUTERS + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		Delete(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UDP Service Write Operations

func (o *traefikSdk) CreateUdpService(name string, config interface{}) (*resty.Response, error) {
	fullURL := o.BaseUrl + UDP_SERVICES + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Post(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *traefikSdk) UpdateUdpService(name string, config interface{}) (*resty.Response, error) {
	fullURL := o.BaseUrl + UDP_SERVICES + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Put(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *traefikSdk) DeleteUdpService(name string) (*resty.Response, error) {
	fullURL := o.BaseUrl + UDP_SERVICES + "/" + name
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		Delete(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
