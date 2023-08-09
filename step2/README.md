# Step 2: Connect the displacement sensor

## Architecture

```mermaid
graph RL
    subgraph k8s[Kubernetes Cluster]
        subgraph master[Master Node IP:10.11.70.107]
            sd(shifu controller)
            dss(deviceshifu-sensor)
        end
        subgraph worker[Worker Node IP:192.168.64.9]
            dsc(deviceshifu-rtspcamera)
        end
    end
    master <-- WireGuard --> worker
    c(camera IP:10.11.70.180) <-- Ethernet --> dsc
    s(displacement sensor) <-- I2C --> dss
```

## build and scp the image

```bash
docker build -t displacement-sensor:v0.0.1 .
docker save displacement-sensor:v0.0.1 > displacement-sensor.tar.gz
scp displacement-sensor.tar.gz  raspberrypi@10.11.70.107:
```

## import the image

```bash
sudo ctr images import displacement-sensor.tar.gz
```

## deploy the sensor

```bash
kubectl apply -f sensor.yaml
```

## check in browser

```bash
http://10.11.70.107:{PORT}/sensor
```
