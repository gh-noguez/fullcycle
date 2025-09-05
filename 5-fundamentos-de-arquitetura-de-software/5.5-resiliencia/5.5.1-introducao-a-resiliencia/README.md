# Introdução à resiliência.

#### Perspectivas de Arquitetura: Resiliência
Depois de falar sobre performance e escalabilidade, é a vez da resiliência, um conceito crucial para a arquitetura de software que, muitas vezes, é mal compreendido.

##### O que é Resiliência?
A resiliência é a capacidade de um sistema se adaptar a falhas de forma intencional. A ideia é que, quando algo dá errado, o software não simplesmente "quebra", mas sim "dobra", encontrando uma forma de continuar funcionando, mesmo que de forma parcial.

- O sistema vai falhar: É uma verdade absoluta. As falhas podem ser causadas por bugs, problemas de rede, falhas de hardware ou indisponibilidade de serviços de terceiros (servidores de CEP, gateways de pagamento, etc.).

- Estratégia intencional: A resiliência não acontece por acaso. Você precisa planejar o comportamento do sistema para quando uma falha ocorrer. Se a sua aplicação não tem um plano B, ela não é resiliente.

- Minimização de riscos: Adotar estratégias de resiliência ajuda a minimizar a perda de dados e transações importantes para o negócio. Por exemplo, se uma gateway de pagamento está fora do ar, a aplicação precisa ter uma forma de não perder a venda, como processar o pagamento depois.

##### Resiliência na Prática
Pense em cenários de falha e como o seu sistema pode reagir:

- Serviço de CEP fora do ar: O sistema de cadastro de usuário simplesmente falha ou permite que o usuário continue e preencha o endereço manualmente mais tarde?

- Gateway de pagamento offline: A transação é cancelada ou o sistema armazena o pedido e tenta processar o pagamento novamente quando o serviço voltar ao normal?

Em resumo, resiliência é a arte de ter um plano B, C e D. É a capacidade do software de se adaptar quando o inesperado acontece, garantindo a melhor experiência possível para o cliente.