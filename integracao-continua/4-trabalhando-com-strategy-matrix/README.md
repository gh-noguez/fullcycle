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

- Doc:

https://github.com/actions/setup-go