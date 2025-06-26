# Combinando Liveness e Readiness.

- Durante a combinação de Liveness e Readiness, é necessário atenção para os tempos e/ou quantidade de threshold do liveness, pois dependendo do tempo que leva para o pod ficar pronto, por diversos motivos, o liveness por sua vez pode reiniciar o pod antes de ele estar pronto, desta forma criando um ciclo de falhas. Desta forma destaca-se perfeitamente o risco de um `livenessProbe` prematuro ou excessivamente zeloso interferir no ciclo de inicialização de um `readinessProbe` que ainda está aguardando a aplicação estar de fato apta a receber tráfego.
- Nesta aula, foi adicionado à aplicação um "ou" para gerar falha quando a duração exceder 20 segundos, assim vamos poder analisar outro caso, onde o pod vai deixar de estar com o status `READY` com o readiness e o liveness irá reinicia-lo, mais uma vez é preciso ter atenção com os tempos para que o liveness tenha tempo o suficinente para que o pod esteja `READY`.


- Comando usados em aula:

- Gerando nova imagem com a condicional para "quebrar" a aplicação após 20 segundos:
```bash
docker build -t felipenoguez/hello-go:v9 .
```

- Enviando imagem para o Dockerhub:
```bash
docker push felipenoguez/hello-go:v9
```


- Aplicando e analisando o estado `READY` e `RESTART`:
```bash
kubectl apply -f deployment-v2.yaml && watch -n1 kubectl get pods
```
