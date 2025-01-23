# Criando aplicação exemplo e imagem.

- Aplicação criada [aqui](server.go)

- Executando aplicação:
```bash
go run server.go
```

- Abrindo no navegador:
```bash
http://localhost:1078/
```

- Criando a imagem da aplicação:
```bash
docker build -t felipenoguez/hello-go .
```

- Executando imagem:
```bash
docker run --rm -p 1078:1078 felipenoguez/hello-go
```

- Enviando a imagem para o Dockerhub:
```bash
docker push felipenoguez/hello-go
```