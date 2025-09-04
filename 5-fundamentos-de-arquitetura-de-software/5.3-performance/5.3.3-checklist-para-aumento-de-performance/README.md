# Checklist para aumento de performance.

#### Principais Razões para Baixa Performance
Muitos desenvolvedores lutam para melhorar a performance sem saber por onde começar. Aqui estão as principais razões para a baixa performance de um sistema e como combatê-las.

##### Causas da Baixa Performance
1. Processamento Ineficiente:
    - Algoritmos, código e a forma como a aplicação lida com as tarefas de forma geral podem ser ineficientes, gerando gargalos.

2. Recursos Computacionais Limitados:
    - A qualidade do hardware onde a aplicação roda afeta diretamente a performance.
    - Existe um trade-off entre custo e performance: hardware mais potente custa mais.

3. Processamento Bloqueante:
    - Quando uma requisição "trava" a aplicação enquanto é processada, ela impede que outras requisições sejam atendidas. Isso diminui o throughput.
    - Solução: Adotar arquiteturas não-bloqueantes (ex: o event loop do Node.js), que permitem lidar com múltiplas requisições simultaneamente, delegando tarefas demoradas a outras threads ou processos.

4. Acesso Serial aos Recursos:
    - Fazer requisições a recursos (como APIs ou bancos de dados) uma após a outra, de forma serial.
    - Isso reduz a capacidade de processamento simultâneo, impactando diretamente o throughput.
    - Com a maioria das máquinas modernas tendo múltiplos núcleos, o acesso serial não aproveita o hardware de forma eficiente.

##### Estratégias para Melhorar a Performance
1. Otimização de Recursos Computacionais:

    - Analisar o que está causando o gargalo:
        - CPU: Baixo poder de processamento.
        - Disco: Lento para operações de I/O.
        - Memória: Falta de RAM, levando ao uso lento da memória swap no disco.
        - Rede: A largura de banda limitada impede que todas as requisições cheguem à aplicação.

2. Otimização de Algoritmos e Queries:
    - Algoritmos: Evitar operações ineficientes (ex: N+1 problem).
    - Queries de Banco de Dados: Evitar SELECT * e usar índices adequadamente.
    - Overhead de Frameworks: Avaliar se um framework muito pesado não está prejudicando a performance.

3. Concorrência e Paralelismo:
    - Adotar linguagens ou estratégias que permitam concorrência (lidar com várias coisas ao mesmo tempo) e paralelismo (fazer várias coisas de forma conjunta em diferentes núcleos).
    - Exemplo: A linguagem Go, que cria uma nova thread (goroutine) para cada acesso, aumentando o throughput.

4. Otimização de Banco de Dados:
    - O banco de dados é um gargalo frequente.
    - Perguntas a fazer: O banco de dados é o certo para o problema? O modelo de dados está otimizado? Os índices estão criados corretamente?
    - Ferramentas: Usar explain nas queries para entender o plano de execução e ferramentas de APM (Application Performance Monitoring) para identificar queries lentas.

5. Caching:
    - O cache é fundamental para alta performance.
    - O que é: Armazenar o resultado de um processamento caro para reutilizá-lo em requisições futuras, em vez de processar tudo novamente.
    - Tipos: Pode ser em disco, em memória ou em servidores separados.
    - Importância: Diminui drasticamente o tempo de resposta e a carga no sistema.

##### Conclusão:
Para aumentar a performance, é preciso analisar cada um desses tópicos de forma individual para identificar a raiz do problema. Não adianta otimizar apenas um aspecto se o gargalo estiver em outro. Ter esse "checklist" em mente é o primeiro passo para repensar suas aplicações e torná-las mais eficientes.