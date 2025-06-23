# Utilizando nodeport.

- Em Kubernetes, um serviço NodePort expõe um serviço (conjunto de pods) a partir do exterior do cluster, abrindo uma porta específica em cada nó do cluster e encaminhando o tráfego para o serviço interno. Essencialmente, o NodePort fornece uma maneira de acessar seus aplicativos de fora do cluster Kubernetes, utilizando um número de porta entre 30000 e 32767 em cada nó. 

- Comando utilizados em aula:

- Aplicando yaml com a service:
```bash
kubectl apply -f service.yaml
```

- listar serviços:
```bash
kubectl get svc
```

- Listar nodes e algumas especificações (como o ip):
```bash
kubectl get nodes -o wide
```

- Acessando aplicação no navegador:
```bash
ip:300001
```