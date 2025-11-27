## アーキテクチャ

レイヤードアーキテクチャ

## ディレクトリ構成

```plain
.
├── application
│   ├── gqlmapper   # gql <-> domainのマッパー
│   ├── service
│   │   ├── todo
│   │   └── user
│   └── utils       # pageInfoなどの変換
├── db
│   ├── generator   # gorm gen用の自動生成関数
│   └── migrations
├── domain
│   ├── derr        # エラー定義
│   ├── dmodel
│   │   ├── agg
│   │   └── vo
│   └── drepository # repoのインターフェース
├── graph
│   ├── graphmodel
│   └── graphtype   # scalarやmarshalizerなど
├── infrastructure
│   ├── auth        # 認証まわりのヘルパー(外部ライブラリを用いるヘルパー)
│   ├── database    # DB接続を提供
│   ├── dbmapper    # DBモデル <-> domainのマッパー
│   ├── lib         # IDやパスワード周りのヘルパー(外部ライブラリを用いるヘルパー)
│   ├── model       # gormモデル
│   ├── query
│   └── repository  # repoのimpl
│       ├── todo
│       └── user
├── presentation    # resolver
│
└── server.go
```
