#!/bin/bash

# generate golang lib
antlr4 -Dlanguage=Go -package abslgo -visitor -o abslgo ./ABSLLexer.g4 
antlr4 -Dlanguage=Go -package abslgo -visitor -o abslgo ./ABSLParser.g4 

# generate javascript lib
antlr4 -Dlanguage=JavaScript -visitor -package absl -o js ./ABSLLexer.g4 
antlr4 -Dlanguage=JavaScript -visitor -package absl -o js ./ABSLParser.g4 