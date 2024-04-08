module github.com/jcchavezs/coraza-http-wasm-benchmarks/coraza-http-wasm

go 1.22.0

require (
	github.com/http-wasm/http-wasm-host-go v0.6.0
	github.com/jcchavezs/coraza-http-wasm-benchmarks v0.0.0-00010101000000-000000000000
	github.com/mccutchen/go-httpbin/v2 v2.13.4
)

replace github.com/jcchavezs/coraza-http-wasm-benchmarks => ../

require github.com/tetratelabs/wazero v1.7.0 // indirect
