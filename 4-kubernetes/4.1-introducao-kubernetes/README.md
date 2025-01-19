# Introdução ao Kubernetes.

<p>
O que é o Kubernetes?
O Kubernetes, frequentemente chamado de K8s, é uma plataforma de orquestração de contêineres de código aberto que automatiza a implantação, gerenciamento e dimensionamento de aplicações conteinerizadas. Em termos simples, ele funciona como um "orquestrador" que coordena um grande número de contêineres em um cluster de máquinas.
<p>

<p> 
O Kubernetes é uma plataforma de orquestração de contêineres que permite gerenciar aplicações em larga escala de forma automatizada e eficiente. Ele funciona através de uma API que pode ser acessada via linha de comando usando o kubectl.
<p>

Características Principais:

- Baseado em Estado: O Kubernetes opera com base em um estado desejado. Você define como deseja que suas aplicações sejam configuradas, e o sistema trabalha para manter esse estado.

- Componentes Essenciais:
- - Kubernetes Master:
- - - kube-apiserver: A porta de entrada para toda a interação com o cluster.
- - - kube-controller-manager: Responsável por executar os controladores que garantem que o estado do cluster esteja conforme o desejado.
- - - kube-scheduler: Decide em qual nó do cluster um pod deve ser executado.
- - Outros Nós:
- - - kubelet: Agente que roda em cada nó do cluster e garante que os pods estejam em execução.
- - - kube-proxy: Implementa a rede de serviço do Kubernetes.


<p>
Conceitos Básicos:

- Cluster: Um conjunto de máquinas (nós) que trabalham juntas para executar suas aplicações.
- Node: Uma máquina individual que faz parte do cluster.
- Pod: A menor unidade escalável no Kubernetes. Um pod pode conter um ou mais contêineres que são co-localizados e compartilham o mesmo namespace de rede.
- Deployment: Um objeto Kubernetes que define a desejada condição de replicação de um conjunto de pods.
- Service: Um conjunto de endpoints e um selector que define quais pods devem ser direcionados.
- Namespace: Um mecanismo para isolar recursos em um cluster.
- Controller: Um processo que executa uma lógica de controle específica (por exemplo, Deployment Controller, ReplicaSet Controller).
<p>

Link: https://kubernetes.io/