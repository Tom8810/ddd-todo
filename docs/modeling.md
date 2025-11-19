### Bounded Context

特になし。todo アプリ程度のサイズであれば、コンテキストを分ける必要はない。

### 業務フローを理解

- 業務フロー

  1. 1 日の開始時に、タスクのやり忘れがないように、今日やることを確認する。

     行動 / アウトカム

     - アプリにログインする → ログインできる
     - タスク一覧を確認する → 一覧取得できる
     - 今日締め切りのタスクに絞って確認する → 締切日フィルタでソートできる

     ドメインルール

     - 締切日が今日のタスクは強調される
     - タスクはステータスごとに表示される
     - タスクはプロパティでソート、フィルタ可能

  2. 随時、発生したタスクを todo として作成する

     行動 / アウトカム

     - タスクが生じたら、todo を作成する → 作成できる
     - 作成時に、締切日や優先度、説明などを指定する → 指定し、保存できる

     ドメインルール

     - タスクを作成したらステータスは waiting になる
     - タイトルだけは必須
     - 優先度は低、中、高
     - 締切日は今日以降

  3. タスクに着手したタイミングでステータスを変更する

     行動 / アウトカム

     - タスクを始めたら、ステータスを doing に変更できる
     - もう一度未着手に戻したい時は、ステータスを waiting に変更できる

     ドメインルール

     - ステータスは waiting/doing/complete

  4. タスクをやっている途中で、タスクの内容が変わった場合、タスクを更新したり、削除したりする

     行動 / アウトカム

     - タスクの内容が変わったら、以下を変更できる
       - タイトル
       - 説明
       - 締切日
       - 優先度
     - タスクが不要になれば、削除できる

     ドメインルール

     - ステータスに関係なく、タスクの内容は変更可能
     - ステータスに関係なく、タスクの削除は変更可能

  5. タスクを完了したら、ステータスを完了にできる

     行動 / アウトカム

     - タスクを完了したら、ステータスを completed に変更できる

- ユースケース一覧
  - todo の CRUD
    - Create
    - Read
      - 一覧を見る
        - フィルタ
        - ソート
      - 詳細を見る
    - Update
      - ステータスを変更する
      - タイトルを変更する
      - 説明を変更する
      - 期限を変更する
    - Delete
  - user の CRUD

### ドメイン概念を抽出

- ドメイン概念一覧
  - タスク(todo)
  - ログイン
  - フィルタ
  - ソート
  - 締切日
  - 優先度
  - 説明
  - タイトル
  - ステータス
  - ユーザー
  - メールアドレス
  - 所有者
  - ユーザー名
  - パスワード
  - ステータス更新
  - タスク内容更新

### ドメイン概念の制約を整理

- ドメイン概念の制約
  - タイトル
    - ~64 文字
    - 必須
    - 作成時＆更新時に VO がチェック
  - 締切日
    - 値がある場合は、今日以降であること
    - 任意
    - 作成時&更新時に値は VO がチェック、null 許可は Entity で担う
  - 優先度
    - 値がある場合は、高、中、低のいずれか
    - 任意
    - 作成時&更新時に値は VO がチェック、null 許可は Entity が担う
  - 説明
    - 任意
  - ステータス
    - waiting、doing、complete のいずれか
    - todo 作成時に waiting になる
    - 必須
    - 作成時&更新時に値は VO がチェック
  - ユーザー名
    - ~64 文字
    - 必須
    - 作成時&更新時に値は VO がチェック
  - メールアドレス
    - メールアドレス形式である
    - 必須
    - 作成時&更新時に値は VO がチェック
  - パスワードは~16 文字、必須
    - 必須
    - ~16 文字
    - 作成時&更新時に値は VO がチェック、ハッシュ化は VO が担う

### エンティティ / VO / 集約を決める

- モデル定義

  - エンティティ
    - Todo: タスク
    - User: ユーザー
  - 値オブジェクト
    - Title: タスク名
    - Description: 説明
    - Status: ステータス
    - Deadline: 締切日
    - Priority: 優先度
    - name: ユーザー名
    - email: メールアドレス
    - password: パスワード
  - アグリゲーション
    Todo Agg

    ```go
    type Todo struct {
      id: TodoId // VO: 型安全のため
      title: Title // VO: 文字数バリデーションのため
      description: *string
      status: Status // VO: ロジックを持たせるため(ステータス変更)
      deadline: *Deadline // VO: 現在以降バリデーションのため
      priority: *Priority // VO: Enumバリデーションのため
      userId: UserId // VO: 型安全のため
      createdAt: time.time
      updatedAt: time.time
    }
    ```

    User Agg

    ```go
    type User struct {
    	id: UserId // VO: 型安全のため
    	name: UserName // VO: 文字数バリデーションのため
        email: Email // VO: メールアドレス形式バリデーションのため
        paasword: Password // VO: ハッシュ化のため
        createdAt   time.Time
        updatedAt   time.Time
    }
    ```

### ユビキタス言語を整備

- todo(タスク): やるべきこと
- status(ステータス): todo の進捗状況。waiting / doing / completed
- user(ユーザー): todo を登録・編集する人

### 振る舞いをモデル化

- Todo
  固有の不変条件：なし(VO の制約のみ)
  - ChangeTitle
    - 入力：Title
    - 制約：Title が Valid(~64 文字)
  - ChangeDeadline
    - 入力：Deadline
    - 制約：null 許容。Deadline が Valid(今日以降)
  - ChangePriority
    - 入力：Priority
    - 制約：null 許容。Priority が Valid(高、中、低のいずれか)
  - ChangeStatus
    - 入力：Status
    - 制約：Status が Valid(waiting / doing / completed のいずれか)
- User
  固有の不変条件：なし(VO の制約のみ)
  - ChangeName
    - 入力：UserName
    - 制約：UserName が Valid(~64 文字)
  - ChangeEmail
    - 入力：Email
    - 制約：Email が Valid(Email 形式である)
  - ChangePassword
    - 入力：Password
    - 制約：Password が Valid(~16 文字)
      - ハッシュ化は VO 内で行う
