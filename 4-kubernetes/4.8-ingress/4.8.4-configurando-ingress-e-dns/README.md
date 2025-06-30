# Configurando ingress e DNS.

- Para esta aula, é necessário configurar o DNS para o ip externo do ingress controller. Wesley também apagou o service e recriou, com o type `ClusterIP`, pois agora o responsável será o ingress, não havendo a necessidade de um IP externo para a aplicação.

- Comandos usados em aula:


- Aplicando yaml com ingress:
```bash
kubectl aplly -f ingress.yaml
```

- Apagando service:
```bash
kubectl delete goserver-service
```

- Aplicando service com a alteração:
```bash
kubectl apply -f service.yaml
```
