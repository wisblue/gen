// The gen tool automatically generates routing code and openapi3 documents
// Gen can also generate client code
//
// Basic usage:
//   1. Install gen tool `go get -u -v github.com/wzshiming/gen/cmd/gen`
//   2. Add gen tool to $PATH
//   3. Execute it `gen run github.com/wzshiming/gen/examples/group/service`
//   4. Open http://127.0.0.1:8080/swagger/?url=./openapi.json# with your browser.

//go:generate gen route github.com/wzshiming/gen/examples/group/service
//go:generate gen openapi -o ../openapi/openapi.json github.com/wzshiming/gen/examples/group/service
//go:generate gen client -o ../client/client_gen.go github.com/wzshiming/gen/examples/group/service

package service
