# Publicando imagem no Dockerhub

FROM nginx:latest

COPY html /usr/share/nginx/html

ENTRYPOINT [ "/docker-entrypoint.sh" ]
CMD [ "nginx", "-g", "daemon off;" ]

# Comando executado no terminal e sua descrição:
# docker build -t felipenoguez/nginx-noguez .

# docker build: Comando principal para construir uma imagem Docker.
# -t felipenoguz/nginx-noguez: Define a tag (-t) da imagem. Neste caso, a tag é felipenoguz/nginx-noguez. A tag é uma maneira de identificar e nomear uma versão específica da imagem. O latest é uma convenção frequentemente usada para indicar a versão mais recente.
# .: Indica o contexto de build. Neste caso, o ponto (.) representa o diretório atual. Isso significa que o Docker usará todos os arquivos e subdiretórios no diretório atual como o contexto para a construção da imagem.

# COPY: É a instrução do Dockerfile para copiar arquivos ou diretórios.
# html: É o caminho relativo no contexto de build que contém os arquivos ou diretórios a serem copiados.
# /usr/share/nginx/html: É o destino dentro da imagem onde os arquivos ou diretórios serão copiados.

# Acessando o container:
# docker run --rm -d -p 8080:80 felipenoguez/nginx-noguez

# Logando no DockerHub
# comando executado:
# docker login, após é só inserir usuário e senha

# Deslogando do DockerHub
#  docker logout

# Enviando a imagem para o DockerHub
# comando utilizado:
# docker push felipenoguez/nginx-noguez