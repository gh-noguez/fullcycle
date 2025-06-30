# Criando headless service.

- clusterIP: None cria um serviço headless, que não atribui um IP virtual e permite que os pods sejam acessados diretamente pelo nome DNS Exemplo de acesso: mysql-0.mysql-service-headless.default.svc.cluster.local onde "mysql-0" é o nome do pod, "mysql-service-headless" é o nome do serviç "default" é o namespace e "svc.cluster.local" é o domínio padrão do Kuberne Isso é útil para StatefulSets, onde cada pod tem um nome único e previsível e pode ser acessado diretamente por outros pods ou serviços Exemplo de acesso: mysql-0.mysql-service-headless.default.svc.cluster.local ou mysql-0.mysql-service-headless.svc.cluster.local:3306

- Comando usados em aula.

- Apagando statefulset para refazer o exemplo de aula:
```bash
kubectl delete statefulset mysql
```

- Aplicando yaml do statefulset:
```bash
kubectl apply -f statefulset.yaml
```

- Aplicando yaml do mysql-service-headless:
```bash
kubectl apply -f mysql-service-headless.yaml
```

- Listando pods:
```bash
kubectl get po
```

- Listando services para verificar se foi criado o service com `CLUSTER-IP` = `NONE`:
```bash
kubectl get svc
```

- Acessando pod para pingar os services:
```bash
kubectl exec -it goserver-7f87c8459c-crtxn /bin/sh
```

- Ping para os pods do mysql:
```bash
ping mysql-headless
ping mysql-0.mysql-headless
ping mysql-1.mysql-headless
ping mysql-2.mysql-headless
ping mysql-1.mysql-headless.default.svc.cluster.local
ping <pod-nam>.<service-name>.<namespace>.svc.cluster.local
```
