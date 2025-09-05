# Circuit breaker.

#### Estratégias de Resiliência: Circuit Breaker
Para continuar a conversa sobre como proteger seus sistemas e os sistemas vizinhos, vamos falar sobre o Circuit Breaker (Disjuntor de Circuito), uma das estratégias de resiliência mais eficazes.

##### O que é o Circuit Breaker?
O circuit breaker é um padrão de design que protege um sistema ao negar requisições para um serviço que está com problemas, em vez de sobrecarregá-lo ainda mais. Ele funciona como um disjuntor elétrico: quando há uma sobrecarga, ele "abre o circuito" para evitar danos maiores.

##### Estados do Circuito
O circuit breaker opera em três estados:

1. Circuito Fechado (Closed):
    - É o estado normal. As requisições são encaminhadas e processadas sem restrições.

2. Circuito Aberto (Open):
    - Quando o sistema detecta que o serviço chamado está falhando (por lentidão, erros ou timeouts), ele "abre o circuito".
    - A partir desse momento, todas as novas requisições para esse serviço são negadas instantaneamente, retornando um erro 500 sem sequer tentar a conexão. Isso evita o efeito dominó e dá tempo para o serviço com problemas se recuperar.

3. Circuito Meio Aberto (Half-Open):
    - Após um período de tempo definido, o circuit breaker entra no estado "meio aberto".
    - Ele permite que um número limitado de requisições passe para o serviço que estava com problemas.
    - Se essas requisições de teste forem bem-sucedidas, o circuito é fechado novamente, e o tráfego volta ao normal.
    - Se falharem, o circuito retorna ao estado "aberto", e o tempo de espera antes da próxima tentativa aumenta (estratégia de backoff exponencial).

##### Implementação do Circuit Breaker
- Em código: Você pode implementar o circuit breaker diretamente no código da sua aplicação usando bibliotecas específicas.
- Na rede: Soluções mais modernas, como uma Service Mesh (ex: Istio, Linkerd), aplicam o circuit breaker na camada de rede. Isso permite que os desenvolvedores configurem a política de resiliência sem precisar escrever código, centralizando o controle e facilitando a gestão.

O circuit breaker é uma defesa poderosa para garantir que a sua aplicação não se torne um "agressor" que derruba outros sistemas em um ambiente distribuído. Ao negar requisições rapidamente, ele mantém a estabilidade do seu sistema e permite que as dependências se curem.