package bodl

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/assert"
)

func TestParserBasicBODL(t *testing.T) {
	basicBOStream, _ := antlr.NewFileStream("./examples/basic.bo")
	input := antlr.NewCommonTokenStream(NewBODLLexer(basicBOStream), 0)
	parser := NewBODLParser(input)
	program := parser.Program()
	children := program.GetChildren()
	assert.True(t, len(children) == 4)
}

func TestParserComplexBODL(t *testing.T) {
	complexBOStream, _ := antlr.NewFileStream("./examples/complex.bo")
	input := antlr.NewCommonTokenStream(NewBODLLexer(complexBOStream), 0)
	parser := NewBODLParser(input)
	program := parser.Program()
	children := program.GetChildren()
	assert.True(t, len(children) == 5)

}
