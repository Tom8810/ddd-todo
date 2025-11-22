## user

| 区分    | ユースケース       | 入力                  | 出力         | 備考               |
| ------- | ------------------ | --------------------- | ------------ | ------------------ |
| Command | RegisterUser       | name, email, password | bool         | 既存メールチェック |
| Query   | GetUser            | userID                | User         |                    |
| Command | UpdateUserProfile  | userID, name          | bool         |                    |
| Command | UpdateUserEmail    | userID, email         | bool         | 既存メールチェック |
| Command | UpdateUserPassword | userID, password      | bool         |                    |
| Command | DeleteUser         | userID                | bool         |                    |
| Command | Login              | email, password       | result, user |                    |

## todo

query 系

| ユースケース  | 入力例         | 出力例 | 備考 |
| ------------- | -------------- | ------ | ---- |
| ListTodos     | userID, filter | []Todo |      |
| GetTodoDetail | todoID         | Todo   |      |

command 系

| ユースケース     | 入力例                                           | 出力例 | 備考                  |
| ---------------- | ------------------------------------------------ | ------ | --------------------- |
| CreateTodo       | userID, title, description, deadline?, priority? | bool   |                       |
| UpdateTodoDetail | todoID, title?, description?, priority?,         | bool   | status は変更させない |
| StartTodo        | todoID                                           | bool   | waiting → doing       |
| SuspendTodo      | todoID                                           | bool   | doing → waiting       |
| CompleteTodo     | todoID                                           | bool   | doing → completed     |
| ReopenTodo       | todoID                                           | bool   | completed → doing     |
| DeleteTodo       | todoID                                           | bool   |                       |
