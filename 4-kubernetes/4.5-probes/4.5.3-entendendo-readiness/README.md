# Entendendo readiness.

- readinessProbe: Readiness probes determinam quando um contêiner está pronto para começar a aceitar tráfego. Isso é útil ao aguardar que um aplicativo execute tarefas iniciais demoradas, como estabelecer conexões de rede, carregar arquivos e aquecer caches.

- Fonte: https://kubernetes.io/docs/concepts/configuration/liveness-readiness-startup-probes/

- Nesta aula, foi alterada a lógica da aplicação para simular a subida dela com um atraso de 10 segundos.
- Ocorreu 2 erros nesta aula:
1. Gerei uma imagem com nome do endpoint errado e mesmo corrigindo a imagem, enviando para o Dockerhub, ao executar o deployment no Kubernetes a imagem se mantinha a antiga, pois usei a mesma tag, sendo assim, a solução foi utilizar `imagePullPolicy: Always`, para forçar a baixar a imagem novamente do registry.
2. Erro `connection refused` ao utilizar o readinessProbe, pois eu configurei a porta do "localhost:8000:, que eu estava fazendo os testes no navegador com os redirecionamentos da `service.yaml` mais `kubectl port-forward svc/goserver-service 8000:80`, pois eu aponto a porta `8000` local para a porta `80` da minha `service` e a `80` da minha `service` tem a `targetPort: 1078`. Mas o readinessProbe opera diretamente no `kubelet` (agente do Kubernetes dentro do pod), o readinessProbe se conecta com o IP interno do Pod e à `containerPort` especificada na própria definição da probe. Resumindo, alterei a porta do `readinessProbe` para `1078` e após a simulação da aplicação subindo durante 10s, o Pod ficou com status `READY`.

- Comandos utilizados em aula:

- Listando os Pods:
```bash
kubectl get po
```


- Analisando eventos do pod com problema, onde é possível identificar o erro no `liveness`:
```bash
kubectl describe pod goserver-7ff7f7d64b-zb84n
```

- Apagando deployment:
```bash
kubectl delete deployment goserver
```

- Aplicando yaml do deployment e exibindo o Pod subindo:
```bash
kubectl apply -f deployment-v2.yaml && watch -n1 kubectl get pods
```

- Criando túnel de encaminhamento entre porta local e o serviço. Acessando no navegador `http://localhost:8000/` para teste:
```bash
kubectl port-forward svc/goserver-service 8000:80
```

- Recriando a imagem com a tag v8 sem cache:
```bash
docker build -t felipenoguez/hello-go:v8 . --no-cache
```

- Enviando a nova imagem para o Dockerhub:
```bash
docker push felipenoguez/hello-go:v8
```

- Criando container local com a imagem para validar que as configurações estavam corretas:
```bash
docker run -it --rm felipenoguez/hello-go:v8 /bin/sh
```


