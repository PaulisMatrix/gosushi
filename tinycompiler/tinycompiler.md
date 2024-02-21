**Teeny Tiny Compiler in go.**

Followed this blog post: https://austinhenley.com/blog/teenytinycompiler1.html

* Rough compilation process:
1. Convert the program into token, small bits of words.
2. Segregate them depending on token type such as identifiers, keywords, operators, etc (lexing done)
3. Parse these tokens one by one according to the language grammer defined and form an AST(Abstract Syntax Tree). But here, in this case we are just defining a separate fun for each grammer for simplicity.
4. The language grammer dictates how your language syntax/semantics will be. How will your statements will get executed and produce the results you expect. (parsing done)
5. Compile the program, i.e execute each instruction to produce the desired results.

* Example Code: 

```
PRINT "How many fibonacci numbers do you want?"
INPUT nums
PRINT ""

LET a = 0
LET b = 1
WHILE nums > 0 REPEAT
    PRINT a
    LET c = a + b
    LET a = b
    LET b = c
    LET nums = nums - 1
ENDWHILE
```


```
program ::= {statement}
statement ::= "PRINT" (expression | string) nl
    | "IF" comparison "THEN" nl {statement} "ENDIF" nl
    | "WHILE" comparison "REPEAT" nl {statement} "ENDWHILE" nl
    | "LABEL" ident nl
    | "GOTO" ident nl
    | "LET" ident "=" expression nl
    | "INPUT" ident nl
comparison ::= expression (("==" | "!=" | ">" | ">=" | "<" | "<=") expression)+
expression ::= term {( "-" | "+" ) term}
term ::= unary {( "/" | "*" ) unary}
unary ::= {"+" | "-"} primary
primary ::= number | ident
nl ::= '\n'+
```

a. `{}` : means zero or more<br>
b. `[]` : means zero or one<br>
c. `+` : means one or more of whatever is to the left<br>
d. `()` : used for grouping<br>

* TODO
0. Add tests for each step(lexing, parsing, emitting)
1. Implement expressions parsing using an AST.
2. All the primitive, basic data types like int, float, string, arrays, etc. with type checking.
3. User defined data types like struct, union, enum, typedef.
4. All boolean operators like AND, OR, XOR, etc.
5. Add support for functions, functions as first class objects, recursive functions, etc. 
6. Add support for std library for file/os/string/math/ operations.
