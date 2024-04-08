# coraza-http-wasm-benchmarks

Run benchmarks:

```shell
# plain httpbin
go run github.com/mccutchen/go-httpbin/v2/cmd/go-httpbin -port 8090
```

```shell
Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    8 109.3      6    7732
Processing:     1    7   1.7      7      18
Waiting:        0    7   1.5      6      15
Total:          3   15 109.4     13    7745

Percentage of the requests served within a certain time (ms)
  50%     13
  66%     14
  75%     15
  80%     16
  90%     17
  95%     19
  98%     21
  99%     21
 100%   7745 (longest request)
 ```

```shell
# coraza
go run coraza/main.go
```

```shell
Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.6      0       6
Processing:    18  214 160.3    191    1246
Waiting:       12  211 158.1    188    1246
Total:         18  214 160.3    191    1246

Percentage of the requests served within a certain time (ms)
  50%    191
  66%    249
  75%    285
  80%    307
  90%    388
  95%    511
  98%    659
  99%    829
 100%   1246 (longest request)
 ```

```shell
# coraza-http-wasm
go run coraza-http-wasm/main.go
```

```shell
Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.9      0      18
Processing:    17 1706 7769.4    287   76478
Waiting:       12 1694 7737.0    283   76478
Total:         17 1706 7769.9    287   76484

Percentage of the requests served within a certain time (ms)
  50%    287
  66%    470
  75%    700
  80%    943
  90%   1879
  95%   3488
  98%  26572
  99%  52233
 100%  76484 (longest request)
 ```

## Benchmark calls

```shell
ab -c 100 -n 5000 http://localhost:8090/
```
