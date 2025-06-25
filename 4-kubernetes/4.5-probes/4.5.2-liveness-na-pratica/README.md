# Liveness na prátiva..

- Nesta aula, foi adicionado o livenessprobe e alguns atributos de threshold.

- Comandos utilizados em aula:

- Aplicando yaml da aplicação e visualizando subida do pod com `watch`:
```bash
kubectl apply -f deployment-v2.yaml && watch -n1 kubectl get pods
```
