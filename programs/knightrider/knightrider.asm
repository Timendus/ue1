; Knight Rider - by Timendus
; Intended as a paper tape loop!
;
; Shows a "Knight Rider" LED pattern on the output register. Seems about right
; at around 50Hz.

; Enable input and output
ONE
IEN
OEN

; Init output register
STO OR0
NAND RR
STO OR1
STO OR2
STO OR3
STO OR4
STO OR5
STO OR6
STO OR7

;;;;;;;;;;;;;;
;; MOVING LEFT

STO OR0
STOC OR1

NOP0  ; Waste some time
NOP0
NOP0
NOP0
NOP0
NOP0
NOP0

STO OR1
STOC OR2

NOP0  ; Waste some time
NOP0
NOP0
NOP0
NOP0
NOP0
NOP0

STO OR2
STOC OR3

NOP0  ; Waste some time
NOP0
NOP0
NOP0
NOP0
NOP0
NOP0

STO OR3
STOC OR4

NOP0  ; Waste some time
NOP0
NOP0
NOP0
NOP0
NOP0
NOP0

STO OR4
STOC OR5

NOP0  ; Waste some time
NOP0
NOP0
NOP0
NOP0
NOP0
NOP0

STO OR5
STOC OR6

NOP0  ; Waste some time
NOP0
NOP0
NOP0
NOP0
NOP0
NOP0

STO OR6
STOC OR7

NOP0  ; Waste some time
NOP0
NOP0
NOP0
NOP0
NOP0
NOP0

;;;;;;;;;;;;;;;
;; MOVING RIGHT

STO OR7
STOC OR6

NOP0  ; Waste some time
NOP0
NOP0
NOP0
NOP0
NOP0
NOP0

STO OR6
STOC OR5

NOP0  ; Waste some time
NOP0
NOP0
NOP0
NOP0
NOP0
NOP0

STO OR5
STOC OR4

NOP0  ; Waste some time
NOP0
NOP0
NOP0
NOP0
NOP0
NOP0

STO OR4
STOC OR3

NOP0  ; Waste some time
NOP0
NOP0
NOP0
NOP0
NOP0
NOP0

STO OR3
STOC OR2

NOP0  ; Waste some time
NOP0
NOP0
NOP0
NOP0
NOP0
NOP0

STO OR2
STOC OR1

NOP0  ; Waste some time
NOP0
NOP0
NOP0

; Aaaaand loop!
