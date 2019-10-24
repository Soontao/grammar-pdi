#!/bin/bash

# generate golang lib
antlr -Dlanguage=Go -package bodlgo -visitor -o bodlgo ./BODLLexer.g4 
antlr -Dlanguage=Go -package bodlgo -visitor -o bodlgo ./BODLParser.g4 

# generate javascript lib
antlr -Dlanguage=JavaScript -visitor -package bodl -o js ./BODLLexer.g4 
antlr -Dlanguage=JavaScript -visitor -package bodl -o js ./BODLParser.g4 