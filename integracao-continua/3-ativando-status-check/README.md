# Ativando status check

- Em settings selecionar a branch default como `develop`.

![alt text](image.png)

- Utilizando a organizçõ criada em aulas anteriores, onde já há as configurações necessárias para seguir com a aula. Configurações das branchs como, `develop`como branch default e `Restrict who can push to matching branches`.
- Em `Branches`, marcar `Require status checks to pass before merging`, marcar `Require branches to be up to date before merging`, buscar pelo nome dado ao jobs no workflow, neste caso, `check`, marcar `Restrict who can push to matching branches`, marcar `Restrict pushes that create matching branches` e replicar para a branch main .

![alt text](image-1.png)

- Com estas alterações no repositório, agora vamos realizar a alteração no [yaml](../../.github/workflows/ci.yaml) de configuração.
- Comandos e próximos passos:

- - Criando branch para a ci.
```bash
git checkout -b feature/ci
```

- - git add e git commit:

```bash
git add .github/workflows/ci.yaml
```

```bash
git commit -m "ci: alterando trigger para o workflow develop" 
```

```bash
git push origin develop 
```

Documentação Github Action:
https://docs.github.com/pt/actions
