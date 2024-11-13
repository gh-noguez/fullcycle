# Padrões e técnicas avançadas com Git e Github

#### Aqui será documentado todos os comandos utilizados em aula com uma breve definição de cada comando.

- Exibir lista de comando do git flow:
```bash
git flow
```

#### Para manter meu repositório e integrar o git glow, segui os passos a seguir:

```bash
git branch git-flow
```
```bash
git checkout git-flow
```

- Foi necessário criar uma branch develop para o git flow identifica-la e assim sendo possível a integração quando solicitado:
```bash
git branch develop
```

```bash
git flow init
```

```bash
git rebase main
```
```bash
git branch -D git-flow
```

- Aqui criamos uma branch feature welcome:
```bash
git flow feature start welcome
```
- Neste momento, foi criado um [index.html](index.html) para realizar os testes de aula, que consiste em criar uma feature, desenvolve-la no seu branch, em seguida enviar para a branch develop e deletando a branch de feature com o comando '`git flow feature finish welcome`', adicionado e enviado para o repositório remoto.
```bash
git add git-github/mao-na-massa/index.html
```

```bash
git commit -m "Feat: Adicionado index.html"
```

```bash
git branch
```

```bash
git status
```

```bash
git log
```

Summary of actions:
- The feature branch 'feature/welcome' was merged into 'develop'
- Feature branch 'feature/welcome' has been locally deleted
- You are now on branch 'develop'
```bash
git flow feature finish welcome
```

- Agora sera enviado para a 'release', exemplo para atualizações e/ou ajustes antes de ser enviado para produção, então será criada a release com o comando baixo e em seguida será criada uma nova feature:
```bash
git flow release start 0.1.0
```

```bash
git flow feature start contato
```

verificando todas as branchs:
  develop
* feature/contato
  main
  release/0.1.0
```bash
git branch
```
- Agora foi adicionado página de contato
```bash
git add .
```

```bash
git commit -m "Feat: Adicionando contato"
```

Summary of actions:
- The feature branch 'feature/contato' was merged into 'develop'
- Feature branch 'feature/contato' has been locally deleted
- You are now on branch 'develop'

```bash
git flow feature finish contato
```

```bash
git log
```

```bash
git checkout release/0.1.0
```

```bash
git status
```

```bash
git add .
```

```bash
git commit -m "Fix: Ajustes no index"
```

```bash
git log
```

- Comentário não precisa ser alterado na primeira mensagem, já vem correto, já o nome da tag é obrigatório ser adicionado
Summary of actions:
- Release branch 'release/0.1.0' has been merged into 'main'
- The release was tagged '0.1.0'
- Release tag '0.1.0' has been back-merged into 'develop'
- Release branch 'release/0.1.0' has been locally deleted
- You are now on branch 'develop'

```bash
git flow release finish 0.1.0
```

```bash
git branch
```

```bash
git log
```

```bash
git tag
```

- Corrigindo bugs
```bash
git flow hotfix start contato
```

- "Correção" no index e feito commit
```bash
git add .
```

```bash
git commit -m "Fix: Index, ajuste de redirecionamento contato"
```

Summary of actions:
- Hotfix branch 'hotfix/contato' has been merged into 'main'
- The hotfix was tagged 'contato'
- Hotfix tag 'contato' has been back-merged into 'develop'
- Hotfix branch 'hotfix/contato' has been locally deleted
- You are now on branch 'develop'

```bash
git flow hotfix finish contato
```

```bash
git branch
```

```bash
git tag
```