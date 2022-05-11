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
      curl -X POST -d '{"name":"name1","age":"24","friends":[]}' -H "Content-Type: application/json" http://localhost:8080/create
   ```
</details>

<details>
  <summary style="color: darkseagreen;">🟢POST /make_friends</summary>

### Добавление в друзья ###
##### request example #####

   ```bash
      curl -X POST -d '{"source_id":1,"target_id":2}' -H "Content-Type: application/json" http://localhost:8080/make_friends
   ```
</details>

<details>
  <summary style="color: red;">🔴DELETE /user</summary>

### Удаление пользователя ###
##### request example #####

   ```bash
      curl -X DELETE -d '{"target_id":2}' -H "Content-Type: application/json" http://localhost:8080/user
   ```
</details>

<details>
  <summary style="color: blue;">🔵GET /friends/{user_id}</summary>

### Получение списка друзей пользователя ###
##### request example #####

   ```bash
      curl -X GET -H "Content-Type: application/json" http://localhost:8080/friends/1
   ```
</details>

