# Escala concorrência e paralelismo.

### Escala Vertical vs. Escala Horizontal
Estes conceitos estão diretamente ligados à capacidade computacional e à forma como uma aplicação pode lidar com mais carga.

1. Escala Vertical (Scale Up)
    - O que é: Aumentar a capacidade computacional da mesma máquina (ex: adicionar mais CPU, RAM ou disco).
    - Como funciona: Em vez de otimizar código ou algoritmos, você simplesmente aumenta os recursos do hardware para que a aplicação possa lidar com mais requisições.
    - Vantagens: Simples de implementar.
    - Limitações: Você está limitado aos recursos máximos de uma única máquina. É uma solução finita.

2. Escala Horizontal (Scale Out)
    - O que é: Aumentar o número de máquinas que rodam a aplicação.
    - Como funciona: Colocar um load balancer na frente, que distribui as requisições entre várias instâncias da aplicação.
    - Vantagens: Praticamente ilimitado, permitindo que o sistema cresça muito.
    - Limitações: Requer que a aplicação seja projetada para isso, geralmente de forma stateless (sem estado).

##### Concorrência vs. Paralelismo
Embora os termos sejam frequentemente usados de forma intercambiável, eles representam conceitos distintos. Uma frase de Rob Pike, um dos criadores da linguagem Go, resume bem a diferença:

`"Concorrência é sobre lidar com muitas coisas ao mesmo tempo. Paralelismo é fazer muitas coisas ao mesmo tempo."`

1. Concorrência
    - O que é: A capacidade de um sistema lidar com várias tarefas simultaneamente, mas não necessariamente executá-las no exato mesmo momento.
    - Exemplo: Uma pessoa atendendo a várias requisições: atende uma ligação, coloca em espera para responder um email, volta para a ligação, e assim por diante.
    - Em software: Uma única thread ou processo pode gerenciar múltiplas requisições, alternando entre elas. Isso é comum em arquiteturas não-bloqueantes.

2. Paralelismo
    - O que é: A capacidade de um sistema executar várias tarefas ao mesmo tempo, em diferentes processadores ou núcleos.
    - Exemplo: Uma pessoa fazendo várias tarefas ao mesmo tempo, como falar ao telefone enquanto digita e bebe água, pois ela tem recursos (mãos, boca) para realizar todas as ações simultaneamente.
    - Em software: Utilizar múltiplas threads (fios de processamento) para que cada requisição seja processada em um núcleo de CPU diferente.

##### Conexão com o Tema da Aula
A concorrência permite que um sistema lide com um grande número de requisições, enquanto o paralelismo permite processá-las em um tempo menor.
- Exemplo de web server:
    - Acesso Serial (Bloqueante): Um único worker processa cinco requisições, uma de cada vez. Se cada uma levar 10ms, o tempo total será de 50ms.
    - Acesso Paralelo (Não-Bloqueante): Com cinco workers (ou threads), todas as cinco requisições podem ser processadas ao mesmo tempo. O tempo total será de 10ms, pois as tarefas são executadas em paralelo.

##### Considerações sobre Threads
- Threads do sistema operacional: Abertura de uma nova thread para cada requisição pode consumir muitos recursos de memória (ex: 1MB por thread). Linguagens como Go usam "green threads" (gerenciadas pelo runtime da própria linguagem) que são mais leves (ex: 2KB por thread), permitindo lidar com um número muito maior de requisições simultâneas.
- Linguagens Não-Bloqueantes: A concorrência é o motivo do sucesso de linguagens como Node.js e frameworks como o Swoole do PHP, pois eles permitem lidar com muitas requisições de forma não-bloqueante.