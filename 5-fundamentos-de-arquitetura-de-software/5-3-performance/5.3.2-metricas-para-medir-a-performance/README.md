# Métricas para medir a performance.

#### Perspectivas de Arquitetura: Performance
Performance é um tema adorado por desenvolvedores, mas que muitas vezes não é bem compreendido ou medido.

##### O que é Performance?
Performance é o desempenho de um software para completar um "workload" (uma carga de trabalho ou tarefa). Para avaliá-la, é essencial usar dados e, mais importante, comparar a performance do software consigo mesmo ao longo do tempo.

##### Métricas de Performance
Existem dois indicadores principais para analisar a performance de um software:

1. Latência (Response Time):

    - É o tempo que leva desde a requisição de um cliente até a resposta do software.

    - Geralmente medida em milissegundos.

    - A latência é impactada por três fatores:

        -   Tempo de processamento da aplicação: Quão eficiente é o seu código.

        - Rede: A qualidade e a distância da conexão entre o cliente e o servidor.

        - Chamadas externas: O tempo de resposta de APIs de terceiros (ex: API dos Correios, banco de dados).

2. Throughput:

    - Mostra a quantidade de requisições que um software consegue processar em um determinado período.

    - Aumentar o throughput significa permitir que o software lide com mais requisições simultâneas.

##### Performance vs. Escalabilidade
- Performance e escalabilidade são conceitos distintos.

- Um sistema pode ser performático, mas não escalável (ex: rápido para 100 usuários, mas falha com 1000).

- Um sistema pode ser escalável, mas não performático (ex: suporta 1000 usuários, mas com um tempo de resposta lento para todos).

##### Como Aumentar a Performance
Para melhorar a performance de um sistema, você deve:

- Diminuir a latência: Reduzir o tempo de resposta da sua aplicação.

- Aumentar o throughput: Capacitar o software para lidar com mais requisições simultâneas.

O throughput e a latência estão intimamente ligados. Se a latência é alta, provavelmente há conexões "presas" na aplicação, o que diminui a capacidade de lidar com novas requisições (throughput). Se você conseguir gerenciar essas duas variáveis, estará no caminho certo para otimizar seu sistema.