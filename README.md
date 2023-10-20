# Pipe Problem

![pipes](https://github.com/scumpsmallbrain/pipes/assets/136015437/774716d5-2b81-405d-93d5-5d660fd989d5)

*v0.0.3*

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

now with working binaries! still raises windows defender warnings i'll work on that lol.

if you have the go compiler installed, you can build with `$ go build` or run with `$ go run .`

only tested on windows so far


---
coded with love by scump smallbrain @2023
