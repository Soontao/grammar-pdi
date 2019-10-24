// Generated from ./BODLParser.g4 by ANTLR 4.7.2
// jshint ignore: start
var antlr4 = require("antlr4/index");
var BODLParserListener = require("./BODLParserListener").BODLParserListener;
var BODLParserVisitor = require("./BODLParserVisitor").BODLParserVisitor;

var grammarFileName = "BODLParser.g4";


var serializedATN = ["\u0003\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964",
  "\u0003;\u0161\u0004\u0002\t\u0002\u0004\u0003\t\u0003\u0004\u0004\t",
  "\u0004\u0004\u0005\t\u0005\u0004\u0006\t\u0006\u0004\u0007\t\u0007\u0004",
  "\b\t\b\u0004\t\t\t\u0004\n\t\n\u0004\u000b\t\u000b\u0004\f\t\f\u0004",
  "\r\t\r\u0004\u000e\t\u000e\u0004\u000f\t\u000f\u0004\u0010\t\u0010\u0004",
  "\u0011\t\u0011\u0004\u0012\t\u0012\u0004\u0013\t\u0013\u0004\u0014\t",
  "\u0014\u0004\u0015\t\u0015\u0004\u0016\t\u0016\u0004\u0017\t\u0017\u0004",
  "\u0018\t\u0018\u0004\u0019\t\u0019\u0004\u001a\t\u001a\u0004\u001b\t",
  "\u001b\u0004\u001c\t\u001c\u0004\u001d\t\u001d\u0004\u001e\t\u001e\u0004",
  "\u001f\t\u001f\u0004 \t \u0004!\t!\u0004\"\t\"\u0003\u0002\u0003\u0002",
  "\u0003\u0002\u0003\u0002\u0003\u0002\u0005\u0002J\n\u0002\u0003\u0003",
  "\u0007\u0003M\n\u0003\f\u0003\u000e\u0003P\u000b\u0003\u0003\u0004\u0005",
  "\u0004S\n\u0004\u0003\u0004\u0003\u0004\u0003\u0004\u0003\u0004\u0003",
  "\u0004\u0005\u0004Z\n\u0004\u0003\u0004\u0003\u0004\u0003\u0004\u0003",
  "\u0004\u0003\u0004\u0003\u0004\u0005\u0004b\n\u0004\u0003\u0005\u0007",
  "\u0005e\n\u0005\f\u0005\u000e\u0005h\u000b\u0005\u0003\u0006\u0005\u0006",
  "k\n\u0006\u0003\u0006\u0005\u0006n\n\u0006\u0003\u0006\u0003\u0006\u0003",
  "\u0006\u0005\u0006s\n\u0006\u0003\u0006\u0005\u0006v\n\u0006\u0003\u0006",
  "\u0003\u0006\u0003\u0007\u0003\u0007\u0003\u0007\u0003\u0007\u0003\u0007",
  "\u0003\b\u0003\b\u0003\b\u0003\b\u0003\b\u0007\b\u0084\n\b\f\b\u000e",
  "\b\u0087\u000b\b\u0003\t\u0005\t\u008a\n\t\u0003\t\u0005\t\u008d\n\t",
  "\u0003\t\u0003\t\u0003\t\u0003\t\u0003\t\u0003\t\u0003\n\u0005\n\u0096",
  "\n\n\u0003\n\u0003\n\u0003\n\u0005\n\u009b\n\n\u0003\n\u0003\n\u0003",
  "\u000b\u0005\u000b\u00a0\n\u000b\u0003\u000b\u0003\u000b\u0003\u000b",
  "\u0003\u000b\u0003\u000b\u0003\u000b\u0003\u000b\u0003\u000b\u0003\f",
  "\u0005\f\u00ab\n\f\u0003\f\u0005\f\u00ae\n\f\u0003\f\u0003\f\u0003\f",
  "\u0005\f\u00b3\n\f\u0003\f\u0005\f\u00b6\n\f\u0003\f\u0005\f\u00b9\n",
  "\f\u0003\r\u0005\r\u00bc\n\r\u0003\r\u0005\r\u00bf\n\r\u0003\r\u0003",
  "\r\u0003\r\u0005\r\u00c4\n\r\u0003\r\u0003\r\u0003\r\u0005\r\u00c9\n",
  "\r\u0003\r\u0005\r\u00cc\n\r\u0003\r\u0003\r\u0003\u000e\u0003\u000e",
  "\u0003\u000e\u0003\u000f\u0003\u000f\u0003\u000f\u0003\u000f\u0003\u000f",
  "\u0003\u0010\u0003\u0010\u0003\u0010\u0003\u0010\u0003\u0010\u0005\u0010",
  "\u00dd\n\u0010\u0003\u0011\u0003\u0011\u0003\u0011\u0003\u0011\u0003",
  "\u0012\u0005\u0012\u00e4\n\u0012\u0003\u0012\u0005\u0012\u00e7\n\u0012",
  "\u0003\u0013\u0007\u0013\u00ea\n\u0013\f\u0013\u000e\u0013\u00ed\u000b",
  "\u0013\u0003\u0014\u0005\u0014\u00f0\n\u0014\u0003\u0014\u0003\u0014",
  "\u0003\u0014\u0003\u0014\u0003\u0014\u0005\u0014\u00f7\n\u0014\u0003",
  "\u0014\u0003\u0014\u0003\u0014\u0003\u0014\u0003\u0014\u0003\u0014\u0003",
  "\u0014\u0003\u0014\u0005\u0014\u0101\n\u0014\u0003\u0014\u0003\u0014",
  "\u0003\u0014\u0003\u0014\u0003\u0014\u0003\u0014\u0003\u0014\u0005\u0014",
  "\u010a\n\u0014\u0003\u0015\u0003\u0015\u0003\u0015\u0003\u0015\u0003",
  "\u0015\u0005\u0015\u0111\n\u0015\u0003\u0016\u0003\u0016\u0005\u0016",
  "\u0115\n\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0005",
  "\u0016\u011b\n\u0016\u0005\u0016\u011d\n\u0016\u0003\u0017\u0003\u0017",
  "\u0003\u0017\u0003\u0017\u0003\u0017\u0003\u0017\u0003\u0017\u0005\u0017",
  "\u0126\n\u0017\u0003\u0018\u0003\u0018\u0003\u0018\u0003\u0018\u0003",
  "\u0018\u0005\u0018\u012d\n\u0018\u0003\u0019\u0003\u0019\u0003\u0019",
  "\u0003\u0019\u0003\u001a\u0003\u001a\u0003\u001a\u0003\u001a\u0003\u001a",
  "\u0003\u001a\u0003\u001a\u0003\u001a\u0003\u001a\u0003\u001a\u0003\u001a",
  "\u0003\u001a\u0005\u001a\u013f\n\u001a\u0003\u001b\u0003\u001b\u0003",
  "\u001b\u0007\u001b\u0144\n\u001b\f\u001b\u000e\u001b\u0147\u000b\u001b",
  "\u0003\u001c\u0003\u001c\u0003\u001c\u0007\u001c\u014c\n\u001c\f\u001c",
  "\u000e\u001c\u014f\u000b\u001c\u0003\u001d\u0003\u001d\u0003\u001e\u0003",
  "\u001e\u0003\u001f\u0007\u001f\u0156\n\u001f\f\u001f\u000e\u001f\u0159",
  "\u000b\u001f\u0003 \u0003 \u0003!\u0003!\u0003\"\u0003\"\u0003\"\u0002",
  "\u0002#\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018",
  "\u001a\u001c\u001e \"$&(*,.02468:<>@B\u0002\u0007\u0003\u0002\u0003",
  "\u000e\u0003\u000235\u0003\u0002\u0010\u0011\u0003\u0002(-\u0003\u0002",
  "./\u0002\u0170\u0002I\u0003\u0002\u0002\u0002\u0004N\u0003\u0002\u0002",
  "\u0002\u0006a\u0003\u0002\u0002\u0002\bf\u0003\u0002\u0002\u0002\nj",
  "\u0003\u0002\u0002\u0002\fy\u0003\u0002\u0002\u0002\u000e\u0085\u0003",
  "\u0002\u0002\u0002\u0010\u0089\u0003\u0002\u0002\u0002\u0012\u0095\u0003",
  "\u0002\u0002\u0002\u0014\u009f\u0003\u0002\u0002\u0002\u0016\u00aa\u0003",
  "\u0002\u0002\u0002\u0018\u00bb\u0003\u0002\u0002\u0002\u001a\u00cf\u0003",
  "\u0002\u0002\u0002\u001c\u00d2\u0003\u0002\u0002\u0002\u001e\u00dc\u0003",
  "\u0002\u0002\u0002 \u00de\u0003\u0002\u0002\u0002\"\u00e3\u0003\u0002",
  "\u0002\u0002$\u00eb\u0003\u0002\u0002\u0002&\u0109\u0003\u0002\u0002",
  "\u0002(\u0110\u0003\u0002\u0002\u0002*\u011c\u0003\u0002\u0002\u0002",
  ",\u0125\u0003\u0002\u0002\u0002.\u012c\u0003\u0002\u0002\u00020\u012e",
  "\u0003\u0002\u0002\u00022\u013e\u0003\u0002\u0002\u00024\u0140\u0003",
  "\u0002\u0002\u00026\u0148\u0003\u0002\u0002\u00028\u0150\u0003\u0002",
  "\u0002\u0002:\u0152\u0003\u0002\u0002\u0002<\u0157\u0003\u0002\u0002",
  "\u0002>\u015a\u0003\u0002\u0002\u0002@\u015c\u0003\u0002\u0002\u0002",
  "B\u015e\u0003\u0002\u0002\u0002DJ\u0007\u0002\u0002\u0003EF\u0005\u0004",
  "\u0003\u0002FG\u0005\b\u0005\u0002GH\u0007\u0002\u0002\u0003HJ\u0003",
  "\u0002\u0002\u0002ID\u0003\u0002\u0002\u0002IE\u0003\u0002\u0002\u0002",
  "J\u0003\u0003\u0002\u0002\u0002KM\u0005\u0006\u0004\u0002LK\u0003\u0002",
  "\u0002\u0002MP\u0003\u0002\u0002\u0002NL\u0003\u0002\u0002\u0002NO\u0003",
  "\u0002\u0002\u0002O\u0005\u0003\u0002\u0002\u0002PN\u0003\u0002\u0002",
  "\u0002QS\u0005<\u001f\u0002RQ\u0003\u0002\u0002\u0002RS\u0003\u0002",
  "\u0002\u0002ST\u0003\u0002\u0002\u0002TU\u0007\f\u0002\u0002UV\u0005",
  "4\u001b\u0002VW\u0007\u0019\u0002\u0002Wb\u0003\u0002\u0002\u0002XZ",
  "\u0005<\u001f\u0002YX\u0003\u0002\u0002\u0002YZ\u0003\u0002\u0002\u0002",
  "Z[\u0003\u0002\u0002\u0002[\\\u0007\f\u0002\u0002\\]\u00054\u001b\u0002",
  "]^\u0007\r\u0002\u0002^_\u0005B\"\u0002_`\u0007\u0019\u0002\u0002`b",
  "\u0003\u0002\u0002\u0002aR\u0003\u0002\u0002\u0002aY\u0003\u0002\u0002",
  "\u0002b\u0007\u0003\u0002\u0002\u0002ce\u0005\n\u0006\u0002dc\u0003",
  "\u0002\u0002\u0002eh\u0003\u0002\u0002\u0002fd\u0003\u0002\u0002\u0002",
  "fg\u0003\u0002\u0002\u0002g\t\u0003\u0002\u0002\u0002hf\u0003\u0002",
  "\u0002\u0002ik\u0005<\u001f\u0002ji\u0003\u0002\u0002\u0002jk\u0003",
  "\u0002\u0002\u0002km\u0003\u0002\u0002\u0002ln\u0005$\u0013\u0002ml",
  "\u0003\u0002\u0002\u0002mn\u0003\u0002\u0002\u0002no\u0003\u0002\u0002",
  "\u0002or\u0007\u0003\u0002\u0002ps\u0005B\"\u0002qs\u0005*\u0016\u0002",
  "rp\u0003\u0002\u0002\u0002rq\u0003\u0002\u0002\u0002su\u0003\u0002\u0002",
  "\u0002tv\u0005\"\u0012\u0002ut\u0003\u0002\u0002\u0002uv\u0003\u0002",
  "\u0002\u0002vw\u0003\u0002\u0002\u0002wx\u0005\f\u0007\u0002x\u000b",
  "\u0003\u0002\u0002\u0002yz\u0007\u0017\u0002\u0002z{\u0005\u000e\b\u0002",
  "{|\u0005<\u001f\u0002|}\u0007\u0018\u0002\u0002}\r\u0003\u0002\u0002",
  "\u0002~\u0084\u0005\u0010\t\u0002\u007f\u0084\u0005\u0014\u000b\u0002",
  "\u0080\u0084\u0005\u0016\f\u0002\u0081\u0084\u0005\u0012\n\u0002\u0082",
  "\u0084\u0005\u0018\r\u0002\u0083~\u0003\u0002\u0002\u0002\u0083\u007f",
  "\u0003\u0002\u0002\u0002\u0083\u0080\u0003\u0002\u0002\u0002\u0083\u0081",
  "\u0003\u0002\u0002\u0002\u0083\u0082\u0003\u0002\u0002\u0002\u0084\u0087",
  "\u0003\u0002\u0002\u0002\u0085\u0083\u0003\u0002\u0002\u0002\u0085\u0086",
  "\u0003\u0002\u0002\u0002\u0086\u000f\u0003\u0002\u0002\u0002\u0087\u0085",
  "\u0003\u0002\u0002\u0002\u0088\u008a\u0005<\u001f\u0002\u0089\u0088",
  "\u0003\u0002\u0002\u0002\u0089\u008a\u0003\u0002\u0002\u0002\u008a\u008c",
  "\u0003\u0002\u0002\u0002\u008b\u008d\u0005$\u0013\u0002\u008c\u008b",
  "\u0003\u0002\u0002\u0002\u008c\u008d\u0003\u0002\u0002\u0002\u008d\u008e",
  "\u0003\u0002\u0002\u0002\u008e\u008f\u0007\u0006\u0002\u0002\u008f\u0090",
  "\u0005B\"\u0002\u0090\u0091\u0007\u001d\u0002\u0002\u0091\u0092\u0005",
  "*\u0016\u0002\u0092\u0093\u0007\u0019\u0002\u0002\u0093\u0011\u0003",
  "\u0002\u0002\u0002\u0094\u0096\u0005<\u001f\u0002\u0095\u0094\u0003",
  "\u0002\u0002\u0002\u0095\u0096\u0003\u0002\u0002\u0002\u0096\u0097\u0003",
  "\u0002\u0002\u0002\u0097\u0098\u0007\b\u0002\u0002\u0098\u009a\u0005",
  "B\"\u0002\u0099\u009b\u0005\"\u0012\u0002\u009a\u0099\u0003\u0002\u0002",
  "\u0002\u009a\u009b\u0003\u0002\u0002\u0002\u009b\u009c\u0003\u0002\u0002",
  "\u0002\u009c\u009d\u0007\u0019\u0002\u0002\u009d\u0013\u0003\u0002\u0002",
  "\u0002\u009e\u00a0\u0005<\u001f\u0002\u009f\u009e\u0003\u0002\u0002",
  "\u0002\u009f\u00a0\u0003\u0002\u0002\u0002\u00a0\u00a1\u0003\u0002\u0002",
  "\u0002\u00a1\u00a2\u0007\t\u0002\u0002\u00a2\u00a3\u0005B\"\u0002\u00a3",
  "\u00a4\u0007\u000e\u0002\u0002\u00a4\u00a5\u00073\u0002\u0002\u00a5",
  "\u00a6\u0007\u001d\u0002\u0002\u00a6\u00a7\u0005(\u0015\u0002\u00a7",
  "\u00a8\u0007\u0019\u0002\u0002\u00a8\u0015\u0003\u0002\u0002\u0002\u00a9",
  "\u00ab\u0005<\u001f\u0002\u00aa\u00a9\u0003\u0002\u0002\u0002\u00aa",
  "\u00ab\u0003\u0002\u0002\u0002\u00ab\u00ad\u0003\u0002\u0002\u0002\u00ac",
  "\u00ae\u0005$\u0013\u0002\u00ad\u00ac\u0003\u0002\u0002\u0002\u00ad",
  "\u00ae\u0003\u0002\u0002\u0002\u00ae\u00af\u0003\u0002\u0002\u0002\u00af",
  "\u00b0\u0007\u0007\u0002\u0002\u00b0\u00b2\u0005B\"\u0002\u00b1\u00b3",
  "\u00052\u001a\u0002\u00b2\u00b1\u0003\u0002\u0002\u0002\u00b2\u00b3",
  "\u0003\u0002\u0002\u0002\u00b3\u00b5\u0003\u0002\u0002\u0002\u00b4\u00b6",
  "\u0005\"\u0012\u0002\u00b5\u00b4\u0003\u0002\u0002\u0002\u00b5\u00b6",
  "\u0003\u0002\u0002\u0002\u00b6\u00b8\u0003\u0002\u0002\u0002\u00b7\u00b9",
  "\u0005\f\u0007\u0002\u00b8\u00b7\u0003\u0002\u0002\u0002\u00b8\u00b9",
  "\u0003\u0002\u0002\u0002\u00b9\u0017\u0003\u0002\u0002\u0002\u00ba\u00bc",
  "\u0005<\u001f\u0002\u00bb\u00ba\u0003\u0002\u0002\u0002\u00bb\u00bc",
  "\u0003\u0002\u0002\u0002\u00bc\u00be\u0003\u0002\u0002\u0002\u00bd\u00bf",
  "\u0005$\u0013\u0002\u00be\u00bd\u0003\u0002\u0002\u0002\u00be\u00bf",
  "\u0003\u0002\u0002\u0002\u00bf\u00c0\u0003\u0002\u0002\u0002\u00c0\u00c1",
  "\u0007\u0005\u0002\u0002\u00c1\u00c3\u0005B\"\u0002\u00c2\u00c4\u0005",
  "2\u001a\u0002\u00c3\u00c2\u0003\u0002\u0002\u0002\u00c3\u00c4\u0003",
  "\u0002\u0002\u0002\u00c4\u00c5\u0003\u0002\u0002\u0002\u00c5\u00c6\u0007",
  "\u0004\u0002\u0002\u00c6\u00c8\u0005B\"\u0002\u00c7\u00c9\u0005\u001a",
  "\u000e\u0002\u00c8\u00c7\u0003\u0002\u0002\u0002\u00c8\u00c9\u0003\u0002",
  "\u0002\u0002\u00c9\u00cb\u0003\u0002\u0002\u0002\u00ca\u00cc\u0005\u001c",
  "\u000f\u0002\u00cb\u00ca\u0003\u0002\u0002\u0002\u00cb\u00cc\u0003\u0002",
  "\u0002\u0002\u00cc\u00cd\u0003\u0002\u0002\u0002\u00cd\u00ce\u0007\u0019",
  "\u0002\u0002\u00ce\u0019\u0003\u0002\u0002\u0002\u00cf\u00d0\u0007\u000b",
  "\u0002\u0002\u00d0\u00d1\u0005B\"\u0002\u00d1\u001b\u0003\u0002\u0002",
  "\u0002\u00d2\u00d3\u0007\u000f\u0002\u0002\u00d3\u00d4\u0007\u0015\u0002",
  "\u0002\u00d4\u00d5\u0005\u001e\u0010\u0002\u00d5\u00d6\u0007\u0016\u0002",
  "\u0002\u00d6\u001d\u0003\u0002\u0002\u0002\u00d7\u00dd\u0005 \u0011",
  "\u0002\u00d8\u00d9\u0005 \u0011\u0002\u00d9\u00da\u0005@!\u0002\u00da",
  "\u00db\u0005\u001e\u0010\u0002\u00db\u00dd\u0003\u0002\u0002\u0002\u00dc",
  "\u00d7\u0003\u0002\u0002\u0002\u00dc\u00d8\u0003\u0002\u0002\u0002\u00dd",
  "\u001f\u0003\u0002\u0002\u0002\u00de\u00df\u0005B\"\u0002\u00df\u00e0",
  "\u0005> \u0002\u00e0\u00e1\u0005:\u001e\u0002\u00e1!\u0003\u0002\u0002",
  "\u0002\u00e2\u00e4\u0007\n\u0002\u0002\u00e3\u00e2\u0003\u0002\u0002",
  "\u0002\u00e3\u00e4\u0003\u0002\u0002\u0002\u00e4\u00e6\u0003\u0002\u0002",
  "\u0002\u00e5\u00e7\u00056\u001c\u0002\u00e6\u00e5\u0003\u0002\u0002",
  "\u0002\u00e6\u00e7\u0003\u0002\u0002\u0002\u00e7#\u0003\u0002\u0002",
  "\u0002\u00e8\u00ea\u0005&\u0014\u0002\u00e9\u00e8\u0003\u0002\u0002",
  "\u0002\u00ea\u00ed\u0003\u0002\u0002\u0002\u00eb\u00e9\u0003\u0002\u0002",
  "\u0002\u00eb\u00ec\u0003\u0002\u0002\u0002\u00ec%\u0003\u0002\u0002",
  "\u0002\u00ed\u00eb\u0003\u0002\u0002\u0002\u00ee\u00f0\u0007\u0012\u0002",
  "\u0002\u00ef\u00ee\u0003\u0002\u0002\u0002\u00ef\u00f0\u0003\u0002\u0002",
  "\u0002\u00f0\u00f1\u0003\u0002\u0002\u0002\u00f1\u00f2\u0007\u0013\u0002",
  "\u0002\u00f2\u00f3\u0005B\"\u0002\u00f3\u00f4\u0007\u0014\u0002\u0002",
  "\u00f4\u010a\u0003\u0002\u0002\u0002\u00f5\u00f7\u0007\u0012\u0002\u0002",
  "\u00f6\u00f5\u0003\u0002\u0002\u0002\u00f6\u00f7\u0003\u0002\u0002\u0002",
  "\u00f7\u00f8\u0003\u0002\u0002\u0002\u00f8\u00f9\u0007\u0013\u0002\u0002",
  "\u00f9\u00fa\u0005B\"\u0002\u00fa\u00fb\u0007\u0015\u0002\u0002\u00fb",
  "\u00fc\u0005B\"\u0002\u00fc\u00fd\u0007\u0016\u0002\u0002\u00fd\u00fe",
  "\u0007\u0014\u0002\u0002\u00fe\u010a\u0003\u0002\u0002\u0002\u00ff\u0101",
  "\u0007\u0012\u0002\u0002\u0100\u00ff\u0003\u0002\u0002\u0002\u0100\u0101",
  "\u0003\u0002\u0002\u0002\u0101\u0102\u0003\u0002\u0002\u0002\u0102\u0103",
  "\u0007\u0013\u0002\u0002\u0103\u0104\u0005B\"\u0002\u0104\u0105\u0007",
  "\u0015\u0002\u0002\u0105\u0106\u0005:\u001e\u0002\u0106\u0107\u0007",
  "\u0016\u0002\u0002\u0107\u0108\u0007\u0014\u0002\u0002\u0108\u010a\u0003",
  "\u0002\u0002\u0002\u0109\u00ef\u0003\u0002\u0002\u0002\u0109\u00f6\u0003",
  "\u0002\u0002\u0002\u0109\u0100\u0003\u0002\u0002\u0002\u010a\'\u0003",
  "\u0002\u0002\u0002\u010b\u0111\u0005*\u0016\u0002\u010c\u010d\u0005",
  "*\u0016\u0002\u010d\u010e\u0007\u001a\u0002\u0002\u010e\u010f\u0005",
  "(\u0015\u0002\u010f\u0111\u0003\u0002\u0002\u0002\u0110\u010b\u0003",
  "\u0002\u0002\u0002\u0110\u010c\u0003\u0002\u0002\u0002\u0111)\u0003",
  "\u0002\u0002\u0002\u0112\u0114\u0005B\"\u0002\u0113\u0115\u0005,\u0017",
  "\u0002\u0114\u0113\u0003\u0002\u0002\u0002\u0114\u0115\u0003\u0002\u0002",
  "\u0002\u0115\u011d\u0003\u0002\u0002\u0002\u0116\u0117\u00054\u001b",
  "\u0002\u0117\u0118\u0007\u001d\u0002\u0002\u0118\u011a\u0005B\"\u0002",
  "\u0119\u011b\u0005,\u0017\u0002\u011a\u0119\u0003\u0002\u0002\u0002",
  "\u011a\u011b\u0003\u0002\u0002\u0002\u011b\u011d\u0003\u0002\u0002\u0002",
  "\u011c\u0112\u0003\u0002\u0002\u0002\u011c\u0116\u0003\u0002\u0002\u0002",
  "\u011d+\u0003\u0002\u0002\u0002\u011e\u011f\u0007\u001b\u0002\u0002",
  "\u011f\u0126\u0005:\u001e\u0002\u0120\u0121\u0007\u001b\u0002\u0002",
  "\u0121\u0122\u0007\u0017\u0002\u0002\u0122\u0123\u0005.\u0018\u0002",
  "\u0123\u0124\u0007\u0018\u0002\u0002\u0124\u0126\u0003\u0002\u0002\u0002",
  "\u0125\u011e\u0003\u0002\u0002\u0002\u0125\u0120\u0003\u0002\u0002\u0002",
  "\u0126-\u0003\u0002\u0002\u0002\u0127\u012d\u00050\u0019\u0002\u0128",
  "\u0129\u00050\u0019\u0002\u0129\u012a\u0007\u001a\u0002\u0002\u012a",
  "\u012b\u0005.\u0018\u0002\u012b\u012d\u0003\u0002\u0002\u0002\u012c",
  "\u0127\u0003\u0002\u0002\u0002\u012c\u0128\u0003\u0002\u0002\u0002\u012d",
  "/\u0003\u0002\u0002\u0002\u012e\u012f\u0005B\"\u0002\u012f\u0130\u0007",
  "\u001b\u0002\u0002\u0130\u0131\u0005:\u001e\u0002\u01311\u0003\u0002",
  "\u0002\u0002\u0132\u0133\u0007\u0013\u0002\u0002\u0133\u0134\u0005:",
  "\u001e\u0002\u0134\u0135\u0007\u001a\u0002\u0002\u0135\u0136\u0005:",
  "\u001e\u0002\u0136\u0137\u0007\u0014\u0002\u0002\u0137\u013f\u0003\u0002",
  "\u0002\u0002\u0138\u0139\u0007\u0013\u0002\u0002\u0139\u013a\u0005:",
  "\u001e\u0002\u013a\u013b\u0007\u001a\u0002\u0002\u013b\u013c\u00071",
  "\u0002\u0002\u013c\u013d\u0007\u0014\u0002\u0002\u013d\u013f\u0003\u0002",
  "\u0002\u0002\u013e\u0132\u0003\u0002\u0002\u0002\u013e\u0138\u0003\u0002",
  "\u0002\u0002\u013f3\u0003\u0002\u0002\u0002\u0140\u0145\u0005B\"\u0002",
  "\u0141\u0142\u0007\u001e\u0002\u0002\u0142\u0144\u00054\u001b\u0002",
  "\u0143\u0141\u0003\u0002\u0002\u0002\u0144\u0147\u0003\u0002\u0002\u0002",
  "\u0145\u0143\u0003\u0002\u0002\u0002\u0145\u0146\u0003\u0002\u0002\u0002",
  "\u01465\u0003\u0002\u0002\u0002\u0147\u0145\u0003\u0002\u0002\u0002",
  "\u0148\u014d\u0005B\"\u0002\u0149\u014a\u0007\u001a\u0002\u0002\u014a",
  "\u014c\u0005B\"\u0002\u014b\u0149\u0003\u0002\u0002\u0002\u014c\u014f",
  "\u0003\u0002\u0002\u0002\u014d\u014b\u0003\u0002\u0002\u0002\u014d\u014e",
  "\u0003\u0002\u0002\u0002\u014e7\u0003\u0002\u0002\u0002\u014f\u014d",
  "\u0003\u0002\u0002\u0002\u0150\u0151\t\u0002\u0002\u0002\u01519\u0003",
  "\u0002\u0002\u0002\u0152\u0153\t\u0003\u0002\u0002\u0153;\u0003\u0002",
  "\u0002\u0002\u0154\u0156\t\u0004\u0002\u0002\u0155\u0154\u0003\u0002",
  "\u0002\u0002\u0156\u0159\u0003\u0002\u0002\u0002\u0157\u0155\u0003\u0002",
  "\u0002\u0002\u0157\u0158\u0003\u0002\u0002\u0002\u0158=\u0003\u0002",
  "\u0002\u0002\u0159\u0157\u0003\u0002\u0002\u0002\u015a\u015b\t\u0005",
  "\u0002\u0002\u015b?\u0003\u0002\u0002\u0002\u015c\u015d\t\u0006\u0002",
  "\u0002\u015dA\u0003\u0002\u0002\u0002\u015e\u015f\u00072\u0002\u0002",
  "\u015fC\u0003\u0002\u0002\u0002/INRYafjmru\u0083\u0085\u0089\u008c\u0095",
  "\u009a\u009f\u00aa\u00ad\u00b2\u00b5\u00b8\u00bb\u00be\u00c3\u00c8\u00cb",
  "\u00dc\u00e3\u00e6\u00eb\u00ef\u00f6\u0100\u0109\u0110\u0114\u011a\u011c",
  "\u0125\u012c\u013e\u0145\u014d\u0157"].join("");


var atn = new antlr4.atn.ATNDeserializer().deserialize(serializedATN);

var decisionsToDFA = atn.decisionToState.map( function(ds, index) { return new antlr4.dfa.DFA(ds, index); });

var sharedContextCache = new antlr4.PredictionContextCache();

var literalNames = [ null, "'businessobject'", "'to'", "'association'",
  "'element'", "'node'", "'action'", "'message'", "'raises'",
  "'using'", "'import'", "'as'", "'text'", "'valuation'",
  null, null, "'///'", "'['", "']'", "'('", "')'", "'{'",
  "'}'", "';'", "','", "'='", "'?'", "':'", "'.'", "'+'",
  "'-'", "'~'", "'!'", "'*'", "'/'", "'%'", "'>>'", "'<<'",
  "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&&'",
  "'||'", "'=>'", "'n'" ];

var symbolicNames = [ null, "BUSINESSOBJECT", "TO", "ASSOCIATION", "ELEMENT",
  "NODE", "ACTION", "MESSAGE", "RAISES", "USING", "IMPORT",
  "AS", "TEXT", "VALUATION", "MultiLineComment", "SingleLineComment",
  "CustomAnnotationStart", "OpenBracket", "CloseBracket",
  "OpenParen", "CloseParen", "OpenBrace", "CloseBrace",
  "SemiColon", "Comma", "Assign", "QuestionMark", "Colon",
  "Dot", "Plus", "Minus", "BitNot", "Not", "Multiply",
  "Divide", "Modulus", "RightShiftArithmetic", "LeftShiftArithmetic",
  "LessThan", "MoreThan", "LessThanEquals", "GreaterThanEquals",
  "Equals_", "NotEquals", "And", "Or", "ARROW", "N",
  "Identifier", "StringLiteral", "BooleanLiteral", "DecimalLiteral",
  "HexIntegerLiteral", "OctalIntegerLiteral", "OctalIntegerLiteral2",
  "BinaryIntegerLiteral", "WhiteSpaces", "LineTerminator" ];

var ruleNames = [ "program", "statements", "importStatement", "definitions",
  "definition", "block", "itemList", "element", "boAction",
  "message", "node", "association", "associationUsingDefinition",
  "valuationDefinition", "valutaionExpressionList", "valutaionExpression",
  "raiseMessage", "annotations", "annotation", "typeList",
  "typeDeclaration", "typeDefaultValue", "valueAssignList",
  "valueAssign", "multiplicity", "memberExpression", "identifierList",
  "keyword", "literal", "comments", "compareOperator",
  "logicOperator", "identifier" ];

function BODLParser(input) {
  antlr4.Parser.call(this, input);
  this._interp = new antlr4.atn.ParserATNSimulator(this, atn, decisionsToDFA, sharedContextCache);
  this.ruleNames = ruleNames;
  this.literalNames = literalNames;
  this.symbolicNames = symbolicNames;
  return this;
}

BODLParser.prototype = Object.create(antlr4.Parser.prototype);
BODLParser.prototype.constructor = BODLParser;

Object.defineProperty(BODLParser.prototype, "atn", {
  get : function() {
    return atn;
  }
});

BODLParser.EOF = antlr4.Token.EOF;
BODLParser.BUSINESSOBJECT = 1;
BODLParser.TO = 2;
BODLParser.ASSOCIATION = 3;
BODLParser.ELEMENT = 4;
BODLParser.NODE = 5;
BODLParser.ACTION = 6;
BODLParser.MESSAGE = 7;
BODLParser.RAISES = 8;
BODLParser.USING = 9;
BODLParser.IMPORT = 10;
BODLParser.AS = 11;
BODLParser.TEXT = 12;
BODLParser.VALUATION = 13;
BODLParser.MultiLineComment = 14;
BODLParser.SingleLineComment = 15;
BODLParser.CustomAnnotationStart = 16;
BODLParser.OpenBracket = 17;
BODLParser.CloseBracket = 18;
BODLParser.OpenParen = 19;
BODLParser.CloseParen = 20;
BODLParser.OpenBrace = 21;
BODLParser.CloseBrace = 22;
BODLParser.SemiColon = 23;
BODLParser.Comma = 24;
BODLParser.Assign = 25;
BODLParser.QuestionMark = 26;
BODLParser.Colon = 27;
BODLParser.Dot = 28;
BODLParser.Plus = 29;
BODLParser.Minus = 30;
BODLParser.BitNot = 31;
BODLParser.Not = 32;
BODLParser.Multiply = 33;
BODLParser.Divide = 34;
BODLParser.Modulus = 35;
BODLParser.RightShiftArithmetic = 36;
BODLParser.LeftShiftArithmetic = 37;
BODLParser.LessThan = 38;
BODLParser.MoreThan = 39;
BODLParser.LessThanEquals = 40;
BODLParser.GreaterThanEquals = 41;
BODLParser.Equals_ = 42;
BODLParser.NotEquals = 43;
BODLParser.And = 44;
BODLParser.Or = 45;
BODLParser.ARROW = 46;
BODLParser.N = 47;
BODLParser.Identifier = 48;
BODLParser.StringLiteral = 49;
BODLParser.BooleanLiteral = 50;
BODLParser.DecimalLiteral = 51;
BODLParser.HexIntegerLiteral = 52;
BODLParser.OctalIntegerLiteral = 53;
BODLParser.OctalIntegerLiteral2 = 54;
BODLParser.BinaryIntegerLiteral = 55;
BODLParser.WhiteSpaces = 56;
BODLParser.LineTerminator = 57;

BODLParser.RULE_program = 0;
BODLParser.RULE_statements = 1;
BODLParser.RULE_importStatement = 2;
BODLParser.RULE_definitions = 3;
BODLParser.RULE_definition = 4;
BODLParser.RULE_block = 5;
BODLParser.RULE_itemList = 6;
BODLParser.RULE_element = 7;
BODLParser.RULE_boAction = 8;
BODLParser.RULE_message = 9;
BODLParser.RULE_node = 10;
BODLParser.RULE_association = 11;
BODLParser.RULE_associationUsingDefinition = 12;
BODLParser.RULE_valuationDefinition = 13;
BODLParser.RULE_valutaionExpressionList = 14;
BODLParser.RULE_valutaionExpression = 15;
BODLParser.RULE_raiseMessage = 16;
BODLParser.RULE_annotations = 17;
BODLParser.RULE_annotation = 18;
BODLParser.RULE_typeList = 19;
BODLParser.RULE_typeDeclaration = 20;
BODLParser.RULE_typeDefaultValue = 21;
BODLParser.RULE_valueAssignList = 22;
BODLParser.RULE_valueAssign = 23;
BODLParser.RULE_multiplicity = 24;
BODLParser.RULE_memberExpression = 25;
BODLParser.RULE_identifierList = 26;
BODLParser.RULE_keyword = 27;
BODLParser.RULE_literal = 28;
BODLParser.RULE_comments = 29;
BODLParser.RULE_compareOperator = 30;
BODLParser.RULE_logicOperator = 31;
BODLParser.RULE_identifier = 32;


function ProgramContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_program;
  return this;
}

ProgramContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
ProgramContext.prototype.constructor = ProgramContext;

ProgramContext.prototype.EOF = function() {
  return this.getToken(BODLParser.EOF, 0);
};

ProgramContext.prototype.statements = function() {
  return this.getTypedRuleContext(StatementsContext, 0);
};

ProgramContext.prototype.definitions = function() {
  return this.getTypedRuleContext(DefinitionsContext, 0);
};

ProgramContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterProgram(this);
  }
};

ProgramContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitProgram(this);
  }
};

ProgramContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitProgram(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.ProgramContext = ProgramContext;

BODLParser.prototype.program = function() {

  var localctx = new ProgramContext(this, this._ctx, this.state);
  this.enterRule(localctx, 0, BODLParser.RULE_program);
  try {
    this.state = 71;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 0, this._ctx);
    switch(la_) {
    case 1:
      this.enterOuterAlt(localctx, 1);
      this.state = 66;
      this.match(BODLParser.EOF);
      break;

    case 2:
      this.enterOuterAlt(localctx, 2);
      this.state = 67;
      this.statements();
      this.state = 68;
      this.definitions();
      this.state = 69;
      this.match(BODLParser.EOF);
      break;

    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function StatementsContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_statements;
  return this;
}

StatementsContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
StatementsContext.prototype.constructor = StatementsContext;

StatementsContext.prototype.importStatement = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTypedRuleContexts(ImportStatementContext);
  } else {
    return this.getTypedRuleContext(ImportStatementContext, i);
  }
};

StatementsContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterStatements(this);
  }
};

StatementsContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitStatements(this);
  }
};

StatementsContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitStatements(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.StatementsContext = StatementsContext;

BODLParser.prototype.statements = function() {

  var localctx = new StatementsContext(this, this._ctx, this.state);
  this.enterRule(localctx, 2, BODLParser.RULE_statements);
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 76;
    this._errHandler.sync(this);
    var _alt = this._interp.adaptivePredict(this._input, 1, this._ctx);
    while(_alt != 2 && _alt != antlr4.atn.ATN.INVALID_ALT_NUMBER) {
      if(_alt === 1) {
        this.state = 73;
        this.importStatement();
      }
      this.state = 78;
      this._errHandler.sync(this);
      _alt = this._interp.adaptivePredict(this._input, 1, this._ctx);
    }

  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function ImportStatementContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_importStatement;
  return this;
}

ImportStatementContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
ImportStatementContext.prototype.constructor = ImportStatementContext;

ImportStatementContext.prototype.IMPORT = function() {
  return this.getToken(BODLParser.IMPORT, 0);
};

ImportStatementContext.prototype.memberExpression = function() {
  return this.getTypedRuleContext(MemberExpressionContext, 0);
};

ImportStatementContext.prototype.SemiColon = function() {
  return this.getToken(BODLParser.SemiColon, 0);
};

ImportStatementContext.prototype.comments = function() {
  return this.getTypedRuleContext(CommentsContext, 0);
};

ImportStatementContext.prototype.AS = function() {
  return this.getToken(BODLParser.AS, 0);
};

ImportStatementContext.prototype.identifier = function() {
  return this.getTypedRuleContext(IdentifierContext, 0);
};

ImportStatementContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterImportStatement(this);
  }
};

ImportStatementContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitImportStatement(this);
  }
};

ImportStatementContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitImportStatement(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.ImportStatementContext = ImportStatementContext;

BODLParser.prototype.importStatement = function() {

  var localctx = new ImportStatementContext(this, this._ctx, this.state);
  this.enterRule(localctx, 4, BODLParser.RULE_importStatement);
  try {
    this.state = 95;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 4, this._ctx);
    switch(la_) {
    case 1:
      this.enterOuterAlt(localctx, 1);
      this.state = 80;
      this._errHandler.sync(this);
      var la_ = this._interp.adaptivePredict(this._input, 2, this._ctx);
      if(la_ === 1) {
        this.state = 79;
        this.comments();

      }
      this.state = 82;
      this.match(BODLParser.IMPORT);
      this.state = 83;
      this.memberExpression();
      this.state = 84;
      this.match(BODLParser.SemiColon);
      break;

    case 2:
      this.enterOuterAlt(localctx, 2);
      this.state = 87;
      this._errHandler.sync(this);
      var la_ = this._interp.adaptivePredict(this._input, 3, this._ctx);
      if(la_ === 1) {
        this.state = 86;
        this.comments();

      }
      this.state = 89;
      this.match(BODLParser.IMPORT);
      this.state = 90;
      this.memberExpression();
      this.state = 91;
      this.match(BODLParser.AS);
      this.state = 92;
      this.identifier();
      this.state = 93;
      this.match(BODLParser.SemiColon);
      break;

    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function DefinitionsContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_definitions;
  return this;
}

DefinitionsContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
DefinitionsContext.prototype.constructor = DefinitionsContext;

DefinitionsContext.prototype.definition = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTypedRuleContexts(DefinitionContext);
  } else {
    return this.getTypedRuleContext(DefinitionContext, i);
  }
};

DefinitionsContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterDefinitions(this);
  }
};

DefinitionsContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitDefinitions(this);
  }
};

DefinitionsContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitDefinitions(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.DefinitionsContext = DefinitionsContext;

BODLParser.prototype.definitions = function() {

  var localctx = new DefinitionsContext(this, this._ctx, this.state);
  this.enterRule(localctx, 6, BODLParser.RULE_definitions);
  var _la = 0; // Token type
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 100;
    this._errHandler.sync(this);
    _la = this._input.LA(1);
    while((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << BODLParser.BUSINESSOBJECT) | (1 << BODLParser.MultiLineComment) | (1 << BODLParser.SingleLineComment) | (1 << BODLParser.CustomAnnotationStart) | (1 << BODLParser.OpenBracket))) !== 0)) {
      this.state = 97;
      this.definition();
      this.state = 102;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function DefinitionContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_definition;
  return this;
}

DefinitionContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
DefinitionContext.prototype.constructor = DefinitionContext;

DefinitionContext.prototype.BUSINESSOBJECT = function() {
  return this.getToken(BODLParser.BUSINESSOBJECT, 0);
};

DefinitionContext.prototype.block = function() {
  return this.getTypedRuleContext(BlockContext, 0);
};

DefinitionContext.prototype.identifier = function() {
  return this.getTypedRuleContext(IdentifierContext, 0);
};

DefinitionContext.prototype.typeDeclaration = function() {
  return this.getTypedRuleContext(TypeDeclarationContext, 0);
};

DefinitionContext.prototype.comments = function() {
  return this.getTypedRuleContext(CommentsContext, 0);
};

DefinitionContext.prototype.annotations = function() {
  return this.getTypedRuleContext(AnnotationsContext, 0);
};

DefinitionContext.prototype.raiseMessage = function() {
  return this.getTypedRuleContext(RaiseMessageContext, 0);
};

DefinitionContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterDefinition(this);
  }
};

DefinitionContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitDefinition(this);
  }
};

DefinitionContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitDefinition(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.DefinitionContext = DefinitionContext;

BODLParser.prototype.definition = function() {

  var localctx = new DefinitionContext(this, this._ctx, this.state);
  this.enterRule(localctx, 8, BODLParser.RULE_definition);
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 104;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 6, this._ctx);
    if(la_ === 1) {
      this.state = 103;
      this.comments();

    }
    this.state = 107;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 7, this._ctx);
    if(la_ === 1) {
      this.state = 106;
      this.annotations();

    }
    this.state = 109;
    this.match(BODLParser.BUSINESSOBJECT);
    this.state = 112;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 8, this._ctx);
    switch(la_) {
    case 1:
      this.state = 110;
      this.identifier();
      break;

    case 2:
      this.state = 111;
      this.typeDeclaration();
      break;

    }
    this.state = 115;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 9, this._ctx);
    if(la_ === 1) {
      this.state = 114;
      this.raiseMessage();

    }
    this.state = 117;
    this.block();
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function BlockContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_block;
  return this;
}

BlockContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
BlockContext.prototype.constructor = BlockContext;

BlockContext.prototype.OpenBrace = function() {
  return this.getToken(BODLParser.OpenBrace, 0);
};

BlockContext.prototype.itemList = function() {
  return this.getTypedRuleContext(ItemListContext, 0);
};

BlockContext.prototype.comments = function() {
  return this.getTypedRuleContext(CommentsContext, 0);
};

BlockContext.prototype.CloseBrace = function() {
  return this.getToken(BODLParser.CloseBrace, 0);
};

BlockContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterBlock(this);
  }
};

BlockContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitBlock(this);
  }
};

BlockContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitBlock(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.BlockContext = BlockContext;

BODLParser.prototype.block = function() {

  var localctx = new BlockContext(this, this._ctx, this.state);
  this.enterRule(localctx, 10, BODLParser.RULE_block);
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 119;
    this.match(BODLParser.OpenBrace);
    this.state = 120;
    this.itemList();
    this.state = 121;
    this.comments();
    this.state = 122;
    this.match(BODLParser.CloseBrace);
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function ItemListContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_itemList;
  return this;
}

ItemListContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
ItemListContext.prototype.constructor = ItemListContext;

ItemListContext.prototype.element = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTypedRuleContexts(ElementContext);
  } else {
    return this.getTypedRuleContext(ElementContext, i);
  }
};

ItemListContext.prototype.message = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTypedRuleContexts(MessageContext);
  } else {
    return this.getTypedRuleContext(MessageContext, i);
  }
};

ItemListContext.prototype.node = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTypedRuleContexts(NodeContext);
  } else {
    return this.getTypedRuleContext(NodeContext, i);
  }
};

ItemListContext.prototype.boAction = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTypedRuleContexts(BoActionContext);
  } else {
    return this.getTypedRuleContext(BoActionContext, i);
  }
};

ItemListContext.prototype.association = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTypedRuleContexts(AssociationContext);
  } else {
    return this.getTypedRuleContext(AssociationContext, i);
  }
};

ItemListContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterItemList(this);
  }
};

ItemListContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitItemList(this);
  }
};

ItemListContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitItemList(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.ItemListContext = ItemListContext;

BODLParser.prototype.itemList = function() {

  var localctx = new ItemListContext(this, this._ctx, this.state);
  this.enterRule(localctx, 12, BODLParser.RULE_itemList);
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 131;
    this._errHandler.sync(this);
    var _alt = this._interp.adaptivePredict(this._input, 11, this._ctx);
    while(_alt != 2 && _alt != antlr4.atn.ATN.INVALID_ALT_NUMBER) {
      if(_alt === 1) {
        this.state = 129;
        this._errHandler.sync(this);
        var la_ = this._interp.adaptivePredict(this._input, 10, this._ctx);
        switch(la_) {
        case 1:
          this.state = 124;
          this.element();
          break;

        case 2:
          this.state = 125;
          this.message();
          break;

        case 3:
          this.state = 126;
          this.node();
          break;

        case 4:
          this.state = 127;
          this.boAction();
          break;

        case 5:
          this.state = 128;
          this.association();
          break;

        }
      }
      this.state = 133;
      this._errHandler.sync(this);
      _alt = this._interp.adaptivePredict(this._input, 11, this._ctx);
    }

  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function ElementContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_element;
  return this;
}

ElementContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
ElementContext.prototype.constructor = ElementContext;

ElementContext.prototype.ELEMENT = function() {
  return this.getToken(BODLParser.ELEMENT, 0);
};

ElementContext.prototype.identifier = function() {
  return this.getTypedRuleContext(IdentifierContext, 0);
};

ElementContext.prototype.Colon = function() {
  return this.getToken(BODLParser.Colon, 0);
};

ElementContext.prototype.typeDeclaration = function() {
  return this.getTypedRuleContext(TypeDeclarationContext, 0);
};

ElementContext.prototype.SemiColon = function() {
  return this.getToken(BODLParser.SemiColon, 0);
};

ElementContext.prototype.comments = function() {
  return this.getTypedRuleContext(CommentsContext, 0);
};

ElementContext.prototype.annotations = function() {
  return this.getTypedRuleContext(AnnotationsContext, 0);
};

ElementContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterElement(this);
  }
};

ElementContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitElement(this);
  }
};

ElementContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitElement(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.ElementContext = ElementContext;

BODLParser.prototype.element = function() {

  var localctx = new ElementContext(this, this._ctx, this.state);
  this.enterRule(localctx, 14, BODLParser.RULE_element);
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 135;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 12, this._ctx);
    if(la_ === 1) {
      this.state = 134;
      this.comments();

    }
    this.state = 138;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 13, this._ctx);
    if(la_ === 1) {
      this.state = 137;
      this.annotations();

    }
    this.state = 140;
    this.match(BODLParser.ELEMENT);
    this.state = 141;
    this.identifier();
    this.state = 142;
    this.match(BODLParser.Colon);
    this.state = 143;
    this.typeDeclaration();
    this.state = 144;
    this.match(BODLParser.SemiColon);
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function BoActionContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_boAction;
  return this;
}

BoActionContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
BoActionContext.prototype.constructor = BoActionContext;

BoActionContext.prototype.ACTION = function() {
  return this.getToken(BODLParser.ACTION, 0);
};

BoActionContext.prototype.identifier = function() {
  return this.getTypedRuleContext(IdentifierContext, 0);
};

BoActionContext.prototype.SemiColon = function() {
  return this.getToken(BODLParser.SemiColon, 0);
};

BoActionContext.prototype.comments = function() {
  return this.getTypedRuleContext(CommentsContext, 0);
};

BoActionContext.prototype.raiseMessage = function() {
  return this.getTypedRuleContext(RaiseMessageContext, 0);
};

BoActionContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterBoAction(this);
  }
};

BoActionContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitBoAction(this);
  }
};

BoActionContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitBoAction(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.BoActionContext = BoActionContext;

BODLParser.prototype.boAction = function() {

  var localctx = new BoActionContext(this, this._ctx, this.state);
  this.enterRule(localctx, 16, BODLParser.RULE_boAction);
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 147;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 14, this._ctx);
    if(la_ === 1) {
      this.state = 146;
      this.comments();

    }
    this.state = 149;
    this.match(BODLParser.ACTION);
    this.state = 150;
    this.identifier();
    this.state = 152;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 15, this._ctx);
    if(la_ === 1) {
      this.state = 151;
      this.raiseMessage();

    }
    this.state = 154;
    this.match(BODLParser.SemiColon);
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function MessageContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_message;
  return this;
}

MessageContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
MessageContext.prototype.constructor = MessageContext;

MessageContext.prototype.MESSAGE = function() {
  return this.getToken(BODLParser.MESSAGE, 0);
};

MessageContext.prototype.identifier = function() {
  return this.getTypedRuleContext(IdentifierContext, 0);
};

MessageContext.prototype.TEXT = function() {
  return this.getToken(BODLParser.TEXT, 0);
};

MessageContext.prototype.StringLiteral = function() {
  return this.getToken(BODLParser.StringLiteral, 0);
};

MessageContext.prototype.Colon = function() {
  return this.getToken(BODLParser.Colon, 0);
};

MessageContext.prototype.typeList = function() {
  return this.getTypedRuleContext(TypeListContext, 0);
};

MessageContext.prototype.SemiColon = function() {
  return this.getToken(BODLParser.SemiColon, 0);
};

MessageContext.prototype.comments = function() {
  return this.getTypedRuleContext(CommentsContext, 0);
};

MessageContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterMessage(this);
  }
};

MessageContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitMessage(this);
  }
};

MessageContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitMessage(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.MessageContext = MessageContext;

BODLParser.prototype.message = function() {

  var localctx = new MessageContext(this, this._ctx, this.state);
  this.enterRule(localctx, 18, BODLParser.RULE_message);
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 157;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 16, this._ctx);
    if(la_ === 1) {
      this.state = 156;
      this.comments();

    }
    this.state = 159;
    this.match(BODLParser.MESSAGE);
    this.state = 160;
    this.identifier();
    this.state = 161;
    this.match(BODLParser.TEXT);
    this.state = 162;
    this.match(BODLParser.StringLiteral);
    this.state = 163;
    this.match(BODLParser.Colon);
    this.state = 164;
    this.typeList();
    this.state = 165;
    this.match(BODLParser.SemiColon);
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function NodeContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_node;
  return this;
}

NodeContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
NodeContext.prototype.constructor = NodeContext;

NodeContext.prototype.NODE = function() {
  return this.getToken(BODLParser.NODE, 0);
};

NodeContext.prototype.identifier = function() {
  return this.getTypedRuleContext(IdentifierContext, 0);
};

NodeContext.prototype.comments = function() {
  return this.getTypedRuleContext(CommentsContext, 0);
};

NodeContext.prototype.annotations = function() {
  return this.getTypedRuleContext(AnnotationsContext, 0);
};

NodeContext.prototype.multiplicity = function() {
  return this.getTypedRuleContext(MultiplicityContext, 0);
};

NodeContext.prototype.raiseMessage = function() {
  return this.getTypedRuleContext(RaiseMessageContext, 0);
};

NodeContext.prototype.block = function() {
  return this.getTypedRuleContext(BlockContext, 0);
};

NodeContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterNode(this);
  }
};

NodeContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitNode(this);
  }
};

NodeContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitNode(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.NodeContext = NodeContext;

BODLParser.prototype.node = function() {

  var localctx = new NodeContext(this, this._ctx, this.state);
  this.enterRule(localctx, 20, BODLParser.RULE_node);
  var _la = 0; // Token type
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 168;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 17, this._ctx);
    if(la_ === 1) {
      this.state = 167;
      this.comments();

    }
    this.state = 171;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 18, this._ctx);
    if(la_ === 1) {
      this.state = 170;
      this.annotations();

    }
    this.state = 173;
    this.match(BODLParser.NODE);
    this.state = 174;
    this.identifier();
    this.state = 176;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 19, this._ctx);
    if(la_ === 1) {
      this.state = 175;
      this.multiplicity();

    }
    this.state = 179;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 20, this._ctx);
    if(la_ === 1) {
      this.state = 178;
      this.raiseMessage();

    }
    this.state = 182;
    this._errHandler.sync(this);
    _la = this._input.LA(1);
    if(_la === BODLParser.OpenBrace) {
      this.state = 181;
      this.block();
    }

  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function AssociationContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_association;
  return this;
}

AssociationContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
AssociationContext.prototype.constructor = AssociationContext;

AssociationContext.prototype.ASSOCIATION = function() {
  return this.getToken(BODLParser.ASSOCIATION, 0);
};

AssociationContext.prototype.identifier = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTypedRuleContexts(IdentifierContext);
  } else {
    return this.getTypedRuleContext(IdentifierContext, i);
  }
};

AssociationContext.prototype.TO = function() {
  return this.getToken(BODLParser.TO, 0);
};

AssociationContext.prototype.SemiColon = function() {
  return this.getToken(BODLParser.SemiColon, 0);
};

AssociationContext.prototype.comments = function() {
  return this.getTypedRuleContext(CommentsContext, 0);
};

AssociationContext.prototype.annotations = function() {
  return this.getTypedRuleContext(AnnotationsContext, 0);
};

AssociationContext.prototype.multiplicity = function() {
  return this.getTypedRuleContext(MultiplicityContext, 0);
};

AssociationContext.prototype.associationUsingDefinition = function() {
  return this.getTypedRuleContext(AssociationUsingDefinitionContext, 0);
};

AssociationContext.prototype.valuationDefinition = function() {
  return this.getTypedRuleContext(ValuationDefinitionContext, 0);
};

AssociationContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterAssociation(this);
  }
};

AssociationContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitAssociation(this);
  }
};

AssociationContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitAssociation(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.AssociationContext = AssociationContext;

BODLParser.prototype.association = function() {

  var localctx = new AssociationContext(this, this._ctx, this.state);
  this.enterRule(localctx, 22, BODLParser.RULE_association);
  var _la = 0; // Token type
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 185;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 22, this._ctx);
    if(la_ === 1) {
      this.state = 184;
      this.comments();

    }
    this.state = 188;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 23, this._ctx);
    if(la_ === 1) {
      this.state = 187;
      this.annotations();

    }
    this.state = 190;
    this.match(BODLParser.ASSOCIATION);
    this.state = 191;
    this.identifier();
    this.state = 193;
    this._errHandler.sync(this);
    _la = this._input.LA(1);
    if(_la === BODLParser.OpenBracket) {
      this.state = 192;
      this.multiplicity();
    }

    this.state = 195;
    this.match(BODLParser.TO);
    this.state = 196;
    this.identifier();
    this.state = 198;
    this._errHandler.sync(this);
    _la = this._input.LA(1);
    if(_la === BODLParser.USING) {
      this.state = 197;
      this.associationUsingDefinition();
    }

    this.state = 201;
    this._errHandler.sync(this);
    _la = this._input.LA(1);
    if(_la === BODLParser.VALUATION) {
      this.state = 200;
      this.valuationDefinition();
    }

    this.state = 203;
    this.match(BODLParser.SemiColon);
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function AssociationUsingDefinitionContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_associationUsingDefinition;
  return this;
}

AssociationUsingDefinitionContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
AssociationUsingDefinitionContext.prototype.constructor = AssociationUsingDefinitionContext;

AssociationUsingDefinitionContext.prototype.USING = function() {
  return this.getToken(BODLParser.USING, 0);
};

AssociationUsingDefinitionContext.prototype.identifier = function() {
  return this.getTypedRuleContext(IdentifierContext, 0);
};

AssociationUsingDefinitionContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterAssociationUsingDefinition(this);
  }
};

AssociationUsingDefinitionContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitAssociationUsingDefinition(this);
  }
};

AssociationUsingDefinitionContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitAssociationUsingDefinition(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.AssociationUsingDefinitionContext = AssociationUsingDefinitionContext;

BODLParser.prototype.associationUsingDefinition = function() {

  var localctx = new AssociationUsingDefinitionContext(this, this._ctx, this.state);
  this.enterRule(localctx, 24, BODLParser.RULE_associationUsingDefinition);
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 205;
    this.match(BODLParser.USING);
    this.state = 206;
    this.identifier();
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function ValuationDefinitionContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_valuationDefinition;
  return this;
}

ValuationDefinitionContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
ValuationDefinitionContext.prototype.constructor = ValuationDefinitionContext;

ValuationDefinitionContext.prototype.VALUATION = function() {
  return this.getToken(BODLParser.VALUATION, 0);
};

ValuationDefinitionContext.prototype.OpenParen = function() {
  return this.getToken(BODLParser.OpenParen, 0);
};

ValuationDefinitionContext.prototype.valutaionExpressionList = function() {
  return this.getTypedRuleContext(ValutaionExpressionListContext, 0);
};

ValuationDefinitionContext.prototype.CloseParen = function() {
  return this.getToken(BODLParser.CloseParen, 0);
};

ValuationDefinitionContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterValuationDefinition(this);
  }
};

ValuationDefinitionContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitValuationDefinition(this);
  }
};

ValuationDefinitionContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitValuationDefinition(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.ValuationDefinitionContext = ValuationDefinitionContext;

BODLParser.prototype.valuationDefinition = function() {

  var localctx = new ValuationDefinitionContext(this, this._ctx, this.state);
  this.enterRule(localctx, 26, BODLParser.RULE_valuationDefinition);
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 208;
    this.match(BODLParser.VALUATION);
    this.state = 209;
    this.match(BODLParser.OpenParen);
    this.state = 210;
    this.valutaionExpressionList();
    this.state = 211;
    this.match(BODLParser.CloseParen);
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function ValutaionExpressionListContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_valutaionExpressionList;
  return this;
}

ValutaionExpressionListContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
ValutaionExpressionListContext.prototype.constructor = ValutaionExpressionListContext;

ValutaionExpressionListContext.prototype.valutaionExpression = function() {
  return this.getTypedRuleContext(ValutaionExpressionContext, 0);
};

ValutaionExpressionListContext.prototype.logicOperator = function() {
  return this.getTypedRuleContext(LogicOperatorContext, 0);
};

ValutaionExpressionListContext.prototype.valutaionExpressionList = function() {
  return this.getTypedRuleContext(ValutaionExpressionListContext, 0);
};

ValutaionExpressionListContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterValutaionExpressionList(this);
  }
};

ValutaionExpressionListContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitValutaionExpressionList(this);
  }
};

ValutaionExpressionListContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitValutaionExpressionList(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.ValutaionExpressionListContext = ValutaionExpressionListContext;

BODLParser.prototype.valutaionExpressionList = function() {

  var localctx = new ValutaionExpressionListContext(this, this._ctx, this.state);
  this.enterRule(localctx, 28, BODLParser.RULE_valutaionExpressionList);
  try {
    this.state = 218;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 27, this._ctx);
    switch(la_) {
    case 1:
      this.enterOuterAlt(localctx, 1);
      this.state = 213;
      this.valutaionExpression();
      break;

    case 2:
      this.enterOuterAlt(localctx, 2);
      this.state = 214;
      this.valutaionExpression();
      this.state = 215;
      this.logicOperator();
      this.state = 216;
      this.valutaionExpressionList();
      break;

    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function ValutaionExpressionContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_valutaionExpression;
  return this;
}

ValutaionExpressionContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
ValutaionExpressionContext.prototype.constructor = ValutaionExpressionContext;

ValutaionExpressionContext.prototype.identifier = function() {
  return this.getTypedRuleContext(IdentifierContext, 0);
};

ValutaionExpressionContext.prototype.compareOperator = function() {
  return this.getTypedRuleContext(CompareOperatorContext, 0);
};

ValutaionExpressionContext.prototype.literal = function() {
  return this.getTypedRuleContext(LiteralContext, 0);
};

ValutaionExpressionContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterValutaionExpression(this);
  }
};

ValutaionExpressionContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitValutaionExpression(this);
  }
};

ValutaionExpressionContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitValutaionExpression(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.ValutaionExpressionContext = ValutaionExpressionContext;

BODLParser.prototype.valutaionExpression = function() {

  var localctx = new ValutaionExpressionContext(this, this._ctx, this.state);
  this.enterRule(localctx, 30, BODLParser.RULE_valutaionExpression);
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 220;
    this.identifier();
    this.state = 221;
    this.compareOperator();
    this.state = 222;
    this.literal();
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function RaiseMessageContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_raiseMessage;
  return this;
}

RaiseMessageContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
RaiseMessageContext.prototype.constructor = RaiseMessageContext;

RaiseMessageContext.prototype.RAISES = function() {
  return this.getToken(BODLParser.RAISES, 0);
};

RaiseMessageContext.prototype.identifierList = function() {
  return this.getTypedRuleContext(IdentifierListContext, 0);
};

RaiseMessageContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterRaiseMessage(this);
  }
};

RaiseMessageContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitRaiseMessage(this);
  }
};

RaiseMessageContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitRaiseMessage(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.RaiseMessageContext = RaiseMessageContext;

BODLParser.prototype.raiseMessage = function() {

  var localctx = new RaiseMessageContext(this, this._ctx, this.state);
  this.enterRule(localctx, 32, BODLParser.RULE_raiseMessage);
  var _la = 0; // Token type
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 225;
    this._errHandler.sync(this);
    _la = this._input.LA(1);
    if(_la === BODLParser.RAISES) {
      this.state = 224;
      this.match(BODLParser.RAISES);
    }

    this.state = 228;
    this._errHandler.sync(this);
    _la = this._input.LA(1);
    if(_la === BODLParser.Identifier) {
      this.state = 227;
      this.identifierList();
    }

  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function AnnotationsContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_annotations;
  return this;
}

AnnotationsContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
AnnotationsContext.prototype.constructor = AnnotationsContext;

AnnotationsContext.prototype.annotation = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTypedRuleContexts(AnnotationContext);
  } else {
    return this.getTypedRuleContext(AnnotationContext, i);
  }
};

AnnotationsContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterAnnotations(this);
  }
};

AnnotationsContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitAnnotations(this);
  }
};

AnnotationsContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitAnnotations(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.AnnotationsContext = AnnotationsContext;

BODLParser.prototype.annotations = function() {

  var localctx = new AnnotationsContext(this, this._ctx, this.state);
  this.enterRule(localctx, 34, BODLParser.RULE_annotations);
  var _la = 0; // Token type
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 233;
    this._errHandler.sync(this);
    _la = this._input.LA(1);
    while(_la === BODLParser.CustomAnnotationStart || _la === BODLParser.OpenBracket) {
      this.state = 230;
      this.annotation();
      this.state = 235;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function AnnotationContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_annotation;
  return this;
}

AnnotationContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
AnnotationContext.prototype.constructor = AnnotationContext;

AnnotationContext.prototype.OpenBracket = function() {
  return this.getToken(BODLParser.OpenBracket, 0);
};

AnnotationContext.prototype.identifier = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTypedRuleContexts(IdentifierContext);
  } else {
    return this.getTypedRuleContext(IdentifierContext, i);
  }
};

AnnotationContext.prototype.CloseBracket = function() {
  return this.getToken(BODLParser.CloseBracket, 0);
};

AnnotationContext.prototype.CustomAnnotationStart = function() {
  return this.getToken(BODLParser.CustomAnnotationStart, 0);
};

AnnotationContext.prototype.OpenParen = function() {
  return this.getToken(BODLParser.OpenParen, 0);
};

AnnotationContext.prototype.CloseParen = function() {
  return this.getToken(BODLParser.CloseParen, 0);
};

AnnotationContext.prototype.literal = function() {
  return this.getTypedRuleContext(LiteralContext, 0);
};

AnnotationContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterAnnotation(this);
  }
};

AnnotationContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitAnnotation(this);
  }
};

AnnotationContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitAnnotation(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.AnnotationContext = AnnotationContext;

BODLParser.prototype.annotation = function() {

  var localctx = new AnnotationContext(this, this._ctx, this.state);
  this.enterRule(localctx, 36, BODLParser.RULE_annotation);
  var _la = 0; // Token type
  try {
    this.state = 263;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 34, this._ctx);
    switch(la_) {
    case 1:
      this.enterOuterAlt(localctx, 1);
      this.state = 237;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      if(_la === BODLParser.CustomAnnotationStart) {
        this.state = 236;
        this.match(BODLParser.CustomAnnotationStart);
      }

      this.state = 239;
      this.match(BODLParser.OpenBracket);
      this.state = 240;
      this.identifier();
      this.state = 241;
      this.match(BODLParser.CloseBracket);
      break;

    case 2:
      this.enterOuterAlt(localctx, 2);
      this.state = 244;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      if(_la === BODLParser.CustomAnnotationStart) {
        this.state = 243;
        this.match(BODLParser.CustomAnnotationStart);
      }

      this.state = 246;
      this.match(BODLParser.OpenBracket);
      this.state = 247;
      this.identifier();
      this.state = 248;
      this.match(BODLParser.OpenParen);
      this.state = 249;
      this.identifier();
      this.state = 250;
      this.match(BODLParser.CloseParen);
      this.state = 251;
      this.match(BODLParser.CloseBracket);
      break;

    case 3:
      this.enterOuterAlt(localctx, 3);
      this.state = 254;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      if(_la === BODLParser.CustomAnnotationStart) {
        this.state = 253;
        this.match(BODLParser.CustomAnnotationStart);
      }

      this.state = 256;
      this.match(BODLParser.OpenBracket);
      this.state = 257;
      this.identifier();
      this.state = 258;
      this.match(BODLParser.OpenParen);
      this.state = 259;
      this.literal();
      this.state = 260;
      this.match(BODLParser.CloseParen);
      this.state = 261;
      this.match(BODLParser.CloseBracket);
      break;

    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function TypeListContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_typeList;
  return this;
}

TypeListContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
TypeListContext.prototype.constructor = TypeListContext;

TypeListContext.prototype.typeDeclaration = function() {
  return this.getTypedRuleContext(TypeDeclarationContext, 0);
};

TypeListContext.prototype.Comma = function() {
  return this.getToken(BODLParser.Comma, 0);
};

TypeListContext.prototype.typeList = function() {
  return this.getTypedRuleContext(TypeListContext, 0);
};

TypeListContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterTypeList(this);
  }
};

TypeListContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitTypeList(this);
  }
};

TypeListContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitTypeList(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.TypeListContext = TypeListContext;

BODLParser.prototype.typeList = function() {

  var localctx = new TypeListContext(this, this._ctx, this.state);
  this.enterRule(localctx, 38, BODLParser.RULE_typeList);
  try {
    this.state = 270;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 35, this._ctx);
    switch(la_) {
    case 1:
      this.enterOuterAlt(localctx, 1);
      this.state = 265;
      this.typeDeclaration();
      break;

    case 2:
      this.enterOuterAlt(localctx, 2);
      this.state = 266;
      this.typeDeclaration();
      this.state = 267;
      this.match(BODLParser.Comma);
      this.state = 268;
      this.typeList();
      break;

    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function TypeDeclarationContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_typeDeclaration;
  return this;
}

TypeDeclarationContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
TypeDeclarationContext.prototype.constructor = TypeDeclarationContext;

TypeDeclarationContext.prototype.identifier = function() {
  return this.getTypedRuleContext(IdentifierContext, 0);
};

TypeDeclarationContext.prototype.typeDefaultValue = function() {
  return this.getTypedRuleContext(TypeDefaultValueContext, 0);
};

TypeDeclarationContext.prototype.memberExpression = function() {
  return this.getTypedRuleContext(MemberExpressionContext, 0);
};

TypeDeclarationContext.prototype.Colon = function() {
  return this.getToken(BODLParser.Colon, 0);
};

TypeDeclarationContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterTypeDeclaration(this);
  }
};

TypeDeclarationContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitTypeDeclaration(this);
  }
};

TypeDeclarationContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitTypeDeclaration(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.TypeDeclarationContext = TypeDeclarationContext;

BODLParser.prototype.typeDeclaration = function() {

  var localctx = new TypeDeclarationContext(this, this._ctx, this.state);
  this.enterRule(localctx, 40, BODLParser.RULE_typeDeclaration);
  var _la = 0; // Token type
  try {
    this.state = 282;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 38, this._ctx);
    switch(la_) {
    case 1:
      this.enterOuterAlt(localctx, 1);
      this.state = 272;
      this.identifier();
      this.state = 274;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      if(_la === BODLParser.Assign) {
        this.state = 273;
        this.typeDefaultValue();
      }

      break;

    case 2:
      this.enterOuterAlt(localctx, 2);
      this.state = 276;
      this.memberExpression();
      this.state = 277;
      this.match(BODLParser.Colon);
      this.state = 278;
      this.identifier();
      this.state = 280;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      if(_la === BODLParser.Assign) {
        this.state = 279;
        this.typeDefaultValue();
      }

      break;

    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function TypeDefaultValueContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_typeDefaultValue;
  return this;
}

TypeDefaultValueContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
TypeDefaultValueContext.prototype.constructor = TypeDefaultValueContext;

TypeDefaultValueContext.prototype.Assign = function() {
  return this.getToken(BODLParser.Assign, 0);
};

TypeDefaultValueContext.prototype.literal = function() {
  return this.getTypedRuleContext(LiteralContext, 0);
};

TypeDefaultValueContext.prototype.OpenBrace = function() {
  return this.getToken(BODLParser.OpenBrace, 0);
};

TypeDefaultValueContext.prototype.valueAssignList = function() {
  return this.getTypedRuleContext(ValueAssignListContext, 0);
};

TypeDefaultValueContext.prototype.CloseBrace = function() {
  return this.getToken(BODLParser.CloseBrace, 0);
};

TypeDefaultValueContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterTypeDefaultValue(this);
  }
};

TypeDefaultValueContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitTypeDefaultValue(this);
  }
};

TypeDefaultValueContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitTypeDefaultValue(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.TypeDefaultValueContext = TypeDefaultValueContext;

BODLParser.prototype.typeDefaultValue = function() {

  var localctx = new TypeDefaultValueContext(this, this._ctx, this.state);
  this.enterRule(localctx, 42, BODLParser.RULE_typeDefaultValue);
  try {
    this.state = 291;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 39, this._ctx);
    switch(la_) {
    case 1:
      this.enterOuterAlt(localctx, 1);
      this.state = 284;
      this.match(BODLParser.Assign);
      this.state = 285;
      this.literal();
      break;

    case 2:
      this.enterOuterAlt(localctx, 2);
      this.state = 286;
      this.match(BODLParser.Assign);
      this.state = 287;
      this.match(BODLParser.OpenBrace);
      this.state = 288;
      this.valueAssignList();
      this.state = 289;
      this.match(BODLParser.CloseBrace);
      break;

    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function ValueAssignListContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_valueAssignList;
  return this;
}

ValueAssignListContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
ValueAssignListContext.prototype.constructor = ValueAssignListContext;

ValueAssignListContext.prototype.valueAssign = function() {
  return this.getTypedRuleContext(ValueAssignContext, 0);
};

ValueAssignListContext.prototype.Comma = function() {
  return this.getToken(BODLParser.Comma, 0);
};

ValueAssignListContext.prototype.valueAssignList = function() {
  return this.getTypedRuleContext(ValueAssignListContext, 0);
};

ValueAssignListContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterValueAssignList(this);
  }
};

ValueAssignListContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitValueAssignList(this);
  }
};

ValueAssignListContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitValueAssignList(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.ValueAssignListContext = ValueAssignListContext;

BODLParser.prototype.valueAssignList = function() {

  var localctx = new ValueAssignListContext(this, this._ctx, this.state);
  this.enterRule(localctx, 44, BODLParser.RULE_valueAssignList);
  try {
    this.state = 298;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 40, this._ctx);
    switch(la_) {
    case 1:
      this.enterOuterAlt(localctx, 1);
      this.state = 293;
      this.valueAssign();
      break;

    case 2:
      this.enterOuterAlt(localctx, 2);
      this.state = 294;
      this.valueAssign();
      this.state = 295;
      this.match(BODLParser.Comma);
      this.state = 296;
      this.valueAssignList();
      break;

    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function ValueAssignContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_valueAssign;
  return this;
}

ValueAssignContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
ValueAssignContext.prototype.constructor = ValueAssignContext;

ValueAssignContext.prototype.identifier = function() {
  return this.getTypedRuleContext(IdentifierContext, 0);
};

ValueAssignContext.prototype.Assign = function() {
  return this.getToken(BODLParser.Assign, 0);
};

ValueAssignContext.prototype.literal = function() {
  return this.getTypedRuleContext(LiteralContext, 0);
};

ValueAssignContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterValueAssign(this);
  }
};

ValueAssignContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitValueAssign(this);
  }
};

ValueAssignContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitValueAssign(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.ValueAssignContext = ValueAssignContext;

BODLParser.prototype.valueAssign = function() {

  var localctx = new ValueAssignContext(this, this._ctx, this.state);
  this.enterRule(localctx, 46, BODLParser.RULE_valueAssign);
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 300;
    this.identifier();
    this.state = 301;
    this.match(BODLParser.Assign);
    this.state = 302;
    this.literal();
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function MultiplicityContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_multiplicity;
  return this;
}

MultiplicityContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
MultiplicityContext.prototype.constructor = MultiplicityContext;

MultiplicityContext.prototype.OpenBracket = function() {
  return this.getToken(BODLParser.OpenBracket, 0);
};

MultiplicityContext.prototype.literal = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTypedRuleContexts(LiteralContext);
  } else {
    return this.getTypedRuleContext(LiteralContext, i);
  }
};

MultiplicityContext.prototype.Comma = function() {
  return this.getToken(BODLParser.Comma, 0);
};

MultiplicityContext.prototype.CloseBracket = function() {
  return this.getToken(BODLParser.CloseBracket, 0);
};

MultiplicityContext.prototype.N = function() {
  return this.getToken(BODLParser.N, 0);
};

MultiplicityContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterMultiplicity(this);
  }
};

MultiplicityContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitMultiplicity(this);
  }
};

MultiplicityContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitMultiplicity(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.MultiplicityContext = MultiplicityContext;

BODLParser.prototype.multiplicity = function() {

  var localctx = new MultiplicityContext(this, this._ctx, this.state);
  this.enterRule(localctx, 48, BODLParser.RULE_multiplicity);
  try {
    this.state = 316;
    this._errHandler.sync(this);
    var la_ = this._interp.adaptivePredict(this._input, 41, this._ctx);
    switch(la_) {
    case 1:
      this.enterOuterAlt(localctx, 1);
      this.state = 304;
      this.match(BODLParser.OpenBracket);
      this.state = 305;
      this.literal();
      this.state = 306;
      this.match(BODLParser.Comma);
      this.state = 307;
      this.literal();
      this.state = 308;
      this.match(BODLParser.CloseBracket);
      break;

    case 2:
      this.enterOuterAlt(localctx, 2);
      this.state = 310;
      this.match(BODLParser.OpenBracket);
      this.state = 311;
      this.literal();
      this.state = 312;
      this.match(BODLParser.Comma);
      this.state = 313;
      this.match(BODLParser.N);
      this.state = 314;
      this.match(BODLParser.CloseBracket);
      break;

    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function MemberExpressionContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_memberExpression;
  return this;
}

MemberExpressionContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
MemberExpressionContext.prototype.constructor = MemberExpressionContext;

MemberExpressionContext.prototype.identifier = function() {
  return this.getTypedRuleContext(IdentifierContext, 0);
};

MemberExpressionContext.prototype.Dot = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTokens(BODLParser.Dot);
  } else {
    return this.getToken(BODLParser.Dot, i);
  }
};


MemberExpressionContext.prototype.memberExpression = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTypedRuleContexts(MemberExpressionContext);
  } else {
    return this.getTypedRuleContext(MemberExpressionContext, i);
  }
};

MemberExpressionContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterMemberExpression(this);
  }
};

MemberExpressionContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitMemberExpression(this);
  }
};

MemberExpressionContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitMemberExpression(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.MemberExpressionContext = MemberExpressionContext;

BODLParser.prototype.memberExpression = function() {

  var localctx = new MemberExpressionContext(this, this._ctx, this.state);
  this.enterRule(localctx, 50, BODLParser.RULE_memberExpression);
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 318;
    this.identifier();
    this.state = 323;
    this._errHandler.sync(this);
    var _alt = this._interp.adaptivePredict(this._input, 42, this._ctx);
    while(_alt != 2 && _alt != antlr4.atn.ATN.INVALID_ALT_NUMBER) {
      if(_alt === 1) {
        this.state = 319;
        this.match(BODLParser.Dot);
        this.state = 320;
        this.memberExpression();
      }
      this.state = 325;
      this._errHandler.sync(this);
      _alt = this._interp.adaptivePredict(this._input, 42, this._ctx);
    }

  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function IdentifierListContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_identifierList;
  return this;
}

IdentifierListContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
IdentifierListContext.prototype.constructor = IdentifierListContext;

IdentifierListContext.prototype.identifier = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTypedRuleContexts(IdentifierContext);
  } else {
    return this.getTypedRuleContext(IdentifierContext, i);
  }
};

IdentifierListContext.prototype.Comma = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTokens(BODLParser.Comma);
  } else {
    return this.getToken(BODLParser.Comma, i);
  }
};


IdentifierListContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterIdentifierList(this);
  }
};

IdentifierListContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitIdentifierList(this);
  }
};

IdentifierListContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitIdentifierList(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.IdentifierListContext = IdentifierListContext;

BODLParser.prototype.identifierList = function() {

  var localctx = new IdentifierListContext(this, this._ctx, this.state);
  this.enterRule(localctx, 52, BODLParser.RULE_identifierList);
  var _la = 0; // Token type
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 326;
    this.identifier();
    this.state = 331;
    this._errHandler.sync(this);
    _la = this._input.LA(1);
    while(_la === BODLParser.Comma) {
      this.state = 327;
      this.match(BODLParser.Comma);
      this.state = 328;
      this.identifier();
      this.state = 333;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function KeywordContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_keyword;
  return this;
}

KeywordContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
KeywordContext.prototype.constructor = KeywordContext;

KeywordContext.prototype.IMPORT = function() {
  return this.getToken(BODLParser.IMPORT, 0);
};

KeywordContext.prototype.BUSINESSOBJECT = function() {
  return this.getToken(BODLParser.BUSINESSOBJECT, 0);
};

KeywordContext.prototype.TO = function() {
  return this.getToken(BODLParser.TO, 0);
};

KeywordContext.prototype.ASSOCIATION = function() {
  return this.getToken(BODLParser.ASSOCIATION, 0);
};

KeywordContext.prototype.ELEMENT = function() {
  return this.getToken(BODLParser.ELEMENT, 0);
};

KeywordContext.prototype.NODE = function() {
  return this.getToken(BODLParser.NODE, 0);
};

KeywordContext.prototype.ACTION = function() {
  return this.getToken(BODLParser.ACTION, 0);
};

KeywordContext.prototype.MESSAGE = function() {
  return this.getToken(BODLParser.MESSAGE, 0);
};

KeywordContext.prototype.RAISES = function() {
  return this.getToken(BODLParser.RAISES, 0);
};

KeywordContext.prototype.USING = function() {
  return this.getToken(BODLParser.USING, 0);
};

KeywordContext.prototype.TEXT = function() {
  return this.getToken(BODLParser.TEXT, 0);
};

KeywordContext.prototype.AS = function() {
  return this.getToken(BODLParser.AS, 0);
};

KeywordContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterKeyword(this);
  }
};

KeywordContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitKeyword(this);
  }
};

KeywordContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitKeyword(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.KeywordContext = KeywordContext;

BODLParser.prototype.keyword = function() {

  var localctx = new KeywordContext(this, this._ctx, this.state);
  this.enterRule(localctx, 54, BODLParser.RULE_keyword);
  var _la = 0; // Token type
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 334;
    _la = this._input.LA(1);
    if(!((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << BODLParser.BUSINESSOBJECT) | (1 << BODLParser.TO) | (1 << BODLParser.ASSOCIATION) | (1 << BODLParser.ELEMENT) | (1 << BODLParser.NODE) | (1 << BODLParser.ACTION) | (1 << BODLParser.MESSAGE) | (1 << BODLParser.RAISES) | (1 << BODLParser.USING) | (1 << BODLParser.IMPORT) | (1 << BODLParser.AS) | (1 << BODLParser.TEXT))) !== 0))) {
      this._errHandler.recoverInline(this);
    } else {
        	this._errHandler.reportMatch(this);
      this.consume();
    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function LiteralContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_literal;
  return this;
}

LiteralContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
LiteralContext.prototype.constructor = LiteralContext;

LiteralContext.prototype.DecimalLiteral = function() {
  return this.getToken(BODLParser.DecimalLiteral, 0);
};

LiteralContext.prototype.BooleanLiteral = function() {
  return this.getToken(BODLParser.BooleanLiteral, 0);
};

LiteralContext.prototype.StringLiteral = function() {
  return this.getToken(BODLParser.StringLiteral, 0);
};

LiteralContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterLiteral(this);
  }
};

LiteralContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitLiteral(this);
  }
};

LiteralContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitLiteral(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.LiteralContext = LiteralContext;

BODLParser.prototype.literal = function() {

  var localctx = new LiteralContext(this, this._ctx, this.state);
  this.enterRule(localctx, 56, BODLParser.RULE_literal);
  var _la = 0; // Token type
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 336;
    _la = this._input.LA(1);
    if(!(((((_la - 49)) & ~0x1f) == 0 && ((1 << (_la - 49)) & ((1 << (BODLParser.StringLiteral - 49)) | (1 << (BODLParser.BooleanLiteral - 49)) | (1 << (BODLParser.DecimalLiteral - 49)))) !== 0))) {
      this._errHandler.recoverInline(this);
    } else {
        	this._errHandler.reportMatch(this);
      this.consume();
    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function CommentsContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_comments;
  return this;
}

CommentsContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
CommentsContext.prototype.constructor = CommentsContext;

CommentsContext.prototype.SingleLineComment = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTokens(BODLParser.SingleLineComment);
  } else {
    return this.getToken(BODLParser.SingleLineComment, i);
  }
};


CommentsContext.prototype.MultiLineComment = function(i) {
  if(i === undefined) {
    i = null;
  }
  if(i === null) {
    return this.getTokens(BODLParser.MultiLineComment);
  } else {
    return this.getToken(BODLParser.MultiLineComment, i);
  }
};


CommentsContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterComments(this);
  }
};

CommentsContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitComments(this);
  }
};

CommentsContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitComments(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.CommentsContext = CommentsContext;

BODLParser.prototype.comments = function() {

  var localctx = new CommentsContext(this, this._ctx, this.state);
  this.enterRule(localctx, 58, BODLParser.RULE_comments);
  var _la = 0; // Token type
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 341;
    this._errHandler.sync(this);
    _la = this._input.LA(1);
    while(_la === BODLParser.MultiLineComment || _la === BODLParser.SingleLineComment) {
      this.state = 338;
      _la = this._input.LA(1);
      if(!(_la === BODLParser.MultiLineComment || _la === BODLParser.SingleLineComment)) {
        this._errHandler.recoverInline(this);
      } else {
            	this._errHandler.reportMatch(this);
        this.consume();
      }
      this.state = 343;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function CompareOperatorContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_compareOperator;
  return this;
}

CompareOperatorContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
CompareOperatorContext.prototype.constructor = CompareOperatorContext;

CompareOperatorContext.prototype.Equals_ = function() {
  return this.getToken(BODLParser.Equals_, 0);
};

CompareOperatorContext.prototype.NotEquals = function() {
  return this.getToken(BODLParser.NotEquals, 0);
};

CompareOperatorContext.prototype.MoreThan = function() {
  return this.getToken(BODLParser.MoreThan, 0);
};

CompareOperatorContext.prototype.GreaterThanEquals = function() {
  return this.getToken(BODLParser.GreaterThanEquals, 0);
};

CompareOperatorContext.prototype.LessThan = function() {
  return this.getToken(BODLParser.LessThan, 0);
};

CompareOperatorContext.prototype.LessThanEquals = function() {
  return this.getToken(BODLParser.LessThanEquals, 0);
};

CompareOperatorContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterCompareOperator(this);
  }
};

CompareOperatorContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitCompareOperator(this);
  }
};

CompareOperatorContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitCompareOperator(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.CompareOperatorContext = CompareOperatorContext;

BODLParser.prototype.compareOperator = function() {

  var localctx = new CompareOperatorContext(this, this._ctx, this.state);
  this.enterRule(localctx, 60, BODLParser.RULE_compareOperator);
  var _la = 0; // Token type
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 344;
    _la = this._input.LA(1);
    if(!(((((_la - 38)) & ~0x1f) == 0 && ((1 << (_la - 38)) & ((1 << (BODLParser.LessThan - 38)) | (1 << (BODLParser.MoreThan - 38)) | (1 << (BODLParser.LessThanEquals - 38)) | (1 << (BODLParser.GreaterThanEquals - 38)) | (1 << (BODLParser.Equals_ - 38)) | (1 << (BODLParser.NotEquals - 38)))) !== 0))) {
      this._errHandler.recoverInline(this);
    } else {
        	this._errHandler.reportMatch(this);
      this.consume();
    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function LogicOperatorContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_logicOperator;
  return this;
}

LogicOperatorContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
LogicOperatorContext.prototype.constructor = LogicOperatorContext;

LogicOperatorContext.prototype.And = function() {
  return this.getToken(BODLParser.And, 0);
};

LogicOperatorContext.prototype.Or = function() {
  return this.getToken(BODLParser.Or, 0);
};

LogicOperatorContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterLogicOperator(this);
  }
};

LogicOperatorContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitLogicOperator(this);
  }
};

LogicOperatorContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitLogicOperator(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.LogicOperatorContext = LogicOperatorContext;

BODLParser.prototype.logicOperator = function() {

  var localctx = new LogicOperatorContext(this, this._ctx, this.state);
  this.enterRule(localctx, 62, BODLParser.RULE_logicOperator);
  var _la = 0; // Token type
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 346;
    _la = this._input.LA(1);
    if(!(_la === BODLParser.And || _la === BODLParser.Or)) {
      this._errHandler.recoverInline(this);
    } else {
        	this._errHandler.reportMatch(this);
      this.consume();
    }
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


function IdentifierContext(parser, parent, invokingState) {
  if(parent === undefined) {
	    parent = null;
  }
  if(invokingState === undefined || invokingState === null) {
    invokingState = -1;
  }
  antlr4.ParserRuleContext.call(this, parent, invokingState);
  this.parser = parser;
  this.ruleIndex = BODLParser.RULE_identifier;
  return this;
}

IdentifierContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
IdentifierContext.prototype.constructor = IdentifierContext;

IdentifierContext.prototype.Identifier = function() {
  return this.getToken(BODLParser.Identifier, 0);
};

IdentifierContext.prototype.enterRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.enterIdentifier(this);
  }
};

IdentifierContext.prototype.exitRule = function(listener) {
  if(listener instanceof BODLParserListener ) {
    listener.exitIdentifier(this);
  }
};

IdentifierContext.prototype.accept = function(visitor) {
  if ( visitor instanceof BODLParserVisitor ) {
    return visitor.visitIdentifier(this);
  } else {
    return visitor.visitChildren(this);
  }
};




BODLParser.IdentifierContext = IdentifierContext;

BODLParser.prototype.identifier = function() {

  var localctx = new IdentifierContext(this, this._ctx, this.state);
  this.enterRule(localctx, 64, BODLParser.RULE_identifier);
  try {
    this.enterOuterAlt(localctx, 1);
    this.state = 348;
    this.match(BODLParser.Identifier);
  } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
  } finally {
    this.exitRule();
  }
  return localctx;
};


exports.BODLParser = BODLParser;
