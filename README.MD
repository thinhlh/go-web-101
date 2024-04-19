# GO WEB 101

## Prerequisite
- Docker/Docker compose
- Golang
- Makefile (optional)
- Copy env
    ```
    cp ./env/.env.example ./env/.env.dev
    ```

## Boot
1. Local (Boot whole stack)
    ```
    make run-local
    ```
2. Dev
    ```
    make pre-run
    make -j2 services-dev
    ```
3. Prod
    ```
    make pre-run
    make -j2 services
    ```