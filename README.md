# ゼロメモ
いつでもどこでもゼロ秒思考メモをかけるWebアプリ

## 機能要件
- アカウント登録・削除・更新
- パスワード変更・リセット
- フォルダ作成・削除・更新
- メモ作成・表示・削除

## 非機能要件
- メモのデザインはカードで(大きさ検討中)
- タイトル、本文の入力を分ける(日付は自動付与)
- 本文の入力は4〜6行までをデフォルトとする
  図をかけるようにするかは検討中
- タイマー機能をつける
  時間になったらもう少し続けるか、終了するかを選べるようにする
- ショートカットキーを使用して、タイマーの開始とリセットを出来るようにする

## Docker
- mysql x2 レプリケーション
- apache x2
- haproxy ロードバランサー

## DBテーブル
- users

|   カラム名    | データ型 |      備考      |
| :------------ | :------- | :------------- |
| user_id       | int64    | auto_increment |
| user_name     | string   |                |
| user_password | string   |                |
| create_at     | date     |                |
| updated_at    | date     |                |
| deleted       | int      |                |

- memos

|  カラム名  | データ型 |    備考    |
| :--------- | :------- | :--------- |
| user_id    | int64    | foreignkey |
| text       | string   |            |
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

- role

| カラム名 | データ型 |      備考      |
| :------- | :------- | :------------- |
| user_id  | int64    | foreignkey     |
| role     | int      | 0:user 1:admin |

