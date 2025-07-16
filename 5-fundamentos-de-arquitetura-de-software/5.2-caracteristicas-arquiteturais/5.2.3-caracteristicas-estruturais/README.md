# Características estruturais.

#### Características Arquiteturais: Lado Estrutural
Agora, vamos focar nas características estruturais, que estão mais ligadas ao processo de desenvolvimento e à flexibilidade da aplicação.

1. Configurabilidade
- O que é: A facilidade com que a aplicação pode ser configurada sem a necessidade de alterar o código-fonte.

- Exemplos:

    - Conexão com Banco de Dados: Usar variáveis de ambiente em vez de "hard-coded" no código.

    - Chaves de API de Gateways de Pagamento: Ser capaz de mudar chaves ou alternar entre diferentes gateways de forma simples.

- Teste: Se precisar mudar o código-fonte para a aplicação rodar em ambientes diferentes (produção, staging), ela não é configurável.

- Importância: Pensar na configurabilidade desde o início é um padrão essencial.

2. Extensibilidade
- O que é: A capacidade da aplicação de crescer e permitir que novas funcionalidades ou integrações sejam "plugadas" sem exigir grandes refatorações na base.

- Cenário: Adicionar uma nova gateway de pagamento (ex: trocar da gateway X para Y). Se for necessário mudar a estrutura principal da aplicação, ela não é extensível.

- Soluções: Trabalhar com interfaces, adaptadores e conceitos como camadas anticorrupção para desacoplar a aplicação de dependências externas.

- Benefícios: Trocar de banco de dados, message broker, ou adicionar novos módulos sem afetar a base do código.

- Sinal de alerta: Se é difícil adicionar novas features ou módulos, ou se refatorar é constante, a estrutura precisa de revisão.

3. Fácil Instalação / Deploy
- O que é: A facilidade de configurar e colocar a aplicação para rodar em diferentes ambientes (teste, staging, produção, máquina de outro desenvolvedor).

- Desafios e Soluções:

    - Padronização do ambiente: Containers (Docker) são a melhor forma atual de garantir que a aplicação rode com o mesmo ambiente e kernel, evitando problemas de "funciona na minha máquina".

    - Configurabilidade: Aplicações difíceis de configurar são difíceis de instalar.

    - Dependências complexas:

        - Ex: Elasticsearch, Kafka, RabbitMQ.

        - Estratégia: Definir se a aplicação cria recursos (índices, tópicos, filas) automaticamente ou se eles são criados manualmente na infraestrutura.

- Importância: Facilita o deploy e a replicação de ambientes.

4. Reutilização de Componentes
- O que é: A capacidade de usar o mesmo código ou bibliotecas em diferentes partes do sistema ou em diferentes sistemas.

- Contextos:

    - Sistemas Monolíticos: É mais fácil reutilizar componentes internamente (bibliotecas de validação, frameworks).

    - Sistemas Distribuídos (Microsserviços): Pode levar a equipes criando soluções duplicadas para o mesmo problema.

- Solução: Criar verticais ou times dedicados para manter bibliotecas compartilhadas, garantindo que todos os produtos da empresa possam utilizá-las.

- Benefícios: Evita duplicação de esforços e facilita a manutenção.

5. Internacionalização
- O que é: Preparar o software para ser usado em diferentes idiomas, culturas e regiões.

- Dificuldades:

    - Front-end: Alterar o idioma pode desconfigurar layouts, e a cultura do usuário impacta o design.

    - Back-end: Preocupações com:

        - Moeda: Moeda base, conversão (automática ou manual), diferentes formas de pagamento (parcelamento, recorrente) e políticas de preço por região.

        - Fuso Horário: Como datas e horas são exibidas e manipuladas.

- Importância: Levantar os pontos de impacto da internacionalização desde o início evita problemas futuros.

6. Fácil Manutenção
- O que é: A capacidade de corrigir bugs e adicionar novas features de forma rápida e simples.

- Desafio: Simplificar soluções complexas. Quanto mais simples o software, mais fácil a manutenção.

- Práticas:

    - SOLID: Aplicar os princípios de design orientado a objetos.

    - Camadas e Interfaces: Usar adaptadores e interfaces para evitar acoplamento.

    - Design Patterns: Utilizar padrões de projeto.

    - Testes: Essencial! Sistemas sem testes são extremamente difíceis de manter e corrigir bugs, além de dificultar a adição de novas funcionalidades com segurança.

- Pontos de atenção: Facilidade para adicionar novas features e corrigir bugs.

7. Portabilidade
- O que é: A capacidade de o sistema ser menos dependente de fornecedores (vendors) e permitir a troca de tecnologias sem grandes impactos.

- Exemplos:

    - Mudar de banco de dados (nunca é "tranquilo", mas deve ser possível sem grandes refatorações de código).

    - Trocar de ferramentas de observabilidade (ex: Elastic Stack para New Relic ou DataDog). OpenTelemetry facilita essa portabilidade.

    - Mudar de gateway de pagamento.

- Importância: Reduz o aprisionamento tecnológico (vendor lock-in).

8. Observabilidade
- O que é: A capacidade de entender o que está acontecendo dentro do sistema em tempo de execução, facilitando a identificação e resolução de problemas.

- Objetivo: Conseguir suportar a aplicação ativamente, detectando problemas antes mesmo que o cliente os perceba.

- Componentes:

    - Logs: Como logar e centralizá-los.

    - Debug: Técnicas e ferramentas de depuração.

    - Benchmarks: Marcar momentos e medir desempenho interno.

    - Spans: Rastreamento distribuído.

    - Métricas: Coleta e análise de dados de desempenho.

- Dica: Focar na observabilidade e consolidar logs em um único padrão para facilitar a operação.

Checklist: Ao arquitetar uma aplicação, use essas características como um checklist. Se alguma não estiver sendo focada, vale a pena parar e pensar se ela já não está sendo tratada por padrão ou se precisa de atenção intencional.