package ast_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/ast/traverser"
	"github.com/z7zmey/php-parser/pkg/ast/visitor"
	"os"
	"strings"
)

func ExampleJSON() {
	stxTree := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.Nullable{
				Expr: &ast.Parameter{
					Type:         nil,
					Var:          nil,
					DefaultValue: nil,
				},
			},
			&ast.Identifier{},
			&ast.ArgumentList{
				Arguments: []ast.Vertex{
					&ast.Argument{},
					&ast.Argument{
						Expr: &ast.ScalarDnumber{},
					},
				},
			},
		},
	}

	jsonStxTree, err := json.Marshal(stxTree)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(nil)
	err = json.Indent(buf, jsonStxTree, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stdout, buf.String())

	// output:
	// {
	//   "FreeFloating": null,
	//   "Position": null,
	//   "Stmts": [
	//     {
	//       "FreeFloating": null,
	//       "Position": null,
	//       "Expr": {
	//         "FreeFloating": null,
	//         "Position": null,
	//         "ByRef": false,
	//         "Variadic": false,
	//         "Type": null,
	//         "Var": null,
	//         "DefaultValue": null
	//       }
	//     },
	//     {
	//       "FreeFloating": null,
	//       "Position": null,
	//       "Value": ""
	//     },
	//     {
	//       "FreeFloating": null,
	//       "Position": null,
	//       "Arguments": [
	//         {
	//           "FreeFloating": null,
	//           "Position": null,
	//           "Variadic": false,
	//           "IsReference": false,
	//           "Expr": null
	//         },
	//         {
	//           "FreeFloating": null,
	//           "Position": null,
	//           "Variadic": false,
	//           "IsReference": false,
	//           "Expr": {
	//             "FreeFloating": null,
	//             "Position": null,
	//             "Value": ""
	//           }
	//         }
	//       ]
	//     }
	//   ]
	// }
}

func ExampleStxTree() {
	stxTree := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.Nullable{
				Expr: &ast.Parameter{
					Type:         nil,
					Var:          nil,
					DefaultValue: nil,
				},
			},
			&ast.Identifier{},
			&ast.ArgumentList{
				Arguments: []ast.Vertex{
					&ast.Argument{},
					&ast.Argument{
						Expr: &ast.ScalarDnumber{},
					},
				},
			},
		},
	}

	traverser.NewDFS(&testVisitor{}).Traverse(stxTree)

	//output:
	//=>  *ast.Root
	//=>    Stmts:
	//=>      *ast.Nullable
	//=>        Expr:
	//=>          *ast.Parameter
	//=>      *ast.Identifier
	//=>      *ast.ArgumentList
	//=>        Arguments:
	//=>          *ast.Argument
	//=>          *ast.Argument
	//=>            Expr:
	//=>              *ast.ScalarDnumber
}

type testVisitor struct {
	visitor.Null
	depth int
}


func (v *testVisitor) Enter(key string, _ bool) {
	v.depth++
	fmt.Fprint(os.Stdout, "=>", strings.Repeat("  ", v.depth), key, ":\n")
}

func (v *testVisitor) Leave(key string, _ bool) {
	v.depth--
}

func (v *testVisitor) EnterNode(n ast.Vertex) bool {
	v.depth++
	n.Accept(v)

	return true
}

func (v *testVisitor) LeaveNode(_ ast.Vertex) {
	v.depth--
}

func (v *testVisitor) Root(_ *ast.Root) {
	fmt.Fprintln(os.Stdout, "=>", strings.Repeat("  ", v.depth-1), "*ast.Root")
}

func (v *testVisitor) Nullable(_ *ast.Nullable) {
	fmt.Fprintln(os.Stdout, "=>", strings.Repeat("  ", v.depth-1), "*ast.Nullable")
}

func (v *testVisitor) Parameter(_ *ast.Parameter) {
	fmt.Fprintln(os.Stdout, "=>", strings.Repeat("  ", v.depth-1), "*ast.Parameter")
}

func (v *testVisitor) Identifier(_ *ast.Identifier) {
	fmt.Fprintln(os.Stdout, "=>", strings.Repeat("  ", v.depth-1), "*ast.Identifier")
}

func (v *testVisitor) ArgumentList(_ *ast.ArgumentList)  {
	fmt.Fprintln(os.Stdout, "=>", strings.Repeat("  ", v.depth-1), "*ast.ArgumentList")
}

func (v *testVisitor) Argument(_ *ast.Argument)  {
	fmt.Fprintln(os.Stdout, "=>", strings.Repeat("  ", v.depth-1), "*ast.Argument")
}

func (v *testVisitor) ScalarDnumber(_ *ast.ScalarDnumber)  {
	fmt.Fprintln(os.Stdout, "=>", strings.Repeat("  ", v.depth-1), "*ast.ScalarDnumber")
}