package traefik

import "github.com/go-resty/resty/v2"

// ExportConfiguration exports current configuration as JSON
func (o *traefikSdk) ExportConfiguration() (*resty.Response, error) {
	url := "/api/rawdata"
	return o.restyGet(url, nil)
}

// ImportConfiguration imports configuration from JSON
func (o *traefikSdk) ImportConfiguration(config interface{}) (*resty.Response, error) {
	fullURL := o.BaseUrl + "/api/config"
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(config).
		Post(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BackupConfiguration creates a timestamped configuration backup
func (o *traefikSdk) BackupConfiguration(timestamp string) (*resty.Response, error) {
	queryParams := map[string]string{
		"timestamp": timestamp,
	}
	return o.restyGet("/api/config/backup", queryParams)
}

// RestoreConfiguration restores configuration from a backup
func (o *traefikSdk) RestoreConfiguration(backup interface{}) (*resty.Response, error) {
	fullURL := o.BaseUrl + "/api/config/restore"
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(backup).
		Post(fullURL)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
