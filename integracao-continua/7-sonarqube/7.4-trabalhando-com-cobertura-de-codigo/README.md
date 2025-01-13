# Trabalhando com cobertura de código.

### Adicionando sonar-project.properties.
- Nesta aula foi adicionado `sonar-project.properties`, onde foi adicionada toda a configuração necessária para enviar para o Sonarqube analisar o código, executando somente o comando:
```bash
sonar-scanner
```

### Pré requisitos:
- Testes na aplicação e gerar o output com a cobertura do código.
- Executar testes.
```bash
go test
```

```bash
go test -v
```

- Gerando output com a cobertura de código em Golang para ser analisado pelo sonarqube;
```bash
go test -coverprofile=coverage.out
```

- Após estes passos, é só executar o comando `sonar-scanner`, aguardar a execução e verificar na interface do Sonarqube o resultado.