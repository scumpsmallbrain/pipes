# Pipe Problem

*v0.0.1*

pipe problem is a logic puzzle devised by my buddy will. there are 7 pipes. each pipe can be empty, represented by 0, half-full, represented by 1, or full, represented by 2.

you can fill a pipe, sending it up one number, but you must subtract 1 from each neighboring pipe. if a neighboring pipe is empty, you cannot fill a pipe. same goes for removing liquid from a pipe.

the goal is to get all pipes to be half-full.

## commands:

```
<a>+	-- add liquid to a pipe, where <a> = the letter of the pipe
<a>-	-- remove liquid from a pipe
q		-- quit game
```

## technical details

a windows x86 build is included in the repo. you can compile to your own machine with the go compiler or run with `$ go run .`



---
coded with love by scump smallbrain @2023