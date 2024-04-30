module github.com/jcchavezs/coraza-http-wasm-benchmarks/coraza

go 1.22.0

require (
	github.com/corazawaf/coraza/v3 v3.1.1-0.20240403190642-dc778126cf45
	github.com/mccutchen/go-httpbin/v2 v2.13.4
)

replace github.com/jcchavezs/coraza-http-wasm-benchmarks => ../
