# Criando Dockerfile

- Criado o [Dockerfile](../2-iniciando-com-ci/goapp/Dockerfile), agora vamos realizar um teste, movendo no terminal para o path onde está o Dockerfile e "buildando" a imagem com o comando:

```bash
docker build -t teste .
```

Testando o container:

```bash
docker run --rm teste
```
