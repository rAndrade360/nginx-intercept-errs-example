# Interceptando erros com o nginx

## Pré requisitos

Para executar esse repo, você precisa ter as seguintes ferramentas instaladas:

 - [docker](https://www.docker.com/)
 - [docker compose](https://docs.docker.com/compose/)
 - [make](https://www.gnu.org/software/make/#documentation)
 - [curl](https://curl.se/)

## Como executar a aplicação

1. Execute o seguinte comando para iniciar a app no docker compose:

```bash
$ make run
```

2. Execute algum dos seguintes comandos para testar os cenários de sucesso e erro:

    - `make success` para a rota `/success`
    - `make not_found` para a rota `/not_found`
    - `make bad_request` para a rota `/bad_request`
    - `make internal_server_error` para a rota `/internal_server_error`