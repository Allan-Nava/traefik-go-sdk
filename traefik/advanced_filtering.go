package traefik

import "github.com/go-resty/resty/v2"

// GetHttpRoutersByRule returns HTTP routers matching a specific rule pattern
func (o *traefikSdk) GetHttpRoutersByRule(rule string) (*resty.Response, error) {
	queryParams := map[string]string{
		"rule": rule,
	}
	return o.restyGet(HTTP_ROUTERS, queryParams)
}

// GetTcpRoutersByEntryPoint returns TCP routers for a specific entrypoint
func (o *traefikSdk) GetTcpRoutersByEntryPoint(entryPoint string) (*resty.Response, error) {
	queryParams := map[string]string{
		"entryPoint": entryPoint,
	}
	return o.restyGet(TCP_ROUTERS, queryParams)
}

// GetUdpRoutersByEntryPoint returns UDP routers for a specific entrypoint
func (o *traefikSdk) GetUdpRoutersByEntryPoint(entryPoint string) (*resty.Response, error) {
	queryParams := map[string]string{
		"entryPoint": entryPoint,
	}
	return o.restyGet(UDP_ROUTERS, queryParams)
}

// GetServicesByRouter returns services associated with a specific router
func (o *traefikSdk) GetServicesByRouter(routerName string) (*resty.Response, error) {
	queryParams := map[string]string{
		"router": routerName,
	}
	return o.restyGet(HTTP_SERVICES, queryParams)
}

// GetHttpMiddlewareByType returns HTTP middlewares of a specific type
func (o *traefikSdk) GetHttpMiddlewareByType(middlewareType string) (*resty.Response, error) {
	queryParams := map[string]string{
		"type": middlewareType,
	}
	return o.restyGet(HTTP_MIDDLEWARES, queryParams)
}
