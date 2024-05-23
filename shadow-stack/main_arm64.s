#include "textflag.h"

TEXT ·regfp(SB),$8-0
	MOVD (R29), R0
	MOVD R0, ret+0(FP)
	RET

TEXT ·trampoline(SB),NOSPLIT|NOFRAME,$0-0
	CALL main·destroyShadowStack(SB)
	MOVD ret+0(FP), R30
	RET
