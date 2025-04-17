# Usagi Electric 1

I wanted to see if I could write a couple of small assembly programs for the
[UE1 vacuum tube computer](https://github.com/Nakazoto/UEVTC) by [Usagi
Electric](https://www.youtube.com/@UsagiElectric). Because the machine is epic
in its simplicity and its complexity at the same time.

[This is the video](https://www.youtube.com/watch?v=JsbzHNOEsZ4) that
nerd-sniped me to start this little project. It has cost me three evenings so
far 😄

## Tools

To familiarize myself with the machine and its assembly language, I first
re-implemented his assembler and emulator in Go. If you are interested in these
tools, the binaries are in the [dist directory](./dist/) of this repository. The
sources are in the [assembler](./assembler/) and [emulator](./emulator/)
directories.

## Assembly programs

If you are interested in the assembly programs that I have found and/or
(re)written, see the [programs directory](./programs).

The directory structures are set up in such a way to also be my test suite for
the assembler. Running `make tests` in the assembler directory should (if you
have Go installed) assemble each `.asm` file and compare the result to the
corresponding `.bin` file.
