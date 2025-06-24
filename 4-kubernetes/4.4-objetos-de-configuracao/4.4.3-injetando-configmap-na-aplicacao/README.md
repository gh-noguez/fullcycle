# Injetando configmap na aplicação.

- Nesta aula ocorreu erro no apontamento do path para o arquivo txt, onde foi realizado debug para idientificar o problema e corrigi-lo.

- Comandos utilizados em aula:


- Aplicando yaml do configmap:
```bash
kubectl apply -f configmap-family.yaml
```

- Gerando novaa imagem com as alterações, criando endpoint e método para leitura do arquivo dinamicamente:
```bash
docker build -t felipenoguez/hello-go:v5 .
```

- Enviando imagem para o Dockerhub:
```bash
docker push felipenoguez/hello-go:v5
```

- Aplicando yaml do deployment com as alterações de volumes:
```bash
kubectl apply -f deployment-v2.yaml
```

- Criando túnel de encaminhamento entre porta local e o serviço. Acessando no navegador `http://localhost:8000/` para teste:
```bash
kubectl port-forward svc/goserver-service 8000:80
```

- Acessando container para debugar:
```bash
kubectl exec -it goserver-7b68b647b7-5l9wz /bin/sh
```

- Visualizando logs do container:
```bash
kubectl logs goserver-7b68b647b7-5l9wz -f
```
