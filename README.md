# Тестовое задание для компании Crtex на вакансию GO разработчика

## Сервис "Transactions"

### Технологии:

![Golang](https://img.shields.io/badge/-Golang-00ADD8?style=for-the-badge&logo=Go&logoColor=white)
![Postgres](https://img.shields.io/badge/-Postgres-336791?style=for-the-badge&logo=PostgreSQL&logoColor=white)
![gRPC](https://img.shields.io/badge/-gRPC-00C5CA?style=for-the-badge&logo=grpc&logoColor=white)
![Docker-compose](https://img.shields.io/badge/-DockerCompose-23A1F1?style=for-the-badge&logo=docker&logoColor=white)

Сервис представляет из себя два микросервиса: `Api_Gateway` принимает REST запросы и возвращает информацию. За информацией `Api_Gateway` обращается к `Gateway_Processor` по gRPC соединению, тот в свою очередь уже обращается к базе данных.

### Запуск проекта

1. Переходим в директорию проекта:
    ```bash
    cd <путь к проекту>/crtex_test_task
    ```
2. Собираем контейнеры:
    ```bash
    make build
    ```
3. Запускаем:
    ```bash
    make run
    ```
4. Применяем миграции:
    ```bash
    make migration
    ```

### Функционал сервиса

- **Создать пользователя**:
  Отправляем POST запрос на:
    ```
    http://localhost:8000/api/users/
    ```
  Пример JSON:
    ```json
    {
        "first_name": "Имя",
        "last_name": "Фамилия",
        "balance": 100
    }
    ```
  В ответ получите:
    ```json
    {
        "id": 1, 
        "first_name": "Имя",
        "last_name": "Фамилия",
        "balance": 100
    }
    ```

- **Получить пользователя**:
  Отправляем GET запрос на:
    ```
    http://localhost:8000/api/users/<user_id>
    ```
  В ответ получите:
    ```json
    {
        "id": 1, 
        "first_name": "Имя",
        "last_name": "Фамилия",
        "balance": 100
    }
    ```

- **Создать транзакцию**:
  Отправляем POST запрос на:
    ```
    http://localhost:8000/api/users/<user_id>/transactions/
    ```
  Пример JSON:
    ```json
    {
        "amount": 1000,
        "transaction_type": "deposit"
    }
    ```
  В ответ получите:
    ```json
    {
        "id": 1,
        "user_id": 1,
        "amount": 1000,
        "balance_before": 100,
        "transaction_type": "deposit"
    }
    ```

- **Посмотреть все транзакции пользователя**:
  Отправляем GET запрос на:
    ```
    http://localhost:8000/api/users/<user_id>/transactions/
    ```
  В ответ получите:
    ```json
    {
        "transactions": [
            {
                "id": 1,
                "user_id": 1,
                "amount": 100,
                "balance_before": 100,
                "transaction_type": "deposit"
            }
        ]
    }
    ```