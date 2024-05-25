# code-review-fullcycle

- Comando utilizado para criar um submódulo (adicionar repositório dentro do repositório)

```bash
git submodule add https://github.com/gh-noguez/code-review-fullcycle /home/felipe/Documents/workspaces/fullcycle/git-github/4-prs-e-code-review/code-review-fullcycle
```

#### Problema ao realizar push na minha organização via terminal

- Configurar o repositório para usar a chave SSH correta:

  - Verifique o URL remoto do repositório. Ele deve estar configurado para usar o SSH.
```bash
git remote -v
```
O URL deve ser algo como git@github.com:organizacao/nome-do-repositorio.git. Se estiver configurado para HTTPS, você pode alterar para SSH com:
```bash
git remote set-url origin git@github.com:organizacao/nome-do-repositorio.git
```

Comandu usado
```bash
git remote set-url origin  git@github.com:gh-noguez/code-review-fullcycle.git
```
