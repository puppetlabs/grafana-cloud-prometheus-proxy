package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/puppetlabs/grafana-cloud-prometheus-proxy/lib/config"
)

// Routes is a struct which contains all http routes for this application
type Routes struct {
	Config config.Config
}

func (r *Routes) getHTTPResponse(url string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("Error(22): %s\n", err)
	}

	req.SetBasicAuth(r.Config.TenantID(), r.Config.APIKey())

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error(30): %s\n", err)
	}

	return resp
}

func (r *Routes) proxyJSON(url string, w http.ResponseWriter, req *http.Request) {
	resp := r.getHTTPResponse(url)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP Error: %d\n", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error(48): %s\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bodyBytes)
}

func (r *Routes) makePrometheusURL(endpoint string) string {
	return fmt.Sprintf("%s/api/v1/%s", r.Config.Address(), endpoint)
}

// Alerts proxies the prometheus alerts endpoint
func (r *Routes) Alerts(w http.ResponseWriter, req *http.Request) {
	url := r.makePrometheusURL("alerts")
	r.proxyJSON(url, w, req)
}

// Rules proxies the prometheus rules endpoint
func (r *Routes) Rules(w http.ResponseWriter, req *http.Request) {
	url := r.makePrometheusURL("alerts")
	r.proxyJSON(url, w, req)
}
