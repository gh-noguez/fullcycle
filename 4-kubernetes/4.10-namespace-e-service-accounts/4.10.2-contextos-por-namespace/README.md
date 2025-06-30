# Contextos por namespace.

- Comandos usados em aula.

- Aplicando yaml de exemplo.
```bash
kubectl apply -f deployment.yaml
```

- Criando namespace para dev:
```bash
kubectl create ns dev
```

- Criando namespace para prod:
```bash
kubectl create ns prod
```

- Aplicando deployment no namespace especificado:
```bash
kubectl apply -f deployment.yaml -n=dev
```

- Aplicando deployment no namespace especificado:
```bash
kubectl apply -f deployment.yaml -n=prod
```

- Busca do namespace por nome:
```bash
kubectl get ns prod
```

- busca de Pod pelo namespace:
```bash
kubectl get po -n=dev
```

- busca de Pod pelo namespace:
```bash
kubectl get po -n=prod
```

- Listar Pods pela label:
```bash
kubectl get po -l app=server
```

- Exibe as configurações que estão definidas atualmente:
```bash
cat ~/.kube/config
```

- Exibe as configurações atuais do cluster:
```bash
kubectl config view
```

- Exibe o contexto atual:
```bash
kubectl config current-context
```

- criando contexto para uso em dev:
```bash
kubectl config set-context dev --namespace=dev --cluster=kind-fullcycle --user=kind-fullcycle
```

- criando contexto para uso em prod:
```bash
kubectl config set-context prod --namespace=prod --cluster=kind-fullcycle --user=kind-fullcycle
```

- Define o contexto padrão a ser usado:
```bash
kubectl config use-context dev
```

- Exibe o contexto atual após aplicar o contexto padrão como dev:
```bash
kubectl config current-context
```
