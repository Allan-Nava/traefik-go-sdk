package traefik

import (
	"github.com/go-resty/resty/v2"
	"log"
)

type traefikSdk struct {
	BaseUrl    string
	restClient *resty.Client
	debug      bool
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
	GetUdpService(name string) (*resty.Response, error)
	GetEntrypoints() (*resty.Response, error)
	GetApiOverview() (*resty.Response, error)
	GetApiRawData() (*resty.Response, error)
	GetApiVersion() (*resty.Response, error)
	CreateHttpRouter(name string, config interface{}) (*resty.Response, error)
	UpdateHttpRouter(name string, config interface{}) (*resty.Response, error)
	DeleteHttpRouter(name string) (*resty.Response, error)
	CreateHttpService(name string, config interface{}) (*resty.Response, error)
	UpdateHttpService(name string, config interface{}) (*resty.Response, error)
	DeleteHttpService(name string) (*resty.Response, error)
	ApplyConfiguration(config interface{}) (*resty.Response, error)
	GetConfiguration() (*resty.Response, error)
	ValidateConfiguration(config interface{}) (*resty.Response, error)
	ResetConfiguration() (*resty.Response, error)
	CreateTcpRouter(name string, config interface{}) (*resty.Response, error)
	UpdateTcpRouter(name string, config interface{}) (*resty.Response, error)
	DeleteTcpRouter(name string) (*resty.Response, error)
	CreateTcpService(name string, config interface{}) (*resty.Response, error)
	UpdateTcpService(name string, config interface{}) (*resty.Response, error)
	DeleteTcpService(name string) (*resty.Response, error)
	CreateUdpRouter(name string, config interface{}) (*resty.Response, error)
	UpdateUdpRouter(name string, config interface{}) (*resty.Response, error)
	DeleteUdpRouter(name string) (*resty.Response, error)
	CreateUdpService(name string, config interface{}) (*resty.Response, error)
	UpdateUdpService(name string, config interface{}) (*resty.Response, error)
	DeleteUdpService(name string) (*resty.Response, error)
	CreateHttpMiddleware(name string, config interface{}) (*resty.Response, error)
	UpdateHttpMiddleware(name string, config interface{}) (*resty.Response, error)
	DeleteHttpMiddleware(name string) (*resty.Response, error)
	CreateTcpMiddleware(name string, config interface{}) (*resty.Response, error)
	UpdateTcpMiddleware(name string, config interface{}) (*resty.Response, error)
	DeleteTcpMiddleware(name string) (*resty.Response, error)
	//
}

// Builder is used to build a new haivision client
func BuildTraefik(url string, debug bool) (ITraefikClient, error) {
	// init haivision
	traefikClient := &traefikSdk{
		BaseUrl:    url,
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
	fullURL := o.BaseUrl + url
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(body).
		Post(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// get request
func (o *traefikSdk) restyGet(url string, queryParams map[string]string) (*resty.Response, error) {
	fullURL := o.BaseUrl + url
	resp, err := o.restClient.R().
		SetQueryParams(queryParams).
		Get(fullURL)
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
