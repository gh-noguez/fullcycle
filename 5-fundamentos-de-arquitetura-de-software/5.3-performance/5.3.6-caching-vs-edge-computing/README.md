# Caching vs Edge Computing.

#### Edge Computing e Cache

A Edge Computing está cada vez mais em evidência por sua importância em otimizar o tráfego de dados na internet. O conceito principal é mover o processamento e o armazenamento de informações para mais próximo do usuário, reduzindo a latência e o congestionamento da rede.

##### O Problema do Tráfego Distante
- Se todos os dados estivessem centralizados em um único local (ex: um datacenter nos EUA), usuários em outras partes do mundo (Brasil, Japão) teriam que esperar o tráfego percorrer grandes distâncias.
- Isso não apenas degrada a experiência do usuário com alta latência, mas também sobrecarrega a infraestrutura da internet e do provedor de nuvem.

##### A Solução da Edge Computing
- A Edge Computing faz com que a informação do usuário esteja mais próxima dele.
- Isso evita que a requisição precise trafegar por longas distâncias, diminuindo o tempo de resposta (response time) e melhorando a performance.
- Pode fornecer serviços além de um simples cache, processando informações na borda e economizando custos de infraestrutura do servidor principal.

##### CDN (Content Delivery Network)
- Uma CDN é uma malha de servidores distribuídos globalmente. Quando um conteúdo (ex: um vídeo da Netflix) é carregado, ele é replicado para os servidores mais próximos dos usuários.
- Quando um usuário solicita o conteúdo, ele é servido a partir do datacenter mais próximo.
- Exemplo: A Akamai, uma das maiores CDNs do mundo, tem milhares de pontos de presença. Quando a Full Cycle carrega um vídeo, a CDN o replica para esses pontos de presença, e os usuários recebem o vídeo do servidor mais próximo, reduzindo a latência e o buffering.

##### Custos da CDN
O custo de uma CDN é geralmente dividido em duas partes:

1. Custo de Midgress: O custo para espalhar o conteúdo do servidor de origem (Origin) para todos os servidores da CDN.

2. Custo da Edge: O custo do tráfego de dados consumido pelos usuários a partir dos servidores da CDN.

##### Serviços de Edge Computing
- Cloudflare Workers: Uma plataforma de Edge Computing que permite a execução de aplicações (geralmente em JavaScript) em servidores próximos ao usuário.
    - Utiliza a engine V8 do Google para criar "mini-containers" leves e rápidos.
    - Permite a execução de código na borda, economizando recursos e tempo de resposta.
- Vercel: Plataforma conhecida por manter o Next.js.
    - Trabalha com um tipo de Edge Computing para entregar sites estáticos e dinâmicos de forma otimizada.
    - Utiliza mecanismos de cache e revalidação de dados para garantir que o conteúdo seja sempre atualizado.

A Edge Computing é uma cultura de desenvolvimento essencial nos dias de hoje para garantir alta performance e escalabilidade, evitando sobrecarga na infraestrutura e proporcionando uma melhor experiência para o usuário final.