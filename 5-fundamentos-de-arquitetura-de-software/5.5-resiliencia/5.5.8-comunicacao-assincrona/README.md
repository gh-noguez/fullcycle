# Comunicação assíncrona.

##### Trabalhando de Forma Assíncrona
Trabalhar de forma assíncrona é uma estratégia de resiliência que previne a perda de dados e permite que o sistema processe mais requisições do que sua capacidade imediata permitiria. O conceito é simples e intuitivo, pois replicamos uma situação do dia a dia: as filas.

##### A Fila da Vida Real
Imagine a fila de um banco ou supermercado. Mesmo que o número de caixas seja menor que o número de clientes, o sistema não "perde" os clientes. Em vez disso, eles entram em uma fila e aguardam sua vez de serem processados.

##### Aplicações Assíncronas
Muitas vezes, uma aplicação pode receber mais requisições do que consegue processar instantaneamente. Se a aplicação trabalha de forma síncrona, ela pode travar ou perder as requisições que chegam quando está sobrecarregada.

- Assincronicidade é Resiliência: Em um sistema assíncrono, as requisições que chegam não são processadas na hora. Elas são enviadas para uma fila e processadas quando o sistema está pronto.

- Benefícios:
    - Prevenção de perda de dados: Nenhum dado é perdido, pois a requisição fica na fila até ser processada.
     - Eficiência de recursos: O sistema consegue lidar com picos de tráfego, pois não precisa ter uma capacidade de processamento imediata para todas as requisições. Com um recurso menor, ele pode dar vazão a mais requisições ao longo do tempo.
    - Desacoplamento: A parte da aplicação que envia a requisição não precisa esperar uma resposta imediata.

##### O Papel do Message Broker
Para implementar a comunicação assíncrona, usamos um Message Broker (Corretor de Mensagens).

- Como funciona:

    1. O emissor envia uma mensagem para o Message Broker.

    2. O Message Broker armazena a mensagem de forma segura.

    3. O serviço que precisa processar a mensagem lê a fila do Message Broker quando está disponível.

- Exemplos: RabbitMQ, Apache Kafka, Amazon SQS.

É crucial entender a fundo o Message Broker escolhido, pois a forma como ele é configurado (ex: garantias de entrega e recebimento) pode afetar a resiliência. Uma configuração incorreta pode, ironicamente, levar à perda de dados.

Trabalhar de forma assíncrona é uma mentalidade de arquitetura que permite criar sistemas mais robustos e resilientes, que não perdem dados e conseguem lidar melhor com a variação do tráfego.