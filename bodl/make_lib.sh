#!/bin/bash

# alias
if ! type "antlr" > /dev/null; then
  if type "antlr4" > /dev/null; then 
    alias antlr=antlr4
  fi
fi

# generate golang lib
antlr -Dlanguage=Go -package bodlgo -visitor -o bodlgo ./BODLLexer.g4 
antlr -Dlanguage=Go -package bodlgo -visitor -o bodlgo ./BODLParser.g4 

# generate javascript lib
antlr -Dlanguage=JavaScript -visitor -package bodl -o js ./BODLLexer.g4 
antlr -Dlanguage=JavaScript -visitor -package bodl -o js ./BODLParser.g4 