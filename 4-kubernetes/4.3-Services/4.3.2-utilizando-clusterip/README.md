# Utilizando ClusterIP.

- Comandos utilizados.


- Aplicando yaml do service:
```bash
kubectl apply -f service.yaml
```

- Listar services:
```bash
kubectl get services
kubectl get svc
```

- Novamente para teste, vamos usar o port-foward para acessar a aplicação (porta usada na aplicação é a 1078):
```bash
kubectl port-forward svc/goserver-service 8000:1078
```

```bash

```

http://localhost:8000/


- Kubernetes possui um sistema de DNS integrado que permite a resolução de nomes para serviços e pods dentro do cluster. Esse sistema simplifica a comunicação entre os componentes da aplicação, substituindo endereços IP complexos por nomes de host mais fáceis de usar.