# day-one-sql

## DAY1
onion0904の7日間チャレンジの一日目です。
詳しくは[zenn](https://zenn.dev/onion0904/articles/ff700890522030)

## 作ったもの

CLI日記(標準入力から操作)

""で囲まれてる部分は自分で変えて入力してください

- 一行目(ユーザー)
    - "username"(ユーザー作成していない人は作成される)
    - delete "username"(ユーザーの削除)

- 二行目(操作選択)
    - upsert(作成,更新)
    - show(取得)
    - delete(削除)

- 三行以降(操作)
    - upsertの場合
        - "2004-09-04" "誕生日"(日付,内容)
        - 日付の更新はできません
    - showの場合
        - all(全て)
        - "2004-09-04"(日付で検索)
        - "誕生日"(内容で検索)
    - deleteの場合
        - "2004-09-04" (日付)

## 入力例
```
go run main.go
input UserOperation: onion0904
input operation: upsert
input date and content: 2025-05-30 徹夜した
```