package main

import (
	"fmt"
	"log"
	"net/http"

	coreruleset "github.com/corazawaf/coraza-coreruleset/v4"
	"github.com/corazawaf/coraza/v3"
	txhttp "github.com/corazawaf/coraza/v3/http"
	"github.com/corazawaf/coraza/v3/types"
	benchmarks "github.com/jcchavezs/coraza-http-wasm-benchmarks"
	"github.com/mccutchen/go-httpbin/v2/httpbin"
)

func main() {
	waf := createWAF()

	app := httpbin.New()

	http.Handle("/", txhttp.WrapHandler(waf, app.Handler()))

	fmt.Println("Server is running. Listening port: 8090")

	log.Fatal(http.ListenAndServe(":8090", nil))
}

func createWAF() coraza.WAF {
	waf, err := coraza.NewWAF(
		coraza.NewWAFConfig().
			WithErrorCallback(logError).
			WithDirectives(benchmarks.Directives).WithRootFS(coreruleset.FS),
	)
	if err != nil {
		log.Fatal(err)
	}
	return waf
}

func logError(error types.MatchedRule) {
	msg := error.ErrorLog()
	fmt.Printf("[logError][%s] %s\n", error.Rule().Severity(), msg)
}
