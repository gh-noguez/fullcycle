# Rate limit.

#### Estratégias de Resiliência: Rate Limiting
Continuando com as estratégias para proteger e ser protegido em sistemas distribuídos, o Rate Limiting é uma ferramenta poderosa para controlar o fluxo de requisições.

##### O que é Rate Limiting?
O Rate Limiting protege seu sistema barrando o tráfego que excede a capacidade que ele foi projetado para suportar.

- Como funciona: Você define um limite máximo de requisições que seu sistema pode processar em um determinado período (ex: 100 requisições por segundo). Se esse limite for ultrapassado, o sistema retorna um erro 500 para as requisições excedentes.

- Objetivo: Garantir que o sistema opere com um nível de qualidade aceitável dentro de sua capacidade, evitando a sobrecarga e o colapso.

##### Implementando Rate Limiting com Inteligência
Uma implementação ingênua de Rate Limiting pode causar problemas. Se o sistema tem um limite global de 100 requisições/segundo, uma aplicação não crítica com um bug em um loop pode consumir todas as 100 requisições, impedindo que um cliente mais importante acesse o serviço.

- Solução: Implementar limites por cliente ou prioridade.

- Estratégia:

    - Identificar a prioridade: Determine quais clientes ou sistemas são mais críticos para o negócio.

    - Atribuir limites: Aloque uma quantidade específica de requisições para cada cliente com base em sua prioridade. Por exemplo, o cliente "Zezinho" (crítico) pode ter um limite de 60 req/s, enquanto outros clientes (menos críticos) compartilham o restante do limite.

- Benefício: Mesmo que um cliente menos crítico tente sobrecarregar o sistema, ele não conseguirá consumir todo o limite, garantindo que os clientes de alta prioridade tenham acesso garantido e que seu negócio continue funcionando de forma fluida.

O Rate Limiting é uma tática de defesa que garante a estabilidade do sistema, controlando o tráfego de entrada e priorizando os clientes mais importantes. No entanto, é importante entender os clientes e as prioridades do seu negócio para implementá-lo de forma eficaz.