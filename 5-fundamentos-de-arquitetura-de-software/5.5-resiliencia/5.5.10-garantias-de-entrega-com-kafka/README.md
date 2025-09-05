# Garantias de entrega com Kafka.

##### Garantia de Entrega com Message Brokers
Garantir a entrega de mensagens é crucial para a resiliência de sistemas assíncronos. Se você não pode perder uma transação de alto valor, é fundamental entender as garantias que o seu message broker oferece.

A escolha da garantia de entrega está sempre ligada a um trade-off entre velocidade e segurança.

##### Os Níveis de Confirmação (ACK) no Kafka
Como exemplo, o Apache Kafka oferece diferentes níveis de confirmação de recebimento (ACK), mas a lógica se aplica a muitos outros brokers.

1. ACK 0 (Fire and Forget):
    - O que é: O produtor (quem envia a mensagem) não espera nenhuma confirmação de que o broker recebeu a mensagem. Ele simplesmente "dispara e esquece".
    - Vantagem: É a opção mais rápida, com a menor latência possível. Ideal para dados que podem ser perdidos sem causar grandes problemas, como a localização de um motorista de Uber a cada segundo.
    - Risco: Você não tem garantia de que a mensagem foi recebida, podendo haver perda de dados.

2. ACK 1 (Leader Confirmation):
    - O que é: O produtor espera uma confirmação de que o broker líder (o que recebeu a mensagem) a registrou.
    - Vantagem: Oferece um equilíbrio entre velocidade e garantia. É mais seguro que o ACK 0.
    - Risco: Se o broker líder cair logo após a confirmação, mas antes de replicar a mensagem para os outros brokers, a mensagem pode ser perdida.

3. ACK -1 ou ACK All (All Replicas Confirmation):
    - O que é: O produtor espera uma confirmação de que o broker líder recebeu a mensagem e que ela foi replicada para todas as outras réplicas no cluster.
    - Vantagem: É a opção mais segura, garantindo que a mensagem está altamente disponível e não será perdida, mesmo se um ou mais brokers falharem.
    - Desvantagem: É a opção mais lenta, pois a latência é maior devido ao tempo necessário para a replicação.

##### A Importância da Escolha
A escolha do nível de garantia de entrega deve ser uma decisão estratégica baseada na criticidade dos dados.
- Uma transação de um milhão de dólares exige o maior nível de garantia (ACK All), mesmo que isso signifique uma performance um pouco menor.
- Um fluxo de dados em tempo real onde a perda de informações é aceitável pode usar ACK 0 para maximizar a velocidade.

Para garantir a resiliência do seu sistema, você precisa entender o Message Broker em profundidade e saber qual nível de garantia se aplica a cada tipo de mensagem. Lembre-se, não há solução mágica, e cada decisão tem um custo-benefício que deve ser avaliado.