# Entendendo Service Accounts


- Uma conta de serviço é um tipo de conta não humana que, no Kubernetes, fornece uma identidade distinta em um cluster do Kubernetes. Aplicação Pods, sistema componentes e entidades dentro e fora do cluster podem usar um cluster específico As credenciais da ServiceAccount para identificar como a ServiceAccount. Essa identidade é útil em várias situações, incluindo autenticação para o servidor da API ou implementação de políticas de segurança baseadas em identidade.

- Fonte/Documentação: https://kubernetes.io/docs/concepts/security/service-accounts/

- Comandos usados em aula.

- Listar serviceaccounts:
```bash
kubectl get serviceaccounts
```

- Listar os pods para pegar o nome e acessar:
```bash
kubectl get po
```

- Analisando as informações do Pod, podemos ver o caminho onde fica armazenado o certificado, namespace e token:
```bash
kubectl describe po goserver-7f87c8459c-crtxn
```

- Acessando Pod para analisar arquivos do path abaixo:
```bash
kubectl exec -it goserver-7f87c8459c-crtxn /bin/sh
```

- Acessando path com certificado, namespace e token da serviceaccount default:
```bash
cd /var/run/secrets/kubernetes.io/serviceaccount/
```

- Listando e abrindo os arquivos:
```bash
ls
cat ca.crt
cat namespace  
cat token
```