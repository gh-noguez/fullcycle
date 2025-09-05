# Health check.

#### Estratégias de Resiliência: Health Check
Uma das formas mais básicas e cruciais de garantir a resiliência em sistemas distribuídos é através do health check (checagem de saúde). Sem ele, é impossível saber a condição de um sistema e tomar decisões inteligentes para protegê-lo e a todo o ecossistema.

##### O que é Health Check?
- O health check é um mecanismo que permite verificar a saúde de uma aplicação.

- Ele fornece os sinais vitais de um sistema, permitindo que outros serviços ou um load balancer determinem se a aplicação está saudável o suficiente para receber mais tráfego.

##### Por que um Health Check é Importante?
- Evitar o "chutar cachorro morto": Um sistema que está lutando para se manter no ar não deve receber mais requisições. O health check identifica essa situação e permite que o tráfego seja direcionado para outro lugar.

- Auto-recuperação (Self Healing): Ao desviar o tráfego de um sistema sobrecarregado, ele ganha tempo para processar as requisições pendentes e se recuperar. Quando volta a ficar saudável, pode voltar a receber tráfego.

##### Como Implementar um Health Check de Qualidade
- Não se contente com o básico: Um health check que retorna apenas um HTML estático (/health) não é suficiente, pois ele não reflete a saúde real da aplicação. Ele pode retornar 200, enquanto a aplicação está com problemas de conexão com o banco de dados ou com outros serviços.

- Seja estratégico: Um health check de qualidade deve verificar as dependências críticas da sua aplicação. Ele pode, por exemplo:

    - Fazer uma consulta simples no banco de dados.

    - Checar a latência das últimas requisições.

    - Verificar a conectividade com outros serviços externos.

- Retorno: Se as checagens forem bem-sucedidas, o health check deve retornar um código de sucesso (ex: 200). Se houver problemas, ele deve retornar um erro (ex: 500), sinalizando que o serviço está indisponível para uso.

A implementação de um health check bem pensado é o primeiro passo para garantir a resiliência de um sistema, evitando o efeito dominó e promovendo a auto-recuperação.