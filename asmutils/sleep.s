#include "textflag.h"

TEXT ·SleepASM(SB), NOSPLIT, $0-16
    MOVQ ms+0(FP), BX          
outer:
    TESTQ BX, BX
    JE done
    MOVQ $500000, CX           
inner:
    DECQ CX
    JNZ inner
    DECQ BX
    JMP outer
done:
    RET
    