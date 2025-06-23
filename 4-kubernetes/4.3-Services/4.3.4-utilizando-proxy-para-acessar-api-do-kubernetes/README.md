# Utilizando proxy para acessar api do kubernetes.

- Comando utilizados:

- Comando utilizado para criar um servidor proxy HTTP local que encaminha suas solicitações entre sua máquina e o plano de controle do K8s, permitindo que você acesse o servidor da API do K8s sem expô-lo diretamente.
```bash
kubectl proxy --port=8088
```

- URLs acessadas:
[http://localhost:8088](http://localhost:8088)
[http://localhost:8088/api/v1/namespaces/default/services/goserver-service](http://localhost:8088/api/v1/namespaces/default/services/goserver-service)
[http://localhost:8088/apis/apps/v1](http://localhost:8088/apis/apps/v1)
[http://localhost:8088/apis/apps/v1/deployments](http://localhost:8088/apis/apps/v1/deployments)


- [Reference guide da api do Kubernetes](https://kubernetes.io/docs/reference/).
