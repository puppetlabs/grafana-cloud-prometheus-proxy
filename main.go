package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/puppetlabs/grafana-cloud-prometheus-proxy/lib/config"
	"github.com/puppetlabs/grafana-cloud-prometheus-proxy/lib/routes"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	grafanaConfig := config.Config{}
	grafanaConfig.Init()

	r := routes.Routes{
		Config: grafanaConfig,
	}

	http.HandleFunc("/rules", r.Rules)
	http.HandleFunc("/alerts", r.Alerts)

	fmt.Println("Listening on port 8080")
	log.Fatal(s.ListenAndServe())

}
