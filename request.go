package infisical

import "net/http"

func CreateApiRequestWithAuthInterceptor(BaseURL, ServiceToken string) *http.Client {
	client := &http.Client{}
	client.Transport = &Interceptor{
		BaseURL:      BaseURL,
		ServiceToken: ServiceToken,
	}

	return client
}

type Interceptor struct {
	BaseURL      string
	ServiceToken string
}

func (i *Interceptor) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "InfisicalNodeSDK")
	req.Header.Set("Authorization", "Bearer "+i.ServiceToken)

	// set url
	req.URL.Scheme = "https"
	req.URL.Host = i.BaseURL

	return http.DefaultTransport.RoundTrip(req)
}
