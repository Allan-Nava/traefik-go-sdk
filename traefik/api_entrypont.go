package traefik

import "github.com/go-resty/resty/v2"
/*
*/
func(o *traefikSdk) GetEntrypoints() (*resty.Response, error){
	resp, err := o.restyGet(ENTRYPOINTS, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}