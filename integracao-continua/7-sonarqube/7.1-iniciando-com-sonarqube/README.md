# Iniciando com Sonarqube.

Nesta aula vamos baixar e executar a imagem Docker para teste.

```bash
docker run -d --name sonarqube -e SONAR_ES_BOOTSTRAP_CHECKS_DISABLE=true -p 9000:9000 sonarqube:latest
```
- Acesso à interface do Sonarqube em `localhost:9000`.
login: admin
password: admin

Link da documentação:
https://docs.sonarsource.com/sonarqube-server/latest/try-out-sonarqube/