

## Инструкция по использованию

| Команды            | Описание                                      |
|--------------------|-----------------------------------------------|
| make server        | Запуск главного сервиса                       |
| make mail          | Запуск сервиса, которые осуществляет рассылки |
| make migrate-up    | Создание миграций бд                          |
| make migrate-down  | Удаление всех таблиц в бд                     |
| make migrate-force | Применяется в случае dirty бд                 |
| make docker-up | Поднятие сервисов                             |
| make docker-down | Завершение работы сервисов                    |
| make docker-build | Пересборка docker compose                     |
| make docker-clear | Удаление неиспользуемых докер контейнеров     |

## API

### Client entity

### POST - `/api/client/create`

Запрос:

```json
{
  "client_property": {
    "operator_code": "750",
    "tag": "vip"
  },
  "phone_number": "79506137355",
  "time_zone": "Europe/Moscow"
}
```

Ответ:

| Код ответа | сообщение               |
|------------|-------------------------|
| 201        | -                       |
| 400        | `{"message": "string"}` |
| 500        | `{"message": "string"}` |

### POST - `/api/client/update`

Запрос:

```json
{
  "id": "c93694e0-72ed-45c4-8a31-852c17e2c066",
  "phone_number":"71506117100",
  "client_property":{
    "tag":"vip",
    "operator_code":"800"
  },
  "time_zone":"Africa/Bangui"
}
```

Ответ:

| Код ответа | сообщение               |
|------------|-------------------------|
| 204        | -                       |
| 400        | `{"message": "string"}` |
| 500        | `{"message": "string"}` |

### Delete - `/api/client/delete`

Запрос:

```json
{
  "id": "bb44bf9d-8892-493b-855d-808426033b44"
}
```

Ответ:

| Код ответа | сообщение               |
|------------|-------------------------|
| 204        | -                       |
| 400        | `{"message": "string"}` |
| 500        | `{"message": "string"}` |

### Message entity

### Get - `/api/message/group`

Ответ:

| Код ответа | сообщение                             |
|------------|---------------------------------------|
| 200        | ```map[string][]entity.MessageInfo``` |
| 500        | `{"message": "string"}`               |

### Get - `/api/message/info/{id}`

Запрос:
```json
 {
  "id": "cba47385-e24b-40ec-9b19-ff3ffbdf3e4a"
}
```

Ответ:

| Код ответа | сообщение                             |
|------------|---------------------------------------|
| 200        | ```map[string][]entity.MessageInfo``` |
| 400        | `{"message": "string"}`               |
| 404        | `{"message": "string"}`               |
| 500        | `{"message": "string"}`               |

### Notification entity

### Post - `api/notification/create`

Запрос:

```json
{
  "message":"Hello world",
  "create_at":"01:00 11.10.2023",
  "expires_at":"01:00 11.10.2023",
  "tags":[
    {
      "tag":"vip"
    },
    {
      "tag":"bomj"
    }
  ],
  "operator_codes":[
    {
      "operator_code":"121"
    },
    {
      "operator_code":"750"
    }
  ]
}
```

Ответ:

| Код ответа | сообщение               |
|------------|-------------------------|
| 201        | -                       |
| 404        | `{"message": "string"}` |
| 500        | `{"message": "string"}` |


### Post - `api/notification/stat`


Запрос:

```json
{
  "create_at": "01:22 10.10.2023"
}
```

Ответ:

| Код ответа | сообщение               |
|------------|-------------------------|
| 200        | ниженаписанный json     |
| 404        | `{"message": "string"}` |
| 500        | `{"message": "string"}` |


```json
[
  {
    "client_property": [
      {
        "operator_code": "string",
        "tag": "string"
      }
    ],
    "create_at": "string",
    "expires_at": "string",
    "id": "string",
    "id_client_properties": [
      "string"
    ],
    "message": "string"
  }
]
```

### Post - `api/notification/update`

Запрос:

```json
{
  "create_at": "string",
  "expires_at": "string",
  "message": "string"
}
```

Ответ:

| Код ответа | сообщение               |
|------------|-------------------------|
| 204        | -                       |
| 400        | `{"message": "string"}` |
| 500        | `{"message": "string"}` |


### Delete - `api/notification/delete`

Запрос:

```json
{
  "create_at": "11:21 11.10.2023"
}
```

Ответ:

| Код ответа | сообщение               |
|------------|-------------------------|
| 204        | -                       |
| 400        | `{"message": "string"}` |
| 500        | `{"message": "string"}` |



### Дополнительные задания

- [x] подготовить docker-compose для запуска всех сервисов проекта одной командой

- [x] сделать так, чтобы по адресу /docs/ открывалась страница со Swagger UI и в нём отображалось описание разработанного API. Пример: https://petstore.swagger.io

Замечания по реализации:

- При создании клиента с несуществующими атрибутами, они добавятся в таблицу client_properties;

- При создании рассылки с created_at меньше времени сейчас и с expires_at меньше времени сейчас будет ошибка;

- При создании рассылки с created_at меньше времени сейчас и с expires_at больше времени сейчас будет добавлена запись в таблицу signal
и в таблице notification будет добавлен true к записи with_signal;

- Записи в таблице signal удяляются после их селекта.
