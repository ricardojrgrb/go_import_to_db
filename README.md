# Importação de um arquivo TXT para o PostgreSQL em Golang

Requisitos:
- Docker: https://docs.docker.com/get-docker/
- Github: https://git-scm.com/downloads

Passos para rodar o projeto(executar no terminal):

1 Passo: realize o clone deste projeto;

2 Passo: acesse o diretório go_import_to_db;

3 Passo: docker compose up -d go_db;

4 Passo: docker compose build;

5 Passo: docker compose up go-app;

Obs01.: Após a realização dos passos indicados, acesse o container go_app, e na aba Logs deverá ser apresentado o log abaixo:

Database created!

Table created!

File processing started!

Finished file processing!

Obs02.: Após obter o log acima, utilize a ferramenta de gestão de banco de dados da sua preferência e visualize os dados persistidos.

Sugestão de ferramenta para acesso ao banco de dados: https://dbeaver.io/download/

Configurações do banco no dbeaver:

![Captura de Tela 2024-07-13 às 17 57 31](https://github.com/user-attachments/assets/04bab849-ee7c-4a30-8926-ad238d481287)
