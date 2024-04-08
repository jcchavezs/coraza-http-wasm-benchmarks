package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "embed"

	"github.com/http-wasm/http-wasm-host-go/handler"
	wasm "github.com/http-wasm/http-wasm-host-go/handler/nethttp"
	benchmarks "github.com/jcchavezs/coraza-http-wasm-benchmarks"
	"github.com/mccutchen/go-httpbin/v2/httpbin"
)

//go:embed coraza-http-wasm.wasm
var guest []byte

func main() {
	app := httpbin.New()

	conf := fmt.Sprintf(`{"directives": ["%s"]}`, strings.Join(strings.Split(benchmarks.Directives, "\n"), `", "`))

	mw, err := wasm.NewMiddleware(context.Background(), guest, handler.GuestConfig([]byte(conf)))
	if err != nil {
		log.Fatal(err)
	}

	h := mw.NewHandler(context.Background(), app.Handler())

	http.Handle("/", h)

	fmt.Println("Server is running. Listening port: 8090")

	log.Fatal(http.ListenAndServe(":8090", nil))
}
