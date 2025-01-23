# Criando cluster multi-node

- Primeiro vamos remover o cluster criado na aula anteior, para isto, vamos utilizar o comando para verificar os clusters que temos na máquina `kind get clusters` e `kind delete clusters nome-do-cluster`:

- Buscando os clusters:
```bash
kind get clusters
```

- Apagando cluster:
```bash
kind delete clusters kind
```

- Após feita a remoção do cluster, vamos criar um arquivo de configuração, o [kind.yaml](k8s/kind.yaml) e executar o comando:

```bash
kind create cluster --config=k8s/kind.yaml --name=fullcycle
```

```bash
kubectl cluster-info --context kind-fullcycle
```

```bash
kubectl get nodes
```

```bash

```


- Link de referência com a documentação:

https://kind.sigs.k8s.io/docs/user/configuration/#nodes

