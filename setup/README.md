# Connection

Host IP: `10.11.70.179/24, 192.168.64.9/24`

Raspberry Pi IP: `10.11.70.107/24`

## Setup Cluster

### WireGuard on On Host

```bash
multipass launch --name k3s --memory 8G --disk 40G --cpus 10
multipass shell k3s
```

```bash
apt-get update && apt install wireguard-tools resolvconf
```


### WireGuard and K3s On Raspberry Pi

add the followingt to `/boot/cmdline.txt`, then reboot

```text
cgroup_memory=1 cgroup_enable=memory
```

Setup WireGuard

```bash
curl -O https://raw.githubusercontent.com/angristan/wireguard-install/master/wireguard-install.sh

chmod +x wireguard-install.sh
./wireguard-install.sh # follow the prompt to create a client
```

Start the K3s master node

```bash
curl -sfL https://get.k3s.io | K3S_TOKEN=token INSTALL_K3S_EXEC="--advertise-address=10.66.66.1 --flannel-iface=wg0"  sh -
```


### WireGuard interface and K3s worker node on Host

Copy the WireGuard config to K3s node at `/etc/wireguard/wg0.conf ` and issue

```bash
wg-quick up wg0
```

bring up the K3s worker node

```bash
curl -sfL https://get.k3s.io | K3S_TOKEN=token K3S_URL=https://10.66.66.1:6443 INSTALL_K3S_EXEC="--node-ip=10.66.66.2 --flannel-iface=wg0" sh -
```

### Label worker node from Raspberry Pi (Master Node)

```bash
kubectl label nodes k3s type=worker
```

## Setup Shifu

From master node, issue:

```bash
kubectl apply -f https://raw.githubusercontent.com/Edgenesis/shifu/v0.24.0/pkg/k8s/crd/install/shifu_install.yml
```

### Deploy test Nginx if needed

```bash
kubectl run nginx --image=nginx -n deviceshifu --overrides='{"spec": { "nodeSelector": {"type": "worker"}}}'


kubectl run nginx1 --image=nginx -n deviceshifu --overrides='{"spec": { "nodeSelector": {"node-role.kubernetes.io/master": "true"}}}'
```

## Raspberry Pi Connection

TODO: add Raspberry Pi GPIO connection diagram
