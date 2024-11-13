# Aula - Otimizando imagens utilizando multistage building

- Construindo a imagem:
` docker build -t felipenoguez/laravel:prod -f Dockerfile.prod .`

- Com o multistage, foi conseguimos gerar uma imagem muito menor