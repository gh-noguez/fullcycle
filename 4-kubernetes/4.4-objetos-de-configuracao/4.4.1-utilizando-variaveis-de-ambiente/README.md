# Utilizando variáveis de ambiente.

- Nesta aula, copiei o Dcokerfile e server.go, realizei a alteração na aplicação Go adicionando as variáveis, crie nova imagem e subi para o Dockerhub.

- Comandos utilizados em aula:

- Gerando nova imagem com as alterações e com a tag v4:
```bash
docker build -t felipenoguez/hello-go:v4 .
```

- Enviando imagem para o Dockerhub:
```bash
docker push felipenoguez/hello-go:v4
```

- Aplicando deployment com as variáveis:
```bash
kubectl apply -f deployment.yaml
```

- Criando túnel de encaminhamento de porta local e o serviço. Acessando no navegador `http://localhost:8000/` para teste:
```bash
kubectl port-forward svc/goserver-service 8000:80
```