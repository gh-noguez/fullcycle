# Características operacionais.

Características Arquiteturais: Lado Operacional
Agora, vamos detalhar as características focadas no lado operacional do software. É crucial pensar intencionalmente nesses aspectos.

1. Disponibilidade
- O que é: Quanto tempo o sistema ficará no ar (ex: 24/7).

- Métricas: Nível de SLA (Service Level Agreement - o que foi combinado com o cliente) e SLO (Service Level Objective - o objetivo interno).

- Importância: É fundamental pensar e gerenciar intencionalmente o nível de disponibilidade desejado para o sistema, em vez de apenas colocá-lo no ar no "piloto automático".

- Impacto: Pensar em disponibilidade desencadeia outras ações, como:

    - Observabilidade: Essencial para verificar a disponibilidade.

    - Técnicas de SRE (Site Reliability Engineering): Criar um "budget" ou "crédito" de indisponibilidade (ex: posso ficar 1 hora indisponível por ano). Monitorar incidentes para garantir que o limite não seja excedido.

2. Recuperação de Desastres
- O que é: Como o sistema se recupera rapidamente quando sai do ar.

- Importância: Crítico para sistemas de missão crítica.

- Processos:

    - Definir ações imediatas quando o sistema cai.

    - Estabelecer como mitigar o problema para evitar recorrência.

    - Considerar cenários como a queda de uma região inteira de nuvem (ex: AWS).

    - Avaliar o custo de soluções multi-região ou multi-cloud para garantir a recuperação.

3. Performance
- O que é: Capacidade do sistema de responder rapidamente e suportar carga.

- Aspectos:

    - Latência: Tempo de resposta.

    - Throughput: Capacidade de receber requisições (ex: 50 transações por segundo, 5000 requisições por segundo).

- Importância: A performance deve ser pensada intencionalmente no projeto. Um sistema projetado para 50 requisições por segundo é diferente de um para 5000 (ex: escolha de banco de dados, uso de CQRS).

- Reflexão: Sempre que um sistema é colocado no ar, deve-se pensar especificamente sobre sua performance.

4. Recuperação / Backup
- O que é: Como o backup do software é feito e gerenciado.

- Ponto Crucial: Testar o backup regularmente. O backup é acionado em situações críticas, e a falha do backup pode ser catastrófica.

- Boas Práticas:

    - Criar políticas e regras para testes periódicos de backup.

    - Armazenar backups em redes separadas para proteção contra ataques (ex: ransomware).

- Cuidado com a Nuvem: Embora a nuvem ofereça facilidades de backup, a responsabilidade de testá-los e garantir sua segurança ainda é nossa.

5. Confiabilidade e Segurança
- O que é: Garantir que o sistema seja robusto contra falhas e ataques.

- Importância: Essencial para sistemas de missão crítica.

- Exemplos de Falhas Comuns:

    - Criação de contas falsas/spam: Sistemas com cadastro simplificado podem ser explorados por robôs, gerando lixo no banco de dados e afetando a confiabilidade.

    - Ataques de Brute Force: Tentativas massivas de login podem sobrecarregar o sistema (banco de dados, CPU), mesmo que as senhas sejam seguras.

        - Solução: Implementar CAPTCHA, Web Application Firewall (WAF) para barrar acessos maliciosos na borda.

- Reflexão: Pensar em como o sistema reagiria a uma tentativa de forçar um endpoint ou a um ataque.

6. Robustez
- O que é: A capacidade do sistema de manter sua funcionalidade e performance sob condições adversas ou crescentes.

- Desafios:

    - A nuvem não é infinita. Empresas com centenas de milhares de instâncias podem esbarrar em limites de capacidade de regiões ou zonas de disponibilidade.

    - Exemplo: Uma região da AWS cai, e as máquinas precisam ser evacuadas para outra região que não tem capacidade para suportar todas as instâncias.

    - Zonas de Disponibilidade: A queda de um datacenter em uma zona de disponibilidade pode expor fraquezas se a rede não estiver bem dividida ou dimensionada.

- Importância: Pensar na robustez significa prever e se preparar para cenários de alta demanda ou falhas em grande escala.

7. Escalabilidade
- O que é: A capacidade da aplicação de crescer para suportar mais carga ou usuários.

-   Tipos de Escala:

    - Vertical: Aumentar os recursos computacionais de uma única máquina (ex: mais CPU, RAM).

    - Horizontal: Adicionar mais máquinas ou recursos computacionais paralelamente.

- Como garantir a escalabilidade horizontal:

    - Trabalhar de forma stateless (sem estado na aplicação).

    - Seguir os "Twelve Factors App" (Heroku)*: Um conjunto de princípios para construir aplicações escaláveis e resilientes na nuvem.

- Consequência da falta de escalabilidade: Se a aplicação não permite escalar horizontalmente, ela atingirá um limite e sofrerá, impedindo o crescimento do negócio.

#### Conclusão sobre Características Operacionais:
Ao desenvolver uma aplicação, você não precisa ser o especialista em SRE, backup ou segurança, mas deve permitir que a aplicação seja operada com esses pilares. Ou seja, a arquitetura deve garantir que ela possa ter escalabilidade, performance, segurança, etc., quando necessário.

#### Twelve-Factor App
- O conceito de "Twelve-Factor App" foi introduzido por desenvolvedores da Heroku e é um conjunto de princípios para construir aplicações modernas e escaláveis, especialmente na nuvem. Esses princípios ajudam a garantir que as aplicações sejam portáveis, resilientes e fáceis de gerenciar. Aqui estão os doze fatores:

1. Codebase: Uma base de código rastreada em um sistema de controle de versão, muitas vezes com múltiplos deploys.

1. Dependencies: Declare explicitamente e isole as dependências da aplicação.

1. Config: Armazene a configuração no ambiente, não no código.

1. Backing Services: Trate serviços de apoio (bancos de dados, filas, etc.) como recursos anexados.

1. Build, Release, Run: Separe estritamente os estágios de build e execução.

1. Processes: Execute a aplicação como um ou mais processos sem estado.

1. Port Binding: Exporte serviços via binding de porta.

1. Concurrency: Escale por meio do modelo de processos.

1. Disposability: Maximize a robustez com inicialização rápida e desligamento gracioso.

1. Dev/Prod Parity: Mantenha o desenvolvimento, staging e produção o mais semelhantes possível.

1. Logs: Trate logs como fluxos de eventos.

1. Admin Processes: Execute tarefas administrativas/gerenciais como processos pontuais.

Esses princípios são projetados para ajudar desenvolvedores a criar aplicações que podem ser facilmente escaladas e mantidas em ambientes de nuvem.* 