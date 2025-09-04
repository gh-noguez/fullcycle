# Caching.

#### Cache: Aumentando a Performance do Sistema
O cache é uma ferramenta poderosa para aumentar a performance, diminuindo o tempo de resposta (response time) e aumentando o throughput. A ideia é armazenar resultados já processados para responder rapidamente ao usuário.

##### Tipos de Cache
Existem várias formas de implementar cache, dependendo da necessidade:
- Cache na Borda (Edge Computing):
    - O cache fica em servidores mais próximos do usuário, evitando que a requisição chegue ao servidor principal.
    - Exemplo: A Cloudflare armazena arquivos HTML, CSS, JavaScript e imagens, servindo-os a partir de um local mais próximo do usuário, o que é mais rápido e reduz o custo de infraestrutura.

- Cache de Dados Estáticos:
    - Armazenar em cache arquivos que não mudam frequentemente, como imagens e folhas de estilo (CSS). Isso economiza recursos do servidor principal e melhora a experiência do usuário.

- Cache de Página Web:
    - Armazenar o HTML de uma página inteira. Quando o usuário acessa, o servidor devolve o HTML já processado, sem precisar fazer consultas ao banco de dados ou executar lógica de negócio. Isso é útil para páginas que mudam com pouca frequência (ex: notícias, página de contato).

- Cache de Funções e Dados:
    - Armazenar o resultado de funções ou algoritmos que são caros de processar. Se os parâmetros de entrada não mudam, o resultado pode ser servido a partir do cache.
    - Também serve para armazenar resultados de consultas a bancos de dados, que são operações custosas.

- Cache de Objetos:
    - Armazenar objetos frequentemente utilizados, como os que mapeiam estruturas de banco de dados (ORM). Isso evita o reprocessamento da lógica de mapeamento, economizando tempo e recursos.

##### Cache Exclusivo vs. Compartilhado
Existem duas abordagens principais para o cache, cada uma com seus prós e contras:

- Cache Exclusivo (Local):
    - O que é: O cache é armazenado localmente em cada máquina/servidor.
    - Vantagens: Baixíssima latência, pois o acesso é local.
    - Desvantagens:
        - Duplicação: O cache é duplicado em todas as máquinas, ocupando mais espaço.
        - Problemas de sessão: Se um usuário fizer login em uma máquina, sua sessão não estará disponível em outra. Isso é problemático em sistemas que usam load balancers, pois o usuário pode ser direcionado para outra máquina e ter que fazer login novamente.

- Cache Compartilhado (Centralizado):
    - O que é: Um servidor de cache centralizado que todas as máquinas acessam.
    - Vantagens:
        - Sem duplicação: O mesmo cache é usado por todas as máquinas.
        - Sessões compartilhadas: A sessão de um usuário é armazenada no cache compartilhado, permitindo que ele acesse a aplicação de qualquer máquina sem ter que fazer login novamente.
    - Desvantagens:
        - Maior latência: Acesso ao cache é feito por rede, o que é mais lento que o acesso local.

- Ferramentas comuns: O Redis é uma das soluções de cache compartilhado mais populares atualmente, por ser um banco de dados em memória extremamente rápido. O Memcached é outra opção.

Aprender a usar o cache de forma eficaz é crucial para qualquer arquitetura de alta performance. É importante ter clareza sobre esses conceitos para tomar as melhores decisões no dia a dia.