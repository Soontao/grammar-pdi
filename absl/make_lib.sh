#!/bin/bash

# generate golang lib
antlr -Dlanguage=Go -package abslgo -visitor -o abslgo ./ABSLLexer.g4 
antlr -Dlanguage=Go -package abslgo -visitor -o abslgo ./ABSLParser.g4 

# generate javascript lib
antlr -Dlanguage=JavaScript -visitor -package absl -o js ./ABSLLexer.g4 
antlr -Dlanguage=JavaScript -visitor -package absl -o js ./ABSLParser.g4 