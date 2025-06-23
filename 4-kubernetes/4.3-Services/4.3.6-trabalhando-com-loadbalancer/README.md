# Trabalhando com Loadbalancer

- O serviço LoadBalancer no Kubernetes é um tipo de serviço que permite expor uma aplicação para acesso externo através de um balanceador de carga provisionado pelo provedor de nuvem. Ele cria um endereço IP externo único e estável para que os clientes possam acessar a aplicação sem precisar conhecer os endereços IP individuais dos pods.

- Comandos usados em aula:

- Aplicando yaml:
```bash
kubectl apply -f service.yaml
```


```bash
kubectl get svc
```
