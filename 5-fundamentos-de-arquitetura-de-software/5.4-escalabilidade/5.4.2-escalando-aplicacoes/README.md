# Escalando aplicações.

Escalabilidade Horizontal: Desafios e Práticas
A escalabilidade horizontal é a abordagem preferida hoje em dia, mas requer uma mudança de paradigma no desenvolvimento e na arquitetura do software. A principal ideia é que as máquinas são descartáveis e podem ser criadas e destruídas a qualquer momento.

1. Descentralização de Dados
- O primeiro passo para escalar horizontalmente é descentralizar os dados e a arquitetura. Um software que armazena dados em sua máquina local não pode ser escalado horizontalmente de forma eficiente.

2. Discos Efêmeros
- O que é: O disco da máquina deve ser tratado como efêmero, o que significa que qualquer coisa salva nele pode ser perdida a qualquer momento.

- Problema: Um blog em WordPress que salva imagens no disco local. Se a máquina cair, as imagens são perdidas. Além disso, se uma nova máquina for adicionada, as imagens não estarão disponíveis.

- Solução: O software deve salvar arquivos em um serviço externo (ex: um bucket S3 da AWS). Dessa forma, todas as máquinas da aplicação poderão acessar os mesmos arquivos, e a perda de uma máquina não resultará em perda de dados.

- Regra: Discos locais devem ser usados apenas para arquivos temporários, que não precisam de persistência.

3. Separação de Servidores
- Servidor de Aplicação vs. Servidor de Assets: O servidor que roda a aplicação (PHP, Java, Go) deve ser separado do servidor que armazena arquivos estáticos (assets), como imagens, CSS e HTML.

- Benefício: Apenas os servidores de aplicação precisam ser escalados e desescalados, enquanto os arquivos estáticos ficam em um local centralizado, acessível por todos.

4. Cache Centralizado
- O que é: O cache não pode mais ser exclusivo da máquina. Ele precisa ser compartilhado em um servidor externo e acessível por todas as máquinas.

- Exemplo: Usar um servidor de cache como o Redis. Quando uma máquina faz uma consulta ao banco de dados e armazena o resultado no Redis, outras máquinas podem se beneficiar desse cache, evitando a duplicação de processamento.

5. Sessões Centralizadas
- Sessões de usuários não podem ser armazenadas localmente na máquina, pois, se o usuário for redirecionado para outra máquina pelo load balancer, ele perderá sua sessão.

- A sessão deve ser armazenada em um local externo, como o servidor de cache compartilhado, para que a aplicação seja stateless (sem estado).

6. Uploads e Gravações de Arquivos Externos
- O mesmo princípio se aplica a uploads e gravação de arquivos importantes (ex: relatórios, extratos).

- Em vez de salvar no disco local, o software deve subir o arquivo para um bucket ou um servidor de arquivos, garantindo que o dado não se perca e seja acessível por qualquer instância da aplicação.

##### Resumo:
Escalar horizontalmente exige descentralização. O software não pode ter estado, nem dados, nem arquivos em sua máquina local. Tudo o que é importante e precisa ser persistido deve estar fora da máquina, em serviços externos. O objetivo é que a máquina possa ser criada e destruída a qualquer momento, sem perda de informação, permitindo aumentar o throughput de forma eficiente.