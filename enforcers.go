package ggi

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type enforcers map[string][]string


// GetEnforcers links all required interface implementations.
// eg. "interfaceName" => ["SomeStruct", "otherStruct", "anotherType"]
//
// By enforcer, I refer to "var _ interfaceName = (*SomeStruct)(nil)"
// as you can not compile the source code before that implementation exist.
//
//  package main
//
//  import (
//    "fmt"
//    "path/filepath"
//    "runtime"
//
//    "github.com/andersfylling/ggi"
//  )
//
//  var (
//    _, b, _, _ = runtime.Caller(0)
//    basepath   = filepath.Dir(b)
//    genpath    = "/generate/testing"
//  )
//
//  func main() {
//    path := basepath[:len(basepath)-len(genpath)]
//    files, err := ggi.GetFiles(path)
//    if err != nil {
//      panic(err)
//    }
//
//    enforcers := ggi.GetEnforcers(files)
//
//    // TODO: ensure the enforcer and type names correlate as desired
//    for k, v := range enforcers {
//      fmt.Println(k, v)
//    }
//  }
func GetEnforcers(files []string) (es enforcers){
	es = make(enforcers)

	for _, file := range files {
		file, err := parser.ParseFile(token.NewFileSet(), file, nil, 0)
		if err != nil {
			panic(err)
		}

		addEnforcers(es, file)
	}

	return es
}

func addEnforcers(es enforcers, file *ast.File) {
	for _, item := range file.Decls {
		var gdecl *ast.GenDecl
		var ok bool
		if gdecl, ok = item.(*ast.GenDecl); !ok {
			continue
		}

		if gdecl.Tok != token.VAR {
			continue
		}

		specs := item.(*ast.GenDecl).Specs
		for i := range specs {
			vs := specs[i].(*ast.ValueSpec)
			if len(vs.Names) == 0 || vs.Names[0].Name != "_" {
				continue
			}

			var cExpr *ast.CallExpr
			if cExpr, ok = vs.Values[0].(*ast.CallExpr); !ok {
				continue
			}

			var pExpr *ast.ParenExpr
			if pExpr, ok = cExpr.Fun.(*ast.ParenExpr); !ok {
				continue
			}

			var sExpr *ast.StarExpr
			if sExpr, ok = pExpr.X.(*ast.StarExpr); !ok {
				continue
			}

			var id *ast.Ident
			if id, ok = sExpr.X.(*ast.Ident); !ok {
				continue
			}

			var id2 *ast.Ident
			if id2, ok = vs.Type.(*ast.Ident); !ok {
				continue
			}

			enf := id2.Name
			typeName := id.Name
			es[enf] = append(es[enf], typeName)
		}
	}
}