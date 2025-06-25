# Criando endpoint healtz.

- Nesta aula foi adicionado endpoint de healthz e gerada nova imagem da aplicação.

- Comandos utilizados em aula:

- Gerando imagem com a v7:
```bash
docker build -t felipenoguez/hello-go:v7 .
```

- Enviando imagem para o Dockerhub:
```bash
docker push felipenoguez/hello-go:v7
```

- Aplicando yam com deployment:
```bash
kubectl apply -f deployment-v2.yaml 
```

- Criando túnel de encaminhamento entre porta local e o serviço. Acessando no navegador `http://localhost:8000/` para teste:
```bash
kubectl port-forward svc/goserver-service 8000:80
```
