# Grafana Cloud Prometheus Proxy

The Grafana Cloud Prometheus Proxy stores api key credentials so users in a private network can access Grafana Cloud Prometheus

Environment Variables:

* CORTEX_ADDRESS The Grafana.com https endpoint (usually ends in /api/prom)
* CORTEX_API_KEY The Grafana.com SSO API Key
* CORTEX_TENANT_ID Your 4-digit tenant ID for Grafana.com

TODO: 

* Add some templates for viewing these in a web browser