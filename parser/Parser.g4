grammar Parser;

//prog: stat+;
stat: expr NEWLINE          # printExpr
    | ID '=' expr NEWLINE   # assign
    | NEWLINE               # blank
    ;

expr: '(' expr ')'          # paren
    | expr ('*'|'/') expr   # multiDiv
    | expr ('+'|'-') expr   # addSub
    | INT                   # int
    | ID                    # id
    ;

INT: [0-9]+;
ID : [a-zA-Z]+;
NEWLINE: '\r'?'\n';
WS: [ \t]+ ->skip;

MUL: '*';
DIV: '/';
SUB: '-';
ADD: '+';