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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 79, 314,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 3, 2, 5, 2, 66, 10,
	2, 3, 2, 5, 2, 69, 10, 2, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 75, 10, 3, 12,
	3, 14, 3, 78, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 5, 4, 90, 10, 4, 3, 5, 5, 5, 93, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5,
	98, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 7, 7, 111, 10, 7, 12, 7, 14, 7, 114, 11, 7, 3, 8, 5, 8, 117, 10,
	8, 3, 8, 5, 8, 120, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 5,
	9, 129, 10, 9, 3, 9, 3, 9, 3, 9, 5, 9, 134, 10, 9, 3, 9, 3, 9, 3, 10, 5,
	10, 139, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 11, 5, 11, 150, 10, 11, 3, 11, 5, 11, 153, 10, 11, 3, 11, 3, 11, 3,
	11, 5, 11, 158, 10, 11, 3, 11, 5, 11, 161, 10, 11, 3, 11, 3, 11, 3, 12,
	5, 12, 166, 10, 12, 3, 12, 5, 12, 169, 10, 12, 3, 12, 3, 12, 3, 12, 5,
	12, 174, 10, 12, 3, 12, 3, 12, 3, 12, 5, 12, 179, 10, 12, 3, 12, 5, 12,
	182, 10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 199, 10, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 17, 5, 17, 206, 10, 17, 3, 17, 5, 17, 209, 10,
	17, 3, 18, 7, 18, 212, 10, 18, 12, 18, 14, 18, 215, 11, 18, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 5, 19, 233, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 5, 20, 240, 10, 20, 3, 21, 3, 21, 5, 21, 244, 10, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 5, 21, 250, 10, 21, 5, 21, 252, 10, 21, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 261, 10, 22, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 5, 23, 268, 10, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 5, 25, 286, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 292, 10, 26,
	3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 298, 10, 27, 3, 28, 3, 28, 3, 29, 3,
	29, 3, 30, 7, 30, 305, 10, 30, 12, 30, 14, 30, 308, 11, 30, 3, 31, 3, 31,
	3, 32, 3, 32, 3, 32, 2, 2, 33, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
	60, 62, 2, 7, 3, 2, 3, 14, 3, 2, 70, 72, 3, 2, 16, 17, 3, 2, 43, 48, 3,
	2, 54, 55, 2, 323, 2, 65, 3, 2, 2, 2, 4, 76, 3, 2, 2, 2, 6, 89, 3, 2, 2,
	2, 8, 92, 3, 2, 2, 2, 10, 101, 3, 2, 2, 2, 12, 112, 3, 2, 2, 2, 14, 116,
	3, 2, 2, 2, 16, 128, 3, 2, 2, 2, 18, 138, 3, 2, 2, 2, 20, 149, 3, 2, 2,
	2, 22, 165, 3, 2, 2, 2, 24, 185, 3, 2, 2, 2, 26, 188, 3, 2, 2, 2, 28, 198,
	3, 2, 2, 2, 30, 200, 3, 2, 2, 2, 32, 205, 3, 2, 2, 2, 34, 213, 3, 2, 2,
	2, 36, 232, 3, 2, 2, 2, 38, 239, 3, 2, 2, 2, 40, 251, 3, 2, 2, 2, 42, 260,
	3, 2, 2, 2, 44, 267, 3, 2, 2, 2, 46, 269, 3, 2, 2, 2, 48, 285, 3, 2, 2,
	2, 50, 291, 3, 2, 2, 2, 52, 297, 3, 2, 2, 2, 54, 299, 3, 2, 2, 2, 56, 301,
	3, 2, 2, 2, 58, 306, 3, 2, 2, 2, 60, 309, 3, 2, 2, 2, 62, 311, 3, 2, 2,
	2, 64, 66, 5, 58, 30, 2, 65, 64, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 68,
	3, 2, 2, 2, 67, 69, 5, 4, 3, 2, 68, 67, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2,
	69, 70, 3, 2, 2, 2, 70, 71, 5, 8, 5, 2, 71, 72, 7, 2, 2, 3, 72, 3, 3, 2,
	2, 2, 73, 75, 5, 6, 4, 2, 74, 73, 3, 2, 2, 2, 75, 78, 3, 2, 2, 2, 76, 74,
	3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 5, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2,
	79, 80, 7, 12, 2, 2, 80, 81, 5, 50, 26, 2, 81, 82, 7, 24, 2, 2, 82, 90,
	3, 2, 2, 2, 83, 84, 7, 12, 2, 2, 84, 85, 5, 50, 26, 2, 85, 86, 7, 13, 2,
	2, 86, 87, 7, 69, 2, 2, 87, 88, 7, 24, 2, 2, 88, 90, 3, 2, 2, 2, 89, 79,
	3, 2, 2, 2, 89, 83, 3, 2, 2, 2, 90, 7, 3, 2, 2, 2, 91, 93, 5, 34, 18, 2,
	92, 91, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 95, 7,
	3, 2, 2, 95, 97, 7, 69, 2, 2, 96, 98, 5, 32, 17, 2, 97, 96, 3, 2, 2, 2,
	97, 98, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 100, 5, 10, 6, 2, 100, 9, 3,
	2, 2, 2, 101, 102, 7, 22, 2, 2, 102, 103, 5, 12, 7, 2, 103, 104, 7, 23,
	2, 2, 104, 11, 3, 2, 2, 2, 105, 111, 5, 14, 8, 2, 106, 111, 5, 18, 10,
	2, 107, 111, 5, 20, 11, 2, 108, 111, 5, 16, 9, 2, 109, 111, 5, 22, 12,
	2, 110, 105, 3, 2, 2, 2, 110, 106, 3, 2, 2, 2, 110, 107, 3, 2, 2, 2, 110,
	108, 3, 2, 2, 2, 110, 109, 3, 2, 2, 2, 111, 114, 3, 2, 2, 2, 112, 110,
	3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 13, 3, 2, 2, 2, 114, 112, 3, 2,
	2, 2, 115, 117, 5, 58, 30, 2, 116, 115, 3, 2, 2, 2, 116, 117, 3, 2, 2,
	2, 117, 119, 3, 2, 2, 2, 118, 120, 5, 34, 18, 2, 119, 118, 3, 2, 2, 2,
	119, 120, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 122, 7, 6, 2, 2, 122,
	123, 7, 69, 2, 2, 123, 124, 7, 28, 2, 2, 124, 125, 5, 40, 21, 2, 125, 126,
	7, 24, 2, 2, 126, 15, 3, 2, 2, 2, 127, 129, 5, 58, 30, 2, 128, 127, 3,
	2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 131, 7, 8, 2,
	2, 131, 133, 7, 69, 2, 2, 132, 134, 5, 32, 17, 2, 133, 132, 3, 2, 2, 2,
	133, 134, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 136, 7, 24, 2, 2, 136,
	17, 3, 2, 2, 2, 137, 139, 5, 58, 30, 2, 138, 137, 3, 2, 2, 2, 138, 139,
	3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 141, 7, 9, 2, 2, 141, 142, 7, 69,
	2, 2, 142, 143, 7, 14, 2, 2, 143, 144, 7, 70, 2, 2, 144, 145, 7, 28, 2,
	2, 145, 146, 5, 38, 20, 2, 146, 147, 7, 24, 2, 2, 147, 19, 3, 2, 2, 2,
	148, 150, 5, 58, 30, 2, 149, 148, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150,
	152, 3, 2, 2, 2, 151, 153, 5, 34, 18, 2, 152, 151, 3, 2, 2, 2, 152, 153,
	3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 155, 7, 7, 2, 2, 155, 157, 7, 69,
	2, 2, 156, 158, 5, 48, 25, 2, 157, 156, 3, 2, 2, 2, 157, 158, 3, 2, 2,
	2, 158, 160, 3, 2, 2, 2, 159, 161, 5, 32, 17, 2, 160, 159, 3, 2, 2, 2,
	160, 161, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 163, 5, 10, 6, 2, 163,
	21, 3, 2, 2, 2, 164, 166, 5, 58, 30, 2, 165, 164, 3, 2, 2, 2, 165, 166,
	3, 2, 2, 2, 166, 168, 3, 2, 2, 2, 167, 169, 5, 34, 18, 2, 168, 167, 3,
	2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 171, 7, 5, 2,
	2, 171, 173, 7, 69, 2, 2, 172, 174, 5, 48, 25, 2, 173, 172, 3, 2, 2, 2,
	173, 174, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 176, 7, 4, 2, 2, 176,
	178, 7, 69, 2, 2, 177, 179, 5, 24, 13, 2, 178, 177, 3, 2, 2, 2, 178, 179,
	3, 2, 2, 2, 179, 181, 3, 2, 2, 2, 180, 182, 5, 26, 14, 2, 181, 180, 3,
	2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 184, 7, 24, 2,
	2, 184, 23, 3, 2, 2, 2, 185, 186, 7, 11, 2, 2, 186, 187, 7, 69, 2, 2, 187,
	25, 3, 2, 2, 2, 188, 189, 7, 15, 2, 2, 189, 190, 7, 20, 2, 2, 190, 191,
	5, 28, 15, 2, 191, 192, 7, 21, 2, 2, 192, 27, 3, 2, 2, 2, 193, 199, 5,
	30, 16, 2, 194, 195, 5, 30, 16, 2, 195, 196, 5, 62, 32, 2, 196, 197, 5,
	28, 15, 2, 197, 199, 3, 2, 2, 2, 198, 193, 3, 2, 2, 2, 198, 194, 3, 2,
	2, 2, 199, 29, 3, 2, 2, 2, 200, 201, 7, 69, 2, 2, 201, 202, 5, 60, 31,
	2, 202, 203, 5, 56, 29, 2, 203, 31, 3, 2, 2, 2, 204, 206, 7, 10, 2, 2,
	205, 204, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 208, 3, 2, 2, 2, 207,
	209, 5, 52, 27, 2, 208, 207, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 33,
	3, 2, 2, 2, 210, 212, 5, 36, 19, 2, 211, 210, 3, 2, 2, 2, 212, 215, 3,
	2, 2, 2, 213, 211, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 35, 3, 2, 2,
	2, 215, 213, 3, 2, 2, 2, 216, 217, 7, 18, 2, 2, 217, 218, 7, 69, 2, 2,
	218, 233, 7, 19, 2, 2, 219, 220, 7, 18, 2, 2, 220, 221, 7, 69, 2, 2, 221,
	222, 7, 20, 2, 2, 222, 223, 7, 69, 2, 2, 223, 224, 7, 21, 2, 2, 224, 233,
	7, 19, 2, 2, 225, 226, 7, 18, 2, 2, 226, 227, 7, 69, 2, 2, 227, 228, 7,
	20, 2, 2, 228, 229, 5, 56, 29, 2, 229, 230, 7, 21, 2, 2, 230, 231, 7, 19,
	2, 2, 231, 233, 3, 2, 2, 2, 232, 216, 3, 2, 2, 2, 232, 219, 3, 2, 2, 2,
	232, 225, 3, 2, 2, 2, 233, 37, 3, 2, 2, 2, 234, 240, 5, 40, 21, 2, 235,
	236, 5, 40, 21, 2, 236, 237, 7, 25, 2, 2, 237, 238, 5, 38, 20, 2, 238,
	240, 3, 2, 2, 2, 239, 234, 3, 2, 2, 2, 239, 235, 3, 2, 2, 2, 240, 39, 3,
	2, 2, 2, 241, 243, 7, 69, 2, 2, 242, 244, 5, 42, 22, 2, 243, 242, 3, 2,
	2, 2, 243, 244, 3, 2, 2, 2, 244, 252, 3, 2, 2, 2, 245, 246, 5, 50, 26,
	2, 246, 247, 7, 28, 2, 2, 247, 249, 7, 69, 2, 2, 248, 250, 5, 42, 22, 2,
	249, 248, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250, 252, 3, 2, 2, 2, 251,
	241, 3, 2, 2, 2, 251, 245, 3, 2, 2, 2, 252, 41, 3, 2, 2, 2, 253, 254, 7,
	26, 2, 2, 254, 261, 5, 56, 29, 2, 255, 256, 7, 26, 2, 2, 256, 257, 7, 22,
	2, 2, 257, 258, 5, 44, 23, 2, 258, 259, 7, 23, 2, 2, 259, 261, 3, 2, 2,
	2, 260, 253, 3, 2, 2, 2, 260, 255, 3, 2, 2, 2, 261, 43, 3, 2, 2, 2, 262,
	268, 5, 46, 24, 2, 263, 264, 5, 46, 24, 2, 264, 265, 7, 25, 2, 2, 265,
	266, 5, 44, 23, 2, 266, 268, 3, 2, 2, 2, 267, 262, 3, 2, 2, 2, 267, 263,
	3, 2, 2, 2, 268, 45, 3, 2, 2, 2, 269, 270, 7, 69, 2, 2, 270, 271, 7, 26,
	2, 2, 271, 272, 5, 56, 29, 2, 272, 47, 3, 2, 2, 2, 273, 274, 7, 18, 2,
	2, 274, 275, 5, 56, 29, 2, 275, 276, 7, 25, 2, 2, 276, 277, 5, 56, 29,
	2, 277, 278, 7, 19, 2, 2, 278, 286, 3, 2, 2, 2, 279, 280, 7, 18, 2, 2,
	280, 281, 5, 56, 29, 2, 281, 282, 7, 25, 2, 2, 282, 283, 7, 68, 2, 2, 283,
	284, 7, 19, 2, 2, 284, 286, 3, 2, 2, 2, 285, 273, 3, 2, 2, 2, 285, 279,
	3, 2, 2, 2, 286, 49, 3, 2, 2, 2, 287, 292, 7, 69, 2, 2, 288, 289, 7, 69,
	2, 2, 289, 290, 7, 30, 2, 2, 290, 292, 5, 50, 26, 2, 291, 287, 3, 2, 2,
	2, 291, 288, 3, 2, 2, 2, 292, 51, 3, 2, 2, 2, 293, 298, 7, 69, 2, 2, 294,
	295, 7, 69, 2, 2, 295, 296, 7, 25, 2, 2, 296, 298, 5, 52, 27, 2, 297, 293,
	3, 2, 2, 2, 297, 294, 3, 2, 2, 2, 298, 53, 3, 2, 2, 2, 299, 300, 9, 2,
	2, 2, 300, 55, 3, 2, 2, 2, 301, 302, 9, 3, 2, 2, 302, 57, 3, 2, 2, 2, 303,
	305, 9, 4, 2, 2, 304, 303, 3, 2, 2, 2, 305, 308, 3, 2, 2, 2, 306, 304,
	3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307, 59, 3, 2, 2, 2, 308, 306, 3, 2,
	2, 2, 309, 310, 9, 5, 2, 2, 310, 61, 3, 2, 2, 2, 311, 312, 9, 6, 2, 2,
	312, 63, 3, 2, 2, 2, 39, 65, 68, 76, 89, 92, 97, 110, 112, 116, 119, 128,
	133, 138, 149, 152, 157, 160, 165, 168, 173, 178, 181, 198, 205, 208, 213,
	232, 239, 243, 249, 251, 260, 267, 285, 291, 297, 306,
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
	"program", "statments", "importStatment", "businessObjectDefinition", "block",
	"itemList", "element", "boAction", "message", "node", "association", "associationUsingDefinition",
	"valuationDefinition", "valutaionExpressionList", "valutaionExpression",
	"messageRaiseDefinition", "annotationList", "annotation", "typeList", "type",
	"typeDefaultValue", "valueAssignList", "valueAssign", "multiplicity", "memberExpression",
	"identifierList", "keyword", "literal", "comments", "compareOperator",
	"logicOperator",
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
	BODLParserRULE_statments                  = 1
	BODLParserRULE_importStatment             = 2
	BODLParserRULE_businessObjectDefinition   = 3
	BODLParserRULE_block                      = 4
	BODLParserRULE_itemList                   = 5
	BODLParserRULE_element                    = 6
	BODLParserRULE_boAction                   = 7
	BODLParserRULE_message                    = 8
	BODLParserRULE_node                       = 9
	BODLParserRULE_association                = 10
	BODLParserRULE_associationUsingDefinition = 11
	BODLParserRULE_valuationDefinition        = 12
	BODLParserRULE_valutaionExpressionList    = 13
	BODLParserRULE_valutaionExpression        = 14
	BODLParserRULE_messageRaiseDefinition     = 15
	BODLParserRULE_annotationList             = 16
	BODLParserRULE_annotation                 = 17
	BODLParserRULE_typeList                   = 18
	BODLParserRULE_type                       = 19
	BODLParserRULE_typeDefaultValue           = 20
	BODLParserRULE_valueAssignList            = 21
	BODLParserRULE_valueAssign                = 22
	BODLParserRULE_multiplicity               = 23
	BODLParserRULE_memberExpression           = 24
	BODLParserRULE_identifierList             = 25
	BODLParserRULE_keyword                    = 26
	BODLParserRULE_literal                    = 27
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

func (s *ProgramContext) BusinessObjectDefinition() IBusinessObjectDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBusinessObjectDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBusinessObjectDefinitionContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(BODLParserEOF, 0)
}

func (s *ProgramContext) Comments() ICommentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentsContext)
}

func (s *ProgramContext) Statments() IStatmentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatmentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatmentsContext)
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
	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(62)
			p.Comments()
		}

	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(65)
			p.Statments()
		}

	}
	{
		p.SetState(68)
		p.BusinessObjectDefinition()
	}
	{
		p.SetState(69)
		p.Match(BODLParserEOF)
	}

	return localctx
}

// IStatmentsContext is an interface to support dynamic dispatch.
type IStatmentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatmentsContext differentiates from other interfaces.
	IsStatmentsContext()
}

type StatmentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatmentsContext() *StatmentsContext {
	var p = new(StatmentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_statments
	return p
}

func (*StatmentsContext) IsStatmentsContext() {}

func NewStatmentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatmentsContext {
	var p = new(StatmentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_statments

	return p
}

func (s *StatmentsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatmentsContext) AllImportStatment() []IImportStatmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportStatmentContext)(nil)).Elem())
	var tst = make([]IImportStatmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportStatmentContext)
		}
	}

	return tst
}

func (s *StatmentsContext) ImportStatment(i int) IImportStatmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportStatmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportStatmentContext)
}

func (s *StatmentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatmentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatmentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterStatments(s)
	}
}

func (s *StatmentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitStatments(s)
	}
}

func (p *BODLParser) Statments() (localctx IStatmentsContext) {
	localctx = NewStatmentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BODLParserRULE_statments)
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
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BODLParserIMPORT {
		{
			p.SetState(71)
			p.ImportStatment()
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 4, BODLParserRULE_importStatment)

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

	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Match(BODLParserIMPORT)
		}
		{
			p.SetState(78)
			p.MemberExpression()
		}
		{
			p.SetState(79)
			p.Match(BODLParserSemiColon)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Match(BODLParserIMPORT)
		}
		{
			p.SetState(82)
			p.MemberExpression()
		}
		{
			p.SetState(83)
			p.Match(BODLParserAS)
		}
		{
			p.SetState(84)
			p.Match(BODLParserIdentifier)
		}
		{
			p.SetState(85)
			p.Match(BODLParserSemiColon)
		}

	}

	return localctx
}

// IBusinessObjectDefinitionContext is an interface to support dynamic dispatch.
type IBusinessObjectDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBusinessObjectDefinitionContext differentiates from other interfaces.
	IsBusinessObjectDefinitionContext()
}

type BusinessObjectDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBusinessObjectDefinitionContext() *BusinessObjectDefinitionContext {
	var p = new(BusinessObjectDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_businessObjectDefinition
	return p
}

func (*BusinessObjectDefinitionContext) IsBusinessObjectDefinitionContext() {}

func NewBusinessObjectDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BusinessObjectDefinitionContext {
	var p = new(BusinessObjectDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_businessObjectDefinition

	return p
}

func (s *BusinessObjectDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *BusinessObjectDefinitionContext) BUSINESSOBJECT() antlr.TerminalNode {
	return s.GetToken(BODLParserBUSINESSOBJECT, 0)
}

func (s *BusinessObjectDefinitionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BODLParserIdentifier, 0)
}

func (s *BusinessObjectDefinitionContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BusinessObjectDefinitionContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
}

func (s *BusinessObjectDefinitionContext) MessageRaiseDefinition() IMessageRaiseDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageRaiseDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageRaiseDefinitionContext)
}

func (s *BusinessObjectDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BusinessObjectDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BusinessObjectDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterBusinessObjectDefinition(s)
	}
}

func (s *BusinessObjectDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitBusinessObjectDefinition(s)
	}
}

func (p *BODLParser) BusinessObjectDefinition() (localctx IBusinessObjectDefinitionContext) {
	localctx = NewBusinessObjectDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BODLParserRULE_businessObjectDefinition)

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
	p.SetState(90)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(89)
			p.AnnotationList()
		}

	}
	{
		p.SetState(92)
		p.Match(BODLParserBUSINESSOBJECT)
	}
	{
		p.SetState(93)
		p.Match(BODLParserIdentifier)
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(94)
			p.MessageRaiseDefinition()
		}

	}
	{
		p.SetState(97)
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
	p.EnterRule(localctx, 8, BODLParserRULE_block)

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
		p.SetState(99)
		p.Match(BODLParserOpenBrace)
	}
	{
		p.SetState(100)
		p.ItemList()
	}
	{
		p.SetState(101)
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
	p.EnterRule(localctx, 10, BODLParserRULE_itemList)
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
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BODLParserASSOCIATION)|(1<<BODLParserELEMENT)|(1<<BODLParserNODE)|(1<<BODLParserACTION)|(1<<BODLParserMESSAGE)|(1<<BODLParserMultiLineComment)|(1<<BODLParserSingleLineComment)|(1<<BODLParserOpenBracket))) != 0 {
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(103)
				p.Element()
			}

		case 2:
			{
				p.SetState(104)
				p.Message()
			}

		case 3:
			{
				p.SetState(105)
				p.Node()
			}

		case 4:
			{
				p.SetState(106)
				p.BoAction()
			}

		case 5:
			{
				p.SetState(107)
				p.Association()
			}

		}

		p.SetState(112)
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
	p.EnterRule(localctx, 12, BODLParserRULE_element)

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
	p.SetState(114)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(113)
			p.Comments()
		}

	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(116)
			p.AnnotationList()
		}

	}
	{
		p.SetState(119)
		p.Match(BODLParserELEMENT)
	}
	{
		p.SetState(120)
		p.Match(BODLParserIdentifier)
	}
	{
		p.SetState(121)
		p.Match(BODLParserColon)
	}
	{
		p.SetState(122)
		p.Type()
	}
	{
		p.SetState(123)
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

func (s *BoActionContext) MessageRaiseDefinition() IMessageRaiseDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageRaiseDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageRaiseDefinitionContext)
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
	p.EnterRule(localctx, 14, BODLParserRULE_boAction)

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
	p.SetState(126)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(125)
			p.Comments()
		}

	}
	{
		p.SetState(128)
		p.Match(BODLParserACTION)
	}
	{
		p.SetState(129)
		p.Match(BODLParserIdentifier)
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(130)
			p.MessageRaiseDefinition()
		}

	}
	{
		p.SetState(133)
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
	p.EnterRule(localctx, 16, BODLParserRULE_message)

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
	p.SetState(136)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(135)
			p.Comments()
		}

	}
	{
		p.SetState(138)
		p.Match(BODLParserMESSAGE)
	}
	{
		p.SetState(139)
		p.Match(BODLParserIdentifier)
	}
	{
		p.SetState(140)
		p.Match(BODLParserTEXT)
	}
	{
		p.SetState(141)
		p.Match(BODLParserStringLiteral)
	}
	{
		p.SetState(142)
		p.Match(BODLParserColon)
	}
	{
		p.SetState(143)
		p.TypeList()
	}
	{
		p.SetState(144)
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

func (s *NodeContext) MessageRaiseDefinition() IMessageRaiseDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageRaiseDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageRaiseDefinitionContext)
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
	p.EnterRule(localctx, 18, BODLParserRULE_node)
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
	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(146)
			p.Comments()
		}

	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(149)
			p.AnnotationList()
		}

	}
	{
		p.SetState(152)
		p.Match(BODLParserNODE)
	}
	{
		p.SetState(153)
		p.Match(BODLParserIdentifier)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserOpenBracket {
		{
			p.SetState(154)
			p.Multiplicity()
		}

	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(157)
			p.MessageRaiseDefinition()
		}

	}
	{
		p.SetState(160)
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
	p.EnterRule(localctx, 20, BODLParserRULE_association)
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
	p.SetState(163)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(162)
			p.Comments()
		}

	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(165)
			p.AnnotationList()
		}

	}
	{
		p.SetState(168)
		p.Match(BODLParserASSOCIATION)
	}
	{
		p.SetState(169)
		p.Match(BODLParserIdentifier)
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserOpenBracket {
		{
			p.SetState(170)
			p.Multiplicity()
		}

	}
	{
		p.SetState(173)
		p.Match(BODLParserTO)
	}
	{
		p.SetState(174)
		p.Match(BODLParserIdentifier)
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserUSING {
		{
			p.SetState(175)
			p.AssociationUsingDefinition()
		}

	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserVALUATION {
		{
			p.SetState(178)
			p.ValuationDefinition()
		}

	}
	{
		p.SetState(181)
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
	p.EnterRule(localctx, 22, BODLParserRULE_associationUsingDefinition)

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
		p.SetState(183)
		p.Match(BODLParserUSING)
	}
	{
		p.SetState(184)
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
	p.EnterRule(localctx, 24, BODLParserRULE_valuationDefinition)

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
		p.SetState(186)
		p.Match(BODLParserVALUATION)
	}
	{
		p.SetState(187)
		p.Match(BODLParserOpenParen)
	}
	{
		p.SetState(188)
		p.ValutaionExpressionList()
	}
	{
		p.SetState(189)
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
	p.EnterRule(localctx, 26, BODLParserRULE_valutaionExpressionList)

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

	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(191)
			p.ValutaionExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(192)
			p.ValutaionExpression()
		}
		{
			p.SetState(193)
			p.LogicOperator()
		}
		{
			p.SetState(194)
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
	p.EnterRule(localctx, 28, BODLParserRULE_valutaionExpression)

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
		p.SetState(198)
		p.Match(BODLParserIdentifier)
	}
	{
		p.SetState(199)
		p.CompareOperator()
	}
	{
		p.SetState(200)
		p.Literal()
	}

	return localctx
}

// IMessageRaiseDefinitionContext is an interface to support dynamic dispatch.
type IMessageRaiseDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageRaiseDefinitionContext differentiates from other interfaces.
	IsMessageRaiseDefinitionContext()
}

type MessageRaiseDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageRaiseDefinitionContext() *MessageRaiseDefinitionContext {
	var p = new(MessageRaiseDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BODLParserRULE_messageRaiseDefinition
	return p
}

func (*MessageRaiseDefinitionContext) IsMessageRaiseDefinitionContext() {}

func NewMessageRaiseDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageRaiseDefinitionContext {
	var p = new(MessageRaiseDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BODLParserRULE_messageRaiseDefinition

	return p
}

func (s *MessageRaiseDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageRaiseDefinitionContext) RAISES() antlr.TerminalNode {
	return s.GetToken(BODLParserRAISES, 0)
}

func (s *MessageRaiseDefinitionContext) IdentifierList() IIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *MessageRaiseDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageRaiseDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageRaiseDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.EnterMessageRaiseDefinition(s)
	}
}

func (s *MessageRaiseDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BODLParserListener); ok {
		listenerT.ExitMessageRaiseDefinition(s)
	}
}

func (p *BODLParser) MessageRaiseDefinition() (localctx IMessageRaiseDefinitionContext) {
	localctx = NewMessageRaiseDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BODLParserRULE_messageRaiseDefinition)
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
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserRAISES {
		{
			p.SetState(202)
			p.Match(BODLParserRAISES)
		}

	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BODLParserIdentifier {
		{
			p.SetState(205)
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
	p.EnterRule(localctx, 32, BODLParserRULE_annotationList)
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
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BODLParserOpenBracket {
		{
			p.SetState(208)
			p.Annotation()
		}

		p.SetState(213)
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
	p.EnterRule(localctx, 34, BODLParserRULE_annotation)

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

	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.Match(BODLParserOpenBracket)
		}
		{
			p.SetState(215)
			p.Match(BODLParserIdentifier)
		}
		{
			p.SetState(216)
			p.Match(BODLParserCloseBracket)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Match(BODLParserOpenBracket)
		}
		{
			p.SetState(218)
			p.Match(BODLParserIdentifier)
		}
		{
			p.SetState(219)
			p.Match(BODLParserOpenParen)
		}
		{
			p.SetState(220)
			p.Match(BODLParserIdentifier)
		}
		{
			p.SetState(221)
			p.Match(BODLParserCloseParen)
		}
		{
			p.SetState(222)
			p.Match(BODLParserCloseBracket)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(223)
			p.Match(BODLParserOpenBracket)
		}
		{
			p.SetState(224)
			p.Match(BODLParserIdentifier)
		}
		{
			p.SetState(225)
			p.Match(BODLParserOpenParen)
		}
		{
			p.SetState(226)
			p.Literal()
		}
		{
			p.SetState(227)
			p.Match(BODLParserCloseParen)
		}
		{
			p.SetState(228)
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
	p.EnterRule(localctx, 36, BODLParserRULE_typeList)

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

	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.Type()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)
			p.Type()
		}
		{
			p.SetState(234)
			p.Match(BODLParserComma)
		}
		{
			p.SetState(235)
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
	p.EnterRule(localctx, 38, BODLParserRULE_type)
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

	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(239)
			p.Match(BODLParserIdentifier)
		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BODLParserAssign {
			{
				p.SetState(240)
				p.TypeDefaultValue()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(243)
			p.MemberExpression()
		}
		{
			p.SetState(244)
			p.Match(BODLParserColon)
		}
		{
			p.SetState(245)
			p.Match(BODLParserIdentifier)
		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BODLParserAssign {
			{
				p.SetState(246)
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
	p.EnterRule(localctx, 40, BODLParserRULE_typeDefaultValue)

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

	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(251)
			p.Match(BODLParserAssign)
		}
		{
			p.SetState(252)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
			p.Match(BODLParserAssign)
		}
		{
			p.SetState(254)
			p.Match(BODLParserOpenBrace)
		}
		{
			p.SetState(255)
			p.ValueAssignList()
		}
		{
			p.SetState(256)
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
	p.EnterRule(localctx, 42, BODLParserRULE_valueAssignList)

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

	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(260)
			p.ValueAssign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
			p.ValueAssign()
		}
		{
			p.SetState(262)
			p.Match(BODLParserComma)
		}
		{
			p.SetState(263)
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
	p.EnterRule(localctx, 44, BODLParserRULE_valueAssign)

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
		p.SetState(267)
		p.Match(BODLParserIdentifier)
	}
	{
		p.SetState(268)
		p.Match(BODLParserAssign)
	}
	{
		p.SetState(269)
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
	p.EnterRule(localctx, 46, BODLParserRULE_multiplicity)

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

	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.Match(BODLParserOpenBracket)
		}
		{
			p.SetState(272)
			p.Literal()
		}
		{
			p.SetState(273)
			p.Match(BODLParserComma)
		}
		{
			p.SetState(274)
			p.Literal()
		}
		{
			p.SetState(275)
			p.Match(BODLParserCloseBracket)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)
			p.Match(BODLParserOpenBracket)
		}
		{
			p.SetState(278)
			p.Literal()
		}
		{
			p.SetState(279)
			p.Match(BODLParserComma)
		}
		{
			p.SetState(280)
			p.Match(BODLParserN)
		}
		{
			p.SetState(281)
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
	p.EnterRule(localctx, 48, BODLParserRULE_memberExpression)

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

	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
			p.Match(BODLParserIdentifier)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(286)
			p.Match(BODLParserIdentifier)
		}
		{
			p.SetState(287)
			p.Match(BODLParserDot)
		}
		{
			p.SetState(288)
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
	p.EnterRule(localctx, 50, BODLParserRULE_identifierList)

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

	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(291)
			p.Match(BODLParserIdentifier)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(292)
			p.Match(BODLParserIdentifier)
		}
		{
			p.SetState(293)
			p.Match(BODLParserComma)
		}
		{
			p.SetState(294)
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
	p.EnterRule(localctx, 52, BODLParserRULE_keyword)
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
		p.SetState(297)
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
	p.EnterRule(localctx, 54, BODLParserRULE_literal)
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
		p.SetState(299)
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
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BODLParserMultiLineComment || _la == BODLParserSingleLineComment {
		{
			p.SetState(301)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BODLParserMultiLineComment || _la == BODLParserSingleLineComment) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(306)
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
		p.SetState(307)
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
		p.SetState(309)
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
