# ゼロメモ
いつでもどこでもゼロ秒思考メモをかけるWebアプリ

## 機能
- タイマー
- アカウント
- フォルダ
- メモ作成・表示・削除

## DBスキーマ
- users

|   カラム名    | データ型 |      備考      |
| :------------ | :------- | :------------- |
| user_id       | int64    | auto_increment |
| user_name     | string   |                |
| user_password | string   |                |
| create_at     | date     |                |
| updated_at    | date     |                |


- memos

|  カラム名  | データ型 |    備考    |
| :--------- | :------- | :--------- |
| user_id    | int64    | foreignkey |
| memo       | string   |            |
| create_at  | date     |            |
| updated_at | date     |            |

- folders

|  カラム名   | データ型 |    備考    |
| :---------- | :------- | :--------- |
| user_id     | int64    | foreignkey |
| folder_id   | int64    |            |
| folder_name | string   |            |
| create_at   | date     |            |
| updated_at  | date     |            |

