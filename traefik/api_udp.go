package traefik

import "github.com/go-resty/resty/v2"
//
/*
*/
func (o *traefikSdk) GetUdpRouters() (*resty.Response, error){
	resp, err := o.restyGet(UDP_ROUTERS, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *traefikSdk) GetUdpRouter(name string) (*resty.Response, error) {
	resp, err := o.restyGet(UDP_ROUTERS, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *traefikSdk)  GetUdpServices() (*resty.Response, error){
	resp, err := o.restyGet(UDP_ROUTERS, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}