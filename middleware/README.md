# middleware

gin の処理手順

```
Request -> Route Parser -> Middleware -> Route Handler -> Middleware -> Response
```

ルートハンドラーの直前などに処理を行う必要がある処理を記述する
