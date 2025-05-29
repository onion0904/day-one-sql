## postgresqlの起動方法

### GUI

```
# imageのpull,containerのrun,初期設定
docker-compose.yml
```
http://localhost:8080にアクセスして、ログインする

ここでpostgresqlの操作ができます。

### CLI

```
# imageのpull,containerのrun,初期設定
docker-compose.yml
# containerの状態を確認 
docker ps
# psのCONTAINER IDを確認して実行
docker exec -it [container id] psql -U onion0904 -d db
```

## postgresqlの書き方

### DBの作成

```
CREATE DATABASE example;
```
これをすることで
```
docker exec -it ac060a1a7817 psql -U onion0904 -d example
```
のようにexampleを操作することができる

### DDL(データ定義言語)

CREATE
DROP
ALTER

### DML(データ操作言語)

SELECT
INSERT
UPDATE
DELETE

### DCL(データ制御言語)

COMMIT
ROLLBACK
GRANT
REVOKE