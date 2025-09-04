# Proxy reverso.

#### Proxy Reverso
O conceito de proxy reverso é fundamental para a arquitetura de sistemas modernos. Ele desempenha um papel crucial na segurança, performance e escalabilidade das aplicações.

##### O que é um Proxy?
- Um proxy tradicional age em nome do cliente (usuário). Ele recebe a requisição do usuário e a encaminha para o servidor de destino. É frequentemente usado em redes corporativas para filtrar o acesso a sites, por exemplo.

##### O que é um Proxy Reverso?
- Um proxy reverso faz o oposto. Ele age em nome dos servidores.
- É um servidor que fica na frente dos servidores web e encaminha as requisições dos clientes (navegadores) para o servidor correto.
- O cliente faz a requisição para o proxy reverso, e este, com base em regras, a encaminha para um dos servidores que estão "atrás" dele.

##### Como um Proxy Reverso Funciona?
O proxy reverso recebe todas as requisições e, com base em informações como a URL, os cabeçalhos HTTP e outros dados, decide para qual servidor a requisição deve ser enviada.

- Exemplo:
    - Um cliente acessa a.com.br. O proxy reverso recebe a requisição, analisa a URL e a encaminha para o Servidor 1.
    - Outro cliente acessa b.com.br. O proxy reverso o encaminha para o Servidor 2.
    - O mesmo se aplica a diferentes serviços ou aplicações rodando em servidores distintos.

##### Funções do Proxy Reverso
- Balanceamento de Carga (Load Balancing): O proxy reverso pode distribuir as requisições de forma inteligente entre vários servidores que executam a mesma aplicação. Isso é essencial para a escala horizontal.

- Centralização da Segurança: Ele pode ser configurado com regras de segurança para barrar ataques como SQL Injection ou Cross-Site Scripting, agindo como um firewall de aplicação.

- Cache: O proxy reverso pode armazenar em cache o conteúdo estático (imagens, CSS, JavaScript), reduzindo a carga nos servidores de aplicação e melhorando o tempo de resposta.

- Compactação e Criptografia: Pode compactar respostas para economizar banda ou lidar com a criptografia SSL/TLS, aliviando essa carga dos servidores de aplicação.

##### Soluções de Proxy Reverso Comuns
- Nginx (Engine-X): O mais popular e amplamente utilizado no mercado. É conhecido por sua alta performance e flexibilidade.

- HAProxy: Focado em alta disponibilidade e balanceamento de carga, ideal para ambientes que exigem grande resiliência.

- Traefik: Um proxy reverso moderno e dinâmico, popular em ambientes de microsserviços e contêineres, pois ele se configura automaticamente com base na orquestração.

Embora existam outras ferramentas como API Gateways (ex: Kong), muitas delas são construídas sobre as bases de proxies como o Nginx. Por isso, a recomendação é aprender a configurar e trabalhar com o Nginx, o que já será de grande ajuda na sua carreira.