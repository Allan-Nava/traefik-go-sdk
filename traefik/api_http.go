package traefik

import "github.com/go-resty/resty/v2"

/*
*/
func (o *traefikSdk) GetHttpRouters() (*resty.Response, error) {
	//log.Println("GetRoutes ", deviceId)
	resp, err := o.restyGet(HTTP_ROUTERS, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
*/
func (o *traefikSdk) GetHttpRouter(routerName string) (*resty.Response, error) {
	resp, err := o.restyGet(GET_HTTP_ROUTER(routerName), nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}


/*
*/
func (o *traefikSdk) GetHttpMiddlewares() (*resty.Response, error) {
	//log.Println("GetRoutes ", deviceId)
	resp, err := o.restyGet(HTTP_MIDDLEWARES, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
*/
func (o *traefikSdk) GetHttpServices() (*resty.Response, error) {
	//log.Println("GetRoutes ", deviceId)
	resp, err := o.restyGet(HTTP_SERVICES, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}