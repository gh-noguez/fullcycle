# Diferença entre Port e targetPort.

- Utilizando o targetPort (que é a porta do container/app), podemos redirecionar ao acessar a porta 80 por exemplo, ele irá redirecionar para a porta 1078 (que é a porta que estou utilizando na apĺicação).

- Comando usados em aula:

- Após alteração do yaml para 80 em "port", foi aplicado o yaml:
```bash
kubectl apply -f service.yaml
```

- Aplicando port-forward
```bash
kubectl port-forward svc/goserver-service 8000:80
```

```bash

```

- Aogra, ao acessar a porta 8000 (definida no comando acima) do computador é apontado para a porta 80 do service que por sua vez, aponta para a porta 1078, que é o targetPort que foi definido na aplicação/container.