# Instalando metrics-server

 - O Servidor Métricas é uma fonte escalável e eficiente de métricas de recursos de contêiner para o Kubernetes tubulações de autoescalonamento integradas.

- Fonte para download do yaml: https://github.com/kubernetes-sigs/metrics-server

- Fonte da issue com a solução para uso em laboratório sem TLS: https://github.com/kubernetes-sigs/metrics-server/issues/525

- Comandos utilizados em aula:

- Download do yaml e alterado o nome do arquivo para `metrics-server`
```bash
wget https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```


- Para o laboratório com o cluster K8s Kind, vamos utilizar o args `- --kubelet-insecure-tls` junto nos args do deployment que foi realizado o download

- Aplicando yaml para instalação do metrics-server.
```bash
kubectl apply -f metrics-server.yaml
```

- listando todos os serviços de API registrados no cluster Kubernetes e verificando a instalação do metrics-server.
```bash
kubectl get apiservices
```

![alt text](image.png)