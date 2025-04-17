; Addition - by Timendus
; Intended as a paper tape loop!
;
; This program adds the binary number on the input switches to the scratch
; register, rings the bell and halts. Assuming we start from a reset scratch
; register, you can input a number using the switches, run the program until it
; halts, input a second number, run the program again until it halts and you
; will see the resulting addition in the output register. You can keep
; repeating the process to add more numbers together.

; Enable input and output
ONE
IEN
OEN

; Add input register to scratch register and store back into scratch register
; Also; show result on output register
LD SR0
ADD IR1
STO SR0
STO OR0
LD SR1
ADD IR2
STO SR1
STO OR1
LD SR2
ADD IR3
STO SR2
STO OR2
LD SR3
ADD IR4
STO SR3
STO OR3
LD SR4
ADD IR5
STO SR4
STO OR4
LD SR5
ADD IR6
STO SR5
STO OR5
LD SR6
ADD IR7
STO SR6
STO OR6

; We're done!
IOC    ; Ring bell
NOPF   ; Halt program

; Aaaaand loop!
