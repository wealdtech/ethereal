grammar Func;

start
   : funcName '(' funcArgs ')' EOF
   ;

funcName
   : NAME
   ;

funcArgs
   : (arg (',' arg)*)?
   ;

arg
   : intArg
   | hexArg
   | stringArg
   | boolArg
   | arrayArg
   ;

intArg
   : '-'? INT
   ;

hexArg
   : HEX
   ;

stringArg
   : STRING
   ;

boolArg
   : 'true'
   | 'false'
   ;

arrayArg
   : '[' funcArgs ']'
   ;

NAME
   : NAMESTART NAMEPART*
   ;

INT
   : DIGIT+
   ;

HEX
   : '0x' HEXDIGIT+
   ;

STRING
   : '"' DOUBLEQUOTEDCHAR* '"'
   | '\'' SINGLEQUOTEDCHAR* '\''
   ;

BOOL
   : TRUE
   | FALSE
   ;

fragment
TRUE
   : 'true'
   | 'True'
   ;

fragment
FALSE
   : 'false'
   | 'False'
   ;

fragment
DOUBLEQUOTEDCHAR
   : ~["\r\n\\] | ('\\' .)
   ;

fragment
SINGLEQUOTEDCHAR
   : ~['\r\n\\] | ('\\' .)
   ;

fragment
NAMESTART
   : LETTER | '$' | '_'
   ;

fragment
NAMEPART
   : LETTER | '$' | '_' | DIGIT
   ;

fragment
DIGIT
   : ('0' .. '9')
   ;

fragment
HEXDIGIT
   : (DIGIT | 'A' .. 'F' | 'a' .. 'f')
   ;

fragment
LETTER
  : ('A' .. 'Z' | 'a' .. 'z')
  ;

WS
   : [ \r\n\t\u000C]+ -> skip
   ;
