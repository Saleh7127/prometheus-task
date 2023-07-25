### Steps
- [make an api server with prometheus client library](https://prometheus.io/docs/tutorials/instrumenting_http_server_in_go/)
- dockerized the application and push to dockerhub
- deploy pod then service for that pod then servicemonitor for that service in cluster
- check graph and targets from prom-ui
- make grafana dashboard

## Resource
- https://www.youtube.com/watch?v=_NtRkBipepg&t=976s&ab_channel=ThatDevOpsGuy
- [part 1 to 3](https://www.youtube.com/watch?v=h4Sl21AKiDg&ab_channel=TechWorldwithNana)
- [Introduction to Kubernetes Monitoring And exploring Prometheus](https://blog.devops.dev/introduction-to-kubernetes-monitoring-and-exploring-prometheus-9bd0358ce504)
- [Prometheus, Kube State metrics and Integrate Grafana with Kubernetes](https://blog.devops.dev/part-1-setup-prometheus-kube-state-metrics-and-integrate-grafana-with-kubernetes-6c21f60d167f)