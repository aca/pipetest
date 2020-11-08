# pipetest

Small script to test command line output with pipe.   
Soak up standard input and test output.

Supports 
- Operators for text strings: "=/==",  "!="
- Operators to compare and examine numbers: "-eq", "-ne", "-gt", "-ge", "-lt", "-le"

---

For example, Instead of 
```sh
#!/usr/bin/env fish
if test (echo -n 1) = 1; echo "true"; end
```

Test like
```sh
#!/usr/bin/env fish
if echo -n 1 | pipetest = 1; echo "true"; end
```
