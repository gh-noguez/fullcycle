# Aula - Instalando framework em um container

Seguindo passo a passo da aula, onde será baixada uma imagem do PHP e configurado o container para rodar uma aplicação em PHP.

- Baixando/executando e acessando o container com PHP.
`docker run -it --name php php:7.4-cli bash`

- Atualizando pacotes:
apt-get update
etc . . .

- RODANDO DOCERKFILE
`docker build -t felipenoguez/laravel:v1 .`

- Executando o container
`docker run --rm -d --name laravel-example -p 8000:8000 <id_container>`

- Executando o container passando outra porta
`docker run --rm -d --name laravel-example -p 8001:8001 <id_container> --host=0.0.0.0 --port=8001`

- Imagem enviada para o Dockerhub:
`docker push felipenoguez/laravel:v1`