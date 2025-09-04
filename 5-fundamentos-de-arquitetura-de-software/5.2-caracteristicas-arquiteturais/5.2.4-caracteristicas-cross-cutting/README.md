# Características cross cutting.

Características Arquiteturais: Cross-Cutting
As características cross-cutting (ou transversais) são aspectos que cruzam toda a aplicação e devem ser sempre levados em consideração no dia a dia.

1. Acessibilidade
- O que é: Garantir que a aplicação possa ser utilizada por pessoas com diferentes habilidades ou deficiências (visuais, auditivas, etc.).

- Foco: Principalmente em aplicações com forte componente de front-end.

- Soluções: Utilizar bibliotecas e padrões que facilitam a acessibilidade (ex: para leitores de tela).

- Importância: Existe um grande movimento e conscientização sobre acessibilidade; é crucial aprender e implementar essas práticas para que a plataforma seja inclusiva.

2. Retenção e Recuperação de Dados
- O que é: Definir por quanto tempo os dados serão armazenados e como serão recuperados.

- Desafio: Armazenamento de dados pode ser caro.

- Estratégias:

    - Tempo de Retenção: Definir por quanto tempo os dados precisam estar "quentes" e disponíveis (ex: 7 dias, 30 dias em Kafka).

    - Tiering de Dados: Mover dados menos acessados, mas necessários para compliance, relatórios ou histórico, para storages mais baratos (ex: compactar e guardar).

    - Precisão vs. Tempo: Em sistemas de métricas (ex: Prometheus), dados mais antigos podem ter menor precisão por serem mais compactados e de difícil acesso.

- Reflexão: Pensar na real necessidade de manter cada dado a longo prazo e como gerenciá-lo de forma eficiente.

3. Autenticação e Autorização
- O que é: Gerenciar quem pode acessar o sistema (autenticação) e o que essa pe ssoa pode fazer (autorização).

- Desafio em Arquiteturas Distribuídas: Torna-se mais complexo do que em monolíticos.

- Soluções:

    - Identity Provider (Provedor de Identidade): Servidor centralizado para gerenciar usuários e suas credenciais.

    - API Gateway: Mecanismo na borda da aplicação que pode centralizar a autenticação, timeout, controle de requisições, etc., antes que as requisições cheguem aos microsserviços. Muitos microsserviços não precisam autenticar o usuário, pois a API Gateway já o fez.

- Importância: Em arquiteturas distribuídas, pensar em como autenticar antes que a requisição chegue ao serviço, utilizando mecanismos como o API Gateway, simplifica a lógica de segurança nos microsserviços.

4. Legalidade e Privacidade
- O que é: Garantir que a aplicação esteja em conformidade com as leis e regulamentações dos países onde opera, especialmente no que tange à privacidade dos dados.

- Exemplos:

    - LGPD (Lei Geral de Proteção de Dados) no Brasil: Exige mecanismos para evitar vazamento de dados do usuário.

    - Testes: Evitar usar cópias de bancos de dados de produção com dados sensíveis em ambientes de teste.

- Estratégias para Privacidade:

    - Separação de Dados Sensíveis: Armazenar dados sensíveis em bancos de dados separados, com níveis de serviço diferentes, ou até criptografados.

    - Anonimização/Pseudonimização: Usar apenas dados básicos no sistema principal.

- Importância: A privacidade é um tema crítico hoje em dia, com muitas empresas exigindo contratos rigorosos.

5. Segurança de Ponta a Ponta
- O que é: Proteger a aplicação em todas as camadas, desde a borda até o banco de dados.

- Práticas Recomendadas:

    - Proteção na borda: Usar Web Application Firewalls (WAF) para barrar ataques (SQL Injection, XSS, etc.) e identificar robôs antes que cheguem ao servidor.

    - Padrões Abertos: NUNCA crie suas próprias rotinas de segurança (criptografia, login). Utilize padrões abertos e consolidados criados por especialistas.

    - Separação: Manter o banco de dados separado da aplicação e os backups em outra rede.

- Importância: Pensar na segurança desde o primeiro acesso do usuário.

6. Usabilidade (Não apenas Front-end)
- O que é: A facilidade de uso da aplicação para todos os seus clientes, sejam eles usuários finais ou outras aplicações.

- Front-end: Envolve design de interface, fluxo de usuário e ferramentas de análise de comportamento.

- Back-end (APIs):

    - Organização e Documentação: APIs devem ser bem organizadas, documentadas (ex: usando padrões Open API) e fáceis de consumir por outras aplicações.

    - Contratos Claros: Fornecer contratos claros para as APIs.

    - Exemplos: Protocol Buffers e seus protofiles.

- Importância: Garanta a melhor experiência possível para o seu cliente, seja ele um usuário humano ou outra aplicação.

#### Conclusão:
Ao arquitetar uma aplicação, é crucial ter um checklist mental dessas características cross-cutting e pensar nelas intencionalmente. A "sorte" não deve ser um fator na qualidade do seu software.