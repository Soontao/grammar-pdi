// Code generated from ../BODLParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package bodl // BODLParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 79, 336,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 3, 2, 3, 2, 5, 2,
	67, 10, 2, 3, 2, 7, 2, 70, 10, 2, 12, 2, 14, 2, 73, 11, 2, 3, 2, 7, 2,
	76, 10, 2, 12, 2, 14, 2, 79, 11, 2, 3, 2, 5, 2, 82, 10, 2, 3, 3, 5, 3,
	85, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 92, 10, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 5, 3, 100, 10, 3, 3, 4, 5, 4, 103, 10, 4, 3, 4, 3,
	4, 3, 4, 5, 4, 108, 10, 4, 3, 4, 5, 4, 111, 10, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 124, 10, 6, 12, 6, 14,
	6, 127, 11, 6, 3, 7, 5, 7, 130, 10, 7, 3, 7, 5, 7, 133, 10, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 5, 8, 142, 10, 8, 3, 8, 3, 8, 3, 8, 5,
	8, 147, 10, 8, 3, 8, 3, 8, 3, 9, 5, 9, 152, 10, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 5, 10, 163, 10, 10, 3, 10, 5, 10, 166,
	10, 10, 3, 10, 3, 10, 3, 10, 5, 10, 171, 10, 10, 3, 10, 5, 10, 174, 10,
	10, 3, 10, 3, 10, 3, 11, 5, 11, 179, 10, 11, 3, 11, 5, 11, 182, 10, 11,
	3, 11, 3, 11, 3, 11, 5, 11, 187, 10, 11, 3, 11, 3, 11, 3, 11, 5, 11, 192,
	10, 11, 3, 11, 5, 11, 195, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14,
	212, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 5, 16, 219, 10, 16, 3,
	16, 5, 16, 222, 10, 16, 3, 17, 7, 17, 225, 10, 17, 12, 17, 14, 17, 228,
	11, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 246, 10, 18, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 253, 10, 19, 3, 20, 3, 20, 5, 20,
	257, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 263, 10, 20, 5, 20, 265,
	10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 274, 10,
	21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 281, 10, 22, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 5, 24, 299, 10, 24, 3, 25, 3, 25, 3, 25, 3, 25,
	5, 25, 305, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 311, 10, 26, 3,
	27, 3, 27, 3, 28, 3, 28, 3, 29, 5, 29, 318, 10, 29, 3, 29, 7, 29, 321,
	10, 29, 12, 29, 14, 29, 324, 11, 29, 3, 30, 7, 30, 327, 10, 30, 12, 30,
	14, 30, 330, 11, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 2, 2, 33, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 2, 7, 3, 2, 3, 14, 3, 2, 70, 72,
	3, 2, 16, 17, 3, 2, 43, 48, 3, 2, 54, 55, 2, 351, 2, 81, 3, 2, 2, 2, 4,
	99, 3, 2, 2, 2, 6, 102, 3, 2, 2, 2, 8, 114, 3, 2, 2, 2, 10, 125, 3, 2,
	2, 2, 12, 129, 3, 2, 2, 2, 14, 141, 3, 2, 2, 2, 16, 151, 3, 2, 2, 2, 18,
	162, 3, 2, 2, 2, 20, 178, 3, 2, 2, 2, 22, 198, 3, 2, 2, 2, 24, 201, 3,
	2, 2, 2, 26, 211, 3, 2, 2, 2, 28, 213, 3, 2, 2, 2, 30, 218, 3, 2, 2, 2,
	32, 226, 3, 2, 2, 2, 34, 245, 3, 2, 2, 2, 36, 252, 3, 2, 2, 2, 38, 264,
	3, 2, 2, 2, 40, 273, 3, 2, 2, 2, 42, 280, 3, 2, 2, 2, 44, 282, 3, 2, 2,
	2, 46, 298, 3, 2, 2, 2, 48, 304, 3, 2, 2, 2, 50, 310, 3, 2, 2, 2, 52, 312,
	3, 2, 2, 2, 54, 314, 3, 2, 2, 2, 56, 317, 3, 2, 2, 2, 58, 328, 3, 2, 2,
	2, 60, 331, 3, 2, 2, 2, 62, 333, 3, 2, 2, 2, 64, 82, 7, 2, 2, 3, 65, 67,
	5, 56, 29, 2, 66, 65, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 71, 3, 2, 2,
	2, 68, 70, 5, 4, 3, 2, 69, 68, 3, 2, 2, 2, 70, 73, 3, 2, 2, 2, 71, 69,
	3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 77, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2,
	74, 76, 5, 6, 4, 2, 75, 74, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3,
	2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 80, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80,
	82, 7, 2, 2, 3, 81, 64, 3, 2, 2, 2, 81, 66, 3, 2, 2, 2, 82, 3, 3, 2, 2,
	2, 83, 85, 5, 58, 30, 2, 84, 83, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 86,
	3, 2, 2, 2, 86, 87, 7, 12, 2, 2, 87, 88, 5, 48, 25, 2, 88, 89, 7, 24, 2,
	2, 89, 100, 3, 2, 2, 2, 90, 92, 5, 58, 30, 2, 91, 90, 3, 2, 2, 2, 91, 92,
	3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 94, 7, 12, 2, 2, 94, 95, 5, 48, 25,
	2, 95, 96, 7, 13, 2, 2, 96, 97, 7, 69, 2, 2, 97, 98, 7, 24, 2, 2, 98, 100,
	3, 2, 2, 2, 99, 84, 3, 2, 2, 2, 99, 91, 3, 2, 2, 2, 100, 5, 3, 2, 2, 2,
	101, 103, 5, 32, 17, 2, 102, 101, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103,
	104, 3, 2, 2, 2, 104, 107, 7, 3, 2, 2, 105, 108, 7, 69, 2, 2, 106, 108,
	5, 38, 20, 2, 107, 105, 3, 2, 2, 2, 107, 106, 3, 2, 2, 2, 108, 110, 3,
	2, 2, 2, 109, 111, 5, 30, 16, 2, 110, 109, 3, 2, 2, 2, 110, 111, 3, 2,
	2, 2, 111, 112, 3, 2, 2, 2, 112, 113, 5, 8, 5, 2, 113, 7, 3, 2, 2, 2, 114,
	115, 7, 22, 2, 2, 115, 116, 5, 10, 6, 2, 116, 117, 7, 23, 2, 2, 117, 9,
	3, 2, 2, 2, 118, 124, 5, 12, 7, 2, 119, 124, 5, 16, 9, 2, 120, 124, 5,
	18, 10, 2, 121, 124, 5, 14, 8, 2, 122, 124, 5, 20, 11, 2, 123, 118, 3,
	2, 2, 2, 123, 119, 3, 2, 2, 2, 123, 120, 3, 2, 2, 2, 123, 121, 3, 2, 2,
	2, 123, 122, 3, 2, 2, 2, 124, 127, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 125,
	126, 3, 2, 2, 2, 126, 11, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 128, 130, 5,
	58, 30, 2, 129, 128, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 132, 3, 2,
	2, 2, 131, 133, 5, 32, 17, 2, 132, 131, 3, 2, 2, 2, 132, 133, 3, 2, 2,
	2, 133, 134, 3, 2, 2, 2, 134, 135, 7, 6, 2, 2, 135, 136, 7, 69, 2, 2, 136,
	137, 7, 28, 2, 2, 137, 138, 5, 38, 20, 2, 138, 139, 7, 24, 2, 2, 139, 13,
	3, 2, 2, 2, 140, 142, 5, 58, 30, 2, 141, 140, 3, 2, 2, 2, 141, 142, 3,
	2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 144, 7, 8, 2, 2, 144, 146, 7, 69, 2,
	2, 145, 147, 5, 30, 16, 2, 146, 145, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2,
	147, 148, 3, 2, 2, 2, 148, 149, 7, 24, 2, 2, 149, 15, 3, 2, 2, 2, 150,
	152, 5, 58, 30, 2, 151, 150, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 153,
	3, 2, 2, 2, 153, 154, 7, 9, 2, 2, 154, 155, 7, 69, 2, 2, 155, 156, 7, 14,
	2, 2, 156, 157, 7, 70, 2, 2, 157, 158, 7, 28, 2, 2, 158, 159, 5, 36, 19,
	2, 159, 160, 7, 24, 2, 2, 160, 17, 3, 2, 2, 2, 161, 163, 5, 58, 30, 2,
	162, 161, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 165, 3, 2, 2, 2, 164,
	166, 5, 32, 17, 2, 165, 164, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 167,
	3, 2, 2, 2, 167, 168, 7, 7, 2, 2, 168, 170, 7, 69, 2, 2, 169, 171, 5, 46,
	24, 2, 170, 169, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 173, 3, 2, 2, 2,
	172, 174, 5, 30, 16, 2, 173, 172, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174,
	175, 3, 2, 2, 2, 175, 176, 5, 8, 5, 2, 176, 19, 3, 2, 2, 2, 177, 179, 5,
	58, 30, 2, 178, 177, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 181, 3, 2,
	2, 2, 180, 182, 5, 32, 17, 2, 181, 180, 3, 2, 2, 2, 181, 182, 3, 2, 2,
	2, 182, 183, 3, 2, 2, 2, 183, 184, 7, 5, 2, 2, 184, 186, 7, 69, 2, 2, 185,
	187, 5, 46, 24, 2, 186, 185, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 188,
	3, 2, 2, 2, 188, 189, 7, 4, 2, 2, 189, 191, 7, 69, 2, 2, 190, 192, 5, 22,
	12, 2, 191, 190, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 194, 3, 2, 2, 2,
	193, 195, 5, 24, 13, 2, 194, 193, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195,
	196, 3, 2, 2, 2, 196, 197, 7, 24, 2, 2, 197, 21, 3, 2, 2, 2, 198, 199,
	7, 11, 2, 2, 199, 200, 7, 69, 2, 2, 200, 23, 3, 2, 2, 2, 201, 202, 7, 15,
	2, 2, 202, 203, 7, 20, 2, 2, 203, 204, 5, 26, 14, 2, 204, 205, 7, 21, 2,
	2, 205, 25, 3, 2, 2, 2, 206, 212, 5, 28, 15, 2, 207, 208, 5, 28, 15, 2,
	208, 209, 5, 62, 32, 2, 209, 210, 5, 26, 14, 2, 210, 212, 3, 2, 2, 2, 211,
	206, 3, 2, 2, 2, 211, 207, 3, 2, 2, 2, 212, 27, 3, 2, 2, 2, 213, 214, 7,
	69, 2, 2, 214, 215, 5, 60, 31, 2, 215, 216, 5, 54, 28, 2, 216, 29, 3, 2,
	2, 2, 217, 219, 7, 10, 2, 2, 218, 217, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2,
	219, 221, 3, 2, 2, 2, 220, 222, 5, 50, 26, 2, 221, 220, 3, 2, 2, 2, 221,
	222, 3, 2, 2, 2, 222, 31, 3, 2, 2, 2, 223, 225, 5, 34, 18, 2, 224, 223,
	3, 2, 2, 2, 225, 228, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2, 226, 227, 3, 2,
	2, 2, 227, 33, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 229, 230, 7, 18, 2, 2,
	230, 231, 7, 69, 2, 2, 231, 246, 7, 19, 2, 2, 232, 233, 7, 18, 2, 2, 233,
	234, 7, 69, 2, 2, 234, 235, 7, 20, 2, 2, 235, 236, 7, 69, 2, 2, 236, 237,
	7, 21, 2, 2, 237, 246, 7, 19, 2, 2, 238, 239, 7, 18, 2, 2, 239, 240, 7,
	69, 2, 2, 240, 241, 7, 20, 2, 2, 241, 242, 5, 54, 28, 2, 242, 243, 7, 21,
	2, 2, 243, 244, 7, 19, 2, 2, 244, 246, 3, 2, 2, 2, 245, 229, 3, 2, 2, 2,
	245, 232, 3, 2, 2, 2, 245, 238, 3, 2, 2, 2, 246, 35, 3, 2, 2, 2, 247, 253,
	5, 38, 20, 2, 248, 249, 5, 38, 20, 2, 249, 250, 7, 25, 2, 2, 250, 251,
	5, 36, 19, 2, 251, 253, 3, 2, 2, 2, 252, 247, 3, 2, 2, 2, 252, 248, 3,
	2, 2, 2, 253, 37, 3, 2, 2, 2, 254, 256, 7, 69, 2, 2, 255, 257, 5, 40, 21,
	2, 256, 255, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 265, 3, 2, 2, 2, 258,
	259, 5, 48, 25, 2, 259, 260, 7, 28, 2, 2, 260, 262, 7, 69, 2, 2, 261, 263,
	5, 40, 21, 2, 262, 261, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263, 265, 3,
	2, 2, 2, 264, 254, 3, 2, 2, 2, 264, 258, 3, 2, 2, 2, 265, 39, 3, 2, 2,
	2, 266, 267, 7, 26, 2, 2, 267, 274, 5, 54, 28, 2, 268, 269, 7, 26, 2, 2,
	269, 270, 7, 22, 2, 2, 270, 271, 5, 42, 22, 2, 271, 272, 7, 23, 2, 2, 272,
	274, 3, 2, 2, 2, 273, 266, 3, 2, 2, 2, 273, 268, 3, 2, 2, 2, 274, 41, 3,
	2, 2, 2, 275, 281, 5, 44, 23, 2, 276, 277, 5, 44, 23, 2, 277, 278, 7, 25,
	2, 2, 278, 279, 5, 42, 22, 2, 279, 281, 3, 2, 2, 2, 280, 275, 3, 2, 2,
	2, 280, 276, 3, 2, 2, 2, 281, 43, 3, 2, 2, 2, 282, 283, 7, 69, 2, 2, 283,
	284, 7, 26, 2, 2, 284, 285, 5, 54, 28, 2, 285, 45, 3, 2, 2, 2, 286, 287,
	7, 18, 2, 2, 287, 288, 5, 54, 28, 2, 288, 289, 7, 25, 2, 2, 289, 290, 5,
	54, 28, 2, 290, 291, 7, 19, 2, 2, 291, 299, 3, 2, 2, 2, 292, 293, 7, 18,
	2, 2, 293, 294, 5, 54, 28, 2, 294, 295, 7, 25, 2, 2, 295, 296, 7, 68, 2,
	2, 296, 297, 7, 19, 2, 2, 297, 299, 3, 2, 2, 2, 298, 286, 3, 2, 2, 2, 298,
	292, 3, 2, 2, 2, 299, 47, 3, 2, 2, 2, 300, 305, 7, 69, 2, 2, 301, 302,
	7, 69, 2, 2, 302, 303, 7, 30, 2, 2, 303, 305, 5, 48, 25, 2, 304, 300, 3,
	2, 2, 2, 304, 301, 3, 2, 2, 2, 305, 49, 3, 2, 2, 2, 306, 311, 7, 69, 2,
	2, 307, 308, 7, 69, 2, 2, 308, 309, 7, 25, 2, 2, 309, 311, 5, 50, 26, 2,
	310, 306, 3, 2, 2, 2, 310, 307, 3, 2, 2, 2, 311, 51, 3, 2, 2, 2, 312, 313,
	9, 2, 2, 2, 313, 53, 3, 2, 2, 2, 314, 315, 9, 3, 2, 2, 315, 55, 3, 2, 2,
	2, 316, 318, 7, 16, 2, 2, 317, 316, 3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318,
	322, 3, 2, 2, 2, 319, 321, 7, 17, 2, 2, 320, 319, 3, 2, 2, 2, 321, 324,
	3, 2, 2, 2, 322, 320, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323, 57, 3, 2,
	2, 2, 324, 322, 3, 2, 2, 2, 325, 327, 9, 4, 2, 2, 326, 325, 3, 2, 2, 2,
	327, 330, 3, 2, 2, 2, 328, 326, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329,
	59, 3, 2, 2, 2, 330, 328, 3, 2, 2, 2, 331, 332, 9, 5, 2, 2, 332, 61, 3,
	2, 2, 2, 333, 334, 9, 6, 2, 2, 334, 63, 3, 2, 2, 2, 45, 66, 71, 77, 81,
	84, 91, 99, 102, 107, 110, 123, 125, 129, 132, 141, 146, 151, 162, 165,
	170, 173, 178, 181, 186, 191, 194, 211, 218, 221, 226, 245, 252, 256, 262,
	264, 273, 280, 298, 304, 310, 317, 322, 328,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'businessobject'", "'to'", "'association'", "'element'", "'node'",
	"'action'", "'message'", "'raises'", "'using'", "'import'", "'as'", "'text'",
	"'valuation'", "", "", "'['", "']'", "'('", "')'", "'{'", "'}'", "';'",
	"','", "'='", "'?'", "':'", "'...'", "'.'", "'++'", "'--'", "'+'", "'-'",
	"'~'", "'!'", "'*'", "'/'", "'%'", "'>>'", "'<<'", "'>>>'", "'<'", "'>'",
	"'<='", "'>='", "'=='", "'!='", "'==='", "'!=='", "'&'", "'^'", "'|'",
	"'&&'", "'||'", "'*='", "'/='", "'%='", "'+='", "'-='", "'<<='", "'>>='",
	"'>>>='", "'&='", "'^='", "'|='", "'=>'", "'n'",
}
var symbolicNames = []string{
	"", "BUSINESSOBJECT", "TO", "ASSOCIATION", "ELEMENT", "NODE", "ACTION",
	"MESSAGE", "RAISES", "USING", "IMPORT", "AS", "TEXT", "VALUATION", "MultiLineComment",
	"SingleLineComment", "OpenBracket", "CloseBracket", "OpenParen", "CloseParen",
	"OpenBrace", "CloseBrace", "SemiColon", "Comma", "Assign", "QuestionMark",
	"Colon", "Ellipsis", "Dot", "PlusPlus", "MinusMinus", "Plus", "Minus",
	"BitNot", "Not", "Multiply", "Divide", "Modulus", "RightShiftArithmetic",
	"LeftShiftArithmetic", "RightShiftLogical", "LessThan", "MoreThan", "LessThanEquals",
	"GreaterThanEquals", "Equals_", "NotEquals", "IdentityEquals", "IdentityNotEquals",
	"BitAnd", "BitXOr", "BitOr", "And", "Or", "MultiplyAssign", "DivideAssign",
	"ModulusAssign", "PlusAssign", "MinusAssign", "LeftShiftArithmeticAssign",
	"RightShiftArithmeticAssign", "RightShiftLogicalAssign", "BitAndAssign",
	"BitXorAssign", "BitOrAssign", "ARROW", "N", "Identifier", "StringLiteral",
	"BooleanLiteral", "DecimalLiteral", "HexIntegerLiteral", "OctalIntegerLiteral",
	"OctalIntegerLiteral2", "BinaryIntegerLiteral", "TemplateStringLiteral",
	"WhiteSpaces", "LineTerminator",
}

var ruleNames = []string{
	"program", "importStatment", "definition", "block", "itemList", "element",
	"boAction", "message", "node", "association", "associationUsingDefinition",
	"valuationDefinition", "valutaionExpressionList", "valutaionExpression",
	"raiseMessage", "annotationList", "annotation", "typeList", "type", "typeDefaultValue",
	"valueAssignList", "valueAssign", "multiplicity", "memberExpression", "identifierList",
	"keyword", "literal", "copyright", "comments", "compareOperator", "logicOperator",
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
	BODLParserEOF                        = antlr.TokenEOF
	BODLParserBUSINESSOBJECT             = 1
	BODLParserTO                         = 2
	BODLParserASSOCIATION                = 3
	BODLParserELEMENT                    = 4
	BODLParserNODE                       = 5
	BODLParserACTION                     = 6
	BODLParserMESSAGE                    = 7
	BODLParserRAISES                     = 8
	BODLParserUSING                      = 9
	BODLParserIMPORT                     = 10
	BODLParserAS                         = 11
	BODLParserTEXT                       = 12
	BODLParserVALUATION                  = 13
	BODLParserMultiLineComment           = 14
	BODLParserSingleLineComment          = 15
	BODLParserOpenBracket                = 16
	BODLParserCloseBracket               = 17
	BODLParserOpenParen                  = 18
	BODLParserCloseParen                 = 19
	BODLParserOpenBrace                  = 20
	BODLParserCloseBrace                 = 21
	BODLParserSemiColon                  = 22
	BODLParserComma                      = 23
	BODLParserAssign                     = 24
	BODLParserQuestionMark               = 25
	BODLParserColon                      = 26
	BODLParserEllipsis                   = 27
	BODLParserDot                        = 28
	BODLParserPlusPlus                   = 29
	BODLParserMinusMinus                 = 30
	BODLParserPlus                       = 31
	BODLParserMinus                      = 32
	BODLParserBitNot                     = 33
	BODLParserNot                        = 34
	BODLParserMultiply                   = 35
	BODLParserDivide                     = 36
	BODLParserModulus                    = 37
	BODLParserRightShiftArithmetic       = 38
	BODLParserLeftShiftArithmetic        = 39
	BODLParserRightShiftLogical          = 40
	BODLParserLessThan                   = 41
	BODLParserMoreThan                   = 42
	BODLParserLessThanEquals             = 43
	BODLParserGreaterThanEquals          = 44
	BODLParserEquals_                    = 45
	BODLParserNotEquals                  = 46
	BODLParserIdentityEquals             = 47
	BODLParserIdentityNotEquals          = 48
	BODLParserBitAnd                     = 49
	BODLParserBitXOr                     = 50
	BODLParserBitOr                      = 51
	BODLParserAnd                        = 52
	BODLParserOr                         = 53
	BODLParserMultiplyAssign             = 54
	BODLParserDivideAssign               = 55
	BODLParserModulusAssign              = 56
	BODLParserPlusAssign                 = 57
	BODLParserMinusAssign                = 58
	BODLParserLeftShiftArithmeticAssign  = 59
	BODLParserRightShiftArithmeticAssign = 60
	BODLParserRightShiftLogicalAssign    = 61
	BODLParserBitAndAssign               = 62
	BODLParserBitXorAssign               = 63
	BODLParserBitOrAssign                = 64
	BODLParserARROW                      = 65
	BODLParserN                          = 66
	BODLParserIdentifier                 = 67
	BODLParserStringLiteral              = 68
	BODLParserBooleanLiteral             = 69
	BODLParserDecimalLiteral             = 70
	BODLParserHexIntegerLiteral          = 71
	BODLParserOctalIntegerLiteral        = 72
	BODLParserOctalIntegerLiteral2       = 73
	BODLParserBinaryIntegerLiteral       = 74
	BODLParserTemplateStringLiteral      = 75
	BODLParserWhiteSpaces                = 76
	BODLParserLineTerminator             = 77
)

// BODLParser rules.
const (
	BODLParserRULE_program                    = 0
	BODLParserRULE_importStatment             = 1
	BODLParserRULE_definition                 = 2
	BODLParserRULE_block                      = 3
	BODLParserRULE_itemList                   = 4
	BODLParserRULE_element                    = 5
	BODLParserRULE_boAction                   = 6
	BODLParserRULE_message                    = 7
	BODLParserRULE_node                       = 8
	BODLParserRULE_association                = 9
	BODLParserRULE_associationUsingDefinition = 10
	BODLParserRULE_valuationDefinition        = 11
	BODLParserRULE_valutaionExpressionList    = 12
	BODLParserRULE_valutaionExpression        = 13
	BODLParserRULE_raiseMessage               = 14
	BODLParserRULE_annotationList             = 15
	BODLParserRULE_annotation                 = 16
	BODLParserRULE_typeList                   = 17
	BODLParserRULE_type                       = 18
	BODLParserRULE_typeDefaultValue           = 19
	BODLParserRULE_valueAssignList            = 20
	BODLParserRULE_valueAssign                = 21
	BODLParserRULE_multiplicity               = 22
	BODLParserRULE_memberExpression           = 23
	BODLParserRULE_identifierList             = 24
	BODLParserRULE_keyword                    = 25
	BODLParserRULE_literal                    = 26
	BODLParserRULE_copyright                  = 27
	BODLParserRULE_comments                   = 28
	BODLParserRULE_compareOperator            = 29
	BODLParserRULE_logicOperator              = 30
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

func (s *ProgramContext) Copyright() ICopyrightContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopyrightContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICopyrightContext)
}

func (s *ProgramContext) AllImportStatment() []IImportStatmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportStatmentContext)(nil)).Elem())
	var tst = make([]IImportStatmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportStatmentContext)
		}
	}

	return tst
}

func (s *ProgramContext) ImportStatment(i int) IImportStatmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportStatmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportStatmentContext)
}

func (s *ProgramContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *ProgramContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
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

func (p *BODLParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BODLParserRULE_program)
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

	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Match(BODLParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(64)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(63)
				p.Copyright()
			}

		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BODLParserIMPORT)|(1<<BODLParserMultiLineComment)|(1<<BODLParserSingleLineComment))) != 0 {
			{
				p.SetState(66)
				p.ImportStatment()
			}

			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == BODLParserBUSINESSOBJECT || _la == BODLParserOpenBracket {
			{
				p.SetState(72)
				p.Definition()
			}

			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(78)
			p.Match(BODLParserEOF)
		}

	}

	return localctx
}

// IImportStatmentContext is an interface to support dynamic dispatch.
type IImportStatmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportStatmentContext differentiates from other interfaces.
	IsImportStatmentContext()
}

type ImportStatmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatmentContext() *ImportStatmentContext {
	var p = new(ImportStatmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_importStatment
	return p
}

func (*ImportStatmentContext) IsImportStatmentContext() {}

func NewImportStatmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatmentContext {
	var p = new(ImportStatmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_importStatment

	return p
}

func (s *ImportStatmentContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatmentContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(BODLParserIMPORT, 0)
}

func (s *ImportStatmentContext) MemberExpression() IMemberExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberExpressionContext)
}

func (s *ImportStatmentContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(BODLParserSemiColon, 0)
}

func (s *ImportStatmentContext) Comments() ICommentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentsContext)
}

func (s *ImportStatmentContext) AS() antlr.TerminalNode {
	return s.GetToken(BODLParserAS, 0)
}

func (s *ImportStatmentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BODLParserIdentifier, 0)
}

func (s *ImportStatmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterImportStatment(s)
	}
}

func (s *ImportStatmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitImportStatment(s)
	}
}

func (p *BODLParser) ImportStatment() (localctx IImportStatmentContext) {
	localctx = NewImportStatmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BODLParserRULE_importStatment)

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

	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(82)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(81)
				p.Comments()
			}

		}
		{
			p.SetState(84)
			p.Match(BODLParserIMPORT)
		}
		{
			p.SetState(85)
			p.MemberExpression()
		}
		{
			p.SetState(86)
			p.Match(BODLParserSemiColon)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(89)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(88)
				p.Comments()
			}

		}
		{
			p.SetState(91)
			p.Match(BODLParserIMPORT)
		}
		{
			p.SetState(92)
			p.MemberExpression()
		}
		{
			p.SetState(93)
			p.Match(BODLParserAS)
		}
		{
			p.SetState(94)
			p.Match(BODLParserIdentifier)
		}
		{
			p.SetState(95)
			p.Match(BODLParserSemiColon)
		}

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

func (s *DefinitionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BODLParserIdentifier, 0)
}

func (s *DefinitionContext) Type() ITypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DefinitionContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
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

func (p *BODLParser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BODLParserRULE_definition)

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

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(99)
			p.AnnotationList()
		}

	}
	{
		p.SetState(102)
		p.Match(BODLParserBUSINESSOBJECT)
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(103)
			p.Match(BODLParserIdentifier)
		}

	case 2:
		{
			p.SetState(104)
			p.Type()
		}

	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(107)
			p.RaiseMessage()
		}

	}
	{
		p.SetState(110)
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

func (p *BODLParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BODLParserRULE_block)

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
		p.SetState(112)
		p.Match(BODLParserOpenBrace)
	}
	{
		p.SetState(113)
		p.ItemList()
	}
	{
		p.SetState(114)
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

func (p *BODLParser) ItemList() (localctx IItemListContext) {
	localctx = NewItemListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BODLParserRULE_itemList)
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
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BODLParserASSOCIATION)|(1<<BODLParserELEMENT)|(1<<BODLParserNODE)|(1<<BODLParserACTION)|(1<<BODLParserMESSAGE)|(1<<BODLParserMultiLineComment)|(1<<BODLParserSingleLineComment)|(1<<BODLParserOpenBracket))) != 0 {
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(116)
				p.Element()
			}

		case 2:
			{
				p.SetState(117)
				p.Message()
			}

		case 3:
			{
				p.SetState(118)
				p.Node()
			}

		case 4:
			{
				p.SetState(119)
				p.BoAction()
			}

		case 5:
			{
				p.SetState(120)
				p.Association()
			}

		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *ElementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BODLParserIdentifier, 0)
}

func (s *ElementContext) Colon() antlr.TerminalNode {
	return s.GetToken(BODLParserColon, 0)
}

func (s *ElementContext) Type() ITypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
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

func (s *ElementContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
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

func (p *BODLParser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BODLParserRULE_element)

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
	p.SetState(127)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(126)
			p.Comments()
		}

	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(129)
			p.AnnotationList()
		}

	}
	{
		p.SetState(132)
		p.Match(BODLParserELEMENT)
	}
	{
		p.SetState(133)
		p.Match(BODLParserIdentifier)
	}
	{
		p.SetState(134)
		p.Match(BODLParserColon)
	}
	{
		p.SetState(135)
		p.Type()
	}
	{
		p.SetState(136)
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

func (s *BoActionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BODLParserIdentifier, 0)
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

func (p *BODLParser) BoAction() (localctx IBoActionContext) {
	localctx = NewBoActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BODLParserRULE_boAction)

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
	p.SetState(139)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(138)
			p.Comments()
		}

	}
	{
		p.SetState(141)
		p.Match(BODLParserACTION)
	}
	{
		p.SetState(142)
		p.Match(BODLParserIdentifier)
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(143)
			p.RaiseMessage()
		}

	}
	{
		p.SetState(146)
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

func (s *MessageContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BODLParserIdentifier, 0)
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

func (p *BODLParser) Message() (localctx IMessageContext) {
	localctx = NewMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BODLParserRULE_message)

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
	p.SetState(149)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(148)
			p.Comments()
		}

	}
	{
		p.SetState(151)
		p.Match(BODLParserMESSAGE)
	}
	{
		p.SetState(152)
		p.Match(BODLParserIdentifier)
	}
	{
		p.SetState(153)
		p.Match(BODLParserTEXT)
	}
	{
		p.SetState(154)
		p.Match(BODLParserStringLiteral)
	}
	{
		p.SetState(155)
		p.Match(BODLParserColon)
	}
	{
		p.SetState(156)
		p.TypeList()
	}
	{
		p.SetState(157)
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

func (s *NodeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BODLParserIdentifier, 0)
}

func (s *NodeContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *NodeContext) Comments() ICommentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentsContext)
}

func (s *NodeContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
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

func (p *BODLParser) Node() (localctx INodeContext) {
	localctx = NewNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BODLParserRULE_node)
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
			p.AnnotationList()
		}

	}
	{
		p.SetState(165)
		p.Match(BODLParserNODE)
	}
	{
		p.SetState(166)
		p.Match(BODLParserIdentifier)
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserOpenBracket {
		{
			p.SetState(167)
			p.Multiplicity()
		}

	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(170)
			p.RaiseMessage()
		}

	}
	{
		p.SetState(173)
		p.Block()
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

func (s *AssociationContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(BODLParserIdentifier)
}

func (s *AssociationContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(BODLParserIdentifier, i)
}

func (s *AssociationContext) TO() antlr.TerminalNode {
	return s.GetToken(BODLParserTO, 0)
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

func (s *AssociationContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
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

func (p *BODLParser) Association() (localctx IAssociationContext) {
	localctx = NewAssociationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BODLParserRULE_association)
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
	p.SetState(176)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(175)
			p.Comments()
		}

	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(178)
			p.AnnotationList()
		}

	}
	{
		p.SetState(181)
		p.Match(BODLParserASSOCIATION)
	}
	{
		p.SetState(182)
		p.Match(BODLParserIdentifier)
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserOpenBracket {
		{
			p.SetState(183)
			p.Multiplicity()
		}

	}
	{
		p.SetState(186)
		p.Match(BODLParserTO)
	}
	{
		p.SetState(187)
		p.Match(BODLParserIdentifier)
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserUSING {
		{
			p.SetState(188)
			p.AssociationUsingDefinition()
		}

	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserVALUATION {
		{
			p.SetState(191)
			p.ValuationDefinition()
		}

	}
	{
		p.SetState(194)
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

func (s *AssociationUsingDefinitionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BODLParserIdentifier, 0)
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

func (p *BODLParser) AssociationUsingDefinition() (localctx IAssociationUsingDefinitionContext) {
	localctx = NewAssociationUsingDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BODLParserRULE_associationUsingDefinition)

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
		p.SetState(196)
		p.Match(BODLParserUSING)
	}
	{
		p.SetState(197)
		p.Match(BODLParserIdentifier)
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

func (p *BODLParser) ValuationDefinition() (localctx IValuationDefinitionContext) {
	localctx = NewValuationDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BODLParserRULE_valuationDefinition)

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
		p.SetState(199)
		p.Match(BODLParserVALUATION)
	}
	{
		p.SetState(200)
		p.Match(BODLParserOpenParen)
	}
	{
		p.SetState(201)
		p.ValutaionExpressionList()
	}
	{
		p.SetState(202)
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

func (p *BODLParser) ValutaionExpressionList() (localctx IValutaionExpressionListContext) {
	localctx = NewValutaionExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BODLParserRULE_valutaionExpressionList)

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

	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(204)
			p.ValutaionExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(205)
			p.ValutaionExpression()
		}
		{
			p.SetState(206)
			p.LogicOperator()
		}
		{
			p.SetState(207)
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

func (s *ValutaionExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BODLParserIdentifier, 0)
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

func (p *BODLParser) ValutaionExpression() (localctx IValutaionExpressionContext) {
	localctx = NewValutaionExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BODLParserRULE_valutaionExpression)

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
		p.SetState(211)
		p.Match(BODLParserIdentifier)
	}
	{
		p.SetState(212)
		p.CompareOperator()
	}
	{
		p.SetState(213)
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

func (p *BODLParser) RaiseMessage() (localctx IRaiseMessageContext) {
	localctx = NewRaiseMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BODLParserRULE_raiseMessage)
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
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserRAISES {
		{
			p.SetState(215)
			p.Match(BODLParserRAISES)
		}

	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserIdentifier {
		{
			p.SetState(218)
			p.IdentifierList()
		}

	}

	return localctx
}

// IAnnotationListContext is an interface to support dynamic dispatch.
type IAnnotationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationListContext differentiates from other interfaces.
	IsAnnotationListContext()
}

type AnnotationListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationListContext() *AnnotationListContext {
	var p = new(AnnotationListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_annotationList
	return p
}

func (*AnnotationListContext) IsAnnotationListContext() {}

func NewAnnotationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationListContext {
	var p = new(AnnotationListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_annotationList

	return p
}

func (s *AnnotationListContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationListContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *AnnotationListContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *AnnotationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterAnnotationList(s)
	}
}

func (s *AnnotationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitAnnotationList(s)
	}
}

func (p *BODLParser) AnnotationList() (localctx IAnnotationListContext) {
	localctx = NewAnnotationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BODLParserRULE_annotationList)
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
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BODLParserOpenBracket {
		{
			p.SetState(221)
			p.Annotation()
		}

		p.SetState(226)
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

func (s *AnnotationContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(BODLParserIdentifier)
}

func (s *AnnotationContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(BODLParserIdentifier, i)
}

func (s *AnnotationContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(BODLParserCloseBracket, 0)
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

func (p *BODLParser) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BODLParserRULE_annotation)

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

	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)
			p.Match(BODLParserOpenBracket)
		}
		{
			p.SetState(228)
			p.Match(BODLParserIdentifier)
		}
		{
			p.SetState(229)
			p.Match(BODLParserCloseBracket)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			p.Match(BODLParserOpenBracket)
		}
		{
			p.SetState(231)
			p.Match(BODLParserIdentifier)
		}
		{
			p.SetState(232)
			p.Match(BODLParserOpenParen)
		}
		{
			p.SetState(233)
			p.Match(BODLParserIdentifier)
		}
		{
			p.SetState(234)
			p.Match(BODLParserCloseParen)
		}
		{
			p.SetState(235)
			p.Match(BODLParserCloseBracket)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(236)
			p.Match(BODLParserOpenBracket)
		}
		{
			p.SetState(237)
			p.Match(BODLParserIdentifier)
		}
		{
			p.SetState(238)
			p.Match(BODLParserOpenParen)
		}
		{
			p.SetState(239)
			p.Literal()
		}
		{
			p.SetState(240)
			p.Match(BODLParserCloseParen)
		}
		{
			p.SetState(241)
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

func (s *TypeListContext) Type() ITypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
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

func (p *BODLParser) TypeList() (localctx ITypeListContext) {
	localctx = NewTypeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BODLParserRULE_typeList)

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

	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(245)
			p.Type()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(246)
			p.Type()
		}
		{
			p.SetState(247)
			p.Match(BODLParserComma)
		}
		{
			p.SetState(248)
			p.TypeList()
		}

	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BODLParserIdentifier, 0)
}

func (s *TypeContext) TypeDefaultValue() ITypeDefaultValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDefaultValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDefaultValueContext)
}

func (s *TypeContext) MemberExpression() IMemberExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberExpressionContext)
}

func (s *TypeContext) Colon() antlr.TerminalNode {
	return s.GetToken(BODLParserColon, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *BODLParser) Type() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BODLParserRULE_type)
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

	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(252)
			p.Match(BODLParserIdentifier)
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BODLParserAssign {
			{
				p.SetState(253)
				p.TypeDefaultValue()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(256)
			p.MemberExpression()
		}
		{
			p.SetState(257)
			p.Match(BODLParserColon)
		}
		{
			p.SetState(258)
			p.Match(BODLParserIdentifier)
		}
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BODLParserAssign {
			{
				p.SetState(259)
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

func (p *BODLParser) TypeDefaultValue() (localctx ITypeDefaultValueContext) {
	localctx = NewTypeDefaultValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BODLParserRULE_typeDefaultValue)

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

	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			p.Match(BODLParserAssign)
		}
		{
			p.SetState(265)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.Match(BODLParserAssign)
		}
		{
			p.SetState(267)
			p.Match(BODLParserOpenBrace)
		}
		{
			p.SetState(268)
			p.ValueAssignList()
		}
		{
			p.SetState(269)
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

func (p *BODLParser) ValueAssignList() (localctx IValueAssignListContext) {
	localctx = NewValueAssignListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BODLParserRULE_valueAssignList)

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

	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(273)
			p.ValueAssign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
			p.ValueAssign()
		}
		{
			p.SetState(275)
			p.Match(BODLParserComma)
		}
		{
			p.SetState(276)
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

func (s *ValueAssignContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BODLParserIdentifier, 0)
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

func (p *BODLParser) ValueAssign() (localctx IValueAssignContext) {
	localctx = NewValueAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BODLParserRULE_valueAssign)

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
		p.SetState(280)
		p.Match(BODLParserIdentifier)
	}
	{
		p.SetState(281)
		p.Match(BODLParserAssign)
	}
	{
		p.SetState(282)
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

func (p *BODLParser) Multiplicity() (localctx IMultiplicityContext) {
	localctx = NewMultiplicityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BODLParserRULE_multiplicity)

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

	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(284)
			p.Match(BODLParserOpenBracket)
		}
		{
			p.SetState(285)
			p.Literal()
		}
		{
			p.SetState(286)
			p.Match(BODLParserComma)
		}
		{
			p.SetState(287)
			p.Literal()
		}
		{
			p.SetState(288)
			p.Match(BODLParserCloseBracket)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(290)
			p.Match(BODLParserOpenBracket)
		}
		{
			p.SetState(291)
			p.Literal()
		}
		{
			p.SetState(292)
			p.Match(BODLParserComma)
		}
		{
			p.SetState(293)
			p.Match(BODLParserN)
		}
		{
			p.SetState(294)
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

func (s *MemberExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BODLParserIdentifier, 0)
}

func (s *MemberExpressionContext) Dot() antlr.TerminalNode {
	return s.GetToken(BODLParserDot, 0)
}

func (s *MemberExpressionContext) MemberExpression() IMemberExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberExpressionContext)(nil)).Elem(), 0)

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

func (p *BODLParser) MemberExpression() (localctx IMemberExpressionContext) {
	localctx = NewMemberExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BODLParserRULE_memberExpression)

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

	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(298)
			p.Match(BODLParserIdentifier)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(299)
			p.Match(BODLParserIdentifier)
		}
		{
			p.SetState(300)
			p.Match(BODLParserDot)
		}
		{
			p.SetState(301)
			p.MemberExpression()
		}

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

func (s *IdentifierListContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BODLParserIdentifier, 0)
}

func (s *IdentifierListContext) Comma() antlr.TerminalNode {
	return s.GetToken(BODLParserComma, 0)
}

func (s *IdentifierListContext) IdentifierList() IIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
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

func (p *BODLParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BODLParserRULE_identifierList)

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

	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(304)
			p.Match(BODLParserIdentifier)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(305)
			p.Match(BODLParserIdentifier)
		}
		{
			p.SetState(306)
			p.Match(BODLParserComma)
		}
		{
			p.SetState(307)
			p.IdentifierList()
		}

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

func (p *BODLParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, BODLParserRULE_keyword)
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
		p.SetState(310)
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

func (p *BODLParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BODLParserRULE_literal)
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
		p.SetState(312)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-68)&-(0x1f+1)) == 0 && ((1<<uint((_la-68)))&((1<<(BODLParserStringLiteral-68))|(1<<(BODLParserBooleanLiteral-68))|(1<<(BODLParserDecimalLiteral-68)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICopyrightContext is an interface to support dynamic dispatch.
type ICopyrightContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCopyrightContext differentiates from other interfaces.
	IsCopyrightContext()
}

type CopyrightContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCopyrightContext() *CopyrightContext {
	var p = new(CopyrightContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_copyright
	return p
}

func (*CopyrightContext) IsCopyrightContext() {}

func NewCopyrightContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopyrightContext {
	var p = new(CopyrightContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_copyright

	return p
}

func (s *CopyrightContext) GetParser() antlr.Parser { return s.parser }

func (s *CopyrightContext) MultiLineComment() antlr.TerminalNode {
	return s.GetToken(BODLParserMultiLineComment, 0)
}

func (s *CopyrightContext) AllSingleLineComment() []antlr.TerminalNode {
	return s.GetTokens(BODLParserSingleLineComment)
}

func (s *CopyrightContext) SingleLineComment(i int) antlr.TerminalNode {
	return s.GetToken(BODLParserSingleLineComment, i)
}

func (s *CopyrightContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyrightContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CopyrightContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterCopyright(s)
	}
}

func (s *CopyrightContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitCopyright(s)
	}
}

func (p *BODLParser) Copyright() (localctx ICopyrightContext) {
	localctx = NewCopyrightContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BODLParserRULE_copyright)

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
	p.SetState(315)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(314)
			p.Match(BODLParserMultiLineComment)
		}

	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(317)
				p.Match(BODLParserSingleLineComment)
			}

		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
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

func (p *BODLParser) Comments() (localctx ICommentsContext) {
	localctx = NewCommentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BODLParserRULE_comments)
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
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BODLParserMultiLineComment || _la == BODLParserSingleLineComment {
		{
			p.SetState(323)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BODLParserMultiLineComment || _la == BODLParserSingleLineComment) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(328)
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

func (p *BODLParser) CompareOperator() (localctx ICompareOperatorContext) {
	localctx = NewCompareOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BODLParserRULE_compareOperator)
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
		p.SetState(329)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(BODLParserLessThan-41))|(1<<(BODLParserMoreThan-41))|(1<<(BODLParserLessThanEquals-41))|(1<<(BODLParserGreaterThanEquals-41))|(1<<(BODLParserEquals_-41))|(1<<(BODLParserNotEquals-41)))) != 0) {
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

func (p *BODLParser) LogicOperator() (localctx ILogicOperatorContext) {
	localctx = NewLogicOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, BODLParserRULE_logicOperator)
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
		p.SetState(331)
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
