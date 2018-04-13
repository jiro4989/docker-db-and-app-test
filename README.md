# docker-db-and-app-test
docker-composeでgolangアプリコンテナからmysqlコンテナのデータを参照するテスト

## 使い方
下記のコマンドを実行する。

```bash
docker-compose up -d
```

### Golangアプリを実行してMySQLデータを確認
コンテナに入るために下記のコマンドを実行する。

```bash
docker exec -it test_app bash
```

アプリのディレクトリに移動するために下記のコマンドを実行する。

```bash
cd src/app
```

アプリを実行する。  
user名とpasswordが出力されれば成功。

```bash
go run main.go
```

### MySQLの初期データの確認
コンテナに入るために下記のコマンドを実行する。

```bash
docker exec -it test_db bash
```

MySQL CLIにログインするために下記のコマンドを実行する。
パスワードは password

```bash
mysql -u root -p
```

テスト用に追加したDBのtest_dbの存在を確認するために下記のコマンドを実行する。

```sql
show databases;
```
