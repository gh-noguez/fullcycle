# Aula - Node vs MySQL

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

Acessando o container do MySQL:
```bash
docker exec -it db bash
```

Acessando o banco com o comando e inserindo a senha definida no docker-compose:
```bash
mysql -uroot -p
```

- Com o comando "showd atabases", podemos ver que o banco "nodedb" já está criado de acordo com a definição do docker-compose.

- Agora, será selecionado a base para criarmos a tabela:

```bash
use nodedb;
```

Agora, será criada a tabela:

```bash
create table people(id int not null auto_increment, name varchar(255), primary key(id));
```

E com podemos ver a tabela com o comando:
```bash
desc people;
```

- Agora, vamos para o container do Node:
```bash
docker exec -it app bash
```

- Feito isto, vamos instalar o MySQL:
```bash
npm install mysql --save
```

- Com isto, podemos adicionar a conexão com o banco de dados no arquivo [index.js](node/index.js) e em seguida executar o comando no container do Node para gerar um insert no banco:

```bash
node index.js
```