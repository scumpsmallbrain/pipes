# Pipe Problem


![pipes](https://github.com/scumpsmallbrain/pipes/assets/136015437/ab435ae2-a591-494c-93ee-7cb00ec43699)

*v0.1.1*

pipe problem is a logic puzzle devised by my buddy will. there are 7 pipes. each pipe can be empty, represented by 0, half-full, represented by 1, or full, represented by 2.

you can fill a pipe, sending it up one number, but you must subtract 1 from each neighboring pipe. if a neighboring pipe is empty, you cannot fill a pipe. same goes for removing liquid from a pipe.

the goal is to get all pipes to be half-full.

## controls:

```
left/right arrow keys: move cursor
up: add to pipe
down: subtract from pipe 
```

## technical details

now with working binaries! still raises windows defender warnings i'll work on that lol.

if you have the go compiler installed, you can build with `$ go build` or run with `$ go run .`

only tested on windows so far


---
coded with love by scump smallbrain @2023
