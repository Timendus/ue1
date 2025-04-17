; Fibonacci
; This program should calculate the Fibonacci sequence.
; Based on program as seen in https://www.youtube.com/watch?v=JsbzHNOEsZ4, but
; modified a bit to make it work for me.

; CPU initialization
ONE
IEN RR
OEN RR
NAND RR

; Empty Scratch Register and Output Register
STO SR0
STO SR1
STO SR2
STO SR3
STO SR4
STO SR5
STO SR6
STO SR7

STO OR0
STO OR1
STO OR2
STO OR3
STO OR4
STO OR5
STO OR6
STO OR7

; Initialze the two 4-bit operands in the scratch register
STOC SR0
STOC SR4

; Add the two 4-bit operands together and store result in output
LD SR0
ADD SR4
STO SR0
LD SR1
ADD SR5
STO SR1
LD SR2
ADD SR6
STO SR2
LD SR3
ADD SR7
STO SR3

; Store result to output register
LD SR0
STO OR0
LD SR1
STO OR1
LD SR2
STO OR2
LD SR3
STO OR3


; Add the two 4-bit operands together and store result in output
LD SR4
ADD SR0
STO SR4
LD SR5
ADD SR1
STO SR5
LD SR6
ADD SR2
STO SR6
LD SR7
ADD SR3
STO SR7

; Store result to output register
LD SR4
STO OR0
LD SR5
STO OR1
LD SR6
STO OR2
LD SR7
STO OR3


; Add the two 4-bit operands together and store result in output
LD SR0
ADD SR4
STO SR0
LD SR1
ADD SR5
STO SR1
LD SR2
ADD SR6
STO SR2
LD SR3
ADD SR7
STO SR3

; Store result to output register
LD SR0
STO OR0
LD SR1
STO OR1
LD SR2
STO OR2
LD SR3
STO OR3


; Add the two 4-bit operands together and store result in output
LD SR4
ADD SR0
STO SR4
LD SR5
ADD SR1
STO SR5
LD SR6
ADD SR2
STO SR6
LD SR7
ADD SR3
STO SR7

; Store result to output register
LD SR4
STO OR0
LD SR5
STO OR1
LD SR6
STO OR2
LD SR7
STO OR3


; Add the two 4-bit operands together and store result in output
LD SR0
ADD SR4
STO SR0
LD SR1
ADD SR5
STO SR1
LD SR2
ADD SR6
STO SR2
LD SR3
ADD SR7
STO SR3

; Store result to output register
LD SR0
STO OR0
LD SR1
STO OR1
LD SR2
STO OR2
LD SR3
STO OR3


; Add the two 4-bit operands together one last time and store result in output
LD SR4
ADD SR0
STO OR0
LD SR5
ADD SR1
STO OR1
LD SR6
ADD SR2
STO OR2
LD SR7
ADD SR3
STO OR3
; Show carry flag too
ADD RR
STO OR4


; We're done!
IOC SR0 ; Ring bell
NOPF SR0 ; Halt program
