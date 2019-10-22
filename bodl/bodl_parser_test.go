// Code generated from ../BODLParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package bodl

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/assert"
)

func TestNewBODLParser(t *testing.T) {
	type args struct {
		input antlr.TokenStream
	}
	basicBOStream, _ := antlr.NewFileStream("./examples/basic.bo")
	tests := []struct {
		name string
		args args
	}{
		{
			"Basic BO Definition",
			args{antlr.NewCommonTokenStream(NewBODLLexer(basicBOStream), 0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
			parser := NewBODLParser(tt.args.input)
			program := parser.Program()
			children := program.GetChildren()
			assert.True(t, len(children) == 4)

		})
	}
}
