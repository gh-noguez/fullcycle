# Criando primeiro cluster com Kind.

- Criar um cluster:
```bash
kind create cluster
```

- Acessar o cluster no `context` `kind` com nome do cluster `kind`:
```bash
kubectl cluster-info --context kind-kind
```

- Verificar se o container do Kind foi criado:

```bash
docker ps
```
