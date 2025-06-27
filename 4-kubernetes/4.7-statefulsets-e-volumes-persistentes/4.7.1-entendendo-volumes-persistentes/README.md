# Entendendo volumes persistentes.

- Nesta aula foi criado um yaml de PV apenas para demonstração/exemplo.

- O sistema de Persistent Volumes (PV) no Kubernetes foi criado para abstrair a complexidade do armazenamento de infraestrutura dos desenvolvedores de aplicações, permitindo que eles simplesmente "solicitem" armazenamento sem se preocupar com onde ou como ele é provisionado.

- Os principais componentes são:

- PersistentVolume (PV):

- O quê é: É uma peça de armazenamento real no cluster. Pense nele como um disco físico ou um volume de armazenamento em nuvem (como um volume EBS na AWS ou um disco persistente no Google Cloud) que foi "montado" e está disponível para ser usado.

- Provisionamento: Pode ser provisionado estaticamente por um administrador (onde o administrador configura o PV manualmente) ou dinamicamente (onde o Kubernetes cria o PV automaticamente sob demanda, usando um StorageClass).

- Ciclo de Vida: Sua vida é independente do Pod. Se um Pod que o está usando for deletado, o PV não é deletado automaticamente (a menos que configurado para isso), mantendo os dados seguros.

- Detalhes: Ele encapsula os detalhes d- a infraestrutura de armazenamento subjacente (NFS, iSCSI, Fibre Channel, volumes específicos de provedores de nuvem, etc.).

- PersistentVolumeClaim (PVC):

- O quê é: É a solicitação de armazenamento feita por um usuário (ou por um Pod) ao Kubernetes. É como um desenvolvedor dizendo: "Preciso de 10GB de armazenamento que possa ser lido e escrito por um Pod de cada vez."

- Consumo: Um PVC "consome" (vincula-se a) um PV que atenda aos seus critérios (tamanho, modo de acesso, tipo de armazenamento). Uma vez que um PV está vinculado a um PVC, ele não pode ser vinculado a outro PVC.

- Abstração: O usuário do Pod só interage com o PVC, não com o PV diretamente, o que proporciona a abstração tão desejada.

- StorageClass:

- O quê é: É um objeto do Kubernetes que descreve diferentes "classes" de armazenamento disponíveis no cluster. Pense nele como um "perfil" ou "tipo" de armazenamento.

- Provisionamento Dinâmico: É fundamental para o provisionamento dinâmico de PVs. Quando um PVC é criado e solicita um StorageClass específico, o Kubernetes (através de um provisionador) cria um PV automaticamente que corresponde àquele StorageClass.

- Características: Define atributos como: o provisionador (qual plugin de armazenamento vai criar o PV), parâmetros específicos do provisionador (tipo de disco, IOPS, etc.), e a política de recuperação (reclaimPolicy) que define o que acontece com o PV e os dados quando o PVC é excluído (ex: Retain, Delete).

- Fluxo Básico (Provisionamento Dinâmico):

- Um desenvolvedor cria um PVC (ex: "Preciso de 5GB de um armazenamento fast-ssd").

- O StorageClass chamado fast-ssd (previamente configurado pelo administrador) define como e onde esse armazenamento deve ser provisionado.

- O Kubernetes, usando o provisionador definido no StorageClass, cria um novo PV na infraestrutura subjacente (ex: um novo volume SSD na AWS).

- O Kubernetes então vincula esse novo PV ao PVC solicitado.

- Finalmente, o Pod do desenvolvedor pode montar o PVC (e, por extensão, o PV vinculado) em seu filesystem, e a aplicação pode ler/escrever dados.

- Este sistema permite uma grande flexibilidade e desacoplamento entre as necessidades de armazenamento das aplicações e a complexidade da infraestrutura subjacente.


- Fonte: https://kubernetes.io/docs/concepts/storage/persistent-volumes/
