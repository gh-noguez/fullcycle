# Criando service accounts e roles.

- Define o serviceAccountName para o Pod, que é o nome do ServiceAccount criado no arquivo security.yaml Isso permite que o Pod use as permissões definidas na Role e RoleBinding associadas ao ServiceAccount Isso é importante para que o Pod possa acessar os recursos permitidos pela Role, como pods, services e deployments no namespace especificado (prod) Isso é útil para isolar as permissões de acesso do Pod, garantindo que ele tenha apenas as permissões necessárias para sua operação e não mais do que isso, seguindo o princípio do menor privilégio Isso também facilita a gestão de permissões, pois as Roles e RoleBindings podem ser atualizadas ou removidas sem afetar o Pod diretamente e o Pod pode ser reimplantado com as novas ou removido se não for mais necessário Isso é especialmente útil em ambientes de produção, onde a segurança e o controle de acesso são críticos para evitar acessos não autorizados ou vazamentos de dados e para garantir que os Pods tenham apenas as permissões necessárias para suas operações e que não possam acessar recursos desnecessários ou sensíveis

- Comandos usados em aula:

- Listando serviceaccount no namespace prod:
```bash
kubectl get serviceaccounts -n prod
```

- Listando a api do Kubernetes:
```bash
kubectl api-resources
```

- Aplicando yaml com a serviceaccount:
```bash
kubectl apply -f security.yaml
```

- Novamente, listando serviceaccount no namespace prod e validando que foi criada uma nova serviceaccount:
```bash
kubectl get serviceaccounts -n prod
```

- Aplicando novamente, agora junto com a Rolebinding
```bash
kubectl apply -f security.yaml 
```

- Aplicando o deployment com a serviceaccount adicionada:
```bash
kubectl apply -f deployment.yaml
```

- Listar Pod para poder pegar o nome:
```bash
kubectl get pod -n prod
```

- Exibindo as configs do Pod:
```bash
kubectl describe po server-5dd59db49c-xbgnp -n prod
```

- Verificar o ServiceAccount no Pod:
```bash
kubectl get pod server-6799fdbc87-4c7b2 -o yaml -n prod | grep serviceAccount
```

- Após acessar o Pod, podemos testar a permissão, acessando e listando os Pods através da API (neste comando recebi erro, pois estou com deployment, serviceaccount, role, rolebind, tudo com namespace, então para que funcionar o culr, após ter inatalado o curl, usei o comando seguinte passando o nome do namespace):
```bash
curl -k https://kubernetes.default.svc/api/v1/pods --header "Authorization: Bearer $(cat /var/run/secrets/kubernetes.io/serviceaccount/token)"
```

- Comando usado no Pod para instalar o curl no Alpine Linux:
```bash
apk add curl
```

- Teste de permissão passando o namespace (com isto. é exibida a lista dos Pods):
```bash
curl -k https://kubernetes.default.svc/api/v1/namespaces/prod/pods --header "Authorization: Bearer $(cat /var/run/secrets/kubernetes.io/serviceaccount/token)"
```