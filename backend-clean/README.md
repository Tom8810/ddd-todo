## アーキテクチャ

クリーンアーキテクチャ

## ディレクトリ構成

```plain
.
├── controller # resolver(layered では presentation)
├── db
│ ├── generator
│ └── migrations
├── domain
│ ├── derr
│ ├── dmodel
│ │ ├── agg
│ │ └── vo
│ └── drepository
├── graph
│ ├── graphmodel
│ └── graphtype
├── infrastructure
│ ├── auth
│ ├── database
│ ├── model
│ ├── query
│ └── repository
│ ├── todo
│ └── user
├── internal
│ └── lib # 外部ライブラリを使用する関数(変更は lib に閉じる)
├── port # port(layered では、特段意味を持たない mapper だった)
│ ├── dbport # domain <-> DB
│ ├── domainport # domain <-> DTO(usecase のみのモデル)
│ │ ├── common
│ │ ├── todo
│ │ └── user
│ └── gqlport # DTO <-> GQL
│ ├── common
│ ├── todo
│ └── user
└── usecase # usecase(layered では service)
├── common_dto # 共通の DTO(モデルは usecase 内で定義するのが良い)
├── todo
│ └── dto # todo の DTO
└── user
└── dto # user の DTO
```
