; Bell test - by Timendus
; Intended as a paper tape loop!
;
; Flip any of the input switches and the bell should ring once, each time you
; make a change.
;
; Unfortunately the bell doesn't ring every time in my emulator, but I'm
; assuming that's an issue with the emulator and not this program.

; CPU initialization
ONE
IEN
OEN
STOC SR0  ; Flag that we have seen no differences (0)

; Compare scratch register to input switches
LD IR1
XOR SR1
SKZ      ; If they are the same, result register is 0, skip next instruction
STO SR0  ; Otherwise, result is 1, store in flag
LD IR2
XOR SR2
SKZ
STO SR0
LD IR3
XOR SR3
SKZ
STO SR0
LD IR4
XOR SR4
SKZ
STO SR0
LD IR5
XOR SR5
SKZ
STO SR0
LD IR6
XOR SR6
SKZ
STO SR0
LD IR7
XOR SR7
SKZ
STO SR0

; Beep if we have seen differences (SR0 is 1)
LD SR0
SKZ
IOC      ; Ring bell

; Copy input register to scratch register
LD IR1
STO SR1
LD IR2
STO SR2
LD IR3
STO SR3
LD IR4
STO SR4
LD IR5
STO SR5
LD IR6
STO SR6
LD IR7
STO SR7

; Aaaaand loop...
