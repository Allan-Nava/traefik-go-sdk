package traefik


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
