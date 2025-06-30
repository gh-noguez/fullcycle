# Criando volume persistente e montando.

### Introdução
- O gerenciamento de armazenamento é uma questão bem diferente do gerenciamento de instâncias computacionais. O subsistema PersistentVolume provê uma API para usuários e administradores que mostra de forma detalhada de como o armazenamento é provido e como ele é consumido. Para isso, nós introduzimos duas novas APIs: PersistentVolume e PersistentVolumeClaim.

- Os modos de acesso são:

1. ReadWriteOnce -- o volume pode ser montado como leitura-escrita por um nó único
1. ReadOnlyMany -- o volume pode ser montado como somente-leitura por vários nós
1. ReadWriteMany -- o volume pode ser montado como leitura-escrita por vários nós

- Fonte: https://kubernetes.io/pt-br/docs/concepts/storage/persistent-volumes/

- Comandos utilizados nesta aula:

```bash
kubectl get storageclass
```

```bash
kubectl apply -f pvc.yaml
```

- Listar volumes. Após aplicarmos o yaml do PVC, podemos ver o status do Volume como `Pending`:
```bash
kubectl get pvc
```

- Aplicando o yam do deployment com o apontamento para o volume:
```bash
kubectl apply -f deployment-v2.yaml
```

- Listar volumes. Após aplicado o yaml do deployment com as alterações para montar o volume podemos ver que o status do Volume é `BOUND`:
```bash
kubectl get pvc
```

- Agora, vamos listar o pod e acessa-lo com seu nome:
```bash
kubectl get po
kubectl exec -it goserver-7f87c8459c-45bj6 /bin/sh
```

- Após os passos anteriores, caso tenha ocorrido tudo certo, dentro do POD podemos ver o diretório que foi criado, acessando este diretório, vamos criar um arquivo para em seguida apagar o POD, para testar a montagem do PV:
```bash
ls
cd data
ls
touch teste.txt
```

- Apagando POD para que seja recriado, para o teste do PV:
```bash
kubectl delete goserver-7f87c8459c-45bj6
```

- Listando o Pod para pegar o novo nome e acessa-lo:
```bash
kubectl get po
kubectl exec -it goserver-7f87c8459c-crtxn /bin/sh
```

- Acessando o diretório e validando que o arquivo criado se mantem:
```bash
ls
cd data
ls
```