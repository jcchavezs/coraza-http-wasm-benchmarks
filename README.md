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
Time taken for tests:   9.784 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      60645000 bytes
HTML transferred:       59195000 bytes
Requests per second:    511.02 [#/sec] (mean)
Time per request:       195.686 [ms] (mean)
Time per request:       1.957 [ms] (mean, across all concurrent requests)
Transfer rate:          6052.92 [Kbytes/sec] received

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
Concurrency Level:      100
Time taken for tests:   75.901 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      60645000 bytes
HTML transferred:       59195000 bytes
Requests per second:    65.88 [#/sec] (mean)
Time per request:       1518.020 [ms] (mean)
Time per request:       15.180 [ms] (mean, across all concurrent requests)
Transfer rate:          780.27 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.9      0      12
Processing:    17 1515 6871.5    218   66850
Waiting:       17 1509 6860.4    216   66850
Total:         17 1516 6872.1    218   66856

Percentage of the requests served within a certain time (ms)
  50%    218
  66%    342
  75%    491
  80%    628
  90%   1910
  95%   2976
  98%  31504
  99%  45067
 100%  66856 (longest request)
 ```

## Benchmark calls

```shell
ab -c 100 -n 5000 http://localhost:8090/
```
