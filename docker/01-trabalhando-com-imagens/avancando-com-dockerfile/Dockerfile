# Avançando com Dockerfile

FROM nginx:latest

WORKDIR /app

RUN apt-get update && \
    apt-get install vim -y

COPY html /usr/share/nginx


# Comando executado no terminal e sua descrição:
# docker build -t felipenoguez/nginx-com-vim:latest .

# docker build: Comando principal para construir uma imagem Docker.
# -t felipenoguez/nginx-com-vi:latest: Define a tag (-t) da imagem. Neste caso, a tag é felipenoguez/nginx-com-vim:latest. A tag é uma maneira de identificar e nomear uma versão específica da imagem. O latest é uma convenção frequentemente usada para indicar a versão mais recente.
# .: Indica o contexto de build. Neste caso, o ponto (.) representa o diretório atual. Isso significa que o Docker usará todos os arquivos e subdiretórios no diretório atual como o contexto para a construção da imagem.

# COPY: É a instrução do Dockerfile para copiar arquivos ou diretórios.
# html: É o caminho relativo no contexto de build que contém os arquivos ou diretórios a serem copiados.
# /usr/share/nginx: É o destino dentro da imagem onde os arquivos ou diretórios serão copiados.

# Acessando o container:
# docker run -it <nome_ou_id_do_container> bash
