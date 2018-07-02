package route

import (
	"strings"

	"github.com/wzshiming/gen/model"
	"github.com/wzshiming/gen/spec"
	"github.com/wzshiming/gen/srcgen"
)

// GenClient is the generating generating
type GenRoute struct {
	api *spec.API
	buf *srcgen.File
	model.GenModel
}

func NewGenRoute(api *spec.API) *GenRoute {
	buf := &srcgen.File{}
	return &GenRoute{
		api:      api,
		buf:      buf,
		GenModel: *model.NewGenModel(api, buf),
	}
}

func (g *GenRoute) Generate(funcName string) (*srcgen.File, error) {
	g.buf.WithPackname(g.api.Package)
	err := g.GenerateRoutes(funcName)
	if err != nil {
		return nil, err
	}

	return g.buf, nil
}

func (g *GenRoute) GenerateRoutes(funcName string) (err error) {
	g.buf.AddImport("", "github.com/gorilla/mux")
	g.buf.AddImport("", "net/http")

	m := map[string]bool{}
	for _, v := range g.api.Operations {
		err = g.GenerateRouteTypes(v, m)
		if err != nil {
			return err
		}
	}

	g.buf.WriteFormat(`
// %s is generated do not edit.
func %s() http.Handler {
	router := mux.NewRouter()

`, funcName, funcName)
	for _, v := range g.api.Operations {
		err = g.GenerateRoute(v)
		if err != nil {
			return err
		}
	}
	g.buf.WriteString(`
	return router
}
`)

	for _, v := range g.api.Requests {
		switch v.In {
		case "security":
		default:
			err = g.GenerateRequestFunction(v)
			if err != nil {
				return err
			}
		}
	}

	for _, v := range g.api.Securitys {
		err = g.GenerateSecurityFunction(v)
		if err != nil {
			return err
		}
	}

	for _, v := range g.api.Operations {
		err = g.GenerateOperationFunction(v)
		if err != nil {
			return err
		}
	}

	return
}

func (g *GenRoute) GenerateRouteTypes(oper *spec.Operation, m map[string]bool) (err error) {
	if oper.Type == nil {
		return
	}
	typ := oper.Type
	if typ.Ref != "" {
		typ = g.api.Types[oper.Type.Ref]
	}
	if m[typ.Name] {
		return
	}
	g.buf.WriteFormat("// %s Define the method scope\n", typ.Name)
	g.buf.WriteFormat("var %s %s\n", GetGlobalVarName(typ.Name), typ.Name)
	m[typ.Name] = true
	return
}

func (g *GenRoute) GenerateRoute(oper *spec.Operation) (err error) {
	name := GetOperationFunctionName(oper.Name)
	g.buf.WriteFormat(`
	// Registered routing %s %s
	router.Path("%s").
		Methods("%s").
		HandlerFunc(`, strings.ToUpper(oper.Method), oper.Path, oper.Path, strings.ToUpper(oper.Method))
	if oper.Type != nil {
		typ := oper.Type
		if typ.Ref != "" {
			typ = g.api.Types[oper.Type.Ref]
		}
		g.buf.WriteFormat("%s.", GetGlobalVarName(typ.Name))
	}
	g.buf.WriteFormat(`%s)
`, name)
	return
}
