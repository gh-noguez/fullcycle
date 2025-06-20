# Entendendo conceito de service.

- Em Kubernetes, um "Service" é um objeto que define um conjunto lógico de pods e uma política para acessá-los. Essencialmente, ele age como um balanceador de carga e um ponto de acesso estável para seus aplicativos, mesmo quando os pods individuais são criados, destruídos ou atualizados. Detalhando. . .

- Abstração da rede:
Um Service fornece um endereço IP e um nome DNS estáveis para um grupo de pods. Isso permite que outros componentes do cluster, ou mesmo clientes externos, se comuniquem com seus aplicativos sem se preocupar com os detalhes de quais pods estão ativos e onde eles estão localizados. 

- Balanceamento de carga:
O Service distribui o tráfego de rede entre os pods que ele seleciona, garantindo que nenhum pod seja sobrecarregado e que o aplicativo permaneça disponível mesmo com falhas em alguns pods. 
Descoberta de serviços:
O Kubernetes possui um sistema DNS interno que registra os Services e seus endereços IP, permitindo que outros pods e componentes descubram e se comuniquem com os Services usando nomes DNS. 

- Tipos de Service:
Existem diferentes tipos de Services, cada um com um propósito específico:

1. ClusterIP: O tipo padrão, que fornece um endereço IP interno ao cluster para acesso interno aos pods. 

2. NodePort: Expondo o Service em uma porta em cada nó do cluster, permitindo acesso externo através de qualquer IP de nó. 
3. LoadBalancer: Cria um balanceador de carga externo, como um provedor de nuvem, para expor o Service na internet. 
4. ExternalName: Mapeia o Service para um nome DNS externo, útil para integração com serviços externos. 

- Seletores:
Os Services usam seletores para identificar os pods que eles devem gerenciar. Um seletor é um conjunto de rótulos que são aplicados aos pods, e o Service só seleciona pods que possuem todos os rótulos especificados no seletor. 

Em resumo: Um Service em Kubernetes é uma ferramenta essencial para criar uma infraestrutura de aplicativos resiliente e escalável, fornecendo um ponto de acesso estável e gerenciando o tráfego para seus pods. Ele abstrai a complexidade da rede e permite que você se concentre no desenvolvimento de seu aplicativo


Fonte: Google

- [Documentação oficial sobre service](https://kubernetes.io/docs/concepts/services-networking/service/).