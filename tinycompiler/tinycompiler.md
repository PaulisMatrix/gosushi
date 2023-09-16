Teeny Tiny Compiler in go.

Followed this blog post: https://austinhenley.com/blog/teenytinycompiler1.html


Example Code: 

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