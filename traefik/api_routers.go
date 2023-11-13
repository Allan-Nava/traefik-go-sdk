package traefik

/*
*/
func (o *traefikSdk) GetHttpRouters() (*resty.Response, error) {
	//log.Println("GetRoutes ", deviceId)
	resp, err := o.restyGet(HTTP_ROUTERS, nil)
	if err != nil {
		return nil, err
	}
	//log.Println("GetRoutes ", resp)
	return resp, nil
}