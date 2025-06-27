# Entendendo utilização de resources.

- Nesta aula foi adicionado resources ao deployment, que reserva um mínimo e o limite de recurso para o Pod.

- vCPU = 1000m (milicores), é possível utilizar porcentagem também: 500m = 0.5

- É possível definir `requests` e `limits` para memória e CPU.
- `Requests` é o mínimo que será reservado para o pod.
- `Limits` é o máximo que o pod poderá utilizar.
- É importante definir esses valores para evitar que o pod consuma todos os recursos do cluster e cause problemas de performance ou até mesmo crash do pod.
- É  possível utilizar porcentagens também, por exemplo: 500m = 0.5 vCPU
- É possível utilizar também o valor em milicores, por exemplo: 1000m = 1 vCPU, 500m = 0.5 vCPU, 250m = 0.25 vCPU
- É possível utilizar também o valor em bytes, por exemplo: 100Mi = 100 megabytes, 50Mi = 50 megabytes, 25Mi = 25 megabytes
- É possível utilizar também o valor em gibabytes, por exemplo: 1Gi = 1 gibabyte, 512Mi = 512 megabytes, 256Mi = 256 megabytes
