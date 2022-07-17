# Golang REST API skeleton

Небольшой каркас проекта, демонстрирующий:
1. Работу с базой данных в Go
1. Использование миграций
1. Использование DI
1. Генерацию REST API по Swagger-спецификации

## Генерация REST API

```bash
swagger generate server urlshortener -f ./swagger.yml --target ./internal/gen --exclude-main --with-context
```


## Миграции

1. [Создать файл миграций](https://github.com/golang-migrate/migrate/blob/master/MIGRATIONS.md) в 
`/internal/migrations` (можно использовать утилиту [CLI](https://github.com/golang-migrate/migrate#cli-usage)).
1. Собрать приложение.
1. Запустить приложение с флагом `migrate`: `./myapp migrate`
