package traefik

import "github.com/go-resty/resty/v2"

/*
*/
func (o *traefikSdk) GetApiOverview() (*resty.Response, error) {
	//o.debugPrint()
	resp, err := o.restyGet(OVERVIEW, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
*/
func (o *traefik) GetApiRawData() (*resty.Response, error) {
	resp, err := o.restyGet(RAW_DATA, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
*/
func (o *traefik) GetApiVersion() (*resty.Response, error) {
	resp, err := o.restyGet(API_VERSION, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
