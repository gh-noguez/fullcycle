# Escalabilidade.

#### Perspectivas de Arquitetura: Escalabilidade
É comum confundir performance com escalabilidade, mas, embora relacionados, são conceitos distintos.

##### O que é Escalabilidade?
Uma definição de Elemar Junior (criador do `arquiteturadesoftware.online`):

`"Escalabilidade é a capacidade de sistemas suportarem o aumento (ou a redução) dos workloads incrementando (ou reduzindo) o custo em menor ou igual proporção."`

Em outras palavras, é a capacidade de aumentar ou diminuir a capacidade de processamento do sistema de acordo com a demanda, controlando os custos.

##### Escalabilidade vs. Performance
- Performance: Foca em diminuir a latência e aumentar o throughput de uma única máquina.
- Escalabilidade: Foca em aumentar ou diminuir o throughput ao adicionar ou remover capacidade computacional.

##### Tipos de Escala
Existem duas formas principais de escalar um sistema:

1. Escala Vertical (Scale Up):
    - O que é: Aumentar os recursos computacionais de uma única máquina (ex: mais CPU, mais RAM, mais disco).
    - Problemas:
        - Limite de Hardware: Há um limite físico para o quanto uma única máquina pode ser poderosa.
        - Ponto de Falha Único: Se a máquina falhar, 100% do sistema fica indisponível.

2. Escala Horizontal (Scale Out):
- O que é: Aumentar o número de máquinas que rodam a aplicação.
- Como funciona: Um load balancer ou proxy reverso é usado para distribuir as requisições entre as várias máquinas.
- Vantagens:
    - Praticamente Ilimitado: Não há o mesmo limite físico de uma única máquina.
    - Tolerância a Falhas: Se uma máquina cair, as outras continuam operando, garantindo a disponibilidade do sistema.
- Desvantagens: Requer que o software seja projetado para isso.

Hoje em dia, a escala horizontal é a abordagem mais comum e preferida devido à sua flexibilidade e resiliência.