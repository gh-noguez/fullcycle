# O problema do ReplicaSet.

- Comandos usados em aula:

- Criando nova imagem da aplicação:
```bash
docker build -t felipenoguez/hello-go:v2 .
```

- Enviando nova versão para o Dockerhub:
```bash
docker push felipenoguez/hello-go:v2
```

- Aplicando ReplicaSet, o problema é que mesmo alterando a versão da imagem para a tag "v2", após aplicar o yaml, nada acontece, sendo necessário apagar o pod para que ele crie um novo POD, ai sim, ele sobe com a versão correta da tag.
```bash
kubectl apply -f 4.2-primeiros-passos-na-pratica/4.2.4-o-problema-do-replicaset/replicaset.yaml
```

- Visualizando que nenhum POD sofreu alteração:
```bash
kubectl get po
```

- Aqui, podemos ver que a imagem segue sem alteração:
```bash
kubectl describe pod goserver-ckjdz
```

- Apagando POD para validar que a imagem foi alterada:
```bash
kubectl delete pod goserver-ckjdz
```

- Verificando a criação de um novo POD:
```bash
kubectl get po
```

- Vendo a descrição do POD criado e validando que a imagem foi alterada para a tag correta:
```bash
kubectl describe pod goserver-8p7vv
```