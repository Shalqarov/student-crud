# student-crud
## Запуск
С помощью этой команды можно запустить докер контейнеры и начать работу с api :
```console
$ make run
```
Для того, чтобы остановить контейнеры введите :
```console
$ make stop
```
## Задание
- Написать API Restful CRUD (Create, Read, Update, Delete). 
- Создать таблицу Student с полями (ID,name,e-mail,
phoneNumbers, и тд). 
- Нужно сделать API с обращением в базу данных(Postgresql) , и контроллер для всех crud операций.

## О выполнении
Я использовал для роутинга и хендлеров библиотеку fiber(https://github.com/gofiber/fiber).

Для запросов в бд использовал Gorm(https://pkg.go.dev/gorm.io/gorm)

### Доступные эндпоинты
- `Post/    /api/v1/students`
- `Get /    /api/v1/students/:id`
- `Get /    /api/v1/students/`
- `Put /    /api/v1/students/:id`
- `Delete / /api/v1/students/:id`

Данные отправляются посредством JSON, например: 
```json
{
    "first_name": "Temirlan",
    "last_name": "Shalkarov",
    "email": "asd@asd",
    "phone_number": "87548652315"
}
```