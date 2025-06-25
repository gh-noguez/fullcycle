# Secrets e variaveis de ambiente.

- Comandos utilizados em aula.

- Gerando versão 6 da imagem:
```bash
docker build -t felipenoguez/hello-go:v6 .
```

- Enviando versão 6 para o Dockerhub
```bash
docker push felipenoguez/hello-go:v6
```

- Convertendo user para baase64:
```bash
echo "Noguez" | base64
```

- Convertendo password para base64:
```bash
echo "123456" | base64
```

- Aplicando yaml do secret:
```bash
kubectl apply -f secret.yaml
```

- Aplicando yaaml do deployment:
```bash
kubectl apply -f deployment-v2.yaml 
```

- Criando túnel de encaminhamento entre porta local e o serviço. Acessando no navegador `http://localhost:8000/` para teste:
```bash
kubectl port-forward svc/goserver-service 8000:80
```

- Listando pods:
```bash
kubectl get po
```

- Acessando pod:
```bash
kubectl exec -it goserver-6dc8b747f7-zl66w /bin/sh
```

- Após acessar o pod, é possível ver o vaslor da variável de ambiente que foi criada:
```bash
echo $USER
```

- Após acessar o pod, é possível ver o vaslor da variável de ambiente que foi criada:
```bash
echo $PASSWORD
```