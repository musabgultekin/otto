package otto

import (
	"github.com/musabgultekin/otto/ast"
	"github.com/musabgultekin/otto/file"
)

type _file struct {
	name string
	src  string
	base int // This will always be 1 or greater
}

type _compiler struct {
	file    *file.File
	program *ast.Program
}

func (cmpl *_compiler) parse() *_nodeProgram {
	if cmpl.program != nil {
		cmpl.file = cmpl.program.File
	}
	return cmpl._parse(cmpl.program)
}
