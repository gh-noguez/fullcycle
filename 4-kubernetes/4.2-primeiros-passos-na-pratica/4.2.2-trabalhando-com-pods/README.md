# Trabalhando com PODs.

- Comandos usados nesta aula:

```bash
kubectl get nodes
```

```bash
kubectl apply -f 4.2-primeiros-passos-na-pratica/4.2.2-trabalhando-com-pods/pod.yaml
```

```bash
kubectl get pods
kubectl get pod
kubectl get po
```

```bash
kubectl port-forward pod/goserver 8001:80
```

```bash
kubectl delete pod goserver
```