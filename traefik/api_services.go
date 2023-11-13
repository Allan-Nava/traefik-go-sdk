package traefik

import "github.com/go-resty/resty/v2"


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