# Requisitos arquiteturais.

#### Requisitos Arquiteturais (RA's)
Requisitos Arquiteturais (RA's) são requisitos, geralmente não-funcionais, que impactam diretamente a arquitetura do software.

#### Importância e Contexto Atual
- Antigamente, os RA's eram formalizados extensivamente em documentos (Excel, etc.) para guiar o planejamento da arquitetura.

- Hoje em dia, com o modelo de squads e equipes mais autônomas, essa formalização é rara de se ver no dia a dia, especialmente em arquitetura de software (mais comum em arquitetura de solução).

- É importante, no mínimo, conhecer o conceito de RA's.

#### Exemplos de Requisitos Arquiteturais
Os RA's estabelecem as "regras do jogo" para o desenvolvimento do software, abrangendo diversas categorias:

- Performance:

    - Tempo máximo de resposta para requisições (ex: não passar de 500 milissegundos).

    - Throughput (ex: uma máquina com "X" recursos deve suportar "Y" transações por segundo).

- Armazenamento de Dados:

    - Tecnologias de banco de dados a serem usadas (ex: DynamoDB, se houver contrato com AWS).

    - Regulações de localização dos dados (ex: dados da Europa devem estar em data centers na Europa).

    - Compliance com leis e regulamentações (ex: LGPD no Brasil).

- Escalabilidade:

    - Como o software vai escalar (horizontalmente, verticalmente).

    - Estratégias de load balancing (ex: round-robin, ou balanceamento entre máquinas com configurações diferentes).

- Segurança:

    - Certificações necessárias (ex: PCI para e-commerce com transações de cartão de crédito).

    - Comunicação criptografada (ex: entre microsserviços usando mTLS).

- Legais:

    - Requisitos para cumprir legislações específicas de cada país (ex: LGPD no Brasil para evitar vazamento de dados).

- Auditoria:

    - Tudo o que acontece no sistema deve ser logado.

    - Definição de onde e como os logs serão armazenados.

    - Tempo de retenção dos dados de auditoria.

- Marketing:

    - Disponibilidade para campanhas de marketing com picos de acesso.

    - Capacidade de rastrear origem de acessos e dispositivos.

    - Personalização de campanhas (chave de acesso mais próxima do usuário).

#### Como Obter os Requisitos Arquiteturais
- Os RA's podem envolver todas as áreas da empresa que o software irá servir.

- Tradicionalmente, a coleta é feita conversando com domain experts, executivos e usuários do software.

- Isso envolve fazer as perguntas certas (ex: "Qual o tempo de resposta esperado?", "O que o jurídico exige?").

- Na prática, muitas vezes esses requisitos são subentendidos ou implícitos, mesmo em grandes organizações.

- Quanto mais clareza nos RA's, mais claro será o processo de desenvolvimento, evitando ruídos e problemas no futuro.

- A coleta de RA's envolve muito diálogo e questionamento, sem "segredos" ou "esqueminhas".

Embora o módulo não se aprofunde em formalizações burocráticas, a compreensão dos RA's é fundamental para garantir um software robusto e adequado às necessidades do negócio.