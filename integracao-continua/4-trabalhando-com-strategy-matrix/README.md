# Trabalhando com strategy Matrix

- No workflow, vamos adicionar o trecho abaixo para poder trabalhar com uma matrix de versões da linguagem Golang:

```bash
build:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        go: [ '1.14', '1.13', '1.15' ]
    name: Go ${{ matrix.go }} sample
    ...
        go-version: ${{ matrix.go }}
```

- Com isto, será realizado o check nas versões adicionadas à matrix.


- Como também tive o erro de ficar travado no "check", tive de alterar em settings, desativando temporariamente o `status check` para poder realizer o merge.

- Após isto, realizei os comandos abaixo:

```bash
git checkout develop
```

```bash
git pull origin develop
```

```bash
git branch -d feature/ci-strategy-matrix 
```

- Doc:

https://github.com/actions/setup-go