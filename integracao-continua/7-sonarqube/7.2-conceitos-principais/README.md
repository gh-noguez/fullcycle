# Principais conceitos:

### Rules (Regras):
As Rules (regras) são as verificações estáticas de código que o SonarQube aplica aos projetos. Cada regra identifica um possível problema, como vulnerabilidades de segurança, bugs ou problemas de manutenção.

- Exemplo: Verificar se variáveis estão sendo usadas sem inicialização ou se há SQL Injection.
- Objetivo: Garantir boas práticas de codificação e detectar problemas antes que cheguem à produção.
- Customização: Pode-se ativar, desativar ou ajustar regras para atender às necessidades do projeto.


### Quality Profiles (Perfis de Qualidade):
Os Quality Profiles agrupam um conjunto de regras para determinar como o código será analisado.

- Exemplo: Um perfil pode ser configurado para linguagens específicas como Java, Python, ou JavaScript.
- Objetivo: Personalizar e padronizar os critérios de qualidade para cada tecnologia usada no projeto.
- Diferenciação: Você pode criar diferentes perfis para projetos com requisitos distintos, como perfis mais rigorosos para sistemas críticos.


### Quality Gates (Portões de Qualidade):
Os Quality Gates definem os critérios que o código deve atender para ser considerado de boa qualidade.

- Exemplo: Garantir que 80% do código tenha cobertura de testes ou que não haja bugs críticos.
- Objetivo: Agir como um "filtro" para impedir a entrega de código que não atenda aos padrões de qualidade.
- Aprovação/Rejeição: Se os critérios do portão não forem cumpridos, o código será reprovado no pipeline de integração contínua.