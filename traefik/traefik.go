package traefik

import (
	"log"
	"github.com/go-resty/resty/v2"
)

type traefikSdk struct {
	BaseUrl        	string
	restClient 		*resty.Client
	debug      		bool
}

type ITraefikClient interface {
	//
	HealthCheck() error
	IsDebug() bool
	GetHttpRouters() (*resty.Response, error)
	GetHttpRouter(routerName string) (*resty.Response, error)
	GetHttpServices() (*resty.Response, error)
	GetHttpMiddlewares() (*resty.Response, error)
	GetTcpServices() (*resty.Response, error)
	GetTcpRouters() (*resty.Response, error)
	GetTcpRouter(name string) (*resty.Response, error)
	GetUdpRouters() (*resty.Response, error)
	GetUdpRouter(name string) (*resty.Response, error)
	GetUdpServices() (*resty.Response, error)
	//
}

// Builder is used to build a new haivision client
func BuildTraefik(url string, debug bool) (ITraefikClient, error) {
	// init haivision
	traefikClient := &traefikSdk{
		BaseUrl:        url,
		restClient: resty.New(),
	}
	//
	if debug {
		traefikClient.restClient.SetDebug(true)
		traefikClient.debug = true
	}
	//
	return traefikClient, nil
}

func (o *traefikSdk) HealthCheck() error {
	_, err := o.restyGet(o.BaseUrl, nil)
	if err != nil {
		return nil
	}
	return nil
}

func (o *traefikSdk) IsDebug() bool {
	return o.debug
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




func (o *traefikSdk) debugPrint(data interface{}) {
	if o.debug {
		log.Println(data)
	}
}