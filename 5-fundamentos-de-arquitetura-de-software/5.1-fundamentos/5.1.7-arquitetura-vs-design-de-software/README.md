# Arquitetura vs design de software.

#### Arquitetura vs. Design de Software: Uma Discussão
Este é um tema polêmico, com diferentes pontos de vista. Embora alguns considerem arquitetura e design de software a mesma coisa, há uma distinção importante.

#### Arquitetura de Software: Escopo Global
- Foca no escopo global do software, uma visão de mais alto nível.

- Envolve a componentização, a forma de comunicação entre esses componentes e as abstrações gerais.

- Pensa no que os componentes têm em comum e como padronizar seu uso.

- Objetivo primário: garantir que os atributos de qualidade (requisitos não-funcionais), restrições de alto nível e objetivos de negócio sejam atendidos pelo sistema.

    - Exemplo: Decidir usar OpenTelemetry para logs e métricas em todo o sistema é uma decisão arquitetural de alto nível.

#### Design de Software: Escopo Local
- Foca em um escopo mais local (baixo nível).

- Preocupa-se em como uma classe terá menos responsabilidades ou como um padrão (pattern) específico será implementado para facilitar uma estratégia.

- Exemplo: A implementação do código para o OpenTelemetry (como ele será escrito dentro de uma classe ou módulo) é uma decisão de design.

#### Relação entre Arquitetura e Design
- Um conceito importante: "Atividades relacionadas à arquitetura de software são sempre de design." (Elemar Jr.)

    - Isso significa que, ao definir a arquitetura (ex: um componente), o design desse componente é intrínseco. Arquitetura e design andam juntos.

- "Entretanto, nem todas as atividades de design são sobre arquitetura."

    - Decisões de design muito locais que não impactam a visão de alto nível (ex: refatorar um trecho de código, criar uma classe de abstração interna que não é visível externamente) não são consideradas arquiteturais.

- "Qualquer decisão de design que não tenha relação com esse objetivo [garantir atributos de qualidade, restrições de alto nível e objetivos de negócio] não é arquitetural."

- Decisões arquiteturais são visíveis fora de um componente. Se algo é uma decisão de design dentro de uma classe e não é visível ou não impacta o sistema externamente, não é arquitetural.

#### Exemplo Polêmico: Clean Architecture
- A Clean Architecture, proposta por Uncle Bob (Robert Martin), apesar do nome, possui muitas decisões que se assemelham mais a decisões de design do que de arquitetura em si.

- Há elementos arquiteturais, mas grande parte trata de detalhes de implementação e organização interna.

#### Conclusão
- É importante reconhecer que existe uma discussão e diferentes pontos de vista sobre a distinção entre arquitetura e design de software.

- Você está livre para formar sua própria opinião, mas o mais importante é saber que essa discussão existe.


- Citação, fonte: https://eximia.co/quais-sao-as-diferencas-entre-arquitetura-e-design-de-software/