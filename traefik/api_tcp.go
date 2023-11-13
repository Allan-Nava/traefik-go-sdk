package traefik

import "github.com/go-resty/resty/v2"
/*
*/
func (o *traefikSdk) GetTcpRouters() (*resty.Response, error) {
	//log.Println("GetRoutes ", deviceId)
	resp, err := o.restyGet(HTTP_ROUTERS, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
*/
func (o *traefikSdk) GetTcpRouter(name string) (*resty.Response, error) {
	resp, err := o.restyGet(GET_HTTP_TCP_ROUTER(name), nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}