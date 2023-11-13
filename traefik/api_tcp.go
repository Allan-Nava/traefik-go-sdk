package traefik

import "github.com/go-resty/resty/v2"

/*
*/
func(o *traefikSdk) GetTcpServices() (*resty.Response, error){
	resp, err := o.restyGet(TCP_SERVICES, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
/*
*/
func (o *traefikSdk) GetTcpRouters() (*resty.Response, error) {
	//log.Println("GetRoutes ", deviceId)
	resp, err := o.restyGet(TCP_ROUTERS, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
*/
func (o *traefikSdk) GetTcpRouter(name string) (*resty.Response, error) {
	resp, err := o.restyGet(GET_TCP_ROUTER(name), nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
/*
*/
func (o *traefikSdk) GetUdpService(name string) (*resty.Response, error){
	resp, err := o.restyGet(GET_UDP_SERVICE(name), nil)
	if err != nil {
		return nil, err
	}
	return resp, nil

}