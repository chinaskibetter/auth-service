# Аутентификационный сервис

Этот проект представляет собой простой сервис аутентификации, написанный на языке программирования Go, который предоставляет два REST API эндпоинта для работы с токенами.

## Используемые технологии

- Go
- JWT
- MongoDB

## Задание

Сервис реализует два REST API эндпоинта:

1. `/get-tokens`: Выдает пару Access и Refresh токенов для пользователя с идентификатором (GUID), указанным в параметре запроса.
2. `/refresh-tokens`: Выполняет Refresh операцию на пару Access и Refresh токенов.

## Требования

- Access токен тип JWT с алгоритмом SHA512.
- Refresh токен хранится в базе данных в виде bcrypt хеша.
- Access и Refresh токены обоюдно связаны.

## Запуск

1. Убедиться, что MongoDB запущена и доступна на `mongodb://localhost:27017`.
2. Установить необходимые зависимости с помощью `go mod tidy`.
3. Запустить приложение с помощью команды `go run main.go`.

## Примеры использования

1. Получение пары Access и Refresh токенов:

```bash
curl -X POST http://localhost:8080/get-tokens -d '{"userId":"user_id"}' -H 'Content-Type: application/json'
```

2. Обновление Access токена по Refresh токену:

```bash
curl -X POST http://localhost:8080/refresh-tokens -H 'refresh_token'
```

Замените `user_id` и `refresh_token` на реальные значения.