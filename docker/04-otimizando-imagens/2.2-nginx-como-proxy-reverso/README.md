# Aula - Nginx como proxy reverso

- Comandos utilizados nesta aula:

```bash
docker build -t felipenoguez/nginx:prod .
```
```bash
docker build -t felipenoguez/laravel:prod . -f Dockerfile.prod
```
```bash
docker network list
```
```bash
docker network create laranet
```
```bash
docker run -d --network laranet --name laravel-container felipenoguez/laravel:prod
```
```bash
docker run -d --network laranet --name nginx -p 8080:80 felipenoguez/nginx:prod
```

## Debug:
- Como eu alterei o nome da aplicação de "laravel" para "exemplo-app", meu container não subia, desta forma fui buscar formar para debugar, onde usei o comando abaixo para ver a subida do container:
```bash
docker run -it --network laranet --name nginx -p 8080:80 felipenoguez/nginx:prod
```
- Erro:
- 2024/02/16 16:28:41 [emerg] 1#1: host not found in upstream "example-app" in /etc/nginx/conf.d/nginx.conf:14
nginx: [emerg] host not found in upstream "example-app" in /etc/nginx/conf.d/nginx.conf:14`

- Ao final, alterei o nome do container (durante o run) e no conf do Nginx para laravel-container, para ficar claro a referência.

- Outro problema que tive ao rodar os containers foi com o cache, então rodei so comando com a flag --no-cache, conforme abaixo:
```bash
docker build --no-cache -t felipenoguez/laravel:prod . -f Dockerfile.prod
```
```bash
docker build --no-cache -t felipenoguez/nginx:prod .
```