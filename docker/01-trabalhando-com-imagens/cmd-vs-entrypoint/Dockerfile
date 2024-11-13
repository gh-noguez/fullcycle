FROM ubuntu:latest

ENTRYPOINT [ "echo", "Hello" ]

CMD [ "World" ]

# FROM ubuntu:latest: Esta linha especifica a imagem base que será usada para construir a nova imagem. Neste caso, está sendo usada a imagem oficial do Ubuntu na versão mais recente (latest). Isso significa que a imagem sendo construída terá o Ubuntu como base.

# ENTRYPOINT ["echo", "Hello"]: Define o ponto de entrada padrão para a imagem. O ENTRYPOINT especifica o comando que será executado quando um contêiner baseado na imagem for iniciado. Neste caso, o ponto de entrada é o comando echo seguido pela string "Hello". Se você executar um contêiner dessa imagem sem especificar um comando, ele imprimirá "Hello".

# CMD ["World"]: Fornece argumentos padrão para o ponto de entrada especificado pelo ENTRYPOINT. Neste caso, "World" é passado como argumento para o comando echo. Se você executar um contêiner dessa imagem e especificar um comando, o argumento padrão será substituído. Se nenhum comando for especificado ao iniciar o contêiner, o CMD será usado como argumento padrão para o ENTRYPOINT.

# Comando executado no terminal:
# docker build -t felipenoguez/hello:latest .