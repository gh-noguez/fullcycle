# Trabalhando com startupProbe

1. Caso o startup probe falhe, o container será reiniciado

- O startupProbe serve como um "período de carência" inicial. Se ele falhar persistentemente, o Kubernetes assume que o container não conseguiu nem mesmo iniciar adequadamente e o reinicia.

2. Caso o startup probe seja bem-sucedido, o liveness probe será executado

- O livenessProbe só começa a ser verificado após o startupProbe ter tido sucesso (ou se não houver um startupProbe configurado, ele começa após o initialDelaySeconds do livenessProbe).

3. E o container será considerado saudável

- Quando o startupProbe é bem-sucedido, ele permite que o livenessProbe e o readinessProbe comecem (ou continuem). O container será considerado "saudável" no sentido de que passou da fase inicial, mas a verificação contínua da saúde é responsabilidade do livenessProbe. O status "Ready" (pronto para receber tráfego) é do readinessProbe.

4. Caso o liveness probe falhe, o container será reiniciado

- Essa é a função principal do livenessProbe: detectar e reagir a containers "presos" ou com problemas que não podem ser resolvidos de outra forma, reiniciando-os.

5. Caso o liveness probe seja bem-sucedido, o container será considerado saudável

- Se o livenessProbe está passando, o Kubernetes considera que o processo principal dentro do container está funcionando como deveria.

6. O startup probe é executado apenas uma vez, no início do container

- Ele é uma verificação de fase única no começo da vida do container. Ele não é executado periodicamente como as outras probes após seu sucesso inicial.

7. O liveness probe é executado periodicamente, enquanto o container estiver em execução

- Ele verifica continuamente a saúde do container.

8. O readiness probe é executado periodicamente, enquanto o container estiver em execução

- Ele verifica continuamente a prontidão do container para receber tráfego.

9. O readiness probe é usado para verificar se o container está pronto para receber tráfego

- Esta é a distinção fundamental. Se o readinessProbe falhar, o Pod é removido dos endpoints do Service, deixando de receber tráfego, mas o container não é reiniciado.

10. O liveness probe é usado para verificar se o container está vivo e funcionando corretamente

- Ele verifica a "vitalidade" do container. Se falhar, o container é reiniciado.

11. O startup probe é usado para verificar se o container está pronto para iniciar

- Mais precisamente, ele verifica se o container conseguiu concluir sua inicialização e está pronto para que as outras probes (Liveness e Readiness) assumam o controle. Ele estende o tempo para que a aplicação "suba" sem que o livenessProbe a reinicie prematuramente.

</br>
- Aplicando e analisando o estado `READY` e `RESTART`:
```bash
kubectl apply -f deployment-v2.yaml && watch -n1 kubectl get pods
```