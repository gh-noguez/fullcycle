# Aula - Criando banco de dados MySQL

- Nesta aula, utilizamos o [docker-compose.yaml](docker-compose.yaml)  para subir o container do banco MySQL.

- Comandos utilizados nesta aula:
```bash
docker-compose up -d
```

```bash
docker logs dbmysql
```

```bash
docker exec -it dbmysql mysql -uroot -p
```

```bash
show databases;
```