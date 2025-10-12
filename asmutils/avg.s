#include "textflag.h"

TEXT ·CalcDailyFeedAverage(SB), NOSPLIT, $0-24
    MOVQ total+0(FP), AX
    MOVQ count+8(FP), BX
    CMPQ BX, $0
    JE zero_div
    CQO
    IDIVQ BX
    MOVQ AX, ret+16(FP)
    RET
zero_div:
    MOVQ $0, ret+16(FP)
    RET
    