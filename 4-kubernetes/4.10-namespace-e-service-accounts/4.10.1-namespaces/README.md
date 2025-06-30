# Namespaces.

- Comandos usados em aula.

- Listar namespaces:
```bash
kubectl get ns
```

- Criando namespace:
```bash
kubectl create ns dev
```

- Aplicando yalm no namespace especificado:
```bash
kubectl apply -f <arquivo-deploy.yaml> -n=dev
```
