# Criando statefulset.

- `Um StatefulSet executa um grupo de Pods e mantém uma identidade permanente para cada um deles. Isso é útil para gerenciar aplicativos que precisam de armazenamento persistente ou de uma identidade de rede estável e única.`

- Fonte/Documentação: https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/

- Comandos usados em aula.

- Aplicando yaml do statefulset:
```bash
kubectl apply -f statefulset.yaml
```

- Listando Pods e visualizando a criação um a um:
```bash
kubectl get po
```

- Escalando manualmente:
```bash
kubectl scale statefulset mysql --replicas=3
```
