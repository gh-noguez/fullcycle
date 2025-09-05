# API Gateway.

O API Gateway é um componente fundamental em arquiteturas de microsserviços e sistemas distribuídos. Ele funciona como um ponto de entrada único para todas as requisições que chegam à sua aplicação, centralizando o tráfego e aplicando regras e políticas antes de encaminhá-lo aos serviços internos.

##### Como Funciona
Imagine o API Gateway como a portaria de um condomínio fechado. Ninguém consegue chegar à porta do seu apartamento sem passar pela portaria primeiro. O API Gateway faz o mesmo, interceptando todas as requisições e tomando decisões com base em suas regras.

##### Funções Principais
O API Gateway tem a capacidade de entender as necessidades de cada sistema e aplicar lógicas complexas, o que o torna uma ferramenta poderosa para a resiliência.

1. Centralização: Todas as requisições passam por ele, permitindo um controle unificado.

2. Autenticação e Autorização: Ele pode validar a autenticação do cliente (ex: validar um token JWT) antes mesmo que a requisição chegue ao serviço, economizando recursos computacionais.

3. Rate Limiting: O API Gateway pode aplicar regras de Rate Limiting para proteger os serviços, barrando requisições que excedem um limite predefinido. Isso evita a sobrecarga e permite que você configure limites diferentes para clientes com prioridades distintas.

4. Health Check: Ele pode monitorar a saúde dos serviços de forma ativa. Se um serviço está lento ou com problemas, o API Gateway pode parar de encaminhar o tráfego para ele e redirecionar as requisições para outras instâncias saudáveis ou retornar um erro instantâneo para o cliente.

5. Transformação de Dados: O API Gateway pode modificar requisições e respostas. Por exemplo, ele pode converter um formato de dados (como XML para JSON) ou adicionar e remover headers, simplificando a lógica de negócio dos serviços.

6. Roteamento Dinâmico: Ele direciona as requisições para o serviço correto com base na URL ou outros critérios. Por exemplo, uma requisição para /produtos pode ser encaminhada para o serviço de produtos, enquanto uma para /usuarios vai para o serviço de usuários.

##### Ferramentas de API Gateway
Uma das soluções mais famosas é o Kong, que utiliza o Nginx por trás dos panos. O Kong é conhecido por sua flexibilidade e a vasta quantidade de plugins que oferece para autenticação, Rate Limiting, Health Check e outras funcionalidades.

##### Benefícios
O API Gateway centraliza preocupações de arquitetura que, de outra forma, teriam que ser implementadas em cada serviço individualmente. Isso torna os serviços mais simples, aumenta a resiliência e a segurança, e melhora a performance geral do sistema.
