# Situações complexas e decisões de alto nível.

#### A Resiliência da Resiliência
Para encerrar o assunto, é crucial entender que a resiliência não é um conceito absoluto. Mesmo com todas as estratégias que discutimos, você sempre terá o risco de falhas maiores e mais complexas. A grande provocação é: como você constrói a resiliência da resiliência?

##### O Risco do "Single Point of Failure" (Ponto Único de Falha)
Toda a sua estratégia de resiliência pode se basear em um único componente, como um message broker (Kafka, RabbitMQ, SQS). Se esse componente crítico falhar, sua resiliência cai por terra. O que acontece se seu broker estiver fora do ar?

- Seu sistema vai travar?

- As mensagens serão perdidas?

- A aplicação inteira ficará indisponível?

Essas são perguntas essenciais. Se seu sistema não consegue nem iniciar sem se conectar ao message broker, ele tem um ponto único de falha.

##### Gerenciamento de Riscos e Custo
Resiliência é, no fundo, uma questão de gerenciamento de riscos. Cada nível adicional de resiliência custa tempo, dinheiro e esforço.

- 99.999% de disponibilidade tem um custo exponencialmente maior do que 99% de disponibilidade.

- A decisão sobre o nível de resiliência não é técnica, mas estratégica. Não cabe ao desenvolvedor decidir quanto a empresa está disposta a gastar para garantir a resiliência. Essa é uma decisão dos líderes (CTO, CEO).

##### Cenários de Alta Probabilidade vs. Baixa Probabilidade
- Cenários comuns: A falha de um serviço ou um pico de tráfego são cenários com alta probabilidade de ocorrência. O desenvolvedor deve ter a responsabilidade de implementar estratégias como retries e health checks.

- Cenários improváveis: A falha de uma região inteira de um provedor de nuvem (como a AWS) ou até a falha do próprio provedor são cenários de baixa probabilidade. No entanto, o impacto pode ser catastrófico. É por isso que algumas empresas investem em estratégias como o multi-cloud, mesmo que o custo seja mais alto.

O ponto principal é que existe um limite para a resiliência. A decisão de onde parar é estratégica e envolve balancear o custo e o esforço com o nível de risco que a empresa está disposta a assumir. O desenvolvedor deve se concentrar em implementar as estratégias de resiliência que protegem o sistema de falhas comuns, enquanto a liderança decide sobre os riscos maiores.

Espero que esta série de aulas tenha proporcionado uma visão clara e abrangente sobre performance, escalabilidade e resiliência. Estes são os pilares essenciais para a arquitetura de software moderna.