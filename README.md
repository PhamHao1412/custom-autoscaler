# 🚀 Custom Autoscaler (Go + Prometheus + Grafana)

A lightweight, educational **autoscaler simulator** written in Go — designed to mimic how a Kubernetes Cluster Autoscaler works, 
but simplified for local demonstration with **Prometheus metrics** and **Grafana dashboards**.

---

## 🌟 Features

- 🧠 **Multi-metric decision logic**: CPU, Memory, and Response Time (ms)
- ⚙️ **Dynamic scaling simulation** using a mock cloud provider
- 📊 **Prometheus integration** for metrics scraping
- 📈 **Grafana dashboard** for real-time visualization
- 🌐 **Optional REST API** to query and control scaling manually

---

## 🧩 Architecture Overview

```markdown
flowchart TD
    A["Autoscaler (Go)"] -->|"Expose /metrics"| B["Prometheus"]
    B -->|"Scrape every 5s"| A
    B --> C["Grafana Dashboard"]
    A -->|"REST API"| D["User / External Tools"]
```

---

## 🏗 Project Structur“pe

```
custom-autoscaler/
├── cmd/
│   └── main.go
├── internal/
│   ├── autoscaler/
│   ├── metrics/
│   ├── cloud/
│   └── app/
├── grafana/
│   └── dashboards/autoscaler.json
├── prometheus/
│   └── prometheus.yml
├── docker-compose.yml
├── Dockerfile
├── Makefile
└── README.md
└── .env

```

---

## 📊 Grafana Dashboard Preview

![Grafana Dashboard Screenshot](./docs/image/grafana_dashboard.png)

> The dashboard visualizes live metrics for **CPU Usage**, **Memory Usage**, **Response Time**, and **Node Count**.

---

## ⚙️ .env example

```.env
AUTOSCALER.INTERVAL_SECONDS=5
AUTOSCALER.SCALE_UP_CPU=80
AUTOSCALER.SCALE_DOWN_CPU=30
AUTOSCALER.SCALE_UP_MEM=85
AUTOSCALER.SCALE_DOWN_MEM=40
AUTOSCALER.SCALE_UP_RESP_TIME=400 # ms
AUTOSCALER.SCALE_DOWN_RESP_TIME=150 # ms
AUTOSCALER.MIN_NODES=1
AUTOSCALER.MAX_NODES=5
AUTOSCALER.PROVIDER=mock
AUTOSCALER.PROMETHEUS_PORT=2112
```

---

## 🧠 How It Works

1. Every interval, the autoscaler:
   - Collects system metrics (CPU, Memory, Response Time)
   - Evaluates the decision logic:
     ```go
     if CPU>80 || Mem>85 || RT>400 -> scale up
     if CPU<30 && Mem<40 && RT<150 -> scale down
     else no-op
     ```
   - Updates the node count in the mock cloud provider  
   - Exposes metrics to Prometheus via `/metrics`

2. Prometheus scrapes the metrics and Grafana visualizes them in real-time.

---
## 🖥 Running Locally

### **Prerequisites**
- **Docker** & **Docker Compose** installed
- Ports **2112**, **9090**, and **3000** must be free

---

### 🐳 **Run the Monitoring Stack**

```bash
make up
```

This command launches the full monitoring environment, including:
- 🧠 **Custom Autoscaler** (Go service)
- 📊 **Prometheus** (metrics collector)
- 📈 **Grafana** (metrics visualization)

Once everything is up, you can access:

| Service | URL | Description |
|----------|-----|-------------|
| Autoscaler | [http://localhost:2112/metrics](http://localhost:2112/metrics) | Exposes Prometheus metrics |
| Prometheus | [http://localhost:9090](http://localhost:9090) | Explore & query raw metrics |
| Grafana | [http://localhost:3000](http://localhost:3000) | Dashboard UI *(login: admin / admin)* |

---

### 📈 **Setup Grafana Dashboards**

Once Grafana is running:

1. **Login** to Grafana at [http://localhost:3000](http://localhost:3000)  
   → Default credentials: `admin / admin`

2. **Add Prometheus as a Data Source**
    - Go to **⚙️ → Data Sources → Add data source**
    - Choose **Prometheus**
    - In the “URL” field, enter:
      ```
      http://prometheus:9090
      ```
    - Click **Save & Test**

3. **Import Dashboard**
    - Go to **+ → Import**
    - Upload or paste the contents of your local file:
      ```
      grafana_dashboard.json
      ```
    - Select **Prometheus** as the data source, then click **Import**

4. **Refresh Dashboards**
    - After import, open each dashboard and click **Refresh 🔄**
    - You should start seeing real-time metrics from the autoscaler:
        - CPU usage
        - Memory usage
        - Response time
        - Node count
        - Scaling actions

---

### 🧰 **Other Make Commands**

| Command | Description |
|----------|-------------|
| `make docker-build` | Build Docker image manually |
| `make logs` | View live container logs |
| `make down` | Stop all running containers |
| `make clean` | Remove containers, networks, and volumes |

---

### 📦 **Example Output**

```
[AUTOSCALER] CPU=82.33 | MEM=69.10 | RT=435.22 | Nodes=3 | Action=scale-up
📊 Prometheus metrics server running at http://localhost:2112/metrics
```

---

## 🪪 License

MIT © 2025  
Developed by Hao Pham
