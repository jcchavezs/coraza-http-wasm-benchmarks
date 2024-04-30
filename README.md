# coraza-http-wasm-benchmarks

Run benchmarks:

```shell
# plain httpbin
go run httpbin/main.go
```

```shell
Concurrency Level:      100
Time taken for tests:   0.806 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      60645000 bytes
HTML transferred:       59195000 bytes
Requests per second:    6203.81 [#/sec] (mean)
Time per request:       16.119 [ms] (mean)
Time per request:       0.161 [ms] (mean, across all concurrent requests)
Transfer rate:          73482.46 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    8  14.2      6     130
Processing:     2    8  14.6      6     130
Waiting:        1    8  13.9      6     130
Total:          8   16  21.1     11     135

Percentage of the requests served within a certain time (ms)
  50%     11
  66%     12
  75%     12
  80%     13
  90%     14
  95%     16
  98%    104
  99%    135
 100%    135 (longest request)
 ```

```shell
# coraza
go run coraza/main.go
```

```shell
Concurrency Level:      100
Time taken for tests:   7.591 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      60645000 bytes
HTML transferred:       59195000 bytes
Requests per second:    658.65 [#/sec] (mean)
Time per request:       151.826 [ms] (mean)
Time per request:       1.518 [ms] (mean, across all concurrent requests)
Transfer rate:          7801.52 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   1.0      0       9
Processing:    14  150 115.5    128     874
Waiting:        8  148 113.4    126     874
Total:         14  150 115.5    128     874

Percentage of the requests served within a certain time (ms)
  50%    128
  66%    173
  75%    207
  80%    228
  90%    287
  95%    386
  98%    495
  99%    542
 100%    874 (longest request)
 ```

```shell
# coraza-http-wasm
go run coraza-http-wasm/main.go
```

```shell
Concurrency Level:      100
Time taken for tests:   72.779 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      60645000 bytes
HTML transferred:       59195000 bytes
Requests per second:    68.70 [#/sec] (mean)
Time per request:       1455.582 [ms] (mean)
Time per request:       14.556 [ms] (mean, across all concurrent requests)
Transfer rate:          813.75 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.9      0      31
Processing:    16 1453 6702.5    332   69471
Waiting:       16 1441 6684.1    329   69471
Total:         16 1453 6703.0    332   69477

Percentage of the requests served within a certain time (ms)
  50%    332
  66%    488
  75%    607
  80%    710
  90%   1195
  95%   2011
  98%  22645
  99%  45749
 100%  69477 (longest request)
 ```

## Benchmark calls

```shell
ab -c 100 -n 5000 http://localhost:8090/
```
