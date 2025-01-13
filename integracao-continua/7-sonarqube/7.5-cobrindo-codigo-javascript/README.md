# Cobrindo código Javascript.

### Nesta aula foi realizada a cobertura de testes com Javascript.

- Iniciando projeto Javascript:
```bash
npm init
```

- Instalando lib:
```bash
npm install jest @types/jest sonar-scanner --only-dev
```

- Executando testes:
```bash
npm run test
```





- Comando/token gerado pelo Sonarqube (obs.: agora não preciso executar todo este comando, apenas usar o token gerado, adicionando no properties e executar o comando `sonar-scanner`)

```bash
sonar-scanner \
  -Dsonar.projectKey=node-test \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://localhost:9000 \
  -Dsonar.token=sqp_5daaf4a1cfac06df50ad145fa5537ae8e4172a6d
```