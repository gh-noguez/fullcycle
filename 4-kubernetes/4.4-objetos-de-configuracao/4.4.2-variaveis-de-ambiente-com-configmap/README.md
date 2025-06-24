# Variáveis de ambiente com configmap.

- Em implantações do Kubernetes, o `valueFrom: configMapKeyRef:` e o `envFrom: configMapRef:` são duas maneiras de injetar dados de um ConfigMap em variáveis de ambiente de um container, mas com propósitos e níveis de granularidade diferentes. `valueFrom: configMapKeyRef:` permite selecionar uma chave específica dentro de um ConfigMap e mapeá-la para uma variável de ambiente, enquanto `envFrom: configMapRef:` injeta todas as chaves e valores de um ConfigMap como variáveis de ambiente. Há dois deployments com o exemplo de uso.


- Comandos utilizados em aula.


- Aplicando yaml com o configmap:
```bash
kubectl apply -f configmap-env.yaml
```

- Aqui temos 2 exemplos de uso do configmap no deployment:
```bash
kubectl apply -f deployment.yaml
kubectl apply -f deployment-v2.yaml
```

- Criando túnel de encaminhamento entre porta local e o serviço. Acessando no navegador `http://localhost:8000/` para teste:
```bash
kubectl port-forward svc/goserver-service 8000:80
```

- Acessando container para validar as alterações:
```bash
kubectl exec -it goserver-7b68b647b7-5l9wz /bin/sh
```

- Listando configmaps:
```bash
kubectl get configmap
```

- Abrindo editor do yaml com o configMap para validar as configurações:
```bash
kubectl edit configmap goserver-env
```

- Listando deployments:
```bash
kubectl get deployments
```

- Abrindo editor do yaml com o deployments para validar as configurações:
```bash
kubectl edit deployment goserver
```