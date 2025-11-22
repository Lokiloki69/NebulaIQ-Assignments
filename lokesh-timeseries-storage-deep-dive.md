# Timeseries Storage Deep Dive – Work Done Till Now (22 Nov)

This is only the progress part of the assignment, whatever I have completed till now. Most of the time went into fixing the environment because Docker was giving random errors (permissions, ports already used, config not mounted, Prometheus failing because of missing file, node exporter not reachable, scraping issues, Linux host IP, etc.). Finally everything is working now.

So this is everything I have done till now.

---

## 1. Observability Stack Setup (Prometheus + Grafana + Node Exporter)

I set up Docker Compose with Prometheus, Grafana and Node Exporter.  
This is the compose screenshot when everything finally started properly:

![DockerCompose](./screenshots/DockerCompose.png)

Prometheus home page:

![PrometheusUIHome](./screenshots/PrometheusUIHome.png)

Grafana login page also working:

![GrafanaLoginPage](./screenshots/GrafanaLoginPage.png)

Node exporter metrics (verified manually):

![NodeExporterMetrics](./screenshots/NodeExporterMetrics.png)

Prometheus targets after fixing hostname resolution, IP issues, and scrape configs:

![PrometheusTargets](./screenshots/PrometheusTargets.png)

---

## 2. Custom Go High-Cardinality Metrics Generator

I wrote a Go program that generates a high number of time series.  
Basically:

- 2000 dynamic endpoints  
- 2 HTTP methods  
- 3 status codes  
- 1000+ increments/sec  

This creates a **huge cardinality explosion** which is exactly what we need for the assignment.

Go program running in terminal:

![GoTerminalRunning](./screenshots/GoTerminalRunning.png)

Metrics visible in browser:

![BrowserShowing](./screenshots/BrowserShowing.png)

Prometheus scrapes the Go generator using the host IP because Linux Docker doesn’t support `host.docker.internal` by default.

---

## 3. Query Testing in Prometheus

Basic metric:


Screenshot:

![QueryResults](./screenshots/QueryResults.png)

---

## 4. General Monitoring Screens I captured

Just some random screens I saved while debugging everything:

![MonitorScreen](./screenshots/MonitorScreen.png)

![LogsView](./screenshots/LogsView.png)

---

## 5. Extra Screens (not directly part of this assignment but in the folder)

Redis CLI:

![RedisCLI](./screenshots/RedisCLI.png)

Redis Integration:

![RedisIntegrationInstalled](./screenshots/RedisIntegrationInstalled.png)

Redis Metrics:

![RedisMetricMonitor](./screenshots/RedisMetricMonitor.png)

Spike chart:

![Spike](./screenshots/Spike.png)

Trace View:

![TraceView](./screenshots/TraceView.png)

---

## What is Completed Till Now

- Docker observability stack completely functional  
- Prometheus scraping all 3 targets (prometheus, node_exporter, go-metrics)  
- High-cardinality metrics generator implemented in Go  
- Verified PromQL queries (`rate`, counters, etc.)  
- Screenshots captured  
- Environment is ready for the next part (TSDB internals, cardinality math, query optimization, tiered storage, etc.)

More work will continue tomorrow.

