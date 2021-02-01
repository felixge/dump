// func cputicks() int64
TEXT ·cputicks(SB),$0-0
	CMPB	runtime·lfenceBeforeRdtsc(SB), $1
	JNE	mfence
	LFENCE
	JMP	done
mfence:
	MFENCE
done:
	RDTSC
	SHLQ	$32, DX
	ADDQ	DX, AX
	MOVQ	AX, ret+0(FP)
	RET
