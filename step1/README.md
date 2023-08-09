# Step 1: Deploy the RTSP Camera

## Architecture

```mermaid
graph RL
    subgraph k8s[Kubernetes Cluster]
        subgraph master[Master Node IP:10.11.70.107]
            sd(shifu controller)
        end
        subgraph worker[Worker Node IP:192.168.64.9]
            dsc(deviceshifu-rtspcamera)
        end
    end
    master <-- WireGuard --> worker
    c(camera IP:10.11.70.180) <-- Ethernet --> dsc
```

## Deploy the deviceShifu

```bash
kubectl apply -f camera.yaml
```

## Check the NodePort of the deviceShifu

```bash
kubectl  get svc -n deviceshifu
```

## Open in browser

```text
192.168.64.9:{PORT}
```

## (Optional)Port-forward to view in browser

```text
kubectl port-forward -n deviceshifu svc/deviceshifu-rtspcamera 30080: --address 0.0.0.0
localhost:30080
```
