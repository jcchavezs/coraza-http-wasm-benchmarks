module github.com/jcchavezs/coraza-http-wasm-benchmarks/coraza

go 1.22.0

require (
	github.com/corazawaf/coraza-coreruleset/v4 v4.1.0
	github.com/corazawaf/coraza/v3 v3.1.1-0.20240403190642-dc778126cf45
	github.com/jcchavezs/coraza-http-wasm-benchmarks v0.0.0-00010101000000-000000000000
	github.com/mccutchen/go-httpbin/v2 v2.13.4
)

replace github.com/jcchavezs/coraza-http-wasm-benchmarks => ../

require (
	github.com/corazawaf/libinjection-go v0.1.3 // indirect
	github.com/magefile/mage v1.15.0 // indirect
	github.com/petar-dambovaliev/aho-corasick v0.0.0-20230725210150-fb29fc3c913e // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/tidwall/gjson v1.17.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	rsc.io/binaryregexp v0.2.0 // indirect
)
