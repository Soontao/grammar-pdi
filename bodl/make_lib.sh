#!/bin/bash

# generate golang lib
antlr4 -Dlanguage=Go -package bodlgo -visitor -o bodlgo ./BODLLexer.g4 
antlr4 -Dlanguage=Go -package bodlgo -visitor -o bodlgo ./BODLParser.g4 

# generate javascript lib
antlr4 -Dlanguage=JavaScript -visitor -package bodl -o js ./BODLLexer.g4 
antlr4 -Dlanguage=JavaScript -visitor -package bodl -o js ./BODLParser.g4 