// Code generated from ./BODLParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package bodlgo // BODLParser
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 59, 362,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 74, 10, 2, 3, 3, 7, 3, 77, 10,
	3, 12, 3, 14, 3, 80, 11, 3, 3, 4, 5, 4, 83, 10, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 5, 4, 90, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 98,
	10, 4, 3, 5, 7, 5, 101, 10, 5, 12, 5, 14, 5, 104, 11, 5, 3, 6, 5, 6, 107,
	10, 6, 3, 6, 5, 6, 110, 10, 6, 3, 6, 3, 6, 3, 6, 5, 6, 115, 10, 6, 3, 6,
	5, 6, 118, 10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 7, 8, 132, 10, 8, 12, 8, 14, 8, 135, 11, 8, 3, 9, 5,
	9, 138, 10, 9, 3, 9, 5, 9, 141, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 10, 5, 10, 150, 10, 10, 3, 10, 5, 10, 153, 10, 10, 3, 10, 3, 10,
	3, 10, 5, 10, 158, 10, 10, 3, 10, 3, 10, 3, 11, 5, 11, 163, 10, 11, 3,
	11, 5, 11, 166, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 12, 5, 12, 177, 10, 12, 3, 12, 5, 12, 180, 10, 12, 3, 12, 3,
	12, 3, 12, 5, 12, 185, 10, 12, 3, 12, 5, 12, 188, 10, 12, 3, 12, 5, 12,
	191, 10, 12, 3, 12, 5, 12, 194, 10, 12, 3, 13, 5, 13, 197, 10, 13, 3, 13,
	5, 13, 200, 10, 13, 3, 13, 3, 13, 3, 13, 5, 13, 205, 10, 13, 3, 13, 3,
	13, 3, 13, 5, 13, 210, 10, 13, 3, 13, 5, 13, 213, 10, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 5, 16, 230, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18,
	5, 18, 237, 10, 18, 3, 18, 5, 18, 240, 10, 18, 3, 19, 7, 19, 243, 10, 19,
	12, 19, 14, 19, 246, 11, 19, 3, 20, 5, 20, 249, 10, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 5, 20, 256, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 5, 20, 266, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 5, 20, 275, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	5, 21, 282, 10, 21, 3, 22, 3, 22, 5, 22, 286, 10, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 5, 22, 292, 10, 22, 5, 22, 294, 10, 22, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 303, 10, 23, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 5, 24, 310, 10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5,
	26, 328, 10, 26, 3, 27, 3, 27, 3, 27, 7, 27, 333, 10, 27, 12, 27, 14, 27,
	336, 11, 27, 3, 28, 3, 28, 3, 28, 7, 28, 341, 10, 28, 12, 28, 14, 28, 344,
	11, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 7, 31, 351, 10, 31, 12, 31,
	14, 31, 354, 11, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 2,
	2, 35, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 2, 7, 3,
	2, 4, 15, 4, 2, 3, 3, 52, 53, 3, 2, 17, 18, 3, 2, 41, 46, 3, 2, 47, 48,
	2, 380, 2, 73, 3, 2, 2, 2, 4, 78, 3, 2, 2, 2, 6, 97, 3, 2, 2, 2, 8, 102,
	3, 2, 2, 2, 10, 106, 3, 2, 2, 2, 12, 121, 3, 2, 2, 2, 14, 133, 3, 2, 2,
	2, 16, 137, 3, 2, 2, 2, 18, 149, 3, 2, 2, 2, 20, 162, 3, 2, 2, 2, 22, 176,
	3, 2, 2, 2, 24, 196, 3, 2, 2, 2, 26, 216, 3, 2, 2, 2, 28, 219, 3, 2, 2,
	2, 30, 229, 3, 2, 2, 2, 32, 231, 3, 2, 2, 2, 34, 236, 3, 2, 2, 2, 36, 244,
	3, 2, 2, 2, 38, 274, 3, 2, 2, 2, 40, 281, 3, 2, 2, 2, 42, 293, 3, 2, 2,
	2, 44, 302, 3, 2, 2, 2, 46, 309, 3, 2, 2, 2, 48, 311, 3, 2, 2, 2, 50, 327,
	3, 2, 2, 2, 52, 329, 3, 2, 2, 2, 54, 337, 3, 2, 2, 2, 56, 345, 3, 2, 2,
	2, 58, 347, 3, 2, 2, 2, 60, 352, 3, 2, 2, 2, 62, 355, 3, 2, 2, 2, 64, 357,
	3, 2, 2, 2, 66, 359, 3, 2, 2, 2, 68, 74, 7, 2, 2, 3, 69, 70, 5, 4, 3, 2,
	70, 71, 5, 8, 5, 2, 71, 72, 7, 2, 2, 3, 72, 74, 3, 2, 2, 2, 73, 68, 3,
	2, 2, 2, 73, 69, 3, 2, 2, 2, 74, 3, 3, 2, 2, 2, 75, 77, 5, 6, 4, 2, 76,
	75, 3, 2, 2, 2, 77, 80, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2,
	2, 79, 5, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 81, 83, 5, 60, 31, 2, 82, 81,
	3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 7, 13, 2, 2,
	85, 86, 5, 52, 27, 2, 86, 87, 7, 26, 2, 2, 87, 98, 3, 2, 2, 2, 88, 90,
	5, 60, 31, 2, 89, 88, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 91, 3, 2, 2,
	2, 91, 92, 7, 13, 2, 2, 92, 93, 5, 52, 27, 2, 93, 94, 7, 14, 2, 2, 94,
	95, 5, 66, 34, 2, 95, 96, 7, 26, 2, 2, 96, 98, 3, 2, 2, 2, 97, 82, 3, 2,
	2, 2, 97, 89, 3, 2, 2, 2, 98, 7, 3, 2, 2, 2, 99, 101, 5, 10, 6, 2, 100,
	99, 3, 2, 2, 2, 101, 104, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 103, 3,
	2, 2, 2, 103, 9, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 105, 107, 5, 60, 31,
	2, 106, 105, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 109, 3, 2, 2, 2, 108,
	110, 5, 36, 19, 2, 109, 108, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 111,
	3, 2, 2, 2, 111, 114, 7, 4, 2, 2, 112, 115, 5, 66, 34, 2, 113, 115, 5,
	42, 22, 2, 114, 112, 3, 2, 2, 2, 114, 113, 3, 2, 2, 2, 115, 117, 3, 2,
	2, 2, 116, 118, 5, 34, 18, 2, 117, 116, 3, 2, 2, 2, 117, 118, 3, 2, 2,
	2, 118, 119, 3, 2, 2, 2, 119, 120, 5, 12, 7, 2, 120, 11, 3, 2, 2, 2, 121,
	122, 7, 24, 2, 2, 122, 123, 5, 14, 8, 2, 123, 124, 5, 60, 31, 2, 124, 125,
	7, 25, 2, 2, 125, 13, 3, 2, 2, 2, 126, 132, 5, 16, 9, 2, 127, 132, 5, 20,
	11, 2, 128, 132, 5, 22, 12, 2, 129, 132, 5, 18, 10, 2, 130, 132, 5, 24,
	13, 2, 131, 126, 3, 2, 2, 2, 131, 127, 3, 2, 2, 2, 131, 128, 3, 2, 2, 2,
	131, 129, 3, 2, 2, 2, 131, 130, 3, 2, 2, 2, 132, 135, 3, 2, 2, 2, 133,
	131, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 15, 3, 2, 2, 2, 135, 133, 3,
	2, 2, 2, 136, 138, 5, 60, 31, 2, 137, 136, 3, 2, 2, 2, 137, 138, 3, 2,
	2, 2, 138, 140, 3, 2, 2, 2, 139, 141, 5, 36, 19, 2, 140, 139, 3, 2, 2,
	2, 140, 141, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 143, 7, 7, 2, 2, 143,
	144, 5, 66, 34, 2, 144, 145, 7, 30, 2, 2, 145, 146, 5, 42, 22, 2, 146,
	147, 7, 26, 2, 2, 147, 17, 3, 2, 2, 2, 148, 150, 5, 60, 31, 2, 149, 148,
	3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 152, 3, 2, 2, 2, 151, 153, 5, 36,
	19, 2, 152, 151, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2,
	154, 155, 7, 9, 2, 2, 155, 157, 5, 66, 34, 2, 156, 158, 5, 34, 18, 2, 157,
	156, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 160,
	7, 26, 2, 2, 160, 19, 3, 2, 2, 2, 161, 163, 5, 60, 31, 2, 162, 161, 3,
	2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 165, 3, 2, 2, 2, 164, 166, 5, 36, 19,
	2, 165, 164, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167,
	168, 7, 10, 2, 2, 168, 169, 5, 66, 34, 2, 169, 170, 7, 15, 2, 2, 170, 171,
	7, 52, 2, 2, 171, 172, 7, 30, 2, 2, 172, 173, 5, 40, 21, 2, 173, 174, 7,
	26, 2, 2, 174, 21, 3, 2, 2, 2, 175, 177, 5, 60, 31, 2, 176, 175, 3, 2,
	2, 2, 176, 177, 3, 2, 2, 2, 177, 179, 3, 2, 2, 2, 178, 180, 5, 36, 19,
	2, 179, 178, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181,
	182, 7, 8, 2, 2, 182, 184, 5, 66, 34, 2, 183, 185, 5, 50, 26, 2, 184, 183,
	3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 187, 3, 2, 2, 2, 186, 188, 5, 34,
	18, 2, 187, 186, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 190, 3, 2, 2, 2,
	189, 191, 5, 12, 7, 2, 190, 189, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191,
	193, 3, 2, 2, 2, 192, 194, 7, 26, 2, 2, 193, 192, 3, 2, 2, 2, 193, 194,
	3, 2, 2, 2, 194, 23, 3, 2, 2, 2, 195, 197, 5, 60, 31, 2, 196, 195, 3, 2,
	2, 2, 196, 197, 3, 2, 2, 2, 197, 199, 3, 2, 2, 2, 198, 200, 5, 36, 19,
	2, 199, 198, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201,
	202, 7, 6, 2, 2, 202, 204, 5, 66, 34, 2, 203, 205, 5, 50, 26, 2, 204, 203,
	3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 207, 7, 5,
	2, 2, 207, 209, 5, 52, 27, 2, 208, 210, 5, 26, 14, 2, 209, 208, 3, 2, 2,
	2, 209, 210, 3, 2, 2, 2, 210, 212, 3, 2, 2, 2, 211, 213, 5, 28, 15, 2,
	212, 211, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214,
	215, 7, 26, 2, 2, 215, 25, 3, 2, 2, 2, 216, 217, 7, 12, 2, 2, 217, 218,
	5, 66, 34, 2, 218, 27, 3, 2, 2, 2, 219, 220, 7, 16, 2, 2, 220, 221, 7,
	22, 2, 2, 221, 222, 5, 30, 16, 2, 222, 223, 7, 23, 2, 2, 223, 29, 3, 2,
	2, 2, 224, 230, 5, 32, 17, 2, 225, 226, 5, 32, 17, 2, 226, 227, 5, 64,
	33, 2, 227, 228, 5, 30, 16, 2, 228, 230, 3, 2, 2, 2, 229, 224, 3, 2, 2,
	2, 229, 225, 3, 2, 2, 2, 230, 31, 3, 2, 2, 2, 231, 232, 5, 66, 34, 2, 232,
	233, 5, 62, 32, 2, 233, 234, 5, 58, 30, 2, 234, 33, 3, 2, 2, 2, 235, 237,
	7, 11, 2, 2, 236, 235, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 239, 3, 2,
	2, 2, 238, 240, 5, 54, 28, 2, 239, 238, 3, 2, 2, 2, 239, 240, 3, 2, 2,
	2, 240, 35, 3, 2, 2, 2, 241, 243, 5, 38, 20, 2, 242, 241, 3, 2, 2, 2, 243,
	246, 3, 2, 2, 2, 244, 242, 3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 37, 3,
	2, 2, 2, 246, 244, 3, 2, 2, 2, 247, 249, 7, 19, 2, 2, 248, 247, 3, 2, 2,
	2, 248, 249, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250, 251, 7, 20, 2, 2, 251,
	252, 5, 66, 34, 2, 252, 253, 7, 21, 2, 2, 253, 275, 3, 2, 2, 2, 254, 256,
	7, 19, 2, 2, 255, 254, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 257, 3, 2,
	2, 2, 257, 258, 7, 20, 2, 2, 258, 259, 5, 66, 34, 2, 259, 260, 7, 22, 2,
	2, 260, 261, 5, 66, 34, 2, 261, 262, 7, 23, 2, 2, 262, 263, 7, 21, 2, 2,
	263, 275, 3, 2, 2, 2, 264, 266, 7, 19, 2, 2, 265, 264, 3, 2, 2, 2, 265,
	266, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267, 268, 7, 20, 2, 2, 268, 269,
	5, 66, 34, 2, 269, 270, 7, 22, 2, 2, 270, 271, 5, 58, 30, 2, 271, 272,
	7, 23, 2, 2, 272, 273, 7, 21, 2, 2, 273, 275, 3, 2, 2, 2, 274, 248, 3,
	2, 2, 2, 274, 255, 3, 2, 2, 2, 274, 265, 3, 2, 2, 2, 275, 39, 3, 2, 2,
	2, 276, 282, 5, 42, 22, 2, 277, 278, 5, 42, 22, 2, 278, 279, 7, 27, 2,
	2, 279, 280, 5, 40, 21, 2, 280, 282, 3, 2, 2, 2, 281, 276, 3, 2, 2, 2,
	281, 277, 3, 2, 2, 2, 282, 41, 3, 2, 2, 2, 283, 285, 5, 66, 34, 2, 284,
	286, 5, 44, 23, 2, 285, 284, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 294,
	3, 2, 2, 2, 287, 288, 5, 52, 27, 2, 288, 289, 7, 30, 2, 2, 289, 291, 5,
	66, 34, 2, 290, 292, 5, 44, 23, 2, 291, 290, 3, 2, 2, 2, 291, 292, 3, 2,
	2, 2, 292, 294, 3, 2, 2, 2, 293, 283, 3, 2, 2, 2, 293, 287, 3, 2, 2, 2,
	294, 43, 3, 2, 2, 2, 295, 296, 7, 28, 2, 2, 296, 303, 5, 58, 30, 2, 297,
	298, 7, 28, 2, 2, 298, 299, 7, 24, 2, 2, 299, 300, 5, 46, 24, 2, 300, 301,
	7, 25, 2, 2, 301, 303, 3, 2, 2, 2, 302, 295, 3, 2, 2, 2, 302, 297, 3, 2,
	2, 2, 303, 45, 3, 2, 2, 2, 304, 310, 5, 48, 25, 2, 305, 306, 5, 48, 25,
	2, 306, 307, 7, 27, 2, 2, 307, 308, 5, 46, 24, 2, 308, 310, 3, 2, 2, 2,
	309, 304, 3, 2, 2, 2, 309, 305, 3, 2, 2, 2, 310, 47, 3, 2, 2, 2, 311, 312,
	5, 66, 34, 2, 312, 313, 7, 28, 2, 2, 313, 314, 5, 58, 30, 2, 314, 49, 3,
	2, 2, 2, 315, 316, 7, 20, 2, 2, 316, 317, 5, 58, 30, 2, 317, 318, 7, 27,
	2, 2, 318, 319, 5, 58, 30, 2, 319, 320, 7, 21, 2, 2, 320, 328, 3, 2, 2,
	2, 321, 322, 7, 20, 2, 2, 322, 323, 5, 58, 30, 2, 323, 324, 7, 27, 2, 2,
	324, 325, 7, 50, 2, 2, 325, 326, 7, 21, 2, 2, 326, 328, 3, 2, 2, 2, 327,
	315, 3, 2, 2, 2, 327, 321, 3, 2, 2, 2, 328, 51, 3, 2, 2, 2, 329, 334, 5,
	66, 34, 2, 330, 331, 7, 31, 2, 2, 331, 333, 5, 52, 27, 2, 332, 330, 3,
	2, 2, 2, 333, 336, 3, 2, 2, 2, 334, 332, 3, 2, 2, 2, 334, 335, 3, 2, 2,
	2, 335, 53, 3, 2, 2, 2, 336, 334, 3, 2, 2, 2, 337, 342, 5, 66, 34, 2, 338,
	339, 7, 27, 2, 2, 339, 341, 5, 66, 34, 2, 340, 338, 3, 2, 2, 2, 341, 344,
	3, 2, 2, 2, 342, 340, 3, 2, 2, 2, 342, 343, 3, 2, 2, 2, 343, 55, 3, 2,
	2, 2, 344, 342, 3, 2, 2, 2, 345, 346, 9, 2, 2, 2, 346, 57, 3, 2, 2, 2,
	347, 348, 9, 3, 2, 2, 348, 59, 3, 2, 2, 2, 349, 351, 9, 4, 2, 2, 350, 349,
	3, 2, 2, 2, 351, 354, 3, 2, 2, 2, 352, 350, 3, 2, 2, 2, 352, 353, 3, 2,
	2, 2, 353, 61, 3, 2, 2, 2, 354, 352, 3, 2, 2, 2, 355, 356, 9, 5, 2, 2,
	356, 63, 3, 2, 2, 2, 357, 358, 9, 6, 2, 2, 358, 65, 3, 2, 2, 2, 359, 360,
	7, 51, 2, 2, 360, 67, 3, 2, 2, 2, 50, 73, 78, 82, 89, 97, 102, 106, 109,
	114, 117, 131, 133, 137, 140, 149, 152, 157, 162, 165, 176, 179, 184, 187,
	190, 193, 196, 199, 204, 209, 212, 229, 236, 239, 244, 248, 255, 265, 274,
	281, 285, 291, 293, 302, 309, 327, 334, 342, 352,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'businessobject'", "'to'", "'association'", "'element'", "'node'",
	"'action'", "'message'", "'raises'", "'using'", "'import'", "'as'", "'text'",
	"'valuation'", "", "", "'///'", "'['", "']'", "'('", "')'", "'{'", "'}'",
	"';'", "','", "'='", "'?'", "':'", "'.'", "'+'", "'-'", "'~'", "'!'", "'*'",
	"'/'", "'%'", "'>>'", "'<<'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='",
	"'&&'", "'||'", "'=>'", "'n'",
}
var symbolicNames = []string{
	"", "BooleanLiteral", "BUSINESSOBJECT", "TO", "ASSOCIATION", "ELEMENT",
	"NODE", "ACTION", "MESSAGE", "RAISES", "USING", "IMPORT", "AS", "TEXT",
	"VALUATION", "MultiLineComment", "SingleLineComment", "CustomAnnotationStart",
	"OpenBracket", "CloseBracket", "OpenParen", "CloseParen", "OpenBrace",
	"CloseBrace", "SemiColon", "Comma", "Assign", "QuestionMark", "Colon",
	"Dot", "Plus", "Minus", "BitNot", "Not", "Multiply", "Divide", "Modulus",
	"RightShiftArithmetic", "LeftShiftArithmetic", "LessThan", "MoreThan",
	"LessThanEquals", "GreaterThanEquals", "Equals_", "NotEquals", "And", "Or",
	"ARROW", "N", "Identifier", "StringLiteral", "DecimalLiteral", "HexIntegerLiteral",
	"OctalIntegerLiteral", "OctalIntegerLiteral2", "BinaryIntegerLiteral",
	"WhiteSpaces", "LineTerminator",
}

var ruleNames = []string{
	"program", "statements", "importStatement", "definitions", "definition",
	"block", "itemList", "element", "boAction", "message", "node", "association",
	"associationUsingDefinition", "valuationDefinition", "valutaionExpressionList",
	"valutaionExpression", "raiseMessage", "annotations", "annotation", "typeList",
	"typeDeclaration", "typeDefaultValue", "valueAssignList", "valueAssign",
	"multiplicity", "memberExpression", "identifierList", "keyword", "literal",
	"comments", "compareOperator", "logicOperator", "identifier",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type BODLParser struct {
	*antlr.BaseParser
}

func NewBODLParser(input antlr.TokenStream) *BODLParser {
	this := new(BODLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "BODLParser.g4"

	return this
}

// BODLParser tokens.
const (
	BODLParserEOF                   = antlr.TokenEOF
	BODLParserBooleanLiteral        = 1
	BODLParserBUSINESSOBJECT        = 2
	BODLParserTO                    = 3
	BODLParserASSOCIATION           = 4
	BODLParserELEMENT               = 5
	BODLParserNODE                  = 6
	BODLParserACTION                = 7
	BODLParserMESSAGE               = 8
	BODLParserRAISES                = 9
	BODLParserUSING                 = 10
	BODLParserIMPORT                = 11
	BODLParserAS                    = 12
	BODLParserTEXT                  = 13
	BODLParserVALUATION             = 14
	BODLParserMultiLineComment      = 15
	BODLParserSingleLineComment     = 16
	BODLParserCustomAnnotationStart = 17
	BODLParserOpenBracket           = 18
	BODLParserCloseBracket          = 19
	BODLParserOpenParen             = 20
	BODLParserCloseParen            = 21
	BODLParserOpenBrace             = 22
	BODLParserCloseBrace            = 23
	BODLParserSemiColon             = 24
	BODLParserComma                 = 25
	BODLParserAssign                = 26
	BODLParserQuestionMark          = 27
	BODLParserColon                 = 28
	BODLParserDot                   = 29
	BODLParserPlus                  = 30
	BODLParserMinus                 = 31
	BODLParserBitNot                = 32
	BODLParserNot                   = 33
	BODLParserMultiply              = 34
	BODLParserDivide                = 35
	BODLParserModulus               = 36
	BODLParserRightShiftArithmetic  = 37
	BODLParserLeftShiftArithmetic   = 38
	BODLParserLessThan              = 39
	BODLParserMoreThan              = 40
	BODLParserLessThanEquals        = 41
	BODLParserGreaterThanEquals     = 42
	BODLParserEquals_               = 43
	BODLParserNotEquals             = 44
	BODLParserAnd                   = 45
	BODLParserOr                    = 46
	BODLParserARROW                 = 47
	BODLParserN                     = 48
	BODLParserIdentifier            = 49
	BODLParserStringLiteral         = 50
	BODLParserDecimalLiteral        = 51
	BODLParserHexIntegerLiteral     = 52
	BODLParserOctalIntegerLiteral   = 53
	BODLParserOctalIntegerLiteral2  = 54
	BODLParserBinaryIntegerLiteral  = 55
	BODLParserWhiteSpaces           = 56
	BODLParserLineTerminator        = 57
)

// BODLParser rules.
const (
	BODLParserRULE_program                    = 0
	BODLParserRULE_statements                 = 1
	BODLParserRULE_importStatement            = 2
	BODLParserRULE_definitions                = 3
	BODLParserRULE_definition                 = 4
	BODLParserRULE_block                      = 5
	BODLParserRULE_itemList                   = 6
	BODLParserRULE_element                    = 7
	BODLParserRULE_boAction                   = 8
	BODLParserRULE_message                    = 9
	BODLParserRULE_node                       = 10
	BODLParserRULE_association                = 11
	BODLParserRULE_associationUsingDefinition = 12
	BODLParserRULE_valuationDefinition        = 13
	BODLParserRULE_valutaionExpressionList    = 14
	BODLParserRULE_valutaionExpression        = 15
	BODLParserRULE_raiseMessage               = 16
	BODLParserRULE_annotations                = 17
	BODLParserRULE_annotation                 = 18
	BODLParserRULE_typeList                   = 19
	BODLParserRULE_typeDeclaration            = 20
	BODLParserRULE_typeDefaultValue           = 21
	BODLParserRULE_valueAssignList            = 22
	BODLParserRULE_valueAssign                = 23
	BODLParserRULE_multiplicity               = 24
	BODLParserRULE_memberExpression           = 25
	BODLParserRULE_identifierList             = 26
	BODLParserRULE_keyword                    = 27
	BODLParserRULE_literal                    = 28
	BODLParserRULE_comments                   = 29
	BODLParserRULE_compareOperator            = 30
	BODLParserRULE_logicOperator              = 31
	BODLParserRULE_identifier                 = 32
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(BODLParserEOF, 0)
}

func (s *ProgramContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ProgramContext) Definitions() IDefinitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefinitionsContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BODLParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Match(BODLParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.Statements()
		}
		{
			p.SetState(68)
			p.Definitions()
		}
		{
			p.SetState(69)
			p.Match(BODLParserEOF)
		}

	}

	return localctx
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllImportStatement() []IImportStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportStatementContext)(nil)).Elem())
	var tst = make([]IImportStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportStatementContext)
		}
	}

	return tst
}

func (s *StatementsContext) ImportStatement(i int) IImportStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BODLParserRULE_statements)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(73)
				p.ImportStatement()
			}

		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_importStatement
	return p
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(BODLParserIMPORT, 0)
}

func (s *ImportStatementContext) MemberExpression() IMemberExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberExpressionContext)
}

func (s *ImportStatementContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(BODLParserSemiColon, 0)
}

func (s *ImportStatementContext) Comments() ICommentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentsContext)
}

func (s *ImportStatementContext) AS() antlr.TerminalNode {
	return s.GetToken(BODLParserAS, 0)
}

func (s *ImportStatementContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (s *ImportStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitImportStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BODLParserRULE_importStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(80)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(79)
				p.Comments()
			}

		}
		{
			p.SetState(82)
			p.Match(BODLParserIMPORT)
		}
		{
			p.SetState(83)
			p.MemberExpression()
		}
		{
			p.SetState(84)
			p.Match(BODLParserSemiColon)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(87)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(86)
				p.Comments()
			}

		}
		{
			p.SetState(89)
			p.Match(BODLParserIMPORT)
		}
		{
			p.SetState(90)
			p.MemberExpression()
		}
		{
			p.SetState(91)
			p.Match(BODLParserAS)
		}
		{
			p.SetState(92)
			p.Identifier()
		}
		{
			p.SetState(93)
			p.Match(BODLParserSemiColon)
		}

	}

	return localctx
}

// IDefinitionsContext is an interface to support dynamic dispatch.
type IDefinitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionsContext differentiates from other interfaces.
	IsDefinitionsContext()
}

type DefinitionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionsContext() *DefinitionsContext {
	var p = new(DefinitionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_definitions
	return p
}

func (*DefinitionsContext) IsDefinitionsContext() {}

func NewDefinitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionsContext {
	var p = new(DefinitionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_definitions

	return p
}

func (s *DefinitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionsContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *DefinitionsContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DefinitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterDefinitions(s)
	}
}

func (s *DefinitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitDefinitions(s)
	}
}

func (s *DefinitionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitDefinitions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) Definitions() (localctx IDefinitionsContext) {
	localctx = NewDefinitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BODLParserRULE_definitions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BODLParserBUSINESSOBJECT)|(1<<BODLParserMultiLineComment)|(1<<BODLParserSingleLineComment)|(1<<BODLParserCustomAnnotationStart)|(1<<BODLParserOpenBracket))) != 0 {
		{
			p.SetState(97)
			p.Definition()
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) BUSINESSOBJECT() antlr.TerminalNode {
	return s.GetToken(BODLParserBUSINESSOBJECT, 0)
}

func (s *DefinitionContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DefinitionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DefinitionContext) TypeDeclaration() ITypeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclarationContext)
}

func (s *DefinitionContext) Comments() ICommentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentsContext)
}

func (s *DefinitionContext) Annotations() IAnnotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationsContext)
}

func (s *DefinitionContext) RaiseMessage() IRaiseMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRaiseMessageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRaiseMessageContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (s *DefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BODLParserRULE_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(104)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(103)
			p.Comments()
		}

	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(106)
			p.Annotations()
		}

	}
	{
		p.SetState(109)
		p.Match(BODLParserBUSINESSOBJECT)
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(110)
			p.Identifier()
		}

	case 2:
		{
			p.SetState(111)
			p.TypeDeclaration()
		}

	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(114)
			p.RaiseMessage()
		}

	}
	{
		p.SetState(117)
		p.Block()
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(BODLParserOpenBrace, 0)
}

func (s *BlockContext) ItemList() IItemListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItemListContext)
}

func (s *BlockContext) Comments() ICommentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentsContext)
}

func (s *BlockContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(BODLParserCloseBrace, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BODLParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(BODLParserOpenBrace)
	}
	{
		p.SetState(120)
		p.ItemList()
	}
	{
		p.SetState(121)
		p.Comments()
	}
	{
		p.SetState(122)
		p.Match(BODLParserCloseBrace)
	}

	return localctx
}

// IItemListContext is an interface to support dynamic dispatch.
type IItemListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItemListContext differentiates from other interfaces.
	IsItemListContext()
}

type ItemListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItemListContext() *ItemListContext {
	var p = new(ItemListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_itemList
	return p
}

func (*ItemListContext) IsItemListContext() {}

func NewItemListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItemListContext {
	var p = new(ItemListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_itemList

	return p
}

func (s *ItemListContext) GetParser() antlr.Parser { return s.parser }

func (s *ItemListContext) AllElement() []IElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementContext)(nil)).Elem())
	var tst = make([]IElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementContext)
		}
	}

	return tst
}

func (s *ItemListContext) Element(i int) IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ItemListContext) AllMessage() []IMessageContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageContext)(nil)).Elem())
	var tst = make([]IMessageContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageContext)
		}
	}

	return tst
}

func (s *ItemListContext) Message(i int) IMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *ItemListContext) AllNode() []INodeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INodeContext)(nil)).Elem())
	var tst = make([]INodeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INodeContext)
		}
	}

	return tst
}

func (s *ItemListContext) Node(i int) INodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INodeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INodeContext)
}

func (s *ItemListContext) AllBoAction() []IBoActionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoActionContext)(nil)).Elem())
	var tst = make([]IBoActionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoActionContext)
		}
	}

	return tst
}

func (s *ItemListContext) BoAction(i int) IBoActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoActionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoActionContext)
}

func (s *ItemListContext) AllAssociation() []IAssociationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssociationContext)(nil)).Elem())
	var tst = make([]IAssociationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssociationContext)
		}
	}

	return tst
}

func (s *ItemListContext) Association(i int) IAssociationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssociationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssociationContext)
}

func (s *ItemListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItemListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterItemList(s)
	}
}

func (s *ItemListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitItemList(s)
	}
}

func (s *ItemListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitItemList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) ItemList() (localctx IItemListContext) {
	localctx = NewItemListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BODLParserRULE_itemList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(124)
					p.Element()
				}

			case 2:
				{
					p.SetState(125)
					p.Message()
				}

			case 3:
				{
					p.SetState(126)
					p.Node()
				}

			case 4:
				{
					p.SetState(127)
					p.BoAction()
				}

			case 5:
				{
					p.SetState(128)
					p.Association()
				}

			}

		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_element
	return p
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) ELEMENT() antlr.TerminalNode {
	return s.GetToken(BODLParserELEMENT, 0)
}

func (s *ElementContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ElementContext) Colon() antlr.TerminalNode {
	return s.GetToken(BODLParserColon, 0)
}

func (s *ElementContext) TypeDeclaration() ITypeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclarationContext)
}

func (s *ElementContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(BODLParserSemiColon, 0)
}

func (s *ElementContext) Comments() ICommentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentsContext)
}

func (s *ElementContext) Annotations() IAnnotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationsContext)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterElement(s)
	}
}

func (s *ElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitElement(s)
	}
}

func (s *ElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BODLParserRULE_element)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(135)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(134)
			p.Comments()
		}

	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(137)
			p.Annotations()
		}

	}
	{
		p.SetState(140)
		p.Match(BODLParserELEMENT)
	}
	{
		p.SetState(141)
		p.Identifier()
	}
	{
		p.SetState(142)
		p.Match(BODLParserColon)
	}
	{
		p.SetState(143)
		p.TypeDeclaration()
	}
	{
		p.SetState(144)
		p.Match(BODLParserSemiColon)
	}

	return localctx
}

// IBoActionContext is an interface to support dynamic dispatch.
type IBoActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoActionContext differentiates from other interfaces.
	IsBoActionContext()
}

type BoActionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoActionContext() *BoActionContext {
	var p = new(BoActionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_boAction
	return p
}

func (*BoActionContext) IsBoActionContext() {}

func NewBoActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoActionContext {
	var p = new(BoActionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_boAction

	return p
}

func (s *BoActionContext) GetParser() antlr.Parser { return s.parser }

func (s *BoActionContext) ACTION() antlr.TerminalNode {
	return s.GetToken(BODLParserACTION, 0)
}

func (s *BoActionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *BoActionContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(BODLParserSemiColon, 0)
}

func (s *BoActionContext) Comments() ICommentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentsContext)
}

func (s *BoActionContext) Annotations() IAnnotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationsContext)
}

func (s *BoActionContext) RaiseMessage() IRaiseMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRaiseMessageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRaiseMessageContext)
}

func (s *BoActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterBoAction(s)
	}
}

func (s *BoActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitBoAction(s)
	}
}

func (s *BoActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitBoAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) BoAction() (localctx IBoActionContext) {
	localctx = NewBoActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BODLParserRULE_boAction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(146)
			p.Comments()
		}

	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(149)
			p.Annotations()
		}

	}
	{
		p.SetState(152)
		p.Match(BODLParserACTION)
	}
	{
		p.SetState(153)
		p.Identifier()
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(154)
			p.RaiseMessage()
		}

	}
	{
		p.SetState(157)
		p.Match(BODLParserSemiColon)
	}

	return localctx
}

// IMessageContext is an interface to support dynamic dispatch.
type IMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageContext differentiates from other interfaces.
	IsMessageContext()
}

type MessageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageContext() *MessageContext {
	var p = new(MessageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_message
	return p
}

func (*MessageContext) IsMessageContext() {}

func NewMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageContext {
	var p = new(MessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_message

	return p
}

func (s *MessageContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(BODLParserMESSAGE, 0)
}

func (s *MessageContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MessageContext) TEXT() antlr.TerminalNode {
	return s.GetToken(BODLParserTEXT, 0)
}

func (s *MessageContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(BODLParserStringLiteral, 0)
}

func (s *MessageContext) Colon() antlr.TerminalNode {
	return s.GetToken(BODLParserColon, 0)
}

func (s *MessageContext) TypeList() ITypeListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *MessageContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(BODLParserSemiColon, 0)
}

func (s *MessageContext) Comments() ICommentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentsContext)
}

func (s *MessageContext) Annotations() IAnnotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationsContext)
}

func (s *MessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterMessage(s)
	}
}

func (s *MessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitMessage(s)
	}
}

func (s *MessageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitMessage(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) Message() (localctx IMessageContext) {
	localctx = NewMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BODLParserRULE_message)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(159)
			p.Comments()
		}

	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(162)
			p.Annotations()
		}

	}
	{
		p.SetState(165)
		p.Match(BODLParserMESSAGE)
	}
	{
		p.SetState(166)
		p.Identifier()
	}
	{
		p.SetState(167)
		p.Match(BODLParserTEXT)
	}
	{
		p.SetState(168)
		p.Match(BODLParserStringLiteral)
	}
	{
		p.SetState(169)
		p.Match(BODLParserColon)
	}
	{
		p.SetState(170)
		p.TypeList()
	}
	{
		p.SetState(171)
		p.Match(BODLParserSemiColon)
	}

	return localctx
}

// INodeContext is an interface to support dynamic dispatch.
type INodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNodeContext differentiates from other interfaces.
	IsNodeContext()
}

type NodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNodeContext() *NodeContext {
	var p = new(NodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_node
	return p
}

func (*NodeContext) IsNodeContext() {}

func NewNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NodeContext {
	var p = new(NodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_node

	return p
}

func (s *NodeContext) GetParser() antlr.Parser { return s.parser }

func (s *NodeContext) NODE() antlr.TerminalNode {
	return s.GetToken(BODLParserNODE, 0)
}

func (s *NodeContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *NodeContext) Comments() ICommentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentsContext)
}

func (s *NodeContext) Annotations() IAnnotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationsContext)
}

func (s *NodeContext) Multiplicity() IMultiplicityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicityContext)
}

func (s *NodeContext) RaiseMessage() IRaiseMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRaiseMessageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRaiseMessageContext)
}

func (s *NodeContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *NodeContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(BODLParserSemiColon, 0)
}

func (s *NodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterNode(s)
	}
}

func (s *NodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitNode(s)
	}
}

func (s *NodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitNode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) Node() (localctx INodeContext) {
	localctx = NewNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BODLParserRULE_node)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(174)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(173)
			p.Comments()
		}

	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(176)
			p.Annotations()
		}

	}
	{
		p.SetState(179)
		p.Match(BODLParserNODE)
	}
	{
		p.SetState(180)
		p.Identifier()
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(181)
			p.Multiplicity()
		}

	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(184)
			p.RaiseMessage()
		}

	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserOpenBrace {
		{
			p.SetState(187)
			p.Block()
		}

	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserSemiColon {
		{
			p.SetState(190)
			p.Match(BODLParserSemiColon)
		}

	}

	return localctx
}

// IAssociationContext is an interface to support dynamic dispatch.
type IAssociationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssociationContext differentiates from other interfaces.
	IsAssociationContext()
}

type AssociationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssociationContext() *AssociationContext {
	var p = new(AssociationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_association
	return p
}

func (*AssociationContext) IsAssociationContext() {}

func NewAssociationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssociationContext {
	var p = new(AssociationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_association

	return p
}

func (s *AssociationContext) GetParser() antlr.Parser { return s.parser }

func (s *AssociationContext) ASSOCIATION() antlr.TerminalNode {
	return s.GetToken(BODLParserASSOCIATION, 0)
}

func (s *AssociationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AssociationContext) TO() antlr.TerminalNode {
	return s.GetToken(BODLParserTO, 0)
}

func (s *AssociationContext) MemberExpression() IMemberExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberExpressionContext)
}

func (s *AssociationContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(BODLParserSemiColon, 0)
}

func (s *AssociationContext) Comments() ICommentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentsContext)
}

func (s *AssociationContext) Annotations() IAnnotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationsContext)
}

func (s *AssociationContext) Multiplicity() IMultiplicityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicityContext)
}

func (s *AssociationContext) AssociationUsingDefinition() IAssociationUsingDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssociationUsingDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssociationUsingDefinitionContext)
}

func (s *AssociationContext) ValuationDefinition() IValuationDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuationDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuationDefinitionContext)
}

func (s *AssociationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssociationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssociationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterAssociation(s)
	}
}

func (s *AssociationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitAssociation(s)
	}
}

func (s *AssociationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitAssociation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) Association() (localctx IAssociationContext) {
	localctx = NewAssociationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BODLParserRULE_association)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(194)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(193)
			p.Comments()
		}

	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(196)
			p.Annotations()
		}

	}
	{
		p.SetState(199)
		p.Match(BODLParserASSOCIATION)
	}
	{
		p.SetState(200)
		p.Identifier()
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserOpenBracket {
		{
			p.SetState(201)
			p.Multiplicity()
		}

	}
	{
		p.SetState(204)
		p.Match(BODLParserTO)
	}
	{
		p.SetState(205)
		p.MemberExpression()
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserUSING {
		{
			p.SetState(206)
			p.AssociationUsingDefinition()
		}

	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserVALUATION {
		{
			p.SetState(209)
			p.ValuationDefinition()
		}

	}
	{
		p.SetState(212)
		p.Match(BODLParserSemiColon)
	}

	return localctx
}

// IAssociationUsingDefinitionContext is an interface to support dynamic dispatch.
type IAssociationUsingDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssociationUsingDefinitionContext differentiates from other interfaces.
	IsAssociationUsingDefinitionContext()
}

type AssociationUsingDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssociationUsingDefinitionContext() *AssociationUsingDefinitionContext {
	var p = new(AssociationUsingDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_associationUsingDefinition
	return p
}

func (*AssociationUsingDefinitionContext) IsAssociationUsingDefinitionContext() {}

func NewAssociationUsingDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssociationUsingDefinitionContext {
	var p = new(AssociationUsingDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_associationUsingDefinition

	return p
}

func (s *AssociationUsingDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssociationUsingDefinitionContext) USING() antlr.TerminalNode {
	return s.GetToken(BODLParserUSING, 0)
}

func (s *AssociationUsingDefinitionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AssociationUsingDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssociationUsingDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssociationUsingDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterAssociationUsingDefinition(s)
	}
}

func (s *AssociationUsingDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitAssociationUsingDefinition(s)
	}
}

func (s *AssociationUsingDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitAssociationUsingDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) AssociationUsingDefinition() (localctx IAssociationUsingDefinitionContext) {
	localctx = NewAssociationUsingDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BODLParserRULE_associationUsingDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(BODLParserUSING)
	}
	{
		p.SetState(215)
		p.Identifier()
	}

	return localctx
}

// IValuationDefinitionContext is an interface to support dynamic dispatch.
type IValuationDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValuationDefinitionContext differentiates from other interfaces.
	IsValuationDefinitionContext()
}

type ValuationDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuationDefinitionContext() *ValuationDefinitionContext {
	var p = new(ValuationDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_valuationDefinition
	return p
}

func (*ValuationDefinitionContext) IsValuationDefinitionContext() {}

func NewValuationDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuationDefinitionContext {
	var p = new(ValuationDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_valuationDefinition

	return p
}

func (s *ValuationDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuationDefinitionContext) VALUATION() antlr.TerminalNode {
	return s.GetToken(BODLParserVALUATION, 0)
}

func (s *ValuationDefinitionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(BODLParserOpenParen, 0)
}

func (s *ValuationDefinitionContext) ValutaionExpressionList() IValutaionExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValutaionExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValutaionExpressionListContext)
}

func (s *ValuationDefinitionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(BODLParserCloseParen, 0)
}

func (s *ValuationDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuationDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuationDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterValuationDefinition(s)
	}
}

func (s *ValuationDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitValuationDefinition(s)
	}
}

func (s *ValuationDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitValuationDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) ValuationDefinition() (localctx IValuationDefinitionContext) {
	localctx = NewValuationDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BODLParserRULE_valuationDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.Match(BODLParserVALUATION)
	}
	{
		p.SetState(218)
		p.Match(BODLParserOpenParen)
	}
	{
		p.SetState(219)
		p.ValutaionExpressionList()
	}
	{
		p.SetState(220)
		p.Match(BODLParserCloseParen)
	}

	return localctx
}

// IValutaionExpressionListContext is an interface to support dynamic dispatch.
type IValutaionExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValutaionExpressionListContext differentiates from other interfaces.
	IsValutaionExpressionListContext()
}

type ValutaionExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValutaionExpressionListContext() *ValutaionExpressionListContext {
	var p = new(ValutaionExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_valutaionExpressionList
	return p
}

func (*ValutaionExpressionListContext) IsValutaionExpressionListContext() {}

func NewValutaionExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValutaionExpressionListContext {
	var p = new(ValutaionExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_valutaionExpressionList

	return p
}

func (s *ValutaionExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ValutaionExpressionListContext) ValutaionExpression() IValutaionExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValutaionExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValutaionExpressionContext)
}

func (s *ValutaionExpressionListContext) LogicOperator() ILogicOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicOperatorContext)
}

func (s *ValutaionExpressionListContext) ValutaionExpressionList() IValutaionExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValutaionExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValutaionExpressionListContext)
}

func (s *ValutaionExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValutaionExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValutaionExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterValutaionExpressionList(s)
	}
}

func (s *ValutaionExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitValutaionExpressionList(s)
	}
}

func (s *ValutaionExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitValutaionExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) ValutaionExpressionList() (localctx IValutaionExpressionListContext) {
	localctx = NewValutaionExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BODLParserRULE_valutaionExpressionList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)
			p.ValutaionExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(223)
			p.ValutaionExpression()
		}
		{
			p.SetState(224)
			p.LogicOperator()
		}
		{
			p.SetState(225)
			p.ValutaionExpressionList()
		}

	}

	return localctx
}

// IValutaionExpressionContext is an interface to support dynamic dispatch.
type IValutaionExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValutaionExpressionContext differentiates from other interfaces.
	IsValutaionExpressionContext()
}

type ValutaionExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValutaionExpressionContext() *ValutaionExpressionContext {
	var p = new(ValutaionExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_valutaionExpression
	return p
}

func (*ValutaionExpressionContext) IsValutaionExpressionContext() {}

func NewValutaionExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValutaionExpressionContext {
	var p = new(ValutaionExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_valutaionExpression

	return p
}

func (s *ValutaionExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ValutaionExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ValutaionExpressionContext) CompareOperator() ICompareOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareOperatorContext)
}

func (s *ValutaionExpressionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ValutaionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValutaionExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValutaionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterValutaionExpression(s)
	}
}

func (s *ValutaionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitValutaionExpression(s)
	}
}

func (s *ValutaionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitValutaionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) ValutaionExpression() (localctx IValutaionExpressionContext) {
	localctx = NewValutaionExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BODLParserRULE_valutaionExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Identifier()
	}
	{
		p.SetState(230)
		p.CompareOperator()
	}
	{
		p.SetState(231)
		p.Literal()
	}

	return localctx
}

// IRaiseMessageContext is an interface to support dynamic dispatch.
type IRaiseMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRaiseMessageContext differentiates from other interfaces.
	IsRaiseMessageContext()
}

type RaiseMessageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRaiseMessageContext() *RaiseMessageContext {
	var p = new(RaiseMessageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_raiseMessage
	return p
}

func (*RaiseMessageContext) IsRaiseMessageContext() {}

func NewRaiseMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RaiseMessageContext {
	var p = new(RaiseMessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_raiseMessage

	return p
}

func (s *RaiseMessageContext) GetParser() antlr.Parser { return s.parser }

func (s *RaiseMessageContext) RAISES() antlr.TerminalNode {
	return s.GetToken(BODLParserRAISES, 0)
}

func (s *RaiseMessageContext) IdentifierList() IIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *RaiseMessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RaiseMessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RaiseMessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterRaiseMessage(s)
	}
}

func (s *RaiseMessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitRaiseMessage(s)
	}
}

func (s *RaiseMessageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitRaiseMessage(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) RaiseMessage() (localctx IRaiseMessageContext) {
	localctx = NewRaiseMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BODLParserRULE_raiseMessage)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserRAISES {
		{
			p.SetState(233)
			p.Match(BODLParserRAISES)
		}

	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserIdentifier {
		{
			p.SetState(236)
			p.IdentifierList()
		}

	}

	return localctx
}

// IAnnotationsContext is an interface to support dynamic dispatch.
type IAnnotationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationsContext differentiates from other interfaces.
	IsAnnotationsContext()
}

type AnnotationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationsContext() *AnnotationsContext {
	var p = new(AnnotationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_annotations
	return p
}

func (*AnnotationsContext) IsAnnotationsContext() {}

func NewAnnotationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationsContext {
	var p = new(AnnotationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_annotations

	return p
}

func (s *AnnotationsContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationsContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *AnnotationsContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *AnnotationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterAnnotations(s)
	}
}

func (s *AnnotationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitAnnotations(s)
	}
}

func (s *AnnotationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitAnnotations(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) Annotations() (localctx IAnnotationsContext) {
	localctx = NewAnnotationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BODLParserRULE_annotations)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BODLParserCustomAnnotationStart || _la == BODLParserOpenBracket {
		{
			p.SetState(239)
			p.Annotation()
		}

		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_annotation
	return p
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(BODLParserOpenBracket, 0)
}

func (s *AnnotationContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *AnnotationContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AnnotationContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(BODLParserCloseBracket, 0)
}

func (s *AnnotationContext) CustomAnnotationStart() antlr.TerminalNode {
	return s.GetToken(BODLParserCustomAnnotationStart, 0)
}

func (s *AnnotationContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(BODLParserOpenParen, 0)
}

func (s *AnnotationContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(BODLParserCloseParen, 0)
}

func (s *AnnotationContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitAnnotation(s)
	}
}

func (s *AnnotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitAnnotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BODLParserRULE_annotation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BODLParserCustomAnnotationStart {
			{
				p.SetState(245)
				p.Match(BODLParserCustomAnnotationStart)
			}

		}
		{
			p.SetState(248)
			p.Match(BODLParserOpenBracket)
		}
		{
			p.SetState(249)
			p.Identifier()
		}
		{
			p.SetState(250)
			p.Match(BODLParserCloseBracket)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BODLParserCustomAnnotationStart {
			{
				p.SetState(252)
				p.Match(BODLParserCustomAnnotationStart)
			}

		}
		{
			p.SetState(255)
			p.Match(BODLParserOpenBracket)
		}
		{
			p.SetState(256)
			p.Identifier()
		}
		{
			p.SetState(257)
			p.Match(BODLParserOpenParen)
		}
		{
			p.SetState(258)
			p.Identifier()
		}
		{
			p.SetState(259)
			p.Match(BODLParserCloseParen)
		}
		{
			p.SetState(260)
			p.Match(BODLParserCloseBracket)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BODLParserCustomAnnotationStart {
			{
				p.SetState(262)
				p.Match(BODLParserCustomAnnotationStart)
			}

		}
		{
			p.SetState(265)
			p.Match(BODLParserOpenBracket)
		}
		{
			p.SetState(266)
			p.Identifier()
		}
		{
			p.SetState(267)
			p.Match(BODLParserOpenParen)
		}
		{
			p.SetState(268)
			p.Literal()
		}
		{
			p.SetState(269)
			p.Match(BODLParserCloseParen)
		}
		{
			p.SetState(270)
			p.Match(BODLParserCloseBracket)
		}

	}

	return localctx
}

// ITypeListContext is an interface to support dynamic dispatch.
type ITypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeListContext differentiates from other interfaces.
	IsTypeListContext()
}

type TypeListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeListContext() *TypeListContext {
	var p = new(TypeListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_typeList
	return p
}

func (*TypeListContext) IsTypeListContext() {}

func NewTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeListContext {
	var p = new(TypeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_typeList

	return p
}

func (s *TypeListContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeListContext) TypeDeclaration() ITypeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclarationContext)
}

func (s *TypeListContext) Comma() antlr.TerminalNode {
	return s.GetToken(BODLParserComma, 0)
}

func (s *TypeListContext) TypeList() ITypeListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *TypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterTypeList(s)
	}
}

func (s *TypeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitTypeList(s)
	}
}

func (s *TypeListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitTypeList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) TypeList() (localctx ITypeListContext) {
	localctx = NewTypeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BODLParserRULE_typeList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.TypeDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(275)
			p.TypeDeclaration()
		}
		{
			p.SetState(276)
			p.Match(BODLParserComma)
		}
		{
			p.SetState(277)
			p.TypeList()
		}

	}

	return localctx
}

// ITypeDeclarationContext is an interface to support dynamic dispatch.
type ITypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclarationContext differentiates from other interfaces.
	IsTypeDeclarationContext()
}

type TypeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclarationContext() *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_typeDeclaration
	return p
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *TypeDeclarationContext) TypeDefaultValue() ITypeDefaultValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDefaultValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDefaultValueContext)
}

func (s *TypeDeclarationContext) MemberExpression() IMemberExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberExpressionContext)
}

func (s *TypeDeclarationContext) Colon() antlr.TerminalNode {
	return s.GetToken(BODLParserColon, 0)
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitTypeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BODLParserRULE_typeDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
			p.Identifier()
		}
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BODLParserAssign {
			{
				p.SetState(282)
				p.TypeDefaultValue()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(285)
			p.MemberExpression()
		}
		{
			p.SetState(286)
			p.Match(BODLParserColon)
		}
		{
			p.SetState(287)
			p.Identifier()
		}
		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BODLParserAssign {
			{
				p.SetState(288)
				p.TypeDefaultValue()
			}

		}

	}

	return localctx
}

// ITypeDefaultValueContext is an interface to support dynamic dispatch.
type ITypeDefaultValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDefaultValueContext differentiates from other interfaces.
	IsTypeDefaultValueContext()
}

type TypeDefaultValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefaultValueContext() *TypeDefaultValueContext {
	var p = new(TypeDefaultValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_typeDefaultValue
	return p
}

func (*TypeDefaultValueContext) IsTypeDefaultValueContext() {}

func NewTypeDefaultValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefaultValueContext {
	var p = new(TypeDefaultValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_typeDefaultValue

	return p
}

func (s *TypeDefaultValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefaultValueContext) Assign() antlr.TerminalNode {
	return s.GetToken(BODLParserAssign, 0)
}

func (s *TypeDefaultValueContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *TypeDefaultValueContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(BODLParserOpenBrace, 0)
}

func (s *TypeDefaultValueContext) ValueAssignList() IValueAssignListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueAssignListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueAssignListContext)
}

func (s *TypeDefaultValueContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(BODLParserCloseBrace, 0)
}

func (s *TypeDefaultValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefaultValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefaultValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterTypeDefaultValue(s)
	}
}

func (s *TypeDefaultValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitTypeDefaultValue(s)
	}
}

func (s *TypeDefaultValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitTypeDefaultValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) TypeDefaultValue() (localctx ITypeDefaultValueContext) {
	localctx = NewTypeDefaultValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BODLParserRULE_typeDefaultValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.Match(BODLParserAssign)
		}
		{
			p.SetState(294)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(295)
			p.Match(BODLParserAssign)
		}
		{
			p.SetState(296)
			p.Match(BODLParserOpenBrace)
		}
		{
			p.SetState(297)
			p.ValueAssignList()
		}
		{
			p.SetState(298)
			p.Match(BODLParserCloseBrace)
		}

	}

	return localctx
}

// IValueAssignListContext is an interface to support dynamic dispatch.
type IValueAssignListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueAssignListContext differentiates from other interfaces.
	IsValueAssignListContext()
}

type ValueAssignListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueAssignListContext() *ValueAssignListContext {
	var p = new(ValueAssignListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_valueAssignList
	return p
}

func (*ValueAssignListContext) IsValueAssignListContext() {}

func NewValueAssignListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueAssignListContext {
	var p = new(ValueAssignListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_valueAssignList

	return p
}

func (s *ValueAssignListContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueAssignListContext) ValueAssign() IValueAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueAssignContext)
}

func (s *ValueAssignListContext) Comma() antlr.TerminalNode {
	return s.GetToken(BODLParserComma, 0)
}

func (s *ValueAssignListContext) ValueAssignList() IValueAssignListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueAssignListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueAssignListContext)
}

func (s *ValueAssignListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueAssignListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueAssignListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterValueAssignList(s)
	}
}

func (s *ValueAssignListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitValueAssignList(s)
	}
}

func (s *ValueAssignListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitValueAssignList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) ValueAssignList() (localctx IValueAssignListContext) {
	localctx = NewValueAssignListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BODLParserRULE_valueAssignList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(302)
			p.ValueAssign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(303)
			p.ValueAssign()
		}
		{
			p.SetState(304)
			p.Match(BODLParserComma)
		}
		{
			p.SetState(305)
			p.ValueAssignList()
		}

	}

	return localctx
}

// IValueAssignContext is an interface to support dynamic dispatch.
type IValueAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueAssignContext differentiates from other interfaces.
	IsValueAssignContext()
}

type ValueAssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueAssignContext() *ValueAssignContext {
	var p = new(ValueAssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_valueAssign
	return p
}

func (*ValueAssignContext) IsValueAssignContext() {}

func NewValueAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueAssignContext {
	var p = new(ValueAssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_valueAssign

	return p
}

func (s *ValueAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueAssignContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ValueAssignContext) Assign() antlr.TerminalNode {
	return s.GetToken(BODLParserAssign, 0)
}

func (s *ValueAssignContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ValueAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterValueAssign(s)
	}
}

func (s *ValueAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitValueAssign(s)
	}
}

func (s *ValueAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitValueAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) ValueAssign() (localctx IValueAssignContext) {
	localctx = NewValueAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BODLParserRULE_valueAssign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.Identifier()
	}
	{
		p.SetState(310)
		p.Match(BODLParserAssign)
	}
	{
		p.SetState(311)
		p.Literal()
	}

	return localctx
}

// IMultiplicityContext is an interface to support dynamic dispatch.
type IMultiplicityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicityContext differentiates from other interfaces.
	IsMultiplicityContext()
}

type MultiplicityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicityContext() *MultiplicityContext {
	var p = new(MultiplicityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_multiplicity
	return p
}

func (*MultiplicityContext) IsMultiplicityContext() {}

func NewMultiplicityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicityContext {
	var p = new(MultiplicityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_multiplicity

	return p
}

func (s *MultiplicityContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicityContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(BODLParserOpenBracket, 0)
}

func (s *MultiplicityContext) AllLiteral() []ILiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILiteralContext)(nil)).Elem())
	var tst = make([]ILiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILiteralContext)
		}
	}

	return tst
}

func (s *MultiplicityContext) Literal(i int) ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *MultiplicityContext) Comma() antlr.TerminalNode {
	return s.GetToken(BODLParserComma, 0)
}

func (s *MultiplicityContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(BODLParserCloseBracket, 0)
}

func (s *MultiplicityContext) N() antlr.TerminalNode {
	return s.GetToken(BODLParserN, 0)
}

func (s *MultiplicityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterMultiplicity(s)
	}
}

func (s *MultiplicityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitMultiplicity(s)
	}
}

func (s *MultiplicityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitMultiplicity(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) Multiplicity() (localctx IMultiplicityContext) {
	localctx = NewMultiplicityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BODLParserRULE_multiplicity)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(313)
			p.Match(BODLParserOpenBracket)
		}
		{
			p.SetState(314)
			p.Literal()
		}
		{
			p.SetState(315)
			p.Match(BODLParserComma)
		}
		{
			p.SetState(316)
			p.Literal()
		}
		{
			p.SetState(317)
			p.Match(BODLParserCloseBracket)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(319)
			p.Match(BODLParserOpenBracket)
		}
		{
			p.SetState(320)
			p.Literal()
		}
		{
			p.SetState(321)
			p.Match(BODLParserComma)
		}
		{
			p.SetState(322)
			p.Match(BODLParserN)
		}
		{
			p.SetState(323)
			p.Match(BODLParserCloseBracket)
		}

	}

	return localctx
}

// IMemberExpressionContext is an interface to support dynamic dispatch.
type IMemberExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberExpressionContext differentiates from other interfaces.
	IsMemberExpressionContext()
}

type MemberExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberExpressionContext() *MemberExpressionContext {
	var p = new(MemberExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_memberExpression
	return p
}

func (*MemberExpressionContext) IsMemberExpressionContext() {}

func NewMemberExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberExpressionContext {
	var p = new(MemberExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_memberExpression

	return p
}

func (s *MemberExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MemberExpressionContext) AllDot() []antlr.TerminalNode {
	return s.GetTokens(BODLParserDot)
}

func (s *MemberExpressionContext) Dot(i int) antlr.TerminalNode {
	return s.GetToken(BODLParserDot, i)
}

func (s *MemberExpressionContext) AllMemberExpression() []IMemberExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMemberExpressionContext)(nil)).Elem())
	var tst = make([]IMemberExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMemberExpressionContext)
		}
	}

	return tst
}

func (s *MemberExpressionContext) MemberExpression(i int) IMemberExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMemberExpressionContext)
}

func (s *MemberExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterMemberExpression(s)
	}
}

func (s *MemberExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitMemberExpression(s)
	}
}

func (s *MemberExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitMemberExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) MemberExpression() (localctx IMemberExpressionContext) {
	localctx = NewMemberExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, BODLParserRULE_memberExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.Identifier()
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(328)
				p.Match(BODLParserDot)
			}
			{
				p.SetState(329)
				p.MemberExpression()
			}

		}
		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())
	}

	return localctx
}

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_identifierList
	return p
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *IdentifierListContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *IdentifierListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(BODLParserComma)
}

func (s *IdentifierListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(BODLParserComma, i)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterIdentifierList(s)
	}
}

func (s *IdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitIdentifierList(s)
	}
}

func (s *IdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BODLParserRULE_identifierList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.Identifier()
	}
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BODLParserComma {
		{
			p.SetState(336)
			p.Match(BODLParserComma)
		}
		{
			p.SetState(337)
			p.Identifier()
		}

		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_keyword
	return p
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(BODLParserIMPORT, 0)
}

func (s *KeywordContext) BUSINESSOBJECT() antlr.TerminalNode {
	return s.GetToken(BODLParserBUSINESSOBJECT, 0)
}

func (s *KeywordContext) TO() antlr.TerminalNode {
	return s.GetToken(BODLParserTO, 0)
}

func (s *KeywordContext) ASSOCIATION() antlr.TerminalNode {
	return s.GetToken(BODLParserASSOCIATION, 0)
}

func (s *KeywordContext) ELEMENT() antlr.TerminalNode {
	return s.GetToken(BODLParserELEMENT, 0)
}

func (s *KeywordContext) NODE() antlr.TerminalNode {
	return s.GetToken(BODLParserNODE, 0)
}

func (s *KeywordContext) ACTION() antlr.TerminalNode {
	return s.GetToken(BODLParserACTION, 0)
}

func (s *KeywordContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(BODLParserMESSAGE, 0)
}

func (s *KeywordContext) RAISES() antlr.TerminalNode {
	return s.GetToken(BODLParserRAISES, 0)
}

func (s *KeywordContext) USING() antlr.TerminalNode {
	return s.GetToken(BODLParserUSING, 0)
}

func (s *KeywordContext) TEXT() antlr.TerminalNode {
	return s.GetToken(BODLParserTEXT, 0)
}

func (s *KeywordContext) AS() antlr.TerminalNode {
	return s.GetToken(BODLParserAS, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (s *KeywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitKeyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BODLParserRULE_keyword)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(343)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BODLParserBUSINESSOBJECT)|(1<<BODLParserTO)|(1<<BODLParserASSOCIATION)|(1<<BODLParserELEMENT)|(1<<BODLParserNODE)|(1<<BODLParserACTION)|(1<<BODLParserMESSAGE)|(1<<BODLParserRAISES)|(1<<BODLParserUSING)|(1<<BODLParserIMPORT)|(1<<BODLParserAS)|(1<<BODLParserTEXT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(BODLParserDecimalLiteral, 0)
}

func (s *LiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(BODLParserBooleanLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(BODLParserStringLiteral, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BODLParserRULE_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BODLParserBooleanLiteral || _la == BODLParserStringLiteral || _la == BODLParserDecimalLiteral) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICommentsContext is an interface to support dynamic dispatch.
type ICommentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentsContext differentiates from other interfaces.
	IsCommentsContext()
}

type CommentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentsContext() *CommentsContext {
	var p = new(CommentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_comments
	return p
}

func (*CommentsContext) IsCommentsContext() {}

func NewCommentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentsContext {
	var p = new(CommentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_comments

	return p
}

func (s *CommentsContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentsContext) AllSingleLineComment() []antlr.TerminalNode {
	return s.GetTokens(BODLParserSingleLineComment)
}

func (s *CommentsContext) SingleLineComment(i int) antlr.TerminalNode {
	return s.GetToken(BODLParserSingleLineComment, i)
}

func (s *CommentsContext) AllMultiLineComment() []antlr.TerminalNode {
	return s.GetTokens(BODLParserMultiLineComment)
}

func (s *CommentsContext) MultiLineComment(i int) antlr.TerminalNode {
	return s.GetToken(BODLParserMultiLineComment, i)
}

func (s *CommentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterComments(s)
	}
}

func (s *CommentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitComments(s)
	}
}

func (s *CommentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitComments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) Comments() (localctx ICommentsContext) {
	localctx = NewCommentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BODLParserRULE_comments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BODLParserMultiLineComment || _la == BODLParserSingleLineComment {
		{
			p.SetState(347)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BODLParserMultiLineComment || _la == BODLParserSingleLineComment) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICompareOperatorContext is an interface to support dynamic dispatch.
type ICompareOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompareOperatorContext differentiates from other interfaces.
	IsCompareOperatorContext()
}

type CompareOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareOperatorContext() *CompareOperatorContext {
	var p = new(CompareOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_compareOperator
	return p
}

func (*CompareOperatorContext) IsCompareOperatorContext() {}

func NewCompareOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareOperatorContext {
	var p = new(CompareOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_compareOperator

	return p
}

func (s *CompareOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareOperatorContext) Equals_() antlr.TerminalNode {
	return s.GetToken(BODLParserEquals_, 0)
}

func (s *CompareOperatorContext) NotEquals() antlr.TerminalNode {
	return s.GetToken(BODLParserNotEquals, 0)
}

func (s *CompareOperatorContext) MoreThan() antlr.TerminalNode {
	return s.GetToken(BODLParserMoreThan, 0)
}

func (s *CompareOperatorContext) GreaterThanEquals() antlr.TerminalNode {
	return s.GetToken(BODLParserGreaterThanEquals, 0)
}

func (s *CompareOperatorContext) LessThan() antlr.TerminalNode {
	return s.GetToken(BODLParserLessThan, 0)
}

func (s *CompareOperatorContext) LessThanEquals() antlr.TerminalNode {
	return s.GetToken(BODLParserLessThanEquals, 0)
}

func (s *CompareOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterCompareOperator(s)
	}
}

func (s *CompareOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitCompareOperator(s)
	}
}

func (s *CompareOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitCompareOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) CompareOperator() (localctx ICompareOperatorContext) {
	localctx = NewCompareOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, BODLParserRULE_compareOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(BODLParserLessThan-39))|(1<<(BODLParserMoreThan-39))|(1<<(BODLParserLessThanEquals-39))|(1<<(BODLParserGreaterThanEquals-39))|(1<<(BODLParserEquals_-39))|(1<<(BODLParserNotEquals-39)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILogicOperatorContext is an interface to support dynamic dispatch.
type ILogicOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicOperatorContext differentiates from other interfaces.
	IsLogicOperatorContext()
}

type LogicOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicOperatorContext() *LogicOperatorContext {
	var p = new(LogicOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_logicOperator
	return p
}

func (*LogicOperatorContext) IsLogicOperatorContext() {}

func NewLogicOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicOperatorContext {
	var p = new(LogicOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_logicOperator

	return p
}

func (s *LogicOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicOperatorContext) And() antlr.TerminalNode {
	return s.GetToken(BODLParserAnd, 0)
}

func (s *LogicOperatorContext) Or() antlr.TerminalNode {
	return s.GetToken(BODLParserOr, 0)
}

func (s *LogicOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterLogicOperator(s)
	}
}

func (s *LogicOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitLogicOperator(s)
	}
}

func (s *LogicOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitLogicOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) LogicOperator() (localctx ILogicOperatorContext) {
	localctx = NewLogicOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, BODLParserRULE_logicOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BODLParserAnd || _la == BODLParserOr) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BODLParserIdentifier, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BODLParserVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BODLParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, BODLParserRULE_identifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(BODLParserIdentifier)
	}

	return localctx
}
