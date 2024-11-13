# Aula - Configurando app node com docker-compose

- Nesta aula, utilizamos o [docker-compose.yaml](docker-compose.yaml)  para subir o container.

- Comandos utilizados nesta aula:
```bash
docker-compose up -d
```
```bash
docker-compose up -d --build
```

```bash
docker exec -it app bash
```

```bash
docker compose build --no-cache
```

<!-- executados dentro do container node
npm init -y 
npm install -->
aqui ainda sigo com problema de não conseguir criar o projeto nem a node_modules pelo Dockerfile - revisar
Para o problema acima,identifiquei que com o volume montado, onde na raiz do projeto no meu host, não existia a pasta node_modules, o container subia executando normalmente os comandos e criando o projeto, mas, era sobrescrito pelo que havia na pasta "./node", que no caso, era nada. Para corrigir a falha, rodei os comandos "npm ini"  e "npm install express" dentro da pasta "./node".