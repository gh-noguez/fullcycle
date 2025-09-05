# Garantias de entrega com retry.

##### Garantia de Entrega e Retries
No universo de sistemas distribuídos, é crucial garantir que as mensagens enviadas cheguem ao seu destino. Para isso, são usadas políticas de retries (retentativas), que são uma forma de resiliência.

##### O Problema do Retry Linear
A forma mais simples de retry é tentar novamente após um período fixo. Por exemplo, você manda uma requisição, e se ela falhar, espera 2 segundos e tenta de novo. O problema é que, se múltiplos sistemas estiverem fazendo o mesmo, todos tentarão novamente no mesmo momento, sobrecarregando ainda mais o serviço que já estava falhando. Isso cria um ciclo vicioso que impede a recuperação do sistema e pode até piorar o cenário.

##### Exponential Backoff
Para evitar o retry simultâneo, uma técnica mais eficaz é o Exponential Backoff. Em vez de esperar um tempo fixo, o tempo de espera aumenta exponencialmente após cada falha:

- 1ª tentativa: Falhou.

- 2ª tentativa (retry): Espera 2 segundos. Falhou.

- 3ª tentativa (retry): Espera 4 segundos. Falhou.

- 4ª tentativa (retry): Espera 8 segundos.

Essa abordagem dá um tempo maior para o sistema se recuperar. No entanto, se vários clientes usam o mesmo algoritmo de backoff, eles ainda podem se sincronizar e bombardear o serviço ao mesmo tempo.

##### Exponential Backoff com Jitter
A estratégia mais robusta é adicionar um Jitter (um ruído, um desvio) ao backoff exponencial.
- O que é: Em vez de esperar exatamente 2, 4, 8 segundos, o algoritmo adiciona um pequeno valor aleatório.
- Exemplo: A próxima tentativa pode ser em 2.1 segundos, a outra em 2.05, a outra em 2.25, e assim por diante.
- Benefício: Essa pequena aleatoriedade quebra a sincronia das retentativas. As requisições chegam em momentos ligeiramente diferentes, o que aumenta a probabilidade de uma delas ser bem-sucedida, pois o sistema de destino não é sobrecarregado por um pico repentino de requisições.

##### Conclusão
Políticas de retry são essenciais para a resiliência de um sistema, mas precisam ser implementadas com inteligência. Usar Exponential Backoff com Jitter é a melhor forma de aumentar as chances de sucesso das retentativas, garantindo que o sistema não perca dados e que os serviços tenham a oportunidade de se recuperar. Não se deve "chutar" os valores de retentativa no código; eles devem ser baseados em uma lógica que distribua a carga de forma mais inteligente.