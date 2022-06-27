<h2>Социальная сеть</h2>

Что было необходимо реализовать? [Текст](task.md) задания.

<details>
  <summary>Клонируем и запускаем.</summary>

1. ```bash
   git clone https://github.com/darow/some-go-api
   ```

2. ```bash
   cd sb_social_network
   go run ./cmd/api
   ```
</details>

## Доступные методы ##

<details>
  <summary style="color: darkseagreen;">🟢POST /users</summary>

### Создание пользователя ###
##### request example #####

```bash
  curl -X POST -d '{"name":"carl","age":12,"friends":[]}' -H "Content-Type: application/json" http://localhost:8080/create
  curl -X POST -d '{"name":"donald","age":29,"friends":[]}' -H "Content-Type: application/json" http://localhost:8080/create
  curl -X POST -d '{"name":"peter","age":43,"friends":[]}' -H "Content-Type: application/json" http://localhost:8080/create
```
</details>

<details>
  <summary style="color: darkseagreen;">🟢POST /make_friends</summary>

### Добавление в друзья ###
##### request example #####
    
```bash
  curl -X POST -d '{"source_id":"62b5b79fe8ac95cfdd5d1d4e","target_id":"62b5b23ef8d2f2c7bbe27894"}' -H "Content-Type: application/json" http://localhost:8080/make_friends
```
</details>

<details>
  <summary style="color: red;">🔴DELETE /user</summary>

### Удаление пользователя ###
##### request example #####

```bash
  curl -X DELETE -d '{"target_id":"62b5b235f8d2f2c7bbe27892"}' -H "Content-Type: application/json" http://localhost:8080/user
```
</details>

<details>
  <summary style="color: blue;">🔵GET /friends/{user_id}</summary>

### Получение списка друзей пользователя ###
##### request example #####

```bash
    curl -X GET -H "Content-Type: application/json" http://localhost:8080/friends/62b5b239f8d2f2c7bbe27893
```
</details>

<details>
  <summary style="color: orange;">🟠PUT /{user_id}</summary>

### Изменение возраста пользователя ###
##### request example #####

```bash
  curl -X PUT -d '{"new age": 14}' -H "Content-Type: application/json" http://localhost:8082/62b5b79fe8ac95cfdd5d1d4e
```
</details>

## Запускаем два приложения и reverse proxy балансировщик нагрузки ##

<details>
  <summary>Запускаем первый сервер:</summary>

```bash
  go run .\cmd\api\ --config-path configs/test.json
```
</details>

<details>
  <summary>Запускаем второй сервер:</summary>

```bash
  go run .\cmd\api\ --config-path configs/test2.json
```
</details>

<details>
  <summary>Запускаем сервер балансировщик нагрузки:</summary>

```bash
  go run .\reverseproxy\simple\
```

Далее в консоли балансировщика можем увидеть какой сервер обработал входящий запрос.

throwing request to http://localhost:8081
или
throwing request to http://localhost:8080
</details>
