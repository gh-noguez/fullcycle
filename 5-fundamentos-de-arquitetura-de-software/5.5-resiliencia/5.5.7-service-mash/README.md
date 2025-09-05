# Service mesh.

##### Service Mesh (Malha de Serviços)
O Service Mesh (ou Malha de Serviços) é uma camada de infraestrutura que controla e gerencia a comunicação entre os serviços de uma aplicação. Ele está se tornando cada vez mais popular e é crucial para garantir a resiliência em ambientes de microsserviços.

##### O Conceito Principal
A ideia central do Service Mesh é usar proxies (sidecars) que rodam ao lado de cada serviço. Em vez de se comunicarem diretamente, os serviços A e B se comunicam através de seus respectivos proxies.

- O Serviço A envia uma requisição para o seu proxy.

- O proxy do Serviço A encaminha a requisição para o proxy do Serviço B.

- O proxy do Serviço B, então, a entrega para o Serviço B.

Dessa forma, toda a comunicação de rede é controlada e medida pelos proxies. Essa camada de infraestrutura oferece visibilidade e controle sobre o tráfego de forma centralizada.

##### Service Mesh e Resiliência
O Service Mesh abstrai a lógica de resiliência do código da aplicação, tratando-a como uma política de rede.

- Rate Limiting, Retries e Circuit Breaker: Em vez de programar essas lógicas em cada serviço, você as configura de forma centralizada na Service Mesh. O proxy de cada serviço se encarrega de aplicar as regras de Rate Limiting, Circuit Breaker e retries automaticamente, sem que o desenvolvedor precise se preocupar.

- Visibilidade: A Service Mesh fornece informações detalhadas sobre o tráfego de rede, como latência, erros e volume. Essa visibilidade permite que você tome decisões baseadas em dados, em vez de "chutar" os limites de requisição ou o número de retries.

##### Outros Benefícios da Service Mesh
- Criptografia Mútua (mTLS): O Service Mesh permite criptografar automaticamente a comunicação entre os serviços, garantindo que mesmo dentro da rede interna, as mensagens sejam seguras. Isso é feito de forma automatizada, gerenciando certificados e chaves, o que seria insano de fazer manualmente em um sistema com muitos microsserviços.

- Injeção de Falhas (Fault Injection): É possível simular falhas na rede (ex: latência artificial, erros) para testar a resiliência dos seus serviços. Isso ajuda a garantir que o sistema se comporte da maneira esperada em cenários de problemas.

A Service Mesh é um recurso de infraestrutura que abstrai complexidades de comunicação e resiliência, permitindo que os desenvolvedores se concentrem na lógica de negócio. Ela transforma a resiliência em uma política de rede, não em uma implementação manual em cada pedaço de código.