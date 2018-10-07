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
- ログインしなくても利用できるようにする  
  WebDBにデーをを保存する

## Docker
- mysql x2 レプリケーション
- zmemo x2 
- haproxy ロードバランサー

## DBテーブル
- users

|   カラム名    | データ型 |    制約     |        備考         |
| :------------ | :------- | :---------- | :------------------ |
| user_id       | string   | primary_key | uuid                |
| user_name     | string   | not null    |                     |
| user_password | string   | not null    |                     |
| created_at    | string   | not null    | yyyy-mm-dd hh:mm:ss |
| updated_at    | string   | not null    | yyyy-mm-dd hh:mm:ss |
| deleted_at    | string   |             | yyyy-mm-dd hh:mm:ss |

- memos

|  カラム名   | データ型 |          制約           |        備考         |
| :---------- | :------- | :---------------------- | :------------------ |
| user_id     | string   | primary_key, foreignkey | uuid                |
| folder_name | string   | unique, not null        |                     |
| text        | string   | not null                |                     |
| created_at  | string   | not null                | yyyy-mm-dd hh:mm:ss |
| updated_at  | string   | not null                | yyyy-mm-dd hh:mm:ss |
| deleted_at  | string   |                         | yyyy-mm-dd hh:mm:ss |

- folders

|  カラム名   | データ型 |          制約           |        備考         |
| :---------- | :------- | :---------------------- | :------------------ |
| user_id     | string   | primary_key, foreignkey | uuid                |
| folder_name | string   | primary_key             |                     |
| created_at  | string   | not null                | yyyy-mm-dd hh:mm:ss |
| updated_at  | string   | not null                | yyyy-mm-dd hh:mm:ss |
| deleted_at  | string   |                         | yyyy-mm-dd hh:mm:ss |

- role

| カラム名 | データ型 |    制約    |      備考      |
| :------- | :------- | :--------- | :------------- |
| user_id  | string   | foreignkey | uuid           |
| role     | string   |            | 0:user 1:admin |

