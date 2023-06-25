# speee_tech_interview

## func fibonacci(num, n_1, n_2 int64) int64
フィボナッチ数を計算する関数。num番目のフィボナッチ数が返される。64bit整数。
```golang:example
ret := fibonacci(5,1,0)
```

## func fibonacciHandler(w http.ResponseWriter, r *http.Request)
リクエストを処理する関数。panicした場合はerror情報を返す。リクエストが正常でない場合もエラーコードで返す。