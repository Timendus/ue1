# Usagi Electric 1

I wanted to see if I could write a couple of small assembly programs for the
[UE1 vacuum tube computer](https://github.com/Nakazoto/UEVTC) by [Usagi
Electric](https://www.youtube.com/@UsagiElectric). Because the machine is epic
in its simplicity and its complexity at the same time.

[This is the video](https://www.youtube.com/watch?v=JsbzHNOEsZ4) that
nerd-sniped me to start this little project. It has cost me three evenings so
far 😄

If you are interested in the assembly programs that I have found and/or
(re)written, see the [programs directory](./programs).

## Tools

To familiarize myself with the machine and its assembly language, I first
re-implemented Usagi Electric's assembler and emulator in Go. If you are
interested in these tools, the binaries are in the [dist directory](./dist/) of
this repository. The sources are in the [assembler](./assembler/) and
[emulator](./emulator/) directories.

### Assembler

Usage:

```
ue1asm <input file> <ouput file>
```

This will assemble the input source file and write the resulting binary to the
output file.

### Emulator

Usage:

```
ue1emu <input file> [<cpu speed in hz>]
```

Input file should be a binary. The CPU speed is optional; it will run the binary
at 50Hz by default. If you specify a different number as the second parameter,
that speed will be used.

When the emulator is running, you can use `h` to halt the system, `s` to step
through the instructions (halting after every instruction) or `r` to resume
running.

Pressing the numbers 1-7 on your keyboard will flip the input switches for bits
7-1.

Press Escape to stop the program.

## Development

There are Makefiles in the directories for the emulator and assembler with
commands for compiling and testing.

The directory structure of the assembly progams is set up in such a way to also
be the test suite for the assembler. Running `make tests` in the assembler
directory should (if you have Go installed) assemble each `.asm` file and
compare the result to the corresponding `.bin` file.
