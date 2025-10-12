#include "textflag.h"

TEXT ·SumFood(SB), NOSPLIT, $0-16
    MOVQ a+0(FP), AX
    ADDQ b+8(FP), AX
    MOVQ AX, ret+16(FP)
    RET
    