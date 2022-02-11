/*
 * @Auther: BinyGo
 * @Description:
 * @Date: 2022-02-11 01:04:33
 * @LastEditTime: 2022-02-11 01:38:26
 */
package compile

import (
	"fmt"
	"go/scanner"
	"go/token"
)

//go词法token化 go/go1.17.6/src/cmd/compile/internal/syntax/tokens.go

var z *int

func ExecSyntax() {
	src := []byte("cos(x) + 2i*sin(x) var s scanner.Scanner // Euler ")

	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}

	a := 1
	z = &a

}
