# Teste de stress do Fortio.

- Nesta aula foi realizado o teste de carga na aplicação Go com Fortio.


- Fonte da ferramenta: https://github.com/fortio/fortio
- Outra ferramenta para teste de carga mencionada em aula: https://k6.io/

- Comandos usados em aula:

- Aplicando deployment com recursos reduzidos para o teste de carga:
```bash
kubectl apply -f deployment-v2.yaml
```

- Criando Pod "temporário" para executar o teste de carga com Fortio:
```bash
kubectl run -it fortio --rm --image=fortio/fortio -- load -qps 800 -t 120s -c 70 "http://goserver-service/healthz"
```

- Listando HPA:
```bash
kubectl get hpa
```

- Monitorando HPA:
```bash
watch -n1 kubectl get hpa
```
