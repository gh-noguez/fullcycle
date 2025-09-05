# Proteger e ser protegido.

#### Estratégias de Resiliência: Proteger e Ser Protegido
A resiliência de um sistema não depende apenas do seu próprio código, mas de como ele interage com o ecossistema de serviços ao seu redor.

##### A Filosofia de "Proteger e Ser Protegido"
Em uma arquitetura distribuída (como microsserviços), é essencial pensar em duas coisas:

1. Proteger a sua aplicação: Garantir que ela continue operando mesmo quando outros serviços falham.

2. Proteger a aplicação dos outros: Evitar que sua aplicação cause falhas em outros serviços.

Um sistema não pode ser egoísta. Não adianta continuar enviando requisições para um serviço que já está lento ou falhando. Isso só piora a situação, causando um efeito dominó onde a lentidão se propaga e pode derrubar todo o ecossistema.

##### O Efeito Dominó
O "efeito dominó" ocorre quando a lentidão de um serviço se espalha, fazendo com que todos os serviços que dependem dele também fiquem lentos e, eventualmente, falhem.

- Cenário: Sistema A chama B, que chama C. Se o Sistema C fica lento, o B fica travado esperando a resposta, e o A também fica travado esperando a resposta do B. O acúmulo de requisições pode derrubar os três serviços.

##### Estratégia de Auto-preservação
- Sistema lento é pior que sistema fora do ar: Um serviço lento prende recursos e pode causar o efeito dominó. Por isso, em certas situações, é melhor um serviço retornar um erro 500 (indicando que está indisponível) do que ficar lento. Isso avisa aos outros sistemas que há um problema, permitindo que eles adotem estratégias de fallback (plano B).

- Harmonia entre sistemas: Assim como você deve evitar sobrecarregar um serviço que está falhando, você também espera que os outros serviços não façam isso com a sua aplicação. A resiliência é um esforço coletivo.

##### Conclusão:
A resiliência não é apenas sobre programar bem em uma linguagem específica. É sobre adotar uma mentalidade e conceitos de arquitetura que garantem que seu sistema possa operar de forma eficiente em um mundo distribuído, onde as falhas são inevitáveis.