#!/bin/bash

# alias
if ! type "antlr" > /dev/null; then
  if type "antlr4" > /dev/null; then 
    alias antlr=antlr4
  fi
fi

# generate golang lib
antlr -Dlanguage=Go -package abslgo -visitor -o abslgo ./ABSLLexer.g4 
antlr -Dlanguage=Go -package abslgo -visitor -o abslgo ./ABSLParser.g4 

# generate javascript lib
antlr -Dlanguage=JavaScript -visitor -package absl -o js ./ABSLLexer.g4 
antlr -Dlanguage=JavaScript -visitor -package absl -o js ./ABSLParser.g4 