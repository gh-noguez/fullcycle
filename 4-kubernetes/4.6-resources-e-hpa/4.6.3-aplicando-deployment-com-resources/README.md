# Aplicando deployment com resources

- Nesta aula foi aplicado o deployment com `resources` (`request` e `limits`). No arquivo de deployment, eu estava com a porta errada no liveness, então acessei o pod e identifiquei o problema para corrigir.

- Aplicando e analisando o estado `READY` e `RESTART`:
```bash
kubectl apply -f deployment-v2.yaml
```

- Analisando falha no pod:
```bash
kubectl logs goserver-5c55789bd7-md288
```
```bash
kubectl describe pod goserver-5c55789bd7-md288
```

- Visualizando recusrsos sendo usados pelo Pod:
```bash
kubectl top pods goserver-5c55789bd7-7xq2d
```