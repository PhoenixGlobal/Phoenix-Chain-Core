	.text
	.file	"<stdin>"
	.globl	makeNIST_P192Lbmi2
	.align	16, 0x90
	.type	makeNIST_P192Lbmi2,@function
makeNIST_P192Lbmi2:                     # @makeNIST_P192Lbmi2
# BB#0:
	movq	$-1, %rax
	movq	$-2, %rdx
	movq	$-1, %rcx
	retq
.Lfunc_end0:
	.size	makeNIST_P192Lbmi2, .Lfunc_end0-makeNIST_P192Lbmi2

	.globl	mcl_fpDbl_mod_NIST_P192Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_mod_NIST_P192Lbmi2,@function
mcl_fpDbl_mod_NIST_P192Lbmi2:           # @mcl_fpDbl_mod_NIST_P192Lbmi2
# BB#0:
	pushq	%r14
	pushq	%rbx
	movq	16(%rsi), %r10
	movq	24(%rsi), %r8
	movq	40(%rsi), %r9
	movq	8(%rsi), %rax
	addq	%r9, %rax
	adcq	$0, %r10
	sbbq	%rcx, %rcx
	andl	$1, %ecx
	movq	32(%rsi), %r11
	movq	(%rsi), %r14
	addq	%r8, %r14
	adcq	%r11, %rax
	adcq	%r9, %r10
	adcq	$0, %rcx
	addq	%r9, %r14
	adcq	%r8, %rax
	adcq	%r11, %r10
	adcq	$0, %rcx
	addq	%rcx, %r14
	adcq	%rax, %rcx
	adcq	$0, %r10
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%r14, %rsi
	addq	$1, %rsi
	movq	%rcx, %rdx
	adcq	$1, %rdx
	movq	%r10, %rbx
	adcq	$0, %rbx
	adcq	$-1, %rax
	andl	$1, %eax
	cmovneq	%r14, %rsi
	movq	%rsi, (%rdi)
	testb	%al, %al
	cmovneq	%rcx, %rdx
	movq	%rdx, 8(%rdi)
	cmovneq	%r10, %rbx
	movq	%rbx, 16(%rdi)
	popq	%rbx
	popq	%r14
	retq
.Lfunc_end1:
	.size	mcl_fpDbl_mod_NIST_P192Lbmi2, .Lfunc_end1-mcl_fpDbl_mod_NIST_P192Lbmi2

	.globl	mcl_fp_sqr_NIST_P192Lbmi2
	.align	16, 0x90
	.type	mcl_fp_sqr_NIST_P192Lbmi2,@function
mcl_fp_sqr_NIST_P192Lbmi2:              # @mcl_fp_sqr_NIST_P192Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rdi, -8(%rsp)          # 8-byte Spill
	movq	16(%rsi), %r8
	movq	(%rsi), %rcx
	movq	8(%rsi), %rsi
	movq	%r8, %rdx
	mulxq	%rsi, %r14, %rbx
	movq	%rbx, -16(%rsp)         # 8-byte Spill
	movq	%rsi, %rdx
	mulxq	%rsi, %r13, %r15
	movq	%rsi, %rdx
	mulxq	%rcx, %r12, %rsi
	addq	%rsi, %r13
	adcq	%r14, %r15
	adcq	$0, %rbx
	movq	%rcx, %rdx
	mulxq	%rcx, %r9, %rax
	addq	%r12, %rax
	movq	%r8, %rdx
	mulxq	%rcx, %rbp, %r11
	adcq	%rbp, %rsi
	movq	%r11, %r10
	adcq	$0, %r10
	addq	%r12, %rax
	adcq	%r13, %rsi
	adcq	%r15, %r10
	adcq	$0, %rbx
	movq	%r8, %rdx
	mulxq	%r8, %rcx, %rdi
	addq	%r14, %r11
	adcq	-16(%rsp), %rcx         # 8-byte Folded Reload
	adcq	$0, %rdi
	addq	%rbp, %rsi
	adcq	%r10, %r11
	adcq	%rbx, %rcx
	adcq	$0, %rdi
	addq	%rdi, %rax
	adcq	$0, %rsi
	sbbq	%rdx, %rdx
	andl	$1, %edx
	addq	%r11, %r9
	adcq	%rcx, %rax
	adcq	%rdi, %rsi
	adcq	$0, %rdx
	addq	%rdi, %r9
	adcq	%r11, %rax
	adcq	%rcx, %rsi
	adcq	$0, %rdx
	addq	%rdx, %r9
	adcq	%rax, %rdx
	adcq	$0, %rsi
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%r9, %rcx
	addq	$1, %rcx
	movq	%rdx, %rdi
	adcq	$1, %rdi
	movq	%rsi, %rbp
	adcq	$0, %rbp
	adcq	$-1, %rax
	andl	$1, %eax
	cmovneq	%r9, %rcx
	movq	-8(%rsp), %rbx          # 8-byte Reload
	movq	%rcx, (%rbx)
	testb	%al, %al
	cmovneq	%rdx, %rdi
	movq	%rdi, 8(%rbx)
	cmovneq	%rsi, %rbp
	movq	%rbp, 16(%rbx)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end2:
	.size	mcl_fp_sqr_NIST_P192Lbmi2, .Lfunc_end2-mcl_fp_sqr_NIST_P192Lbmi2

	.globl	mcl_fp_mulNIST_P192Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mulNIST_P192Lbmi2,@function
mcl_fp_mulNIST_P192Lbmi2:               # @mcl_fp_mulNIST_P192Lbmi2
# BB#0:
	pushq	%r14
	pushq	%rbx
	subq	$56, %rsp
	movq	%rdi, %r14
	leaq	8(%rsp), %rdi
	callq	mcl_fpDbl_mulPre3Lbmi2@PLT
	movq	24(%rsp), %r9
	movq	32(%rsp), %r8
	movq	48(%rsp), %rdi
	movq	16(%rsp), %rbx
	addq	%rdi, %rbx
	adcq	$0, %r9
	sbbq	%rcx, %rcx
	andl	$1, %ecx
	movq	40(%rsp), %rsi
	movq	8(%rsp), %rdx
	addq	%r8, %rdx
	adcq	%rsi, %rbx
	adcq	%rdi, %r9
	adcq	$0, %rcx
	addq	%rdi, %rdx
	adcq	%r8, %rbx
	adcq	%rsi, %r9
	adcq	$0, %rcx
	addq	%rcx, %rdx
	adcq	%rbx, %rcx
	adcq	$0, %r9
	sbbq	%rsi, %rsi
	andl	$1, %esi
	movq	%rdx, %rdi
	addq	$1, %rdi
	movq	%rcx, %rbx
	adcq	$1, %rbx
	movq	%r9, %rax
	adcq	$0, %rax
	adcq	$-1, %rsi
	andl	$1, %esi
	cmovneq	%rdx, %rdi
	movq	%rdi, (%r14)
	testb	%sil, %sil
	cmovneq	%rcx, %rbx
	movq	%rbx, 8(%r14)
	cmovneq	%r9, %rax
	movq	%rax, 16(%r14)
	addq	$56, %rsp
	popq	%rbx
	popq	%r14
	retq
.Lfunc_end3:
	.size	mcl_fp_mulNIST_P192Lbmi2, .Lfunc_end3-mcl_fp_mulNIST_P192Lbmi2

	.globl	mcl_fpDbl_mod_NIST_P521Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_mod_NIST_P521Lbmi2,@function
mcl_fpDbl_mod_NIST_P521Lbmi2:           # @mcl_fpDbl_mod_NIST_P521Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%r12
	pushq	%rbx
	movq	120(%rsi), %r9
	movq	128(%rsi), %r14
	movq	%r14, %r8
	shldq	$55, %r9, %r8
	movq	112(%rsi), %r10
	shldq	$55, %r10, %r9
	movq	104(%rsi), %r11
	shldq	$55, %r11, %r10
	movq	96(%rsi), %r15
	shldq	$55, %r15, %r11
	movq	88(%rsi), %r12
	shldq	$55, %r12, %r15
	movq	80(%rsi), %rcx
	shldq	$55, %rcx, %r12
	movq	64(%rsi), %rbx
	movq	72(%rsi), %rax
	shldq	$55, %rax, %rcx
	shrq	$9, %r14
	shldq	$55, %rbx, %rax
	andl	$511, %ebx              # imm = 0x1FF
	addq	(%rsi), %rax
	adcq	8(%rsi), %rcx
	adcq	16(%rsi), %r12
	adcq	24(%rsi), %r15
	adcq	32(%rsi), %r11
	adcq	40(%rsi), %r10
	adcq	48(%rsi), %r9
	adcq	56(%rsi), %r8
	adcq	%r14, %rbx
	movq	%rbx, %rsi
	shrq	$9, %rsi
	andl	$1, %esi
	addq	%rax, %rsi
	adcq	$0, %rcx
	adcq	$0, %r12
	adcq	$0, %r15
	adcq	$0, %r11
	adcq	$0, %r10
	adcq	$0, %r9
	adcq	$0, %r8
	adcq	$0, %rbx
	movq	%rsi, %rax
	andq	%r12, %rax
	andq	%r15, %rax
	andq	%r11, %rax
	andq	%r10, %rax
	andq	%r9, %rax
	andq	%r8, %rax
	movq	%rbx, %rdx
	orq	$-512, %rdx             # imm = 0xFFFFFFFFFFFFFE00
	andq	%rax, %rdx
	andq	%rcx, %rdx
	cmpq	$-1, %rdx
	je	.LBB4_1
# BB#3:                                 # %nonzero
	movq	%rsi, (%rdi)
	movq	%rcx, 8(%rdi)
	movq	%r12, 16(%rdi)
	movq	%r15, 24(%rdi)
	movq	%r11, 32(%rdi)
	movq	%r10, 40(%rdi)
	movq	%r9, 48(%rdi)
	movq	%r8, 56(%rdi)
	andl	$511, %ebx              # imm = 0x1FF
	movq	%rbx, 64(%rdi)
	jmp	.LBB4_2
.LBB4_1:                                # %zero
	movq	$0, 64(%rdi)
	movq	$0, 56(%rdi)
	movq	$0, 48(%rdi)
	movq	$0, 40(%rdi)
	movq	$0, 32(%rdi)
	movq	$0, 24(%rdi)
	movq	$0, 16(%rdi)
	movq	$0, 8(%rdi)
	movq	$0, (%rdi)
.LBB4_2:                                # %zero
	popq	%rbx
	popq	%r12
	popq	%r14
	popq	%r15
	retq
.Lfunc_end4:
	.size	mcl_fpDbl_mod_NIST_P521Lbmi2, .Lfunc_end4-mcl_fpDbl_mod_NIST_P521Lbmi2

	.globl	mcl_fp_mulUnitPre1Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mulUnitPre1Lbmi2,@function
mcl_fp_mulUnitPre1Lbmi2:                # @mcl_fp_mulUnitPre1Lbmi2
# BB#0:
	mulxq	(%rsi), %rcx, %rax
	movq	%rcx, (%rdi)
	movq	%rax, 8(%rdi)
	retq
.Lfunc_end5:
	.size	mcl_fp_mulUnitPre1Lbmi2, .Lfunc_end5-mcl_fp_mulUnitPre1Lbmi2

	.globl	mcl_fpDbl_mulPre1Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_mulPre1Lbmi2,@function
mcl_fpDbl_mulPre1Lbmi2:                 # @mcl_fpDbl_mulPre1Lbmi2
# BB#0:
	movq	(%rdx), %rdx
	mulxq	(%rsi), %rcx, %rax
	movq	%rcx, (%rdi)
	movq	%rax, 8(%rdi)
	retq
.Lfunc_end6:
	.size	mcl_fpDbl_mulPre1Lbmi2, .Lfunc_end6-mcl_fpDbl_mulPre1Lbmi2

	.globl	mcl_fpDbl_sqrPre1Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sqrPre1Lbmi2,@function
mcl_fpDbl_sqrPre1Lbmi2:                 # @mcl_fpDbl_sqrPre1Lbmi2
# BB#0:
	movq	(%rsi), %rdx
	mulxq	%rdx, %rcx, %rax
	movq	%rcx, (%rdi)
	movq	%rax, 8(%rdi)
	retq
.Lfunc_end7:
	.size	mcl_fpDbl_sqrPre1Lbmi2, .Lfunc_end7-mcl_fpDbl_sqrPre1Lbmi2

	.globl	mcl_fp_mont1Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mont1Lbmi2,@function
mcl_fp_mont1Lbmi2:                      # @mcl_fp_mont1Lbmi2
# BB#0:
	movq	%rdx, %rax
	movq	(%rsi), %rdx
	mulxq	(%rax), %rsi, %r8
	movq	-8(%rcx), %rdx
	imulq	%rsi, %rdx
	movq	(%rcx), %rcx
	mulxq	%rcx, %rdx, %rax
	addq	%rsi, %rdx
	adcq	%r8, %rax
	sbbq	%rdx, %rdx
	andl	$1, %edx
	movq	%rax, %rsi
	subq	%rcx, %rsi
	sbbq	$0, %rdx
	testb	$1, %dl
	cmovneq	%rax, %rsi
	movq	%rsi, (%rdi)
	retq
.Lfunc_end8:
	.size	mcl_fp_mont1Lbmi2, .Lfunc_end8-mcl_fp_mont1Lbmi2

	.globl	mcl_fp_montNF1Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montNF1Lbmi2,@function
mcl_fp_montNF1Lbmi2:                    # @mcl_fp_montNF1Lbmi2
# BB#0:
	movq	%rdx, %rax
	movq	(%rsi), %rdx
	mulxq	(%rax), %rsi, %r8
	movq	-8(%rcx), %rdx
	imulq	%rsi, %rdx
	movq	(%rcx), %rcx
	mulxq	%rcx, %rdx, %rax
	addq	%rsi, %rdx
	adcq	%r8, %rax
	movq	%rax, %rdx
	subq	%rcx, %rdx
	cmovsq	%rax, %rdx
	movq	%rdx, (%rdi)
	retq
.Lfunc_end9:
	.size	mcl_fp_montNF1Lbmi2, .Lfunc_end9-mcl_fp_montNF1Lbmi2

	.globl	mcl_fp_montRed1Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montRed1Lbmi2,@function
mcl_fp_montRed1Lbmi2:                   # @mcl_fp_montRed1Lbmi2
# BB#0:
	movq	(%rsi), %rcx
	movq	-8(%rdx), %rax
	imulq	%rcx, %rax
	movq	(%rdx), %r8
	movq	%rax, %rdx
	mulxq	%r8, %rax, %rdx
	addq	%rcx, %rax
	adcq	8(%rsi), %rdx
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%rdx, %rcx
	subq	%r8, %rcx
	sbbq	$0, %rax
	testb	$1, %al
	cmovneq	%rdx, %rcx
	movq	%rcx, (%rdi)
	retq
.Lfunc_end10:
	.size	mcl_fp_montRed1Lbmi2, .Lfunc_end10-mcl_fp_montRed1Lbmi2

	.globl	mcl_fp_addPre1Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addPre1Lbmi2,@function
mcl_fp_addPre1Lbmi2:                    # @mcl_fp_addPre1Lbmi2
# BB#0:
	movq	(%rdx), %rax
	addq	(%rsi), %rax
	movq	%rax, (%rdi)
	sbbq	%rax, %rax
	andl	$1, %eax
	retq
.Lfunc_end11:
	.size	mcl_fp_addPre1Lbmi2, .Lfunc_end11-mcl_fp_addPre1Lbmi2

	.globl	mcl_fp_subPre1Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subPre1Lbmi2,@function
mcl_fp_subPre1Lbmi2:                    # @mcl_fp_subPre1Lbmi2
# BB#0:
	movq	(%rsi), %rcx
	xorl	%eax, %eax
	subq	(%rdx), %rcx
	movq	%rcx, (%rdi)
	sbbq	$0, %rax
	andl	$1, %eax
	retq
.Lfunc_end12:
	.size	mcl_fp_subPre1Lbmi2, .Lfunc_end12-mcl_fp_subPre1Lbmi2

	.globl	mcl_fp_shr1_1Lbmi2
	.align	16, 0x90
	.type	mcl_fp_shr1_1Lbmi2,@function
mcl_fp_shr1_1Lbmi2:                     # @mcl_fp_shr1_1Lbmi2
# BB#0:
	movq	(%rsi), %rax
	shrq	%rax
	movq	%rax, (%rdi)
	retq
.Lfunc_end13:
	.size	mcl_fp_shr1_1Lbmi2, .Lfunc_end13-mcl_fp_shr1_1Lbmi2

	.globl	mcl_fp_add1Lbmi2
	.align	16, 0x90
	.type	mcl_fp_add1Lbmi2,@function
mcl_fp_add1Lbmi2:                       # @mcl_fp_add1Lbmi2
# BB#0:
	movq	(%rdx), %rax
	addq	(%rsi), %rax
	movq	%rax, (%rdi)
	sbbq	%rdx, %rdx
	andl	$1, %edx
	subq	(%rcx), %rax
	sbbq	$0, %rdx
	testb	$1, %dl
	jne	.LBB14_2
# BB#1:                                 # %nocarry
	movq	%rax, (%rdi)
.LBB14_2:                               # %carry
	retq
.Lfunc_end14:
	.size	mcl_fp_add1Lbmi2, .Lfunc_end14-mcl_fp_add1Lbmi2

	.globl	mcl_fp_addNF1Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addNF1Lbmi2,@function
mcl_fp_addNF1Lbmi2:                     # @mcl_fp_addNF1Lbmi2
# BB#0:
	movq	(%rdx), %rax
	addq	(%rsi), %rax
	movq	%rax, %rdx
	subq	(%rcx), %rdx
	cmovsq	%rax, %rdx
	movq	%rdx, (%rdi)
	retq
.Lfunc_end15:
	.size	mcl_fp_addNF1Lbmi2, .Lfunc_end15-mcl_fp_addNF1Lbmi2

	.globl	mcl_fp_sub1Lbmi2
	.align	16, 0x90
	.type	mcl_fp_sub1Lbmi2,@function
mcl_fp_sub1Lbmi2:                       # @mcl_fp_sub1Lbmi2
# BB#0:
	movq	(%rsi), %rax
	xorl	%esi, %esi
	subq	(%rdx), %rax
	movq	%rax, (%rdi)
	sbbq	$0, %rsi
	testb	$1, %sil
	jne	.LBB16_2
# BB#1:                                 # %nocarry
	retq
.LBB16_2:                               # %carry
	addq	(%rcx), %rax
	movq	%rax, (%rdi)
	retq
.Lfunc_end16:
	.size	mcl_fp_sub1Lbmi2, .Lfunc_end16-mcl_fp_sub1Lbmi2

	.globl	mcl_fp_subNF1Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subNF1Lbmi2,@function
mcl_fp_subNF1Lbmi2:                     # @mcl_fp_subNF1Lbmi2
# BB#0:
	movq	(%rsi), %rax
	subq	(%rdx), %rax
	movq	%rax, %rdx
	sarq	$63, %rdx
	andq	(%rcx), %rdx
	addq	%rax, %rdx
	movq	%rdx, (%rdi)
	retq
.Lfunc_end17:
	.size	mcl_fp_subNF1Lbmi2, .Lfunc_end17-mcl_fp_subNF1Lbmi2

	.globl	mcl_fpDbl_add1Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_add1Lbmi2,@function
mcl_fpDbl_add1Lbmi2:                    # @mcl_fpDbl_add1Lbmi2
# BB#0:
	movq	(%rdx), %rax
	movq	8(%rdx), %rdx
	addq	(%rsi), %rax
	adcq	8(%rsi), %rdx
	movq	%rax, (%rdi)
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%rdx, %rsi
	subq	(%rcx), %rsi
	sbbq	$0, %rax
	testb	$1, %al
	cmovneq	%rdx, %rsi
	movq	%rsi, 8(%rdi)
	retq
.Lfunc_end18:
	.size	mcl_fpDbl_add1Lbmi2, .Lfunc_end18-mcl_fpDbl_add1Lbmi2

	.globl	mcl_fpDbl_sub1Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sub1Lbmi2,@function
mcl_fpDbl_sub1Lbmi2:                    # @mcl_fpDbl_sub1Lbmi2
# BB#0:
	movq	(%rsi), %rax
	movq	8(%rsi), %r8
	xorl	%esi, %esi
	subq	(%rdx), %rax
	sbbq	8(%rdx), %r8
	movq	%rax, (%rdi)
	movl	$0, %eax
	sbbq	$0, %rax
	testb	$1, %al
	cmovneq	(%rcx), %rsi
	addq	%r8, %rsi
	movq	%rsi, 8(%rdi)
	retq
.Lfunc_end19:
	.size	mcl_fpDbl_sub1Lbmi2, .Lfunc_end19-mcl_fpDbl_sub1Lbmi2

	.globl	mcl_fp_mulUnitPre2Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mulUnitPre2Lbmi2,@function
mcl_fp_mulUnitPre2Lbmi2:                # @mcl_fp_mulUnitPre2Lbmi2
# BB#0:
	mulxq	8(%rsi), %rax, %rcx
	mulxq	(%rsi), %rdx, %rsi
	movq	%rdx, (%rdi)
	addq	%rax, %rsi
	movq	%rsi, 8(%rdi)
	adcq	$0, %rcx
	movq	%rcx, 16(%rdi)
	retq
.Lfunc_end20:
	.size	mcl_fp_mulUnitPre2Lbmi2, .Lfunc_end20-mcl_fp_mulUnitPre2Lbmi2

	.globl	mcl_fpDbl_mulPre2Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_mulPre2Lbmi2,@function
mcl_fpDbl_mulPre2Lbmi2:                 # @mcl_fpDbl_mulPre2Lbmi2
# BB#0:
	movq	%rdx, %r10
	movq	(%rsi), %r11
	movq	8(%rsi), %r8
	movq	(%r10), %rsi
	movq	%r11, %rdx
	mulxq	%rsi, %rdx, %r9
	movq	%rdx, (%rdi)
	movq	%r8, %rdx
	mulxq	%rsi, %rsi, %rax
	addq	%r9, %rsi
	adcq	$0, %rax
	movq	8(%r10), %rcx
	movq	%r11, %rdx
	mulxq	%rcx, %rdx, %r9
	addq	%rsi, %rdx
	movq	%rdx, 8(%rdi)
	movq	%r8, %rdx
	mulxq	%rcx, %rdx, %rcx
	adcq	%rax, %rdx
	sbbq	%rax, %rax
	andl	$1, %eax
	addq	%r9, %rdx
	movq	%rdx, 16(%rdi)
	adcq	%rcx, %rax
	movq	%rax, 24(%rdi)
	retq
.Lfunc_end21:
	.size	mcl_fpDbl_mulPre2Lbmi2, .Lfunc_end21-mcl_fpDbl_mulPre2Lbmi2

	.globl	mcl_fpDbl_sqrPre2Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sqrPre2Lbmi2,@function
mcl_fpDbl_sqrPre2Lbmi2:                 # @mcl_fpDbl_sqrPre2Lbmi2
# BB#0:
	movq	(%rsi), %rax
	movq	8(%rsi), %rcx
	movq	%rax, %rdx
	mulxq	%rax, %rdx, %rsi
	movq	%rdx, (%rdi)
	movq	%rcx, %rdx
	mulxq	%rax, %rdx, %r8
	addq	%rdx, %rsi
	movq	%r8, %rax
	adcq	$0, %rax
	addq	%rdx, %rsi
	movq	%rsi, 8(%rdi)
	movq	%rcx, %rdx
	mulxq	%rcx, %rdx, %rcx
	adcq	%rax, %rdx
	sbbq	%rax, %rax
	andl	$1, %eax
	addq	%r8, %rdx
	movq	%rdx, 16(%rdi)
	adcq	%rcx, %rax
	movq	%rax, 24(%rdi)
	retq
.Lfunc_end22:
	.size	mcl_fpDbl_sqrPre2Lbmi2, .Lfunc_end22-mcl_fpDbl_sqrPre2Lbmi2

	.globl	mcl_fp_mont2Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mont2Lbmi2,@function
mcl_fp_mont2Lbmi2:                      # @mcl_fp_mont2Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	(%rsi), %r8
	movq	8(%rsi), %r9
	movq	(%rdx), %rax
	movq	8(%rdx), %r11
	movq	%r9, %rdx
	mulxq	%rax, %r10, %r13
	movq	%r8, %rdx
	mulxq	%rax, %r14, %rsi
	addq	%r10, %rsi
	adcq	$0, %r13
	movq	-8(%rcx), %rbp
	movq	(%rcx), %r10
	movq	%r14, %rdx
	imulq	%rbp, %rdx
	movq	8(%rcx), %r15
	mulxq	%r15, %r12, %rcx
	mulxq	%r10, %rdx, %rbx
	addq	%r12, %rbx
	adcq	$0, %rcx
	addq	%r14, %rdx
	adcq	%rsi, %rbx
	adcq	%r13, %rcx
	sbbq	%rsi, %rsi
	andl	$1, %esi
	movq	%r11, %rdx
	mulxq	%r9, %r9, %r14
	movq	%r11, %rdx
	mulxq	%r8, %r8, %rax
	addq	%r9, %rax
	adcq	$0, %r14
	addq	%rbx, %r8
	adcq	%rcx, %rax
	adcq	%rsi, %r14
	sbbq	%rsi, %rsi
	andl	$1, %esi
	imulq	%r8, %rbp
	movq	%rbp, %rdx
	mulxq	%r15, %rcx, %rbx
	movq	%rbp, %rdx
	mulxq	%r10, %rdx, %rbp
	addq	%rcx, %rbp
	adcq	$0, %rbx
	addq	%r8, %rdx
	adcq	%rax, %rbp
	adcq	%r14, %rbx
	adcq	$0, %rsi
	movq	%rbp, %rax
	subq	%r10, %rax
	movq	%rbx, %rcx
	sbbq	%r15, %rcx
	sbbq	$0, %rsi
	andl	$1, %esi
	cmovneq	%rbx, %rcx
	testb	%sil, %sil
	cmovneq	%rbp, %rax
	movq	%rax, (%rdi)
	movq	%rcx, 8(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end23:
	.size	mcl_fp_mont2Lbmi2, .Lfunc_end23-mcl_fp_mont2Lbmi2

	.globl	mcl_fp_montNF2Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montNF2Lbmi2,@function
mcl_fp_montNF2Lbmi2:                    # @mcl_fp_montNF2Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	(%rsi), %r8
	movq	8(%rsi), %r9
	movq	(%rdx), %rax
	movq	8(%rdx), %r11
	movq	%r9, %rdx
	mulxq	%rax, %r10, %rsi
	movq	%r8, %rdx
	mulxq	%rax, %r15, %r13
	addq	%r10, %r13
	adcq	$0, %rsi
	movq	-8(%rcx), %rbp
	movq	(%rcx), %r10
	movq	%r15, %rdx
	imulq	%rbp, %rdx
	movq	8(%rcx), %r14
	mulxq	%r10, %rcx, %r12
	addq	%r15, %rcx
	mulxq	%r14, %rbx, %rcx
	adcq	%r13, %rbx
	adcq	$0, %rsi
	addq	%r12, %rbx
	adcq	%rcx, %rsi
	movq	%r11, %rdx
	mulxq	%r9, %r9, %rcx
	movq	%r11, %rdx
	mulxq	%r8, %r8, %rax
	addq	%r9, %rax
	adcq	$0, %rcx
	addq	%rbx, %r8
	adcq	%rsi, %rax
	adcq	$0, %rcx
	imulq	%r8, %rbp
	movq	%rbp, %rdx
	mulxq	%r14, %rbx, %rsi
	movq	%rbp, %rdx
	mulxq	%r10, %rbp, %rdx
	addq	%r8, %rbp
	adcq	%rax, %rbx
	adcq	$0, %rcx
	addq	%rdx, %rbx
	adcq	%rsi, %rcx
	movq	%rbx, %rax
	subq	%r10, %rax
	movq	%rcx, %rdx
	sbbq	%r14, %rdx
	cmovsq	%rbx, %rax
	movq	%rax, (%rdi)
	cmovsq	%rcx, %rdx
	movq	%rdx, 8(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end24:
	.size	mcl_fp_montNF2Lbmi2, .Lfunc_end24-mcl_fp_montNF2Lbmi2

	.globl	mcl_fp_montRed2Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montRed2Lbmi2,@function
mcl_fp_montRed2Lbmi2:                   # @mcl_fp_montRed2Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%rbx
	movq	-8(%rdx), %r15
	movq	(%rdx), %r8
	movq	(%rsi), %r10
	movq	%r10, %rcx
	imulq	%r15, %rcx
	movq	8(%rdx), %r9
	movq	%rcx, %rdx
	mulxq	%r9, %r11, %r14
	movq	%rcx, %rdx
	mulxq	%r8, %rcx, %rax
	addq	%r11, %rax
	adcq	$0, %r14
	movq	24(%rsi), %r11
	addq	%r10, %rcx
	adcq	8(%rsi), %rax
	adcq	16(%rsi), %r14
	adcq	$0, %r11
	sbbq	%rcx, %rcx
	andl	$1, %ecx
	imulq	%rax, %r15
	movq	%r15, %rdx
	mulxq	%r9, %r10, %rbx
	movq	%r15, %rdx
	mulxq	%r8, %rsi, %rdx
	addq	%r10, %rdx
	adcq	$0, %rbx
	addq	%rax, %rsi
	adcq	%r14, %rdx
	adcq	%r11, %rbx
	adcq	$0, %rcx
	movq	%rdx, %rax
	subq	%r8, %rax
	movq	%rbx, %rsi
	sbbq	%r9, %rsi
	sbbq	$0, %rcx
	andl	$1, %ecx
	cmovneq	%rbx, %rsi
	testb	%cl, %cl
	cmovneq	%rdx, %rax
	movq	%rax, (%rdi)
	movq	%rsi, 8(%rdi)
	popq	%rbx
	popq	%r14
	popq	%r15
	retq
.Lfunc_end25:
	.size	mcl_fp_montRed2Lbmi2, .Lfunc_end25-mcl_fp_montRed2Lbmi2

	.globl	mcl_fp_addPre2Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addPre2Lbmi2,@function
mcl_fp_addPre2Lbmi2:                    # @mcl_fp_addPre2Lbmi2
# BB#0:
	movq	(%rdx), %rax
	movq	8(%rdx), %rcx
	addq	(%rsi), %rax
	adcq	8(%rsi), %rcx
	movq	%rax, (%rdi)
	movq	%rcx, 8(%rdi)
	sbbq	%rax, %rax
	andl	$1, %eax
	retq
.Lfunc_end26:
	.size	mcl_fp_addPre2Lbmi2, .Lfunc_end26-mcl_fp_addPre2Lbmi2

	.globl	mcl_fp_subPre2Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subPre2Lbmi2,@function
mcl_fp_subPre2Lbmi2:                    # @mcl_fp_subPre2Lbmi2
# BB#0:
	movq	(%rsi), %rcx
	movq	8(%rsi), %rsi
	xorl	%eax, %eax
	subq	(%rdx), %rcx
	sbbq	8(%rdx), %rsi
	movq	%rcx, (%rdi)
	movq	%rsi, 8(%rdi)
	sbbq	$0, %rax
	andl	$1, %eax
	retq
.Lfunc_end27:
	.size	mcl_fp_subPre2Lbmi2, .Lfunc_end27-mcl_fp_subPre2Lbmi2

	.globl	mcl_fp_shr1_2Lbmi2
	.align	16, 0x90
	.type	mcl_fp_shr1_2Lbmi2,@function
mcl_fp_shr1_2Lbmi2:                     # @mcl_fp_shr1_2Lbmi2
# BB#0:
	movq	(%rsi), %rax
	movq	8(%rsi), %rcx
	shrdq	$1, %rcx, %rax
	movq	%rax, (%rdi)
	shrq	%rcx
	movq	%rcx, 8(%rdi)
	retq
.Lfunc_end28:
	.size	mcl_fp_shr1_2Lbmi2, .Lfunc_end28-mcl_fp_shr1_2Lbmi2

	.globl	mcl_fp_add2Lbmi2
	.align	16, 0x90
	.type	mcl_fp_add2Lbmi2,@function
mcl_fp_add2Lbmi2:                       # @mcl_fp_add2Lbmi2
# BB#0:
	movq	(%rdx), %rax
	movq	8(%rdx), %rdx
	addq	(%rsi), %rax
	adcq	8(%rsi), %rdx
	movq	%rax, (%rdi)
	movq	%rdx, 8(%rdi)
	sbbq	%rsi, %rsi
	andl	$1, %esi
	subq	(%rcx), %rax
	sbbq	8(%rcx), %rdx
	sbbq	$0, %rsi
	testb	$1, %sil
	jne	.LBB29_2
# BB#1:                                 # %nocarry
	movq	%rax, (%rdi)
	movq	%rdx, 8(%rdi)
.LBB29_2:                               # %carry
	retq
.Lfunc_end29:
	.size	mcl_fp_add2Lbmi2, .Lfunc_end29-mcl_fp_add2Lbmi2

	.globl	mcl_fp_addNF2Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addNF2Lbmi2,@function
mcl_fp_addNF2Lbmi2:                     # @mcl_fp_addNF2Lbmi2
# BB#0:
	movq	(%rdx), %rax
	movq	8(%rdx), %r8
	addq	(%rsi), %rax
	adcq	8(%rsi), %r8
	movq	%rax, %rsi
	subq	(%rcx), %rsi
	movq	%r8, %rdx
	sbbq	8(%rcx), %rdx
	testq	%rdx, %rdx
	cmovsq	%rax, %rsi
	movq	%rsi, (%rdi)
	cmovsq	%r8, %rdx
	movq	%rdx, 8(%rdi)
	retq
.Lfunc_end30:
	.size	mcl_fp_addNF2Lbmi2, .Lfunc_end30-mcl_fp_addNF2Lbmi2

	.globl	mcl_fp_sub2Lbmi2
	.align	16, 0x90
	.type	mcl_fp_sub2Lbmi2,@function
mcl_fp_sub2Lbmi2:                       # @mcl_fp_sub2Lbmi2
# BB#0:
	movq	(%rsi), %rax
	movq	8(%rsi), %r8
	xorl	%esi, %esi
	subq	(%rdx), %rax
	sbbq	8(%rdx), %r8
	movq	%rax, (%rdi)
	movq	%r8, 8(%rdi)
	sbbq	$0, %rsi
	testb	$1, %sil
	jne	.LBB31_2
# BB#1:                                 # %nocarry
	retq
.LBB31_2:                               # %carry
	movq	8(%rcx), %rdx
	addq	(%rcx), %rax
	movq	%rax, (%rdi)
	adcq	%r8, %rdx
	movq	%rdx, 8(%rdi)
	retq
.Lfunc_end31:
	.size	mcl_fp_sub2Lbmi2, .Lfunc_end31-mcl_fp_sub2Lbmi2

	.globl	mcl_fp_subNF2Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subNF2Lbmi2,@function
mcl_fp_subNF2Lbmi2:                     # @mcl_fp_subNF2Lbmi2
# BB#0:
	movq	(%rsi), %r8
	movq	8(%rsi), %rsi
	subq	(%rdx), %r8
	sbbq	8(%rdx), %rsi
	movq	%rsi, %rdx
	sarq	$63, %rdx
	movq	8(%rcx), %rax
	andq	%rdx, %rax
	andq	(%rcx), %rdx
	addq	%r8, %rdx
	movq	%rdx, (%rdi)
	adcq	%rsi, %rax
	movq	%rax, 8(%rdi)
	retq
.Lfunc_end32:
	.size	mcl_fp_subNF2Lbmi2, .Lfunc_end32-mcl_fp_subNF2Lbmi2

	.globl	mcl_fpDbl_add2Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_add2Lbmi2,@function
mcl_fpDbl_add2Lbmi2:                    # @mcl_fpDbl_add2Lbmi2
# BB#0:
	movq	24(%rdx), %r8
	movq	24(%rsi), %r9
	movq	16(%rdx), %r10
	movq	(%rdx), %rax
	movq	8(%rdx), %rdx
	addq	(%rsi), %rax
	adcq	8(%rsi), %rdx
	adcq	16(%rsi), %r10
	movq	%rax, (%rdi)
	movq	%rdx, 8(%rdi)
	adcq	%r8, %r9
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%r10, %rdx
	subq	(%rcx), %rdx
	movq	%r9, %rsi
	sbbq	8(%rcx), %rsi
	sbbq	$0, %rax
	andl	$1, %eax
	cmovneq	%r10, %rdx
	movq	%rdx, 16(%rdi)
	testb	%al, %al
	cmovneq	%r9, %rsi
	movq	%rsi, 24(%rdi)
	retq
.Lfunc_end33:
	.size	mcl_fpDbl_add2Lbmi2, .Lfunc_end33-mcl_fpDbl_add2Lbmi2

	.globl	mcl_fpDbl_sub2Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sub2Lbmi2,@function
mcl_fpDbl_sub2Lbmi2:                    # @mcl_fpDbl_sub2Lbmi2
# BB#0:
	movq	24(%rdx), %r8
	movq	24(%rsi), %r9
	movq	16(%rsi), %r10
	movq	(%rsi), %r11
	movq	8(%rsi), %rsi
	xorl	%eax, %eax
	subq	(%rdx), %r11
	sbbq	8(%rdx), %rsi
	sbbq	16(%rdx), %r10
	movq	%r11, (%rdi)
	movq	%rsi, 8(%rdi)
	sbbq	%r8, %r9
	movl	$0, %edx
	sbbq	$0, %rdx
	andl	$1, %edx
	movq	(%rcx), %rsi
	cmoveq	%rax, %rsi
	testb	%dl, %dl
	cmovneq	8(%rcx), %rax
	addq	%r10, %rsi
	movq	%rsi, 16(%rdi)
	adcq	%r9, %rax
	movq	%rax, 24(%rdi)
	retq
.Lfunc_end34:
	.size	mcl_fpDbl_sub2Lbmi2, .Lfunc_end34-mcl_fpDbl_sub2Lbmi2

	.globl	mcl_fp_mulUnitPre3Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mulUnitPre3Lbmi2,@function
mcl_fp_mulUnitPre3Lbmi2:                # @mcl_fp_mulUnitPre3Lbmi2
# BB#0:
	mulxq	16(%rsi), %r8, %rcx
	mulxq	8(%rsi), %r9, %rax
	mulxq	(%rsi), %rdx, %rsi
	movq	%rdx, (%rdi)
	addq	%r9, %rsi
	movq	%rsi, 8(%rdi)
	adcq	%r8, %rax
	movq	%rax, 16(%rdi)
	adcq	$0, %rcx
	movq	%rcx, 24(%rdi)
	retq
.Lfunc_end35:
	.size	mcl_fp_mulUnitPre3Lbmi2, .Lfunc_end35-mcl_fp_mulUnitPre3Lbmi2

	.globl	mcl_fpDbl_mulPre3Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_mulPre3Lbmi2,@function
mcl_fpDbl_mulPre3Lbmi2:                 # @mcl_fpDbl_mulPre3Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%rbx
	movq	%rdx, %r9
	movq	(%rsi), %r10
	movq	8(%rsi), %r8
	movq	(%r9), %rax
	movq	%r10, %rdx
	mulxq	%rax, %rdx, %r14
	movq	16(%rsi), %r11
	movq	%rdx, (%rdi)
	movq	%r11, %rdx
	mulxq	%rax, %rsi, %rbx
	movq	%r8, %rdx
	mulxq	%rax, %rax, %rcx
	addq	%r14, %rax
	adcq	%rsi, %rcx
	adcq	$0, %rbx
	movq	8(%r9), %rsi
	movq	%r10, %rdx
	mulxq	%rsi, %rdx, %r14
	addq	%rax, %rdx
	movq	%rdx, 8(%rdi)
	movq	%r11, %rdx
	mulxq	%rsi, %rax, %r15
	movq	%r8, %rdx
	mulxq	%rsi, %rsi, %rdx
	adcq	%rcx, %rsi
	adcq	%rbx, %rax
	sbbq	%rcx, %rcx
	andl	$1, %ecx
	addq	%r14, %rsi
	adcq	%rdx, %rax
	adcq	%r15, %rcx
	movq	16(%r9), %rbx
	movq	%r10, %rdx
	mulxq	%rbx, %rdx, %r9
	addq	%rsi, %rdx
	movq	%rdx, 16(%rdi)
	movq	%r11, %rdx
	mulxq	%rbx, %rsi, %r10
	movq	%r8, %rdx
	mulxq	%rbx, %rbx, %rdx
	adcq	%rax, %rbx
	adcq	%rcx, %rsi
	sbbq	%rax, %rax
	andl	$1, %eax
	addq	%r9, %rbx
	movq	%rbx, 24(%rdi)
	adcq	%rdx, %rsi
	movq	%rsi, 32(%rdi)
	adcq	%r10, %rax
	movq	%rax, 40(%rdi)
	popq	%rbx
	popq	%r14
	popq	%r15
	retq
.Lfunc_end36:
	.size	mcl_fpDbl_mulPre3Lbmi2, .Lfunc_end36-mcl_fpDbl_mulPre3Lbmi2

	.globl	mcl_fpDbl_sqrPre3Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sqrPre3Lbmi2,@function
mcl_fpDbl_sqrPre3Lbmi2:                 # @mcl_fpDbl_sqrPre3Lbmi2
# BB#0:
	pushq	%r14
	pushq	%rbx
	movq	16(%rsi), %r10
	movq	(%rsi), %rcx
	movq	8(%rsi), %rsi
	movq	%rcx, %rdx
	mulxq	%rcx, %rdx, %rax
	movq	%rdx, (%rdi)
	movq	%r10, %rdx
	mulxq	%rcx, %r11, %r8
	movq	%rsi, %rdx
	mulxq	%rcx, %rdx, %r14
	addq	%rdx, %rax
	movq	%r14, %rbx
	adcq	%r11, %rbx
	movq	%r8, %rcx
	adcq	$0, %rcx
	addq	%rdx, %rax
	movq	%rax, 8(%rdi)
	movq	%r10, %rdx
	mulxq	%rsi, %rax, %r9
	movq	%rsi, %rdx
	mulxq	%rsi, %rsi, %rdx
	adcq	%rbx, %rsi
	adcq	%rax, %rcx
	sbbq	%rbx, %rbx
	andl	$1, %ebx
	addq	%r14, %rsi
	adcq	%rdx, %rcx
	adcq	%r9, %rbx
	addq	%r11, %rsi
	movq	%rsi, 16(%rdi)
	movq	%r10, %rdx
	mulxq	%r10, %rsi, %rdx
	adcq	%rax, %rcx
	adcq	%rbx, %rsi
	sbbq	%rax, %rax
	andl	$1, %eax
	addq	%r8, %rcx
	movq	%rcx, 24(%rdi)
	adcq	%r9, %rsi
	movq	%rsi, 32(%rdi)
	adcq	%rdx, %rax
	movq	%rax, 40(%rdi)
	popq	%rbx
	popq	%r14
	retq
.Lfunc_end37:
	.size	mcl_fpDbl_sqrPre3Lbmi2, .Lfunc_end37-mcl_fpDbl_sqrPre3Lbmi2

	.globl	mcl_fp_mont3Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mont3Lbmi2,@function
mcl_fp_mont3Lbmi2:                      # @mcl_fp_mont3Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rdx, %r15
	movq	%r15, -32(%rsp)         # 8-byte Spill
	movq	%rdi, -24(%rsp)         # 8-byte Spill
	movq	16(%rsi), %rdi
	movq	%rdi, -56(%rsp)         # 8-byte Spill
	movq	(%r15), %rax
	movq	%rdi, %rdx
	mulxq	%rax, %r14, %r11
	movq	(%rsi), %r12
	movq	%r12, -48(%rsp)         # 8-byte Spill
	movq	8(%rsi), %rdx
	movq	%rdx, -16(%rsp)         # 8-byte Spill
	mulxq	%rax, %rbx, %r8
	movq	%r12, %rdx
	mulxq	%rax, %r9, %rdi
	addq	%rbx, %rdi
	adcq	%r14, %r8
	adcq	$0, %r11
	movq	-8(%rcx), %r13
	movq	(%rcx), %rbx
	movq	%rbx, -8(%rsp)          # 8-byte Spill
	movq	%r9, %rdx
	imulq	%r13, %rdx
	movq	8(%rcx), %rax
	movq	%rax, -40(%rsp)         # 8-byte Spill
	mulxq	%rax, %rax, %r10
	mulxq	%rbx, %rsi, %rbx
	addq	%rax, %rbx
	movq	16(%rcx), %rbp
	mulxq	%rbp, %rcx, %rax
	movq	%rbp, %r14
	adcq	%r10, %rcx
	adcq	$0, %rax
	addq	%r9, %rsi
	adcq	%rdi, %rbx
	movq	8(%r15), %rdx
	adcq	%r8, %rcx
	adcq	%r11, %rax
	sbbq	%r9, %r9
	andl	$1, %r9d
	movq	-56(%rsp), %r15         # 8-byte Reload
	mulxq	%r15, %r11, %rdi
	mulxq	-16(%rsp), %r10, %rsi   # 8-byte Folded Reload
	mulxq	%r12, %r8, %rbp
	addq	%r10, %rbp
	adcq	%r11, %rsi
	adcq	$0, %rdi
	addq	%rbx, %r8
	adcq	%rcx, %rbp
	adcq	%rax, %rsi
	adcq	%r9, %rdi
	sbbq	%r11, %r11
	andl	$1, %r11d
	movq	%r8, %rdx
	imulq	%r13, %rdx
	mulxq	%r14, %r9, %rcx
	movq	%r14, %r12
	movq	-40(%rsp), %r14         # 8-byte Reload
	mulxq	%r14, %r10, %rax
	mulxq	-8(%rsp), %rdx, %rbx    # 8-byte Folded Reload
	addq	%r10, %rbx
	adcq	%r9, %rax
	adcq	$0, %rcx
	addq	%r8, %rdx
	adcq	%rbp, %rbx
	adcq	%rsi, %rax
	adcq	%rdi, %rcx
	adcq	$0, %r11
	movq	-32(%rsp), %rdx         # 8-byte Reload
	movq	16(%rdx), %rdx
	mulxq	%r15, %r9, %rsi
	mulxq	-16(%rsp), %r10, %r15   # 8-byte Folded Reload
	mulxq	-48(%rsp), %r8, %rdi    # 8-byte Folded Reload
	addq	%r10, %rdi
	adcq	%r9, %r15
	adcq	$0, %rsi
	addq	%rbx, %r8
	adcq	%rax, %rdi
	adcq	%rcx, %r15
	adcq	%r11, %rsi
	sbbq	%rbx, %rbx
	andl	$1, %ebx
	imulq	%r8, %r13
	movq	%r13, %rdx
	mulxq	%r12, %r9, %rbp
	movq	%r13, %rdx
	mulxq	%r14, %r10, %rax
	movq	%r13, %rdx
	movq	-8(%rsp), %rcx          # 8-byte Reload
	mulxq	%rcx, %r11, %rdx
	addq	%r10, %rdx
	adcq	%r9, %rax
	adcq	$0, %rbp
	addq	%r8, %r11
	adcq	%rdi, %rdx
	adcq	%r15, %rax
	adcq	%rsi, %rbp
	adcq	$0, %rbx
	movq	%rdx, %rsi
	subq	%rcx, %rsi
	movq	%rax, %rdi
	sbbq	%r14, %rdi
	movq	%rbp, %rcx
	sbbq	%r12, %rcx
	sbbq	$0, %rbx
	andl	$1, %ebx
	cmovneq	%rbp, %rcx
	testb	%bl, %bl
	cmovneq	%rdx, %rsi
	movq	-24(%rsp), %rdx         # 8-byte Reload
	movq	%rsi, (%rdx)
	cmovneq	%rax, %rdi
	movq	%rdi, 8(%rdx)
	movq	%rcx, 16(%rdx)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end38:
	.size	mcl_fp_mont3Lbmi2, .Lfunc_end38-mcl_fp_mont3Lbmi2

	.globl	mcl_fp_montNF3Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montNF3Lbmi2,@function
mcl_fp_montNF3Lbmi2:                    # @mcl_fp_montNF3Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rcx, %r8
	movq	%rdx, %r10
	movq	%r10, -16(%rsp)         # 8-byte Spill
	movq	%rdi, -8(%rsp)          # 8-byte Spill
	movq	(%rsi), %rcx
	movq	%rcx, -24(%rsp)         # 8-byte Spill
	movq	8(%rsi), %rdi
	movq	%rdi, -32(%rsp)         # 8-byte Spill
	movq	(%r10), %rax
	movq	%rdi, %rdx
	mulxq	%rax, %rbx, %r14
	movq	%rcx, %rdx
	mulxq	%rax, %r15, %r12
	movq	16(%rsi), %r11
	addq	%rbx, %r12
	movq	%r11, %rdx
	mulxq	%rax, %rsi, %rbx
	adcq	%r14, %rsi
	adcq	$0, %rbx
	movq	-8(%r8), %r9
	movq	(%r8), %r14
	movq	%r15, %rdx
	imulq	%r9, %rdx
	mulxq	%r14, %rbp, %r13
	addq	%r15, %rbp
	movq	8(%r8), %r15
	mulxq	%r15, %rdi, %rbp
	adcq	%r12, %rdi
	movq	16(%r8), %r12
	mulxq	%r12, %rax, %r8
	adcq	%rsi, %rax
	adcq	$0, %rbx
	addq	%r13, %rdi
	movq	8(%r10), %rdx
	adcq	%rbp, %rax
	adcq	%r8, %rbx
	movq	-32(%rsp), %r10         # 8-byte Reload
	mulxq	%r10, %rsi, %r8
	mulxq	%rcx, %r13, %rbp
	addq	%rsi, %rbp
	mulxq	%r11, %rcx, %rsi
	adcq	%r8, %rcx
	adcq	$0, %rsi
	addq	%rdi, %r13
	adcq	%rax, %rbp
	adcq	%rbx, %rcx
	adcq	$0, %rsi
	movq	%r13, %rdx
	imulq	%r9, %rdx
	mulxq	%r14, %rdi, %rbx
	addq	%r13, %rdi
	mulxq	%r15, %rax, %rdi
	adcq	%rbp, %rax
	mulxq	%r12, %rbp, %rdx
	adcq	%rcx, %rbp
	adcq	$0, %rsi
	addq	%rbx, %rax
	adcq	%rdi, %rbp
	adcq	%rdx, %rsi
	movq	-16(%rsp), %rcx         # 8-byte Reload
	movq	16(%rcx), %rdx
	mulxq	%r10, %rbx, %r8
	mulxq	-24(%rsp), %r10, %rdi   # 8-byte Folded Reload
	addq	%rbx, %rdi
	mulxq	%r11, %rcx, %rbx
	adcq	%r8, %rcx
	adcq	$0, %rbx
	addq	%rax, %r10
	adcq	%rbp, %rdi
	adcq	%rsi, %rcx
	adcq	$0, %rbx
	imulq	%r10, %r9
	movq	%r9, %rdx
	mulxq	%r14, %rdx, %r8
	addq	%r10, %rdx
	movq	%r9, %rdx
	mulxq	%r12, %rbp, %rsi
	movq	%r9, %rdx
	mulxq	%r15, %rax, %rdx
	adcq	%rdi, %rax
	adcq	%rcx, %rbp
	adcq	$0, %rbx
	addq	%r8, %rax
	adcq	%rdx, %rbp
	adcq	%rsi, %rbx
	movq	%rax, %rcx
	subq	%r14, %rcx
	movq	%rbp, %rdx
	sbbq	%r15, %rdx
	movq	%rbx, %rsi
	sbbq	%r12, %rsi
	movq	%rsi, %rdi
	sarq	$63, %rdi
	cmovsq	%rax, %rcx
	movq	-8(%rsp), %rax          # 8-byte Reload
	movq	%rcx, (%rax)
	cmovsq	%rbp, %rdx
	movq	%rdx, 8(%rax)
	cmovsq	%rbx, %rsi
	movq	%rsi, 16(%rax)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end39:
	.size	mcl_fp_montNF3Lbmi2, .Lfunc_end39-mcl_fp_montNF3Lbmi2

	.globl	mcl_fp_montRed3Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montRed3Lbmi2,@function
mcl_fp_montRed3Lbmi2:                   # @mcl_fp_montRed3Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rdx, %rcx
	movq	%rdi, -8(%rsp)          # 8-byte Spill
	movq	-8(%rcx), %r15
	movq	(%rcx), %r9
	movq	(%rsi), %rbx
	movq	%rbx, %rdx
	imulq	%r15, %rdx
	movq	16(%rcx), %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	mulxq	%rax, %r14, %r11
	movq	%rax, %rbp
	movq	8(%rcx), %r10
	mulxq	%r10, %rax, %r13
	mulxq	%r9, %rdx, %rcx
	addq	%rax, %rcx
	adcq	%r14, %r13
	adcq	$0, %r11
	movq	40(%rsi), %r14
	movq	32(%rsi), %r12
	addq	%rbx, %rdx
	adcq	8(%rsi), %rcx
	adcq	16(%rsi), %r13
	adcq	24(%rsi), %r11
	adcq	$0, %r12
	adcq	$0, %r14
	sbbq	%rsi, %rsi
	andl	$1, %esi
	movq	%rcx, %rdx
	imulq	%r15, %rdx
	mulxq	%rbp, %rbp, %rdi
	mulxq	%r10, %r8, %rbx
	mulxq	%r9, %rdx, %rax
	addq	%r8, %rax
	adcq	%rbp, %rbx
	adcq	$0, %rdi
	addq	%rcx, %rdx
	adcq	%r13, %rax
	adcq	%r11, %rbx
	adcq	%r12, %rdi
	adcq	$0, %r14
	adcq	$0, %rsi
	imulq	%rax, %r15
	movq	%r15, %rdx
	movq	-16(%rsp), %r13         # 8-byte Reload
	mulxq	%r13, %r8, %rcx
	movq	%r15, %rdx
	mulxq	%r10, %r11, %r12
	movq	%r15, %rdx
	mulxq	%r9, %r15, %rdx
	addq	%r11, %rdx
	adcq	%r8, %r12
	adcq	$0, %rcx
	addq	%rax, %r15
	adcq	%rbx, %rdx
	adcq	%rdi, %r12
	adcq	%r14, %rcx
	adcq	$0, %rsi
	movq	%rdx, %rax
	subq	%r9, %rax
	movq	%r12, %rdi
	sbbq	%r10, %rdi
	movq	%rcx, %rbp
	sbbq	%r13, %rbp
	sbbq	$0, %rsi
	andl	$1, %esi
	cmovneq	%rcx, %rbp
	testb	%sil, %sil
	cmovneq	%rdx, %rax
	movq	-8(%rsp), %rcx          # 8-byte Reload
	movq	%rax, (%rcx)
	cmovneq	%r12, %rdi
	movq	%rdi, 8(%rcx)
	movq	%rbp, 16(%rcx)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end40:
	.size	mcl_fp_montRed3Lbmi2, .Lfunc_end40-mcl_fp_montRed3Lbmi2

	.globl	mcl_fp_addPre3Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addPre3Lbmi2,@function
mcl_fp_addPre3Lbmi2:                    # @mcl_fp_addPre3Lbmi2
# BB#0:
	movq	16(%rdx), %rax
	movq	(%rdx), %rcx
	movq	8(%rdx), %rdx
	addq	(%rsi), %rcx
	adcq	8(%rsi), %rdx
	adcq	16(%rsi), %rax
	movq	%rcx, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%rax, 16(%rdi)
	sbbq	%rax, %rax
	andl	$1, %eax
	retq
.Lfunc_end41:
	.size	mcl_fp_addPre3Lbmi2, .Lfunc_end41-mcl_fp_addPre3Lbmi2

	.globl	mcl_fp_subPre3Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subPre3Lbmi2,@function
mcl_fp_subPre3Lbmi2:                    # @mcl_fp_subPre3Lbmi2
# BB#0:
	movq	16(%rsi), %r8
	movq	(%rsi), %rcx
	movq	8(%rsi), %rsi
	xorl	%eax, %eax
	subq	(%rdx), %rcx
	sbbq	8(%rdx), %rsi
	sbbq	16(%rdx), %r8
	movq	%rcx, (%rdi)
	movq	%rsi, 8(%rdi)
	movq	%r8, 16(%rdi)
	sbbq	$0, %rax
	andl	$1, %eax
	retq
.Lfunc_end42:
	.size	mcl_fp_subPre3Lbmi2, .Lfunc_end42-mcl_fp_subPre3Lbmi2

	.globl	mcl_fp_shr1_3Lbmi2
	.align	16, 0x90
	.type	mcl_fp_shr1_3Lbmi2,@function
mcl_fp_shr1_3Lbmi2:                     # @mcl_fp_shr1_3Lbmi2
# BB#0:
	movq	16(%rsi), %rax
	movq	(%rsi), %rcx
	movq	8(%rsi), %rdx
	shrdq	$1, %rdx, %rcx
	movq	%rcx, (%rdi)
	shrdq	$1, %rax, %rdx
	movq	%rdx, 8(%rdi)
	shrq	%rax
	movq	%rax, 16(%rdi)
	retq
.Lfunc_end43:
	.size	mcl_fp_shr1_3Lbmi2, .Lfunc_end43-mcl_fp_shr1_3Lbmi2

	.globl	mcl_fp_add3Lbmi2
	.align	16, 0x90
	.type	mcl_fp_add3Lbmi2,@function
mcl_fp_add3Lbmi2:                       # @mcl_fp_add3Lbmi2
# BB#0:
	movq	16(%rdx), %r8
	movq	(%rdx), %rax
	movq	8(%rdx), %rdx
	addq	(%rsi), %rax
	adcq	8(%rsi), %rdx
	adcq	16(%rsi), %r8
	movq	%rax, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%r8, 16(%rdi)
	sbbq	%rsi, %rsi
	andl	$1, %esi
	subq	(%rcx), %rax
	sbbq	8(%rcx), %rdx
	sbbq	16(%rcx), %r8
	sbbq	$0, %rsi
	testb	$1, %sil
	jne	.LBB44_2
# BB#1:                                 # %nocarry
	movq	%rax, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%r8, 16(%rdi)
.LBB44_2:                               # %carry
	retq
.Lfunc_end44:
	.size	mcl_fp_add3Lbmi2, .Lfunc_end44-mcl_fp_add3Lbmi2

	.globl	mcl_fp_addNF3Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addNF3Lbmi2,@function
mcl_fp_addNF3Lbmi2:                     # @mcl_fp_addNF3Lbmi2
# BB#0:
	movq	16(%rdx), %r8
	movq	(%rdx), %r10
	movq	8(%rdx), %r9
	addq	(%rsi), %r10
	adcq	8(%rsi), %r9
	adcq	16(%rsi), %r8
	movq	%r10, %rsi
	subq	(%rcx), %rsi
	movq	%r9, %rdx
	sbbq	8(%rcx), %rdx
	movq	%r8, %rax
	sbbq	16(%rcx), %rax
	movq	%rax, %rcx
	sarq	$63, %rcx
	cmovsq	%r10, %rsi
	movq	%rsi, (%rdi)
	cmovsq	%r9, %rdx
	movq	%rdx, 8(%rdi)
	cmovsq	%r8, %rax
	movq	%rax, 16(%rdi)
	retq
.Lfunc_end45:
	.size	mcl_fp_addNF3Lbmi2, .Lfunc_end45-mcl_fp_addNF3Lbmi2

	.globl	mcl_fp_sub3Lbmi2
	.align	16, 0x90
	.type	mcl_fp_sub3Lbmi2,@function
mcl_fp_sub3Lbmi2:                       # @mcl_fp_sub3Lbmi2
# BB#0:
	movq	16(%rsi), %r8
	movq	(%rsi), %rax
	movq	8(%rsi), %r9
	xorl	%esi, %esi
	subq	(%rdx), %rax
	sbbq	8(%rdx), %r9
	sbbq	16(%rdx), %r8
	movq	%rax, (%rdi)
	movq	%r9, 8(%rdi)
	movq	%r8, 16(%rdi)
	sbbq	$0, %rsi
	testb	$1, %sil
	jne	.LBB46_2
# BB#1:                                 # %nocarry
	retq
.LBB46_2:                               # %carry
	movq	8(%rcx), %rdx
	movq	16(%rcx), %rsi
	addq	(%rcx), %rax
	movq	%rax, (%rdi)
	adcq	%r9, %rdx
	movq	%rdx, 8(%rdi)
	adcq	%r8, %rsi
	movq	%rsi, 16(%rdi)
	retq
.Lfunc_end46:
	.size	mcl_fp_sub3Lbmi2, .Lfunc_end46-mcl_fp_sub3Lbmi2

	.globl	mcl_fp_subNF3Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subNF3Lbmi2,@function
mcl_fp_subNF3Lbmi2:                     # @mcl_fp_subNF3Lbmi2
# BB#0:
	movq	16(%rsi), %r10
	movq	(%rsi), %r8
	movq	8(%rsi), %r9
	subq	(%rdx), %r8
	sbbq	8(%rdx), %r9
	sbbq	16(%rdx), %r10
	movq	%r10, %rdx
	sarq	$63, %rdx
	movq	%rdx, %rsi
	shldq	$1, %r10, %rsi
	andq	(%rcx), %rsi
	movq	16(%rcx), %rax
	andq	%rdx, %rax
	andq	8(%rcx), %rdx
	addq	%r8, %rsi
	movq	%rsi, (%rdi)
	adcq	%r9, %rdx
	movq	%rdx, 8(%rdi)
	adcq	%r10, %rax
	movq	%rax, 16(%rdi)
	retq
.Lfunc_end47:
	.size	mcl_fp_subNF3Lbmi2, .Lfunc_end47-mcl_fp_subNF3Lbmi2

	.globl	mcl_fpDbl_add3Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_add3Lbmi2,@function
mcl_fpDbl_add3Lbmi2:                    # @mcl_fpDbl_add3Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%rbx
	movq	40(%rdx), %r10
	movq	40(%rsi), %r8
	movq	32(%rdx), %r11
	movq	24(%rdx), %r14
	movq	24(%rsi), %r15
	movq	32(%rsi), %r9
	movq	16(%rdx), %rbx
	movq	(%rdx), %rax
	movq	8(%rdx), %rdx
	addq	(%rsi), %rax
	adcq	8(%rsi), %rdx
	adcq	16(%rsi), %rbx
	movq	%rax, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%rbx, 16(%rdi)
	adcq	%r14, %r15
	adcq	%r11, %r9
	adcq	%r10, %r8
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%r15, %rdx
	subq	(%rcx), %rdx
	movq	%r9, %rsi
	sbbq	8(%rcx), %rsi
	movq	%r8, %rbx
	sbbq	16(%rcx), %rbx
	sbbq	$0, %rax
	andl	$1, %eax
	cmovneq	%r15, %rdx
	movq	%rdx, 24(%rdi)
	testb	%al, %al
	cmovneq	%r9, %rsi
	movq	%rsi, 32(%rdi)
	cmovneq	%r8, %rbx
	movq	%rbx, 40(%rdi)
	popq	%rbx
	popq	%r14
	popq	%r15
	retq
.Lfunc_end48:
	.size	mcl_fpDbl_add3Lbmi2, .Lfunc_end48-mcl_fpDbl_add3Lbmi2

	.globl	mcl_fpDbl_sub3Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sub3Lbmi2,@function
mcl_fpDbl_sub3Lbmi2:                    # @mcl_fpDbl_sub3Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%r12
	pushq	%rbx
	movq	40(%rdx), %r10
	movq	40(%rsi), %r8
	movq	32(%rsi), %r9
	movq	24(%rsi), %r11
	movq	16(%rsi), %r14
	movq	(%rsi), %rbx
	movq	8(%rsi), %rax
	xorl	%esi, %esi
	subq	(%rdx), %rbx
	sbbq	8(%rdx), %rax
	movq	24(%rdx), %r15
	movq	32(%rdx), %r12
	sbbq	16(%rdx), %r14
	movq	%rbx, (%rdi)
	movq	%rax, 8(%rdi)
	movq	%r14, 16(%rdi)
	sbbq	%r15, %r11
	sbbq	%r12, %r9
	sbbq	%r10, %r8
	movl	$0, %eax
	sbbq	$0, %rax
	andl	$1, %eax
	movq	(%rcx), %rdx
	cmoveq	%rsi, %rdx
	testb	%al, %al
	movq	16(%rcx), %rax
	cmoveq	%rsi, %rax
	cmovneq	8(%rcx), %rsi
	addq	%r11, %rdx
	movq	%rdx, 24(%rdi)
	adcq	%r9, %rsi
	movq	%rsi, 32(%rdi)
	adcq	%r8, %rax
	movq	%rax, 40(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r14
	popq	%r15
	retq
.Lfunc_end49:
	.size	mcl_fpDbl_sub3Lbmi2, .Lfunc_end49-mcl_fpDbl_sub3Lbmi2

	.globl	mcl_fp_mulUnitPre4Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mulUnitPre4Lbmi2,@function
mcl_fp_mulUnitPre4Lbmi2:                # @mcl_fp_mulUnitPre4Lbmi2
# BB#0:
	mulxq	24(%rsi), %r8, %r11
	mulxq	16(%rsi), %r9, %rax
	mulxq	8(%rsi), %r10, %rcx
	mulxq	(%rsi), %rdx, %rsi
	movq	%rdx, (%rdi)
	addq	%r10, %rsi
	movq	%rsi, 8(%rdi)
	adcq	%r9, %rcx
	movq	%rcx, 16(%rdi)
	adcq	%r8, %rax
	movq	%rax, 24(%rdi)
	adcq	$0, %r11
	movq	%r11, 32(%rdi)
	retq
.Lfunc_end50:
	.size	mcl_fp_mulUnitPre4Lbmi2, .Lfunc_end50-mcl_fp_mulUnitPre4Lbmi2

	.globl	mcl_fpDbl_mulPre4Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_mulPre4Lbmi2,@function
mcl_fpDbl_mulPre4Lbmi2:                 # @mcl_fpDbl_mulPre4Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	(%rsi), %r14
	movq	8(%rsi), %r10
	movq	(%rdx), %rcx
	movq	%rdx, %rbp
	movq	%r14, %rdx
	mulxq	%rcx, %rdx, %r15
	movq	24(%rsi), %r11
	movq	16(%rsi), %r9
	movq	%rdx, (%rdi)
	movq	%r10, %rdx
	mulxq	%rcx, %rbx, %r12
	addq	%r15, %rbx
	movq	%r9, %rdx
	mulxq	%rcx, %r13, %r15
	adcq	%r12, %r13
	movq	%r11, %rdx
	mulxq	%rcx, %rcx, %r12
	adcq	%r15, %rcx
	adcq	$0, %r12
	movq	8(%rbp), %rax
	movq	%r14, %rdx
	mulxq	%rax, %r8, %rdx
	movq	%rdx, -8(%rsp)          # 8-byte Spill
	addq	%rbx, %r8
	movq	%r10, %rdx
	mulxq	%rax, %r15, %rdx
	movq	%rdx, -16(%rsp)         # 8-byte Spill
	adcq	%r13, %r15
	movq	%r9, %rdx
	mulxq	%rax, %rbx, %r13
	adcq	%rcx, %rbx
	movq	%r11, %rdx
	mulxq	%rax, %rcx, %rax
	adcq	%r12, %rcx
	sbbq	%r12, %r12
	andl	$1, %r12d
	addq	-8(%rsp), %r15          # 8-byte Folded Reload
	adcq	-16(%rsp), %rbx         # 8-byte Folded Reload
	adcq	%r13, %rcx
	movq	%r8, 8(%rdi)
	adcq	%rax, %r12
	movq	%rbp, %r13
	movq	16(%r13), %rax
	movq	%r14, %rdx
	mulxq	%rax, %rdx, %r8
	addq	%r15, %rdx
	movq	%rdx, 16(%rdi)
	movq	%r10, %rdx
	mulxq	%rax, %rbp, %r10
	adcq	%rbx, %rbp
	movq	%r11, %rdx
	mulxq	%rax, %r14, %r11
	movq	%r9, %rdx
	mulxq	%rax, %r15, %rdx
	adcq	%rcx, %r15
	adcq	%r12, %r14
	sbbq	%rcx, %rcx
	andl	$1, %ecx
	addq	%r8, %rbp
	adcq	%r10, %r15
	adcq	%rdx, %r14
	adcq	%r11, %rcx
	movq	24(%r13), %rdx
	mulxq	24(%rsi), %rbx, %r8
	mulxq	(%rsi), %rax, %r9
	addq	%rbp, %rax
	mulxq	16(%rsi), %rbp, %r10
	mulxq	8(%rsi), %rsi, %rdx
	movq	%rax, 24(%rdi)
	adcq	%r15, %rsi
	adcq	%r14, %rbp
	adcq	%rcx, %rbx
	sbbq	%rax, %rax
	andl	$1, %eax
	addq	%r9, %rsi
	movq	%rsi, 32(%rdi)
	adcq	%rdx, %rbp
	movq	%rbp, 40(%rdi)
	adcq	%r10, %rbx
	movq	%rbx, 48(%rdi)
	adcq	%r8, %rax
	movq	%rax, 56(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end51:
	.size	mcl_fpDbl_mulPre4Lbmi2, .Lfunc_end51-mcl_fpDbl_mulPre4Lbmi2

	.globl	mcl_fpDbl_sqrPre4Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sqrPre4Lbmi2,@function
mcl_fpDbl_sqrPre4Lbmi2:                 # @mcl_fpDbl_sqrPre4Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	24(%rsi), %r8
	movq	16(%rsi), %r9
	movq	(%rsi), %rcx
	movq	8(%rsi), %rax
	movq	%rcx, %rdx
	mulxq	%rcx, %rdx, %r11
	movq	%rdx, (%rdi)
	movq	%r9, %rdx
	mulxq	%rcx, %rbp, %r10
	movq	%rbp, -16(%rsp)         # 8-byte Spill
	movq	%r10, -8(%rsp)          # 8-byte Spill
	movq	%rax, %rdx
	mulxq	%rcx, %r12, %r15
	addq	%r12, %r11
	movq	%r15, %rbx
	adcq	%rbp, %rbx
	movq	%r8, %rdx
	mulxq	%rcx, %rcx, %r13
	adcq	%r10, %rcx
	adcq	$0, %r13
	addq	%r12, %r11
	movq	%rax, %rdx
	mulxq	%rax, %rbp, %r12
	adcq	%rbx, %rbp
	movq	%r8, %rdx
	mulxq	%rax, %r10, %rbx
	movq	%r9, %rdx
	mulxq	%rax, %r14, %rdx
	adcq	%r14, %rcx
	adcq	%r13, %r10
	sbbq	%rax, %rax
	andl	$1, %eax
	addq	%r15, %rbp
	adcq	%r12, %rcx
	adcq	%rdx, %r10
	movq	%rdx, %r12
	adcq	%rbx, %rax
	movq	%r11, 8(%rdi)
	addq	-16(%rsp), %rbp         # 8-byte Folded Reload
	movq	%rbp, 16(%rdi)
	movq	%r8, %rdx
	mulxq	%r9, %r11, %r8
	movq	%r9, %rdx
	mulxq	%r9, %r15, %rdx
	adcq	%r14, %rcx
	adcq	%r10, %r15
	adcq	%rax, %r11
	sbbq	%rax, %rax
	andl	$1, %eax
	addq	-8(%rsp), %rcx          # 8-byte Folded Reload
	adcq	%r12, %r15
	adcq	%rdx, %r11
	adcq	%r8, %rax
	movq	24(%rsi), %rdx
	mulxq	16(%rsi), %rbx, %r8
	mulxq	8(%rsi), %rbp, %r9
	mulxq	(%rsi), %rsi, %r10
	addq	%rcx, %rsi
	movq	%rsi, 24(%rdi)
	adcq	%r15, %rbp
	adcq	%r11, %rbx
	mulxq	%rdx, %rdx, %rcx
	adcq	%rax, %rdx
	sbbq	%rax, %rax
	andl	$1, %eax
	addq	%r10, %rbp
	movq	%rbp, 32(%rdi)
	adcq	%r9, %rbx
	movq	%rbx, 40(%rdi)
	adcq	%r8, %rdx
	movq	%rdx, 48(%rdi)
	adcq	%rcx, %rax
	movq	%rax, 56(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end52:
	.size	mcl_fpDbl_sqrPre4Lbmi2, .Lfunc_end52-mcl_fpDbl_sqrPre4Lbmi2

	.globl	mcl_fp_mont4Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mont4Lbmi2,@function
mcl_fp_mont4Lbmi2:                      # @mcl_fp_mont4Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	movq	%rdi, -88(%rsp)         # 8-byte Spill
	movq	24(%rsi), %rdi
	movq	%rdi, -40(%rsp)         # 8-byte Spill
	movq	(%rdx), %rax
	movq	%rdi, %rdx
	mulxq	%rax, %r10, %r15
	movq	16(%rsi), %rdx
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	mulxq	%rax, %rbx, %r11
	movq	(%rsi), %rdi
	movq	%rdi, -56(%rsp)         # 8-byte Spill
	movq	8(%rsi), %rdx
	movq	%rdx, -64(%rsp)         # 8-byte Spill
	mulxq	%rax, %rbp, %r14
	movq	%rdi, %rdx
	mulxq	%rax, %r13, %r12
	addq	%rbp, %r12
	adcq	%rbx, %r14
	adcq	%r10, %r11
	adcq	$0, %r15
	movq	-8(%rcx), %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	movq	(%rcx), %rdi
	movq	%rdi, -24(%rsp)         # 8-byte Spill
	movq	%r13, %rdx
	imulq	%rax, %rdx
	movq	24(%rcx), %rsi
	movq	%rsi, -72(%rsp)         # 8-byte Spill
	movq	16(%rcx), %rbp
	movq	%rbp, -8(%rsp)          # 8-byte Spill
	movq	8(%rcx), %rax
	movq	%rax, -80(%rsp)         # 8-byte Spill
	mulxq	%rsi, %r10, %r8
	mulxq	%rbp, %r9, %rbx
	mulxq	%rax, %rsi, %rcx
	mulxq	%rdi, %rdx, %rbp
	addq	%rsi, %rbp
	adcq	%r9, %rcx
	adcq	%r10, %rbx
	adcq	$0, %r8
	addq	%r13, %rdx
	adcq	%r12, %rbp
	adcq	%r14, %rcx
	adcq	%r11, %rbx
	adcq	%r15, %r8
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	-32(%rsp), %rdx         # 8-byte Reload
	movq	8(%rdx), %rdx
	mulxq	-40(%rsp), %r12, %r14   # 8-byte Folded Reload
	mulxq	-48(%rsp), %r15, %r11   # 8-byte Folded Reload
	mulxq	-64(%rsp), %r9, %rdi    # 8-byte Folded Reload
	mulxq	-56(%rsp), %r10, %rsi   # 8-byte Folded Reload
	addq	%r9, %rsi
	adcq	%r15, %rdi
	adcq	%r12, %r11
	adcq	$0, %r14
	addq	%rbp, %r10
	adcq	%rcx, %rsi
	adcq	%rbx, %rdi
	adcq	%r8, %r11
	adcq	%rax, %r14
	sbbq	%rbx, %rbx
	andl	$1, %ebx
	movq	%r10, %rdx
	imulq	-16(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	-72(%rsp), %r15, %r9    # 8-byte Folded Reload
	mulxq	-8(%rsp), %r12, %r8     # 8-byte Folded Reload
	movq	-80(%rsp), %r13         # 8-byte Reload
	mulxq	%r13, %rbp, %rcx
	mulxq	-24(%rsp), %rdx, %rax   # 8-byte Folded Reload
	addq	%rbp, %rax
	adcq	%r12, %rcx
	adcq	%r15, %r8
	adcq	$0, %r9
	addq	%r10, %rdx
	adcq	%rsi, %rax
	adcq	%rdi, %rcx
	adcq	%r11, %r8
	adcq	%r14, %r9
	adcq	$0, %rbx
	movq	-32(%rsp), %rdx         # 8-byte Reload
	movq	16(%rdx), %rdx
	mulxq	-40(%rsp), %r15, %r11   # 8-byte Folded Reload
	mulxq	-48(%rsp), %r12, %r14   # 8-byte Folded Reload
	mulxq	-64(%rsp), %rsi, %rbp   # 8-byte Folded Reload
	mulxq	-56(%rsp), %r10, %rdi   # 8-byte Folded Reload
	addq	%rsi, %rdi
	adcq	%r12, %rbp
	adcq	%r15, %r14
	adcq	$0, %r11
	addq	%rax, %r10
	adcq	%rcx, %rdi
	adcq	%r8, %rbp
	adcq	%r9, %r14
	adcq	%rbx, %r11
	sbbq	%rbx, %rbx
	movq	%r10, %rdx
	imulq	-16(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	%r13, %rcx, %rsi
	mulxq	-24(%rsp), %r8, %rax    # 8-byte Folded Reload
	addq	%rcx, %rax
	mulxq	-8(%rsp), %rcx, %r15    # 8-byte Folded Reload
	adcq	%rsi, %rcx
	movq	-72(%rsp), %r13         # 8-byte Reload
	mulxq	%r13, %r9, %rsi
	adcq	%r15, %r9
	adcq	$0, %rsi
	andl	$1, %ebx
	addq	%r10, %r8
	adcq	%rdi, %rax
	adcq	%rbp, %rcx
	adcq	%r14, %r9
	adcq	%r11, %rsi
	adcq	$0, %rbx
	movq	-32(%rsp), %rdx         # 8-byte Reload
	movq	24(%rdx), %rdx
	mulxq	-40(%rsp), %r11, %r8    # 8-byte Folded Reload
	mulxq	-48(%rsp), %r15, %rdi   # 8-byte Folded Reload
	mulxq	-64(%rsp), %r12, %r14   # 8-byte Folded Reload
	mulxq	-56(%rsp), %r10, %rbp   # 8-byte Folded Reload
	addq	%r12, %rbp
	adcq	%r15, %r14
	adcq	%r11, %rdi
	adcq	$0, %r8
	addq	%rax, %r10
	adcq	%rcx, %rbp
	adcq	%r9, %r14
	adcq	%rsi, %rdi
	adcq	%rbx, %r8
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	-16(%rsp), %rdx         # 8-byte Reload
	imulq	%r10, %rdx
	mulxq	%r13, %rcx, %rsi
	movq	%rcx, -16(%rsp)         # 8-byte Spill
	mulxq	-8(%rsp), %r11, %rbx    # 8-byte Folded Reload
	mulxq	-80(%rsp), %r15, %rcx   # 8-byte Folded Reload
	movq	-24(%rsp), %r9          # 8-byte Reload
	mulxq	%r9, %r12, %r13
	addq	%r15, %r13
	adcq	%r11, %rcx
	adcq	-16(%rsp), %rbx         # 8-byte Folded Reload
	adcq	$0, %rsi
	addq	%r10, %r12
	adcq	%rbp, %r13
	adcq	%r14, %rcx
	adcq	%rdi, %rbx
	adcq	%r8, %rsi
	adcq	$0, %rax
	movq	%r13, %rdi
	subq	%r9, %rdi
	movq	%rcx, %rbp
	sbbq	-80(%rsp), %rbp         # 8-byte Folded Reload
	movq	%rbx, %r8
	sbbq	-8(%rsp), %r8           # 8-byte Folded Reload
	movq	%rsi, %rdx
	sbbq	-72(%rsp), %rdx         # 8-byte Folded Reload
	sbbq	$0, %rax
	andl	$1, %eax
	cmovneq	%rsi, %rdx
	testb	%al, %al
	cmovneq	%r13, %rdi
	movq	-88(%rsp), %rax         # 8-byte Reload
	movq	%rdi, (%rax)
	cmovneq	%rcx, %rbp
	movq	%rbp, 8(%rax)
	cmovneq	%rbx, %r8
	movq	%r8, 16(%rax)
	movq	%rdx, 24(%rax)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end53:
	.size	mcl_fp_mont4Lbmi2, .Lfunc_end53-mcl_fp_mont4Lbmi2

	.globl	mcl_fp_montNF4Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montNF4Lbmi2,@function
mcl_fp_montNF4Lbmi2:                    # @mcl_fp_montNF4Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	movq	%rdi, -80(%rsp)         # 8-byte Spill
	movq	(%rsi), %rdi
	movq	%rdi, -64(%rsp)         # 8-byte Spill
	movq	8(%rsi), %rbp
	movq	%rbp, -72(%rsp)         # 8-byte Spill
	movq	(%rdx), %rax
	movq	%rdx, %r15
	movq	%rbp, %rdx
	mulxq	%rax, %rbp, %r9
	movq	%rdi, %rdx
	mulxq	%rax, %r12, %rbx
	movq	16(%rsi), %rdx
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	addq	%rbp, %rbx
	mulxq	%rax, %r14, %rbp
	adcq	%r9, %r14
	movq	24(%rsi), %rdx
	movq	%rdx, -16(%rsp)         # 8-byte Spill
	mulxq	%rax, %r8, %rdi
	adcq	%rbp, %r8
	adcq	$0, %rdi
	movq	-8(%rcx), %r13
	movq	(%rcx), %rax
	movq	%rax, -56(%rsp)         # 8-byte Spill
	movq	%r12, %rdx
	imulq	%r13, %rdx
	mulxq	%rax, %rax, %r11
	addq	%r12, %rax
	movq	8(%rcx), %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	mulxq	%rax, %rbp, %r10
	adcq	%rbx, %rbp
	movq	16(%rcx), %rax
	movq	%rax, -40(%rsp)         # 8-byte Spill
	mulxq	%rax, %rsi, %rbx
	adcq	%r14, %rsi
	movq	24(%rcx), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	mulxq	%rax, %rcx, %rdx
	adcq	%r8, %rcx
	adcq	$0, %rdi
	addq	%r11, %rbp
	adcq	%r10, %rsi
	adcq	%rbx, %rcx
	adcq	%rdx, %rdi
	movq	8(%r15), %rdx
	movq	-72(%rsp), %r12         # 8-byte Reload
	mulxq	%r12, %rbx, %r9
	movq	-64(%rsp), %r15         # 8-byte Reload
	mulxq	%r15, %r10, %r11
	addq	%rbx, %r11
	mulxq	-48(%rsp), %rax, %r8    # 8-byte Folded Reload
	adcq	%r9, %rax
	mulxq	-16(%rsp), %r9, %rbx    # 8-byte Folded Reload
	adcq	%r8, %r9
	adcq	$0, %rbx
	addq	%rbp, %r10
	adcq	%rsi, %r11
	adcq	%rcx, %rax
	adcq	%rdi, %r9
	adcq	$0, %rbx
	movq	%r10, %rdx
	imulq	%r13, %rdx
	movq	-56(%rsp), %r14         # 8-byte Reload
	mulxq	%r14, %rcx, %r8
	addq	%r10, %rcx
	mulxq	-24(%rsp), %r10, %rdi   # 8-byte Folded Reload
	adcq	%r11, %r10
	mulxq	-40(%rsp), %rcx, %rsi   # 8-byte Folded Reload
	adcq	%rax, %rcx
	mulxq	-8(%rsp), %rax, %rdx    # 8-byte Folded Reload
	adcq	%r9, %rax
	adcq	$0, %rbx
	addq	%r8, %r10
	adcq	%rdi, %rcx
	adcq	%rsi, %rax
	adcq	%rdx, %rbx
	movq	-32(%rsp), %rdx         # 8-byte Reload
	movq	16(%rdx), %rdx
	mulxq	%r12, %rsi, %r8
	mulxq	%r15, %r11, %rbp
	addq	%rsi, %rbp
	movq	-48(%rsp), %r12         # 8-byte Reload
	mulxq	%r12, %rdi, %r9
	adcq	%r8, %rdi
	mulxq	-16(%rsp), %r8, %rsi    # 8-byte Folded Reload
	adcq	%r9, %r8
	adcq	$0, %rsi
	addq	%r10, %r11
	adcq	%rcx, %rbp
	adcq	%rax, %rdi
	adcq	%rbx, %r8
	adcq	$0, %rsi
	movq	%r11, %rdx
	imulq	%r13, %rdx
	mulxq	%r14, %rax, %r10
	addq	%r11, %rax
	movq	-24(%rsp), %r14         # 8-byte Reload
	mulxq	%r14, %r9, %rbx
	adcq	%rbp, %r9
	movq	-40(%rsp), %r15         # 8-byte Reload
	mulxq	%r15, %rax, %rbp
	adcq	%rdi, %rax
	mulxq	-8(%rsp), %rcx, %rdx    # 8-byte Folded Reload
	adcq	%r8, %rcx
	adcq	$0, %rsi
	addq	%r10, %r9
	adcq	%rbx, %rax
	adcq	%rbp, %rcx
	adcq	%rdx, %rsi
	movq	-32(%rsp), %rdx         # 8-byte Reload
	movq	24(%rdx), %rdx
	mulxq	-72(%rsp), %rbx, %r8    # 8-byte Folded Reload
	mulxq	-64(%rsp), %r11, %rbp   # 8-byte Folded Reload
	addq	%rbx, %rbp
	mulxq	%r12, %rdi, %r10
	adcq	%r8, %rdi
	mulxq	-16(%rsp), %r8, %rbx    # 8-byte Folded Reload
	adcq	%r10, %r8
	adcq	$0, %rbx
	addq	%r9, %r11
	adcq	%rax, %rbp
	adcq	%rcx, %rdi
	adcq	%rsi, %r8
	adcq	$0, %rbx
	imulq	%r11, %r13
	movq	%r13, %rdx
	movq	-56(%rsp), %r12         # 8-byte Reload
	mulxq	%r12, %rcx, %r9
	addq	%r11, %rcx
	movq	%r13, %rdx
	mulxq	%r14, %r11, %r10
	adcq	%rbp, %r11
	movq	%r13, %rdx
	movq	%r15, %rsi
	mulxq	%rsi, %rax, %rcx
	adcq	%rdi, %rax
	movq	%r13, %rdx
	movq	-8(%rsp), %rbp          # 8-byte Reload
	mulxq	%rbp, %r15, %rdx
	adcq	%r8, %r15
	adcq	$0, %rbx
	addq	%r9, %r11
	adcq	%r10, %rax
	adcq	%rcx, %r15
	adcq	%rdx, %rbx
	movq	%r11, %rcx
	subq	%r12, %rcx
	movq	%rax, %rdx
	sbbq	%r14, %rdx
	movq	%r15, %rdi
	sbbq	%rsi, %rdi
	movq	%rbx, %rsi
	sbbq	%rbp, %rsi
	cmovsq	%r11, %rcx
	movq	-80(%rsp), %rbp         # 8-byte Reload
	movq	%rcx, (%rbp)
	cmovsq	%rax, %rdx
	movq	%rdx, 8(%rbp)
	cmovsq	%r15, %rdi
	movq	%rdi, 16(%rbp)
	cmovsq	%rbx, %rsi
	movq	%rsi, 24(%rbp)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end54:
	.size	mcl_fp_montNF4Lbmi2, .Lfunc_end54-mcl_fp_montNF4Lbmi2

	.globl	mcl_fp_montRed4Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montRed4Lbmi2,@function
mcl_fp_montRed4Lbmi2:                   # @mcl_fp_montRed4Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rdx, %rcx
	movq	%rdi, -56(%rsp)         # 8-byte Spill
	movq	-8(%rcx), %r13
	movq	(%rcx), %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	movq	(%rsi), %r10
	movq	%r10, %rdx
	imulq	%r13, %rdx
	movq	24(%rcx), %rdi
	movq	%rdi, -48(%rsp)         # 8-byte Spill
	mulxq	%rdi, %r9, %r15
	movq	%rdi, %r14
	movq	16(%rcx), %rdi
	movq	%rdi, -8(%rsp)          # 8-byte Spill
	mulxq	%rdi, %rdi, %rbx
	movq	8(%rcx), %rcx
	movq	%rcx, -16(%rsp)         # 8-byte Spill
	mulxq	%rcx, %rcx, %r8
	mulxq	%rax, %rdx, %rbp
	addq	%rcx, %rbp
	adcq	%rdi, %r8
	adcq	%r9, %rbx
	adcq	$0, %r15
	movq	56(%rsi), %r11
	movq	48(%rsi), %rcx
	addq	%r10, %rdx
	movq	40(%rsi), %r12
	adcq	8(%rsi), %rbp
	adcq	16(%rsi), %r8
	adcq	24(%rsi), %rbx
	adcq	32(%rsi), %r15
	adcq	$0, %r12
	adcq	$0, %rcx
	movq	%rcx, -32(%rsp)         # 8-byte Spill
	adcq	$0, %r11
	sbbq	%rsi, %rsi
	andl	$1, %esi
	movq	%rbp, %rdx
	imulq	%r13, %rdx
	mulxq	%r14, %rax, %r9
	movq	%rax, -40(%rsp)         # 8-byte Spill
	mulxq	-8(%rsp), %r14, %rdi    # 8-byte Folded Reload
	mulxq	-16(%rsp), %r10, %rcx   # 8-byte Folded Reload
	mulxq	-24(%rsp), %rdx, %rax   # 8-byte Folded Reload
	addq	%r10, %rax
	adcq	%r14, %rcx
	adcq	-40(%rsp), %rdi         # 8-byte Folded Reload
	adcq	$0, %r9
	addq	%rbp, %rdx
	adcq	%r8, %rax
	adcq	%rbx, %rcx
	adcq	%r15, %rdi
	adcq	%r12, %r9
	adcq	$0, -32(%rsp)           # 8-byte Folded Spill
	adcq	$0, %r11
	movq	%r11, -40(%rsp)         # 8-byte Spill
	adcq	$0, %rsi
	movq	%rax, %rdx
	imulq	%r13, %rdx
	movq	-48(%rsp), %r15         # 8-byte Reload
	mulxq	%r15, %rbp, %r8
	movq	%rbp, -64(%rsp)         # 8-byte Spill
	movq	-8(%rsp), %r11          # 8-byte Reload
	mulxq	%r11, %rbx, %r10
	movq	%rbx, -72(%rsp)         # 8-byte Spill
	mulxq	-16(%rsp), %r12, %rbp   # 8-byte Folded Reload
	movq	-24(%rsp), %r14         # 8-byte Reload
	mulxq	%r14, %rdx, %rbx
	addq	%r12, %rbx
	adcq	-72(%rsp), %rbp         # 8-byte Folded Reload
	adcq	-64(%rsp), %r10         # 8-byte Folded Reload
	adcq	$0, %r8
	addq	%rax, %rdx
	adcq	%rcx, %rbx
	adcq	%rdi, %rbp
	adcq	%r9, %r10
	adcq	-32(%rsp), %r8          # 8-byte Folded Reload
	adcq	$0, -40(%rsp)           # 8-byte Folded Spill
	adcq	$0, %rsi
	imulq	%rbx, %r13
	movq	%r13, %rdx
	mulxq	%r15, %rax, %rdi
	movq	%rax, -32(%rsp)         # 8-byte Spill
	movq	%r13, %rdx
	mulxq	%r11, %r9, %rax
	movq	%r13, %rdx
	movq	-16(%rsp), %r11         # 8-byte Reload
	mulxq	%r11, %r12, %rcx
	movq	%r13, %rdx
	mulxq	%r14, %r15, %r13
	addq	%r12, %r13
	adcq	%r9, %rcx
	adcq	-32(%rsp), %rax         # 8-byte Folded Reload
	adcq	$0, %rdi
	addq	%rbx, %r15
	adcq	%rbp, %r13
	adcq	%r10, %rcx
	adcq	%r8, %rax
	adcq	-40(%rsp), %rdi         # 8-byte Folded Reload
	adcq	$0, %rsi
	movq	%r13, %rdx
	subq	%r14, %rdx
	movq	%rcx, %rbp
	sbbq	%r11, %rbp
	movq	%rax, %r8
	sbbq	-8(%rsp), %r8           # 8-byte Folded Reload
	movq	%rdi, %rbx
	sbbq	-48(%rsp), %rbx         # 8-byte Folded Reload
	sbbq	$0, %rsi
	andl	$1, %esi
	cmovneq	%rdi, %rbx
	testb	%sil, %sil
	cmovneq	%r13, %rdx
	movq	-56(%rsp), %rsi         # 8-byte Reload
	movq	%rdx, (%rsi)
	cmovneq	%rcx, %rbp
	movq	%rbp, 8(%rsi)
	cmovneq	%rax, %r8
	movq	%r8, 16(%rsi)
	movq	%rbx, 24(%rsi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end55:
	.size	mcl_fp_montRed4Lbmi2, .Lfunc_end55-mcl_fp_montRed4Lbmi2

	.globl	mcl_fp_addPre4Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addPre4Lbmi2,@function
mcl_fp_addPre4Lbmi2:                    # @mcl_fp_addPre4Lbmi2
# BB#0:
	movq	24(%rdx), %r8
	movq	24(%rsi), %r9
	movq	16(%rdx), %rax
	movq	(%rdx), %rcx
	movq	8(%rdx), %rdx
	addq	(%rsi), %rcx
	adcq	8(%rsi), %rdx
	adcq	16(%rsi), %rax
	movq	%rcx, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%rax, 16(%rdi)
	adcq	%r8, %r9
	movq	%r9, 24(%rdi)
	sbbq	%rax, %rax
	andl	$1, %eax
	retq
.Lfunc_end56:
	.size	mcl_fp_addPre4Lbmi2, .Lfunc_end56-mcl_fp_addPre4Lbmi2

	.globl	mcl_fp_subPre4Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subPre4Lbmi2,@function
mcl_fp_subPre4Lbmi2:                    # @mcl_fp_subPre4Lbmi2
# BB#0:
	movq	24(%rdx), %r8
	movq	24(%rsi), %r9
	movq	16(%rsi), %r10
	movq	(%rsi), %rcx
	movq	8(%rsi), %rsi
	xorl	%eax, %eax
	subq	(%rdx), %rcx
	sbbq	8(%rdx), %rsi
	sbbq	16(%rdx), %r10
	movq	%rcx, (%rdi)
	movq	%rsi, 8(%rdi)
	movq	%r10, 16(%rdi)
	sbbq	%r8, %r9
	movq	%r9, 24(%rdi)
	sbbq	$0, %rax
	andl	$1, %eax
	retq
.Lfunc_end57:
	.size	mcl_fp_subPre4Lbmi2, .Lfunc_end57-mcl_fp_subPre4Lbmi2

	.globl	mcl_fp_shr1_4Lbmi2
	.align	16, 0x90
	.type	mcl_fp_shr1_4Lbmi2,@function
mcl_fp_shr1_4Lbmi2:                     # @mcl_fp_shr1_4Lbmi2
# BB#0:
	movq	24(%rsi), %rax
	movq	16(%rsi), %rcx
	movq	(%rsi), %rdx
	movq	8(%rsi), %rsi
	shrdq	$1, %rsi, %rdx
	movq	%rdx, (%rdi)
	shrdq	$1, %rcx, %rsi
	movq	%rsi, 8(%rdi)
	shrdq	$1, %rax, %rcx
	movq	%rcx, 16(%rdi)
	shrq	%rax
	movq	%rax, 24(%rdi)
	retq
.Lfunc_end58:
	.size	mcl_fp_shr1_4Lbmi2, .Lfunc_end58-mcl_fp_shr1_4Lbmi2

	.globl	mcl_fp_add4Lbmi2
	.align	16, 0x90
	.type	mcl_fp_add4Lbmi2,@function
mcl_fp_add4Lbmi2:                       # @mcl_fp_add4Lbmi2
# BB#0:
	movq	24(%rdx), %r10
	movq	24(%rsi), %r8
	movq	16(%rdx), %r9
	movq	(%rdx), %rax
	movq	8(%rdx), %rdx
	addq	(%rsi), %rax
	adcq	8(%rsi), %rdx
	adcq	16(%rsi), %r9
	movq	%rax, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%r9, 16(%rdi)
	adcq	%r10, %r8
	movq	%r8, 24(%rdi)
	sbbq	%rsi, %rsi
	andl	$1, %esi
	subq	(%rcx), %rax
	sbbq	8(%rcx), %rdx
	sbbq	16(%rcx), %r9
	sbbq	24(%rcx), %r8
	sbbq	$0, %rsi
	testb	$1, %sil
	jne	.LBB59_2
# BB#1:                                 # %nocarry
	movq	%rax, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%r9, 16(%rdi)
	movq	%r8, 24(%rdi)
.LBB59_2:                               # %carry
	retq
.Lfunc_end59:
	.size	mcl_fp_add4Lbmi2, .Lfunc_end59-mcl_fp_add4Lbmi2

	.globl	mcl_fp_addNF4Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addNF4Lbmi2,@function
mcl_fp_addNF4Lbmi2:                     # @mcl_fp_addNF4Lbmi2
# BB#0:
	pushq	%rbx
	movq	24(%rdx), %r8
	movq	16(%rdx), %r9
	movq	(%rdx), %r11
	movq	8(%rdx), %r10
	addq	(%rsi), %r11
	adcq	8(%rsi), %r10
	adcq	16(%rsi), %r9
	adcq	24(%rsi), %r8
	movq	%r11, %rsi
	subq	(%rcx), %rsi
	movq	%r10, %rdx
	sbbq	8(%rcx), %rdx
	movq	%r9, %rax
	sbbq	16(%rcx), %rax
	movq	%r8, %rbx
	sbbq	24(%rcx), %rbx
	testq	%rbx, %rbx
	cmovsq	%r11, %rsi
	movq	%rsi, (%rdi)
	cmovsq	%r10, %rdx
	movq	%rdx, 8(%rdi)
	cmovsq	%r9, %rax
	movq	%rax, 16(%rdi)
	cmovsq	%r8, %rbx
	movq	%rbx, 24(%rdi)
	popq	%rbx
	retq
.Lfunc_end60:
	.size	mcl_fp_addNF4Lbmi2, .Lfunc_end60-mcl_fp_addNF4Lbmi2

	.globl	mcl_fp_sub4Lbmi2
	.align	16, 0x90
	.type	mcl_fp_sub4Lbmi2,@function
mcl_fp_sub4Lbmi2:                       # @mcl_fp_sub4Lbmi2
# BB#0:
	movq	24(%rdx), %r10
	movq	24(%rsi), %r8
	movq	16(%rsi), %r9
	movq	(%rsi), %rax
	movq	8(%rsi), %r11
	xorl	%esi, %esi
	subq	(%rdx), %rax
	sbbq	8(%rdx), %r11
	sbbq	16(%rdx), %r9
	movq	%rax, (%rdi)
	movq	%r11, 8(%rdi)
	movq	%r9, 16(%rdi)
	sbbq	%r10, %r8
	movq	%r8, 24(%rdi)
	sbbq	$0, %rsi
	testb	$1, %sil
	jne	.LBB61_2
# BB#1:                                 # %nocarry
	retq
.LBB61_2:                               # %carry
	movq	24(%rcx), %r10
	movq	8(%rcx), %rsi
	movq	16(%rcx), %rdx
	addq	(%rcx), %rax
	movq	%rax, (%rdi)
	adcq	%r11, %rsi
	movq	%rsi, 8(%rdi)
	adcq	%r9, %rdx
	movq	%rdx, 16(%rdi)
	adcq	%r8, %r10
	movq	%r10, 24(%rdi)
	retq
.Lfunc_end61:
	.size	mcl_fp_sub4Lbmi2, .Lfunc_end61-mcl_fp_sub4Lbmi2

	.globl	mcl_fp_subNF4Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subNF4Lbmi2,@function
mcl_fp_subNF4Lbmi2:                     # @mcl_fp_subNF4Lbmi2
# BB#0:
	pushq	%rbx
	movq	24(%rsi), %r11
	movq	16(%rsi), %r8
	movq	(%rsi), %r9
	movq	8(%rsi), %r10
	subq	(%rdx), %r9
	sbbq	8(%rdx), %r10
	sbbq	16(%rdx), %r8
	sbbq	24(%rdx), %r11
	movq	%r11, %rdx
	sarq	$63, %rdx
	movq	24(%rcx), %rsi
	andq	%rdx, %rsi
	movq	16(%rcx), %rax
	andq	%rdx, %rax
	movq	8(%rcx), %rbx
	andq	%rdx, %rbx
	andq	(%rcx), %rdx
	addq	%r9, %rdx
	movq	%rdx, (%rdi)
	adcq	%r10, %rbx
	movq	%rbx, 8(%rdi)
	adcq	%r8, %rax
	movq	%rax, 16(%rdi)
	adcq	%r11, %rsi
	movq	%rsi, 24(%rdi)
	popq	%rbx
	retq
.Lfunc_end62:
	.size	mcl_fp_subNF4Lbmi2, .Lfunc_end62-mcl_fp_subNF4Lbmi2

	.globl	mcl_fpDbl_add4Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_add4Lbmi2,@function
mcl_fpDbl_add4Lbmi2:                    # @mcl_fpDbl_add4Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	56(%rdx), %r9
	movq	56(%rsi), %r8
	movq	48(%rdx), %r10
	movq	48(%rsi), %r12
	movq	40(%rdx), %r11
	movq	32(%rdx), %r14
	movq	24(%rdx), %r15
	movq	16(%rdx), %rbx
	movq	(%rdx), %rax
	movq	8(%rdx), %rdx
	addq	(%rsi), %rax
	adcq	8(%rsi), %rdx
	adcq	16(%rsi), %rbx
	movq	40(%rsi), %r13
	movq	24(%rsi), %rbp
	movq	32(%rsi), %rsi
	movq	%rax, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%rbx, 16(%rdi)
	adcq	%r15, %rbp
	movq	%rbp, 24(%rdi)
	adcq	%r14, %rsi
	adcq	%r11, %r13
	adcq	%r10, %r12
	adcq	%r9, %r8
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%rsi, %rdx
	subq	(%rcx), %rdx
	movq	%r13, %rbp
	sbbq	8(%rcx), %rbp
	movq	%r12, %rbx
	sbbq	16(%rcx), %rbx
	movq	%r8, %r9
	sbbq	24(%rcx), %r9
	sbbq	$0, %rax
	andl	$1, %eax
	cmovneq	%rsi, %rdx
	movq	%rdx, 32(%rdi)
	testb	%al, %al
	cmovneq	%r13, %rbp
	movq	%rbp, 40(%rdi)
	cmovneq	%r12, %rbx
	movq	%rbx, 48(%rdi)
	cmovneq	%r8, %r9
	movq	%r9, 56(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end63:
	.size	mcl_fpDbl_add4Lbmi2, .Lfunc_end63-mcl_fpDbl_add4Lbmi2

	.globl	mcl_fpDbl_sub4Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sub4Lbmi2,@function
mcl_fpDbl_sub4Lbmi2:                    # @mcl_fpDbl_sub4Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%r12
	pushq	%rbx
	movq	56(%rdx), %r9
	movq	56(%rsi), %r8
	movq	48(%rdx), %r10
	movq	24(%rdx), %r11
	movq	(%rsi), %rbx
	xorl	%eax, %eax
	subq	(%rdx), %rbx
	movq	%rbx, (%rdi)
	movq	8(%rsi), %rbx
	sbbq	8(%rdx), %rbx
	movq	%rbx, 8(%rdi)
	movq	16(%rsi), %rbx
	sbbq	16(%rdx), %rbx
	movq	%rbx, 16(%rdi)
	movq	24(%rsi), %rbx
	sbbq	%r11, %rbx
	movq	40(%rdx), %r11
	movq	32(%rdx), %rdx
	movq	%rbx, 24(%rdi)
	movq	32(%rsi), %r12
	sbbq	%rdx, %r12
	movq	48(%rsi), %r14
	movq	40(%rsi), %r15
	sbbq	%r11, %r15
	sbbq	%r10, %r14
	sbbq	%r9, %r8
	movl	$0, %edx
	sbbq	$0, %rdx
	andl	$1, %edx
	movq	(%rcx), %rsi
	cmoveq	%rax, %rsi
	testb	%dl, %dl
	movq	16(%rcx), %rdx
	cmoveq	%rax, %rdx
	movq	24(%rcx), %rbx
	cmoveq	%rax, %rbx
	cmovneq	8(%rcx), %rax
	addq	%r12, %rsi
	movq	%rsi, 32(%rdi)
	adcq	%r15, %rax
	movq	%rax, 40(%rdi)
	adcq	%r14, %rdx
	movq	%rdx, 48(%rdi)
	adcq	%r8, %rbx
	movq	%rbx, 56(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r14
	popq	%r15
	retq
.Lfunc_end64:
	.size	mcl_fpDbl_sub4Lbmi2, .Lfunc_end64-mcl_fpDbl_sub4Lbmi2

	.globl	mcl_fp_mulUnitPre5Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mulUnitPre5Lbmi2,@function
mcl_fp_mulUnitPre5Lbmi2:                # @mcl_fp_mulUnitPre5Lbmi2
# BB#0:
	pushq	%r14
	pushq	%rbx
	mulxq	32(%rsi), %r8, %r11
	mulxq	24(%rsi), %r9, %rax
	mulxq	16(%rsi), %r10, %rcx
	mulxq	8(%rsi), %r14, %rbx
	mulxq	(%rsi), %rdx, %rsi
	movq	%rdx, (%rdi)
	addq	%r14, %rsi
	movq	%rsi, 8(%rdi)
	adcq	%r10, %rbx
	movq	%rbx, 16(%rdi)
	adcq	%r9, %rcx
	movq	%rcx, 24(%rdi)
	adcq	%r8, %rax
	movq	%rax, 32(%rdi)
	adcq	$0, %r11
	movq	%r11, 40(%rdi)
	popq	%rbx
	popq	%r14
	retq
.Lfunc_end65:
	.size	mcl_fp_mulUnitPre5Lbmi2, .Lfunc_end65-mcl_fp_mulUnitPre5Lbmi2

	.globl	mcl_fpDbl_mulPre5Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_mulPre5Lbmi2,@function
mcl_fpDbl_mulPre5Lbmi2:                 # @mcl_fpDbl_mulPre5Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	movq	%rdi, -8(%rsp)          # 8-byte Spill
	movq	(%rsi), %r11
	movq	8(%rsi), %r10
	movq	(%rdx), %rcx
	movq	%r10, %rdx
	mulxq	%rcx, %rax, %r14
	movq	%r11, %rdx
	mulxq	%rcx, %rdx, %rbx
	movq	%rdx, -24(%rsp)         # 8-byte Spill
	movq	24(%rsi), %rbp
	movq	%rbp, -16(%rsp)         # 8-byte Spill
	movq	16(%rsi), %r15
	addq	%rax, %rbx
	movq	%r15, %rdx
	mulxq	%rcx, %rax, %r13
	adcq	%r14, %rax
	movq	%rbp, %rdx
	mulxq	%rcx, %r8, %r12
	adcq	%r13, %r8
	movq	32(%rsi), %r14
	movq	%r14, %rdx
	mulxq	%rcx, %r9, %r13
	adcq	%r12, %r9
	movq	-24(%rsp), %rcx         # 8-byte Reload
	movq	%rcx, (%rdi)
	adcq	$0, %r13
	movq	-48(%rsp), %rdi         # 8-byte Reload
	movq	8(%rdi), %rbp
	movq	%r11, %rdx
	mulxq	%rbp, %r12, %r11
	addq	%rbx, %r12
	movq	%r10, %rdx
	mulxq	%rbp, %rbx, %rcx
	movq	%rcx, -24(%rsp)         # 8-byte Spill
	adcq	%rax, %rbx
	movq	%r15, %rdx
	mulxq	%rbp, %rcx, %r10
	adcq	%r8, %rcx
	movq	-16(%rsp), %rdx         # 8-byte Reload
	mulxq	%rbp, %rax, %r8
	adcq	%r9, %rax
	movq	%r14, %rdx
	mulxq	%rbp, %r15, %rdx
	adcq	%r13, %r15
	sbbq	%r14, %r14
	andl	$1, %r14d
	addq	%r11, %rbx
	movq	-8(%rsp), %rbp          # 8-byte Reload
	movq	%r12, 8(%rbp)
	adcq	-24(%rsp), %rcx         # 8-byte Folded Reload
	adcq	%r10, %rax
	adcq	%r8, %r15
	adcq	%rdx, %r14
	movq	(%rsi), %rdx
	movq	%rdx, -24(%rsp)         # 8-byte Spill
	movq	8(%rsi), %r8
	movq	%r8, -16(%rsp)          # 8-byte Spill
	movq	16(%rdi), %rbp
	mulxq	%rbp, %r12, %rdx
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	addq	%rbx, %r12
	movq	%r8, %rdx
	mulxq	%rbp, %rbx, %rdx
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	adcq	%rcx, %rbx
	movq	16(%rsi), %r11
	movq	%r11, %rdx
	mulxq	%rbp, %rcx, %rdx
	movq	%rdx, -56(%rsp)         # 8-byte Spill
	adcq	%rax, %rcx
	movq	24(%rsi), %r13
	movq	%r13, %rdx
	mulxq	%rbp, %r9, %r10
	adcq	%r15, %r9
	movq	32(%rsi), %r15
	movq	%r15, %rdx
	mulxq	%rbp, %r8, %rdx
	adcq	%r14, %r8
	sbbq	%r14, %r14
	andl	$1, %r14d
	addq	-32(%rsp), %rbx         # 8-byte Folded Reload
	adcq	-40(%rsp), %rcx         # 8-byte Folded Reload
	adcq	-56(%rsp), %r9          # 8-byte Folded Reload
	adcq	%r10, %r8
	adcq	%rdx, %r14
	movq	-8(%rsp), %r10          # 8-byte Reload
	movq	%r12, 16(%r10)
	movq	%rdi, %rbp
	movq	24(%rbp), %rax
	movq	-24(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %r12, %rdi
	addq	%rbx, %r12
	movq	-16(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %rbx, %rdx
	movq	%rdx, -16(%rsp)         # 8-byte Spill
	adcq	%rcx, %rbx
	movq	%r11, %rdx
	mulxq	%rax, %rcx, %r11
	adcq	%r9, %rcx
	movq	%r13, %rdx
	mulxq	%rax, %r13, %r9
	adcq	%r8, %r13
	movq	%r15, %rdx
	mulxq	%rax, %r8, %rdx
	adcq	%r14, %r8
	sbbq	%r14, %r14
	andl	$1, %r14d
	addq	%rdi, %rbx
	movq	%r12, 24(%r10)
	movq	%r10, %rdi
	adcq	-16(%rsp), %rcx         # 8-byte Folded Reload
	adcq	%r11, %r13
	adcq	%r9, %r8
	adcq	%rdx, %r14
	movq	32(%rbp), %rdx
	mulxq	8(%rsi), %rax, %r9
	mulxq	(%rsi), %rbp, %r10
	addq	%rbx, %rbp
	adcq	%rcx, %rax
	mulxq	16(%rsi), %rbx, %r15
	adcq	%r13, %rbx
	mulxq	32(%rsi), %rcx, %r11
	mulxq	24(%rsi), %rsi, %rdx
	movq	%rbp, 32(%rdi)
	adcq	%r8, %rsi
	adcq	%r14, %rcx
	sbbq	%rbp, %rbp
	andl	$1, %ebp
	addq	%r10, %rax
	movq	%rax, 40(%rdi)
	adcq	%r9, %rbx
	movq	%rbx, 48(%rdi)
	adcq	%r15, %rsi
	movq	%rsi, 56(%rdi)
	adcq	%rdx, %rcx
	movq	%rcx, 64(%rdi)
	adcq	%r11, %rbp
	movq	%rbp, 72(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end66:
	.size	mcl_fpDbl_mulPre5Lbmi2, .Lfunc_end66-mcl_fpDbl_mulPre5Lbmi2

	.globl	mcl_fpDbl_sqrPre5Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sqrPre5Lbmi2,@function
mcl_fpDbl_sqrPre5Lbmi2:                 # @mcl_fpDbl_sqrPre5Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	16(%rsi), %r11
	movq	(%rsi), %rax
	movq	8(%rsi), %rcx
	movq	%r11, %rdx
	mulxq	%rax, %rbx, %r15
	movq	32(%rsi), %r9
	movq	%r9, -8(%rsp)           # 8-byte Spill
	movq	24(%rsi), %r13
	movq	%rcx, %rdx
	mulxq	%rax, %r12, %rbp
	movq	%rbp, -16(%rsp)         # 8-byte Spill
	movq	%rax, %rdx
	mulxq	%rax, %rdx, %r14
	movq	%rdx, -24(%rsp)         # 8-byte Spill
	addq	%r12, %r14
	adcq	%rbp, %rbx
	movq	%r13, %rdx
	mulxq	%rax, %r8, %r10
	adcq	%r15, %r8
	movq	%r9, %rdx
	mulxq	%rax, %rbp, %r15
	adcq	%r10, %rbp
	movq	-24(%rsp), %rax         # 8-byte Reload
	movq	%rax, (%rdi)
	adcq	$0, %r15
	addq	%r12, %r14
	movq	%rcx, %rdx
	mulxq	%rcx, %rax, %rdx
	movq	%rdx, -24(%rsp)         # 8-byte Spill
	adcq	%rbx, %rax
	movq	%r11, %rdx
	mulxq	%rcx, %rbx, %r10
	adcq	%r8, %rbx
	movq	%r13, %rdx
	mulxq	%rcx, %r13, %r8
	adcq	%rbp, %r13
	movq	%r9, %rdx
	mulxq	%rcx, %r12, %rcx
	adcq	%r15, %r12
	sbbq	%r15, %r15
	andl	$1, %r15d
	addq	-16(%rsp), %rax         # 8-byte Folded Reload
	movq	%r14, 8(%rdi)
	adcq	-24(%rsp), %rbx         # 8-byte Folded Reload
	adcq	%r10, %r13
	adcq	%r8, %r12
	adcq	%rcx, %r15
	movq	(%rsi), %r9
	movq	8(%rsi), %r10
	movq	%r9, %rdx
	mulxq	%r11, %rbp, %rcx
	movq	%rcx, -16(%rsp)         # 8-byte Spill
	addq	%rax, %rbp
	movq	%r10, %rdx
	mulxq	%r11, %rax, %r8
	adcq	%rbx, %rax
	movq	%r11, %rdx
	mulxq	%r11, %r14, %rcx
	movq	%rcx, -24(%rsp)         # 8-byte Spill
	adcq	%r13, %r14
	movq	24(%rsi), %rcx
	movq	%rcx, %rdx
	mulxq	%r11, %rbx, %r13
	adcq	%r12, %rbx
	movq	-8(%rsp), %rdx          # 8-byte Reload
	mulxq	%r11, %r12, %rdx
	adcq	%r15, %r12
	sbbq	%r15, %r15
	andl	$1, %r15d
	addq	-16(%rsp), %rax         # 8-byte Folded Reload
	adcq	%r8, %r14
	movq	%rbp, 16(%rdi)
	adcq	-24(%rsp), %rbx         # 8-byte Folded Reload
	adcq	%r13, %r12
	adcq	%rdx, %r15
	movq	%r10, %rdx
	mulxq	%rcx, %r10, %rdx
	movq	%rdx, -8(%rsp)          # 8-byte Spill
	movq	%r9, %rdx
	mulxq	%rcx, %r13, %rdx
	movq	%rdx, -16(%rsp)         # 8-byte Spill
	addq	%rax, %r13
	movq	16(%rsi), %r8
	movq	32(%rsi), %rax
	adcq	%r14, %r10
	movq	%r8, %rdx
	mulxq	%rcx, %r9, %r14
	adcq	%rbx, %r9
	movq	%rcx, %rdx
	mulxq	%rcx, %r11, %rbp
	adcq	%r12, %r11
	movq	%rax, %rdx
	mulxq	%rcx, %r12, %rdx
	adcq	%r15, %r12
	sbbq	%rbx, %rbx
	andl	$1, %ebx
	addq	-16(%rsp), %r10         # 8-byte Folded Reload
	movq	%r13, 24(%rdi)
	adcq	-8(%rsp), %r9           # 8-byte Folded Reload
	adcq	%r14, %r11
	adcq	%rbp, %r12
	adcq	%rdx, %rbx
	movq	%rax, %rdx
	mulxq	24(%rsi), %rbp, %r14
	movq	%rax, %rdx
	mulxq	(%rsi), %rcx, %r15
	addq	%r10, %rcx
	movq	%rax, %rdx
	mulxq	8(%rsi), %rsi, %r10
	movq	%rcx, 32(%rdi)
	adcq	%r9, %rsi
	movq	%r8, %rdx
	mulxq	%rax, %rcx, %r8
	adcq	%r11, %rcx
	adcq	%r12, %rbp
	movq	%rax, %rdx
	mulxq	%rax, %rdx, %rax
	adcq	%rbx, %rdx
	sbbq	%rbx, %rbx
	andl	$1, %ebx
	addq	%r15, %rsi
	movq	%rsi, 40(%rdi)
	adcq	%r10, %rcx
	movq	%rcx, 48(%rdi)
	adcq	%r8, %rbp
	movq	%rbp, 56(%rdi)
	adcq	%r14, %rdx
	movq	%rdx, 64(%rdi)
	adcq	%rax, %rbx
	movq	%rbx, 72(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end67:
	.size	mcl_fpDbl_sqrPre5Lbmi2, .Lfunc_end67-mcl_fpDbl_sqrPre5Lbmi2

	.globl	mcl_fp_mont5Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mont5Lbmi2,@function
mcl_fp_mont5Lbmi2:                      # @mcl_fp_mont5Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rdx, -56(%rsp)         # 8-byte Spill
	movq	%rdi, -112(%rsp)        # 8-byte Spill
	movq	32(%rsi), %rdi
	movq	%rdi, -64(%rsp)         # 8-byte Spill
	movq	(%rdx), %rax
	movq	%rdi, %rdx
	mulxq	%rax, %r10, %rbx
	movq	24(%rsi), %rdx
	movq	%rdx, -72(%rsp)         # 8-byte Spill
	mulxq	%rax, %r12, %r14
	movq	16(%rsi), %rdx
	movq	%rdx, -80(%rsp)         # 8-byte Spill
	mulxq	%rax, %r13, %r11
	movq	(%rsi), %rbp
	movq	%rbp, -88(%rsp)         # 8-byte Spill
	movq	8(%rsi), %rdx
	movq	%rdx, -96(%rsp)         # 8-byte Spill
	mulxq	%rax, %rdi, %r9
	movq	%rbp, %rdx
	mulxq	%rax, %r15, %r8
	addq	%rdi, %r8
	adcq	%r13, %r9
	adcq	%r12, %r11
	adcq	%r10, %r14
	adcq	$0, %rbx
	movq	%rbx, -104(%rsp)        # 8-byte Spill
	movq	-8(%rcx), %rax
	movq	%rax, -48(%rsp)         # 8-byte Spill
	movq	%r15, %rdx
	imulq	%rax, %rdx
	movq	(%rcx), %rsi
	movq	%rsi, -32(%rsp)         # 8-byte Spill
	movq	32(%rcx), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	mulxq	%rax, %rax, %r12
	movq	%rax, -120(%rsp)        # 8-byte Spill
	movq	24(%rcx), %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	mulxq	%rax, %r13, %r10
	movq	8(%rcx), %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	mulxq	%rax, %rdi, %rbp
	mulxq	%rsi, %rax, %rbx
	addq	%rdi, %rbx
	movq	16(%rcx), %rcx
	movq	%rcx, -40(%rsp)         # 8-byte Spill
	mulxq	%rcx, %rdi, %rcx
	adcq	%rbp, %rdi
	adcq	%r13, %rcx
	adcq	-120(%rsp), %r10        # 8-byte Folded Reload
	adcq	$0, %r12
	addq	%r15, %rax
	adcq	%r8, %rbx
	adcq	%r9, %rdi
	adcq	%r11, %rcx
	adcq	%r14, %r10
	adcq	-104(%rsp), %r12        # 8-byte Folded Reload
	sbbq	%rbp, %rbp
	andl	$1, %ebp
	movq	-56(%rsp), %rax         # 8-byte Reload
	movq	8(%rax), %rdx
	mulxq	-64(%rsp), %rax, %r14   # 8-byte Folded Reload
	movq	%rax, -104(%rsp)        # 8-byte Spill
	mulxq	-72(%rsp), %rax, %r15   # 8-byte Folded Reload
	movq	%rax, -120(%rsp)        # 8-byte Spill
	mulxq	-80(%rsp), %r13, %r9    # 8-byte Folded Reload
	mulxq	-96(%rsp), %r8, %rsi    # 8-byte Folded Reload
	mulxq	-88(%rsp), %r11, %rax   # 8-byte Folded Reload
	addq	%r8, %rax
	adcq	%r13, %rsi
	adcq	-120(%rsp), %r9         # 8-byte Folded Reload
	adcq	-104(%rsp), %r15        # 8-byte Folded Reload
	adcq	$0, %r14
	addq	%rbx, %r11
	adcq	%rdi, %rax
	adcq	%rcx, %rsi
	adcq	%r10, %r9
	adcq	%r12, %r15
	adcq	%rbp, %r14
	sbbq	%r12, %r12
	andl	$1, %r12d
	movq	%r11, %rdx
	imulq	-48(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	-8(%rsp), %rcx, %r10    # 8-byte Folded Reload
	movq	%rcx, -104(%rsp)        # 8-byte Spill
	mulxq	-16(%rsp), %rcx, %rdi   # 8-byte Folded Reload
	movq	%rcx, -120(%rsp)        # 8-byte Spill
	mulxq	-40(%rsp), %r13, %rcx   # 8-byte Folded Reload
	mulxq	-24(%rsp), %r8, %rbx    # 8-byte Folded Reload
	mulxq	-32(%rsp), %rdx, %rbp   # 8-byte Folded Reload
	addq	%r8, %rbp
	adcq	%r13, %rbx
	adcq	-120(%rsp), %rcx        # 8-byte Folded Reload
	adcq	-104(%rsp), %rdi        # 8-byte Folded Reload
	adcq	$0, %r10
	addq	%r11, %rdx
	adcq	%rax, %rbp
	adcq	%rsi, %rbx
	adcq	%r9, %rcx
	adcq	%r15, %rdi
	adcq	%r14, %r10
	adcq	$0, %r12
	movq	-56(%rsp), %rax         # 8-byte Reload
	movq	16(%rax), %rdx
	mulxq	-64(%rsp), %rax, %r15   # 8-byte Folded Reload
	movq	%rax, -104(%rsp)        # 8-byte Spill
	mulxq	-72(%rsp), %rax, %r11   # 8-byte Folded Reload
	movq	%rax, -120(%rsp)        # 8-byte Spill
	mulxq	-80(%rsp), %r13, %r9    # 8-byte Folded Reload
	mulxq	-96(%rsp), %rsi, %r8    # 8-byte Folded Reload
	mulxq	-88(%rsp), %r14, %rax   # 8-byte Folded Reload
	addq	%rsi, %rax
	adcq	%r13, %r8
	adcq	-120(%rsp), %r9         # 8-byte Folded Reload
	adcq	-104(%rsp), %r11        # 8-byte Folded Reload
	adcq	$0, %r15
	addq	%rbp, %r14
	adcq	%rbx, %rax
	adcq	%rcx, %r8
	adcq	%rdi, %r9
	adcq	%r10, %r11
	adcq	%r12, %r15
	sbbq	%r13, %r13
	andl	$1, %r13d
	movq	%r14, %rdx
	imulq	-48(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	-8(%rsp), %rcx, %r12    # 8-byte Folded Reload
	movq	%rcx, -104(%rsp)        # 8-byte Spill
	mulxq	-16(%rsp), %rcx, %r10   # 8-byte Folded Reload
	movq	%rcx, -120(%rsp)        # 8-byte Spill
	mulxq	-40(%rsp), %rdi, %rsi   # 8-byte Folded Reload
	mulxq	-24(%rsp), %rcx, %rbx   # 8-byte Folded Reload
	mulxq	-32(%rsp), %rdx, %rbp   # 8-byte Folded Reload
	addq	%rcx, %rbp
	adcq	%rdi, %rbx
	adcq	-120(%rsp), %rsi        # 8-byte Folded Reload
	adcq	-104(%rsp), %r10        # 8-byte Folded Reload
	adcq	$0, %r12
	addq	%r14, %rdx
	adcq	%rax, %rbp
	adcq	%r8, %rbx
	adcq	%r9, %rsi
	adcq	%r11, %r10
	adcq	%r15, %r12
	adcq	$0, %r13
	movq	-56(%rsp), %rax         # 8-byte Reload
	movq	24(%rax), %rdx
	mulxq	-64(%rsp), %rcx, %rax   # 8-byte Folded Reload
	movq	%rcx, -120(%rsp)        # 8-byte Spill
	movq	%rax, -104(%rsp)        # 8-byte Spill
	mulxq	-72(%rsp), %r11, %r14   # 8-byte Folded Reload
	mulxq	-80(%rsp), %r8, %r9     # 8-byte Folded Reload
	mulxq	-96(%rsp), %rax, %rdi   # 8-byte Folded Reload
	mulxq	-88(%rsp), %r15, %rcx   # 8-byte Folded Reload
	addq	%rax, %rcx
	adcq	%r8, %rdi
	adcq	%r11, %r9
	adcq	-120(%rsp), %r14        # 8-byte Folded Reload
	movq	-104(%rsp), %rax        # 8-byte Reload
	adcq	$0, %rax
	addq	%rbp, %r15
	adcq	%rbx, %rcx
	adcq	%rsi, %rdi
	adcq	%r10, %r9
	adcq	%r12, %r14
	adcq	%r13, %rax
	movq	%rax, -104(%rsp)        # 8-byte Spill
	sbbq	%r12, %r12
	andl	$1, %r12d
	movq	%r15, %rdx
	imulq	-48(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	-8(%rsp), %rax, %rbp    # 8-byte Folded Reload
	movq	%rax, -120(%rsp)        # 8-byte Spill
	mulxq	-16(%rsp), %r13, %r10   # 8-byte Folded Reload
	mulxq	-40(%rsp), %rbx, %r8    # 8-byte Folded Reload
	mulxq	-24(%rsp), %rsi, %r11   # 8-byte Folded Reload
	mulxq	-32(%rsp), %rdx, %rax   # 8-byte Folded Reload
	addq	%rsi, %rax
	adcq	%rbx, %r11
	adcq	%r13, %r8
	adcq	-120(%rsp), %r10        # 8-byte Folded Reload
	adcq	$0, %rbp
	addq	%r15, %rdx
	adcq	%rcx, %rax
	adcq	%rdi, %r11
	adcq	%r9, %r8
	adcq	%r14, %r10
	adcq	-104(%rsp), %rbp        # 8-byte Folded Reload
	adcq	$0, %r12
	movq	-56(%rsp), %rcx         # 8-byte Reload
	movq	32(%rcx), %rdx
	mulxq	-64(%rsp), %rcx, %r14   # 8-byte Folded Reload
	movq	%rcx, -56(%rsp)         # 8-byte Spill
	mulxq	-72(%rsp), %rcx, %rbx   # 8-byte Folded Reload
	movq	%rcx, -64(%rsp)         # 8-byte Spill
	mulxq	-80(%rsp), %rsi, %r15   # 8-byte Folded Reload
	mulxq	-96(%rsp), %rcx, %r9    # 8-byte Folded Reload
	mulxq	-88(%rsp), %r13, %rdi   # 8-byte Folded Reload
	addq	%rcx, %rdi
	adcq	%rsi, %r9
	adcq	-64(%rsp), %r15         # 8-byte Folded Reload
	adcq	-56(%rsp), %rbx         # 8-byte Folded Reload
	adcq	$0, %r14
	addq	%rax, %r13
	adcq	%r11, %rdi
	adcq	%r8, %r9
	adcq	%r10, %r15
	adcq	%rbp, %rbx
	adcq	%r12, %r14
	sbbq	%rax, %rax
	movq	-48(%rsp), %rdx         # 8-byte Reload
	imulq	%r13, %rdx
	mulxq	-32(%rsp), %r10, %rcx   # 8-byte Folded Reload
	mulxq	-24(%rsp), %r8, %rsi    # 8-byte Folded Reload
	addq	%rcx, %r8
	mulxq	-40(%rsp), %rbp, %r11   # 8-byte Folded Reload
	adcq	%rsi, %rbp
	mulxq	-16(%rsp), %rcx, %r12   # 8-byte Folded Reload
	adcq	%r11, %rcx
	mulxq	-8(%rsp), %rsi, %r11    # 8-byte Folded Reload
	adcq	%r12, %rsi
	adcq	$0, %r11
	andl	$1, %eax
	addq	%r13, %r10
	adcq	%rdi, %r8
	adcq	%r9, %rbp
	adcq	%r15, %rcx
	adcq	%rbx, %rsi
	adcq	%r14, %r11
	adcq	$0, %rax
	movq	%r8, %rdi
	subq	-32(%rsp), %rdi         # 8-byte Folded Reload
	movq	%rbp, %rbx
	sbbq	-24(%rsp), %rbx         # 8-byte Folded Reload
	movq	%rcx, %r9
	sbbq	-40(%rsp), %r9          # 8-byte Folded Reload
	movq	%rsi, %rdx
	sbbq	-16(%rsp), %rdx         # 8-byte Folded Reload
	movq	%r11, %r10
	sbbq	-8(%rsp), %r10          # 8-byte Folded Reload
	sbbq	$0, %rax
	andl	$1, %eax
	cmovneq	%rsi, %rdx
	testb	%al, %al
	cmovneq	%r8, %rdi
	movq	-112(%rsp), %rax        # 8-byte Reload
	movq	%rdi, (%rax)
	cmovneq	%rbp, %rbx
	movq	%rbx, 8(%rax)
	cmovneq	%rcx, %r9
	movq	%r9, 16(%rax)
	movq	%rdx, 24(%rax)
	cmovneq	%r11, %r10
	movq	%r10, 32(%rax)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end68:
	.size	mcl_fp_mont5Lbmi2, .Lfunc_end68-mcl_fp_mont5Lbmi2

	.globl	mcl_fp_montNF5Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montNF5Lbmi2,@function
mcl_fp_montNF5Lbmi2:                    # @mcl_fp_montNF5Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rdx, -16(%rsp)         # 8-byte Spill
	movq	%rdi, -104(%rsp)        # 8-byte Spill
	movq	(%rsi), %r13
	movq	%r13, -64(%rsp)         # 8-byte Spill
	movq	8(%rsi), %rbp
	movq	%rbp, -24(%rsp)         # 8-byte Spill
	movq	(%rdx), %rax
	movq	%rbp, %rdx
	mulxq	%rax, %rbp, %r9
	movq	%r13, %rdx
	mulxq	%rax, %r8, %r10
	movq	16(%rsi), %rdx
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	addq	%rbp, %r10
	mulxq	%rax, %rbp, %rbx
	adcq	%r9, %rbp
	movq	24(%rsi), %rdx
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	mulxq	%rax, %r15, %r9
	adcq	%rbx, %r15
	movq	32(%rsi), %rdx
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	mulxq	%rax, %rax, %r11
	adcq	%r9, %rax
	adcq	$0, %r11
	movq	-8(%rcx), %rsi
	movq	%rsi, -72(%rsp)         # 8-byte Spill
	movq	%r8, %rdx
	imulq	%rsi, %rdx
	movq	(%rcx), %rsi
	movq	%rsi, -88(%rsp)         # 8-byte Spill
	mulxq	%rsi, %rbx, %r14
	addq	%r8, %rbx
	movq	8(%rcx), %rsi
	movq	%rsi, -80(%rsp)         # 8-byte Spill
	mulxq	%rsi, %rbx, %r12
	adcq	%r10, %rbx
	movq	16(%rcx), %rsi
	movq	%rsi, -96(%rsp)         # 8-byte Spill
	mulxq	%rsi, %r10, %rdi
	adcq	%rbp, %r10
	movq	24(%rcx), %rsi
	movq	%rsi, -56(%rsp)         # 8-byte Spill
	mulxq	%rsi, %r9, %rbp
	adcq	%r15, %r9
	movq	32(%rcx), %rcx
	movq	%rcx, -8(%rsp)          # 8-byte Spill
	mulxq	%rcx, %r8, %rcx
	adcq	%rax, %r8
	adcq	$0, %r11
	addq	%r14, %rbx
	adcq	%r12, %r10
	adcq	%rdi, %r9
	adcq	%rbp, %r8
	adcq	%rcx, %r11
	movq	-16(%rsp), %rax         # 8-byte Reload
	movq	8(%rax), %rdx
	mulxq	-24(%rsp), %rcx, %rsi   # 8-byte Folded Reload
	mulxq	%r13, %r14, %rax
	addq	%rcx, %rax
	mulxq	-32(%rsp), %rcx, %rdi   # 8-byte Folded Reload
	adcq	%rsi, %rcx
	mulxq	-40(%rsp), %rsi, %r15   # 8-byte Folded Reload
	adcq	%rdi, %rsi
	mulxq	-48(%rsp), %rdi, %rbp   # 8-byte Folded Reload
	adcq	%r15, %rdi
	adcq	$0, %rbp
	addq	%rbx, %r14
	adcq	%r10, %rax
	adcq	%r9, %rcx
	adcq	%r8, %rsi
	adcq	%r11, %rdi
	adcq	$0, %rbp
	movq	%r14, %rdx
	movq	-72(%rsp), %r12         # 8-byte Reload
	imulq	%r12, %rdx
	mulxq	-88(%rsp), %rbx, %r15   # 8-byte Folded Reload
	addq	%r14, %rbx
	movq	-80(%rsp), %r13         # 8-byte Reload
	mulxq	%r13, %r8, %rbx
	adcq	%rax, %r8
	mulxq	-96(%rsp), %r9, %rax    # 8-byte Folded Reload
	adcq	%rcx, %r9
	mulxq	-56(%rsp), %r10, %rcx   # 8-byte Folded Reload
	adcq	%rsi, %r10
	mulxq	-8(%rsp), %r11, %rdx    # 8-byte Folded Reload
	adcq	%rdi, %r11
	adcq	$0, %rbp
	addq	%r15, %r8
	adcq	%rbx, %r9
	adcq	%rax, %r10
	adcq	%rcx, %r11
	adcq	%rdx, %rbp
	movq	-16(%rsp), %rax         # 8-byte Reload
	movq	16(%rax), %rdx
	mulxq	-24(%rsp), %rcx, %rax   # 8-byte Folded Reload
	mulxq	-64(%rsp), %r14, %rsi   # 8-byte Folded Reload
	addq	%rcx, %rsi
	mulxq	-32(%rsp), %rbx, %rcx   # 8-byte Folded Reload
	adcq	%rax, %rbx
	mulxq	-40(%rsp), %rdi, %r15   # 8-byte Folded Reload
	adcq	%rcx, %rdi
	mulxq	-48(%rsp), %rcx, %rax   # 8-byte Folded Reload
	adcq	%r15, %rcx
	adcq	$0, %rax
	addq	%r8, %r14
	adcq	%r9, %rsi
	adcq	%r10, %rbx
	adcq	%r11, %rdi
	adcq	%rbp, %rcx
	adcq	$0, %rax
	movq	%r14, %rdx
	imulq	%r12, %rdx
	movq	-88(%rsp), %r12         # 8-byte Reload
	mulxq	%r12, %rbp, %r15
	addq	%r14, %rbp
	mulxq	%r13, %r8, %rbp
	adcq	%rsi, %r8
	movq	-96(%rsp), %r13         # 8-byte Reload
	mulxq	%r13, %r9, %rsi
	adcq	%rbx, %r9
	mulxq	-56(%rsp), %r10, %rbx   # 8-byte Folded Reload
	adcq	%rdi, %r10
	mulxq	-8(%rsp), %r11, %rdx    # 8-byte Folded Reload
	adcq	%rcx, %r11
	adcq	$0, %rax
	addq	%r15, %r8
	adcq	%rbp, %r9
	adcq	%rsi, %r10
	adcq	%rbx, %r11
	adcq	%rdx, %rax
	movq	-16(%rsp), %rcx         # 8-byte Reload
	movq	24(%rcx), %rdx
	mulxq	-24(%rsp), %rdi, %rsi   # 8-byte Folded Reload
	mulxq	-64(%rsp), %r14, %rcx   # 8-byte Folded Reload
	addq	%rdi, %rcx
	mulxq	-32(%rsp), %rbx, %rdi   # 8-byte Folded Reload
	adcq	%rsi, %rbx
	mulxq	-40(%rsp), %rsi, %r15   # 8-byte Folded Reload
	adcq	%rdi, %rsi
	mulxq	-48(%rsp), %rdi, %rbp   # 8-byte Folded Reload
	adcq	%r15, %rdi
	adcq	$0, %rbp
	addq	%r8, %r14
	adcq	%r9, %rcx
	adcq	%r10, %rbx
	adcq	%r11, %rsi
	adcq	%rax, %rdi
	adcq	$0, %rbp
	movq	%r14, %rdx
	imulq	-72(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	%r12, %rax, %r11
	addq	%r14, %rax
	mulxq	-80(%rsp), %r8, %r14    # 8-byte Folded Reload
	adcq	%rcx, %r8
	mulxq	%r13, %r9, %rax
	adcq	%rbx, %r9
	movq	-56(%rsp), %r12         # 8-byte Reload
	mulxq	%r12, %r10, %rbx
	adcq	%rsi, %r10
	mulxq	-8(%rsp), %rcx, %rdx    # 8-byte Folded Reload
	adcq	%rdi, %rcx
	adcq	$0, %rbp
	addq	%r11, %r8
	adcq	%r14, %r9
	adcq	%rax, %r10
	adcq	%rbx, %rcx
	adcq	%rdx, %rbp
	movq	-16(%rsp), %rax         # 8-byte Reload
	movq	32(%rax), %rdx
	mulxq	-24(%rsp), %rdi, %rbx   # 8-byte Folded Reload
	mulxq	-64(%rsp), %r14, %rsi   # 8-byte Folded Reload
	addq	%rdi, %rsi
	mulxq	-32(%rsp), %rdi, %rax   # 8-byte Folded Reload
	adcq	%rbx, %rdi
	mulxq	-40(%rsp), %rbx, %r15   # 8-byte Folded Reload
	adcq	%rax, %rbx
	mulxq	-48(%rsp), %r11, %rax   # 8-byte Folded Reload
	adcq	%r15, %r11
	adcq	$0, %rax
	addq	%r8, %r14
	adcq	%r9, %rsi
	adcq	%r10, %rdi
	adcq	%rcx, %rbx
	adcq	%rbp, %r11
	adcq	$0, %rax
	movq	-72(%rsp), %rdx         # 8-byte Reload
	imulq	%r14, %rdx
	movq	-88(%rsp), %r10         # 8-byte Reload
	mulxq	%r10, %rcx, %rbp
	movq	%rbp, -16(%rsp)         # 8-byte Spill
	addq	%r14, %rcx
	movq	-80(%rsp), %r9          # 8-byte Reload
	mulxq	%r9, %r14, %rcx
	movq	%rcx, -24(%rsp)         # 8-byte Spill
	adcq	%rsi, %r14
	movq	%r13, %r8
	mulxq	%r8, %r15, %r13
	adcq	%rdi, %r15
	mulxq	%r12, %rbp, %rcx
	adcq	%rbx, %rbp
	movq	-8(%rsp), %rbx          # 8-byte Reload
	mulxq	%rbx, %r12, %rdx
	adcq	%r11, %r12
	adcq	$0, %rax
	addq	-16(%rsp), %r14         # 8-byte Folded Reload
	adcq	-24(%rsp), %r15         # 8-byte Folded Reload
	adcq	%r13, %rbp
	adcq	%rcx, %r12
	adcq	%rdx, %rax
	movq	%r14, %rcx
	subq	%r10, %rcx
	movq	%r15, %rsi
	sbbq	%r9, %rsi
	movq	%rbp, %rdi
	sbbq	%r8, %rdi
	movq	%r12, %r8
	sbbq	-56(%rsp), %r8          # 8-byte Folded Reload
	movq	%rax, %rdx
	sbbq	%rbx, %rdx
	movq	%rdx, %rbx
	sarq	$63, %rbx
	cmovsq	%r14, %rcx
	movq	-104(%rsp), %rbx        # 8-byte Reload
	movq	%rcx, (%rbx)
	cmovsq	%r15, %rsi
	movq	%rsi, 8(%rbx)
	cmovsq	%rbp, %rdi
	movq	%rdi, 16(%rbx)
	cmovsq	%r12, %r8
	movq	%r8, 24(%rbx)
	cmovsq	%rax, %rdx
	movq	%rdx, 32(%rbx)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end69:
	.size	mcl_fp_montNF5Lbmi2, .Lfunc_end69-mcl_fp_montNF5Lbmi2

	.globl	mcl_fp_montRed5Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montRed5Lbmi2,@function
mcl_fp_montRed5Lbmi2:                   # @mcl_fp_montRed5Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rdx, %rcx
	movq	%rdi, -80(%rsp)         # 8-byte Spill
	movq	-8(%rcx), %rax
	movq	%rax, -64(%rsp)         # 8-byte Spill
	movq	(%rcx), %rbx
	movq	%rbx, -8(%rsp)          # 8-byte Spill
	movq	(%rsi), %r9
	movq	%r9, %rdx
	imulq	%rax, %rdx
	movq	%rax, %r15
	movq	32(%rcx), %rax
	movq	%rax, -40(%rsp)         # 8-byte Spill
	mulxq	%rax, %r8, %r13
	movq	24(%rcx), %r12
	movq	%r12, -32(%rsp)         # 8-byte Spill
	mulxq	%r12, %r10, %r14
	movq	16(%rcx), %rax
	movq	%rax, -48(%rsp)         # 8-byte Spill
	movq	8(%rcx), %rcx
	movq	%rcx, -16(%rsp)         # 8-byte Spill
	mulxq	%rax, %rdi, %rbp
	mulxq	%rcx, %rax, %r11
	mulxq	%rbx, %rdx, %rcx
	addq	%rax, %rcx
	adcq	%rdi, %r11
	adcq	%r10, %rbp
	adcq	%r8, %r14
	adcq	$0, %r13
	addq	%r9, %rdx
	movq	72(%rsi), %rax
	movq	64(%rsi), %rdx
	adcq	8(%rsi), %rcx
	adcq	16(%rsi), %r11
	adcq	24(%rsi), %rbp
	adcq	32(%rsi), %r14
	adcq	40(%rsi), %r13
	movq	56(%rsi), %rdi
	movq	48(%rsi), %rsi
	adcq	$0, %rsi
	movq	%rsi, -88(%rsp)         # 8-byte Spill
	adcq	$0, %rdi
	movq	%rdi, -72(%rsp)         # 8-byte Spill
	adcq	$0, %rdx
	movq	%rdx, -56(%rsp)         # 8-byte Spill
	adcq	$0, %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	sbbq	%rsi, %rsi
	andl	$1, %esi
	movq	%rcx, %rdx
	imulq	%r15, %rdx
	mulxq	-40(%rsp), %rax, %r15   # 8-byte Folded Reload
	movq	%rax, -96(%rsp)         # 8-byte Spill
	mulxq	%r12, %rax, %r10
	movq	%rax, -104(%rsp)        # 8-byte Spill
	movq	-48(%rsp), %r12         # 8-byte Reload
	mulxq	%r12, %rbx, %r8
	mulxq	-16(%rsp), %r9, %rdi    # 8-byte Folded Reload
	mulxq	-8(%rsp), %rdx, %rax    # 8-byte Folded Reload
	addq	%r9, %rax
	adcq	%rbx, %rdi
	adcq	-104(%rsp), %r8         # 8-byte Folded Reload
	adcq	-96(%rsp), %r10         # 8-byte Folded Reload
	adcq	$0, %r15
	addq	%rcx, %rdx
	adcq	%r11, %rax
	adcq	%rbp, %rdi
	adcq	%r14, %r8
	adcq	%r13, %r10
	adcq	-88(%rsp), %r15         # 8-byte Folded Reload
	adcq	$0, -72(%rsp)           # 8-byte Folded Spill
	adcq	$0, -56(%rsp)           # 8-byte Folded Spill
	adcq	$0, -24(%rsp)           # 8-byte Folded Spill
	adcq	$0, %rsi
	movq	%rax, %rdx
	imulq	-64(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	-40(%rsp), %rcx, %r13   # 8-byte Folded Reload
	movq	%rcx, -88(%rsp)         # 8-byte Spill
	mulxq	-32(%rsp), %rcx, %r14   # 8-byte Folded Reload
	movq	%rcx, -96(%rsp)         # 8-byte Spill
	mulxq	%r12, %r11, %rbx
	mulxq	-16(%rsp), %r9, %rbp    # 8-byte Folded Reload
	mulxq	-8(%rsp), %rdx, %rcx    # 8-byte Folded Reload
	addq	%r9, %rcx
	adcq	%r11, %rbp
	adcq	-96(%rsp), %rbx         # 8-byte Folded Reload
	adcq	-88(%rsp), %r14         # 8-byte Folded Reload
	adcq	$0, %r13
	addq	%rax, %rdx
	adcq	%rdi, %rcx
	adcq	%r8, %rbp
	adcq	%r10, %rbx
	adcq	%r15, %r14
	adcq	-72(%rsp), %r13         # 8-byte Folded Reload
	adcq	$0, -56(%rsp)           # 8-byte Folded Spill
	adcq	$0, -24(%rsp)           # 8-byte Folded Spill
	adcq	$0, %rsi
	movq	%rcx, %rdx
	imulq	-64(%rsp), %rdx         # 8-byte Folded Reload
	movq	-40(%rsp), %r9          # 8-byte Reload
	mulxq	%r9, %rax, %r12
	movq	%rax, -72(%rsp)         # 8-byte Spill
	mulxq	-32(%rsp), %rax, %r10   # 8-byte Folded Reload
	movq	%rax, -88(%rsp)         # 8-byte Spill
	mulxq	-48(%rsp), %r8, %r11    # 8-byte Folded Reload
	mulxq	-16(%rsp), %rdi, %r15   # 8-byte Folded Reload
	mulxq	-8(%rsp), %rdx, %rax    # 8-byte Folded Reload
	addq	%rdi, %rax
	adcq	%r8, %r15
	adcq	-88(%rsp), %r11         # 8-byte Folded Reload
	adcq	-72(%rsp), %r10         # 8-byte Folded Reload
	adcq	$0, %r12
	addq	%rcx, %rdx
	adcq	%rbp, %rax
	adcq	%rbx, %r15
	adcq	%r14, %r11
	adcq	%r13, %r10
	adcq	-56(%rsp), %r12         # 8-byte Folded Reload
	adcq	$0, -24(%rsp)           # 8-byte Folded Spill
	adcq	$0, %rsi
	movq	-64(%rsp), %rdx         # 8-byte Reload
	imulq	%rax, %rdx
	mulxq	%r9, %rdi, %rcx
	movq	%rdi, -56(%rsp)         # 8-byte Spill
	mulxq	-32(%rsp), %rbp, %rdi   # 8-byte Folded Reload
	movq	%rbp, -64(%rsp)         # 8-byte Spill
	mulxq	-48(%rsp), %r13, %rbp   # 8-byte Folded Reload
	mulxq	-8(%rsp), %r8, %r9      # 8-byte Folded Reload
	movq	-16(%rsp), %r14         # 8-byte Reload
	mulxq	%r14, %rbx, %rdx
	addq	%r9, %rbx
	adcq	%r13, %rdx
	adcq	-64(%rsp), %rbp         # 8-byte Folded Reload
	adcq	-56(%rsp), %rdi         # 8-byte Folded Reload
	adcq	$0, %rcx
	addq	%rax, %r8
	adcq	%r15, %rbx
	adcq	%r11, %rdx
	adcq	%r10, %rbp
	adcq	%r12, %rdi
	adcq	-24(%rsp), %rcx         # 8-byte Folded Reload
	adcq	$0, %rsi
	movq	%rbx, %rax
	subq	-8(%rsp), %rax          # 8-byte Folded Reload
	movq	%rdx, %r8
	sbbq	%r14, %r8
	movq	%rbp, %r9
	sbbq	-48(%rsp), %r9          # 8-byte Folded Reload
	movq	%rdi, %r10
	sbbq	-32(%rsp), %r10         # 8-byte Folded Reload
	movq	%rcx, %r11
	sbbq	-40(%rsp), %r11         # 8-byte Folded Reload
	sbbq	$0, %rsi
	andl	$1, %esi
	cmovneq	%rcx, %r11
	testb	%sil, %sil
	cmovneq	%rbx, %rax
	movq	-80(%rsp), %rcx         # 8-byte Reload
	movq	%rax, (%rcx)
	cmovneq	%rdx, %r8
	movq	%r8, 8(%rcx)
	cmovneq	%rbp, %r9
	movq	%r9, 16(%rcx)
	cmovneq	%rdi, %r10
	movq	%r10, 24(%rcx)
	movq	%r11, 32(%rcx)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end70:
	.size	mcl_fp_montRed5Lbmi2, .Lfunc_end70-mcl_fp_montRed5Lbmi2

	.globl	mcl_fp_addPre5Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addPre5Lbmi2,@function
mcl_fp_addPre5Lbmi2:                    # @mcl_fp_addPre5Lbmi2
# BB#0:
	movq	32(%rdx), %r8
	movq	24(%rdx), %r9
	movq	24(%rsi), %r11
	movq	32(%rsi), %r10
	movq	16(%rdx), %rcx
	movq	(%rdx), %rax
	movq	8(%rdx), %rdx
	addq	(%rsi), %rax
	adcq	8(%rsi), %rdx
	adcq	16(%rsi), %rcx
	movq	%rax, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%rcx, 16(%rdi)
	adcq	%r9, %r11
	movq	%r11, 24(%rdi)
	adcq	%r8, %r10
	movq	%r10, 32(%rdi)
	sbbq	%rax, %rax
	andl	$1, %eax
	retq
.Lfunc_end71:
	.size	mcl_fp_addPre5Lbmi2, .Lfunc_end71-mcl_fp_addPre5Lbmi2

	.globl	mcl_fp_subPre5Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subPre5Lbmi2,@function
mcl_fp_subPre5Lbmi2:                    # @mcl_fp_subPre5Lbmi2
# BB#0:
	pushq	%rbx
	movq	32(%rsi), %r10
	movq	24(%rdx), %r8
	movq	32(%rdx), %r9
	movq	24(%rsi), %r11
	movq	16(%rsi), %rcx
	movq	(%rsi), %rbx
	movq	8(%rsi), %rsi
	xorl	%eax, %eax
	subq	(%rdx), %rbx
	sbbq	8(%rdx), %rsi
	sbbq	16(%rdx), %rcx
	movq	%rbx, (%rdi)
	movq	%rsi, 8(%rdi)
	movq	%rcx, 16(%rdi)
	sbbq	%r8, %r11
	movq	%r11, 24(%rdi)
	sbbq	%r9, %r10
	movq	%r10, 32(%rdi)
	sbbq	$0, %rax
	andl	$1, %eax
	popq	%rbx
	retq
.Lfunc_end72:
	.size	mcl_fp_subPre5Lbmi2, .Lfunc_end72-mcl_fp_subPre5Lbmi2

	.globl	mcl_fp_shr1_5Lbmi2
	.align	16, 0x90
	.type	mcl_fp_shr1_5Lbmi2,@function
mcl_fp_shr1_5Lbmi2:                     # @mcl_fp_shr1_5Lbmi2
# BB#0:
	movq	32(%rsi), %r8
	movq	24(%rsi), %rcx
	movq	16(%rsi), %rdx
	movq	(%rsi), %rax
	movq	8(%rsi), %rsi
	shrdq	$1, %rsi, %rax
	movq	%rax, (%rdi)
	shrdq	$1, %rdx, %rsi
	movq	%rsi, 8(%rdi)
	shrdq	$1, %rcx, %rdx
	movq	%rdx, 16(%rdi)
	shrdq	$1, %r8, %rcx
	movq	%rcx, 24(%rdi)
	shrq	%r8
	movq	%r8, 32(%rdi)
	retq
.Lfunc_end73:
	.size	mcl_fp_shr1_5Lbmi2, .Lfunc_end73-mcl_fp_shr1_5Lbmi2

	.globl	mcl_fp_add5Lbmi2
	.align	16, 0x90
	.type	mcl_fp_add5Lbmi2,@function
mcl_fp_add5Lbmi2:                       # @mcl_fp_add5Lbmi2
# BB#0:
	pushq	%rbx
	movq	32(%rdx), %r11
	movq	24(%rdx), %rbx
	movq	24(%rsi), %r9
	movq	32(%rsi), %r8
	movq	16(%rdx), %r10
	movq	(%rdx), %rax
	movq	8(%rdx), %rdx
	addq	(%rsi), %rax
	adcq	8(%rsi), %rdx
	adcq	16(%rsi), %r10
	movq	%rax, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%r10, 16(%rdi)
	adcq	%rbx, %r9
	movq	%r9, 24(%rdi)
	adcq	%r11, %r8
	movq	%r8, 32(%rdi)
	sbbq	%rsi, %rsi
	andl	$1, %esi
	subq	(%rcx), %rax
	sbbq	8(%rcx), %rdx
	sbbq	16(%rcx), %r10
	sbbq	24(%rcx), %r9
	sbbq	32(%rcx), %r8
	sbbq	$0, %rsi
	testb	$1, %sil
	jne	.LBB74_2
# BB#1:                                 # %nocarry
	movq	%rax, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%r10, 16(%rdi)
	movq	%r9, 24(%rdi)
	movq	%r8, 32(%rdi)
.LBB74_2:                               # %carry
	popq	%rbx
	retq
.Lfunc_end74:
	.size	mcl_fp_add5Lbmi2, .Lfunc_end74-mcl_fp_add5Lbmi2

	.globl	mcl_fp_addNF5Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addNF5Lbmi2,@function
mcl_fp_addNF5Lbmi2:                     # @mcl_fp_addNF5Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%rbx
	movq	32(%rdx), %r8
	movq	24(%rdx), %r9
	movq	16(%rdx), %r10
	movq	(%rdx), %r14
	movq	8(%rdx), %r11
	addq	(%rsi), %r14
	adcq	8(%rsi), %r11
	adcq	16(%rsi), %r10
	adcq	24(%rsi), %r9
	adcq	32(%rsi), %r8
	movq	%r14, %rsi
	subq	(%rcx), %rsi
	movq	%r11, %rdx
	sbbq	8(%rcx), %rdx
	movq	%r10, %rbx
	sbbq	16(%rcx), %rbx
	movq	%r9, %r15
	sbbq	24(%rcx), %r15
	movq	%r8, %rax
	sbbq	32(%rcx), %rax
	movq	%rax, %rcx
	sarq	$63, %rcx
	cmovsq	%r14, %rsi
	movq	%rsi, (%rdi)
	cmovsq	%r11, %rdx
	movq	%rdx, 8(%rdi)
	cmovsq	%r10, %rbx
	movq	%rbx, 16(%rdi)
	cmovsq	%r9, %r15
	movq	%r15, 24(%rdi)
	cmovsq	%r8, %rax
	movq	%rax, 32(%rdi)
	popq	%rbx
	popq	%r14
	popq	%r15
	retq
.Lfunc_end75:
	.size	mcl_fp_addNF5Lbmi2, .Lfunc_end75-mcl_fp_addNF5Lbmi2

	.globl	mcl_fp_sub5Lbmi2
	.align	16, 0x90
	.type	mcl_fp_sub5Lbmi2,@function
mcl_fp_sub5Lbmi2:                       # @mcl_fp_sub5Lbmi2
# BB#0:
	pushq	%r14
	pushq	%rbx
	movq	32(%rsi), %r8
	movq	24(%rdx), %r11
	movq	32(%rdx), %r14
	movq	24(%rsi), %r9
	movq	16(%rsi), %r10
	movq	(%rsi), %rax
	movq	8(%rsi), %rsi
	xorl	%ebx, %ebx
	subq	(%rdx), %rax
	sbbq	8(%rdx), %rsi
	sbbq	16(%rdx), %r10
	movq	%rax, (%rdi)
	movq	%rsi, 8(%rdi)
	movq	%r10, 16(%rdi)
	sbbq	%r11, %r9
	movq	%r9, 24(%rdi)
	sbbq	%r14, %r8
	movq	%r8, 32(%rdi)
	sbbq	$0, %rbx
	testb	$1, %bl
	je	.LBB76_2
# BB#1:                                 # %carry
	movq	32(%rcx), %r11
	movq	24(%rcx), %r14
	movq	8(%rcx), %rdx
	movq	16(%rcx), %rbx
	addq	(%rcx), %rax
	movq	%rax, (%rdi)
	adcq	%rsi, %rdx
	movq	%rdx, 8(%rdi)
	adcq	%r10, %rbx
	movq	%rbx, 16(%rdi)
	adcq	%r9, %r14
	movq	%r14, 24(%rdi)
	adcq	%r8, %r11
	movq	%r11, 32(%rdi)
.LBB76_2:                               # %nocarry
	popq	%rbx
	popq	%r14
	retq
.Lfunc_end76:
	.size	mcl_fp_sub5Lbmi2, .Lfunc_end76-mcl_fp_sub5Lbmi2

	.globl	mcl_fp_subNF5Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subNF5Lbmi2,@function
mcl_fp_subNF5Lbmi2:                     # @mcl_fp_subNF5Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%rbx
	movq	32(%rsi), %r11
	movq	24(%rsi), %r8
	movq	16(%rsi), %r9
	movq	(%rsi), %r10
	movq	8(%rsi), %r14
	subq	(%rdx), %r10
	sbbq	8(%rdx), %r14
	sbbq	16(%rdx), %r9
	sbbq	24(%rdx), %r8
	sbbq	32(%rdx), %r11
	movq	%r11, %rax
	sarq	$63, %rax
	movq	%rax, %rdx
	shldq	$1, %r11, %rdx
	movq	8(%rcx), %rbx
	andq	%rdx, %rbx
	andq	(%rcx), %rdx
	movq	32(%rcx), %r15
	andq	%rax, %r15
	rorxq	$63, %rax, %rsi
	andq	24(%rcx), %rax
	andq	16(%rcx), %rsi
	addq	%r10, %rdx
	movq	%rdx, (%rdi)
	adcq	%r14, %rbx
	movq	%rbx, 8(%rdi)
	adcq	%r9, %rsi
	movq	%rsi, 16(%rdi)
	adcq	%r8, %rax
	movq	%rax, 24(%rdi)
	adcq	%r11, %r15
	movq	%r15, 32(%rdi)
	popq	%rbx
	popq	%r14
	popq	%r15
	retq
.Lfunc_end77:
	.size	mcl_fp_subNF5Lbmi2, .Lfunc_end77-mcl_fp_subNF5Lbmi2

	.globl	mcl_fpDbl_add5Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_add5Lbmi2,@function
mcl_fpDbl_add5Lbmi2:                    # @mcl_fpDbl_add5Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	72(%rdx), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	movq	64(%rdx), %r11
	movq	56(%rdx), %r14
	movq	48(%rdx), %r15
	movq	24(%rsi), %rbp
	movq	32(%rsi), %r13
	movq	16(%rdx), %r12
	movq	(%rdx), %rbx
	movq	8(%rdx), %rax
	addq	(%rsi), %rbx
	adcq	8(%rsi), %rax
	adcq	16(%rsi), %r12
	adcq	24(%rdx), %rbp
	adcq	32(%rdx), %r13
	movq	40(%rdx), %r9
	movq	%rbx, (%rdi)
	movq	72(%rsi), %r8
	movq	%rax, 8(%rdi)
	movq	64(%rsi), %r10
	movq	%r12, 16(%rdi)
	movq	56(%rsi), %r12
	movq	%rbp, 24(%rdi)
	movq	48(%rsi), %rbp
	movq	40(%rsi), %rbx
	movq	%r13, 32(%rdi)
	adcq	%r9, %rbx
	adcq	%r15, %rbp
	adcq	%r14, %r12
	adcq	%r11, %r10
	adcq	-8(%rsp), %r8           # 8-byte Folded Reload
	sbbq	%rsi, %rsi
	andl	$1, %esi
	movq	%rbx, %rax
	subq	(%rcx), %rax
	movq	%rbp, %rdx
	sbbq	8(%rcx), %rdx
	movq	%r12, %r9
	sbbq	16(%rcx), %r9
	movq	%r10, %r11
	sbbq	24(%rcx), %r11
	movq	%r8, %r14
	sbbq	32(%rcx), %r14
	sbbq	$0, %rsi
	andl	$1, %esi
	cmovneq	%rbx, %rax
	movq	%rax, 40(%rdi)
	testb	%sil, %sil
	cmovneq	%rbp, %rdx
	movq	%rdx, 48(%rdi)
	cmovneq	%r12, %r9
	movq	%r9, 56(%rdi)
	cmovneq	%r10, %r11
	movq	%r11, 64(%rdi)
	cmovneq	%r8, %r14
	movq	%r14, 72(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end78:
	.size	mcl_fpDbl_add5Lbmi2, .Lfunc_end78-mcl_fpDbl_add5Lbmi2

	.globl	mcl_fpDbl_sub5Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sub5Lbmi2,@function
mcl_fpDbl_sub5Lbmi2:                    # @mcl_fpDbl_sub5Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	72(%rdx), %r9
	movq	64(%rdx), %r10
	movq	56(%rdx), %r14
	movq	16(%rsi), %r8
	movq	(%rsi), %r15
	movq	8(%rsi), %r11
	xorl	%eax, %eax
	subq	(%rdx), %r15
	sbbq	8(%rdx), %r11
	sbbq	16(%rdx), %r8
	movq	24(%rsi), %r12
	sbbq	24(%rdx), %r12
	movq	%r15, (%rdi)
	movq	32(%rsi), %rbx
	sbbq	32(%rdx), %rbx
	movq	%r11, 8(%rdi)
	movq	48(%rdx), %r15
	movq	40(%rdx), %rdx
	movq	%r8, 16(%rdi)
	movq	72(%rsi), %r8
	movq	%r12, 24(%rdi)
	movq	64(%rsi), %r11
	movq	%rbx, 32(%rdi)
	movq	40(%rsi), %rbp
	sbbq	%rdx, %rbp
	movq	56(%rsi), %r12
	movq	48(%rsi), %r13
	sbbq	%r15, %r13
	sbbq	%r14, %r12
	sbbq	%r10, %r11
	sbbq	%r9, %r8
	movl	$0, %edx
	sbbq	$0, %rdx
	andl	$1, %edx
	movq	(%rcx), %rsi
	cmoveq	%rax, %rsi
	testb	%dl, %dl
	movq	16(%rcx), %rdx
	cmoveq	%rax, %rdx
	movq	8(%rcx), %rbx
	cmoveq	%rax, %rbx
	movq	32(%rcx), %r9
	cmoveq	%rax, %r9
	cmovneq	24(%rcx), %rax
	addq	%rbp, %rsi
	movq	%rsi, 40(%rdi)
	adcq	%r13, %rbx
	movq	%rbx, 48(%rdi)
	adcq	%r12, %rdx
	movq	%rdx, 56(%rdi)
	adcq	%r11, %rax
	movq	%rax, 64(%rdi)
	adcq	%r8, %r9
	movq	%r9, 72(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end79:
	.size	mcl_fpDbl_sub5Lbmi2, .Lfunc_end79-mcl_fpDbl_sub5Lbmi2

	.globl	mcl_fp_mulUnitPre6Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mulUnitPre6Lbmi2,@function
mcl_fp_mulUnitPre6Lbmi2:                # @mcl_fp_mulUnitPre6Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%r12
	pushq	%rbx
	mulxq	40(%rsi), %r8, %r11
	mulxq	32(%rsi), %r9, %r12
	mulxq	24(%rsi), %r10, %rcx
	mulxq	16(%rsi), %r14, %rbx
	mulxq	8(%rsi), %r15, %rax
	mulxq	(%rsi), %rdx, %rsi
	movq	%rdx, (%rdi)
	addq	%r15, %rsi
	movq	%rsi, 8(%rdi)
	adcq	%r14, %rax
	movq	%rax, 16(%rdi)
	adcq	%r10, %rbx
	movq	%rbx, 24(%rdi)
	adcq	%r9, %rcx
	movq	%rcx, 32(%rdi)
	adcq	%r8, %r12
	movq	%r12, 40(%rdi)
	adcq	$0, %r11
	movq	%r11, 48(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r14
	popq	%r15
	retq
.Lfunc_end80:
	.size	mcl_fp_mulUnitPre6Lbmi2, .Lfunc_end80-mcl_fp_mulUnitPre6Lbmi2

	.globl	mcl_fpDbl_mulPre6Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_mulPre6Lbmi2,@function
mcl_fpDbl_mulPre6Lbmi2:                 # @mcl_fpDbl_mulPre6Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rdx, %r11
	movq	%r11, -16(%rsp)         # 8-byte Spill
	movq	%rdi, -8(%rsp)          # 8-byte Spill
	movq	(%rsi), %r15
	movq	8(%rsi), %rcx
	movq	%rcx, -24(%rsp)         # 8-byte Spill
	movq	(%r11), %rax
	movq	%rcx, %rdx
	mulxq	%rax, %rcx, %r14
	movq	%r15, %rdx
	mulxq	%rax, %rdx, %rbp
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	movq	24(%rsi), %rbx
	movq	%rbx, -32(%rsp)         # 8-byte Spill
	movq	16(%rsi), %rdx
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	addq	%rcx, %rbp
	mulxq	%rax, %rcx, %r12
	adcq	%r14, %rcx
	movq	%rbx, %rdx
	mulxq	%rax, %rbx, %r14
	adcq	%r12, %rbx
	movq	32(%rsi), %r12
	movq	%r12, %rdx
	mulxq	%rax, %r8, %r13
	adcq	%r14, %r8
	movq	40(%rsi), %r14
	movq	%r14, %rdx
	mulxq	%rax, %r9, %r10
	adcq	%r13, %r9
	movq	-48(%rsp), %rax         # 8-byte Reload
	movq	%rax, (%rdi)
	adcq	$0, %r10
	movq	8(%r11), %rdi
	movq	%r15, %rdx
	mulxq	%rdi, %r13, %rax
	movq	%rax, -48(%rsp)         # 8-byte Spill
	addq	%rbp, %r13
	movq	-24(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %rbp, %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	adcq	%rcx, %rbp
	movq	-40(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %rax, %r11
	adcq	%rbx, %rax
	movq	-32(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %rbx, %rcx
	movq	%rcx, -32(%rsp)         # 8-byte Spill
	adcq	%r8, %rbx
	movq	%r12, %rdx
	mulxq	%rdi, %rcx, %r8
	adcq	%r9, %rcx
	movq	%r14, %rdx
	mulxq	%rdi, %r12, %rdx
	adcq	%r10, %r12
	sbbq	%r15, %r15
	andl	$1, %r15d
	addq	-48(%rsp), %rbp         # 8-byte Folded Reload
	adcq	-24(%rsp), %rax         # 8-byte Folded Reload
	adcq	%r11, %rbx
	movq	-8(%rsp), %rdi          # 8-byte Reload
	movq	%r13, 8(%rdi)
	adcq	-32(%rsp), %rcx         # 8-byte Folded Reload
	adcq	%r8, %r12
	adcq	%rdx, %r15
	movq	(%rsi), %rdx
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	movq	8(%rsi), %r8
	movq	%r8, -24(%rsp)          # 8-byte Spill
	movq	-16(%rsp), %r14         # 8-byte Reload
	movq	16(%r14), %rdi
	mulxq	%rdi, %r13, %rdx
	movq	%rdx, -72(%rsp)         # 8-byte Spill
	addq	%rbp, %r13
	movq	%r8, %rdx
	mulxq	%rdi, %r8, %rdx
	movq	%rdx, -80(%rsp)         # 8-byte Spill
	adcq	%rax, %r8
	movq	16(%rsi), %rdx
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	mulxq	%rdi, %r11, %rax
	movq	%rax, -88(%rsp)         # 8-byte Spill
	adcq	%rbx, %r11
	movq	24(%rsi), %rdx
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	mulxq	%rdi, %rax, %rbx
	adcq	%rcx, %rax
	movq	32(%rsi), %rdx
	movq	%rdx, -56(%rsp)         # 8-byte Spill
	mulxq	%rdi, %r10, %rcx
	adcq	%r12, %r10
	movq	40(%rsi), %rdx
	movq	%rdx, -64(%rsp)         # 8-byte Spill
	mulxq	%rdi, %r9, %rdx
	adcq	%r15, %r9
	sbbq	%rbp, %rbp
	andl	$1, %ebp
	addq	-72(%rsp), %r8          # 8-byte Folded Reload
	adcq	-80(%rsp), %r11         # 8-byte Folded Reload
	adcq	-88(%rsp), %rax         # 8-byte Folded Reload
	adcq	%rbx, %r10
	adcq	%rcx, %r9
	adcq	%rdx, %rbp
	movq	-8(%rsp), %rcx          # 8-byte Reload
	movq	%r13, 16(%rcx)
	movq	24(%r14), %rdi
	movq	-32(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %r12, %rcx
	movq	%rcx, -32(%rsp)         # 8-byte Spill
	addq	%r8, %r12
	movq	-24(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %rbx, %rcx
	movq	%rcx, -24(%rsp)         # 8-byte Spill
	adcq	%r11, %rbx
	movq	-40(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %rcx, %r11
	adcq	%rax, %rcx
	movq	-48(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %r14, %rax
	movq	%rax, -40(%rsp)         # 8-byte Spill
	adcq	%r10, %r14
	movq	-56(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %r8, %rax
	adcq	%r9, %r8
	movq	-64(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %r13, %rdx
	adcq	%rbp, %r13
	sbbq	%r15, %r15
	andl	$1, %r15d
	addq	-32(%rsp), %rbx         # 8-byte Folded Reload
	adcq	-24(%rsp), %rcx         # 8-byte Folded Reload
	adcq	%r11, %r14
	movq	-8(%rsp), %rdi          # 8-byte Reload
	movq	%r12, 24(%rdi)
	adcq	-40(%rsp), %r8          # 8-byte Folded Reload
	adcq	%rax, %r13
	adcq	%rdx, %r15
	movq	(%rsi), %rdx
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	movq	8(%rsi), %rbp
	movq	%rbp, -24(%rsp)         # 8-byte Spill
	movq	-16(%rsp), %rax         # 8-byte Reload
	movq	32(%rax), %rdi
	mulxq	%rdi, %r12, %rax
	movq	%rax, -40(%rsp)         # 8-byte Spill
	addq	%rbx, %r12
	movq	%rbp, %rdx
	mulxq	%rdi, %rbx, %rax
	movq	%rax, -48(%rsp)         # 8-byte Spill
	adcq	%rcx, %rbx
	movq	16(%rsi), %r11
	movq	%r11, %rdx
	mulxq	%rdi, %rax, %rcx
	movq	%rcx, -56(%rsp)         # 8-byte Spill
	adcq	%r14, %rax
	movq	24(%rsi), %r14
	movq	%r14, %rdx
	mulxq	%rdi, %rbp, %rcx
	movq	%rcx, -64(%rsp)         # 8-byte Spill
	adcq	%r8, %rbp
	movq	32(%rsi), %r8
	movq	%r8, %rdx
	mulxq	%rdi, %rcx, %r10
	adcq	%r13, %rcx
	movq	40(%rsi), %r13
	movq	%r13, %rdx
	mulxq	%rdi, %r9, %rdx
	adcq	%r15, %r9
	sbbq	%rsi, %rsi
	andl	$1, %esi
	addq	-40(%rsp), %rbx         # 8-byte Folded Reload
	adcq	-48(%rsp), %rax         # 8-byte Folded Reload
	adcq	-56(%rsp), %rbp         # 8-byte Folded Reload
	adcq	-64(%rsp), %rcx         # 8-byte Folded Reload
	adcq	%r10, %r9
	adcq	%rdx, %rsi
	movq	-8(%rsp), %r10          # 8-byte Reload
	movq	%r12, 32(%r10)
	movq	-16(%rsp), %rdx         # 8-byte Reload
	movq	40(%rdx), %rdi
	movq	-32(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %r15, %rdx
	movq	%rdx, -16(%rsp)         # 8-byte Spill
	addq	%rbx, %r15
	movq	-24(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %rbx, %r12
	adcq	%rax, %rbx
	movq	%r11, %rdx
	mulxq	%rdi, %rax, %r11
	adcq	%rbp, %rax
	movq	%r14, %rdx
	mulxq	%rdi, %rbp, %r14
	adcq	%rcx, %rbp
	movq	%r8, %rdx
	mulxq	%rdi, %rcx, %r8
	adcq	%r9, %rcx
	movq	%r13, %rdx
	mulxq	%rdi, %rdi, %r9
	adcq	%rsi, %rdi
	sbbq	%rsi, %rsi
	andl	$1, %esi
	addq	-16(%rsp), %rbx         # 8-byte Folded Reload
	movq	%r15, 40(%r10)
	movq	%rbx, 48(%r10)
	adcq	%r12, %rax
	movq	%rax, 56(%r10)
	adcq	%r11, %rbp
	movq	%rbp, 64(%r10)
	adcq	%r14, %rcx
	movq	%rcx, 72(%r10)
	adcq	%r8, %rdi
	movq	%rdi, 80(%r10)
	adcq	%r9, %rsi
	movq	%rsi, 88(%r10)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end81:
	.size	mcl_fpDbl_mulPre6Lbmi2, .Lfunc_end81-mcl_fpDbl_mulPre6Lbmi2

	.globl	mcl_fpDbl_sqrPre6Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sqrPre6Lbmi2,@function
mcl_fpDbl_sqrPre6Lbmi2:                 # @mcl_fpDbl_sqrPre6Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rdi, %r9
	movq	%r9, -8(%rsp)           # 8-byte Spill
	movq	16(%rsi), %rdx
	movq	%rdx, -24(%rsp)         # 8-byte Spill
	movq	(%rsi), %rcx
	movq	8(%rsi), %rax
	mulxq	%rcx, %r10, %r8
	movq	24(%rsi), %rbp
	movq	%rbp, -32(%rsp)         # 8-byte Spill
	movq	%rax, %rdx
	mulxq	%rcx, %r11, %rbx
	movq	%rbx, -16(%rsp)         # 8-byte Spill
	movq	%rcx, %rdx
	mulxq	%rcx, %rdx, %r14
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	addq	%r11, %r14
	adcq	%rbx, %r10
	movq	%rbp, %rdx
	mulxq	%rcx, %r15, %rbp
	adcq	%r8, %r15
	movq	32(%rsi), %rbx
	movq	%rbx, %rdx
	mulxq	%rcx, %r8, %r13
	adcq	%rbp, %r8
	movq	40(%rsi), %rdi
	movq	%rdi, %rdx
	mulxq	%rcx, %rcx, %r12
	adcq	%r13, %rcx
	movq	-40(%rsp), %rdx         # 8-byte Reload
	movq	%rdx, (%r9)
	adcq	$0, %r12
	addq	%r11, %r14
	movq	%rax, %rdx
	mulxq	%rax, %rbp, %rdx
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	adcq	%r10, %rbp
	movq	-24(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %r13, %r10
	adcq	%r15, %r13
	movq	-32(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %r15, %rdx
	movq	%rdx, -24(%rsp)         # 8-byte Spill
	adcq	%r8, %r15
	movq	%rbx, %rdx
	mulxq	%rax, %rbx, %r8
	adcq	%rcx, %rbx
	movq	%rdi, %rdx
	mulxq	%rax, %r11, %rax
	adcq	%r12, %r11
	sbbq	%r12, %r12
	andl	$1, %r12d
	addq	-16(%rsp), %rbp         # 8-byte Folded Reload
	adcq	-40(%rsp), %r13         # 8-byte Folded Reload
	movq	%r14, 8(%r9)
	adcq	%r10, %r15
	adcq	-24(%rsp), %rbx         # 8-byte Folded Reload
	adcq	%r8, %r11
	adcq	%rax, %r12
	movq	(%rsi), %rdx
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	movq	8(%rsi), %rdi
	movq	%rdi, -24(%rsp)         # 8-byte Spill
	movq	16(%rsi), %rcx
	mulxq	%rcx, %rax, %rdx
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	addq	%rbp, %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	movq	%rdi, %rdx
	mulxq	%rcx, %rbp, %rax
	movq	%rax, -56(%rsp)         # 8-byte Spill
	adcq	%r13, %rbp
	movq	%rcx, %rdx
	mulxq	%rcx, %r13, %rax
	movq	%rax, -64(%rsp)         # 8-byte Spill
	adcq	%r15, %r13
	movq	24(%rsi), %rax
	movq	%rax, %rdx
	mulxq	%rcx, %r8, %rdi
	movq	%rdi, -40(%rsp)         # 8-byte Spill
	adcq	%r8, %rbx
	movq	32(%rsi), %r10
	movq	%r10, %rdx
	mulxq	%rcx, %r14, %r15
	adcq	%r11, %r14
	movq	40(%rsi), %r11
	movq	%r11, %rdx
	mulxq	%rcx, %r9, %rdx
	adcq	%r12, %r9
	sbbq	%rcx, %rcx
	andl	$1, %ecx
	addq	-48(%rsp), %rbp         # 8-byte Folded Reload
	adcq	-56(%rsp), %r13         # 8-byte Folded Reload
	adcq	-64(%rsp), %rbx         # 8-byte Folded Reload
	adcq	%rdi, %r14
	adcq	%r15, %r9
	adcq	%rdx, %rcx
	movq	-32(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %rdi, %rdx
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	addq	%rbp, %rdi
	movq	-24(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %r15, %rbp
	adcq	%r13, %r15
	adcq	%r8, %rbx
	movq	%rax, %rdx
	mulxq	%rax, %r8, %rdx
	movq	%rdx, -24(%rsp)         # 8-byte Spill
	adcq	%r14, %r8
	movq	%r10, %rdx
	mulxq	%rax, %r12, %r10
	adcq	%r9, %r12
	movq	%r11, %rdx
	mulxq	%rax, %r13, %rax
	adcq	%rcx, %r13
	sbbq	%r9, %r9
	andl	$1, %r9d
	addq	-32(%rsp), %r15         # 8-byte Folded Reload
	adcq	%rbp, %rbx
	movq	-8(%rsp), %rdx          # 8-byte Reload
	movq	-16(%rsp), %rbp         # 8-byte Reload
	movq	%rbp, 16(%rdx)
	movq	%rdi, 24(%rdx)
	adcq	-40(%rsp), %r8          # 8-byte Folded Reload
	adcq	-24(%rsp), %r12         # 8-byte Folded Reload
	adcq	%r10, %r13
	adcq	%rax, %r9
	movq	(%rsi), %rcx
	movq	8(%rsi), %rdi
	movq	%rdi, -24(%rsp)         # 8-byte Spill
	movq	32(%rsi), %rax
	movq	%rcx, %rdx
	mulxq	%rax, %rdx, %rbp
	movq	%rbp, -40(%rsp)         # 8-byte Spill
	addq	%r15, %rdx
	movq	%rdx, -16(%rsp)         # 8-byte Spill
	movq	%rdi, %rdx
	mulxq	%rax, %r15, %rdx
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	adcq	%rbx, %r15
	movq	16(%rsi), %r10
	movq	%r10, %rdx
	mulxq	%rax, %r14, %rbx
	adcq	%r8, %r14
	movq	24(%rsi), %r8
	movq	%r8, %rdx
	mulxq	%rax, %rbp, %rdi
	adcq	%r12, %rbp
	movq	%rax, %rdx
	mulxq	%rax, %r11, %r12
	adcq	%r13, %r11
	movq	40(%rsi), %rsi
	movq	%rsi, %rdx
	mulxq	%rax, %r13, %rdx
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	adcq	%r13, %r9
	sbbq	%rax, %rax
	andl	$1, %eax
	addq	-40(%rsp), %r15         # 8-byte Folded Reload
	adcq	-48(%rsp), %r14         # 8-byte Folded Reload
	adcq	%rbx, %rbp
	adcq	%rdi, %r11
	adcq	%r12, %r9
	adcq	%rdx, %rax
	movq	%rcx, %rdx
	mulxq	%rsi, %r12, %rcx
	addq	%r15, %r12
	movq	-24(%rsp), %rdx         # 8-byte Reload
	mulxq	%rsi, %rdi, %r15
	adcq	%r14, %rdi
	movq	%r10, %rdx
	mulxq	%rsi, %rbx, %r10
	adcq	%rbp, %rbx
	movq	%r8, %rdx
	mulxq	%rsi, %rbp, %r8
	adcq	%r11, %rbp
	adcq	%r13, %r9
	movq	%rsi, %rdx
	mulxq	%rsi, %rsi, %r11
	adcq	%rax, %rsi
	sbbq	%rax, %rax
	andl	$1, %eax
	addq	%rcx, %rdi
	movq	-8(%rsp), %rdx          # 8-byte Reload
	movq	-16(%rsp), %rcx         # 8-byte Reload
	movq	%rcx, 32(%rdx)
	movq	%r12, 40(%rdx)
	movq	%rdi, 48(%rdx)
	adcq	%r15, %rbx
	movq	%rbx, 56(%rdx)
	adcq	%r10, %rbp
	movq	%rbp, 64(%rdx)
	adcq	%r8, %r9
	movq	%r9, 72(%rdx)
	adcq	-32(%rsp), %rsi         # 8-byte Folded Reload
	movq	%rsi, 80(%rdx)
	adcq	%r11, %rax
	movq	%rax, 88(%rdx)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end82:
	.size	mcl_fpDbl_sqrPre6Lbmi2, .Lfunc_end82-mcl_fpDbl_sqrPre6Lbmi2

	.globl	mcl_fp_mont6Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mont6Lbmi2,@function
mcl_fp_mont6Lbmi2:                      # @mcl_fp_mont6Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	subq	$32, %rsp
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	movq	%rdi, -104(%rsp)        # 8-byte Spill
	movq	40(%rsi), %rdi
	movq	%rdi, -40(%rsp)         # 8-byte Spill
	movq	(%rdx), %rax
	movq	%rdi, %rdx
	mulxq	%rax, %r11, %r14
	movq	32(%rsi), %rdx
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	mulxq	%rax, %r15, %rbx
	movq	24(%rsi), %rdx
	movq	%rdx, -80(%rsp)         # 8-byte Spill
	movq	16(%rsi), %rdi
	movq	%rdi, -72(%rsp)         # 8-byte Spill
	movq	(%rsi), %rbp
	movq	%rbp, -56(%rsp)         # 8-byte Spill
	movq	8(%rsi), %rsi
	movq	%rsi, -64(%rsp)         # 8-byte Spill
	mulxq	%rax, %r8, %r12
	movq	%rdi, %rdx
	mulxq	%rax, %r9, %r10
	movq	%rsi, %rdx
	mulxq	%rax, %rdi, %r13
	movq	%rbp, %rdx
	mulxq	%rax, %rdx, %rbp
	movq	%rdx, -96(%rsp)         # 8-byte Spill
	addq	%rdi, %rbp
	adcq	%r9, %r13
	adcq	%r8, %r10
	adcq	%r15, %r12
	adcq	%r11, %rbx
	movq	%rbx, %rdi
	adcq	$0, %r14
	movq	%r14, -88(%rsp)         # 8-byte Spill
	movq	-8(%rcx), %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	imulq	%rax, %rdx
	movq	(%rcx), %rsi
	movq	%rsi, (%rsp)            # 8-byte Spill
	movq	40(%rcx), %rax
	movq	%rax, 24(%rsp)          # 8-byte Spill
	mulxq	%rax, %rax, %r9
	movq	%rax, -112(%rsp)        # 8-byte Spill
	movq	16(%rcx), %rax
	movq	%rax, 16(%rsp)          # 8-byte Spill
	mulxq	%rax, %r8, %r11
	movq	8(%rcx), %rax
	movq	%rax, 8(%rsp)           # 8-byte Spill
	mulxq	%rax, %rax, %r14
	mulxq	%rsi, %r15, %rsi
	addq	%rax, %rsi
	adcq	%r8, %r14
	movq	24(%rcx), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	mulxq	%rax, %rbx, %r8
	adcq	%r11, %rbx
	movq	32(%rcx), %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	mulxq	%rax, %rax, %rcx
	adcq	%r8, %rax
	adcq	-112(%rsp), %rcx        # 8-byte Folded Reload
	adcq	$0, %r9
	addq	-96(%rsp), %r15         # 8-byte Folded Reload
	adcq	%rbp, %rsi
	adcq	%r13, %r14
	adcq	%r10, %rbx
	adcq	%r12, %rax
	adcq	%rdi, %rcx
	adcq	-88(%rsp), %r9          # 8-byte Folded Reload
	movq	%r9, -96(%rsp)          # 8-byte Spill
	sbbq	%r12, %r12
	andl	$1, %r12d
	movq	-32(%rsp), %rdx         # 8-byte Reload
	movq	8(%rdx), %rdx
	mulxq	-40(%rsp), %rdi, %rbp   # 8-byte Folded Reload
	movq	%rdi, -112(%rsp)        # 8-byte Spill
	movq	%rbp, -88(%rsp)         # 8-byte Spill
	mulxq	-48(%rsp), %rdi, %r13   # 8-byte Folded Reload
	movq	%rdi, -120(%rsp)        # 8-byte Spill
	mulxq	-80(%rsp), %rdi, %r15   # 8-byte Folded Reload
	movq	%rdi, -128(%rsp)        # 8-byte Spill
	mulxq	-64(%rsp), %r8, %rdi    # 8-byte Folded Reload
	mulxq	-56(%rsp), %rbp, %r10   # 8-byte Folded Reload
	addq	%r8, %r10
	mulxq	-72(%rsp), %r9, %r11    # 8-byte Folded Reload
	adcq	%rdi, %r9
	adcq	-128(%rsp), %r11        # 8-byte Folded Reload
	adcq	-120(%rsp), %r15        # 8-byte Folded Reload
	adcq	-112(%rsp), %r13        # 8-byte Folded Reload
	movq	-88(%rsp), %rdx         # 8-byte Reload
	adcq	$0, %rdx
	addq	%rsi, %rbp
	adcq	%r14, %r10
	adcq	%rbx, %r9
	adcq	%rax, %r11
	adcq	%rcx, %r15
	adcq	-96(%rsp), %r13         # 8-byte Folded Reload
	adcq	%r12, %rdx
	movq	%rdx, -88(%rsp)         # 8-byte Spill
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%rax, -96(%rsp)         # 8-byte Spill
	movq	%rbp, %rdx
	imulq	-24(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	24(%rsp), %rax, %r12    # 8-byte Folded Reload
	movq	%rax, -112(%rsp)        # 8-byte Spill
	mulxq	-16(%rsp), %rax, %r14   # 8-byte Folded Reload
	movq	%rax, -120(%rsp)        # 8-byte Spill
	mulxq	8(%rsp), %rcx, %rsi     # 8-byte Folded Reload
	mulxq	(%rsp), %rax, %r8       # 8-byte Folded Reload
	addq	%rcx, %r8
	mulxq	16(%rsp), %rdi, %rbx    # 8-byte Folded Reload
	adcq	%rsi, %rdi
	mulxq	-8(%rsp), %rcx, %rsi    # 8-byte Folded Reload
	adcq	%rbx, %rcx
	adcq	-120(%rsp), %rsi        # 8-byte Folded Reload
	adcq	-112(%rsp), %r14        # 8-byte Folded Reload
	adcq	$0, %r12
	addq	%rbp, %rax
	adcq	%r10, %r8
	adcq	%r9, %rdi
	adcq	%r11, %rcx
	adcq	%r15, %rsi
	adcq	%r13, %r14
	adcq	-88(%rsp), %r12         # 8-byte Folded Reload
	adcq	$0, -96(%rsp)           # 8-byte Folded Spill
	movq	-32(%rsp), %rdx         # 8-byte Reload
	movq	16(%rdx), %rdx
	mulxq	-40(%rsp), %rbp, %rax   # 8-byte Folded Reload
	movq	%rbp, -112(%rsp)        # 8-byte Spill
	movq	%rax, -88(%rsp)         # 8-byte Spill
	mulxq	-48(%rsp), %rax, %r13   # 8-byte Folded Reload
	movq	%rax, -120(%rsp)        # 8-byte Spill
	mulxq	-80(%rsp), %rbp, %r15   # 8-byte Folded Reload
	mulxq	-64(%rsp), %r9, %rbx    # 8-byte Folded Reload
	mulxq	-56(%rsp), %rax, %r11   # 8-byte Folded Reload
	addq	%r9, %r11
	mulxq	-72(%rsp), %r9, %r10    # 8-byte Folded Reload
	adcq	%rbx, %r9
	adcq	%rbp, %r10
	adcq	-120(%rsp), %r15        # 8-byte Folded Reload
	adcq	-112(%rsp), %r13        # 8-byte Folded Reload
	movq	-88(%rsp), %rdx         # 8-byte Reload
	adcq	$0, %rdx
	addq	%r8, %rax
	movq	%rax, -112(%rsp)        # 8-byte Spill
	movq	%rax, %rbp
	adcq	%rdi, %r11
	adcq	%rcx, %r9
	adcq	%rsi, %r10
	adcq	%r14, %r15
	adcq	%r12, %r13
	adcq	-96(%rsp), %rdx         # 8-byte Folded Reload
	movq	%rdx, -88(%rsp)         # 8-byte Spill
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%rax, -96(%rsp)         # 8-byte Spill
	movq	%rbp, %rdx
	imulq	-24(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	24(%rsp), %rax, %r8     # 8-byte Folded Reload
	movq	%rax, -120(%rsp)        # 8-byte Spill
	mulxq	-16(%rsp), %r12, %r14   # 8-byte Folded Reload
	mulxq	8(%rsp), %rcx, %rsi     # 8-byte Folded Reload
	mulxq	(%rsp), %rax, %rbx      # 8-byte Folded Reload
	addq	%rcx, %rbx
	mulxq	16(%rsp), %rbp, %rdi    # 8-byte Folded Reload
	adcq	%rsi, %rbp
	mulxq	-8(%rsp), %rcx, %rsi    # 8-byte Folded Reload
	adcq	%rdi, %rcx
	adcq	%r12, %rsi
	adcq	-120(%rsp), %r14        # 8-byte Folded Reload
	adcq	$0, %r8
	addq	-112(%rsp), %rax        # 8-byte Folded Reload
	adcq	%r11, %rbx
	adcq	%r9, %rbp
	adcq	%r10, %rcx
	adcq	%r15, %rsi
	adcq	%r13, %r14
	adcq	-88(%rsp), %r8          # 8-byte Folded Reload
	movq	%r8, -112(%rsp)         # 8-byte Spill
	movq	-96(%rsp), %r13         # 8-byte Reload
	adcq	$0, %r13
	movq	-32(%rsp), %rdx         # 8-byte Reload
	movq	24(%rdx), %rdx
	mulxq	-40(%rsp), %rdi, %rax   # 8-byte Folded Reload
	movq	%rdi, -96(%rsp)         # 8-byte Spill
	movq	%rax, -88(%rsp)         # 8-byte Spill
	mulxq	-48(%rsp), %rdi, %rax   # 8-byte Folded Reload
	movq	%rdi, -120(%rsp)        # 8-byte Spill
	mulxq	-80(%rsp), %r15, %r12   # 8-byte Folded Reload
	mulxq	-64(%rsp), %r8, %r11    # 8-byte Folded Reload
	mulxq	-56(%rsp), %r10, %rdi   # 8-byte Folded Reload
	addq	%r8, %rdi
	mulxq	-72(%rsp), %r8, %r9     # 8-byte Folded Reload
	adcq	%r11, %r8
	adcq	%r15, %r9
	adcq	-120(%rsp), %r12        # 8-byte Folded Reload
	adcq	-96(%rsp), %rax         # 8-byte Folded Reload
	movq	-88(%rsp), %rdx         # 8-byte Reload
	adcq	$0, %rdx
	addq	%rbx, %r10
	adcq	%rbp, %rdi
	adcq	%rcx, %r8
	adcq	%rsi, %r9
	adcq	%r14, %r12
	adcq	-112(%rsp), %rax        # 8-byte Folded Reload
	movq	%rax, -96(%rsp)         # 8-byte Spill
	adcq	%r13, %rdx
	movq	%rdx, -88(%rsp)         # 8-byte Spill
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%rax, %r14
	movq	%r10, %rdx
	imulq	-24(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	24(%rsp), %rax, %r13    # 8-byte Folded Reload
	movq	%rax, -112(%rsp)        # 8-byte Spill
	mulxq	-16(%rsp), %rax, %r11   # 8-byte Folded Reload
	movq	%rax, -120(%rsp)        # 8-byte Spill
	mulxq	8(%rsp), %rbp, %rsi     # 8-byte Folded Reload
	mulxq	(%rsp), %rcx, %rbx      # 8-byte Folded Reload
	addq	%rbp, %rbx
	mulxq	16(%rsp), %rbp, %rax    # 8-byte Folded Reload
	adcq	%rsi, %rbp
	mulxq	-8(%rsp), %rsi, %r15    # 8-byte Folded Reload
	adcq	%rax, %rsi
	adcq	-120(%rsp), %r15        # 8-byte Folded Reload
	adcq	-112(%rsp), %r11        # 8-byte Folded Reload
	adcq	$0, %r13
	addq	%r10, %rcx
	adcq	%rdi, %rbx
	adcq	%r8, %rbp
	adcq	%r9, %rsi
	adcq	%r12, %r15
	adcq	-96(%rsp), %r11         # 8-byte Folded Reload
	adcq	-88(%rsp), %r13         # 8-byte Folded Reload
	movq	%r14, %rdi
	adcq	$0, %rdi
	movq	-32(%rsp), %rcx         # 8-byte Reload
	movq	32(%rcx), %rdx
	mulxq	-40(%rsp), %rcx, %rax   # 8-byte Folded Reload
	movq	%rcx, -96(%rsp)         # 8-byte Spill
	movq	%rax, -88(%rsp)         # 8-byte Spill
	mulxq	-48(%rsp), %rax, %r12   # 8-byte Folded Reload
	movq	%rax, -112(%rsp)        # 8-byte Spill
	mulxq	-80(%rsp), %rax, %r14   # 8-byte Folded Reload
	movq	%rax, -120(%rsp)        # 8-byte Spill
	mulxq	-64(%rsp), %rcx, %r9    # 8-byte Folded Reload
	mulxq	-56(%rsp), %rax, %r8    # 8-byte Folded Reload
	addq	%rcx, %r8
	mulxq	-72(%rsp), %rcx, %r10   # 8-byte Folded Reload
	adcq	%r9, %rcx
	adcq	-120(%rsp), %r10        # 8-byte Folded Reload
	adcq	-112(%rsp), %r14        # 8-byte Folded Reload
	adcq	-96(%rsp), %r12         # 8-byte Folded Reload
	movq	-88(%rsp), %rdx         # 8-byte Reload
	adcq	$0, %rdx
	addq	%rbx, %rax
	movq	%rax, -112(%rsp)        # 8-byte Spill
	movq	%rax, %rbx
	adcq	%rbp, %r8
	adcq	%rsi, %rcx
	adcq	%r15, %r10
	adcq	%r11, %r14
	adcq	%r13, %r12
	adcq	%rdi, %rdx
	movq	%rdx, -88(%rsp)         # 8-byte Spill
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%rax, -96(%rsp)         # 8-byte Spill
	movq	%rbx, %rdx
	imulq	-24(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	24(%rsp), %rax, %r15    # 8-byte Folded Reload
	movq	%rax, -120(%rsp)        # 8-byte Spill
	mulxq	-16(%rsp), %r13, %r11   # 8-byte Folded Reload
	mulxq	8(%rsp), %rsi, %rax     # 8-byte Folded Reload
	mulxq	(%rsp), %rdi, %rbx      # 8-byte Folded Reload
	addq	%rsi, %rbx
	mulxq	16(%rsp), %rbp, %r9     # 8-byte Folded Reload
	adcq	%rax, %rbp
	mulxq	-8(%rsp), %rax, %rsi    # 8-byte Folded Reload
	adcq	%r9, %rax
	adcq	%r13, %rsi
	adcq	-120(%rsp), %r11        # 8-byte Folded Reload
	adcq	$0, %r15
	addq	-112(%rsp), %rdi        # 8-byte Folded Reload
	adcq	%r8, %rbx
	adcq	%rcx, %rbp
	adcq	%r10, %rax
	adcq	%r14, %rsi
	adcq	%r12, %r11
	adcq	-88(%rsp), %r15         # 8-byte Folded Reload
	movq	-96(%rsp), %r8          # 8-byte Reload
	adcq	$0, %r8
	movq	-32(%rsp), %rcx         # 8-byte Reload
	movq	40(%rcx), %rdx
	mulxq	-40(%rsp), %rdi, %rcx   # 8-byte Folded Reload
	movq	%rdi, -88(%rsp)         # 8-byte Spill
	movq	%rcx, -32(%rsp)         # 8-byte Spill
	mulxq	-48(%rsp), %rdi, %rcx   # 8-byte Folded Reload
	movq	%rdi, -48(%rsp)         # 8-byte Spill
	movq	%rcx, -40(%rsp)         # 8-byte Spill
	mulxq	-80(%rsp), %rcx, %r14   # 8-byte Folded Reload
	movq	%rcx, -80(%rsp)         # 8-byte Spill
	mulxq	-72(%rsp), %rdi, %r12   # 8-byte Folded Reload
	mulxq	-64(%rsp), %rcx, %r10   # 8-byte Folded Reload
	mulxq	-56(%rsp), %r13, %r9    # 8-byte Folded Reload
	addq	%rcx, %r9
	adcq	%rdi, %r10
	adcq	-80(%rsp), %r12         # 8-byte Folded Reload
	adcq	-48(%rsp), %r14         # 8-byte Folded Reload
	movq	-40(%rsp), %rdx         # 8-byte Reload
	adcq	-88(%rsp), %rdx         # 8-byte Folded Reload
	movq	-32(%rsp), %rcx         # 8-byte Reload
	adcq	$0, %rcx
	addq	%rbx, %r13
	adcq	%rbp, %r9
	adcq	%rax, %r10
	adcq	%rsi, %r12
	adcq	%r11, %r14
	adcq	%r15, %rdx
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	adcq	%r8, %rcx
	movq	%rcx, -32(%rsp)         # 8-byte Spill
	sbbq	%rcx, %rcx
	movq	-24(%rsp), %rdx         # 8-byte Reload
	imulq	%r13, %rdx
	mulxq	(%rsp), %r11, %rax      # 8-byte Folded Reload
	mulxq	8(%rsp), %rdi, %rbx     # 8-byte Folded Reload
	addq	%rax, %rdi
	mulxq	16(%rsp), %rsi, %rax    # 8-byte Folded Reload
	adcq	%rbx, %rsi
	mulxq	-8(%rsp), %rbx, %rbp    # 8-byte Folded Reload
	adcq	%rax, %rbx
	mulxq	-16(%rsp), %rax, %r15   # 8-byte Folded Reload
	adcq	%rbp, %rax
	mulxq	24(%rsp), %rbp, %rdx    # 8-byte Folded Reload
	adcq	%r15, %rbp
	adcq	$0, %rdx
	andl	$1, %ecx
	addq	%r13, %r11
	adcq	%r9, %rdi
	adcq	%r10, %rsi
	adcq	%r12, %rbx
	adcq	%r14, %rax
	adcq	-40(%rsp), %rbp         # 8-byte Folded Reload
	adcq	-32(%rsp), %rdx         # 8-byte Folded Reload
	adcq	$0, %rcx
	movq	%rdi, %r8
	subq	(%rsp), %r8             # 8-byte Folded Reload
	movq	%rsi, %r9
	sbbq	8(%rsp), %r9            # 8-byte Folded Reload
	movq	%rbx, %r10
	sbbq	16(%rsp), %r10          # 8-byte Folded Reload
	movq	%rax, %r11
	sbbq	-8(%rsp), %r11          # 8-byte Folded Reload
	movq	%rbp, %r14
	sbbq	-16(%rsp), %r14         # 8-byte Folded Reload
	movq	%rdx, %r15
	sbbq	24(%rsp), %r15          # 8-byte Folded Reload
	sbbq	$0, %rcx
	andl	$1, %ecx
	cmovneq	%rax, %r11
	testb	%cl, %cl
	cmovneq	%rdi, %r8
	movq	-104(%rsp), %rax        # 8-byte Reload
	movq	%r8, (%rax)
	cmovneq	%rsi, %r9
	movq	%r9, 8(%rax)
	cmovneq	%rbx, %r10
	movq	%r10, 16(%rax)
	movq	%r11, 24(%rax)
	cmovneq	%rbp, %r14
	movq	%r14, 32(%rax)
	cmovneq	%rdx, %r15
	movq	%r15, 40(%rax)
	addq	$32, %rsp
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end83:
	.size	mcl_fp_mont6Lbmi2, .Lfunc_end83-mcl_fp_mont6Lbmi2

	.globl	mcl_fp_montNF6Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montNF6Lbmi2,@function
mcl_fp_montNF6Lbmi2:                    # @mcl_fp_montNF6Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rdx, -72(%rsp)         # 8-byte Spill
	movq	%rdi, -120(%rsp)        # 8-byte Spill
	movq	(%rsi), %rax
	movq	%rax, -64(%rsp)         # 8-byte Spill
	movq	8(%rsi), %rdi
	movq	%rdi, -80(%rsp)         # 8-byte Spill
	movq	(%rdx), %rbp
	movq	%rdi, %rdx
	mulxq	%rbp, %rdi, %rbx
	movq	%rax, %rdx
	mulxq	%rbp, %r9, %r14
	movq	16(%rsi), %rdx
	movq	%rdx, -88(%rsp)         # 8-byte Spill
	addq	%rdi, %r14
	mulxq	%rbp, %rdi, %r8
	adcq	%rbx, %rdi
	movq	24(%rsi), %rdx
	movq	%rdx, -96(%rsp)         # 8-byte Spill
	mulxq	%rbp, %rbx, %r10
	adcq	%r8, %rbx
	movq	32(%rsi), %rdx
	movq	%rdx, -104(%rsp)        # 8-byte Spill
	mulxq	%rbp, %r8, %r11
	adcq	%r10, %r8
	movq	40(%rsi), %rdx
	movq	%rdx, -112(%rsp)        # 8-byte Spill
	mulxq	%rbp, %rsi, %r15
	adcq	%r11, %rsi
	adcq	$0, %r15
	movq	-8(%rcx), %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	movq	%r9, %rdx
	imulq	%rax, %rdx
	movq	(%rcx), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	mulxq	%rax, %rbp, %rax
	movq	%rax, -128(%rsp)        # 8-byte Spill
	addq	%r9, %rbp
	movq	8(%rcx), %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	mulxq	%rax, %r12, %r9
	adcq	%r14, %r12
	movq	16(%rcx), %rax
	movq	%rax, -32(%rsp)         # 8-byte Spill
	mulxq	%rax, %r14, %rax
	adcq	%rdi, %r14
	movq	24(%rcx), %rdi
	movq	%rdi, -40(%rsp)         # 8-byte Spill
	mulxq	%rdi, %r13, %rdi
	adcq	%rbx, %r13
	movq	32(%rcx), %rbp
	movq	%rbp, -48(%rsp)         # 8-byte Spill
	mulxq	%rbp, %r11, %rbx
	adcq	%r8, %r11
	movq	40(%rcx), %rcx
	movq	%rcx, -56(%rsp)         # 8-byte Spill
	mulxq	%rcx, %r10, %rcx
	adcq	%rsi, %r10
	adcq	$0, %r15
	addq	-128(%rsp), %r12        # 8-byte Folded Reload
	adcq	%r9, %r14
	adcq	%rax, %r13
	adcq	%rdi, %r11
	adcq	%rbx, %r10
	adcq	%rcx, %r15
	movq	-72(%rsp), %rax         # 8-byte Reload
	movq	8(%rax), %rdx
	mulxq	-80(%rsp), %rcx, %rsi   # 8-byte Folded Reload
	mulxq	-64(%rsp), %rbx, %rax   # 8-byte Folded Reload
	addq	%rcx, %rax
	mulxq	-88(%rsp), %rcx, %rdi   # 8-byte Folded Reload
	adcq	%rsi, %rcx
	mulxq	-96(%rsp), %rsi, %r8    # 8-byte Folded Reload
	adcq	%rdi, %rsi
	mulxq	-104(%rsp), %rdi, %rbp  # 8-byte Folded Reload
	movq	%rbp, -128(%rsp)        # 8-byte Spill
	adcq	%r8, %rdi
	mulxq	-112(%rsp), %r8, %r9    # 8-byte Folded Reload
	adcq	-128(%rsp), %r8         # 8-byte Folded Reload
	adcq	$0, %r9
	addq	%r12, %rbx
	adcq	%r14, %rax
	adcq	%r13, %rcx
	adcq	%r11, %rsi
	adcq	%r10, %rdi
	adcq	%r15, %r8
	adcq	$0, %r9
	movq	%rbx, %rdx
	imulq	-16(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	-8(%rsp), %rbp, %r13    # 8-byte Folded Reload
	addq	%rbx, %rbp
	mulxq	-24(%rsp), %r11, %rbx   # 8-byte Folded Reload
	adcq	%rax, %r11
	mulxq	-32(%rsp), %r14, %rax   # 8-byte Folded Reload
	adcq	%rcx, %r14
	mulxq	-40(%rsp), %r10, %rcx   # 8-byte Folded Reload
	adcq	%rsi, %r10
	mulxq	-48(%rsp), %r15, %rsi   # 8-byte Folded Reload
	adcq	%rdi, %r15
	mulxq	-56(%rsp), %r12, %rdx   # 8-byte Folded Reload
	adcq	%r8, %r12
	adcq	$0, %r9
	addq	%r13, %r11
	adcq	%rbx, %r14
	adcq	%rax, %r10
	adcq	%rcx, %r15
	adcq	%rsi, %r12
	adcq	%rdx, %r9
	movq	-72(%rsp), %rax         # 8-byte Reload
	movq	16(%rax), %rdx
	mulxq	-80(%rsp), %rcx, %rax   # 8-byte Folded Reload
	mulxq	-64(%rsp), %r13, %rdi   # 8-byte Folded Reload
	addq	%rcx, %rdi
	mulxq	-88(%rsp), %rbx, %rcx   # 8-byte Folded Reload
	adcq	%rax, %rbx
	mulxq	-96(%rsp), %rsi, %rbp   # 8-byte Folded Reload
	adcq	%rcx, %rsi
	mulxq	-104(%rsp), %rax, %rcx  # 8-byte Folded Reload
	movq	%rcx, -128(%rsp)        # 8-byte Spill
	adcq	%rbp, %rax
	mulxq	-112(%rsp), %r8, %rcx   # 8-byte Folded Reload
	adcq	-128(%rsp), %r8         # 8-byte Folded Reload
	adcq	$0, %rcx
	addq	%r11, %r13
	adcq	%r14, %rdi
	adcq	%r10, %rbx
	adcq	%r15, %rsi
	adcq	%r12, %rax
	adcq	%r9, %r8
	adcq	$0, %rcx
	movq	%r13, %rdx
	imulq	-16(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	-8(%rsp), %rbp, %r12    # 8-byte Folded Reload
	addq	%r13, %rbp
	mulxq	-24(%rsp), %r11, %rbp   # 8-byte Folded Reload
	adcq	%rdi, %r11
	mulxq	-32(%rsp), %r9, %rdi    # 8-byte Folded Reload
	adcq	%rbx, %r9
	mulxq	-40(%rsp), %r10, %rbx   # 8-byte Folded Reload
	adcq	%rsi, %r10
	mulxq	-48(%rsp), %r14, %rsi   # 8-byte Folded Reload
	adcq	%rax, %r14
	mulxq	-56(%rsp), %r15, %rax   # 8-byte Folded Reload
	adcq	%r8, %r15
	adcq	$0, %rcx
	addq	%r12, %r11
	adcq	%rbp, %r9
	adcq	%rdi, %r10
	adcq	%rbx, %r14
	adcq	%rsi, %r15
	adcq	%rax, %rcx
	movq	-72(%rsp), %rax         # 8-byte Reload
	movq	24(%rax), %rdx
	mulxq	-80(%rsp), %rsi, %rax   # 8-byte Folded Reload
	mulxq	-64(%rsp), %r13, %rbx   # 8-byte Folded Reload
	addq	%rsi, %rbx
	mulxq	-88(%rsp), %rdi, %rbp   # 8-byte Folded Reload
	adcq	%rax, %rdi
	mulxq	-96(%rsp), %rsi, %r8    # 8-byte Folded Reload
	adcq	%rbp, %rsi
	mulxq	-104(%rsp), %rax, %rbp  # 8-byte Folded Reload
	adcq	%r8, %rax
	mulxq	-112(%rsp), %r8, %r12   # 8-byte Folded Reload
	adcq	%rbp, %r8
	adcq	$0, %r12
	addq	%r11, %r13
	adcq	%r9, %rbx
	adcq	%r10, %rdi
	adcq	%r14, %rsi
	adcq	%r15, %rax
	adcq	%rcx, %r8
	adcq	$0, %r12
	movq	%r13, %rdx
	imulq	-16(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	-8(%rsp), %rbp, %rcx    # 8-byte Folded Reload
	addq	%r13, %rbp
	mulxq	-24(%rsp), %r11, %rbp   # 8-byte Folded Reload
	adcq	%rbx, %r11
	mulxq	-32(%rsp), %r9, %rbx    # 8-byte Folded Reload
	adcq	%rdi, %r9
	mulxq	-40(%rsp), %r10, %rdi   # 8-byte Folded Reload
	adcq	%rsi, %r10
	mulxq	-48(%rsp), %r14, %rsi   # 8-byte Folded Reload
	adcq	%rax, %r14
	mulxq	-56(%rsp), %r15, %rax   # 8-byte Folded Reload
	adcq	%r8, %r15
	adcq	$0, %r12
	addq	%rcx, %r11
	adcq	%rbp, %r9
	adcq	%rbx, %r10
	adcq	%rdi, %r14
	adcq	%rsi, %r15
	adcq	%rax, %r12
	movq	-72(%rsp), %rax         # 8-byte Reload
	movq	32(%rax), %rdx
	mulxq	-80(%rsp), %rsi, %rcx   # 8-byte Folded Reload
	mulxq	-64(%rsp), %r13, %rax   # 8-byte Folded Reload
	addq	%rsi, %rax
	mulxq	-88(%rsp), %rbx, %rsi   # 8-byte Folded Reload
	adcq	%rcx, %rbx
	mulxq	-96(%rsp), %rdi, %rcx   # 8-byte Folded Reload
	adcq	%rsi, %rdi
	mulxq	-104(%rsp), %rsi, %rbp  # 8-byte Folded Reload
	adcq	%rcx, %rsi
	mulxq	-112(%rsp), %r8, %rcx   # 8-byte Folded Reload
	adcq	%rbp, %r8
	adcq	$0, %rcx
	addq	%r11, %r13
	adcq	%r9, %rax
	adcq	%r10, %rbx
	adcq	%r14, %rdi
	adcq	%r15, %rsi
	adcq	%r12, %r8
	adcq	$0, %rcx
	movq	%r13, %rdx
	imulq	-16(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	-8(%rsp), %rbp, %r9     # 8-byte Folded Reload
	addq	%r13, %rbp
	mulxq	-24(%rsp), %r13, %rbp   # 8-byte Folded Reload
	adcq	%rax, %r13
	mulxq	-32(%rsp), %r11, %rax   # 8-byte Folded Reload
	adcq	%rbx, %r11
	mulxq	-40(%rsp), %r10, %rbx   # 8-byte Folded Reload
	adcq	%rdi, %r10
	mulxq	-48(%rsp), %r14, %rdi   # 8-byte Folded Reload
	adcq	%rsi, %r14
	mulxq	-56(%rsp), %rsi, %rdx   # 8-byte Folded Reload
	adcq	%r8, %rsi
	adcq	$0, %rcx
	addq	%r9, %r13
	adcq	%rbp, %r11
	adcq	%rax, %r10
	adcq	%rbx, %r14
	adcq	%rdi, %rsi
	adcq	%rdx, %rcx
	movq	-72(%rsp), %rax         # 8-byte Reload
	movq	40(%rax), %rdx
	mulxq	-80(%rsp), %rdi, %rax   # 8-byte Folded Reload
	mulxq	-64(%rsp), %r8, %rbx    # 8-byte Folded Reload
	addq	%rdi, %rbx
	mulxq	-88(%rsp), %rdi, %rbp   # 8-byte Folded Reload
	adcq	%rax, %rdi
	mulxq	-96(%rsp), %r15, %rax   # 8-byte Folded Reload
	adcq	%rbp, %r15
	mulxq	-104(%rsp), %r12, %rbp  # 8-byte Folded Reload
	adcq	%rax, %r12
	mulxq	-112(%rsp), %r9, %rax   # 8-byte Folded Reload
	adcq	%rbp, %r9
	adcq	$0, %rax
	addq	%r13, %r8
	adcq	%r11, %rbx
	adcq	%r10, %rdi
	adcq	%r14, %r15
	adcq	%rsi, %r12
	adcq	%rcx, %r9
	adcq	$0, %rax
	movq	-16(%rsp), %rdx         # 8-byte Reload
	imulq	%r8, %rdx
	mulxq	-8(%rsp), %rcx, %rsi    # 8-byte Folded Reload
	movq	%rsi, -16(%rsp)         # 8-byte Spill
	addq	%r8, %rcx
	movq	-24(%rsp), %r11         # 8-byte Reload
	mulxq	%r11, %r8, %rcx
	movq	%rcx, -64(%rsp)         # 8-byte Spill
	adcq	%rbx, %r8
	movq	-32(%rsp), %r10         # 8-byte Reload
	mulxq	%r10, %rsi, %rcx
	movq	%rcx, -72(%rsp)         # 8-byte Spill
	adcq	%rdi, %rsi
	movq	-40(%rsp), %r13         # 8-byte Reload
	mulxq	%r13, %rdi, %rcx
	movq	%rcx, -80(%rsp)         # 8-byte Spill
	adcq	%r15, %rdi
	movq	-48(%rsp), %rcx         # 8-byte Reload
	mulxq	%rcx, %r15, %rbx
	adcq	%r12, %r15
	movq	-56(%rsp), %r14         # 8-byte Reload
	mulxq	%r14, %r12, %rbp
	adcq	%r9, %r12
	adcq	$0, %rax
	addq	-16(%rsp), %r8          # 8-byte Folded Reload
	adcq	-64(%rsp), %rsi         # 8-byte Folded Reload
	adcq	-72(%rsp), %rdi         # 8-byte Folded Reload
	adcq	-80(%rsp), %r15         # 8-byte Folded Reload
	adcq	%rbx, %r12
	adcq	%rbp, %rax
	movq	%r8, %rbp
	subq	-8(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rsi, %rbx
	sbbq	%r11, %rbx
	movq	%rdi, %r11
	sbbq	%r10, %r11
	movq	%r15, %r10
	sbbq	%r13, %r10
	movq	%r12, %r9
	sbbq	%rcx, %r9
	movq	%rax, %rcx
	sbbq	%r14, %rcx
	movq	%rcx, %rdx
	sarq	$63, %rdx
	cmovsq	%r8, %rbp
	movq	-120(%rsp), %rdx        # 8-byte Reload
	movq	%rbp, (%rdx)
	cmovsq	%rsi, %rbx
	movq	%rbx, 8(%rdx)
	cmovsq	%rdi, %r11
	movq	%r11, 16(%rdx)
	cmovsq	%r15, %r10
	movq	%r10, 24(%rdx)
	cmovsq	%r12, %r9
	movq	%r9, 32(%rdx)
	cmovsq	%rax, %rcx
	movq	%rcx, 40(%rdx)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end84:
	.size	mcl_fp_montNF6Lbmi2, .Lfunc_end84-mcl_fp_montNF6Lbmi2

	.globl	mcl_fp_montRed6Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montRed6Lbmi2,@function
mcl_fp_montRed6Lbmi2:                   # @mcl_fp_montRed6Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	pushq	%rax
	movq	%rdx, %rcx
	movq	%rdi, -104(%rsp)        # 8-byte Spill
	movq	-8(%rcx), %rax
	movq	%rax, -56(%rsp)         # 8-byte Spill
	movq	(%rcx), %rdi
	movq	%rdi, -8(%rsp)          # 8-byte Spill
	movq	(%rsi), %r14
	movq	%r14, %rdx
	imulq	%rax, %rdx
	movq	40(%rcx), %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	mulxq	%rax, %rbx, %r12
	movq	32(%rcx), %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	mulxq	%rax, %r10, %r11
	movq	24(%rcx), %rax
	movq	%rax, -32(%rsp)         # 8-byte Spill
	movq	16(%rcx), %rbp
	movq	%rbp, -40(%rsp)         # 8-byte Spill
	movq	8(%rcx), %rcx
	movq	%rcx, (%rsp)            # 8-byte Spill
	mulxq	%rax, %r15, %r8
	mulxq	%rbp, %r13, %rbp
	mulxq	%rcx, %rax, %r9
	mulxq	%rdi, %rdx, %rcx
	addq	%rax, %rcx
	adcq	%r13, %r9
	adcq	%r15, %rbp
	adcq	%r10, %r8
	adcq	%rbx, %r11
	adcq	$0, %r12
	addq	%r14, %rdx
	adcq	8(%rsi), %rcx
	adcq	16(%rsi), %r9
	adcq	24(%rsi), %rbp
	adcq	32(%rsi), %r8
	adcq	40(%rsi), %r11
	movq	%r11, -88(%rsp)         # 8-byte Spill
	adcq	48(%rsi), %r12
	movq	%r12, -80(%rsp)         # 8-byte Spill
	movq	88(%rsi), %r10
	movq	80(%rsi), %rdx
	movq	72(%rsi), %rdi
	movq	64(%rsi), %rax
	movq	56(%rsi), %rsi
	adcq	$0, %rsi
	movq	%rsi, -112(%rsp)        # 8-byte Spill
	adcq	$0, %rax
	movq	%rax, -96(%rsp)         # 8-byte Spill
	adcq	$0, %rdi
	movq	%rdi, -72(%rsp)         # 8-byte Spill
	adcq	$0, %rdx
	movq	%rdx, -64(%rsp)         # 8-byte Spill
	adcq	$0, %r10
	movq	%r10, -48(%rsp)         # 8-byte Spill
	sbbq	%r12, %r12
	andl	$1, %r12d
	movq	%rcx, %rdx
	imulq	-56(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	-24(%rsp), %rax, %r13   # 8-byte Folded Reload
	movq	%rax, -120(%rsp)        # 8-byte Spill
	mulxq	-16(%rsp), %rax, %r15   # 8-byte Folded Reload
	movq	%rax, -128(%rsp)        # 8-byte Spill
	mulxq	-32(%rsp), %r11, %r14   # 8-byte Folded Reload
	mulxq	-40(%rsp), %rbx, %r10   # 8-byte Folded Reload
	mulxq	(%rsp), %rsi, %rdi      # 8-byte Folded Reload
	mulxq	-8(%rsp), %rdx, %rax    # 8-byte Folded Reload
	addq	%rsi, %rax
	adcq	%rbx, %rdi
	adcq	%r11, %r10
	adcq	-128(%rsp), %r14        # 8-byte Folded Reload
	adcq	-120(%rsp), %r15        # 8-byte Folded Reload
	adcq	$0, %r13
	addq	%rcx, %rdx
	adcq	%r9, %rax
	adcq	%rbp, %rdi
	adcq	%r8, %r10
	adcq	-88(%rsp), %r14         # 8-byte Folded Reload
	adcq	-80(%rsp), %r15         # 8-byte Folded Reload
	adcq	-112(%rsp), %r13        # 8-byte Folded Reload
	adcq	$0, -96(%rsp)           # 8-byte Folded Spill
	adcq	$0, -72(%rsp)           # 8-byte Folded Spill
	adcq	$0, -64(%rsp)           # 8-byte Folded Spill
	adcq	$0, -48(%rsp)           # 8-byte Folded Spill
	adcq	$0, %r12
	movq	%rax, %rdx
	movq	-56(%rsp), %r11         # 8-byte Reload
	imulq	%r11, %rdx
	mulxq	-24(%rsp), %rsi, %rcx   # 8-byte Folded Reload
	movq	%rsi, -112(%rsp)        # 8-byte Spill
	movq	%rcx, -80(%rsp)         # 8-byte Spill
	mulxq	-16(%rsp), %rsi, %rcx   # 8-byte Folded Reload
	movq	%rsi, -120(%rsp)        # 8-byte Spill
	movq	%rcx, -88(%rsp)         # 8-byte Spill
	mulxq	-32(%rsp), %rcx, %rbx   # 8-byte Folded Reload
	movq	%rcx, -128(%rsp)        # 8-byte Spill
	mulxq	-40(%rsp), %rcx, %r9    # 8-byte Folded Reload
	mulxq	(%rsp), %rsi, %rbp      # 8-byte Folded Reload
	mulxq	-8(%rsp), %rdx, %r8     # 8-byte Folded Reload
	addq	%rsi, %r8
	adcq	%rcx, %rbp
	adcq	-128(%rsp), %r9         # 8-byte Folded Reload
	adcq	-120(%rsp), %rbx        # 8-byte Folded Reload
	movq	-88(%rsp), %rsi         # 8-byte Reload
	adcq	-112(%rsp), %rsi        # 8-byte Folded Reload
	movq	-80(%rsp), %rcx         # 8-byte Reload
	adcq	$0, %rcx
	addq	%rax, %rdx
	adcq	%rdi, %r8
	adcq	%r10, %rbp
	adcq	%r14, %r9
	adcq	%r15, %rbx
	adcq	%r13, %rsi
	movq	%rsi, -88(%rsp)         # 8-byte Spill
	adcq	-96(%rsp), %rcx         # 8-byte Folded Reload
	movq	%rcx, -80(%rsp)         # 8-byte Spill
	adcq	$0, -72(%rsp)           # 8-byte Folded Spill
	movq	-64(%rsp), %rax         # 8-byte Reload
	adcq	$0, %rax
	adcq	$0, -48(%rsp)           # 8-byte Folded Spill
	adcq	$0, %r12
	movq	%r8, %rdx
	imulq	%r11, %rdx
	mulxq	-24(%rsp), %rsi, %rcx   # 8-byte Folded Reload
	movq	%rsi, -64(%rsp)         # 8-byte Spill
	movq	%rcx, -96(%rsp)         # 8-byte Spill
	mulxq	-16(%rsp), %rcx, %r11   # 8-byte Folded Reload
	movq	%rcx, -112(%rsp)        # 8-byte Spill
	mulxq	-32(%rsp), %r10, %r14   # 8-byte Folded Reload
	mulxq	-40(%rsp), %r13, %r15   # 8-byte Folded Reload
	mulxq	(%rsp), %rsi, %rdi      # 8-byte Folded Reload
	mulxq	-8(%rsp), %rdx, %rcx    # 8-byte Folded Reload
	addq	%rsi, %rcx
	adcq	%r13, %rdi
	adcq	%r10, %r15
	adcq	-112(%rsp), %r14        # 8-byte Folded Reload
	adcq	-64(%rsp), %r11         # 8-byte Folded Reload
	movq	-96(%rsp), %rsi         # 8-byte Reload
	adcq	$0, %rsi
	addq	%r8, %rdx
	adcq	%rbp, %rcx
	adcq	%r9, %rdi
	adcq	%rbx, %r15
	adcq	-88(%rsp), %r14         # 8-byte Folded Reload
	adcq	-80(%rsp), %r11         # 8-byte Folded Reload
	adcq	-72(%rsp), %rsi         # 8-byte Folded Reload
	movq	%rsi, -96(%rsp)         # 8-byte Spill
	adcq	$0, %rax
	movq	%rax, -64(%rsp)         # 8-byte Spill
	movq	-48(%rsp), %rax         # 8-byte Reload
	adcq	$0, %rax
	adcq	$0, %r12
	movq	%rcx, %rdx
	imulq	-56(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	-24(%rsp), %rbp, %rsi   # 8-byte Folded Reload
	movq	%rbp, -48(%rsp)         # 8-byte Spill
	movq	%rsi, -72(%rsp)         # 8-byte Spill
	mulxq	-16(%rsp), %rbp, %rsi   # 8-byte Folded Reload
	movq	%rbp, -88(%rsp)         # 8-byte Spill
	movq	%rsi, -80(%rsp)         # 8-byte Spill
	mulxq	-32(%rsp), %rsi, %r13   # 8-byte Folded Reload
	movq	%rsi, -112(%rsp)        # 8-byte Spill
	movq	-40(%rsp), %r9          # 8-byte Reload
	mulxq	%r9, %r10, %rbp
	mulxq	(%rsp), %rsi, %r8       # 8-byte Folded Reload
	mulxq	-8(%rsp), %rdx, %rbx    # 8-byte Folded Reload
	addq	%rsi, %rbx
	adcq	%r10, %r8
	adcq	-112(%rsp), %rbp        # 8-byte Folded Reload
	adcq	-88(%rsp), %r13         # 8-byte Folded Reload
	movq	-80(%rsp), %r10         # 8-byte Reload
	adcq	-48(%rsp), %r10         # 8-byte Folded Reload
	movq	-72(%rsp), %rsi         # 8-byte Reload
	adcq	$0, %rsi
	addq	%rcx, %rdx
	adcq	%rdi, %rbx
	adcq	%r15, %r8
	adcq	%r14, %rbp
	adcq	%r11, %r13
	adcq	-96(%rsp), %r10         # 8-byte Folded Reload
	movq	%r10, -80(%rsp)         # 8-byte Spill
	adcq	-64(%rsp), %rsi         # 8-byte Folded Reload
	movq	%rsi, -72(%rsp)         # 8-byte Spill
	adcq	$0, %rax
	movq	%rax, -48(%rsp)         # 8-byte Spill
	adcq	$0, %r12
	movq	-56(%rsp), %rdx         # 8-byte Reload
	imulq	%rbx, %rdx
	mulxq	-24(%rsp), %rax, %r10   # 8-byte Folded Reload
	movq	%rax, -56(%rsp)         # 8-byte Spill
	mulxq	%r9, %rsi, %r14
	mulxq	-8(%rsp), %r11, %rdi    # 8-byte Folded Reload
	mulxq	(%rsp), %rax, %r9       # 8-byte Folded Reload
	addq	%rdi, %rax
	adcq	%rsi, %r9
	movq	-32(%rsp), %r15         # 8-byte Reload
	mulxq	%r15, %rsi, %rdi
	adcq	%r14, %rsi
	mulxq	-16(%rsp), %rdx, %r14   # 8-byte Folded Reload
	adcq	%rdi, %rdx
	adcq	-56(%rsp), %r14         # 8-byte Folded Reload
	adcq	$0, %r10
	addq	%rbx, %r11
	adcq	%r8, %rax
	adcq	%rbp, %r9
	adcq	%r13, %rsi
	adcq	-80(%rsp), %rdx         # 8-byte Folded Reload
	adcq	-72(%rsp), %r14         # 8-byte Folded Reload
	adcq	-48(%rsp), %r10         # 8-byte Folded Reload
	adcq	$0, %r12
	movq	%rax, %rcx
	subq	-8(%rsp), %rcx          # 8-byte Folded Reload
	movq	%r9, %rdi
	sbbq	(%rsp), %rdi            # 8-byte Folded Reload
	movq	%rsi, %rbp
	sbbq	-40(%rsp), %rbp         # 8-byte Folded Reload
	movq	%rdx, %rbx
	sbbq	%r15, %rbx
	movq	%r14, %r8
	sbbq	-16(%rsp), %r8          # 8-byte Folded Reload
	movq	%r10, %r15
	sbbq	-24(%rsp), %r15         # 8-byte Folded Reload
	sbbq	$0, %r12
	andl	$1, %r12d
	cmovneq	%r10, %r15
	testb	%r12b, %r12b
	cmovneq	%rax, %rcx
	movq	-104(%rsp), %rax        # 8-byte Reload
	movq	%rcx, (%rax)
	cmovneq	%r9, %rdi
	movq	%rdi, 8(%rax)
	cmovneq	%rsi, %rbp
	movq	%rbp, 16(%rax)
	cmovneq	%rdx, %rbx
	movq	%rbx, 24(%rax)
	cmovneq	%r14, %r8
	movq	%r8, 32(%rax)
	movq	%r15, 40(%rax)
	addq	$8, %rsp
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end85:
	.size	mcl_fp_montRed6Lbmi2, .Lfunc_end85-mcl_fp_montRed6Lbmi2

	.globl	mcl_fp_addPre6Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addPre6Lbmi2,@function
mcl_fp_addPre6Lbmi2:                    # @mcl_fp_addPre6Lbmi2
# BB#0:
	pushq	%r14
	pushq	%rbx
	movq	40(%rdx), %r8
	movq	40(%rsi), %r11
	movq	32(%rdx), %r9
	movq	24(%rdx), %r10
	movq	24(%rsi), %rax
	movq	32(%rsi), %r14
	movq	16(%rdx), %rbx
	movq	(%rdx), %rcx
	movq	8(%rdx), %rdx
	addq	(%rsi), %rcx
	adcq	8(%rsi), %rdx
	adcq	16(%rsi), %rbx
	movq	%rcx, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%rbx, 16(%rdi)
	adcq	%r10, %rax
	movq	%rax, 24(%rdi)
	adcq	%r9, %r14
	movq	%r14, 32(%rdi)
	adcq	%r8, %r11
	movq	%r11, 40(%rdi)
	sbbq	%rax, %rax
	andl	$1, %eax
	popq	%rbx
	popq	%r14
	retq
.Lfunc_end86:
	.size	mcl_fp_addPre6Lbmi2, .Lfunc_end86-mcl_fp_addPre6Lbmi2

	.globl	mcl_fp_subPre6Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subPre6Lbmi2,@function
mcl_fp_subPre6Lbmi2:                    # @mcl_fp_subPre6Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%rbx
	movq	40(%rdx), %r8
	movq	40(%rsi), %r9
	movq	32(%rsi), %r10
	movq	24(%rsi), %r11
	movq	16(%rsi), %rcx
	movq	(%rsi), %rbx
	movq	8(%rsi), %rsi
	xorl	%eax, %eax
	subq	(%rdx), %rbx
	sbbq	8(%rdx), %rsi
	movq	24(%rdx), %r14
	movq	32(%rdx), %r15
	sbbq	16(%rdx), %rcx
	movq	%rbx, (%rdi)
	movq	%rsi, 8(%rdi)
	movq	%rcx, 16(%rdi)
	sbbq	%r14, %r11
	movq	%r11, 24(%rdi)
	sbbq	%r15, %r10
	movq	%r10, 32(%rdi)
	sbbq	%r8, %r9
	movq	%r9, 40(%rdi)
	sbbq	$0, %rax
	andl	$1, %eax
	popq	%rbx
	popq	%r14
	popq	%r15
	retq
.Lfunc_end87:
	.size	mcl_fp_subPre6Lbmi2, .Lfunc_end87-mcl_fp_subPre6Lbmi2

	.globl	mcl_fp_shr1_6Lbmi2
	.align	16, 0x90
	.type	mcl_fp_shr1_6Lbmi2,@function
mcl_fp_shr1_6Lbmi2:                     # @mcl_fp_shr1_6Lbmi2
# BB#0:
	movq	40(%rsi), %r8
	movq	32(%rsi), %r9
	movq	24(%rsi), %rdx
	movq	16(%rsi), %rax
	movq	(%rsi), %rcx
	movq	8(%rsi), %rsi
	shrdq	$1, %rsi, %rcx
	movq	%rcx, (%rdi)
	shrdq	$1, %rax, %rsi
	movq	%rsi, 8(%rdi)
	shrdq	$1, %rdx, %rax
	movq	%rax, 16(%rdi)
	shrdq	$1, %r9, %rdx
	movq	%rdx, 24(%rdi)
	shrdq	$1, %r8, %r9
	movq	%r9, 32(%rdi)
	shrq	%r8
	movq	%r8, 40(%rdi)
	retq
.Lfunc_end88:
	.size	mcl_fp_shr1_6Lbmi2, .Lfunc_end88-mcl_fp_shr1_6Lbmi2

	.globl	mcl_fp_add6Lbmi2
	.align	16, 0x90
	.type	mcl_fp_add6Lbmi2,@function
mcl_fp_add6Lbmi2:                       # @mcl_fp_add6Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%rbx
	movq	40(%rdx), %r14
	movq	40(%rsi), %r8
	movq	32(%rdx), %r15
	movq	24(%rdx), %rbx
	movq	24(%rsi), %r10
	movq	32(%rsi), %r9
	movq	16(%rdx), %r11
	movq	(%rdx), %rax
	movq	8(%rdx), %rdx
	addq	(%rsi), %rax
	adcq	8(%rsi), %rdx
	adcq	16(%rsi), %r11
	movq	%rax, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%r11, 16(%rdi)
	adcq	%rbx, %r10
	movq	%r10, 24(%rdi)
	adcq	%r15, %r9
	movq	%r9, 32(%rdi)
	adcq	%r14, %r8
	movq	%r8, 40(%rdi)
	sbbq	%rsi, %rsi
	andl	$1, %esi
	subq	(%rcx), %rax
	sbbq	8(%rcx), %rdx
	sbbq	16(%rcx), %r11
	sbbq	24(%rcx), %r10
	sbbq	32(%rcx), %r9
	sbbq	40(%rcx), %r8
	sbbq	$0, %rsi
	testb	$1, %sil
	jne	.LBB89_2
# BB#1:                                 # %nocarry
	movq	%rax, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%r11, 16(%rdi)
	movq	%r10, 24(%rdi)
	movq	%r9, 32(%rdi)
	movq	%r8, 40(%rdi)
.LBB89_2:                               # %carry
	popq	%rbx
	popq	%r14
	popq	%r15
	retq
.Lfunc_end89:
	.size	mcl_fp_add6Lbmi2, .Lfunc_end89-mcl_fp_add6Lbmi2

	.globl	mcl_fp_addNF6Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addNF6Lbmi2,@function
mcl_fp_addNF6Lbmi2:                     # @mcl_fp_addNF6Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	40(%rdx), %r8
	movq	32(%rdx), %r9
	movq	24(%rdx), %r10
	movq	16(%rdx), %r11
	movq	(%rdx), %r15
	movq	8(%rdx), %r14
	addq	(%rsi), %r15
	adcq	8(%rsi), %r14
	adcq	16(%rsi), %r11
	adcq	24(%rsi), %r10
	adcq	32(%rsi), %r9
	adcq	40(%rsi), %r8
	movq	%r15, %rsi
	subq	(%rcx), %rsi
	movq	%r14, %rbx
	sbbq	8(%rcx), %rbx
	movq	%r11, %rdx
	sbbq	16(%rcx), %rdx
	movq	%r10, %r13
	sbbq	24(%rcx), %r13
	movq	%r9, %r12
	sbbq	32(%rcx), %r12
	movq	%r8, %rax
	sbbq	40(%rcx), %rax
	movq	%rax, %rcx
	sarq	$63, %rcx
	cmovsq	%r15, %rsi
	movq	%rsi, (%rdi)
	cmovsq	%r14, %rbx
	movq	%rbx, 8(%rdi)
	cmovsq	%r11, %rdx
	movq	%rdx, 16(%rdi)
	cmovsq	%r10, %r13
	movq	%r13, 24(%rdi)
	cmovsq	%r9, %r12
	movq	%r12, 32(%rdi)
	cmovsq	%r8, %rax
	movq	%rax, 40(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	retq
.Lfunc_end90:
	.size	mcl_fp_addNF6Lbmi2, .Lfunc_end90-mcl_fp_addNF6Lbmi2

	.globl	mcl_fp_sub6Lbmi2
	.align	16, 0x90
	.type	mcl_fp_sub6Lbmi2,@function
mcl_fp_sub6Lbmi2:                       # @mcl_fp_sub6Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%r12
	pushq	%rbx
	movq	40(%rdx), %r14
	movq	40(%rsi), %r8
	movq	32(%rsi), %r9
	movq	24(%rsi), %r10
	movq	16(%rsi), %r11
	movq	(%rsi), %rax
	movq	8(%rsi), %rsi
	xorl	%ebx, %ebx
	subq	(%rdx), %rax
	sbbq	8(%rdx), %rsi
	movq	24(%rdx), %r15
	movq	32(%rdx), %r12
	sbbq	16(%rdx), %r11
	movq	%rax, (%rdi)
	movq	%rsi, 8(%rdi)
	movq	%r11, 16(%rdi)
	sbbq	%r15, %r10
	movq	%r10, 24(%rdi)
	sbbq	%r12, %r9
	movq	%r9, 32(%rdi)
	sbbq	%r14, %r8
	movq	%r8, 40(%rdi)
	sbbq	$0, %rbx
	testb	$1, %bl
	je	.LBB91_2
# BB#1:                                 # %carry
	movq	40(%rcx), %r14
	movq	32(%rcx), %r15
	movq	24(%rcx), %r12
	movq	8(%rcx), %rbx
	movq	16(%rcx), %rdx
	addq	(%rcx), %rax
	movq	%rax, (%rdi)
	adcq	%rsi, %rbx
	movq	%rbx, 8(%rdi)
	adcq	%r11, %rdx
	movq	%rdx, 16(%rdi)
	adcq	%r10, %r12
	movq	%r12, 24(%rdi)
	adcq	%r9, %r15
	movq	%r15, 32(%rdi)
	adcq	%r8, %r14
	movq	%r14, 40(%rdi)
.LBB91_2:                               # %nocarry
	popq	%rbx
	popq	%r12
	popq	%r14
	popq	%r15
	retq
.Lfunc_end91:
	.size	mcl_fp_sub6Lbmi2, .Lfunc_end91-mcl_fp_sub6Lbmi2

	.globl	mcl_fp_subNF6Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subNF6Lbmi2,@function
mcl_fp_subNF6Lbmi2:                     # @mcl_fp_subNF6Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	40(%rsi), %r15
	movq	32(%rsi), %r8
	movq	24(%rsi), %r9
	movq	16(%rsi), %r10
	movq	(%rsi), %r11
	movq	8(%rsi), %r14
	subq	(%rdx), %r11
	sbbq	8(%rdx), %r14
	sbbq	16(%rdx), %r10
	sbbq	24(%rdx), %r9
	sbbq	32(%rdx), %r8
	sbbq	40(%rdx), %r15
	movq	%r15, %rdx
	sarq	$63, %rdx
	movq	%rdx, %rbx
	addq	%rbx, %rbx
	movq	%rdx, %rsi
	adcq	%rsi, %rsi
	andq	8(%rcx), %rsi
	movq	%r15, %rax
	shrq	$63, %rax
	orq	%rbx, %rax
	andq	(%rcx), %rax
	movq	40(%rcx), %r12
	andq	%rdx, %r12
	movq	32(%rcx), %r13
	andq	%rdx, %r13
	movq	24(%rcx), %rbx
	andq	%rdx, %rbx
	andq	16(%rcx), %rdx
	addq	%r11, %rax
	movq	%rax, (%rdi)
	adcq	%r14, %rsi
	movq	%rsi, 8(%rdi)
	adcq	%r10, %rdx
	movq	%rdx, 16(%rdi)
	adcq	%r9, %rbx
	movq	%rbx, 24(%rdi)
	adcq	%r8, %r13
	movq	%r13, 32(%rdi)
	adcq	%r15, %r12
	movq	%r12, 40(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	retq
.Lfunc_end92:
	.size	mcl_fp_subNF6Lbmi2, .Lfunc_end92-mcl_fp_subNF6Lbmi2

	.globl	mcl_fpDbl_add6Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_add6Lbmi2,@function
mcl_fpDbl_add6Lbmi2:                    # @mcl_fpDbl_add6Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	88(%rdx), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	movq	80(%rdx), %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	movq	72(%rdx), %r14
	movq	64(%rdx), %r15
	movq	24(%rsi), %rbp
	movq	32(%rsi), %r13
	movq	16(%rdx), %r12
	movq	(%rdx), %rbx
	movq	8(%rdx), %rax
	addq	(%rsi), %rbx
	adcq	8(%rsi), %rax
	adcq	16(%rsi), %r12
	adcq	24(%rdx), %rbp
	adcq	32(%rdx), %r13
	movq	56(%rdx), %r11
	movq	48(%rdx), %r9
	movq	40(%rdx), %rdx
	movq	%rbx, (%rdi)
	movq	88(%rsi), %r8
	movq	%rax, 8(%rdi)
	movq	80(%rsi), %r10
	movq	%r12, 16(%rdi)
	movq	72(%rsi), %r12
	movq	%rbp, 24(%rdi)
	movq	40(%rsi), %rax
	adcq	%rdx, %rax
	movq	64(%rsi), %rdx
	movq	%r13, 32(%rdi)
	movq	56(%rsi), %r13
	movq	48(%rsi), %rbp
	adcq	%r9, %rbp
	movq	%rax, 40(%rdi)
	adcq	%r11, %r13
	adcq	%r15, %rdx
	adcq	%r14, %r12
	adcq	-16(%rsp), %r10         # 8-byte Folded Reload
	adcq	-8(%rsp), %r8           # 8-byte Folded Reload
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%rbp, %rsi
	subq	(%rcx), %rsi
	movq	%r13, %rbx
	sbbq	8(%rcx), %rbx
	movq	%rdx, %r9
	sbbq	16(%rcx), %r9
	movq	%r12, %r11
	sbbq	24(%rcx), %r11
	movq	%r10, %r14
	sbbq	32(%rcx), %r14
	movq	%r8, %r15
	sbbq	40(%rcx), %r15
	sbbq	$0, %rax
	andl	$1, %eax
	cmovneq	%rbp, %rsi
	movq	%rsi, 48(%rdi)
	testb	%al, %al
	cmovneq	%r13, %rbx
	movq	%rbx, 56(%rdi)
	cmovneq	%rdx, %r9
	movq	%r9, 64(%rdi)
	cmovneq	%r12, %r11
	movq	%r11, 72(%rdi)
	cmovneq	%r10, %r14
	movq	%r14, 80(%rdi)
	cmovneq	%r8, %r15
	movq	%r15, 88(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end93:
	.size	mcl_fpDbl_add6Lbmi2, .Lfunc_end93-mcl_fpDbl_add6Lbmi2

	.globl	mcl_fpDbl_sub6Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sub6Lbmi2,@function
mcl_fpDbl_sub6Lbmi2:                    # @mcl_fpDbl_sub6Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	88(%rdx), %r9
	movq	80(%rdx), %r10
	movq	72(%rdx), %r14
	movq	16(%rsi), %r8
	movq	(%rsi), %r15
	movq	8(%rsi), %r11
	xorl	%eax, %eax
	subq	(%rdx), %r15
	sbbq	8(%rdx), %r11
	sbbq	16(%rdx), %r8
	movq	24(%rsi), %rbx
	sbbq	24(%rdx), %rbx
	movq	32(%rsi), %r12
	sbbq	32(%rdx), %r12
	movq	64(%rdx), %r13
	movq	%r15, (%rdi)
	movq	56(%rdx), %rbp
	movq	%r11, 8(%rdi)
	movq	48(%rdx), %r15
	movq	40(%rdx), %rdx
	movq	%r8, 16(%rdi)
	movq	88(%rsi), %r8
	movq	%rbx, 24(%rdi)
	movq	40(%rsi), %rbx
	sbbq	%rdx, %rbx
	movq	80(%rsi), %r11
	movq	%r12, 32(%rdi)
	movq	48(%rsi), %rdx
	sbbq	%r15, %rdx
	movq	72(%rsi), %r15
	movq	%rbx, 40(%rdi)
	movq	64(%rsi), %r12
	movq	56(%rsi), %rsi
	sbbq	%rbp, %rsi
	sbbq	%r13, %r12
	sbbq	%r14, %r15
	sbbq	%r10, %r11
	sbbq	%r9, %r8
	movl	$0, %ebp
	sbbq	$0, %rbp
	andl	$1, %ebp
	movq	(%rcx), %r14
	cmoveq	%rax, %r14
	testb	%bpl, %bpl
	movq	16(%rcx), %r9
	cmoveq	%rax, %r9
	movq	8(%rcx), %rbp
	cmoveq	%rax, %rbp
	movq	40(%rcx), %r10
	cmoveq	%rax, %r10
	movq	32(%rcx), %rbx
	cmoveq	%rax, %rbx
	cmovneq	24(%rcx), %rax
	addq	%rdx, %r14
	movq	%r14, 48(%rdi)
	adcq	%rsi, %rbp
	movq	%rbp, 56(%rdi)
	adcq	%r12, %r9
	movq	%r9, 64(%rdi)
	adcq	%r15, %rax
	movq	%rax, 72(%rdi)
	adcq	%r11, %rbx
	movq	%rbx, 80(%rdi)
	adcq	%r8, %r10
	movq	%r10, 88(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end94:
	.size	mcl_fpDbl_sub6Lbmi2, .Lfunc_end94-mcl_fpDbl_sub6Lbmi2

	.globl	mcl_fp_mulUnitPre7Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mulUnitPre7Lbmi2,@function
mcl_fp_mulUnitPre7Lbmi2:                # @mcl_fp_mulUnitPre7Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	mulxq	48(%rsi), %r8, %r11
	mulxq	40(%rsi), %r9, %r13
	mulxq	32(%rsi), %r10, %rcx
	mulxq	8(%rsi), %r12, %r14
	mulxq	(%rsi), %r15, %rbx
	addq	%r12, %rbx
	mulxq	24(%rsi), %r12, %rax
	mulxq	16(%rsi), %rdx, %rsi
	movq	%r15, (%rdi)
	movq	%rbx, 8(%rdi)
	adcq	%r14, %rdx
	movq	%rdx, 16(%rdi)
	adcq	%r12, %rsi
	movq	%rsi, 24(%rdi)
	adcq	%r10, %rax
	movq	%rax, 32(%rdi)
	adcq	%r9, %rcx
	movq	%rcx, 40(%rdi)
	adcq	%r8, %r13
	movq	%r13, 48(%rdi)
	adcq	$0, %r11
	movq	%r11, 56(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	retq
.Lfunc_end95:
	.size	mcl_fp_mulUnitPre7Lbmi2, .Lfunc_end95-mcl_fp_mulUnitPre7Lbmi2

	.globl	mcl_fpDbl_mulPre7Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_mulPre7Lbmi2,@function
mcl_fpDbl_mulPre7Lbmi2:                 # @mcl_fpDbl_mulPre7Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rdx, %r14
	movq	%r14, -8(%rsp)          # 8-byte Spill
	movq	%rsi, %r8
	movq	%rdi, %r13
	movq	%r13, -16(%rsp)         # 8-byte Spill
	movq	(%r8), %rcx
	movq	%rcx, -56(%rsp)         # 8-byte Spill
	movq	8(%r8), %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	movq	(%r14), %rsi
	movq	%rax, %rdx
	mulxq	%rsi, %rbp, %rax
	movq	%rcx, %rdx
	mulxq	%rsi, %rdx, %rcx
	movq	%rdx, -64(%rsp)         # 8-byte Spill
	movq	24(%r8), %rdi
	movq	%rdi, -32(%rsp)         # 8-byte Spill
	movq	16(%r8), %rdx
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	addq	%rbp, %rcx
	mulxq	%rsi, %rbx, %rbp
	adcq	%rax, %rbx
	movq	%rdi, %rdx
	mulxq	%rsi, %r12, %rax
	adcq	%rbp, %r12
	movq	32(%r8), %rdx
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	mulxq	%rsi, %r9, %rbp
	adcq	%rax, %r9
	movq	40(%r8), %rdi
	movq	%rdi, %rdx
	mulxq	%rsi, %r10, %rax
	adcq	%rbp, %r10
	movq	48(%r8), %r15
	movq	%r15, %rdx
	mulxq	%rsi, %rsi, %r11
	adcq	%rax, %rsi
	movq	-64(%rsp), %rax         # 8-byte Reload
	movq	%rax, (%r13)
	adcq	$0, %r11
	movq	8(%r14), %r13
	movq	-56(%rsp), %rdx         # 8-byte Reload
	mulxq	%r13, %r14, %rax
	movq	%rax, -56(%rsp)         # 8-byte Spill
	addq	%rcx, %r14
	movq	-24(%rsp), %rdx         # 8-byte Reload
	mulxq	%r13, %rcx, %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	adcq	%rbx, %rcx
	movq	-40(%rsp), %rdx         # 8-byte Reload
	mulxq	%r13, %rbx, %rax
	movq	%rax, -40(%rsp)         # 8-byte Spill
	adcq	%r12, %rbx
	movq	-32(%rsp), %rdx         # 8-byte Reload
	mulxq	%r13, %rbp, %rax
	movq	%rax, -32(%rsp)         # 8-byte Spill
	adcq	%r9, %rbp
	movq	-48(%rsp), %rdx         # 8-byte Reload
	mulxq	%r13, %rax, %r9
	adcq	%r10, %rax
	movq	%rdi, %rdx
	mulxq	%r13, %r10, %rdi
	adcq	%rsi, %r10
	movq	%r15, %rdx
	mulxq	%r13, %r13, %rdx
	adcq	%r11, %r13
	sbbq	%r12, %r12
	andl	$1, %r12d
	addq	-56(%rsp), %rcx         # 8-byte Folded Reload
	adcq	-24(%rsp), %rbx         # 8-byte Folded Reload
	adcq	-40(%rsp), %rbp         # 8-byte Folded Reload
	adcq	-32(%rsp), %rax         # 8-byte Folded Reload
	adcq	%r9, %r10
	movq	-16(%rsp), %rsi         # 8-byte Reload
	movq	%r14, 8(%rsi)
	adcq	%rdi, %r13
	adcq	%rdx, %r12
	movq	(%r8), %rsi
	movq	%rsi, -32(%rsp)         # 8-byte Spill
	movq	8(%r8), %r11
	movq	%r11, -24(%rsp)         # 8-byte Spill
	movq	-8(%rsp), %rdx          # 8-byte Reload
	movq	16(%rdx), %rdi
	movq	%rsi, %rdx
	mulxq	%rdi, %r9, %rdx
	movq	%rdx, -72(%rsp)         # 8-byte Spill
	addq	%rcx, %r9
	movq	%r11, %rdx
	mulxq	%rdi, %r14, %rcx
	movq	%rcx, -80(%rsp)         # 8-byte Spill
	adcq	%rbx, %r14
	movq	16(%r8), %rdx
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	mulxq	%rdi, %rsi, %rcx
	movq	%rcx, -88(%rsp)         # 8-byte Spill
	adcq	%rbp, %rsi
	movq	24(%r8), %rdx
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	mulxq	%rdi, %rbp, %rcx
	movq	%rcx, -96(%rsp)         # 8-byte Spill
	adcq	%rax, %rbp
	movq	32(%r8), %rdx
	movq	%rdx, -56(%rsp)         # 8-byte Spill
	mulxq	%rdi, %r11, %rax
	movq	%rax, -104(%rsp)        # 8-byte Spill
	adcq	%r10, %r11
	movq	40(%r8), %rdx
	movq	%rdx, -64(%rsp)         # 8-byte Spill
	mulxq	%rdi, %r15, %rax
	adcq	%r13, %r15
	movq	48(%r8), %r13
	movq	%r13, %rdx
	mulxq	%rdi, %rcx, %rdx
	adcq	%r12, %rcx
	sbbq	%rbx, %rbx
	andl	$1, %ebx
	addq	-72(%rsp), %r14         # 8-byte Folded Reload
	adcq	-80(%rsp), %rsi         # 8-byte Folded Reload
	adcq	-88(%rsp), %rbp         # 8-byte Folded Reload
	adcq	-96(%rsp), %r11         # 8-byte Folded Reload
	adcq	-104(%rsp), %r15        # 8-byte Folded Reload
	adcq	%rax, %rcx
	adcq	%rdx, %rbx
	movq	-16(%rsp), %rax         # 8-byte Reload
	movq	%r9, 16(%rax)
	movq	-8(%rsp), %rax          # 8-byte Reload
	movq	24(%rax), %rdi
	movq	-32(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %r9, %rax
	movq	%rax, -32(%rsp)         # 8-byte Spill
	addq	%r14, %r9
	movq	-24(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %rax, %rdx
	movq	%rdx, -24(%rsp)         # 8-byte Spill
	adcq	%rsi, %rax
	movq	-40(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %r14, %rdx
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	adcq	%rbp, %r14
	movq	-48(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %r10, %rdx
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	adcq	%r11, %r10
	movq	-56(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %rbp, %rsi
	adcq	%r15, %rbp
	movq	-64(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %r11, %r15
	adcq	%rcx, %r11
	movq	%r13, %rdx
	mulxq	%rdi, %r13, %rcx
	adcq	%rbx, %r13
	sbbq	%r12, %r12
	andl	$1, %r12d
	addq	-32(%rsp), %rax         # 8-byte Folded Reload
	adcq	-24(%rsp), %r14         # 8-byte Folded Reload
	adcq	-40(%rsp), %r10         # 8-byte Folded Reload
	adcq	-48(%rsp), %rbp         # 8-byte Folded Reload
	adcq	%rsi, %r11
	movq	-16(%rsp), %rdi         # 8-byte Reload
	movq	%r9, 24(%rdi)
	adcq	%r15, %r13
	adcq	%rcx, %r12
	movq	(%r8), %rdx
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	movq	8(%r8), %rbx
	movq	%rbx, -24(%rsp)         # 8-byte Spill
	movq	-8(%rsp), %rcx          # 8-byte Reload
	movq	32(%rcx), %rcx
	mulxq	%rcx, %rsi, %rdx
	movq	%rdx, -80(%rsp)         # 8-byte Spill
	addq	%rax, %rsi
	movq	%rbx, %rdx
	mulxq	%rcx, %r9, %rax
	movq	%rax, -88(%rsp)         # 8-byte Spill
	adcq	%r14, %r9
	movq	16(%r8), %rdx
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	mulxq	%rcx, %rax, %rdx
	movq	%rdx, -96(%rsp)         # 8-byte Spill
	adcq	%r10, %rax
	movq	24(%r8), %rdx
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	mulxq	%rcx, %r15, %rdx
	movq	%rdx, -104(%rsp)        # 8-byte Spill
	adcq	%rbp, %r15
	movq	32(%r8), %rdx
	movq	%rdx, -56(%rsp)         # 8-byte Spill
	mulxq	%rcx, %r10, %rbp
	adcq	%r11, %r10
	movq	40(%r8), %rdx
	movq	%rdx, -64(%rsp)         # 8-byte Spill
	mulxq	%rcx, %r11, %rbx
	adcq	%r13, %r11
	movq	48(%r8), %rdx
	movq	%rdx, -72(%rsp)         # 8-byte Spill
	mulxq	%rcx, %r14, %rcx
	adcq	%r12, %r14
	sbbq	%r12, %r12
	andl	$1, %r12d
	addq	-80(%rsp), %r9          # 8-byte Folded Reload
	adcq	-88(%rsp), %rax         # 8-byte Folded Reload
	adcq	-96(%rsp), %r15         # 8-byte Folded Reload
	adcq	-104(%rsp), %r10        # 8-byte Folded Reload
	adcq	%rbp, %r11
	adcq	%rbx, %r14
	adcq	%rcx, %r12
	movq	%rsi, 32(%rdi)
	movq	-8(%rsp), %rsi          # 8-byte Reload
	movq	40(%rsi), %rdi
	movq	-32(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %r13, %rcx
	movq	%rcx, -32(%rsp)         # 8-byte Spill
	addq	%r9, %r13
	movq	-24(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %rcx, %rdx
	movq	%rdx, -24(%rsp)         # 8-byte Spill
	adcq	%rax, %rcx
	movq	-40(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %rax, %rdx
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	adcq	%r15, %rax
	movq	-48(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %rbx, %rdx
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	adcq	%r10, %rbx
	movq	-56(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %rbp, %r15
	adcq	%r11, %rbp
	movq	-64(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %r9, %r11
	adcq	%r14, %r9
	movq	-72(%rsp), %rdx         # 8-byte Reload
	mulxq	%rdi, %r10, %rdx
	adcq	%r12, %r10
	sbbq	%rdi, %rdi
	andl	$1, %edi
	addq	-32(%rsp), %rcx         # 8-byte Folded Reload
	adcq	-24(%rsp), %rax         # 8-byte Folded Reload
	adcq	-40(%rsp), %rbx         # 8-byte Folded Reload
	adcq	-48(%rsp), %rbp         # 8-byte Folded Reload
	adcq	%r15, %r9
	movq	-16(%rsp), %r14         # 8-byte Reload
	movq	%r13, 40(%r14)
	adcq	%r11, %r10
	adcq	%rdx, %rdi
	movq	48(%rsi), %rdx
	mulxq	(%r8), %r11, %rsi
	movq	%rsi, -8(%rsp)          # 8-byte Spill
	addq	%rcx, %r11
	mulxq	8(%r8), %rsi, %r15
	adcq	%rax, %rsi
	mulxq	16(%r8), %rcx, %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	adcq	%rbx, %rcx
	mulxq	24(%r8), %rbx, %r12
	adcq	%rbp, %rbx
	mulxq	32(%r8), %rbp, %r13
	adcq	%r9, %rbp
	mulxq	40(%r8), %rax, %r9
	adcq	%r10, %rax
	mulxq	48(%r8), %rdx, %r8
	adcq	%rdi, %rdx
	sbbq	%r10, %r10
	andl	$1, %r10d
	addq	-8(%rsp), %rsi          # 8-byte Folded Reload
	adcq	%r15, %rcx
	movq	%r11, 48(%r14)
	movq	%rsi, 56(%r14)
	movq	%rcx, 64(%r14)
	adcq	-24(%rsp), %rbx         # 8-byte Folded Reload
	movq	%rbx, 72(%r14)
	adcq	%r12, %rbp
	movq	%rbp, 80(%r14)
	adcq	%r13, %rax
	movq	%rax, 88(%r14)
	adcq	%r9, %rdx
	movq	%rdx, 96(%r14)
	adcq	%r8, %r10
	movq	%r10, 104(%r14)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end96:
	.size	mcl_fpDbl_mulPre7Lbmi2, .Lfunc_end96-mcl_fpDbl_mulPre7Lbmi2

	.globl	mcl_fpDbl_sqrPre7Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sqrPre7Lbmi2,@function
mcl_fpDbl_sqrPre7Lbmi2:                 # @mcl_fpDbl_sqrPre7Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rdi, -8(%rsp)          # 8-byte Spill
	movq	16(%rsi), %rdx
	movq	%rdx, -24(%rsp)         # 8-byte Spill
	movq	(%rsi), %rcx
	movq	8(%rsi), %rax
	mulxq	%rcx, %r8, %r10
	movq	24(%rsi), %rbx
	movq	%rbx, -32(%rsp)         # 8-byte Spill
	movq	%rax, %rdx
	mulxq	%rcx, %r12, %rbp
	movq	%rbp, -16(%rsp)         # 8-byte Spill
	movq	%rcx, %rdx
	mulxq	%rcx, %rdx, %rdi
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	addq	%r12, %rdi
	adcq	%rbp, %r8
	movq	%rbx, %rdx
	mulxq	%rcx, %rbp, %r9
	adcq	%r10, %rbp
	movq	32(%rsi), %rdx
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	mulxq	%rcx, %r11, %r14
	adcq	%r9, %r11
	movq	40(%rsi), %rdx
	movq	%rdx, -56(%rsp)         # 8-byte Spill
	mulxq	%rcx, %r10, %r15
	adcq	%r14, %r10
	movq	48(%rsi), %r14
	movq	%r14, %rdx
	mulxq	%rcx, %rcx, %r13
	adcq	%r15, %rcx
	movq	-8(%rsp), %rdx          # 8-byte Reload
	movq	-48(%rsp), %rbx         # 8-byte Reload
	movq	%rbx, (%rdx)
	adcq	$0, %r13
	addq	%r12, %rdi
	movq	%rax, %rdx
	mulxq	%rax, %r12, %rdx
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	adcq	%r8, %r12
	movq	-24(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %r8, %rdx
	movq	%rdx, -24(%rsp)         # 8-byte Spill
	adcq	%rbp, %r8
	movq	-32(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %r9, %rbp
	adcq	%r11, %r9
	movq	-40(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %r15, %rdx
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	adcq	%r10, %r15
	movq	-56(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %r11, %rbx
	adcq	%rcx, %r11
	movq	%r14, %rdx
	mulxq	%rax, %r14, %rax
	adcq	%r13, %r14
	sbbq	%r13, %r13
	andl	$1, %r13d
	addq	-16(%rsp), %r12         # 8-byte Folded Reload
	adcq	-48(%rsp), %r8          # 8-byte Folded Reload
	adcq	-24(%rsp), %r9          # 8-byte Folded Reload
	adcq	%rbp, %r15
	movq	-8(%rsp), %rcx          # 8-byte Reload
	movq	%rdi, 8(%rcx)
	adcq	-32(%rsp), %r11         # 8-byte Folded Reload
	adcq	%rbx, %r14
	adcq	%rax, %r13
	movq	(%rsi), %rdx
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	movq	8(%rsi), %rcx
	movq	%rcx, -24(%rsp)         # 8-byte Spill
	movq	16(%rsi), %rbx
	mulxq	%rbx, %rax, %rdx
	movq	%rdx, -64(%rsp)         # 8-byte Spill
	addq	%r12, %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	movq	%rcx, %rdx
	mulxq	%rbx, %r10, %rax
	movq	%rax, -72(%rsp)         # 8-byte Spill
	adcq	%r8, %r10
	movq	%rbx, %rdx
	mulxq	%rbx, %r12, %rax
	movq	%rax, -80(%rsp)         # 8-byte Spill
	adcq	%r9, %r12
	movq	24(%rsi), %rax
	movq	%rax, %rdx
	mulxq	%rbx, %r8, %rdi
	movq	%rdi, -56(%rsp)         # 8-byte Spill
	adcq	%r8, %r15
	movq	32(%rsi), %rdx
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	mulxq	%rbx, %rcx, %rdx
	movq	%rdx, -88(%rsp)         # 8-byte Spill
	adcq	%r11, %rcx
	movq	40(%rsi), %rdx
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	mulxq	%rbx, %rbp, %r11
	adcq	%r14, %rbp
	movq	48(%rsi), %r14
	movq	%r14, %rdx
	mulxq	%rbx, %r9, %rdx
	adcq	%r13, %r9
	sbbq	%rbx, %rbx
	andl	$1, %ebx
	addq	-64(%rsp), %r10         # 8-byte Folded Reload
	adcq	-72(%rsp), %r12         # 8-byte Folded Reload
	adcq	-80(%rsp), %r15         # 8-byte Folded Reload
	adcq	%rdi, %rcx
	adcq	-88(%rsp), %rbp         # 8-byte Folded Reload
	adcq	%r11, %r9
	adcq	%rdx, %rbx
	movq	-32(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %rdi, %rdx
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	addq	%r10, %rdi
	movq	-24(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %r11, %rdx
	movq	%rdx, -24(%rsp)         # 8-byte Spill
	adcq	%r12, %r11
	adcq	%r8, %r15
	movq	%rax, %rdx
	mulxq	%rax, %r8, %rdx
	movq	%rdx, -64(%rsp)         # 8-byte Spill
	adcq	%rcx, %r8
	movq	-40(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %r13, %rcx
	movq	%rcx, -40(%rsp)         # 8-byte Spill
	adcq	%rbp, %r13
	movq	-48(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %r12, %rbp
	adcq	%r9, %r12
	movq	%r14, %rdx
	mulxq	%rax, %rcx, %rax
	adcq	%rbx, %rcx
	sbbq	%r10, %r10
	andl	$1, %r10d
	addq	-32(%rsp), %r11         # 8-byte Folded Reload
	adcq	-24(%rsp), %r15         # 8-byte Folded Reload
	adcq	-56(%rsp), %r8          # 8-byte Folded Reload
	adcq	-64(%rsp), %r13         # 8-byte Folded Reload
	movq	-8(%rsp), %rdx          # 8-byte Reload
	movq	-16(%rsp), %rbx         # 8-byte Reload
	movq	%rbx, 16(%rdx)
	movq	%rdi, 24(%rdx)
	adcq	-40(%rsp), %r12         # 8-byte Folded Reload
	adcq	%rbp, %rcx
	adcq	%rax, %r10
	movq	(%rsi), %rdx
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	movq	8(%rsi), %rdi
	movq	%rdi, -24(%rsp)         # 8-byte Spill
	movq	32(%rsi), %rbx
	mulxq	%rbx, %rax, %rdx
	movq	%rdx, -80(%rsp)         # 8-byte Spill
	addq	%r11, %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	movq	%rdi, %rdx
	mulxq	%rbx, %r9, %rax
	movq	%rax, -88(%rsp)         # 8-byte Spill
	adcq	%r15, %r9
	movq	16(%rsi), %rdx
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	mulxq	%rbx, %r15, %rax
	movq	%rax, -96(%rsp)         # 8-byte Spill
	adcq	%r8, %r15
	movq	24(%rsi), %rdx
	movq	%rdx, -48(%rsp)         # 8-byte Spill
	mulxq	%rbx, %r8, %rbp
	adcq	%r13, %r8
	movq	%rbx, %rdx
	mulxq	%rbx, %r13, %r14
	adcq	%r12, %r13
	movq	40(%rsi), %rax
	movq	%rax, %rdx
	mulxq	%rbx, %rdx, %rdi
	movq	%rdx, -72(%rsp)         # 8-byte Spill
	movq	%rdi, -56(%rsp)         # 8-byte Spill
	adcq	%rdx, %rcx
	movq	48(%rsi), %rdx
	movq	%rdx, -64(%rsp)         # 8-byte Spill
	mulxq	%rbx, %r11, %rdx
	adcq	%r10, %r11
	sbbq	%r12, %r12
	andl	$1, %r12d
	addq	-80(%rsp), %r9          # 8-byte Folded Reload
	adcq	-88(%rsp), %r15         # 8-byte Folded Reload
	adcq	-96(%rsp), %r8          # 8-byte Folded Reload
	adcq	%rbp, %r13
	adcq	%r14, %rcx
	adcq	%rdi, %r11
	adcq	%rdx, %r12
	movq	-32(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %r14, %rdi
	addq	%r9, %r14
	movq	-24(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %rbx, %rdx
	movq	%rdx, -24(%rsp)         # 8-byte Spill
	adcq	%r15, %rbx
	movq	-40(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %rbp, %rdx
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	adcq	%r8, %rbp
	movq	-48(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %r10, %r15
	adcq	%r13, %r10
	adcq	-72(%rsp), %rcx         # 8-byte Folded Reload
	movq	%rax, %rdx
	mulxq	%rax, %r9, %r13
	adcq	%r11, %r9
	movq	-64(%rsp), %rdx         # 8-byte Reload
	mulxq	%rax, %rax, %r11
	adcq	%r12, %rax
	sbbq	%r8, %r8
	andl	$1, %r8d
	addq	%rdi, %rbx
	adcq	-24(%rsp), %rbp         # 8-byte Folded Reload
	adcq	-32(%rsp), %r10         # 8-byte Folded Reload
	adcq	%r15, %rcx
	movq	-8(%rsp), %rdi          # 8-byte Reload
	movq	-16(%rsp), %rdx         # 8-byte Reload
	movq	%rdx, 32(%rdi)
	movq	%r14, 40(%rdi)
	adcq	-56(%rsp), %r9          # 8-byte Folded Reload
	adcq	%r13, %rax
	adcq	%r11, %r8
	movq	48(%rsi), %rdx
	mulxq	(%rsi), %r12, %r11
	addq	%rbx, %r12
	mulxq	8(%rsi), %rbx, %r14
	adcq	%rbp, %rbx
	mulxq	16(%rsi), %rbp, %r15
	adcq	%r10, %rbp
	mulxq	24(%rsi), %rdi, %r10
	adcq	%rcx, %rdi
	mulxq	32(%rsi), %rcx, %r13
	adcq	%r9, %rcx
	mulxq	40(%rsi), %rsi, %r9
	adcq	%rax, %rsi
	mulxq	%rdx, %rdx, %rax
	adcq	%r8, %rdx
	sbbq	%r8, %r8
	andl	$1, %r8d
	addq	%r11, %rbx
	adcq	%r14, %rbp
	movq	-8(%rsp), %r11          # 8-byte Reload
	movq	%r12, 48(%r11)
	movq	%rbx, 56(%r11)
	movq	%rbp, 64(%r11)
	adcq	%r15, %rdi
	movq	%rdi, 72(%r11)
	adcq	%r10, %rcx
	movq	%rcx, 80(%r11)
	adcq	%r13, %rsi
	movq	%rsi, 88(%r11)
	adcq	%r9, %rdx
	movq	%rdx, 96(%r11)
	adcq	%rax, %r8
	movq	%r8, 104(%r11)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end97:
	.size	mcl_fpDbl_sqrPre7Lbmi2, .Lfunc_end97-mcl_fpDbl_sqrPre7Lbmi2

	.globl	mcl_fp_mont7Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mont7Lbmi2,@function
mcl_fp_mont7Lbmi2:                      # @mcl_fp_mont7Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	subq	$56, %rsp
	movq	%rdx, -16(%rsp)         # 8-byte Spill
	movq	%rdi, -120(%rsp)        # 8-byte Spill
	movq	48(%rsi), %rdi
	movq	%rdi, -24(%rsp)         # 8-byte Spill
	movq	(%rdx), %rax
	movq	%rdi, %rdx
	mulxq	%rax, %rdx, %r8
	movq	%rdx, 48(%rsp)          # 8-byte Spill
	movq	40(%rsi), %rdx
	movq	%rdx, -32(%rsp)         # 8-byte Spill
	mulxq	%rax, %rdx, %r9
	movq	%rdx, 40(%rsp)          # 8-byte Spill
	movq	32(%rsi), %rdx
	movq	%rdx, -72(%rsp)         # 8-byte Spill
	movq	24(%rsi), %r11
	movq	%r11, -64(%rsp)         # 8-byte Spill
	movq	16(%rsi), %r10
	movq	%r10, -56(%rsp)         # 8-byte Spill
	movq	(%rsi), %r15
	movq	%r15, -40(%rsp)         # 8-byte Spill
	movq	8(%rsi), %rsi
	movq	%rsi, -48(%rsp)         # 8-byte Spill
	mulxq	%rax, %r13, %rdi
	movq	%r11, %rdx
	mulxq	%rax, %r14, %rbp
	movq	%r10, %rdx
	mulxq	%rax, %r12, %rbx
	movq	%rsi, %rdx
	mulxq	%rax, %r10, %rsi
	movq	%r15, %rdx
	mulxq	%rax, %r15, %r11
	addq	%r10, %r11
	adcq	%r12, %rsi
	movq	%rsi, -112(%rsp)        # 8-byte Spill
	adcq	%r14, %rbx
	movq	%rbx, -104(%rsp)        # 8-byte Spill
	adcq	%r13, %rbp
	movq	%rbp, -96(%rsp)         # 8-byte Spill
	adcq	40(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, -88(%rsp)         # 8-byte Spill
	adcq	48(%rsp), %r9           # 8-byte Folded Reload
	movq	%r9, -80(%rsp)          # 8-byte Spill
	adcq	$0, %r8
	movq	%r8, %r13
	movq	-8(%rcx), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	movq	%r15, %rdx
	imulq	%rax, %rdx
	movq	(%rcx), %rdi
	movq	%rdi, 24(%rsp)          # 8-byte Spill
	movq	32(%rcx), %rax
	movq	%rax, 48(%rsp)          # 8-byte Spill
	mulxq	%rax, %rbx, %r9
	movq	16(%rcx), %rsi
	movq	%rsi, 40(%rsp)          # 8-byte Spill
	mulxq	%rsi, %r14, %rbp
	movq	8(%rcx), %rax
	movq	%rax, 32(%rsp)          # 8-byte Spill
	mulxq	%rax, %rsi, %rax
	mulxq	%rdi, %r8, %r12
	addq	%rsi, %r12
	adcq	%r14, %rax
	movq	%rax, %rsi
	movq	24(%rcx), %rax
	movq	%rax, 16(%rsp)          # 8-byte Spill
	mulxq	%rax, %r10, %r14
	adcq	%rbp, %r10
	adcq	%rbx, %r14
	movq	40(%rcx), %rax
	movq	%rax, 8(%rsp)           # 8-byte Spill
	mulxq	%rax, %rbp, %rdi
	adcq	%r9, %rbp
	movq	48(%rcx), %rax
	movq	%rax, (%rsp)            # 8-byte Spill
	mulxq	%rax, %rax, %rbx
	adcq	%rdi, %rax
	adcq	$0, %rbx
	addq	%r15, %r8
	adcq	%r11, %r12
	adcq	-112(%rsp), %rsi        # 8-byte Folded Reload
	movq	%rsi, -112(%rsp)        # 8-byte Spill
	adcq	-104(%rsp), %r10        # 8-byte Folded Reload
	adcq	-96(%rsp), %r14         # 8-byte Folded Reload
	adcq	-88(%rsp), %rbp         # 8-byte Folded Reload
	adcq	-80(%rsp), %rax         # 8-byte Folded Reload
	adcq	%r13, %rbx
	sbbq	%r11, %r11
	andl	$1, %r11d
	movq	-16(%rsp), %rcx         # 8-byte Reload
	movq	8(%rcx), %rdx
	mulxq	-24(%rsp), %rdi, %rcx   # 8-byte Folded Reload
	movq	%rdi, -96(%rsp)         # 8-byte Spill
	movq	%rcx, -80(%rsp)         # 8-byte Spill
	mulxq	-32(%rsp), %rsi, %rcx   # 8-byte Folded Reload
	movq	%rsi, -128(%rsp)        # 8-byte Spill
	movq	%rcx, -88(%rsp)         # 8-byte Spill
	mulxq	-48(%rsp), %r9, %r8     # 8-byte Folded Reload
	mulxq	-40(%rsp), %rsi, %rcx   # 8-byte Folded Reload
	movq	%rsi, -104(%rsp)        # 8-byte Spill
	addq	%r9, %rcx
	movq	%rcx, %rdi
	mulxq	-56(%rsp), %rcx, %r9    # 8-byte Folded Reload
	adcq	%r8, %rcx
	movq	%rcx, %rsi
	mulxq	-64(%rsp), %r13, %rcx   # 8-byte Folded Reload
	adcq	%r9, %r13
	mulxq	-72(%rsp), %r8, %r15    # 8-byte Folded Reload
	adcq	%rcx, %r8
	adcq	-128(%rsp), %r15        # 8-byte Folded Reload
	movq	-88(%rsp), %rdx         # 8-byte Reload
	adcq	-96(%rsp), %rdx         # 8-byte Folded Reload
	movq	-80(%rsp), %rcx         # 8-byte Reload
	adcq	$0, %rcx
	movq	-104(%rsp), %r9         # 8-byte Reload
	addq	%r12, %r9
	movq	%r9, -104(%rsp)         # 8-byte Spill
	adcq	-112(%rsp), %rdi        # 8-byte Folded Reload
	movq	%rdi, %r12
	adcq	%r10, %rsi
	movq	%rsi, -128(%rsp)        # 8-byte Spill
	adcq	%r14, %r13
	adcq	%rbp, %r8
	adcq	%rax, %r15
	adcq	%rbx, %rdx
	movq	%rdx, -88(%rsp)         # 8-byte Spill
	adcq	%r11, %rcx
	movq	%rcx, -80(%rsp)         # 8-byte Spill
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%rax, -112(%rsp)        # 8-byte Spill
	movq	%r9, %rdx
	imulq	-8(%rsp), %rdx          # 8-byte Folded Reload
	mulxq	(%rsp), %r10, %rax      # 8-byte Folded Reload
	movq	%rax, -96(%rsp)         # 8-byte Spill
	mulxq	16(%rsp), %rcx, %rax    # 8-byte Folded Reload
	mulxq	32(%rsp), %rdi, %rbx    # 8-byte Folded Reload
	mulxq	24(%rsp), %r14, %r9     # 8-byte Folded Reload
	addq	%rdi, %r9
	mulxq	40(%rsp), %rbp, %r11    # 8-byte Folded Reload
	adcq	%rbx, %rbp
	adcq	%rcx, %r11
	mulxq	48(%rsp), %rbx, %rsi    # 8-byte Folded Reload
	adcq	%rax, %rbx
	mulxq	8(%rsp), %rax, %rcx     # 8-byte Folded Reload
	adcq	%rsi, %rax
	adcq	%r10, %rcx
	movq	-96(%rsp), %rdx         # 8-byte Reload
	adcq	$0, %rdx
	addq	-104(%rsp), %r14        # 8-byte Folded Reload
	adcq	%r12, %r9
	adcq	-128(%rsp), %rbp        # 8-byte Folded Reload
	adcq	%r13, %r11
	adcq	%r8, %rbx
	adcq	%r15, %rax
	adcq	-88(%rsp), %rcx         # 8-byte Folded Reload
	adcq	-80(%rsp), %rdx         # 8-byte Folded Reload
	movq	%rdx, -96(%rsp)         # 8-byte Spill
	adcq	$0, -112(%rsp)          # 8-byte Folded Spill
	movq	-16(%rsp), %rdx         # 8-byte Reload
	movq	16(%rdx), %rdx
	mulxq	-24(%rsp), %rdi, %rsi   # 8-byte Folded Reload
	movq	%rdi, -104(%rsp)        # 8-byte Spill
	movq	%rsi, -80(%rsp)         # 8-byte Spill
	mulxq	-32(%rsp), %rdi, %rsi   # 8-byte Folded Reload
	movq	%rdi, -128(%rsp)        # 8-byte Spill
	movq	%rsi, -88(%rsp)         # 8-byte Spill
	mulxq	-56(%rsp), %rdi, %r10   # 8-byte Folded Reload
	mulxq	-48(%rsp), %rsi, %r13   # 8-byte Folded Reload
	mulxq	-40(%rsp), %r8, %r15    # 8-byte Folded Reload
	addq	%rsi, %r15
	adcq	%rdi, %r13
	mulxq	-64(%rsp), %r12, %rsi   # 8-byte Folded Reload
	adcq	%r10, %r12
	mulxq	-72(%rsp), %r10, %r14   # 8-byte Folded Reload
	adcq	%rsi, %r10
	adcq	-128(%rsp), %r14        # 8-byte Folded Reload
	movq	-88(%rsp), %rsi         # 8-byte Reload
	adcq	-104(%rsp), %rsi        # 8-byte Folded Reload
	movq	-80(%rsp), %rdx         # 8-byte Reload
	adcq	$0, %rdx
	addq	%r9, %r8
	movq	%r8, -104(%rsp)         # 8-byte Spill
	adcq	%rbp, %r15
	adcq	%r11, %r13
	adcq	%rbx, %r12
	adcq	%rax, %r10
	adcq	%rcx, %r14
	adcq	-96(%rsp), %rsi         # 8-byte Folded Reload
	movq	%rsi, -88(%rsp)         # 8-byte Spill
	adcq	-112(%rsp), %rdx        # 8-byte Folded Reload
	movq	%rdx, -80(%rsp)         # 8-byte Spill
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%rax, %rbx
	movq	%r8, %rdx
	imulq	-8(%rsp), %rdx          # 8-byte Folded Reload
	mulxq	(%rsp), %rcx, %rax      # 8-byte Folded Reload
	movq	%rcx, -112(%rsp)        # 8-byte Spill
	movq	%rax, -96(%rsp)         # 8-byte Spill
	mulxq	16(%rsp), %rcx, %rax    # 8-byte Folded Reload
	mulxq	32(%rsp), %rbp, %rsi    # 8-byte Folded Reload
	mulxq	24(%rsp), %r11, %r8     # 8-byte Folded Reload
	addq	%rbp, %r8
	mulxq	40(%rsp), %rbp, %r9     # 8-byte Folded Reload
	adcq	%rsi, %rbp
	adcq	%rcx, %r9
	mulxq	48(%rsp), %rsi, %rdi    # 8-byte Folded Reload
	adcq	%rax, %rsi
	mulxq	8(%rsp), %rax, %rcx     # 8-byte Folded Reload
	adcq	%rdi, %rax
	adcq	-112(%rsp), %rcx        # 8-byte Folded Reload
	movq	-96(%rsp), %rdx         # 8-byte Reload
	adcq	$0, %rdx
	addq	-104(%rsp), %r11        # 8-byte Folded Reload
	adcq	%r15, %r8
	adcq	%r13, %rbp
	adcq	%r12, %r9
	adcq	%r10, %rsi
	adcq	%r14, %rax
	adcq	-88(%rsp), %rcx         # 8-byte Folded Reload
	adcq	-80(%rsp), %rdx         # 8-byte Folded Reload
	movq	%rdx, -96(%rsp)         # 8-byte Spill
	adcq	$0, %rbx
	movq	%rbx, -88(%rsp)         # 8-byte Spill
	movq	-16(%rsp), %rdx         # 8-byte Reload
	movq	24(%rdx), %rdx
	mulxq	-24(%rsp), %rbx, %rdi   # 8-byte Folded Reload
	movq	%rbx, -112(%rsp)        # 8-byte Spill
	movq	%rdi, -80(%rsp)         # 8-byte Spill
	mulxq	-32(%rsp), %rdi, %r13   # 8-byte Folded Reload
	movq	%rdi, -128(%rsp)        # 8-byte Spill
	mulxq	-56(%rsp), %r10, %r11   # 8-byte Folded Reload
	mulxq	-48(%rsp), %rdi, %r15   # 8-byte Folded Reload
	mulxq	-40(%rsp), %rbx, %r12   # 8-byte Folded Reload
	movq	%rbx, -104(%rsp)        # 8-byte Spill
	addq	%rdi, %r12
	adcq	%r10, %r15
	mulxq	-64(%rsp), %rbx, %rdi   # 8-byte Folded Reload
	adcq	%r11, %rbx
	mulxq	-72(%rsp), %r10, %r14   # 8-byte Folded Reload
	adcq	%rdi, %r10
	adcq	-128(%rsp), %r14        # 8-byte Folded Reload
	adcq	-112(%rsp), %r13        # 8-byte Folded Reload
	movq	-80(%rsp), %rdx         # 8-byte Reload
	adcq	$0, %rdx
	movq	-104(%rsp), %rdi        # 8-byte Reload
	addq	%r8, %rdi
	movq	%rdi, -104(%rsp)        # 8-byte Spill
	adcq	%rbp, %r12
	adcq	%r9, %r15
	adcq	%rsi, %rbx
	adcq	%rax, %r10
	adcq	%rcx, %r14
	adcq	-96(%rsp), %r13         # 8-byte Folded Reload
	adcq	-88(%rsp), %rdx         # 8-byte Folded Reload
	movq	%rdx, -80(%rsp)         # 8-byte Spill
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%rax, -96(%rsp)         # 8-byte Spill
	movq	%rdi, %rdx
	imulq	-8(%rsp), %rdx          # 8-byte Folded Reload
	mulxq	(%rsp), %rcx, %rax      # 8-byte Folded Reload
	movq	%rcx, -112(%rsp)        # 8-byte Spill
	movq	%rax, -88(%rsp)         # 8-byte Spill
	mulxq	16(%rsp), %rcx, %rax    # 8-byte Folded Reload
	mulxq	32(%rsp), %rbp, %rsi    # 8-byte Folded Reload
	mulxq	24(%rsp), %r11, %r8     # 8-byte Folded Reload
	addq	%rbp, %r8
	mulxq	40(%rsp), %rbp, %r9     # 8-byte Folded Reload
	adcq	%rsi, %rbp
	adcq	%rcx, %r9
	mulxq	48(%rsp), %rsi, %rdi    # 8-byte Folded Reload
	adcq	%rax, %rsi
	mulxq	8(%rsp), %rax, %rcx     # 8-byte Folded Reload
	adcq	%rdi, %rax
	adcq	-112(%rsp), %rcx        # 8-byte Folded Reload
	movq	-88(%rsp), %rdx         # 8-byte Reload
	adcq	$0, %rdx
	addq	-104(%rsp), %r11        # 8-byte Folded Reload
	adcq	%r12, %r8
	adcq	%r15, %rbp
	adcq	%rbx, %r9
	adcq	%r10, %rsi
	adcq	%r14, %rax
	adcq	%r13, %rcx
	adcq	-80(%rsp), %rdx         # 8-byte Folded Reload
	movq	%rdx, -88(%rsp)         # 8-byte Spill
	adcq	$0, -96(%rsp)           # 8-byte Folded Spill
	movq	-16(%rsp), %rdx         # 8-byte Reload
	movq	32(%rdx), %rdx
	mulxq	-24(%rsp), %rbx, %rdi   # 8-byte Folded Reload
	movq	%rbx, -104(%rsp)        # 8-byte Spill
	movq	%rdi, -80(%rsp)         # 8-byte Spill
	mulxq	-32(%rsp), %rdi, %r11   # 8-byte Folded Reload
	movq	%rdi, -112(%rsp)        # 8-byte Spill
	mulxq	-56(%rsp), %r10, %r13   # 8-byte Folded Reload
	mulxq	-48(%rsp), %rdi, %r15   # 8-byte Folded Reload
	mulxq	-40(%rsp), %rbx, %r12   # 8-byte Folded Reload
	addq	%rdi, %r12
	adcq	%r10, %r15
	mulxq	-64(%rsp), %r10, %rdi   # 8-byte Folded Reload
	adcq	%r13, %r10
	mulxq	-72(%rsp), %r13, %r14   # 8-byte Folded Reload
	adcq	%rdi, %r13
	adcq	-112(%rsp), %r14        # 8-byte Folded Reload
	adcq	-104(%rsp), %r11        # 8-byte Folded Reload
	movq	-80(%rsp), %rdx         # 8-byte Reload
	adcq	$0, %rdx
	addq	%r8, %rbx
	movq	%rbx, -112(%rsp)        # 8-byte Spill
	adcq	%rbp, %r12
	adcq	%r9, %r15
	adcq	%rsi, %r10
	adcq	%rax, %r13
	adcq	%rcx, %r14
	adcq	-88(%rsp), %r11         # 8-byte Folded Reload
	movq	%r11, -128(%rsp)        # 8-byte Spill
	adcq	-96(%rsp), %rdx         # 8-byte Folded Reload
	movq	%rdx, -80(%rsp)         # 8-byte Spill
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%rax, -104(%rsp)        # 8-byte Spill
	movq	%rbx, %rdx
	imulq	-8(%rsp), %rdx          # 8-byte Folded Reload
	mulxq	(%rsp), %rcx, %rax      # 8-byte Folded Reload
	movq	%rcx, -88(%rsp)         # 8-byte Spill
	movq	%rax, -96(%rsp)         # 8-byte Spill
	mulxq	16(%rsp), %rcx, %rax    # 8-byte Folded Reload
	mulxq	32(%rsp), %rbp, %rsi    # 8-byte Folded Reload
	mulxq	24(%rsp), %r9, %r11     # 8-byte Folded Reload
	addq	%rbp, %r11
	mulxq	40(%rsp), %rbp, %r8     # 8-byte Folded Reload
	adcq	%rsi, %rbp
	adcq	%rcx, %r8
	mulxq	48(%rsp), %rsi, %rdi    # 8-byte Folded Reload
	adcq	%rax, %rsi
	mulxq	8(%rsp), %rax, %rcx     # 8-byte Folded Reload
	adcq	%rdi, %rax
	adcq	-88(%rsp), %rcx         # 8-byte Folded Reload
	movq	-96(%rsp), %rdx         # 8-byte Reload
	adcq	$0, %rdx
	addq	-112(%rsp), %r9         # 8-byte Folded Reload
	adcq	%r12, %r11
	adcq	%r15, %rbp
	adcq	%r10, %r8
	adcq	%r13, %rsi
	adcq	%r14, %rax
	adcq	-128(%rsp), %rcx        # 8-byte Folded Reload
	adcq	-80(%rsp), %rdx         # 8-byte Folded Reload
	movq	%rdx, -96(%rsp)         # 8-byte Spill
	adcq	$0, -104(%rsp)          # 8-byte Folded Spill
	movq	-16(%rsp), %rdx         # 8-byte Reload
	movq	40(%rdx), %rdx
	mulxq	-24(%rsp), %rbx, %rdi   # 8-byte Folded Reload
	movq	%rbx, -112(%rsp)        # 8-byte Spill
	movq	%rdi, -80(%rsp)         # 8-byte Spill
	mulxq	-32(%rsp), %rbx, %rdi   # 8-byte Folded Reload
	movq	%rbx, -128(%rsp)        # 8-byte Spill
	movq	%rdi, -88(%rsp)         # 8-byte Spill
	mulxq	-56(%rsp), %rbx, %r10   # 8-byte Folded Reload
	mulxq	-48(%rsp), %rdi, %r13   # 8-byte Folded Reload
	mulxq	-40(%rsp), %r9, %r12    # 8-byte Folded Reload
	addq	%rdi, %r12
	adcq	%rbx, %r13
	mulxq	-64(%rsp), %r15, %rdi   # 8-byte Folded Reload
	adcq	%r10, %r15
	mulxq	-72(%rsp), %r10, %r14   # 8-byte Folded Reload
	adcq	%rdi, %r10
	adcq	-128(%rsp), %r14        # 8-byte Folded Reload
	movq	-88(%rsp), %rdi         # 8-byte Reload
	adcq	-112(%rsp), %rdi        # 8-byte Folded Reload
	movq	-80(%rsp), %rdx         # 8-byte Reload
	adcq	$0, %rdx
	addq	%r11, %r9
	movq	%r9, -112(%rsp)         # 8-byte Spill
	adcq	%rbp, %r12
	adcq	%r8, %r13
	adcq	%rsi, %r15
	adcq	%rax, %r10
	adcq	%rcx, %r14
	adcq	-96(%rsp), %rdi         # 8-byte Folded Reload
	movq	%rdi, -88(%rsp)         # 8-byte Spill
	adcq	-104(%rsp), %rdx        # 8-byte Folded Reload
	movq	%rdx, -80(%rsp)         # 8-byte Spill
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	%rax, -104(%rsp)        # 8-byte Spill
	movq	%r9, %rdx
	imulq	-8(%rsp), %rdx          # 8-byte Folded Reload
	mulxq	(%rsp), %rcx, %rax      # 8-byte Folded Reload
	movq	%rcx, -128(%rsp)        # 8-byte Spill
	movq	%rax, -96(%rsp)         # 8-byte Spill
	mulxq	16(%rsp), %rcx, %rax    # 8-byte Folded Reload
	mulxq	32(%rsp), %rdi, %rsi    # 8-byte Folded Reload
	mulxq	24(%rsp), %r11, %rbx    # 8-byte Folded Reload
	addq	%rdi, %rbx
	mulxq	40(%rsp), %r8, %r9      # 8-byte Folded Reload
	adcq	%rsi, %r8
	adcq	%rcx, %r9
	mulxq	48(%rsp), %rdi, %rbp    # 8-byte Folded Reload
	adcq	%rax, %rdi
	mulxq	8(%rsp), %rcx, %rsi     # 8-byte Folded Reload
	adcq	%rbp, %rcx
	adcq	-128(%rsp), %rsi        # 8-byte Folded Reload
	movq	-96(%rsp), %rax         # 8-byte Reload
	adcq	$0, %rax
	addq	-112(%rsp), %r11        # 8-byte Folded Reload
	adcq	%r12, %rbx
	adcq	%r13, %r8
	adcq	%r15, %r9
	adcq	%r10, %rdi
	adcq	%r14, %rcx
	adcq	-88(%rsp), %rsi         # 8-byte Folded Reload
	adcq	-80(%rsp), %rax         # 8-byte Folded Reload
	movq	%rax, -96(%rsp)         # 8-byte Spill
	movq	-104(%rsp), %r12        # 8-byte Reload
	adcq	$0, %r12
	movq	-16(%rsp), %rax         # 8-byte Reload
	movq	48(%rax), %rdx
	mulxq	-24(%rsp), %rbp, %rax   # 8-byte Folded Reload
	movq	%rbp, -80(%rsp)         # 8-byte Spill
	movq	%rax, -16(%rsp)         # 8-byte Spill
	mulxq	-32(%rsp), %rbp, %rax   # 8-byte Folded Reload
	movq	%rbp, -88(%rsp)         # 8-byte Spill
	movq	%rax, -24(%rsp)         # 8-byte Spill
	mulxq	-72(%rsp), %rbp, %rax   # 8-byte Folded Reload
	movq	%rbp, -72(%rsp)         # 8-byte Spill
	movq	%rax, -32(%rsp)         # 8-byte Spill
	mulxq	-64(%rsp), %r13, %rbp   # 8-byte Folded Reload
	mulxq	-56(%rsp), %r14, %r15   # 8-byte Folded Reload
	mulxq	-48(%rsp), %rax, %r11   # 8-byte Folded Reload
	mulxq	-40(%rsp), %rdx, %r10   # 8-byte Folded Reload
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	addq	%rax, %r10
	adcq	%r14, %r11
	adcq	%r13, %r15
	adcq	-72(%rsp), %rbp         # 8-byte Folded Reload
	movq	-32(%rsp), %r14         # 8-byte Reload
	adcq	-88(%rsp), %r14         # 8-byte Folded Reload
	movq	-24(%rsp), %rdx         # 8-byte Reload
	adcq	-80(%rsp), %rdx         # 8-byte Folded Reload
	movq	-16(%rsp), %rax         # 8-byte Reload
	adcq	$0, %rax
	movq	-40(%rsp), %r13         # 8-byte Reload
	addq	%rbx, %r13
	movq	%r13, -40(%rsp)         # 8-byte Spill
	adcq	%r8, %r10
	adcq	%r9, %r11
	adcq	%rdi, %r15
	adcq	%rcx, %rbp
	movq	%rbp, -48(%rsp)         # 8-byte Spill
	adcq	%rsi, %r14
	movq	%r14, -32(%rsp)         # 8-byte Spill
	adcq	-96(%rsp), %rdx         # 8-byte Folded Reload
	movq	%rdx, -24(%rsp)         # 8-byte Spill
	adcq	%r12, %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	sbbq	%rdi, %rdi
	movq	-8(%rsp), %rdx          # 8-byte Reload
	imulq	%r13, %rdx
	mulxq	16(%rsp), %rbp, %rsi    # 8-byte Folded Reload
	mulxq	32(%rsp), %rcx, %rbx    # 8-byte Folded Reload
	mulxq	24(%rsp), %r13, %rax    # 8-byte Folded Reload
	addq	%rcx, %rax
	mulxq	40(%rsp), %rcx, %r9     # 8-byte Folded Reload
	adcq	%rbx, %rcx
	adcq	%rbp, %r9
	mulxq	48(%rsp), %rbp, %rbx    # 8-byte Folded Reload
	adcq	%rsi, %rbp
	mulxq	8(%rsp), %rsi, %r14     # 8-byte Folded Reload
	adcq	%rbx, %rsi
	mulxq	(%rsp), %rdx, %rbx      # 8-byte Folded Reload
	adcq	%r14, %rdx
	adcq	$0, %rbx
	andl	$1, %edi
	addq	-40(%rsp), %r13         # 8-byte Folded Reload
	adcq	%r10, %rax
	adcq	%r11, %rcx
	adcq	%r15, %r9
	adcq	-48(%rsp), %rbp         # 8-byte Folded Reload
	adcq	-32(%rsp), %rsi         # 8-byte Folded Reload
	adcq	-24(%rsp), %rdx         # 8-byte Folded Reload
	adcq	-16(%rsp), %rbx         # 8-byte Folded Reload
	adcq	$0, %rdi
	movq	%rax, %r8
	subq	24(%rsp), %r8           # 8-byte Folded Reload
	movq	%rcx, %r10
	sbbq	32(%rsp), %r10          # 8-byte Folded Reload
	movq	%r9, %r11
	sbbq	40(%rsp), %r11          # 8-byte Folded Reload
	movq	%rbp, %r14
	sbbq	16(%rsp), %r14          # 8-byte Folded Reload
	movq	%rsi, %r15
	sbbq	48(%rsp), %r15          # 8-byte Folded Reload
	movq	%rdx, %r12
	sbbq	8(%rsp), %r12           # 8-byte Folded Reload
	movq	%rbx, %r13
	sbbq	(%rsp), %r13            # 8-byte Folded Reload
	sbbq	$0, %rdi
	andl	$1, %edi
	cmovneq	%rbx, %r13
	testb	%dil, %dil
	cmovneq	%rax, %r8
	movq	-120(%rsp), %rax        # 8-byte Reload
	movq	%r8, (%rax)
	cmovneq	%rcx, %r10
	movq	%r10, 8(%rax)
	cmovneq	%r9, %r11
	movq	%r11, 16(%rax)
	cmovneq	%rbp, %r14
	movq	%r14, 24(%rax)
	cmovneq	%rsi, %r15
	movq	%r15, 32(%rax)
	cmovneq	%rdx, %r12
	movq	%r12, 40(%rax)
	movq	%r13, 48(%rax)
	addq	$56, %rsp
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end98:
	.size	mcl_fp_mont7Lbmi2, .Lfunc_end98-mcl_fp_mont7Lbmi2

	.globl	mcl_fp_montNF7Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montNF7Lbmi2,@function
mcl_fp_montNF7Lbmi2:                    # @mcl_fp_montNF7Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	subq	$40, %rsp
	movq	%rdx, -40(%rsp)         # 8-byte Spill
	movq	%rdi, -104(%rsp)        # 8-byte Spill
	movq	(%rsi), %rax
	movq	%rax, -32(%rsp)         # 8-byte Spill
	movq	8(%rsi), %rdi
	movq	%rdi, -48(%rsp)         # 8-byte Spill
	movq	(%rdx), %rbp
	movq	%rdi, %rdx
	mulxq	%rbp, %rdi, %rbx
	movq	%rax, %rdx
	mulxq	%rbp, %r8, %r14
	movq	16(%rsi), %rdx
	movq	%rdx, -56(%rsp)         # 8-byte Spill
	addq	%rdi, %r14
	mulxq	%rbp, %r15, %rax
	adcq	%rbx, %r15
	movq	24(%rsi), %rdx
	movq	%rdx, -64(%rsp)         # 8-byte Spill
	mulxq	%rbp, %rbx, %rdi
	adcq	%rax, %rbx
	movq	32(%rsi), %rdx
	movq	%rdx, -72(%rsp)         # 8-byte Spill
	mulxq	%rbp, %r11, %rax
	adcq	%rdi, %r11
	movq	40(%rsi), %rdx
	movq	%rdx, -80(%rsp)         # 8-byte Spill
	mulxq	%rbp, %r9, %rdi
	adcq	%rax, %r9
	movq	48(%rsi), %rdx
	movq	%rdx, -88(%rsp)         # 8-byte Spill
	mulxq	%rbp, %r10, %rbp
	adcq	%rdi, %r10
	adcq	$0, %rbp
	movq	-8(%rcx), %rax
	movq	%rax, 8(%rsp)           # 8-byte Spill
	movq	%r8, %rdx
	imulq	%rax, %rdx
	movq	(%rcx), %rax
	movq	%rax, 24(%rsp)          # 8-byte Spill
	mulxq	%rax, %rax, %rsi
	movq	%rsi, -96(%rsp)         # 8-byte Spill
	addq	%r8, %rax
	movq	8(%rcx), %rax
	movq	%rax, 16(%rsp)          # 8-byte Spill
	mulxq	%rax, %r8, %rsi
	movq	%rsi, -112(%rsp)        # 8-byte Spill
	adcq	%r14, %r8
	movq	16(%rcx), %rax
	movq	%rax, (%rsp)            # 8-byte Spill
	mulxq	%rax, %rsi, %r13
	adcq	%r15, %rsi
	movq	24(%rcx), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	mulxq	%rax, %r12, %rax
	adcq	%rbx, %r12
	movq	32(%rcx), %rdi
	movq	%rdi, -16(%rsp)         # 8-byte Spill
	mulxq	%rdi, %r15, %rbx
	adcq	%r11, %r15
	movq	40(%rcx), %rdi
	movq	%rdi, -24(%rsp)         # 8-byte Spill
	mulxq	%rdi, %r14, %rdi
	adcq	%r9, %r14
	movq	48(%rcx), %rcx
	movq	%rcx, 32(%rsp)          # 8-byte Spill
	mulxq	%rcx, %r11, %rcx
	adcq	%r10, %r11
	adcq	$0, %rbp
	addq	-96(%rsp), %r8          # 8-byte Folded Reload
	movq	%r8, -96(%rsp)          # 8-byte Spill
	adcq	-112(%rsp), %rsi        # 8-byte Folded Reload
	movq	%rsi, -112(%rsp)        # 8-byte Spill
	adcq	%r13, %r12
	adcq	%rax, %r15
	adcq	%rbx, %r14
	adcq	%rdi, %r11
	adcq	%rcx, %rbp
	movq	-40(%rsp), %rax         # 8-byte Reload
	movq	8(%rax), %rdx
	mulxq	-48(%rsp), %rcx, %rsi   # 8-byte Folded Reload
	mulxq	-32(%rsp), %r13, %rax   # 8-byte Folded Reload
	addq	%rcx, %rax
	mulxq	-56(%rsp), %rcx, %rdi   # 8-byte Folded Reload
	adcq	%rsi, %rcx
	mulxq	-64(%rsp), %rsi, %r8    # 8-byte Folded Reload
	adcq	%rdi, %rsi
	mulxq	-72(%rsp), %rdi, %r9    # 8-byte Folded Reload
	adcq	%r8, %rdi
	mulxq	-80(%rsp), %r8, %rbx    # 8-byte Folded Reload
	adcq	%r9, %r8
	mulxq	-88(%rsp), %r9, %r10    # 8-byte Folded Reload
	adcq	%rbx, %r9
	adcq	$0, %r10
	addq	-96(%rsp), %r13         # 8-byte Folded Reload
	adcq	-112(%rsp), %rax        # 8-byte Folded Reload
	adcq	%r12, %rcx
	adcq	%r15, %rsi
	adcq	%r14, %rdi
	adcq	%r11, %r8
	adcq	%rbp, %r9
	adcq	$0, %r10
	movq	%r13, %rdx
	imulq	8(%rsp), %rdx           # 8-byte Folded Reload
	mulxq	24(%rsp), %rbp, %rbx    # 8-byte Folded Reload
	movq	%rbx, -96(%rsp)         # 8-byte Spill
	addq	%r13, %rbp
	mulxq	16(%rsp), %rbp, %r14    # 8-byte Folded Reload
	adcq	%rax, %rbp
	mulxq	(%rsp), %rax, %r11      # 8-byte Folded Reload
	adcq	%rcx, %rax
	mulxq	-8(%rsp), %r12, %rcx    # 8-byte Folded Reload
	adcq	%rsi, %r12
	mulxq	-16(%rsp), %r15, %rbx   # 8-byte Folded Reload
	adcq	%rdi, %r15
	mulxq	-24(%rsp), %r13, %rdi   # 8-byte Folded Reload
	adcq	%r8, %r13
	mulxq	32(%rsp), %rsi, %rdx    # 8-byte Folded Reload
	adcq	%r9, %rsi
	adcq	$0, %r10
	addq	-96(%rsp), %rbp         # 8-byte Folded Reload
	movq	%rbp, -96(%rsp)         # 8-byte Spill
	adcq	%r14, %rax
	movq	%rax, -112(%rsp)        # 8-byte Spill
	adcq	%r11, %r12
	adcq	%rcx, %r15
	adcq	%rbx, %r13
	adcq	%rdi, %rsi
	adcq	%rdx, %r10
	movq	-40(%rsp), %rax         # 8-byte Reload
	movq	16(%rax), %rdx
	mulxq	-48(%rsp), %rcx, %rax   # 8-byte Folded Reload
	mulxq	-32(%rsp), %r14, %rdi   # 8-byte Folded Reload
	addq	%rcx, %rdi
	mulxq	-56(%rsp), %rbp, %rcx   # 8-byte Folded Reload
	adcq	%rax, %rbp
	mulxq	-64(%rsp), %rbx, %r8    # 8-byte Folded Reload
	adcq	%rcx, %rbx
	mulxq	-72(%rsp), %rax, %r9    # 8-byte Folded Reload
	adcq	%r8, %rax
	mulxq	-80(%rsp), %r8, %rcx    # 8-byte Folded Reload
	movq	%rcx, -120(%rsp)        # 8-byte Spill
	adcq	%r9, %r8
	mulxq	-88(%rsp), %r9, %r11    # 8-byte Folded Reload
	adcq	-120(%rsp), %r9         # 8-byte Folded Reload
	adcq	$0, %r11
	addq	-96(%rsp), %r14         # 8-byte Folded Reload
	adcq	-112(%rsp), %rdi        # 8-byte Folded Reload
	adcq	%r12, %rbp
	adcq	%r15, %rbx
	adcq	%r13, %rax
	adcq	%rsi, %r8
	adcq	%r10, %r9
	adcq	$0, %r11
	movq	%r14, %rdx
	imulq	8(%rsp), %rdx           # 8-byte Folded Reload
	mulxq	24(%rsp), %rsi, %rcx    # 8-byte Folded Reload
	movq	%rcx, -96(%rsp)         # 8-byte Spill
	addq	%r14, %rsi
	mulxq	16(%rsp), %rsi, %r13    # 8-byte Folded Reload
	adcq	%rdi, %rsi
	mulxq	(%rsp), %rdi, %r15      # 8-byte Folded Reload
	adcq	%rbp, %rdi
	mulxq	-8(%rsp), %rcx, %rbp    # 8-byte Folded Reload
	adcq	%rbx, %rcx
	mulxq	-16(%rsp), %r14, %rbx   # 8-byte Folded Reload
	adcq	%rax, %r14
	mulxq	-24(%rsp), %r12, %rax   # 8-byte Folded Reload
	adcq	%r8, %r12
	mulxq	32(%rsp), %r10, %rdx    # 8-byte Folded Reload
	adcq	%r9, %r10
	adcq	$0, %r11
	addq	-96(%rsp), %rsi         # 8-byte Folded Reload
	movq	%rsi, -96(%rsp)         # 8-byte Spill
	adcq	%r13, %rdi
	movq	%rdi, -112(%rsp)        # 8-byte Spill
	adcq	%r15, %rcx
	adcq	%rbp, %r14
	adcq	%rbx, %r12
	adcq	%rax, %r10
	adcq	%rdx, %r11
	movq	-40(%rsp), %rax         # 8-byte Reload
	movq	24(%rax), %rdx
	mulxq	-48(%rsp), %rsi, %rax   # 8-byte Folded Reload
	mulxq	-32(%rsp), %r15, %rbp   # 8-byte Folded Reload
	addq	%rsi, %rbp
	mulxq	-56(%rsp), %rbx, %rdi   # 8-byte Folded Reload
	adcq	%rax, %rbx
	mulxq	-64(%rsp), %rsi, %rax   # 8-byte Folded Reload
	adcq	%rdi, %rsi
	mulxq	-72(%rsp), %rdi, %r9    # 8-byte Folded Reload
	adcq	%rax, %rdi
	mulxq	-80(%rsp), %r8, %rax    # 8-byte Folded Reload
	adcq	%r9, %r8
	mulxq	-88(%rsp), %r9, %r13    # 8-byte Folded Reload
	adcq	%rax, %r9
	adcq	$0, %r13
	addq	-96(%rsp), %r15         # 8-byte Folded Reload
	adcq	-112(%rsp), %rbp        # 8-byte Folded Reload
	adcq	%rcx, %rbx
	adcq	%r14, %rsi
	adcq	%r12, %rdi
	adcq	%r10, %r8
	adcq	%r11, %r9
	adcq	$0, %r13
	movq	%r15, %rdx
	imulq	8(%rsp), %rdx           # 8-byte Folded Reload
	mulxq	24(%rsp), %rcx, %rax    # 8-byte Folded Reload
	movq	%rax, -96(%rsp)         # 8-byte Spill
	addq	%r15, %rcx
	mulxq	16(%rsp), %rcx, %r11    # 8-byte Folded Reload
	adcq	%rbp, %rcx
	mulxq	(%rsp), %rbp, %r10      # 8-byte Folded Reload
	adcq	%rbx, %rbp
	mulxq	-8(%rsp), %rax, %rbx    # 8-byte Folded Reload
	adcq	%rsi, %rax
	mulxq	-16(%rsp), %r14, %rsi   # 8-byte Folded Reload
	adcq	%rdi, %r14
	mulxq	-24(%rsp), %r15, %rdi   # 8-byte Folded Reload
	adcq	%r8, %r15
	mulxq	32(%rsp), %r12, %rdx    # 8-byte Folded Reload
	adcq	%r9, %r12
	adcq	$0, %r13
	addq	-96(%rsp), %rcx         # 8-byte Folded Reload
	adcq	%r11, %rbp
	movq	%rbp, -96(%rsp)         # 8-byte Spill
	adcq	%r10, %rax
	movq	%rax, -112(%rsp)        # 8-byte Spill
	adcq	%rbx, %r14
	adcq	%rsi, %r15
	adcq	%rdi, %r12
	adcq	%rdx, %r13
	movq	-40(%rsp), %rax         # 8-byte Reload
	movq	32(%rax), %rdx
	mulxq	-48(%rsp), %rsi, %rdi   # 8-byte Folded Reload
	mulxq	-32(%rsp), %r11, %r8    # 8-byte Folded Reload
	addq	%rsi, %r8
	mulxq	-56(%rsp), %rbx, %rsi   # 8-byte Folded Reload
	adcq	%rdi, %rbx
	mulxq	-64(%rsp), %rbp, %rdi   # 8-byte Folded Reload
	adcq	%rsi, %rbp
	mulxq	-72(%rsp), %rsi, %r9    # 8-byte Folded Reload
	adcq	%rdi, %rsi
	mulxq	-80(%rsp), %rdi, %rax   # 8-byte Folded Reload
	adcq	%r9, %rdi
	mulxq	-88(%rsp), %r9, %r10    # 8-byte Folded Reload
	adcq	%rax, %r9
	adcq	$0, %r10
	addq	%rcx, %r11
	adcq	-96(%rsp), %r8          # 8-byte Folded Reload
	adcq	-112(%rsp), %rbx        # 8-byte Folded Reload
	adcq	%r14, %rbp
	adcq	%r15, %rsi
	adcq	%r12, %rdi
	adcq	%r13, %r9
	adcq	$0, %r10
	movq	%r11, %rdx
	imulq	8(%rsp), %rdx           # 8-byte Folded Reload
	mulxq	24(%rsp), %rcx, %rax    # 8-byte Folded Reload
	movq	%rax, -96(%rsp)         # 8-byte Spill
	addq	%r11, %rcx
	mulxq	16(%rsp), %rcx, %r13    # 8-byte Folded Reload
	adcq	%r8, %rcx
	mulxq	(%rsp), %rax, %r8       # 8-byte Folded Reload
	adcq	%rbx, %rax
	mulxq	-8(%rsp), %rbx, %r11    # 8-byte Folded Reload
	adcq	%rbp, %rbx
	mulxq	-16(%rsp), %r14, %rbp   # 8-byte Folded Reload
	adcq	%rsi, %r14
	mulxq	-24(%rsp), %r15, %rsi   # 8-byte Folded Reload
	adcq	%rdi, %r15
	mulxq	32(%rsp), %r12, %rdx    # 8-byte Folded Reload
	adcq	%r9, %r12
	adcq	$0, %r10
	addq	-96(%rsp), %rcx         # 8-byte Folded Reload
	adcq	%r13, %rax
	movq	%rax, -96(%rsp)         # 8-byte Spill
	adcq	%r8, %rbx
	movq	%rbx, -112(%rsp)        # 8-byte Spill
	adcq	%r11, %r14
	adcq	%rbp, %r15
	adcq	%rsi, %r12
	adcq	%rdx, %r10
	movq	-40(%rsp), %rax         # 8-byte Reload
	movq	40(%rax), %rdx
	mulxq	-48(%rsp), %rsi, %rax   # 8-byte Folded Reload
	mulxq	-32(%rsp), %r11, %rbp   # 8-byte Folded Reload
	addq	%rsi, %rbp
	mulxq	-56(%rsp), %rbx, %rdi   # 8-byte Folded Reload
	adcq	%rax, %rbx
	mulxq	-64(%rsp), %rsi, %rax   # 8-byte Folded Reload
	adcq	%rdi, %rsi
	mulxq	-72(%rsp), %rdi, %r9    # 8-byte Folded Reload
	adcq	%rax, %rdi
	mulxq	-80(%rsp), %r8, %rax    # 8-byte Folded Reload
	adcq	%r9, %r8
	mulxq	-88(%rsp), %r9, %r13    # 8-byte Folded Reload
	adcq	%rax, %r9
	adcq	$0, %r13
	addq	%rcx, %r11
	adcq	-96(%rsp), %rbp         # 8-byte Folded Reload
	adcq	-112(%rsp), %rbx        # 8-byte Folded Reload
	adcq	%r14, %rsi
	adcq	%r15, %rdi
	adcq	%r12, %r8
	adcq	%r10, %r9
	adcq	$0, %r13
	movq	%r11, %rdx
	imulq	8(%rsp), %rdx           # 8-byte Folded Reload
	mulxq	24(%rsp), %rcx, %rax    # 8-byte Folded Reload
	movq	%rax, -112(%rsp)        # 8-byte Spill
	addq	%r11, %rcx
	mulxq	16(%rsp), %rcx, %rax    # 8-byte Folded Reload
	movq	%rax, -120(%rsp)        # 8-byte Spill
	adcq	%rbp, %rcx
	mulxq	(%rsp), %rax, %rbp      # 8-byte Folded Reload
	movq	%rbp, -128(%rsp)        # 8-byte Spill
	adcq	%rbx, %rax
	movq	%rax, -96(%rsp)         # 8-byte Spill
	mulxq	-8(%rsp), %r14, %rbp    # 8-byte Folded Reload
	adcq	%rsi, %r14
	mulxq	-16(%rsp), %r11, %r12   # 8-byte Folded Reload
	adcq	%rdi, %r11
	mulxq	-24(%rsp), %r10, %rbx   # 8-byte Folded Reload
	adcq	%r8, %r10
	mulxq	32(%rsp), %rdi, %rax    # 8-byte Folded Reload
	adcq	%r9, %rdi
	adcq	$0, %r13
	addq	-112(%rsp), %rcx        # 8-byte Folded Reload
	movq	%rcx, -112(%rsp)        # 8-byte Spill
	movq	-96(%rsp), %rcx         # 8-byte Reload
	adcq	-120(%rsp), %rcx        # 8-byte Folded Reload
	movq	%rcx, -96(%rsp)         # 8-byte Spill
	adcq	-128(%rsp), %r14        # 8-byte Folded Reload
	adcq	%rbp, %r11
	adcq	%r12, %r10
	adcq	%rbx, %rdi
	adcq	%rax, %r13
	movq	-40(%rsp), %rax         # 8-byte Reload
	movq	48(%rax), %rdx
	mulxq	-48(%rsp), %rbp, %r9    # 8-byte Folded Reload
	mulxq	-32(%rsp), %r8, %rax    # 8-byte Folded Reload
	addq	%rbp, %rax
	mulxq	-56(%rsp), %rbx, %rcx   # 8-byte Folded Reload
	adcq	%r9, %rbx
	mulxq	-64(%rsp), %rbp, %r9    # 8-byte Folded Reload
	adcq	%rcx, %rbp
	mulxq	-72(%rsp), %rcx, %r12   # 8-byte Folded Reload
	adcq	%r9, %rcx
	mulxq	-80(%rsp), %r15, %rsi   # 8-byte Folded Reload
	movq	%rsi, -32(%rsp)         # 8-byte Spill
	adcq	%r12, %r15
	mulxq	-88(%rsp), %r12, %r9    # 8-byte Folded Reload
	adcq	-32(%rsp), %r12         # 8-byte Folded Reload
	adcq	$0, %r9
	addq	-112(%rsp), %r8         # 8-byte Folded Reload
	adcq	-96(%rsp), %rax         # 8-byte Folded Reload
	adcq	%r14, %rbx
	adcq	%r11, %rbp
	adcq	%r10, %rcx
	adcq	%rdi, %r15
	adcq	%r13, %r12
	adcq	$0, %r9
	movq	8(%rsp), %rdx           # 8-byte Reload
	imulq	%r8, %rdx
	mulxq	24(%rsp), %rdi, %rsi    # 8-byte Folded Reload
	movq	%rsi, 8(%rsp)           # 8-byte Spill
	addq	%r8, %rdi
	mulxq	16(%rsp), %r8, %rsi     # 8-byte Folded Reload
	movq	%rsi, -32(%rsp)         # 8-byte Spill
	adcq	%rax, %r8
	movq	(%rsp), %r11            # 8-byte Reload
	mulxq	%r11, %rsi, %rax
	movq	%rax, -40(%rsp)         # 8-byte Spill
	adcq	%rbx, %rsi
	movq	-8(%rsp), %r14          # 8-byte Reload
	mulxq	%r14, %rdi, %rax
	movq	%rax, -48(%rsp)         # 8-byte Spill
	adcq	%rbp, %rdi
	movq	-16(%rsp), %rbp         # 8-byte Reload
	mulxq	%rbp, %rax, %rbx
	movq	%rbx, -56(%rsp)         # 8-byte Spill
	adcq	%rcx, %rax
	movq	-24(%rsp), %rbx         # 8-byte Reload
	mulxq	%rbx, %rcx, %r13
	adcq	%r15, %rcx
	mulxq	32(%rsp), %rdx, %r15    # 8-byte Folded Reload
	adcq	%r12, %rdx
	adcq	$0, %r9
	addq	8(%rsp), %r8            # 8-byte Folded Reload
	adcq	-32(%rsp), %rsi         # 8-byte Folded Reload
	adcq	-40(%rsp), %rdi         # 8-byte Folded Reload
	adcq	-48(%rsp), %rax         # 8-byte Folded Reload
	adcq	-56(%rsp), %rcx         # 8-byte Folded Reload
	adcq	%r13, %rdx
	adcq	%r15, %r9
	movq	%r8, %r13
	subq	24(%rsp), %r13          # 8-byte Folded Reload
	movq	%rsi, %r12
	sbbq	16(%rsp), %r12          # 8-byte Folded Reload
	movq	%rdi, %r10
	sbbq	%r11, %r10
	movq	%rax, %r11
	sbbq	%r14, %r11
	movq	%rcx, %r14
	sbbq	%rbp, %r14
	movq	%rdx, %r15
	sbbq	%rbx, %r15
	movq	%r9, %rbp
	sbbq	32(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rbp, %rbx
	sarq	$63, %rbx
	cmovsq	%r8, %r13
	movq	-104(%rsp), %rbx        # 8-byte Reload
	movq	%r13, (%rbx)
	cmovsq	%rsi, %r12
	movq	%r12, 8(%rbx)
	cmovsq	%rdi, %r10
	movq	%r10, 16(%rbx)
	cmovsq	%rax, %r11
	movq	%r11, 24(%rbx)
	cmovsq	%rcx, %r14
	movq	%r14, 32(%rbx)
	cmovsq	%rdx, %r15
	movq	%r15, 40(%rbx)
	cmovsq	%r9, %rbp
	movq	%rbp, 48(%rbx)
	addq	$40, %rsp
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end99:
	.size	mcl_fp_montNF7Lbmi2, .Lfunc_end99-mcl_fp_montNF7Lbmi2

	.globl	mcl_fp_montRed7Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montRed7Lbmi2,@function
mcl_fp_montRed7Lbmi2:                   # @mcl_fp_montRed7Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	subq	$64, %rsp
	movq	%rdx, %rcx
	movq	%rdi, -88(%rsp)         # 8-byte Spill
	movq	-8(%rcx), %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	movq	(%rcx), %rbx
	movq	%rbx, 32(%rsp)          # 8-byte Spill
	movq	(%rsi), %rdx
	movq	%rdx, 48(%rsp)          # 8-byte Spill
	imulq	%rax, %rdx
	movq	48(%rcx), %rax
	movq	%rax, 16(%rsp)          # 8-byte Spill
	mulxq	%rax, %rdi, %rax
	movq	%rdi, 40(%rsp)          # 8-byte Spill
	movq	%rax, -48(%rsp)         # 8-byte Spill
	movq	40(%rcx), %r8
	movq	%r8, (%rsp)             # 8-byte Spill
	movq	32(%rcx), %r9
	movq	%r9, 24(%rsp)           # 8-byte Spill
	movq	24(%rcx), %rbp
	movq	%rbp, 8(%rsp)           # 8-byte Spill
	movq	16(%rcx), %rdi
	movq	%rdi, 56(%rsp)          # 8-byte Spill
	movq	8(%rcx), %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	mulxq	%r8, %r10, %r11
	mulxq	%r9, %r14, %r9
	mulxq	%rbp, %r8, %r13
	mulxq	%rdi, %rcx, %r12
	mulxq	%rax, %rbp, %r15
	mulxq	%rbx, %rdx, %rdi
	addq	%rbp, %rdi
	adcq	%rcx, %r15
	adcq	%r8, %r12
	adcq	%r14, %r13
	adcq	%r10, %r9
	adcq	40(%rsp), %r11          # 8-byte Folded Reload
	movq	-48(%rsp), %rcx         # 8-byte Reload
	adcq	$0, %rcx
	addq	48(%rsp), %rdx          # 8-byte Folded Reload
	adcq	8(%rsi), %rdi
	adcq	16(%rsi), %r15
	adcq	24(%rsi), %r12
	adcq	32(%rsi), %r13
	adcq	40(%rsi), %r9
	movq	%r9, -96(%rsp)          # 8-byte Spill
	adcq	48(%rsi), %r11
	movq	%r11, -72(%rsp)         # 8-byte Spill
	adcq	56(%rsi), %rcx
	movq	%rcx, -48(%rsp)         # 8-byte Spill
	movq	104(%rsi), %r8
	movq	96(%rsi), %rdx
	movq	88(%rsi), %rbp
	movq	80(%rsi), %rbx
	movq	72(%rsi), %rcx
	movq	64(%rsi), %rsi
	adcq	$0, %rsi
	movq	%rsi, -104(%rsp)        # 8-byte Spill
	adcq	$0, %rcx
	movq	%rcx, -80(%rsp)         # 8-byte Spill
	adcq	$0, %rbx
	movq	%rbx, -40(%rsp)         # 8-byte Spill
	adcq	$0, %rbp
	movq	%rbp, -32(%rsp)         # 8-byte Spill
	adcq	$0, %rdx
	movq	%rdx, -8(%rsp)          # 8-byte Spill
	adcq	$0, %r8
	movq	%r8, 48(%rsp)           # 8-byte Spill
	sbbq	%rcx, %rcx
	andl	$1, %ecx
	movq	%rcx, 40(%rsp)          # 8-byte Spill
	movq	%rdi, %rdx
	movq	-24(%rsp), %r9          # 8-byte Reload
	imulq	%r9, %rdx
	mulxq	16(%rsp), %rsi, %rcx    # 8-byte Folded Reload
	movq	%rsi, -112(%rsp)        # 8-byte Spill
	movq	%rcx, -56(%rsp)         # 8-byte Spill
	mulxq	(%rsp), %rsi, %rcx      # 8-byte Folded Reload
	movq	%rsi, -120(%rsp)        # 8-byte Spill
	movq	%rcx, -64(%rsp)         # 8-byte Spill
	movq	24(%rsp), %rbx          # 8-byte Reload
	mulxq	%rbx, %rcx, %rbp
	movq	%rcx, -128(%rsp)        # 8-byte Spill
	mulxq	8(%rsp), %r10, %r14     # 8-byte Folded Reload
	mulxq	56(%rsp), %rsi, %r11    # 8-byte Folded Reload
	mulxq	%rax, %rcx, %r8
	mulxq	32(%rsp), %rdx, %rax    # 8-byte Folded Reload
	addq	%rcx, %rax
	adcq	%rsi, %r8
	adcq	%r10, %r11
	adcq	-128(%rsp), %r14        # 8-byte Folded Reload
	adcq	-120(%rsp), %rbp        # 8-byte Folded Reload
	movq	-64(%rsp), %rsi         # 8-byte Reload
	adcq	-112(%rsp), %rsi        # 8-byte Folded Reload
	movq	-56(%rsp), %rcx         # 8-byte Reload
	adcq	$0, %rcx
	addq	%rdi, %rdx
	adcq	%r15, %rax
	adcq	%r12, %r8
	adcq	%r13, %r11
	adcq	-96(%rsp), %r14         # 8-byte Folded Reload
	adcq	-72(%rsp), %rbp         # 8-byte Folded Reload
	adcq	-48(%rsp), %rsi         # 8-byte Folded Reload
	movq	%rsi, -64(%rsp)         # 8-byte Spill
	adcq	-104(%rsp), %rcx        # 8-byte Folded Reload
	movq	%rcx, -56(%rsp)         # 8-byte Spill
	adcq	$0, -80(%rsp)           # 8-byte Folded Spill
	adcq	$0, -40(%rsp)           # 8-byte Folded Spill
	adcq	$0, -32(%rsp)           # 8-byte Folded Spill
	adcq	$0, -8(%rsp)            # 8-byte Folded Spill
	adcq	$0, 48(%rsp)            # 8-byte Folded Spill
	adcq	$0, 40(%rsp)            # 8-byte Folded Spill
	movq	%rax, %rdx
	imulq	%r9, %rdx
	mulxq	16(%rsp), %rsi, %rcx    # 8-byte Folded Reload
	movq	%rsi, -96(%rsp)         # 8-byte Spill
	movq	%rcx, -48(%rsp)         # 8-byte Spill
	movq	(%rsp), %r15            # 8-byte Reload
	mulxq	%r15, %rsi, %rcx
	movq	%rsi, -104(%rsp)        # 8-byte Spill
	movq	%rcx, -72(%rsp)         # 8-byte Spill
	mulxq	%rbx, %rcx, %r13
	movq	%rcx, -112(%rsp)        # 8-byte Spill
	mulxq	8(%rsp), %rbx, %r12     # 8-byte Folded Reload
	mulxq	56(%rsp), %rdi, %r9     # 8-byte Folded Reload
	mulxq	-16(%rsp), %rsi, %r10   # 8-byte Folded Reload
	mulxq	32(%rsp), %rdx, %rcx    # 8-byte Folded Reload
	addq	%rsi, %rcx
	adcq	%rdi, %r10
	adcq	%rbx, %r9
	adcq	-112(%rsp), %r12        # 8-byte Folded Reload
	adcq	-104(%rsp), %r13        # 8-byte Folded Reload
	movq	-72(%rsp), %rdi         # 8-byte Reload
	adcq	-96(%rsp), %rdi         # 8-byte Folded Reload
	movq	-48(%rsp), %rsi         # 8-byte Reload
	adcq	$0, %rsi
	addq	%rax, %rdx
	adcq	%r8, %rcx
	adcq	%r11, %r10
	adcq	%r14, %r9
	adcq	%rbp, %r12
	adcq	-64(%rsp), %r13         # 8-byte Folded Reload
	adcq	-56(%rsp), %rdi         # 8-byte Folded Reload
	movq	%rdi, -72(%rsp)         # 8-byte Spill
	adcq	-80(%rsp), %rsi         # 8-byte Folded Reload
	movq	%rsi, -48(%rsp)         # 8-byte Spill
	adcq	$0, -40(%rsp)           # 8-byte Folded Spill
	adcq	$0, -32(%rsp)           # 8-byte Folded Spill
	adcq	$0, -8(%rsp)            # 8-byte Folded Spill
	adcq	$0, 48(%rsp)            # 8-byte Folded Spill
	adcq	$0, 40(%rsp)            # 8-byte Folded Spill
	movq	%rcx, %rdx
	imulq	-24(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	16(%rsp), %rsi, %rax    # 8-byte Folded Reload
	movq	%rsi, -80(%rsp)         # 8-byte Spill
	movq	%rax, -56(%rsp)         # 8-byte Spill
	mulxq	%r15, %rsi, %rax
	movq	%rsi, -96(%rsp)         # 8-byte Spill
	movq	%rax, -64(%rsp)         # 8-byte Spill
	mulxq	24(%rsp), %r8, %r15     # 8-byte Folded Reload
	mulxq	8(%rsp), %r14, %rbp     # 8-byte Folded Reload
	mulxq	56(%rsp), %rdi, %rbx    # 8-byte Folded Reload
	mulxq	-16(%rsp), %rsi, %r11   # 8-byte Folded Reload
	mulxq	32(%rsp), %rdx, %rax    # 8-byte Folded Reload
	addq	%rsi, %rax
	adcq	%rdi, %r11
	adcq	%r14, %rbx
	adcq	%r8, %rbp
	adcq	-96(%rsp), %r15         # 8-byte Folded Reload
	movq	-64(%rsp), %rdi         # 8-byte Reload
	adcq	-80(%rsp), %rdi         # 8-byte Folded Reload
	movq	-56(%rsp), %rsi         # 8-byte Reload
	adcq	$0, %rsi
	addq	%rcx, %rdx
	adcq	%r10, %rax
	adcq	%r9, %r11
	adcq	%r12, %rbx
	adcq	%r13, %rbp
	adcq	-72(%rsp), %r15         # 8-byte Folded Reload
	adcq	-48(%rsp), %rdi         # 8-byte Folded Reload
	movq	%rdi, -64(%rsp)         # 8-byte Spill
	adcq	-40(%rsp), %rsi         # 8-byte Folded Reload
	movq	%rsi, -56(%rsp)         # 8-byte Spill
	adcq	$0, -32(%rsp)           # 8-byte Folded Spill
	adcq	$0, -8(%rsp)            # 8-byte Folded Spill
	adcq	$0, 48(%rsp)            # 8-byte Folded Spill
	adcq	$0, 40(%rsp)            # 8-byte Folded Spill
	movq	%rax, %rdx
	imulq	-24(%rsp), %rdx         # 8-byte Folded Reload
	movq	16(%rsp), %r10          # 8-byte Reload
	mulxq	%r10, %rsi, %rcx
	movq	%rsi, -72(%rsp)         # 8-byte Spill
	movq	%rcx, -40(%rsp)         # 8-byte Spill
	mulxq	(%rsp), %rsi, %rcx      # 8-byte Folded Reload
	movq	%rsi, -80(%rsp)         # 8-byte Spill
	movq	%rcx, -48(%rsp)         # 8-byte Spill
	mulxq	24(%rsp), %rsi, %rcx    # 8-byte Folded Reload
	movq	%rsi, -96(%rsp)         # 8-byte Spill
	mulxq	8(%rsp), %r12, %r13     # 8-byte Folded Reload
	mulxq	56(%rsp), %r8, %r14     # 8-byte Folded Reload
	mulxq	-16(%rsp), %rsi, %r9    # 8-byte Folded Reload
	mulxq	32(%rsp), %rdx, %rdi    # 8-byte Folded Reload
	addq	%rsi, %rdi
	adcq	%r8, %r9
	adcq	%r12, %r14
	adcq	-96(%rsp), %r13         # 8-byte Folded Reload
	adcq	-80(%rsp), %rcx         # 8-byte Folded Reload
	movq	-48(%rsp), %r8          # 8-byte Reload
	adcq	-72(%rsp), %r8          # 8-byte Folded Reload
	movq	-40(%rsp), %rsi         # 8-byte Reload
	adcq	$0, %rsi
	addq	%rax, %rdx
	adcq	%r11, %rdi
	adcq	%rbx, %r9
	adcq	%rbp, %r14
	adcq	%r15, %r13
	adcq	-64(%rsp), %rcx         # 8-byte Folded Reload
	movq	%rcx, -64(%rsp)         # 8-byte Spill
	adcq	-56(%rsp), %r8          # 8-byte Folded Reload
	movq	%r8, -48(%rsp)          # 8-byte Spill
	adcq	-32(%rsp), %rsi         # 8-byte Folded Reload
	movq	%rsi, -40(%rsp)         # 8-byte Spill
	adcq	$0, -8(%rsp)            # 8-byte Folded Spill
	adcq	$0, 48(%rsp)            # 8-byte Folded Spill
	adcq	$0, 40(%rsp)            # 8-byte Folded Spill
	movq	%rdi, %rdx
	imulq	-24(%rsp), %rdx         # 8-byte Folded Reload
	mulxq	%r10, %rcx, %rax
	movq	%rcx, -72(%rsp)         # 8-byte Spill
	movq	%rax, -32(%rsp)         # 8-byte Spill
	mulxq	(%rsp), %rcx, %rax      # 8-byte Folded Reload
	movq	%rcx, -80(%rsp)         # 8-byte Spill
	movq	%rax, -56(%rsp)         # 8-byte Spill
	mulxq	24(%rsp), %rax, %rcx    # 8-byte Folded Reload
	movq	%rax, -96(%rsp)         # 8-byte Spill
	movq	8(%rsp), %r12           # 8-byte Reload
	mulxq	%r12, %rax, %r15
	movq	%rax, -104(%rsp)        # 8-byte Spill
	mulxq	56(%rsp), %rsi, %r11    # 8-byte Folded Reload
	movq	-16(%rsp), %r10         # 8-byte Reload
	mulxq	%r10, %rax, %rbp
	movq	32(%rsp), %rbx          # 8-byte Reload
	mulxq	%rbx, %rdx, %r8
	addq	%rax, %r8
	adcq	%rsi, %rbp
	adcq	-104(%rsp), %r11        # 8-byte Folded Reload
	adcq	-96(%rsp), %r15         # 8-byte Folded Reload
	adcq	-80(%rsp), %rcx         # 8-byte Folded Reload
	movq	-56(%rsp), %rsi         # 8-byte Reload
	adcq	-72(%rsp), %rsi         # 8-byte Folded Reload
	movq	-32(%rsp), %rax         # 8-byte Reload
	adcq	$0, %rax
	addq	%rdi, %rdx
	adcq	%r9, %r8
	adcq	%r14, %rbp
	adcq	%r13, %r11
	adcq	-64(%rsp), %r15         # 8-byte Folded Reload
	adcq	-48(%rsp), %rcx         # 8-byte Folded Reload
	movq	%rcx, -48(%rsp)         # 8-byte Spill
	adcq	-40(%rsp), %rsi         # 8-byte Folded Reload
	movq	%rsi, -56(%rsp)         # 8-byte Spill
	adcq	-8(%rsp), %rax          # 8-byte Folded Reload
	movq	%rax, -32(%rsp)         # 8-byte Spill
	adcq	$0, 48(%rsp)            # 8-byte Folded Spill
	adcq	$0, 40(%rsp)            # 8-byte Folded Spill
	movq	-24(%rsp), %rdx         # 8-byte Reload
	imulq	%r8, %rdx
	mulxq	%r12, %rax, %r13
	mulxq	%r10, %rcx, %rdi
	mulxq	%rbx, %r12, %r14
	addq	%rcx, %r14
	mulxq	56(%rsp), %rcx, %r10    # 8-byte Folded Reload
	adcq	%rdi, %rcx
	adcq	%rax, %r10
	mulxq	24(%rsp), %rax, %r9     # 8-byte Folded Reload
	adcq	%r13, %rax
	mulxq	(%rsp), %rdi, %r13      # 8-byte Folded Reload
	adcq	%r9, %rdi
	mulxq	16(%rsp), %rdx, %rsi    # 8-byte Folded Reload
	adcq	%r13, %rdx
	adcq	$0, %rsi
	addq	%r8, %r12
	adcq	%rbp, %r14
	adcq	%r11, %rcx
	adcq	%r15, %r10
	adcq	-48(%rsp), %rax         # 8-byte Folded Reload
	adcq	-56(%rsp), %rdi         # 8-byte Folded Reload
	adcq	-32(%rsp), %rdx         # 8-byte Folded Reload
	adcq	48(%rsp), %rsi          # 8-byte Folded Reload
	movq	40(%rsp), %rbx          # 8-byte Reload
	adcq	$0, %rbx
	movq	%r14, %rbp
	subq	32(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rcx, %r13
	sbbq	-16(%rsp), %r13         # 8-byte Folded Reload
	movq	%r10, %r8
	sbbq	56(%rsp), %r8           # 8-byte Folded Reload
	movq	%rax, %r9
	sbbq	8(%rsp), %r9            # 8-byte Folded Reload
	movq	%rdi, %r11
	sbbq	24(%rsp), %r11          # 8-byte Folded Reload
	movq	%rdx, %r15
	sbbq	(%rsp), %r15            # 8-byte Folded Reload
	movq	%rsi, %r12
	sbbq	16(%rsp), %r12          # 8-byte Folded Reload
	sbbq	$0, %rbx
	andl	$1, %ebx
	cmovneq	%rsi, %r12
	testb	%bl, %bl
	cmovneq	%r14, %rbp
	movq	-88(%rsp), %rsi         # 8-byte Reload
	movq	%rbp, (%rsi)
	cmovneq	%rcx, %r13
	movq	%r13, 8(%rsi)
	cmovneq	%r10, %r8
	movq	%r8, 16(%rsi)
	cmovneq	%rax, %r9
	movq	%r9, 24(%rsi)
	cmovneq	%rdi, %r11
	movq	%r11, 32(%rsi)
	cmovneq	%rdx, %r15
	movq	%r15, 40(%rsi)
	movq	%r12, 48(%rsi)
	addq	$64, %rsp
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end100:
	.size	mcl_fp_montRed7Lbmi2, .Lfunc_end100-mcl_fp_montRed7Lbmi2

	.globl	mcl_fp_addPre7Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addPre7Lbmi2,@function
mcl_fp_addPre7Lbmi2:                    # @mcl_fp_addPre7Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%r12
	pushq	%rbx
	movq	48(%rdx), %r8
	movq	48(%rsi), %r14
	movq	40(%rdx), %r9
	movq	40(%rsi), %r15
	movq	32(%rdx), %r10
	movq	24(%rdx), %r11
	movq	16(%rdx), %r12
	movq	(%rdx), %rcx
	movq	8(%rdx), %rdx
	addq	(%rsi), %rcx
	adcq	8(%rsi), %rdx
	movq	24(%rsi), %rax
	movq	32(%rsi), %rbx
	adcq	16(%rsi), %r12
	movq	%rcx, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%r12, 16(%rdi)
	adcq	%r11, %rax
	movq	%rax, 24(%rdi)
	adcq	%r10, %rbx
	movq	%rbx, 32(%rdi)
	adcq	%r9, %r15
	movq	%r15, 40(%rdi)
	adcq	%r8, %r14
	movq	%r14, 48(%rdi)
	sbbq	%rax, %rax
	andl	$1, %eax
	popq	%rbx
	popq	%r12
	popq	%r14
	popq	%r15
	retq
.Lfunc_end101:
	.size	mcl_fp_addPre7Lbmi2, .Lfunc_end101-mcl_fp_addPre7Lbmi2

	.globl	mcl_fp_subPre7Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subPre7Lbmi2,@function
mcl_fp_subPre7Lbmi2:                    # @mcl_fp_subPre7Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%r12
	pushq	%rbx
	movq	48(%rdx), %r8
	movq	48(%rsi), %r10
	movq	40(%rdx), %r9
	movq	40(%rsi), %r15
	movq	24(%rdx), %r11
	movq	32(%rdx), %r14
	movq	(%rsi), %rbx
	movq	8(%rsi), %r12
	xorl	%eax, %eax
	subq	(%rdx), %rbx
	sbbq	8(%rdx), %r12
	movq	16(%rsi), %rcx
	sbbq	16(%rdx), %rcx
	movq	32(%rsi), %rdx
	movq	24(%rsi), %rsi
	movq	%rbx, (%rdi)
	movq	%r12, 8(%rdi)
	movq	%rcx, 16(%rdi)
	sbbq	%r11, %rsi
	movq	%rsi, 24(%rdi)
	sbbq	%r14, %rdx
	movq	%rdx, 32(%rdi)
	sbbq	%r9, %r15
	movq	%r15, 40(%rdi)
	sbbq	%r8, %r10
	movq	%r10, 48(%rdi)
	sbbq	$0, %rax
	andl	$1, %eax
	popq	%rbx
	popq	%r12
	popq	%r14
	popq	%r15
	retq
.Lfunc_end102:
	.size	mcl_fp_subPre7Lbmi2, .Lfunc_end102-mcl_fp_subPre7Lbmi2

	.globl	mcl_fp_shr1_7Lbmi2
	.align	16, 0x90
	.type	mcl_fp_shr1_7Lbmi2,@function
mcl_fp_shr1_7Lbmi2:                     # @mcl_fp_shr1_7Lbmi2
# BB#0:
	movq	48(%rsi), %r8
	movq	40(%rsi), %r9
	movq	32(%rsi), %r10
	movq	24(%rsi), %rax
	movq	16(%rsi), %rcx
	movq	(%rsi), %rdx
	movq	8(%rsi), %rsi
	shrdq	$1, %rsi, %rdx
	movq	%rdx, (%rdi)
	shrdq	$1, %rcx, %rsi
	movq	%rsi, 8(%rdi)
	shrdq	$1, %rax, %rcx
	movq	%rcx, 16(%rdi)
	shrdq	$1, %r10, %rax
	movq	%rax, 24(%rdi)
	shrdq	$1, %r9, %r10
	movq	%r10, 32(%rdi)
	shrdq	$1, %r8, %r9
	movq	%r9, 40(%rdi)
	shrq	%r8
	movq	%r8, 48(%rdi)
	retq
.Lfunc_end103:
	.size	mcl_fp_shr1_7Lbmi2, .Lfunc_end103-mcl_fp_shr1_7Lbmi2

	.globl	mcl_fp_add7Lbmi2
	.align	16, 0x90
	.type	mcl_fp_add7Lbmi2,@function
mcl_fp_add7Lbmi2:                       # @mcl_fp_add7Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	48(%rdx), %r14
	movq	48(%rsi), %r8
	movq	40(%rdx), %r15
	movq	40(%rsi), %r9
	movq	32(%rdx), %r12
	movq	24(%rdx), %r13
	movq	16(%rdx), %r10
	movq	(%rdx), %r11
	movq	8(%rdx), %rdx
	addq	(%rsi), %r11
	adcq	8(%rsi), %rdx
	movq	24(%rsi), %rax
	movq	32(%rsi), %rbx
	adcq	16(%rsi), %r10
	movq	%r11, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%r10, 16(%rdi)
	adcq	%r13, %rax
	movq	%rax, 24(%rdi)
	adcq	%r12, %rbx
	movq	%rbx, 32(%rdi)
	adcq	%r15, %r9
	movq	%r9, 40(%rdi)
	adcq	%r14, %r8
	movq	%r8, 48(%rdi)
	sbbq	%rsi, %rsi
	andl	$1, %esi
	subq	(%rcx), %r11
	sbbq	8(%rcx), %rdx
	sbbq	16(%rcx), %r10
	sbbq	24(%rcx), %rax
	sbbq	32(%rcx), %rbx
	sbbq	40(%rcx), %r9
	sbbq	48(%rcx), %r8
	sbbq	$0, %rsi
	testb	$1, %sil
	jne	.LBB104_2
# BB#1:                                 # %nocarry
	movq	%r11, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%r10, 16(%rdi)
	movq	%rax, 24(%rdi)
	movq	%rbx, 32(%rdi)
	movq	%r9, 40(%rdi)
	movq	%r8, 48(%rdi)
.LBB104_2:                              # %carry
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	retq
.Lfunc_end104:
	.size	mcl_fp_add7Lbmi2, .Lfunc_end104-mcl_fp_add7Lbmi2

	.globl	mcl_fp_addNF7Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addNF7Lbmi2,@function
mcl_fp_addNF7Lbmi2:                     # @mcl_fp_addNF7Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	48(%rdx), %r9
	movq	40(%rdx), %rbp
	movq	32(%rdx), %r10
	movq	24(%rdx), %r11
	movq	16(%rdx), %r14
	movq	(%rdx), %r12
	movq	8(%rdx), %r15
	addq	(%rsi), %r12
	adcq	8(%rsi), %r15
	adcq	16(%rsi), %r14
	adcq	24(%rsi), %r11
	adcq	32(%rsi), %r10
	adcq	40(%rsi), %rbp
	movq	%rbp, -8(%rsp)          # 8-byte Spill
	adcq	48(%rsi), %r9
	movq	%r12, %rsi
	subq	(%rcx), %rsi
	movq	%r15, %rdx
	sbbq	8(%rcx), %rdx
	movq	%r14, %rax
	sbbq	16(%rcx), %rax
	movq	%r11, %rbx
	sbbq	24(%rcx), %rbx
	movq	%r10, %r13
	sbbq	32(%rcx), %r13
	sbbq	40(%rcx), %rbp
	movq	%r9, %r8
	sbbq	48(%rcx), %r8
	movq	%r8, %rcx
	sarq	$63, %rcx
	cmovsq	%r12, %rsi
	movq	%rsi, (%rdi)
	cmovsq	%r15, %rdx
	movq	%rdx, 8(%rdi)
	cmovsq	%r14, %rax
	movq	%rax, 16(%rdi)
	cmovsq	%r11, %rbx
	movq	%rbx, 24(%rdi)
	cmovsq	%r10, %r13
	movq	%r13, 32(%rdi)
	cmovsq	-8(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rbp, 40(%rdi)
	cmovsq	%r9, %r8
	movq	%r8, 48(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end105:
	.size	mcl_fp_addNF7Lbmi2, .Lfunc_end105-mcl_fp_addNF7Lbmi2

	.globl	mcl_fp_sub7Lbmi2
	.align	16, 0x90
	.type	mcl_fp_sub7Lbmi2,@function
mcl_fp_sub7Lbmi2:                       # @mcl_fp_sub7Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	48(%rdx), %r14
	movq	48(%rsi), %r8
	movq	40(%rdx), %r15
	movq	40(%rsi), %r9
	movq	32(%rdx), %r12
	movq	(%rsi), %rax
	movq	8(%rsi), %r11
	xorl	%ebx, %ebx
	subq	(%rdx), %rax
	sbbq	8(%rdx), %r11
	movq	16(%rsi), %r13
	sbbq	16(%rdx), %r13
	movq	32(%rsi), %r10
	movq	24(%rsi), %rsi
	sbbq	24(%rdx), %rsi
	movq	%rax, (%rdi)
	movq	%r11, 8(%rdi)
	movq	%r13, 16(%rdi)
	movq	%rsi, 24(%rdi)
	sbbq	%r12, %r10
	movq	%r10, 32(%rdi)
	sbbq	%r15, %r9
	movq	%r9, 40(%rdi)
	sbbq	%r14, %r8
	movq	%r8, 48(%rdi)
	sbbq	$0, %rbx
	testb	$1, %bl
	je	.LBB106_2
# BB#1:                                 # %carry
	movq	48(%rcx), %r14
	movq	40(%rcx), %r15
	movq	32(%rcx), %r12
	movq	24(%rcx), %rbx
	movq	8(%rcx), %rdx
	movq	16(%rcx), %rbp
	addq	(%rcx), %rax
	movq	%rax, (%rdi)
	adcq	%r11, %rdx
	movq	%rdx, 8(%rdi)
	adcq	%r13, %rbp
	movq	%rbp, 16(%rdi)
	adcq	%rsi, %rbx
	movq	%rbx, 24(%rdi)
	adcq	%r10, %r12
	movq	%r12, 32(%rdi)
	adcq	%r9, %r15
	movq	%r15, 40(%rdi)
	adcq	%r8, %r14
	movq	%r14, 48(%rdi)
.LBB106_2:                              # %nocarry
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end106:
	.size	mcl_fp_sub7Lbmi2, .Lfunc_end106-mcl_fp_sub7Lbmi2

	.globl	mcl_fp_subNF7Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subNF7Lbmi2,@function
mcl_fp_subNF7Lbmi2:                     # @mcl_fp_subNF7Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rcx, %r8
	movq	48(%rsi), %r12
	movq	40(%rsi), %rax
	movq	32(%rsi), %r9
	movq	24(%rsi), %r10
	movq	16(%rsi), %r11
	movq	(%rsi), %r14
	movq	8(%rsi), %r15
	subq	(%rdx), %r14
	sbbq	8(%rdx), %r15
	sbbq	16(%rdx), %r11
	sbbq	24(%rdx), %r10
	sbbq	32(%rdx), %r9
	sbbq	40(%rdx), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	sbbq	48(%rdx), %r12
	movq	%r12, %rax
	sarq	$63, %rax
	movq	%rax, %rsi
	shldq	$1, %r12, %rsi
	andq	(%r8), %rsi
	movq	48(%r8), %r13
	andq	%rax, %r13
	movq	40(%r8), %rbx
	andq	%rax, %rbx
	movq	32(%r8), %rdx
	andq	%rax, %rdx
	movq	24(%r8), %rbp
	andq	%rax, %rbp
	movq	16(%r8), %rcx
	andq	%rax, %rcx
	andq	8(%r8), %rax
	addq	%r14, %rsi
	adcq	%r15, %rax
	movq	%rsi, (%rdi)
	movq	%rax, 8(%rdi)
	adcq	%r11, %rcx
	movq	%rcx, 16(%rdi)
	adcq	%r10, %rbp
	movq	%rbp, 24(%rdi)
	adcq	%r9, %rdx
	movq	%rdx, 32(%rdi)
	adcq	-8(%rsp), %rbx          # 8-byte Folded Reload
	movq	%rbx, 40(%rdi)
	adcq	%r12, %r13
	movq	%r13, 48(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end107:
	.size	mcl_fp_subNF7Lbmi2, .Lfunc_end107-mcl_fp_subNF7Lbmi2

	.globl	mcl_fpDbl_add7Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_add7Lbmi2,@function
mcl_fpDbl_add7Lbmi2:                    # @mcl_fpDbl_add7Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rcx, %r8
	movq	104(%rdx), %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	movq	96(%rdx), %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	movq	88(%rdx), %r11
	movq	80(%rdx), %r14
	movq	24(%rsi), %r15
	movq	32(%rsi), %r12
	movq	16(%rdx), %r9
	movq	(%rdx), %rax
	movq	8(%rdx), %rbx
	addq	(%rsi), %rax
	adcq	8(%rsi), %rbx
	adcq	16(%rsi), %r9
	adcq	24(%rdx), %r15
	adcq	32(%rdx), %r12
	movq	72(%rdx), %r13
	movq	64(%rdx), %rbp
	movq	%rax, (%rdi)
	movq	56(%rdx), %r10
	movq	%rbx, 8(%rdi)
	movq	48(%rdx), %rcx
	movq	40(%rdx), %rdx
	movq	%r9, 16(%rdi)
	movq	104(%rsi), %r9
	movq	%r15, 24(%rdi)
	movq	40(%rsi), %rbx
	adcq	%rdx, %rbx
	movq	96(%rsi), %r15
	movq	%r12, 32(%rdi)
	movq	48(%rsi), %rdx
	adcq	%rcx, %rdx
	movq	88(%rsi), %rax
	movq	%rbx, 40(%rdi)
	movq	56(%rsi), %rcx
	adcq	%r10, %rcx
	movq	80(%rsi), %r12
	movq	%rdx, 48(%rdi)
	movq	72(%rsi), %rdx
	movq	64(%rsi), %rsi
	adcq	%rbp, %rsi
	adcq	%r13, %rdx
	adcq	%r14, %r12
	adcq	%r11, %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	adcq	-24(%rsp), %r15         # 8-byte Folded Reload
	movq	%r15, -24(%rsp)         # 8-byte Spill
	adcq	-16(%rsp), %r9          # 8-byte Folded Reload
	sbbq	%rbp, %rbp
	andl	$1, %ebp
	movq	%rcx, %rbx
	subq	(%r8), %rbx
	movq	%rsi, %r10
	sbbq	8(%r8), %r10
	movq	%rdx, %r11
	sbbq	16(%r8), %r11
	movq	%r12, %r14
	sbbq	24(%r8), %r14
	movq	-8(%rsp), %r13          # 8-byte Reload
	sbbq	32(%r8), %r13
	sbbq	40(%r8), %r15
	movq	%r9, %rax
	sbbq	48(%r8), %rax
	sbbq	$0, %rbp
	andl	$1, %ebp
	cmovneq	%rcx, %rbx
	movq	%rbx, 56(%rdi)
	testb	%bpl, %bpl
	cmovneq	%rsi, %r10
	movq	%r10, 64(%rdi)
	cmovneq	%rdx, %r11
	movq	%r11, 72(%rdi)
	cmovneq	%r12, %r14
	movq	%r14, 80(%rdi)
	cmovneq	-8(%rsp), %r13          # 8-byte Folded Reload
	movq	%r13, 88(%rdi)
	cmovneq	-24(%rsp), %r15         # 8-byte Folded Reload
	movq	%r15, 96(%rdi)
	cmovneq	%r9, %rax
	movq	%rax, 104(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end108:
	.size	mcl_fpDbl_add7Lbmi2, .Lfunc_end108-mcl_fpDbl_add7Lbmi2

	.globl	mcl_fpDbl_sub7Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sub7Lbmi2,@function
mcl_fpDbl_sub7Lbmi2:                    # @mcl_fpDbl_sub7Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rcx, %r8
	movq	104(%rdx), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	movq	96(%rdx), %r10
	movq	88(%rdx), %r14
	movq	16(%rsi), %rax
	movq	(%rsi), %r15
	movq	8(%rsi), %r11
	xorl	%ecx, %ecx
	subq	(%rdx), %r15
	sbbq	8(%rdx), %r11
	sbbq	16(%rdx), %rax
	movq	24(%rsi), %rbx
	sbbq	24(%rdx), %rbx
	movq	32(%rsi), %r12
	sbbq	32(%rdx), %r12
	movq	80(%rdx), %r13
	movq	72(%rdx), %rbp
	movq	%r15, (%rdi)
	movq	64(%rdx), %r9
	movq	%r11, 8(%rdi)
	movq	56(%rdx), %r15
	movq	%rax, 16(%rdi)
	movq	48(%rdx), %r11
	movq	40(%rdx), %rdx
	movq	%rbx, 24(%rdi)
	movq	40(%rsi), %rbx
	sbbq	%rdx, %rbx
	movq	104(%rsi), %rax
	movq	%r12, 32(%rdi)
	movq	48(%rsi), %r12
	sbbq	%r11, %r12
	movq	96(%rsi), %r11
	movq	%rbx, 40(%rdi)
	movq	56(%rsi), %rdx
	sbbq	%r15, %rdx
	movq	88(%rsi), %r15
	movq	%r12, 48(%rdi)
	movq	64(%rsi), %rbx
	sbbq	%r9, %rbx
	movq	80(%rsi), %r12
	movq	72(%rsi), %r9
	sbbq	%rbp, %r9
	sbbq	%r13, %r12
	sbbq	%r14, %r15
	sbbq	%r10, %r11
	sbbq	-8(%rsp), %rax          # 8-byte Folded Reload
	movq	%rax, -8(%rsp)          # 8-byte Spill
	movl	$0, %ebp
	sbbq	$0, %rbp
	andl	$1, %ebp
	movq	(%r8), %r10
	cmoveq	%rcx, %r10
	testb	%bpl, %bpl
	movq	16(%r8), %rbp
	cmoveq	%rcx, %rbp
	movq	8(%r8), %rsi
	cmoveq	%rcx, %rsi
	movq	48(%r8), %r14
	cmoveq	%rcx, %r14
	movq	40(%r8), %r13
	cmoveq	%rcx, %r13
	movq	32(%r8), %rax
	cmoveq	%rcx, %rax
	cmovneq	24(%r8), %rcx
	addq	%rdx, %r10
	adcq	%rbx, %rsi
	movq	%r10, 56(%rdi)
	movq	%rsi, 64(%rdi)
	adcq	%r9, %rbp
	movq	%rbp, 72(%rdi)
	adcq	%r12, %rcx
	movq	%rcx, 80(%rdi)
	adcq	%r15, %rax
	movq	%rax, 88(%rdi)
	adcq	%r11, %r13
	movq	%r13, 96(%rdi)
	adcq	-8(%rsp), %r14          # 8-byte Folded Reload
	movq	%r14, 104(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end109:
	.size	mcl_fpDbl_sub7Lbmi2, .Lfunc_end109-mcl_fpDbl_sub7Lbmi2

	.align	16, 0x90
	.type	.LmulPv512x64,@function
.LmulPv512x64:                          # @mulPv512x64
# BB#0:
	mulxq	(%rsi), %rcx, %rax
	movq	%rcx, (%rdi)
	mulxq	8(%rsi), %rcx, %r8
	addq	%rax, %rcx
	movq	%rcx, 8(%rdi)
	mulxq	16(%rsi), %rcx, %r9
	adcq	%r8, %rcx
	movq	%rcx, 16(%rdi)
	mulxq	24(%rsi), %rax, %rcx
	adcq	%r9, %rax
	movq	%rax, 24(%rdi)
	mulxq	32(%rsi), %rax, %r8
	adcq	%rcx, %rax
	movq	%rax, 32(%rdi)
	mulxq	40(%rsi), %rcx, %r9
	adcq	%r8, %rcx
	movq	%rcx, 40(%rdi)
	mulxq	48(%rsi), %rax, %rcx
	adcq	%r9, %rax
	movq	%rax, 48(%rdi)
	mulxq	56(%rsi), %rax, %rdx
	adcq	%rcx, %rax
	movq	%rax, 56(%rdi)
	adcq	$0, %rdx
	movq	%rdx, 64(%rdi)
	movq	%rdi, %rax
	retq
.Lfunc_end110:
	.size	.LmulPv512x64, .Lfunc_end110-.LmulPv512x64

	.globl	mcl_fp_mulUnitPre8Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mulUnitPre8Lbmi2,@function
mcl_fp_mulUnitPre8Lbmi2:                # @mcl_fp_mulUnitPre8Lbmi2
# BB#0:
	pushq	%rbx
	subq	$80, %rsp
	movq	%rdi, %rbx
	leaq	8(%rsp), %rdi
	callq	.LmulPv512x64
	movq	72(%rsp), %r8
	movq	64(%rsp), %r9
	movq	56(%rsp), %r10
	movq	48(%rsp), %r11
	movq	40(%rsp), %rdi
	movq	32(%rsp), %rax
	movq	24(%rsp), %rcx
	movq	8(%rsp), %rdx
	movq	16(%rsp), %rsi
	movq	%rdx, (%rbx)
	movq	%rsi, 8(%rbx)
	movq	%rcx, 16(%rbx)
	movq	%rax, 24(%rbx)
	movq	%rdi, 32(%rbx)
	movq	%r11, 40(%rbx)
	movq	%r10, 48(%rbx)
	movq	%r9, 56(%rbx)
	movq	%r8, 64(%rbx)
	addq	$80, %rsp
	popq	%rbx
	retq
.Lfunc_end111:
	.size	mcl_fp_mulUnitPre8Lbmi2, .Lfunc_end111-mcl_fp_mulUnitPre8Lbmi2

	.globl	mcl_fpDbl_mulPre8Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_mulPre8Lbmi2,@function
mcl_fpDbl_mulPre8Lbmi2:                 # @mcl_fpDbl_mulPre8Lbmi2
# BB#0:
	pushq	%rbp
	movq	%rsp, %rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	subq	$200, %rsp
	movq	%rdx, %rbx
	movq	%rsi, %r15
	movq	%rdi, %r14
	callq	mcl_fpDbl_mulPre4Lbmi2@PLT
	leaq	64(%r14), %rdi
	leaq	32(%r15), %rsi
	leaq	32(%rbx), %rdx
	callq	mcl_fpDbl_mulPre4Lbmi2@PLT
	movq	56(%rbx), %r10
	movq	48(%rbx), %rcx
	movq	(%rbx), %rdx
	movq	8(%rbx), %rsi
	addq	32(%rbx), %rdx
	adcq	40(%rbx), %rsi
	adcq	16(%rbx), %rcx
	adcq	24(%rbx), %r10
	pushfq
	popq	%r8
	xorl	%r9d, %r9d
	movq	56(%r15), %rdi
	movq	48(%r15), %r13
	movq	(%r15), %r12
	movq	8(%r15), %rbx
	addq	32(%r15), %r12
	adcq	40(%r15), %rbx
	adcq	16(%r15), %r13
	adcq	24(%r15), %rdi
	movl	$0, %eax
	cmovbq	%r10, %rax
	movq	%rax, -176(%rbp)        # 8-byte Spill
	movl	$0, %eax
	cmovbq	%rcx, %rax
	movq	%rax, -184(%rbp)        # 8-byte Spill
	movl	$0, %eax
	cmovbq	%rsi, %rax
	movq	%rax, -192(%rbp)        # 8-byte Spill
	movl	$0, %eax
	cmovbq	%rdx, %rax
	movq	%rax, -200(%rbp)        # 8-byte Spill
	sbbq	%r15, %r15
	movq	%r12, -136(%rbp)
	movq	%rbx, -128(%rbp)
	movq	%r13, -120(%rbp)
	movq	%rdi, -112(%rbp)
	movq	%rdx, -168(%rbp)
	movq	%rsi, -160(%rbp)
	movq	%rcx, -152(%rbp)
	movq	%r10, -144(%rbp)
	pushq	%r8
	popfq
	cmovaeq	%r9, %rdi
	movq	%rdi, -216(%rbp)        # 8-byte Spill
	cmovaeq	%r9, %r13
	cmovaeq	%r9, %rbx
	cmovaeq	%r9, %r12
	sbbq	%rax, %rax
	movq	%rax, -208(%rbp)        # 8-byte Spill
	leaq	-104(%rbp), %rdi
	leaq	-136(%rbp), %rsi
	leaq	-168(%rbp), %rdx
	callq	mcl_fpDbl_mulPre4Lbmi2@PLT
	addq	-200(%rbp), %r12        # 8-byte Folded Reload
	adcq	-192(%rbp), %rbx        # 8-byte Folded Reload
	adcq	-184(%rbp), %r13        # 8-byte Folded Reload
	movq	-216(%rbp), %r10        # 8-byte Reload
	adcq	-176(%rbp), %r10        # 8-byte Folded Reload
	sbbq	%rax, %rax
	andl	$1, %eax
	movq	-208(%rbp), %rdx        # 8-byte Reload
	andl	%edx, %r15d
	andl	$1, %r15d
	addq	-72(%rbp), %r12
	adcq	-64(%rbp), %rbx
	adcq	-56(%rbp), %r13
	adcq	-48(%rbp), %r10
	adcq	%rax, %r15
	movq	-80(%rbp), %rax
	movq	-88(%rbp), %rcx
	movq	-104(%rbp), %rsi
	movq	-96(%rbp), %rdx
	subq	(%r14), %rsi
	sbbq	8(%r14), %rdx
	sbbq	16(%r14), %rcx
	sbbq	24(%r14), %rax
	movq	32(%r14), %rdi
	movq	%rdi, -184(%rbp)        # 8-byte Spill
	movq	40(%r14), %r8
	movq	%r8, -176(%rbp)         # 8-byte Spill
	sbbq	%rdi, %r12
	sbbq	%r8, %rbx
	movq	48(%r14), %rdi
	movq	%rdi, -192(%rbp)        # 8-byte Spill
	sbbq	%rdi, %r13
	movq	56(%r14), %rdi
	movq	%rdi, -200(%rbp)        # 8-byte Spill
	sbbq	%rdi, %r10
	sbbq	$0, %r15
	movq	64(%r14), %r11
	subq	%r11, %rsi
	movq	72(%r14), %rdi
	movq	%rdi, -208(%rbp)        # 8-byte Spill
	sbbq	%rdi, %rdx
	movq	80(%r14), %rdi
	movq	%rdi, -216(%rbp)        # 8-byte Spill
	sbbq	%rdi, %rcx
	movq	88(%r14), %rdi
	movq	%rdi, -224(%rbp)        # 8-byte Spill
	sbbq	%rdi, %rax
	movq	96(%r14), %rdi
	movq	%rdi, -232(%rbp)        # 8-byte Spill
	sbbq	%rdi, %r12
	movq	104(%r14), %rdi
	sbbq	%rdi, %rbx
	movq	112(%r14), %r8
	sbbq	%r8, %r13
	movq	120(%r14), %r9
	sbbq	%r9, %r10
	sbbq	$0, %r15
	addq	-184(%rbp), %rsi        # 8-byte Folded Reload
	adcq	-176(%rbp), %rdx        # 8-byte Folded Reload
	movq	%rsi, 32(%r14)
	adcq	-192(%rbp), %rcx        # 8-byte Folded Reload
	movq	%rdx, 40(%r14)
	adcq	-200(%rbp), %rax        # 8-byte Folded Reload
	movq	%rcx, 48(%r14)
	adcq	%r11, %r12
	movq	%rax, 56(%r14)
	movq	%r12, 64(%r14)
	adcq	-208(%rbp), %rbx        # 8-byte Folded Reload
	movq	%rbx, 72(%r14)
	adcq	-216(%rbp), %r13        # 8-byte Folded Reload
	movq	%r13, 80(%r14)
	adcq	-224(%rbp), %r10        # 8-byte Folded Reload
	movq	%r10, 88(%r14)
	adcq	-232(%rbp), %r15        # 8-byte Folded Reload
	movq	%r15, 96(%r14)
	adcq	$0, %rdi
	movq	%rdi, 104(%r14)
	adcq	$0, %r8
	movq	%r8, 112(%r14)
	adcq	$0, %r9
	movq	%r9, 120(%r14)
	addq	$200, %rsp
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end112:
	.size	mcl_fpDbl_mulPre8Lbmi2, .Lfunc_end112-mcl_fpDbl_mulPre8Lbmi2

	.globl	mcl_fpDbl_sqrPre8Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sqrPre8Lbmi2,@function
mcl_fpDbl_sqrPre8Lbmi2:                 # @mcl_fpDbl_sqrPre8Lbmi2
# BB#0:
	pushq	%rbp
	movq	%rsp, %rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	subq	$200, %rsp
	movq	%rsi, %r14
	movq	%rdi, %rbx
	movq	%r14, %rdx
	callq	mcl_fpDbl_mulPre4Lbmi2@PLT
	leaq	64(%rbx), %rdi
	leaq	32(%r14), %rsi
	movq	%rsi, %rdx
	callq	mcl_fpDbl_mulPre4Lbmi2@PLT
	movq	(%r14), %r12
	movq	8(%r14), %r15
	addq	32(%r14), %r12
	adcq	40(%r14), %r15
	pushfq
	popq	%rax
	movq	%r12, -136(%rbp)
	movq	%r12, -168(%rbp)
	addq	%r12, %r12
	movq	%r15, -128(%rbp)
	movq	%r15, -160(%rbp)
	adcq	%r15, %r15
	pushfq
	popq	%rcx
	movq	56(%r14), %r13
	movq	48(%r14), %rdx
	pushq	%rax
	popfq
	adcq	16(%r14), %rdx
	adcq	24(%r14), %r13
	pushfq
	popq	%r8
	pushfq
	popq	%rsi
	pushfq
	popq	%rdi
	sbbq	%rax, %rax
	movq	%rax, -184(%rbp)        # 8-byte Spill
	xorl	%eax, %eax
	pushq	%rdi
	popfq
	cmovaeq	%rax, %r15
	movq	%r15, -176(%rbp)        # 8-byte Spill
	cmovaeq	%rax, %r12
	movq	%rdx, -120(%rbp)
	movq	%rdx, -152(%rbp)
	movq	%rdx, %r15
	pushq	%rcx
	popfq
	adcq	%r15, %r15
	movq	%r13, %r14
	movq	%r13, -112(%rbp)
	movq	%r13, -144(%rbp)
	adcq	%r13, %r13
	pushq	%rsi
	popfq
	cmovaeq	%rax, %r13
	cmovaeq	%rax, %r15
	shrq	$63, %r14
	pushq	%r8
	popfq
	cmovaeq	%rax, %r14
	leaq	-104(%rbp), %rdi
	leaq	-136(%rbp), %rsi
	leaq	-168(%rbp), %rdx
	callq	mcl_fpDbl_mulPre4Lbmi2@PLT
	movq	-184(%rbp), %rax        # 8-byte Reload
	andl	$1, %eax
	addq	-72(%rbp), %r12
	movq	-176(%rbp), %r8         # 8-byte Reload
	adcq	-64(%rbp), %r8
	adcq	-56(%rbp), %r15
	adcq	-48(%rbp), %r13
	adcq	%r14, %rax
	movq	%rax, %rdi
	movq	-80(%rbp), %rax
	movq	-88(%rbp), %rcx
	movq	-104(%rbp), %rsi
	movq	-96(%rbp), %rdx
	subq	(%rbx), %rsi
	sbbq	8(%rbx), %rdx
	sbbq	16(%rbx), %rcx
	sbbq	24(%rbx), %rax
	movq	32(%rbx), %r10
	movq	%r10, -184(%rbp)        # 8-byte Spill
	movq	40(%rbx), %r9
	movq	%r9, -176(%rbp)         # 8-byte Spill
	sbbq	%r10, %r12
	sbbq	%r9, %r8
	movq	%r8, %r10
	movq	48(%rbx), %r8
	movq	%r8, -192(%rbp)         # 8-byte Spill
	sbbq	%r8, %r15
	movq	56(%rbx), %r8
	movq	%r8, -200(%rbp)         # 8-byte Spill
	sbbq	%r8, %r13
	sbbq	$0, %rdi
	movq	64(%rbx), %r11
	subq	%r11, %rsi
	movq	72(%rbx), %r8
	movq	%r8, -208(%rbp)         # 8-byte Spill
	sbbq	%r8, %rdx
	movq	80(%rbx), %r8
	movq	%r8, -216(%rbp)         # 8-byte Spill
	sbbq	%r8, %rcx
	movq	88(%rbx), %r8
	movq	%r8, -224(%rbp)         # 8-byte Spill
	sbbq	%r8, %rax
	movq	96(%rbx), %r8
	movq	%r8, -232(%rbp)         # 8-byte Spill
	sbbq	%r8, %r12
	movq	104(%rbx), %r14
	sbbq	%r14, %r10
	movq	112(%rbx), %r8
	sbbq	%r8, %r15
	movq	120(%rbx), %r9
	sbbq	%r9, %r13
	sbbq	$0, %rdi
	addq	-184(%rbp), %rsi        # 8-byte Folded Reload
	adcq	-176(%rbp), %rdx        # 8-byte Folded Reload
	movq	%rsi, 32(%rbx)
	adcq	-192(%rbp), %rcx        # 8-byte Folded Reload
	movq	%rdx, 40(%rbx)
	adcq	-200(%rbp), %rax        # 8-byte Folded Reload
	movq	%rcx, 48(%rbx)
	adcq	%r11, %r12
	movq	%rax, 56(%rbx)
	movq	%r12, 64(%rbx)
	adcq	-208(%rbp), %r10        # 8-byte Folded Reload
	movq	%r10, 72(%rbx)
	adcq	-216(%rbp), %r15        # 8-byte Folded Reload
	movq	%r15, 80(%rbx)
	adcq	-224(%rbp), %r13        # 8-byte Folded Reload
	movq	%r13, 88(%rbx)
	adcq	-232(%rbp), %rdi        # 8-byte Folded Reload
	movq	%rdi, 96(%rbx)
	adcq	$0, %r14
	movq	%r14, 104(%rbx)
	adcq	$0, %r8
	movq	%r8, 112(%rbx)
	adcq	$0, %r9
	movq	%r9, 120(%rbx)
	addq	$200, %rsp
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end113:
	.size	mcl_fpDbl_sqrPre8Lbmi2, .Lfunc_end113-mcl_fpDbl_sqrPre8Lbmi2

	.globl	mcl_fp_mont8Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mont8Lbmi2,@function
mcl_fp_mont8Lbmi2:                      # @mcl_fp_mont8Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	subq	$1256, %rsp             # imm = 0x4E8
	movq	%rcx, %r13
	movq	%r13, 40(%rsp)          # 8-byte Spill
	movq	%rdx, 16(%rsp)          # 8-byte Spill
	movq	%rsi, 24(%rsp)          # 8-byte Spill
	movq	%rdi, (%rsp)            # 8-byte Spill
	movq	-8(%r13), %rbx
	movq	%rbx, 32(%rsp)          # 8-byte Spill
	movq	(%rdx), %rdx
	leaq	1184(%rsp), %rdi
	callq	.LmulPv512x64
	movq	1184(%rsp), %r15
	movq	1192(%rsp), %r14
	movq	%r15, %rdx
	imulq	%rbx, %rdx
	movq	1248(%rsp), %rax
	movq	%rax, 96(%rsp)          # 8-byte Spill
	movq	1240(%rsp), %rax
	movq	%rax, 80(%rsp)          # 8-byte Spill
	movq	1232(%rsp), %rax
	movq	%rax, 88(%rsp)          # 8-byte Spill
	movq	1224(%rsp), %rax
	movq	%rax, 72(%rsp)          # 8-byte Spill
	movq	1216(%rsp), %r12
	movq	1208(%rsp), %rbx
	movq	1200(%rsp), %rbp
	leaq	1112(%rsp), %rdi
	movq	%r13, %rsi
	callq	.LmulPv512x64
	addq	1112(%rsp), %r15
	adcq	1120(%rsp), %r14
	adcq	1128(%rsp), %rbp
	movq	%rbp, 8(%rsp)           # 8-byte Spill
	adcq	1136(%rsp), %rbx
	movq	%rbx, 48(%rsp)          # 8-byte Spill
	adcq	1144(%rsp), %r12
	movq	%r12, 64(%rsp)          # 8-byte Spill
	movq	72(%rsp), %r13          # 8-byte Reload
	adcq	1152(%rsp), %r13
	movq	88(%rsp), %rbx          # 8-byte Reload
	adcq	1160(%rsp), %rbx
	movq	80(%rsp), %rbp          # 8-byte Reload
	adcq	1168(%rsp), %rbp
	movq	96(%rsp), %rax          # 8-byte Reload
	adcq	1176(%rsp), %rax
	movq	%rax, 96(%rsp)          # 8-byte Spill
	sbbq	%r15, %r15
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	8(%rax), %rdx
	leaq	1040(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	andl	$1, %r15d
	addq	1040(%rsp), %r14
	movq	8(%rsp), %rax           # 8-byte Reload
	adcq	1048(%rsp), %rax
	movq	%rax, 8(%rsp)           # 8-byte Spill
	movq	48(%rsp), %rax          # 8-byte Reload
	adcq	1056(%rsp), %rax
	movq	%rax, %r12
	movq	64(%rsp), %rax          # 8-byte Reload
	adcq	1064(%rsp), %rax
	movq	%rax, 64(%rsp)          # 8-byte Spill
	adcq	1072(%rsp), %r13
	movq	%r13, 72(%rsp)          # 8-byte Spill
	adcq	1080(%rsp), %rbx
	movq	%rbx, 88(%rsp)          # 8-byte Spill
	adcq	1088(%rsp), %rbp
	movq	96(%rsp), %rax          # 8-byte Reload
	adcq	1096(%rsp), %rax
	movq	%rax, 96(%rsp)          # 8-byte Spill
	adcq	1104(%rsp), %r15
	movq	%r15, 56(%rsp)          # 8-byte Spill
	sbbq	%r15, %r15
	movq	%r14, %rdx
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	968(%rsp), %rdi
	movq	40(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	andl	$1, %r15d
	addq	968(%rsp), %r14
	movq	8(%rsp), %r13           # 8-byte Reload
	adcq	976(%rsp), %r13
	adcq	984(%rsp), %r12
	movq	%r12, 48(%rsp)          # 8-byte Spill
	movq	64(%rsp), %r14          # 8-byte Reload
	adcq	992(%rsp), %r14
	movq	72(%rsp), %rbx          # 8-byte Reload
	adcq	1000(%rsp), %rbx
	movq	88(%rsp), %rax          # 8-byte Reload
	adcq	1008(%rsp), %rax
	movq	%rax, 88(%rsp)          # 8-byte Spill
	adcq	1016(%rsp), %rbp
	movq	%rbp, %r12
	movq	96(%rsp), %rbp          # 8-byte Reload
	adcq	1024(%rsp), %rbp
	movq	56(%rsp), %rax          # 8-byte Reload
	adcq	1032(%rsp), %rax
	movq	%rax, 56(%rsp)          # 8-byte Spill
	adcq	$0, %r15
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	16(%rax), %rdx
	leaq	896(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	movq	%r13, %rcx
	addq	896(%rsp), %rcx
	movq	48(%rsp), %r13          # 8-byte Reload
	adcq	904(%rsp), %r13
	adcq	912(%rsp), %r14
	adcq	920(%rsp), %rbx
	movq	%rbx, 72(%rsp)          # 8-byte Spill
	movq	88(%rsp), %rax          # 8-byte Reload
	adcq	928(%rsp), %rax
	movq	%rax, 88(%rsp)          # 8-byte Spill
	adcq	936(%rsp), %r12
	movq	%r12, 80(%rsp)          # 8-byte Spill
	adcq	944(%rsp), %rbp
	movq	%rbp, 96(%rsp)          # 8-byte Spill
	movq	56(%rsp), %r12          # 8-byte Reload
	adcq	952(%rsp), %r12
	adcq	960(%rsp), %r15
	sbbq	%rbx, %rbx
	movq	%rcx, %rdx
	movq	%rcx, %rbp
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	824(%rsp), %rdi
	movq	40(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	andl	$1, %ebx
	addq	824(%rsp), %rbp
	adcq	832(%rsp), %r13
	movq	%r13, 48(%rsp)          # 8-byte Spill
	adcq	840(%rsp), %r14
	movq	%r14, 64(%rsp)          # 8-byte Spill
	movq	72(%rsp), %r13          # 8-byte Reload
	adcq	848(%rsp), %r13
	movq	88(%rsp), %rbp          # 8-byte Reload
	adcq	856(%rsp), %rbp
	movq	80(%rsp), %r14          # 8-byte Reload
	adcq	864(%rsp), %r14
	movq	96(%rsp), %rax          # 8-byte Reload
	adcq	872(%rsp), %rax
	movq	%rax, 96(%rsp)          # 8-byte Spill
	adcq	880(%rsp), %r12
	adcq	888(%rsp), %r15
	adcq	$0, %rbx
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	24(%rax), %rdx
	leaq	752(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	movq	48(%rsp), %rax          # 8-byte Reload
	addq	752(%rsp), %rax
	movq	64(%rsp), %rcx          # 8-byte Reload
	adcq	760(%rsp), %rcx
	movq	%rcx, 64(%rsp)          # 8-byte Spill
	adcq	768(%rsp), %r13
	movq	%r13, 72(%rsp)          # 8-byte Spill
	adcq	776(%rsp), %rbp
	movq	%rbp, 88(%rsp)          # 8-byte Spill
	adcq	784(%rsp), %r14
	movq	%r14, 80(%rsp)          # 8-byte Spill
	movq	96(%rsp), %rbp          # 8-byte Reload
	adcq	792(%rsp), %rbp
	adcq	800(%rsp), %r12
	adcq	808(%rsp), %r15
	adcq	816(%rsp), %rbx
	movq	%rbx, 48(%rsp)          # 8-byte Spill
	sbbq	%r13, %r13
	movq	%rax, %rdx
	movq	%rax, %rbx
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	680(%rsp), %rdi
	movq	40(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	movq	%r13, %rax
	andl	$1, %eax
	addq	680(%rsp), %rbx
	movq	64(%rsp), %r14          # 8-byte Reload
	adcq	688(%rsp), %r14
	movq	72(%rsp), %rcx          # 8-byte Reload
	adcq	696(%rsp), %rcx
	movq	%rcx, 72(%rsp)          # 8-byte Spill
	movq	88(%rsp), %r13          # 8-byte Reload
	adcq	704(%rsp), %r13
	movq	80(%rsp), %rbx          # 8-byte Reload
	adcq	712(%rsp), %rbx
	adcq	720(%rsp), %rbp
	movq	%rbp, 96(%rsp)          # 8-byte Spill
	movq	%r12, %rbp
	adcq	728(%rsp), %rbp
	adcq	736(%rsp), %r15
	movq	48(%rsp), %r12          # 8-byte Reload
	adcq	744(%rsp), %r12
	adcq	$0, %rax
	movq	%rax, 64(%rsp)          # 8-byte Spill
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	32(%rax), %rdx
	leaq	608(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	movq	%r14, %rax
	addq	608(%rsp), %rax
	movq	72(%rsp), %r14          # 8-byte Reload
	adcq	616(%rsp), %r14
	adcq	624(%rsp), %r13
	movq	%r13, 88(%rsp)          # 8-byte Spill
	adcq	632(%rsp), %rbx
	movq	%rbx, %r13
	movq	96(%rsp), %rcx          # 8-byte Reload
	adcq	640(%rsp), %rcx
	movq	%rcx, 96(%rsp)          # 8-byte Spill
	adcq	648(%rsp), %rbp
	movq	%rbp, 56(%rsp)          # 8-byte Spill
	adcq	656(%rsp), %r15
	adcq	664(%rsp), %r12
	movq	%r12, 48(%rsp)          # 8-byte Spill
	movq	64(%rsp), %rcx          # 8-byte Reload
	adcq	672(%rsp), %rcx
	movq	%rcx, 64(%rsp)          # 8-byte Spill
	sbbq	%rbp, %rbp
	movq	%rax, %rdx
	movq	%rax, %rbx
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	536(%rsp), %rdi
	movq	40(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	movq	%rbp, %rax
	andl	$1, %eax
	addq	536(%rsp), %rbx
	adcq	544(%rsp), %r14
	movq	%r14, 72(%rsp)          # 8-byte Spill
	movq	88(%rsp), %rbx          # 8-byte Reload
	adcq	552(%rsp), %rbx
	adcq	560(%rsp), %r13
	movq	96(%rsp), %rbp          # 8-byte Reload
	adcq	568(%rsp), %rbp
	movq	56(%rsp), %r12          # 8-byte Reload
	adcq	576(%rsp), %r12
	adcq	584(%rsp), %r15
	movq	48(%rsp), %rcx          # 8-byte Reload
	adcq	592(%rsp), %rcx
	movq	%rcx, 48(%rsp)          # 8-byte Spill
	movq	64(%rsp), %r14          # 8-byte Reload
	adcq	600(%rsp), %r14
	adcq	$0, %rax
	movq	%rax, 88(%rsp)          # 8-byte Spill
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	40(%rax), %rdx
	leaq	464(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	movq	72(%rsp), %rax          # 8-byte Reload
	addq	464(%rsp), %rax
	adcq	472(%rsp), %rbx
	adcq	480(%rsp), %r13
	movq	%r13, 80(%rsp)          # 8-byte Spill
	adcq	488(%rsp), %rbp
	movq	%rbp, 96(%rsp)          # 8-byte Spill
	adcq	496(%rsp), %r12
	adcq	504(%rsp), %r15
	movq	%r15, 72(%rsp)          # 8-byte Spill
	movq	48(%rsp), %r15          # 8-byte Reload
	adcq	512(%rsp), %r15
	adcq	520(%rsp), %r14
	movq	%r14, 64(%rsp)          # 8-byte Spill
	movq	88(%rsp), %r14          # 8-byte Reload
	adcq	528(%rsp), %r14
	sbbq	%r13, %r13
	movq	%rax, %rdx
	movq	%rax, %rbp
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	392(%rsp), %rdi
	movq	40(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	movq	%r13, %rax
	andl	$1, %eax
	addq	392(%rsp), %rbp
	adcq	400(%rsp), %rbx
	movq	%rbx, 88(%rsp)          # 8-byte Spill
	movq	80(%rsp), %rbp          # 8-byte Reload
	adcq	408(%rsp), %rbp
	movq	96(%rsp), %rbx          # 8-byte Reload
	adcq	416(%rsp), %rbx
	adcq	424(%rsp), %r12
	movq	72(%rsp), %r13          # 8-byte Reload
	adcq	432(%rsp), %r13
	adcq	440(%rsp), %r15
	movq	%r15, 48(%rsp)          # 8-byte Spill
	movq	64(%rsp), %r15          # 8-byte Reload
	adcq	448(%rsp), %r15
	adcq	456(%rsp), %r14
	adcq	$0, %rax
	movq	%rax, 72(%rsp)          # 8-byte Spill
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	48(%rax), %rdx
	leaq	320(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	movq	88(%rsp), %rax          # 8-byte Reload
	addq	320(%rsp), %rax
	adcq	328(%rsp), %rbp
	movq	%rbp, 80(%rsp)          # 8-byte Spill
	adcq	336(%rsp), %rbx
	movq	%rbx, 96(%rsp)          # 8-byte Spill
	movq	%r12, %rbp
	adcq	344(%rsp), %rbp
	adcq	352(%rsp), %r13
	movq	48(%rsp), %r12          # 8-byte Reload
	adcq	360(%rsp), %r12
	adcq	368(%rsp), %r15
	movq	%r15, 64(%rsp)          # 8-byte Spill
	adcq	376(%rsp), %r14
	movq	%r14, 88(%rsp)          # 8-byte Spill
	movq	72(%rsp), %rcx          # 8-byte Reload
	adcq	384(%rsp), %rcx
	movq	%rcx, 72(%rsp)          # 8-byte Spill
	sbbq	%r15, %r15
	movq	%rax, %rdx
	movq	%rax, %rbx
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	248(%rsp), %rdi
	movq	40(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	andl	$1, %r15d
	addq	248(%rsp), %rbx
	movq	80(%rsp), %rax          # 8-byte Reload
	adcq	256(%rsp), %rax
	movq	%rax, 80(%rsp)          # 8-byte Spill
	movq	96(%rsp), %r14          # 8-byte Reload
	adcq	264(%rsp), %r14
	adcq	272(%rsp), %rbp
	movq	%rbp, 56(%rsp)          # 8-byte Spill
	movq	%r13, %rbx
	adcq	280(%rsp), %rbx
	movq	%r12, %rbp
	adcq	288(%rsp), %rbp
	movq	64(%rsp), %r13          # 8-byte Reload
	adcq	296(%rsp), %r13
	movq	88(%rsp), %rax          # 8-byte Reload
	adcq	304(%rsp), %rax
	movq	%rax, 88(%rsp)          # 8-byte Spill
	movq	72(%rsp), %r12          # 8-byte Reload
	adcq	312(%rsp), %r12
	adcq	$0, %r15
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	56(%rax), %rdx
	leaq	176(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	movq	80(%rsp), %rax          # 8-byte Reload
	addq	176(%rsp), %rax
	adcq	184(%rsp), %r14
	movq	%r14, 96(%rsp)          # 8-byte Spill
	movq	56(%rsp), %rcx          # 8-byte Reload
	adcq	192(%rsp), %rcx
	movq	%rcx, 56(%rsp)          # 8-byte Spill
	adcq	200(%rsp), %rbx
	movq	%rbx, 72(%rsp)          # 8-byte Spill
	adcq	208(%rsp), %rbp
	adcq	216(%rsp), %r13
	movq	%r13, 64(%rsp)          # 8-byte Spill
	movq	88(%rsp), %r14          # 8-byte Reload
	adcq	224(%rsp), %r14
	adcq	232(%rsp), %r12
	adcq	240(%rsp), %r15
	sbbq	%rbx, %rbx
	movq	32(%rsp), %rdx          # 8-byte Reload
	imulq	%rax, %rdx
	movq	%rax, %r13
	leaq	104(%rsp), %rdi
	movq	40(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	andl	$1, %ebx
	addq	104(%rsp), %r13
	movq	96(%rsp), %rcx          # 8-byte Reload
	adcq	112(%rsp), %rcx
	movq	56(%rsp), %rdx          # 8-byte Reload
	adcq	120(%rsp), %rdx
	movq	72(%rsp), %rsi          # 8-byte Reload
	adcq	128(%rsp), %rsi
	movq	%rbp, %rdi
	adcq	136(%rsp), %rdi
	movq	%rdi, 48(%rsp)          # 8-byte Spill
	movq	64(%rsp), %r8           # 8-byte Reload
	adcq	144(%rsp), %r8
	movq	%r8, 64(%rsp)           # 8-byte Spill
	movq	%r14, %r9
	adcq	152(%rsp), %r9
	movq	%r9, 88(%rsp)           # 8-byte Spill
	adcq	160(%rsp), %r12
	adcq	168(%rsp), %r15
	adcq	$0, %rbx
	movq	%rcx, %rax
	movq	%rcx, %r11
	movq	40(%rsp), %rbp          # 8-byte Reload
	subq	(%rbp), %rax
	movq	%rdx, %rcx
	movq	%rdx, %r14
	sbbq	8(%rbp), %rcx
	movq	%rsi, %rdx
	movq	%rsi, %r13
	sbbq	16(%rbp), %rdx
	movq	%rdi, %rsi
	sbbq	24(%rbp), %rsi
	movq	%r8, %rdi
	sbbq	32(%rbp), %rdi
	movq	%r9, %r10
	sbbq	40(%rbp), %r10
	movq	%r12, %r8
	sbbq	48(%rbp), %r8
	movq	%r15, %r9
	sbbq	56(%rbp), %r9
	sbbq	$0, %rbx
	andl	$1, %ebx
	cmovneq	%r15, %r9
	testb	%bl, %bl
	cmovneq	%r11, %rax
	movq	(%rsp), %rbx            # 8-byte Reload
	movq	%rax, (%rbx)
	cmovneq	%r14, %rcx
	movq	%rcx, 8(%rbx)
	cmovneq	%r13, %rdx
	movq	%rdx, 16(%rbx)
	cmovneq	48(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 24(%rbx)
	cmovneq	64(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 32(%rbx)
	cmovneq	88(%rsp), %r10          # 8-byte Folded Reload
	movq	%r10, 40(%rbx)
	cmovneq	%r12, %r8
	movq	%r8, 48(%rbx)
	movq	%r9, 56(%rbx)
	addq	$1256, %rsp             # imm = 0x4E8
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end114:
	.size	mcl_fp_mont8Lbmi2, .Lfunc_end114-mcl_fp_mont8Lbmi2

	.globl	mcl_fp_montNF8Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montNF8Lbmi2,@function
mcl_fp_montNF8Lbmi2:                    # @mcl_fp_montNF8Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	subq	$1240, %rsp             # imm = 0x4D8
	movq	%rcx, 56(%rsp)          # 8-byte Spill
	movq	%rdx, 16(%rsp)          # 8-byte Spill
	movq	%rsi, 24(%rsp)          # 8-byte Spill
	movq	%rdi, (%rsp)            # 8-byte Spill
	movq	-8(%rcx), %rbx
	movq	%rbx, 32(%rsp)          # 8-byte Spill
	movq	(%rdx), %rdx
	leaq	1168(%rsp), %rdi
	callq	.LmulPv512x64
	movq	1168(%rsp), %r15
	movq	1176(%rsp), %r12
	movq	%r15, %rdx
	imulq	%rbx, %rdx
	movq	1232(%rsp), %rax
	movq	%rax, 48(%rsp)          # 8-byte Spill
	movq	1224(%rsp), %rax
	movq	%rax, 80(%rsp)          # 8-byte Spill
	movq	1216(%rsp), %r13
	movq	1208(%rsp), %rax
	movq	%rax, 72(%rsp)          # 8-byte Spill
	movq	1200(%rsp), %r14
	movq	1192(%rsp), %rbp
	movq	1184(%rsp), %rbx
	leaq	1096(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	addq	1096(%rsp), %r15
	adcq	1104(%rsp), %r12
	movq	%r12, 64(%rsp)          # 8-byte Spill
	adcq	1112(%rsp), %rbx
	adcq	1120(%rsp), %rbp
	adcq	1128(%rsp), %r14
	movq	%r14, %r12
	movq	72(%rsp), %r14          # 8-byte Reload
	adcq	1136(%rsp), %r14
	adcq	1144(%rsp), %r13
	movq	80(%rsp), %rax          # 8-byte Reload
	adcq	1152(%rsp), %rax
	movq	%rax, 80(%rsp)          # 8-byte Spill
	movq	48(%rsp), %rax          # 8-byte Reload
	adcq	1160(%rsp), %rax
	movq	%rax, 48(%rsp)          # 8-byte Spill
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	8(%rax), %rdx
	leaq	1024(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	movq	1088(%rsp), %r15
	movq	64(%rsp), %rax          # 8-byte Reload
	addq	1024(%rsp), %rax
	adcq	1032(%rsp), %rbx
	movq	%rbx, 8(%rsp)           # 8-byte Spill
	movq	%rbp, %rbx
	adcq	1040(%rsp), %rbx
	adcq	1048(%rsp), %r12
	adcq	1056(%rsp), %r14
	movq	%r14, 72(%rsp)          # 8-byte Spill
	movq	%r13, %rbp
	adcq	1064(%rsp), %rbp
	movq	80(%rsp), %rcx          # 8-byte Reload
	adcq	1072(%rsp), %rcx
	movq	%rcx, 80(%rsp)          # 8-byte Spill
	movq	48(%rsp), %r14          # 8-byte Reload
	adcq	1080(%rsp), %r14
	adcq	$0, %r15
	movq	%rax, %rdx
	movq	%rax, %r13
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	952(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	addq	952(%rsp), %r13
	movq	8(%rsp), %rax           # 8-byte Reload
	adcq	960(%rsp), %rax
	movq	%rax, 8(%rsp)           # 8-byte Spill
	adcq	968(%rsp), %rbx
	movq	%rbx, 64(%rsp)          # 8-byte Spill
	movq	%r12, %rbx
	adcq	976(%rsp), %rbx
	movq	72(%rsp), %r12          # 8-byte Reload
	adcq	984(%rsp), %r12
	adcq	992(%rsp), %rbp
	movq	%rbp, 40(%rsp)          # 8-byte Spill
	movq	80(%rsp), %r13          # 8-byte Reload
	adcq	1000(%rsp), %r13
	movq	%r14, %rbp
	adcq	1008(%rsp), %rbp
	adcq	1016(%rsp), %r15
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	16(%rax), %rdx
	leaq	880(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	movq	944(%rsp), %r14
	movq	8(%rsp), %rax           # 8-byte Reload
	addq	880(%rsp), %rax
	movq	64(%rsp), %rcx          # 8-byte Reload
	adcq	888(%rsp), %rcx
	movq	%rcx, 64(%rsp)          # 8-byte Spill
	adcq	896(%rsp), %rbx
	adcq	904(%rsp), %r12
	movq	%r12, 72(%rsp)          # 8-byte Spill
	movq	40(%rsp), %rcx          # 8-byte Reload
	adcq	912(%rsp), %rcx
	movq	%rcx, 40(%rsp)          # 8-byte Spill
	adcq	920(%rsp), %r13
	movq	%r13, 80(%rsp)          # 8-byte Spill
	adcq	928(%rsp), %rbp
	movq	%rbp, 48(%rsp)          # 8-byte Spill
	adcq	936(%rsp), %r15
	adcq	$0, %r14
	movq	%rax, %rdx
	movq	%rax, %rbp
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	808(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	addq	808(%rsp), %rbp
	movq	64(%rsp), %r13          # 8-byte Reload
	adcq	816(%rsp), %r13
	movq	%rbx, %r12
	adcq	824(%rsp), %r12
	movq	72(%rsp), %rbx          # 8-byte Reload
	adcq	832(%rsp), %rbx
	movq	40(%rsp), %rbp          # 8-byte Reload
	adcq	840(%rsp), %rbp
	movq	80(%rsp), %rax          # 8-byte Reload
	adcq	848(%rsp), %rax
	movq	%rax, 80(%rsp)          # 8-byte Spill
	movq	48(%rsp), %rax          # 8-byte Reload
	adcq	856(%rsp), %rax
	movq	%rax, 48(%rsp)          # 8-byte Spill
	adcq	864(%rsp), %r15
	adcq	872(%rsp), %r14
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	24(%rax), %rdx
	leaq	736(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	movq	800(%rsp), %rax
	movq	%r13, %rcx
	addq	736(%rsp), %rcx
	adcq	744(%rsp), %r12
	movq	%r12, 40(%rsp)          # 8-byte Spill
	adcq	752(%rsp), %rbx
	movq	%rbx, 72(%rsp)          # 8-byte Spill
	adcq	760(%rsp), %rbp
	movq	%rbp, %r13
	movq	80(%rsp), %rbp          # 8-byte Reload
	adcq	768(%rsp), %rbp
	movq	48(%rsp), %rbx          # 8-byte Reload
	adcq	776(%rsp), %rbx
	adcq	784(%rsp), %r15
	adcq	792(%rsp), %r14
	adcq	$0, %rax
	movq	%rax, 64(%rsp)          # 8-byte Spill
	movq	%rcx, %rdx
	movq	%rcx, %r12
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	664(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	addq	664(%rsp), %r12
	movq	40(%rsp), %rax          # 8-byte Reload
	adcq	672(%rsp), %rax
	movq	%rax, 40(%rsp)          # 8-byte Spill
	movq	72(%rsp), %rax          # 8-byte Reload
	adcq	680(%rsp), %rax
	movq	%rax, 72(%rsp)          # 8-byte Spill
	adcq	688(%rsp), %r13
	adcq	696(%rsp), %rbp
	movq	%rbp, 80(%rsp)          # 8-byte Spill
	adcq	704(%rsp), %rbx
	adcq	712(%rsp), %r15
	adcq	720(%rsp), %r14
	movq	64(%rsp), %r12          # 8-byte Reload
	adcq	728(%rsp), %r12
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	32(%rax), %rdx
	leaq	592(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	movq	656(%rsp), %rcx
	movq	40(%rsp), %rax          # 8-byte Reload
	addq	592(%rsp), %rax
	movq	72(%rsp), %rbp          # 8-byte Reload
	adcq	600(%rsp), %rbp
	adcq	608(%rsp), %r13
	movq	%r13, 40(%rsp)          # 8-byte Spill
	movq	80(%rsp), %r13          # 8-byte Reload
	adcq	616(%rsp), %r13
	adcq	624(%rsp), %rbx
	adcq	632(%rsp), %r15
	adcq	640(%rsp), %r14
	adcq	648(%rsp), %r12
	movq	%r12, 64(%rsp)          # 8-byte Spill
	adcq	$0, %rcx
	movq	%rcx, 80(%rsp)          # 8-byte Spill
	movq	%rax, %rdx
	movq	%rax, %r12
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	520(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	addq	520(%rsp), %r12
	adcq	528(%rsp), %rbp
	movq	%rbp, 72(%rsp)          # 8-byte Spill
	movq	40(%rsp), %r12          # 8-byte Reload
	adcq	536(%rsp), %r12
	movq	%r13, %rbp
	adcq	544(%rsp), %rbp
	adcq	552(%rsp), %rbx
	adcq	560(%rsp), %r15
	adcq	568(%rsp), %r14
	movq	64(%rsp), %r13          # 8-byte Reload
	adcq	576(%rsp), %r13
	movq	80(%rsp), %rax          # 8-byte Reload
	adcq	584(%rsp), %rax
	movq	%rax, 80(%rsp)          # 8-byte Spill
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	40(%rax), %rdx
	leaq	448(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	movq	512(%rsp), %rcx
	movq	72(%rsp), %rax          # 8-byte Reload
	addq	448(%rsp), %rax
	adcq	456(%rsp), %r12
	movq	%r12, 40(%rsp)          # 8-byte Spill
	adcq	464(%rsp), %rbp
	adcq	472(%rsp), %rbx
	adcq	480(%rsp), %r15
	adcq	488(%rsp), %r14
	adcq	496(%rsp), %r13
	movq	%r13, 64(%rsp)          # 8-byte Spill
	movq	80(%rsp), %r13          # 8-byte Reload
	adcq	504(%rsp), %r13
	adcq	$0, %rcx
	movq	%rcx, 72(%rsp)          # 8-byte Spill
	movq	%rax, %rdx
	movq	%rax, %r12
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	376(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	addq	376(%rsp), %r12
	movq	40(%rsp), %rax          # 8-byte Reload
	adcq	384(%rsp), %rax
	movq	%rax, 40(%rsp)          # 8-byte Spill
	adcq	392(%rsp), %rbp
	adcq	400(%rsp), %rbx
	adcq	408(%rsp), %r15
	adcq	416(%rsp), %r14
	movq	64(%rsp), %r12          # 8-byte Reload
	adcq	424(%rsp), %r12
	adcq	432(%rsp), %r13
	movq	72(%rsp), %rax          # 8-byte Reload
	adcq	440(%rsp), %rax
	movq	%rax, 72(%rsp)          # 8-byte Spill
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	48(%rax), %rdx
	leaq	304(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	movq	368(%rsp), %rcx
	movq	40(%rsp), %rax          # 8-byte Reload
	addq	304(%rsp), %rax
	adcq	312(%rsp), %rbp
	movq	%rbp, 80(%rsp)          # 8-byte Spill
	adcq	320(%rsp), %rbx
	adcq	328(%rsp), %r15
	adcq	336(%rsp), %r14
	adcq	344(%rsp), %r12
	movq	%r12, 64(%rsp)          # 8-byte Spill
	adcq	352(%rsp), %r13
	movq	72(%rsp), %rbp          # 8-byte Reload
	adcq	360(%rsp), %rbp
	adcq	$0, %rcx
	movq	%rcx, 48(%rsp)          # 8-byte Spill
	movq	%rax, %rdx
	movq	%rax, %r12
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	232(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	addq	232(%rsp), %r12
	movq	80(%rsp), %rax          # 8-byte Reload
	adcq	240(%rsp), %rax
	movq	%rax, 80(%rsp)          # 8-byte Spill
	adcq	248(%rsp), %rbx
	adcq	256(%rsp), %r15
	adcq	264(%rsp), %r14
	movq	64(%rsp), %r12          # 8-byte Reload
	adcq	272(%rsp), %r12
	adcq	280(%rsp), %r13
	adcq	288(%rsp), %rbp
	movq	%rbp, 72(%rsp)          # 8-byte Spill
	movq	48(%rsp), %rbp          # 8-byte Reload
	adcq	296(%rsp), %rbp
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	56(%rax), %rdx
	leaq	160(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	movq	224(%rsp), %rcx
	movq	80(%rsp), %rax          # 8-byte Reload
	addq	160(%rsp), %rax
	adcq	168(%rsp), %rbx
	movq	%rbx, 48(%rsp)          # 8-byte Spill
	adcq	176(%rsp), %r15
	adcq	184(%rsp), %r14
	adcq	192(%rsp), %r12
	movq	%r12, 64(%rsp)          # 8-byte Spill
	adcq	200(%rsp), %r13
	movq	72(%rsp), %rbx          # 8-byte Reload
	adcq	208(%rsp), %rbx
	adcq	216(%rsp), %rbp
	movq	%rbp, %r12
	adcq	$0, %rcx
	movq	%rcx, 80(%rsp)          # 8-byte Spill
	movq	32(%rsp), %rdx          # 8-byte Reload
	imulq	%rax, %rdx
	movq	%rax, %rbp
	leaq	88(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv512x64
	addq	88(%rsp), %rbp
	movq	48(%rsp), %r11          # 8-byte Reload
	adcq	96(%rsp), %r11
	adcq	104(%rsp), %r15
	adcq	112(%rsp), %r14
	movq	64(%rsp), %rsi          # 8-byte Reload
	adcq	120(%rsp), %rsi
	movq	%rsi, 64(%rsp)          # 8-byte Spill
	adcq	128(%rsp), %r13
	adcq	136(%rsp), %rbx
	movq	%rbx, 72(%rsp)          # 8-byte Spill
	adcq	144(%rsp), %r12
	movq	80(%rsp), %r8           # 8-byte Reload
	adcq	152(%rsp), %r8
	movq	%r11, %rax
	movq	56(%rsp), %rbp          # 8-byte Reload
	subq	(%rbp), %rax
	movq	%r15, %rcx
	sbbq	8(%rbp), %rcx
	movq	%r14, %rdx
	sbbq	16(%rbp), %rdx
	sbbq	24(%rbp), %rsi
	movq	%r13, %rdi
	sbbq	32(%rbp), %rdi
	movq	%rbx, %r9
	sbbq	40(%rbp), %r9
	movq	%r12, %r10
	sbbq	48(%rbp), %r10
	movq	%rbp, %rbx
	movq	%r8, %rbp
	sbbq	56(%rbx), %rbp
	testq	%rbp, %rbp
	cmovsq	%r11, %rax
	movq	(%rsp), %rbx            # 8-byte Reload
	movq	%rax, (%rbx)
	cmovsq	%r15, %rcx
	movq	%rcx, 8(%rbx)
	cmovsq	%r14, %rdx
	movq	%rdx, 16(%rbx)
	cmovsq	64(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 24(%rbx)
	cmovsq	%r13, %rdi
	movq	%rdi, 32(%rbx)
	cmovsq	72(%rsp), %r9           # 8-byte Folded Reload
	movq	%r9, 40(%rbx)
	cmovsq	%r12, %r10
	movq	%r10, 48(%rbx)
	cmovsq	%r8, %rbp
	movq	%rbp, 56(%rbx)
	addq	$1240, %rsp             # imm = 0x4D8
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end115:
	.size	mcl_fp_montNF8Lbmi2, .Lfunc_end115-mcl_fp_montNF8Lbmi2

	.globl	mcl_fp_montRed8Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montRed8Lbmi2,@function
mcl_fp_montRed8Lbmi2:                   # @mcl_fp_montRed8Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	subq	$776, %rsp              # imm = 0x308
	movq	%rdx, %rax
	movq	%rax, 112(%rsp)         # 8-byte Spill
	movq	%rdi, 72(%rsp)          # 8-byte Spill
	movq	-8(%rax), %rcx
	movq	%rcx, 128(%rsp)         # 8-byte Spill
	movq	(%rsi), %r15
	movq	8(%rsi), %rdx
	movq	%rdx, 184(%rsp)         # 8-byte Spill
	movq	%r15, %rdx
	imulq	%rcx, %rdx
	movq	120(%rsi), %rcx
	movq	%rcx, 96(%rsp)          # 8-byte Spill
	movq	112(%rsi), %rcx
	movq	%rcx, 136(%rsp)         # 8-byte Spill
	movq	104(%rsi), %rcx
	movq	%rcx, 120(%rsp)         # 8-byte Spill
	movq	96(%rsi), %rcx
	movq	%rcx, 168(%rsp)         # 8-byte Spill
	movq	88(%rsi), %rcx
	movq	%rcx, 176(%rsp)         # 8-byte Spill
	movq	80(%rsi), %rcx
	movq	%rcx, 160(%rsp)         # 8-byte Spill
	movq	72(%rsi), %rcx
	movq	%rcx, 192(%rsp)         # 8-byte Spill
	movq	64(%rsi), %r13
	movq	56(%rsi), %rcx
	movq	%rcx, 144(%rsp)         # 8-byte Spill
	movq	48(%rsi), %r14
	movq	40(%rsi), %rcx
	movq	%rcx, 152(%rsp)         # 8-byte Spill
	movq	32(%rsi), %r12
	movq	24(%rsi), %rbx
	movq	16(%rsi), %rbp
	movq	%rax, %rcx
	movq	(%rcx), %rax
	movq	%rax, 16(%rsp)          # 8-byte Spill
	movq	56(%rcx), %rax
	movq	%rax, 64(%rsp)          # 8-byte Spill
	movq	48(%rcx), %rax
	movq	%rax, 56(%rsp)          # 8-byte Spill
	movq	40(%rcx), %rax
	movq	%rax, 48(%rsp)          # 8-byte Spill
	movq	32(%rcx), %rax
	movq	%rax, 40(%rsp)          # 8-byte Spill
	movq	24(%rcx), %rax
	movq	%rax, 32(%rsp)          # 8-byte Spill
	movq	16(%rcx), %rax
	movq	%rax, 24(%rsp)          # 8-byte Spill
	movq	8(%rcx), %rax
	movq	%rax, 8(%rsp)           # 8-byte Spill
	movq	%rcx, %rsi
	leaq	704(%rsp), %rdi
	callq	.LmulPv512x64
	addq	704(%rsp), %r15
	movq	184(%rsp), %rcx         # 8-byte Reload
	adcq	712(%rsp), %rcx
	adcq	720(%rsp), %rbp
	movq	%rbp, 80(%rsp)          # 8-byte Spill
	adcq	728(%rsp), %rbx
	movq	%rbx, 88(%rsp)          # 8-byte Spill
	adcq	736(%rsp), %r12
	movq	%r12, 104(%rsp)         # 8-byte Spill
	movq	152(%rsp), %rax         # 8-byte Reload
	adcq	744(%rsp), %rax
	movq	%rax, 152(%rsp)         # 8-byte Spill
	adcq	752(%rsp), %r14
	movq	%r14, %r12
	movq	144(%rsp), %rax         # 8-byte Reload
	adcq	760(%rsp), %rax
	movq	%rax, 144(%rsp)         # 8-byte Spill
	adcq	768(%rsp), %r13
	movq	%r13, 184(%rsp)         # 8-byte Spill
	adcq	$0, 192(%rsp)           # 8-byte Folded Spill
	movq	160(%rsp), %r15         # 8-byte Reload
	adcq	$0, %r15
	adcq	$0, 176(%rsp)           # 8-byte Folded Spill
	adcq	$0, 168(%rsp)           # 8-byte Folded Spill
	adcq	$0, 120(%rsp)           # 8-byte Folded Spill
	movq	136(%rsp), %r13         # 8-byte Reload
	adcq	$0, %r13
	movq	96(%rsp), %r14          # 8-byte Reload
	adcq	$0, %r14
	sbbq	%rbx, %rbx
	movq	%rcx, %rbp
	movq	%rbp, %rdx
	imulq	128(%rsp), %rdx         # 8-byte Folded Reload
	leaq	632(%rsp), %rdi
	movq	112(%rsp), %rsi         # 8-byte Reload
	callq	.LmulPv512x64
	andl	$1, %ebx
	movq	%rbx, %rax
	addq	632(%rsp), %rbp
	movq	80(%rsp), %rsi          # 8-byte Reload
	adcq	640(%rsp), %rsi
	movq	88(%rsp), %rcx          # 8-byte Reload
	adcq	648(%rsp), %rcx
	movq	%rcx, 88(%rsp)          # 8-byte Spill
	movq	104(%rsp), %rcx         # 8-byte Reload
	adcq	656(%rsp), %rcx
	movq	%rcx, 104(%rsp)         # 8-byte Spill
	movq	152(%rsp), %rcx         # 8-byte Reload
	adcq	664(%rsp), %rcx
	movq	%rcx, 152(%rsp)         # 8-byte Spill
	adcq	672(%rsp), %r12
	movq	144(%rsp), %rcx         # 8-byte Reload
	adcq	680(%rsp), %rcx
	movq	%rcx, 144(%rsp)         # 8-byte Spill
	movq	184(%rsp), %rcx         # 8-byte Reload
	adcq	688(%rsp), %rcx
	movq	%rcx, 184(%rsp)         # 8-byte Spill
	movq	192(%rsp), %rcx         # 8-byte Reload
	adcq	696(%rsp), %rcx
	movq	%rcx, 192(%rsp)         # 8-byte Spill
	adcq	$0, %r15
	movq	%r15, 160(%rsp)         # 8-byte Spill
	movq	176(%rsp), %rbx         # 8-byte Reload
	adcq	$0, %rbx
	movq	168(%rsp), %r15         # 8-byte Reload
	adcq	$0, %r15
	adcq	$0, 120(%rsp)           # 8-byte Folded Spill
	adcq	$0, %r13
	movq	%r13, 136(%rsp)         # 8-byte Spill
	adcq	$0, %r14
	movq	%r14, 96(%rsp)          # 8-byte Spill
	movq	%rax, %rbp
	adcq	$0, %rbp
	movq	%rsi, %rdx
	movq	%rsi, %r14
	imulq	128(%rsp), %rdx         # 8-byte Folded Reload
	leaq	560(%rsp), %rdi
	movq	112(%rsp), %r13         # 8-byte Reload
	movq	%r13, %rsi
	callq	.LmulPv512x64
	addq	560(%rsp), %r14
	movq	88(%rsp), %rcx          # 8-byte Reload
	adcq	568(%rsp), %rcx
	movq	104(%rsp), %rax         # 8-byte Reload
	adcq	576(%rsp), %rax
	movq	%rax, 104(%rsp)         # 8-byte Spill
	movq	152(%rsp), %rax         # 8-byte Reload
	adcq	584(%rsp), %rax
	movq	%rax, 152(%rsp)         # 8-byte Spill
	adcq	592(%rsp), %r12
	movq	%r12, 88(%rsp)          # 8-byte Spill
	movq	144(%rsp), %r14         # 8-byte Reload
	adcq	600(%rsp), %r14
	movq	184(%rsp), %rax         # 8-byte Reload
	adcq	608(%rsp), %rax
	movq	%rax, 184(%rsp)         # 8-byte Spill
	movq	192(%rsp), %rax         # 8-byte Reload
	adcq	616(%rsp), %rax
	movq	%rax, 192(%rsp)         # 8-byte Spill
	movq	160(%rsp), %rax         # 8-byte Reload
	adcq	624(%rsp), %rax
	movq	%rax, 160(%rsp)         # 8-byte Spill
	adcq	$0, %rbx
	movq	%rbx, 176(%rsp)         # 8-byte Spill
	adcq	$0, %r15
	movq	%r15, 168(%rsp)         # 8-byte Spill
	movq	120(%rsp), %rbx         # 8-byte Reload
	adcq	$0, %rbx
	movq	136(%rsp), %r15         # 8-byte Reload
	adcq	$0, %r15
	adcq	$0, 96(%rsp)            # 8-byte Folded Spill
	adcq	$0, %rbp
	movq	%rbp, 80(%rsp)          # 8-byte Spill
	movq	%rcx, %rbp
	movq	%rbp, %rdx
	movq	128(%rsp), %r12         # 8-byte Reload
	imulq	%r12, %rdx
	leaq	488(%rsp), %rdi
	movq	%r13, %rsi
	callq	.LmulPv512x64
	addq	488(%rsp), %rbp
	movq	104(%rsp), %rax         # 8-byte Reload
	adcq	496(%rsp), %rax
	movq	152(%rsp), %rbp         # 8-byte Reload
	adcq	504(%rsp), %rbp
	movq	88(%rsp), %rcx          # 8-byte Reload
	adcq	512(%rsp), %rcx
	movq	%rcx, 88(%rsp)          # 8-byte Spill
	adcq	520(%rsp), %r14
	movq	184(%rsp), %rcx         # 8-byte Reload
	adcq	528(%rsp), %rcx
	movq	%rcx, 184(%rsp)         # 8-byte Spill
	movq	192(%rsp), %rcx         # 8-byte Reload
	adcq	536(%rsp), %rcx
	movq	%rcx, 192(%rsp)         # 8-byte Spill
	movq	160(%rsp), %r13         # 8-byte Reload
	adcq	544(%rsp), %r13
	movq	176(%rsp), %rcx         # 8-byte Reload
	adcq	552(%rsp), %rcx
	movq	%rcx, 176(%rsp)         # 8-byte Spill
	adcq	$0, 168(%rsp)           # 8-byte Folded Spill
	adcq	$0, %rbx
	movq	%rbx, 120(%rsp)         # 8-byte Spill
	movq	%r15, %rbx
	adcq	$0, %rbx
	adcq	$0, 96(%rsp)            # 8-byte Folded Spill
	adcq	$0, 80(%rsp)            # 8-byte Folded Spill
	movq	%rax, %rdx
	movq	%rax, %r15
	imulq	%r12, %rdx
	leaq	416(%rsp), %rdi
	movq	112(%rsp), %rsi         # 8-byte Reload
	callq	.LmulPv512x64
	addq	416(%rsp), %r15
	adcq	424(%rsp), %rbp
	movq	%rbp, %rax
	movq	88(%rsp), %rcx          # 8-byte Reload
	adcq	432(%rsp), %rcx
	movq	%rcx, 88(%rsp)          # 8-byte Spill
	movq	%r14, %r12
	adcq	440(%rsp), %r12
	movq	184(%rsp), %r14         # 8-byte Reload
	adcq	448(%rsp), %r14
	movq	192(%rsp), %rbp         # 8-byte Reload
	adcq	456(%rsp), %rbp
	adcq	464(%rsp), %r13
	movq	176(%rsp), %rcx         # 8-byte Reload
	adcq	472(%rsp), %rcx
	movq	%rcx, 176(%rsp)         # 8-byte Spill
	movq	168(%rsp), %rcx         # 8-byte Reload
	adcq	480(%rsp), %rcx
	movq	%rcx, 168(%rsp)         # 8-byte Spill
	adcq	$0, 120(%rsp)           # 8-byte Folded Spill
	adcq	$0, %rbx
	movq	%rbx, 136(%rsp)         # 8-byte Spill
	movq	96(%rsp), %r15          # 8-byte Reload
	adcq	$0, %r15
	adcq	$0, 80(%rsp)            # 8-byte Folded Spill
	movq	%rax, %rbx
	movq	%rbx, %rdx
	imulq	128(%rsp), %rdx         # 8-byte Folded Reload
	leaq	344(%rsp), %rdi
	movq	112(%rsp), %rsi         # 8-byte Reload
	callq	.LmulPv512x64
	addq	344(%rsp), %rbx
	movq	88(%rsp), %rax          # 8-byte Reload
	adcq	352(%rsp), %rax
	adcq	360(%rsp), %r12
	movq	%r12, 144(%rsp)         # 8-byte Spill
	adcq	368(%rsp), %r14
	movq	%r14, 184(%rsp)         # 8-byte Spill
	adcq	376(%rsp), %rbp
	movq	%rbp, 192(%rsp)         # 8-byte Spill
	adcq	384(%rsp), %r13
	movq	%r13, 160(%rsp)         # 8-byte Spill
	movq	176(%rsp), %r13         # 8-byte Reload
	adcq	392(%rsp), %r13
	movq	168(%rsp), %r12         # 8-byte Reload
	adcq	400(%rsp), %r12
	movq	120(%rsp), %r14         # 8-byte Reload
	adcq	408(%rsp), %r14
	movq	136(%rsp), %rbp         # 8-byte Reload
	adcq	$0, %rbp
	movq	%r15, %rbx
	adcq	$0, %rbx
	adcq	$0, 80(%rsp)            # 8-byte Folded Spill
	movq	%rax, %rdx
	movq	%rax, %r15
	imulq	128(%rsp), %rdx         # 8-byte Folded Reload
	leaq	272(%rsp), %rdi
	movq	112(%rsp), %rsi         # 8-byte Reload
	callq	.LmulPv512x64
	addq	272(%rsp), %r15
	movq	144(%rsp), %rcx         # 8-byte Reload
	adcq	280(%rsp), %rcx
	movq	184(%rsp), %rax         # 8-byte Reload
	adcq	288(%rsp), %rax
	movq	%rax, 184(%rsp)         # 8-byte Spill
	movq	192(%rsp), %rax         # 8-byte Reload
	adcq	296(%rsp), %rax
	movq	%rax, 192(%rsp)         # 8-byte Spill
	movq	160(%rsp), %rax         # 8-byte Reload
	adcq	304(%rsp), %rax
	movq	%rax, 160(%rsp)         # 8-byte Spill
	adcq	312(%rsp), %r13
	movq	%r13, 176(%rsp)         # 8-byte Spill
	adcq	320(%rsp), %r12
	movq	%r12, 168(%rsp)         # 8-byte Spill
	adcq	328(%rsp), %r14
	movq	%r14, %r13
	adcq	336(%rsp), %rbp
	movq	%rbp, %r12
	adcq	$0, %rbx
	movq	%rbx, %r14
	movq	80(%rsp), %r15          # 8-byte Reload
	adcq	$0, %r15
	movq	128(%rsp), %rdx         # 8-byte Reload
	movq	%rcx, %rbx
	imulq	%rbx, %rdx
	leaq	200(%rsp), %rdi
	movq	112(%rsp), %rsi         # 8-byte Reload
	callq	.LmulPv512x64
	addq	200(%rsp), %rbx
	movq	184(%rsp), %rax         # 8-byte Reload
	adcq	208(%rsp), %rax
	movq	%rax, 184(%rsp)         # 8-byte Spill
	movq	192(%rsp), %r8          # 8-byte Reload
	adcq	216(%rsp), %r8
	movq	%r8, 192(%rsp)          # 8-byte Spill
	movq	160(%rsp), %rdx         # 8-byte Reload
	adcq	224(%rsp), %rdx
	movq	176(%rsp), %rsi         # 8-byte Reload
	adcq	232(%rsp), %rsi
	movq	168(%rsp), %rdi         # 8-byte Reload
	adcq	240(%rsp), %rdi
	movq	%r13, %rbp
	adcq	248(%rsp), %rbp
	movq	%r12, %rbx
	adcq	256(%rsp), %rbx
	movq	%rbx, 136(%rsp)         # 8-byte Spill
	movq	%r14, %r9
	adcq	264(%rsp), %r9
	adcq	$0, %r15
	movq	%r15, %r10
	subq	16(%rsp), %rax          # 8-byte Folded Reload
	movq	%r8, %rcx
	sbbq	8(%rsp), %rcx           # 8-byte Folded Reload
	movq	%rdx, %r13
	sbbq	24(%rsp), %r13          # 8-byte Folded Reload
	movq	%rsi, %r12
	sbbq	32(%rsp), %r12          # 8-byte Folded Reload
	movq	%rdi, %r14
	sbbq	40(%rsp), %r14          # 8-byte Folded Reload
	movq	%rbp, %r11
	sbbq	48(%rsp), %r11          # 8-byte Folded Reload
	movq	%rbx, %r8
	sbbq	56(%rsp), %r8           # 8-byte Folded Reload
	movq	%r9, %r15
	sbbq	64(%rsp), %r9           # 8-byte Folded Reload
	sbbq	$0, %r10
	andl	$1, %r10d
	cmovneq	%r15, %r9
	testb	%r10b, %r10b
	cmovneq	184(%rsp), %rax         # 8-byte Folded Reload
	movq	72(%rsp), %rbx          # 8-byte Reload
	movq	%rax, (%rbx)
	cmovneq	192(%rsp), %rcx         # 8-byte Folded Reload
	movq	%rcx, 8(%rbx)
	cmovneq	%rdx, %r13
	movq	%r13, 16(%rbx)
	cmovneq	%rsi, %r12
	movq	%r12, 24(%rbx)
	cmovneq	%rdi, %r14
	movq	%r14, 32(%rbx)
	cmovneq	%rbp, %r11
	movq	%r11, 40(%rbx)
	cmovneq	136(%rsp), %r8          # 8-byte Folded Reload
	movq	%r8, 48(%rbx)
	movq	%r9, 56(%rbx)
	addq	$776, %rsp              # imm = 0x308
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end116:
	.size	mcl_fp_montRed8Lbmi2, .Lfunc_end116-mcl_fp_montRed8Lbmi2

	.globl	mcl_fp_addPre8Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addPre8Lbmi2,@function
mcl_fp_addPre8Lbmi2:                    # @mcl_fp_addPre8Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	56(%rdx), %r8
	movq	56(%rsi), %r15
	movq	48(%rdx), %r9
	movq	48(%rsi), %r12
	movq	40(%rdx), %r10
	movq	32(%rdx), %r11
	movq	24(%rdx), %r14
	movq	16(%rdx), %rbx
	movq	(%rdx), %rcx
	movq	8(%rdx), %rdx
	addq	(%rsi), %rcx
	adcq	8(%rsi), %rdx
	adcq	16(%rsi), %rbx
	movq	40(%rsi), %r13
	movq	24(%rsi), %rax
	movq	32(%rsi), %rsi
	movq	%rcx, (%rdi)
	movq	%rdx, 8(%rdi)
	movq	%rbx, 16(%rdi)
	adcq	%r14, %rax
	movq	%rax, 24(%rdi)
	adcq	%r11, %rsi
	movq	%rsi, 32(%rdi)
	adcq	%r10, %r13
	movq	%r13, 40(%rdi)
	adcq	%r9, %r12
	movq	%r12, 48(%rdi)
	adcq	%r8, %r15
	movq	%r15, 56(%rdi)
	sbbq	%rax, %rax
	andl	$1, %eax
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	retq
.Lfunc_end117:
	.size	mcl_fp_addPre8Lbmi2, .Lfunc_end117-mcl_fp_addPre8Lbmi2

	.globl	mcl_fp_subPre8Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subPre8Lbmi2,@function
mcl_fp_subPre8Lbmi2:                    # @mcl_fp_subPre8Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	56(%rdx), %r8
	movq	56(%rsi), %r15
	movq	48(%rdx), %r9
	movq	40(%rdx), %r10
	movq	24(%rdx), %r11
	movq	32(%rdx), %r14
	movq	(%rsi), %rbx
	movq	8(%rsi), %r12
	xorl	%eax, %eax
	subq	(%rdx), %rbx
	sbbq	8(%rdx), %r12
	movq	16(%rsi), %rcx
	sbbq	16(%rdx), %rcx
	movq	48(%rsi), %r13
	movq	40(%rsi), %rdx
	movq	32(%rsi), %rbp
	movq	24(%rsi), %rsi
	movq	%rbx, (%rdi)
	movq	%r12, 8(%rdi)
	movq	%rcx, 16(%rdi)
	sbbq	%r11, %rsi
	movq	%rsi, 24(%rdi)
	sbbq	%r14, %rbp
	movq	%rbp, 32(%rdi)
	sbbq	%r10, %rdx
	movq	%rdx, 40(%rdi)
	sbbq	%r9, %r13
	movq	%r13, 48(%rdi)
	sbbq	%r8, %r15
	movq	%r15, 56(%rdi)
	sbbq	$0, %rax
	andl	$1, %eax
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end118:
	.size	mcl_fp_subPre8Lbmi2, .Lfunc_end118-mcl_fp_subPre8Lbmi2

	.globl	mcl_fp_shr1_8Lbmi2
	.align	16, 0x90
	.type	mcl_fp_shr1_8Lbmi2,@function
mcl_fp_shr1_8Lbmi2:                     # @mcl_fp_shr1_8Lbmi2
# BB#0:
	movq	56(%rsi), %r8
	movq	48(%rsi), %r9
	movq	40(%rsi), %r10
	movq	32(%rsi), %r11
	movq	24(%rsi), %rcx
	movq	16(%rsi), %rdx
	movq	(%rsi), %rax
	movq	8(%rsi), %rsi
	shrdq	$1, %rsi, %rax
	movq	%rax, (%rdi)
	shrdq	$1, %rdx, %rsi
	movq	%rsi, 8(%rdi)
	shrdq	$1, %rcx, %rdx
	movq	%rdx, 16(%rdi)
	shrdq	$1, %r11, %rcx
	movq	%rcx, 24(%rdi)
	shrdq	$1, %r10, %r11
	movq	%r11, 32(%rdi)
	shrdq	$1, %r9, %r10
	movq	%r10, 40(%rdi)
	shrdq	$1, %r8, %r9
	movq	%r9, 48(%rdi)
	shrq	%r8
	movq	%r8, 56(%rdi)
	retq
.Lfunc_end119:
	.size	mcl_fp_shr1_8Lbmi2, .Lfunc_end119-mcl_fp_shr1_8Lbmi2

	.globl	mcl_fp_add8Lbmi2
	.align	16, 0x90
	.type	mcl_fp_add8Lbmi2,@function
mcl_fp_add8Lbmi2:                       # @mcl_fp_add8Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	56(%rdx), %r15
	movq	56(%rsi), %r8
	movq	48(%rdx), %r12
	movq	48(%rsi), %r9
	movq	40(%rsi), %r13
	movq	24(%rsi), %r11
	movq	32(%rsi), %r10
	movq	(%rdx), %r14
	movq	8(%rdx), %rbx
	addq	(%rsi), %r14
	adcq	8(%rsi), %rbx
	movq	16(%rdx), %rax
	adcq	16(%rsi), %rax
	adcq	24(%rdx), %r11
	movq	40(%rdx), %rsi
	adcq	32(%rdx), %r10
	movq	%r14, (%rdi)
	movq	%rbx, 8(%rdi)
	movq	%rax, 16(%rdi)
	movq	%r11, 24(%rdi)
	movq	%r10, 32(%rdi)
	adcq	%r13, %rsi
	movq	%rsi, 40(%rdi)
	adcq	%r12, %r9
	movq	%r9, 48(%rdi)
	adcq	%r15, %r8
	movq	%r8, 56(%rdi)
	sbbq	%rdx, %rdx
	andl	$1, %edx
	subq	(%rcx), %r14
	sbbq	8(%rcx), %rbx
	sbbq	16(%rcx), %rax
	sbbq	24(%rcx), %r11
	sbbq	32(%rcx), %r10
	sbbq	40(%rcx), %rsi
	sbbq	48(%rcx), %r9
	sbbq	56(%rcx), %r8
	sbbq	$0, %rdx
	testb	$1, %dl
	jne	.LBB120_2
# BB#1:                                 # %nocarry
	movq	%r14, (%rdi)
	movq	%rbx, 8(%rdi)
	movq	%rax, 16(%rdi)
	movq	%r11, 24(%rdi)
	movq	%r10, 32(%rdi)
	movq	%rsi, 40(%rdi)
	movq	%r9, 48(%rdi)
	movq	%r8, 56(%rdi)
.LBB120_2:                              # %carry
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	retq
.Lfunc_end120:
	.size	mcl_fp_add8Lbmi2, .Lfunc_end120-mcl_fp_add8Lbmi2

	.globl	mcl_fp_addNF8Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addNF8Lbmi2,@function
mcl_fp_addNF8Lbmi2:                     # @mcl_fp_addNF8Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	56(%rdx), %r8
	movq	48(%rdx), %rbp
	movq	40(%rdx), %rbx
	movq	32(%rdx), %rax
	movq	24(%rdx), %r11
	movq	16(%rdx), %r15
	movq	(%rdx), %r13
	movq	8(%rdx), %r12
	addq	(%rsi), %r13
	adcq	8(%rsi), %r12
	adcq	16(%rsi), %r15
	adcq	24(%rsi), %r11
	adcq	32(%rsi), %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	movq	%rax, %r10
	adcq	40(%rsi), %rbx
	movq	%rbx, -16(%rsp)         # 8-byte Spill
	movq	%rbx, %r9
	adcq	48(%rsi), %rbp
	movq	%rbp, -8(%rsp)          # 8-byte Spill
	movq	%rbp, %rax
	adcq	56(%rsi), %r8
	movq	%r13, %rsi
	subq	(%rcx), %rsi
	movq	%r12, %rdx
	sbbq	8(%rcx), %rdx
	movq	%r15, %rbx
	sbbq	16(%rcx), %rbx
	movq	%r11, %r14
	sbbq	24(%rcx), %r14
	movq	%r10, %rbp
	sbbq	32(%rcx), %rbp
	movq	%r9, %r10
	sbbq	40(%rcx), %r10
	movq	%rax, %r9
	sbbq	48(%rcx), %r9
	movq	%r8, %rax
	sbbq	56(%rcx), %rax
	testq	%rax, %rax
	cmovsq	%r13, %rsi
	movq	%rsi, (%rdi)
	cmovsq	%r12, %rdx
	movq	%rdx, 8(%rdi)
	cmovsq	%r15, %rbx
	movq	%rbx, 16(%rdi)
	cmovsq	%r11, %r14
	movq	%r14, 24(%rdi)
	cmovsq	-24(%rsp), %rbp         # 8-byte Folded Reload
	movq	%rbp, 32(%rdi)
	cmovsq	-16(%rsp), %r10         # 8-byte Folded Reload
	movq	%r10, 40(%rdi)
	cmovsq	-8(%rsp), %r9           # 8-byte Folded Reload
	movq	%r9, 48(%rdi)
	cmovsq	%r8, %rax
	movq	%rax, 56(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end121:
	.size	mcl_fp_addNF8Lbmi2, .Lfunc_end121-mcl_fp_addNF8Lbmi2

	.globl	mcl_fp_sub8Lbmi2
	.align	16, 0x90
	.type	mcl_fp_sub8Lbmi2,@function
mcl_fp_sub8Lbmi2:                       # @mcl_fp_sub8Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	56(%rdx), %r12
	movq	56(%rsi), %r8
	movq	48(%rdx), %r13
	movq	(%rsi), %rax
	movq	8(%rsi), %r10
	xorl	%ebx, %ebx
	subq	(%rdx), %rax
	sbbq	8(%rdx), %r10
	movq	16(%rsi), %r11
	sbbq	16(%rdx), %r11
	movq	24(%rsi), %r15
	sbbq	24(%rdx), %r15
	movq	32(%rsi), %r14
	sbbq	32(%rdx), %r14
	movq	48(%rsi), %r9
	movq	40(%rsi), %rsi
	sbbq	40(%rdx), %rsi
	movq	%rax, (%rdi)
	movq	%r10, 8(%rdi)
	movq	%r11, 16(%rdi)
	movq	%r15, 24(%rdi)
	movq	%r14, 32(%rdi)
	movq	%rsi, 40(%rdi)
	sbbq	%r13, %r9
	movq	%r9, 48(%rdi)
	sbbq	%r12, %r8
	movq	%r8, 56(%rdi)
	sbbq	$0, %rbx
	testb	$1, %bl
	je	.LBB122_2
# BB#1:                                 # %carry
	addq	(%rcx), %rax
	movq	%rax, (%rdi)
	movq	8(%rcx), %rax
	adcq	%r10, %rax
	movq	%rax, 8(%rdi)
	movq	16(%rcx), %rax
	adcq	%r11, %rax
	movq	%rax, 16(%rdi)
	movq	24(%rcx), %rax
	adcq	%r15, %rax
	movq	%rax, 24(%rdi)
	movq	32(%rcx), %rax
	adcq	%r14, %rax
	movq	%rax, 32(%rdi)
	movq	40(%rcx), %rax
	adcq	%rsi, %rax
	movq	%rax, 40(%rdi)
	movq	48(%rcx), %rax
	adcq	%r9, %rax
	movq	%rax, 48(%rdi)
	movq	56(%rcx), %rax
	adcq	%r8, %rax
	movq	%rax, 56(%rdi)
.LBB122_2:                              # %nocarry
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	retq
.Lfunc_end122:
	.size	mcl_fp_sub8Lbmi2, .Lfunc_end122-mcl_fp_sub8Lbmi2

	.globl	mcl_fp_subNF8Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subNF8Lbmi2,@function
mcl_fp_subNF8Lbmi2:                     # @mcl_fp_subNF8Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rcx, %r8
	movq	%rdi, %r9
	movq	56(%rsi), %r14
	movq	48(%rsi), %rax
	movq	40(%rsi), %rcx
	movq	32(%rsi), %rdi
	movq	24(%rsi), %r11
	movq	16(%rsi), %r15
	movq	(%rsi), %r13
	movq	8(%rsi), %r12
	subq	(%rdx), %r13
	sbbq	8(%rdx), %r12
	sbbq	16(%rdx), %r15
	sbbq	24(%rdx), %r11
	sbbq	32(%rdx), %rdi
	movq	%rdi, -24(%rsp)         # 8-byte Spill
	sbbq	40(%rdx), %rcx
	movq	%rcx, -16(%rsp)         # 8-byte Spill
	sbbq	48(%rdx), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	sbbq	56(%rdx), %r14
	movq	%r14, %rsi
	sarq	$63, %rsi
	movq	56(%r8), %r10
	andq	%rsi, %r10
	movq	48(%r8), %rbx
	andq	%rsi, %rbx
	movq	40(%r8), %rdi
	andq	%rsi, %rdi
	movq	32(%r8), %rbp
	andq	%rsi, %rbp
	movq	24(%r8), %rdx
	andq	%rsi, %rdx
	movq	16(%r8), %rcx
	andq	%rsi, %rcx
	movq	8(%r8), %rax
	andq	%rsi, %rax
	andq	(%r8), %rsi
	addq	%r13, %rsi
	adcq	%r12, %rax
	movq	%rsi, (%r9)
	adcq	%r15, %rcx
	movq	%rax, 8(%r9)
	movq	%rcx, 16(%r9)
	adcq	%r11, %rdx
	movq	%rdx, 24(%r9)
	adcq	-24(%rsp), %rbp         # 8-byte Folded Reload
	movq	%rbp, 32(%r9)
	adcq	-16(%rsp), %rdi         # 8-byte Folded Reload
	movq	%rdi, 40(%r9)
	adcq	-8(%rsp), %rbx          # 8-byte Folded Reload
	movq	%rbx, 48(%r9)
	adcq	%r14, %r10
	movq	%r10, 56(%r9)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end123:
	.size	mcl_fp_subNF8Lbmi2, .Lfunc_end123-mcl_fp_subNF8Lbmi2

	.globl	mcl_fpDbl_add8Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_add8Lbmi2,@function
mcl_fpDbl_add8Lbmi2:                    # @mcl_fpDbl_add8Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rcx, %r8
	movq	120(%rdx), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	movq	112(%rdx), %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	movq	104(%rdx), %rax
	movq	%rax, -32(%rsp)         # 8-byte Spill
	movq	96(%rdx), %r14
	movq	24(%rsi), %r15
	movq	32(%rsi), %r11
	movq	16(%rdx), %r12
	movq	(%rdx), %rbx
	movq	8(%rdx), %rax
	addq	(%rsi), %rbx
	adcq	8(%rsi), %rax
	adcq	16(%rsi), %r12
	adcq	24(%rdx), %r15
	adcq	32(%rdx), %r11
	movq	88(%rdx), %rbp
	movq	80(%rdx), %r13
	movq	%rbx, (%rdi)
	movq	72(%rdx), %r10
	movq	%rax, 8(%rdi)
	movq	64(%rdx), %r9
	movq	%r12, 16(%rdi)
	movq	40(%rdx), %r12
	movq	%r15, 24(%rdi)
	movq	40(%rsi), %rbx
	adcq	%r12, %rbx
	movq	56(%rdx), %r15
	movq	48(%rdx), %r12
	movq	%r11, 32(%rdi)
	movq	48(%rsi), %rdx
	adcq	%r12, %rdx
	movq	120(%rsi), %r12
	movq	%rbx, 40(%rdi)
	movq	56(%rsi), %rax
	adcq	%r15, %rax
	movq	112(%rsi), %rcx
	movq	%rdx, 48(%rdi)
	movq	64(%rsi), %rbx
	adcq	%r9, %rbx
	movq	104(%rsi), %rdx
	movq	%rax, 56(%rdi)
	movq	72(%rsi), %r9
	adcq	%r10, %r9
	movq	80(%rsi), %r11
	adcq	%r13, %r11
	movq	96(%rsi), %rax
	movq	88(%rsi), %r15
	adcq	%rbp, %r15
	adcq	%r14, %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	movq	%rdx, %rax
	adcq	-32(%rsp), %rax         # 8-byte Folded Reload
	movq	%rax, -32(%rsp)         # 8-byte Spill
	adcq	-24(%rsp), %rcx         # 8-byte Folded Reload
	movq	%rcx, -24(%rsp)         # 8-byte Spill
	adcq	-8(%rsp), %r12          # 8-byte Folded Reload
	movq	%r12, -8(%rsp)          # 8-byte Spill
	sbbq	%rbp, %rbp
	andl	$1, %ebp
	movq	%rbx, %rsi
	subq	(%r8), %rsi
	movq	%r9, %rdx
	sbbq	8(%r8), %rdx
	movq	%r11, %r10
	sbbq	16(%r8), %r10
	movq	%r15, %r14
	sbbq	24(%r8), %r14
	movq	-16(%rsp), %r13         # 8-byte Reload
	sbbq	32(%r8), %r13
	movq	%rax, %r12
	sbbq	40(%r8), %r12
	movq	%rcx, %rax
	sbbq	48(%r8), %rax
	movq	-8(%rsp), %rcx          # 8-byte Reload
	sbbq	56(%r8), %rcx
	sbbq	$0, %rbp
	andl	$1, %ebp
	cmovneq	%rbx, %rsi
	movq	%rsi, 64(%rdi)
	testb	%bpl, %bpl
	cmovneq	%r9, %rdx
	movq	%rdx, 72(%rdi)
	cmovneq	%r11, %r10
	movq	%r10, 80(%rdi)
	cmovneq	%r15, %r14
	movq	%r14, 88(%rdi)
	cmovneq	-16(%rsp), %r13         # 8-byte Folded Reload
	movq	%r13, 96(%rdi)
	cmovneq	-32(%rsp), %r12         # 8-byte Folded Reload
	movq	%r12, 104(%rdi)
	cmovneq	-24(%rsp), %rax         # 8-byte Folded Reload
	movq	%rax, 112(%rdi)
	cmovneq	-8(%rsp), %rcx          # 8-byte Folded Reload
	movq	%rcx, 120(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end124:
	.size	mcl_fpDbl_add8Lbmi2, .Lfunc_end124-mcl_fpDbl_add8Lbmi2

	.globl	mcl_fpDbl_sub8Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sub8Lbmi2,@function
mcl_fpDbl_sub8Lbmi2:                    # @mcl_fpDbl_sub8Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rcx, %r15
	movq	120(%rdx), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	movq	112(%rdx), %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	movq	104(%rdx), %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	movq	16(%rsi), %r9
	movq	(%rsi), %r12
	movq	8(%rsi), %r14
	xorl	%r8d, %r8d
	subq	(%rdx), %r12
	sbbq	8(%rdx), %r14
	sbbq	16(%rdx), %r9
	movq	24(%rsi), %rbx
	sbbq	24(%rdx), %rbx
	movq	32(%rsi), %r13
	sbbq	32(%rdx), %r13
	movq	96(%rdx), %rbp
	movq	88(%rdx), %r11
	movq	%r12, (%rdi)
	movq	80(%rdx), %r12
	movq	%r14, 8(%rdi)
	movq	72(%rdx), %r10
	movq	%r9, 16(%rdi)
	movq	40(%rdx), %r9
	movq	%rbx, 24(%rdi)
	movq	40(%rsi), %rbx
	sbbq	%r9, %rbx
	movq	48(%rdx), %r9
	movq	%r13, 32(%rdi)
	movq	48(%rsi), %r14
	sbbq	%r9, %r14
	movq	64(%rdx), %r13
	movq	56(%rdx), %r9
	movq	%rbx, 40(%rdi)
	movq	56(%rsi), %rdx
	sbbq	%r9, %rdx
	movq	120(%rsi), %rcx
	movq	%r14, 48(%rdi)
	movq	64(%rsi), %rbx
	sbbq	%r13, %rbx
	movq	112(%rsi), %rax
	movq	%rdx, 56(%rdi)
	movq	72(%rsi), %r9
	sbbq	%r10, %r9
	movq	80(%rsi), %r13
	sbbq	%r12, %r13
	movq	88(%rsi), %r12
	sbbq	%r11, %r12
	movq	104(%rsi), %rdx
	movq	96(%rsi), %r14
	sbbq	%rbp, %r14
	sbbq	-24(%rsp), %rdx         # 8-byte Folded Reload
	movq	%rdx, -24(%rsp)         # 8-byte Spill
	sbbq	-16(%rsp), %rax         # 8-byte Folded Reload
	movq	%rax, -16(%rsp)         # 8-byte Spill
	sbbq	-8(%rsp), %rcx          # 8-byte Folded Reload
	movq	%rcx, -8(%rsp)          # 8-byte Spill
	movl	$0, %ebp
	sbbq	$0, %rbp
	andl	$1, %ebp
	movq	(%r15), %r11
	cmoveq	%r8, %r11
	testb	%bpl, %bpl
	movq	16(%r15), %rbp
	cmoveq	%r8, %rbp
	movq	8(%r15), %rsi
	cmoveq	%r8, %rsi
	movq	56(%r15), %r10
	cmoveq	%r8, %r10
	movq	48(%r15), %rdx
	cmoveq	%r8, %rdx
	movq	40(%r15), %rcx
	cmoveq	%r8, %rcx
	movq	32(%r15), %rax
	cmoveq	%r8, %rax
	cmovneq	24(%r15), %r8
	addq	%rbx, %r11
	adcq	%r9, %rsi
	movq	%r11, 64(%rdi)
	adcq	%r13, %rbp
	movq	%rsi, 72(%rdi)
	movq	%rbp, 80(%rdi)
	adcq	%r12, %r8
	movq	%r8, 88(%rdi)
	adcq	%r14, %rax
	movq	%rax, 96(%rdi)
	adcq	-24(%rsp), %rcx         # 8-byte Folded Reload
	movq	%rcx, 104(%rdi)
	adcq	-16(%rsp), %rdx         # 8-byte Folded Reload
	movq	%rdx, 112(%rdi)
	adcq	-8(%rsp), %r10          # 8-byte Folded Reload
	movq	%r10, 120(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end125:
	.size	mcl_fpDbl_sub8Lbmi2, .Lfunc_end125-mcl_fpDbl_sub8Lbmi2

	.align	16, 0x90
	.type	.LmulPv576x64,@function
.LmulPv576x64:                          # @mulPv576x64
# BB#0:
	mulxq	(%rsi), %rcx, %rax
	movq	%rcx, (%rdi)
	mulxq	8(%rsi), %rcx, %r8
	addq	%rax, %rcx
	movq	%rcx, 8(%rdi)
	mulxq	16(%rsi), %rcx, %r9
	adcq	%r8, %rcx
	movq	%rcx, 16(%rdi)
	mulxq	24(%rsi), %rax, %rcx
	adcq	%r9, %rax
	movq	%rax, 24(%rdi)
	mulxq	32(%rsi), %rax, %r8
	adcq	%rcx, %rax
	movq	%rax, 32(%rdi)
	mulxq	40(%rsi), %rcx, %r9
	adcq	%r8, %rcx
	movq	%rcx, 40(%rdi)
	mulxq	48(%rsi), %rax, %rcx
	adcq	%r9, %rax
	movq	%rax, 48(%rdi)
	mulxq	56(%rsi), %rax, %r8
	adcq	%rcx, %rax
	movq	%rax, 56(%rdi)
	mulxq	64(%rsi), %rax, %rcx
	adcq	%r8, %rax
	movq	%rax, 64(%rdi)
	adcq	$0, %rcx
	movq	%rcx, 72(%rdi)
	movq	%rdi, %rax
	retq
.Lfunc_end126:
	.size	.LmulPv576x64, .Lfunc_end126-.LmulPv576x64

	.globl	mcl_fp_mulUnitPre9Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mulUnitPre9Lbmi2,@function
mcl_fp_mulUnitPre9Lbmi2:                # @mcl_fp_mulUnitPre9Lbmi2
# BB#0:
	pushq	%r14
	pushq	%rbx
	subq	$88, %rsp
	movq	%rdi, %rbx
	leaq	8(%rsp), %rdi
	callq	.LmulPv576x64
	movq	80(%rsp), %r8
	movq	72(%rsp), %r9
	movq	64(%rsp), %r10
	movq	56(%rsp), %r11
	movq	48(%rsp), %r14
	movq	40(%rsp), %rax
	movq	32(%rsp), %rcx
	movq	24(%rsp), %rdx
	movq	8(%rsp), %rsi
	movq	16(%rsp), %rdi
	movq	%rsi, (%rbx)
	movq	%rdi, 8(%rbx)
	movq	%rdx, 16(%rbx)
	movq	%rcx, 24(%rbx)
	movq	%rax, 32(%rbx)
	movq	%r14, 40(%rbx)
	movq	%r11, 48(%rbx)
	movq	%r10, 56(%rbx)
	movq	%r9, 64(%rbx)
	movq	%r8, 72(%rbx)
	addq	$88, %rsp
	popq	%rbx
	popq	%r14
	retq
.Lfunc_end127:
	.size	mcl_fp_mulUnitPre9Lbmi2, .Lfunc_end127-mcl_fp_mulUnitPre9Lbmi2

	.globl	mcl_fpDbl_mulPre9Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_mulPre9Lbmi2,@function
mcl_fpDbl_mulPre9Lbmi2:                 # @mcl_fpDbl_mulPre9Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	subq	$808, %rsp              # imm = 0x328
	movq	%rdx, %rax
	movq	%rax, 64(%rsp)          # 8-byte Spill
	movq	%rsi, 72(%rsp)          # 8-byte Spill
	movq	%rdi, %r12
	movq	%r12, 80(%rsp)          # 8-byte Spill
	movq	(%rax), %rdx
	movq	%rax, %rbx
	leaq	728(%rsp), %rdi
	movq	%rsi, %rbp
	callq	.LmulPv576x64
	movq	800(%rsp), %r13
	movq	792(%rsp), %rax
	movq	%rax, 56(%rsp)          # 8-byte Spill
	movq	784(%rsp), %rax
	movq	%rax, 48(%rsp)          # 8-byte Spill
	movq	776(%rsp), %rax
	movq	%rax, 40(%rsp)          # 8-byte Spill
	movq	768(%rsp), %rax
	movq	%rax, 32(%rsp)          # 8-byte Spill
	movq	760(%rsp), %rax
	movq	%rax, 24(%rsp)          # 8-byte Spill
	movq	752(%rsp), %rax
	movq	%rax, 16(%rsp)          # 8-byte Spill
	movq	744(%rsp), %rax
	movq	%rax, 8(%rsp)           # 8-byte Spill
	movq	728(%rsp), %rax
	movq	736(%rsp), %r14
	movq	%rax, (%r12)
	movq	8(%rbx), %rdx
	leaq	648(%rsp), %rdi
	movq	%rbp, %rsi
	callq	.LmulPv576x64
	movq	720(%rsp), %r8
	movq	712(%rsp), %rcx
	movq	704(%rsp), %rdx
	movq	696(%rsp), %rsi
	movq	688(%rsp), %rdi
	movq	680(%rsp), %rbp
	addq	648(%rsp), %r14
	movq	672(%rsp), %rax
	movq	656(%rsp), %rbx
	movq	664(%rsp), %r15
	movq	%r14, 8(%r12)
	adcq	8(%rsp), %rbx           # 8-byte Folded Reload
	adcq	16(%rsp), %r15          # 8-byte Folded Reload
	adcq	24(%rsp), %rax          # 8-byte Folded Reload
	movq	%rax, %r14
	adcq	32(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rbp, 16(%rsp)          # 8-byte Spill
	adcq	40(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 24(%rsp)          # 8-byte Spill
	adcq	48(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 32(%rsp)          # 8-byte Spill
	adcq	56(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, 40(%rsp)          # 8-byte Spill
	adcq	%r13, %rcx
	movq	%rcx, 48(%rsp)          # 8-byte Spill
	adcq	$0, %r8
	movq	%r8, 56(%rsp)           # 8-byte Spill
	movq	64(%rsp), %r13          # 8-byte Reload
	movq	16(%r13), %rdx
	leaq	568(%rsp), %rdi
	movq	72(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	640(%rsp), %r8
	movq	632(%rsp), %r9
	movq	624(%rsp), %r10
	movq	616(%rsp), %rdi
	movq	608(%rsp), %rbp
	movq	600(%rsp), %rcx
	addq	568(%rsp), %rbx
	movq	592(%rsp), %rdx
	movq	576(%rsp), %r12
	movq	584(%rsp), %rsi
	movq	80(%rsp), %rax          # 8-byte Reload
	movq	%rbx, 16(%rax)
	adcq	%r15, %r12
	adcq	%r14, %rsi
	movq	%rsi, (%rsp)            # 8-byte Spill
	adcq	16(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, 8(%rsp)           # 8-byte Spill
	adcq	24(%rsp), %rcx          # 8-byte Folded Reload
	movq	%rcx, 16(%rsp)          # 8-byte Spill
	adcq	32(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rbp, 24(%rsp)          # 8-byte Spill
	adcq	40(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 32(%rsp)          # 8-byte Spill
	adcq	48(%rsp), %r10          # 8-byte Folded Reload
	movq	%r10, 40(%rsp)          # 8-byte Spill
	adcq	56(%rsp), %r9           # 8-byte Folded Reload
	movq	%r9, 48(%rsp)           # 8-byte Spill
	adcq	$0, %r8
	movq	%r8, 56(%rsp)           # 8-byte Spill
	movq	24(%r13), %rdx
	leaq	488(%rsp), %rdi
	movq	72(%rsp), %r15          # 8-byte Reload
	movq	%r15, %rsi
	callq	.LmulPv576x64
	movq	560(%rsp), %r8
	movq	552(%rsp), %rcx
	movq	544(%rsp), %rdx
	movq	536(%rsp), %rsi
	movq	528(%rsp), %rdi
	movq	520(%rsp), %rbp
	addq	488(%rsp), %r12
	movq	512(%rsp), %rax
	movq	496(%rsp), %rbx
	movq	504(%rsp), %r13
	movq	80(%rsp), %r14          # 8-byte Reload
	movq	%r12, 24(%r14)
	adcq	(%rsp), %rbx            # 8-byte Folded Reload
	adcq	8(%rsp), %r13           # 8-byte Folded Reload
	adcq	16(%rsp), %rax          # 8-byte Folded Reload
	movq	%rax, 8(%rsp)           # 8-byte Spill
	adcq	24(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rbp, 16(%rsp)          # 8-byte Spill
	adcq	32(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 24(%rsp)          # 8-byte Spill
	adcq	40(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 32(%rsp)          # 8-byte Spill
	adcq	48(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, 40(%rsp)          # 8-byte Spill
	adcq	56(%rsp), %rcx          # 8-byte Folded Reload
	movq	%rcx, 48(%rsp)          # 8-byte Spill
	adcq	$0, %r8
	movq	%r8, 56(%rsp)           # 8-byte Spill
	movq	64(%rsp), %r12          # 8-byte Reload
	movq	32(%r12), %rdx
	leaq	408(%rsp), %rdi
	movq	%r15, %rsi
	callq	.LmulPv576x64
	movq	480(%rsp), %r8
	movq	472(%rsp), %r9
	movq	464(%rsp), %rdx
	movq	456(%rsp), %rsi
	movq	448(%rsp), %rdi
	movq	440(%rsp), %rbp
	addq	408(%rsp), %rbx
	movq	432(%rsp), %rax
	movq	416(%rsp), %r15
	movq	424(%rsp), %rcx
	movq	%rbx, 32(%r14)
	adcq	%r13, %r15
	adcq	8(%rsp), %rcx           # 8-byte Folded Reload
	movq	%rcx, (%rsp)            # 8-byte Spill
	adcq	16(%rsp), %rax          # 8-byte Folded Reload
	movq	%rax, 8(%rsp)           # 8-byte Spill
	adcq	24(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rbp, 16(%rsp)          # 8-byte Spill
	adcq	32(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 24(%rsp)          # 8-byte Spill
	adcq	40(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 32(%rsp)          # 8-byte Spill
	adcq	48(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, 40(%rsp)          # 8-byte Spill
	adcq	56(%rsp), %r9           # 8-byte Folded Reload
	movq	%r9, 48(%rsp)           # 8-byte Spill
	adcq	$0, %r8
	movq	%r8, 56(%rsp)           # 8-byte Spill
	movq	%r12, %r14
	movq	40(%r14), %rdx
	leaq	328(%rsp), %rdi
	movq	72(%rsp), %r13          # 8-byte Reload
	movq	%r13, %rsi
	callq	.LmulPv576x64
	movq	400(%rsp), %r8
	movq	392(%rsp), %r9
	movq	384(%rsp), %rsi
	movq	376(%rsp), %rdi
	movq	368(%rsp), %rbx
	movq	360(%rsp), %rbp
	addq	328(%rsp), %r15
	movq	352(%rsp), %rcx
	movq	336(%rsp), %r12
	movq	344(%rsp), %rdx
	movq	80(%rsp), %rax          # 8-byte Reload
	movq	%r15, 40(%rax)
	adcq	(%rsp), %r12            # 8-byte Folded Reload
	adcq	8(%rsp), %rdx           # 8-byte Folded Reload
	movq	%rdx, (%rsp)            # 8-byte Spill
	adcq	16(%rsp), %rcx          # 8-byte Folded Reload
	movq	%rcx, 8(%rsp)           # 8-byte Spill
	adcq	24(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rbp, 16(%rsp)          # 8-byte Spill
	adcq	32(%rsp), %rbx          # 8-byte Folded Reload
	movq	%rbx, 24(%rsp)          # 8-byte Spill
	adcq	40(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 32(%rsp)          # 8-byte Spill
	adcq	48(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 40(%rsp)          # 8-byte Spill
	adcq	56(%rsp), %r9           # 8-byte Folded Reload
	movq	%r9, 48(%rsp)           # 8-byte Spill
	adcq	$0, %r8
	movq	%r8, 56(%rsp)           # 8-byte Spill
	movq	48(%r14), %rdx
	leaq	248(%rsp), %rdi
	movq	%r13, %rsi
	movq	%r13, %r15
	callq	.LmulPv576x64
	movq	320(%rsp), %r8
	movq	312(%rsp), %r9
	movq	304(%rsp), %rsi
	movq	296(%rsp), %rdi
	movq	288(%rsp), %rbx
	movq	280(%rsp), %rbp
	addq	248(%rsp), %r12
	movq	272(%rsp), %rcx
	movq	256(%rsp), %r13
	movq	264(%rsp), %rdx
	movq	80(%rsp), %rax          # 8-byte Reload
	movq	%r12, 48(%rax)
	adcq	(%rsp), %r13            # 8-byte Folded Reload
	adcq	8(%rsp), %rdx           # 8-byte Folded Reload
	movq	%rdx, (%rsp)            # 8-byte Spill
	adcq	16(%rsp), %rcx          # 8-byte Folded Reload
	movq	%rcx, 8(%rsp)           # 8-byte Spill
	adcq	24(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rbp, 16(%rsp)          # 8-byte Spill
	adcq	32(%rsp), %rbx          # 8-byte Folded Reload
	movq	%rbx, 24(%rsp)          # 8-byte Spill
	adcq	40(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 32(%rsp)          # 8-byte Spill
	adcq	48(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 40(%rsp)          # 8-byte Spill
	adcq	56(%rsp), %r9           # 8-byte Folded Reload
	movq	%r9, 48(%rsp)           # 8-byte Spill
	adcq	$0, %r8
	movq	%r8, 56(%rsp)           # 8-byte Spill
	movq	56(%r14), %rdx
	leaq	168(%rsp), %rdi
	movq	%r15, %rsi
	callq	.LmulPv576x64
	movq	240(%rsp), %rcx
	movq	232(%rsp), %rdx
	movq	224(%rsp), %rsi
	movq	216(%rsp), %rdi
	movq	208(%rsp), %rbx
	addq	168(%rsp), %r13
	movq	200(%rsp), %r12
	movq	192(%rsp), %rbp
	movq	176(%rsp), %r14
	movq	184(%rsp), %r15
	movq	80(%rsp), %rax          # 8-byte Reload
	movq	%r13, 56(%rax)
	adcq	(%rsp), %r14            # 8-byte Folded Reload
	adcq	8(%rsp), %r15           # 8-byte Folded Reload
	adcq	16(%rsp), %rbp          # 8-byte Folded Reload
	adcq	24(%rsp), %r12          # 8-byte Folded Reload
	adcq	32(%rsp), %rbx          # 8-byte Folded Reload
	movq	%rbx, %r13
	adcq	40(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 32(%rsp)          # 8-byte Spill
	adcq	48(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 40(%rsp)          # 8-byte Spill
	adcq	56(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, 48(%rsp)          # 8-byte Spill
	adcq	$0, %rcx
	movq	%rcx, 56(%rsp)          # 8-byte Spill
	movq	64(%rsp), %rax          # 8-byte Reload
	movq	64(%rax), %rdx
	leaq	88(%rsp), %rdi
	movq	72(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	addq	88(%rsp), %r14
	adcq	96(%rsp), %r15
	movq	160(%rsp), %r8
	adcq	104(%rsp), %rbp
	movq	152(%rsp), %r9
	movq	144(%rsp), %rdx
	movq	136(%rsp), %rsi
	movq	128(%rsp), %rdi
	movq	120(%rsp), %rbx
	movq	112(%rsp), %rax
	movq	80(%rsp), %rcx          # 8-byte Reload
	movq	%r14, 64(%rcx)
	movq	%r15, 72(%rcx)
	adcq	%r12, %rax
	movq	%rbp, 80(%rcx)
	movq	%rax, 88(%rcx)
	adcq	%r13, %rbx
	movq	%rbx, 96(%rcx)
	adcq	32(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 104(%rcx)
	adcq	40(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 112(%rcx)
	adcq	48(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, 120(%rcx)
	adcq	56(%rsp), %r9           # 8-byte Folded Reload
	movq	%r9, 128(%rcx)
	adcq	$0, %r8
	movq	%r8, 136(%rcx)
	addq	$808, %rsp              # imm = 0x328
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end128:
	.size	mcl_fpDbl_mulPre9Lbmi2, .Lfunc_end128-mcl_fpDbl_mulPre9Lbmi2

	.globl	mcl_fpDbl_sqrPre9Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sqrPre9Lbmi2,@function
mcl_fpDbl_sqrPre9Lbmi2:                 # @mcl_fpDbl_sqrPre9Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	subq	$808, %rsp              # imm = 0x328
	movq	%rsi, %r15
	movq	%r15, 80(%rsp)          # 8-byte Spill
	movq	%rdi, %r14
	movq	%r14, 72(%rsp)          # 8-byte Spill
	movq	(%r15), %rdx
	leaq	728(%rsp), %rdi
	callq	.LmulPv576x64
	movq	800(%rsp), %rax
	movq	%rax, 64(%rsp)          # 8-byte Spill
	movq	792(%rsp), %rax
	movq	%rax, 56(%rsp)          # 8-byte Spill
	movq	784(%rsp), %rax
	movq	%rax, 48(%rsp)          # 8-byte Spill
	movq	776(%rsp), %rax
	movq	%rax, 40(%rsp)          # 8-byte Spill
	movq	768(%rsp), %rax
	movq	%rax, 32(%rsp)          # 8-byte Spill
	movq	760(%rsp), %rax
	movq	%rax, 24(%rsp)          # 8-byte Spill
	movq	752(%rsp), %rax
	movq	%rax, 16(%rsp)          # 8-byte Spill
	movq	744(%rsp), %rax
	movq	%rax, 8(%rsp)           # 8-byte Spill
	movq	728(%rsp), %rax
	movq	736(%rsp), %r12
	movq	%rax, (%r14)
	movq	8(%r15), %rdx
	leaq	648(%rsp), %rdi
	movq	%r15, %rsi
	callq	.LmulPv576x64
	movq	720(%rsp), %r8
	movq	712(%rsp), %rcx
	movq	704(%rsp), %rdx
	movq	696(%rsp), %rsi
	movq	688(%rsp), %rdi
	movq	680(%rsp), %rbp
	addq	648(%rsp), %r12
	movq	672(%rsp), %rax
	movq	656(%rsp), %rbx
	movq	664(%rsp), %r13
	movq	%r12, 8(%r14)
	adcq	8(%rsp), %rbx           # 8-byte Folded Reload
	adcq	16(%rsp), %r13          # 8-byte Folded Reload
	adcq	24(%rsp), %rax          # 8-byte Folded Reload
	movq	%rax, 16(%rsp)          # 8-byte Spill
	adcq	32(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rbp, 24(%rsp)          # 8-byte Spill
	adcq	40(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 32(%rsp)          # 8-byte Spill
	adcq	48(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 40(%rsp)          # 8-byte Spill
	adcq	56(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, 48(%rsp)          # 8-byte Spill
	adcq	64(%rsp), %rcx          # 8-byte Folded Reload
	movq	%rcx, 56(%rsp)          # 8-byte Spill
	adcq	$0, %r8
	movq	%r8, 64(%rsp)           # 8-byte Spill
	movq	16(%r15), %rdx
	leaq	568(%rsp), %rdi
	movq	%r15, %rsi
	callq	.LmulPv576x64
	movq	640(%rsp), %r8
	movq	632(%rsp), %rcx
	movq	624(%rsp), %rdx
	movq	616(%rsp), %rsi
	movq	608(%rsp), %rdi
	movq	600(%rsp), %rbp
	addq	568(%rsp), %rbx
	movq	592(%rsp), %rax
	movq	576(%rsp), %r14
	movq	584(%rsp), %r12
	movq	72(%rsp), %r15          # 8-byte Reload
	movq	%rbx, 16(%r15)
	adcq	%r13, %r14
	adcq	16(%rsp), %r12          # 8-byte Folded Reload
	adcq	24(%rsp), %rax          # 8-byte Folded Reload
	movq	%rax, 16(%rsp)          # 8-byte Spill
	adcq	32(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rbp, 24(%rsp)          # 8-byte Spill
	adcq	40(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 32(%rsp)          # 8-byte Spill
	adcq	48(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 40(%rsp)          # 8-byte Spill
	adcq	56(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, 48(%rsp)          # 8-byte Spill
	adcq	64(%rsp), %rcx          # 8-byte Folded Reload
	movq	%rcx, 56(%rsp)          # 8-byte Spill
	adcq	$0, %r8
	movq	%r8, 64(%rsp)           # 8-byte Spill
	movq	80(%rsp), %rsi          # 8-byte Reload
	movq	24(%rsi), %rdx
	leaq	488(%rsp), %rdi
	callq	.LmulPv576x64
	movq	560(%rsp), %r8
	movq	552(%rsp), %rcx
	movq	544(%rsp), %rdx
	movq	536(%rsp), %rsi
	movq	528(%rsp), %rdi
	movq	520(%rsp), %rbp
	addq	488(%rsp), %r14
	movq	512(%rsp), %rax
	movq	496(%rsp), %rbx
	movq	504(%rsp), %r13
	movq	%r14, 24(%r15)
	adcq	%r12, %rbx
	adcq	16(%rsp), %r13          # 8-byte Folded Reload
	adcq	24(%rsp), %rax          # 8-byte Folded Reload
	movq	%rax, 16(%rsp)          # 8-byte Spill
	adcq	32(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rbp, 24(%rsp)          # 8-byte Spill
	adcq	40(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 32(%rsp)          # 8-byte Spill
	adcq	48(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 40(%rsp)          # 8-byte Spill
	adcq	56(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, 48(%rsp)          # 8-byte Spill
	adcq	64(%rsp), %rcx          # 8-byte Folded Reload
	movq	%rcx, 56(%rsp)          # 8-byte Spill
	adcq	$0, %r8
	movq	%r8, 64(%rsp)           # 8-byte Spill
	movq	80(%rsp), %rsi          # 8-byte Reload
	movq	32(%rsi), %rdx
	leaq	408(%rsp), %rdi
	callq	.LmulPv576x64
	movq	480(%rsp), %r8
	movq	472(%rsp), %rcx
	movq	464(%rsp), %rdx
	movq	456(%rsp), %rsi
	movq	448(%rsp), %rdi
	movq	440(%rsp), %rbp
	addq	408(%rsp), %rbx
	movq	432(%rsp), %rax
	movq	416(%rsp), %r14
	movq	424(%rsp), %r12
	movq	%rbx, 32(%r15)
	adcq	%r13, %r14
	adcq	16(%rsp), %r12          # 8-byte Folded Reload
	adcq	24(%rsp), %rax          # 8-byte Folded Reload
	movq	%rax, 16(%rsp)          # 8-byte Spill
	adcq	32(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rbp, 24(%rsp)          # 8-byte Spill
	adcq	40(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 32(%rsp)          # 8-byte Spill
	adcq	48(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 40(%rsp)          # 8-byte Spill
	adcq	56(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, 48(%rsp)          # 8-byte Spill
	adcq	64(%rsp), %rcx          # 8-byte Folded Reload
	movq	%rcx, 56(%rsp)          # 8-byte Spill
	adcq	$0, %r8
	movq	%r8, 64(%rsp)           # 8-byte Spill
	movq	80(%rsp), %rsi          # 8-byte Reload
	movq	40(%rsi), %rdx
	leaq	328(%rsp), %rdi
	callq	.LmulPv576x64
	movq	400(%rsp), %r8
	movq	392(%rsp), %rcx
	movq	384(%rsp), %rdx
	movq	376(%rsp), %rsi
	movq	368(%rsp), %rdi
	movq	360(%rsp), %rbp
	addq	328(%rsp), %r14
	movq	352(%rsp), %rax
	movq	336(%rsp), %rbx
	movq	344(%rsp), %r13
	movq	%r14, 40(%r15)
	adcq	%r12, %rbx
	adcq	16(%rsp), %r13          # 8-byte Folded Reload
	adcq	24(%rsp), %rax          # 8-byte Folded Reload
	movq	%rax, 16(%rsp)          # 8-byte Spill
	adcq	32(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rbp, 24(%rsp)          # 8-byte Spill
	adcq	40(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 32(%rsp)          # 8-byte Spill
	adcq	48(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 40(%rsp)          # 8-byte Spill
	adcq	56(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, 48(%rsp)          # 8-byte Spill
	adcq	64(%rsp), %rcx          # 8-byte Folded Reload
	movq	%rcx, 56(%rsp)          # 8-byte Spill
	adcq	$0, %r8
	movq	%r8, 64(%rsp)           # 8-byte Spill
	movq	80(%rsp), %rsi          # 8-byte Reload
	movq	48(%rsi), %rdx
	leaq	248(%rsp), %rdi
	callq	.LmulPv576x64
	movq	320(%rsp), %r8
	movq	312(%rsp), %rcx
	movq	304(%rsp), %rdx
	movq	296(%rsp), %rsi
	movq	288(%rsp), %rdi
	movq	280(%rsp), %rbp
	addq	248(%rsp), %rbx
	movq	272(%rsp), %rax
	movq	256(%rsp), %r12
	movq	264(%rsp), %r14
	movq	%rbx, 48(%r15)
	adcq	%r13, %r12
	adcq	16(%rsp), %r14          # 8-byte Folded Reload
	adcq	24(%rsp), %rax          # 8-byte Folded Reload
	movq	%rax, 16(%rsp)          # 8-byte Spill
	adcq	32(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rbp, 24(%rsp)          # 8-byte Spill
	adcq	40(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 32(%rsp)          # 8-byte Spill
	adcq	48(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 40(%rsp)          # 8-byte Spill
	adcq	56(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, 48(%rsp)          # 8-byte Spill
	adcq	64(%rsp), %rcx          # 8-byte Folded Reload
	movq	%rcx, 56(%rsp)          # 8-byte Spill
	adcq	$0, %r8
	movq	%r8, 64(%rsp)           # 8-byte Spill
	movq	80(%rsp), %rsi          # 8-byte Reload
	movq	56(%rsi), %rdx
	leaq	168(%rsp), %rdi
	callq	.LmulPv576x64
	movq	240(%rsp), %r8
	movq	232(%rsp), %rdx
	movq	224(%rsp), %rsi
	movq	216(%rsp), %rdi
	movq	208(%rsp), %rbx
	movq	200(%rsp), %rcx
	addq	168(%rsp), %r12
	movq	192(%rsp), %r15
	movq	176(%rsp), %r13
	movq	184(%rsp), %rbp
	movq	72(%rsp), %rax          # 8-byte Reload
	movq	%r12, 56(%rax)
	adcq	%r14, %r13
	adcq	16(%rsp), %rbp          # 8-byte Folded Reload
	adcq	24(%rsp), %r15          # 8-byte Folded Reload
	adcq	32(%rsp), %rcx          # 8-byte Folded Reload
	movq	%rcx, %r12
	adcq	40(%rsp), %rbx          # 8-byte Folded Reload
	movq	%rbx, %r14
	adcq	48(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 40(%rsp)          # 8-byte Spill
	adcq	56(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 48(%rsp)          # 8-byte Spill
	adcq	64(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, 56(%rsp)          # 8-byte Spill
	adcq	$0, %r8
	movq	%r8, 64(%rsp)           # 8-byte Spill
	movq	80(%rsp), %rsi          # 8-byte Reload
	movq	64(%rsi), %rdx
	leaq	88(%rsp), %rdi
	callq	.LmulPv576x64
	addq	88(%rsp), %r13
	adcq	96(%rsp), %rbp
	movq	160(%rsp), %r8
	adcq	104(%rsp), %r15
	movq	152(%rsp), %r9
	movq	144(%rsp), %rdx
	movq	136(%rsp), %rsi
	movq	128(%rsp), %rdi
	movq	120(%rsp), %rbx
	movq	112(%rsp), %rax
	movq	72(%rsp), %rcx          # 8-byte Reload
	movq	%r13, 64(%rcx)
	movq	%rbp, 72(%rcx)
	adcq	%r12, %rax
	movq	%r15, 80(%rcx)
	movq	%rax, 88(%rcx)
	adcq	%r14, %rbx
	movq	%rbx, 96(%rcx)
	adcq	40(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 104(%rcx)
	adcq	48(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rsi, 112(%rcx)
	adcq	56(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, 120(%rcx)
	adcq	64(%rsp), %r9           # 8-byte Folded Reload
	movq	%r9, 128(%rcx)
	adcq	$0, %r8
	movq	%r8, 136(%rcx)
	addq	$808, %rsp              # imm = 0x328
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end129:
	.size	mcl_fpDbl_sqrPre9Lbmi2, .Lfunc_end129-mcl_fpDbl_sqrPre9Lbmi2

	.globl	mcl_fp_mont9Lbmi2
	.align	16, 0x90
	.type	mcl_fp_mont9Lbmi2,@function
mcl_fp_mont9Lbmi2:                      # @mcl_fp_mont9Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	subq	$1560, %rsp             # imm = 0x618
	movq	%rcx, 56(%rsp)          # 8-byte Spill
	movq	%rdx, 32(%rsp)          # 8-byte Spill
	movq	%rsi, 24(%rsp)          # 8-byte Spill
	movq	%rdi, (%rsp)            # 8-byte Spill
	movq	-8(%rcx), %rbx
	movq	%rbx, 16(%rsp)          # 8-byte Spill
	movq	(%rdx), %rdx
	leaq	1480(%rsp), %rdi
	callq	.LmulPv576x64
	movq	1480(%rsp), %r14
	movq	1488(%rsp), %r15
	movq	%r14, %rdx
	imulq	%rbx, %rdx
	movq	1552(%rsp), %rax
	movq	%rax, 112(%rsp)         # 8-byte Spill
	movq	1544(%rsp), %rax
	movq	%rax, 104(%rsp)         # 8-byte Spill
	movq	1536(%rsp), %rax
	movq	%rax, 88(%rsp)          # 8-byte Spill
	movq	1528(%rsp), %r12
	movq	1520(%rsp), %r13
	movq	1512(%rsp), %rbx
	movq	1504(%rsp), %rax
	movq	%rax, 80(%rsp)          # 8-byte Spill
	movq	1496(%rsp), %rbp
	leaq	1400(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	addq	1400(%rsp), %r14
	adcq	1408(%rsp), %r15
	adcq	1416(%rsp), %rbp
	movq	%rbp, 96(%rsp)          # 8-byte Spill
	movq	80(%rsp), %rax          # 8-byte Reload
	adcq	1424(%rsp), %rax
	movq	%rax, 80(%rsp)          # 8-byte Spill
	adcq	1432(%rsp), %rbx
	movq	%rbx, 40(%rsp)          # 8-byte Spill
	adcq	1440(%rsp), %r13
	movq	%r13, 64(%rsp)          # 8-byte Spill
	adcq	1448(%rsp), %r12
	movq	%r12, 48(%rsp)          # 8-byte Spill
	movq	88(%rsp), %rbx          # 8-byte Reload
	adcq	1456(%rsp), %rbx
	movq	104(%rsp), %r14         # 8-byte Reload
	adcq	1464(%rsp), %r14
	movq	112(%rsp), %r13         # 8-byte Reload
	adcq	1472(%rsp), %r13
	sbbq	%rbp, %rbp
	movq	32(%rsp), %rax          # 8-byte Reload
	movq	8(%rax), %rdx
	leaq	1320(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	andl	$1, %ebp
	addq	1320(%rsp), %r15
	movq	96(%rsp), %rax          # 8-byte Reload
	adcq	1328(%rsp), %rax
	movq	%rax, 96(%rsp)          # 8-byte Spill
	movq	80(%rsp), %rax          # 8-byte Reload
	adcq	1336(%rsp), %rax
	movq	%rax, 80(%rsp)          # 8-byte Spill
	movq	40(%rsp), %r12          # 8-byte Reload
	adcq	1344(%rsp), %r12
	movq	64(%rsp), %rax          # 8-byte Reload
	adcq	1352(%rsp), %rax
	movq	%rax, 64(%rsp)          # 8-byte Spill
	movq	48(%rsp), %rax          # 8-byte Reload
	adcq	1360(%rsp), %rax
	movq	%rax, 48(%rsp)          # 8-byte Spill
	adcq	1368(%rsp), %rbx
	adcq	1376(%rsp), %r14
	movq	%r14, 104(%rsp)         # 8-byte Spill
	adcq	1384(%rsp), %r13
	movq	%r13, 112(%rsp)         # 8-byte Spill
	adcq	1392(%rsp), %rbp
	sbbq	%r14, %r14
	movq	%r15, %rdx
	imulq	16(%rsp), %rdx          # 8-byte Folded Reload
	leaq	1240(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	%r14, %rax
	andl	$1, %eax
	addq	1240(%rsp), %r15
	movq	96(%rsp), %rcx          # 8-byte Reload
	adcq	1248(%rsp), %rcx
	movq	%rcx, 96(%rsp)          # 8-byte Spill
	movq	80(%rsp), %r14          # 8-byte Reload
	adcq	1256(%rsp), %r14
	adcq	1264(%rsp), %r12
	movq	%r12, 40(%rsp)          # 8-byte Spill
	movq	64(%rsp), %r12          # 8-byte Reload
	adcq	1272(%rsp), %r12
	movq	48(%rsp), %r13          # 8-byte Reload
	adcq	1280(%rsp), %r13
	adcq	1288(%rsp), %rbx
	movq	%rbx, 88(%rsp)          # 8-byte Spill
	movq	104(%rsp), %r15         # 8-byte Reload
	adcq	1296(%rsp), %r15
	movq	112(%rsp), %rbx         # 8-byte Reload
	adcq	1304(%rsp), %rbx
	adcq	1312(%rsp), %rbp
	adcq	$0, %rax
	movq	%rax, 72(%rsp)          # 8-byte Spill
	movq	32(%rsp), %rax          # 8-byte Reload
	movq	16(%rax), %rdx
	leaq	1160(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	96(%rsp), %rax          # 8-byte Reload
	addq	1160(%rsp), %rax
	adcq	1168(%rsp), %r14
	movq	%r14, 80(%rsp)          # 8-byte Spill
	movq	40(%rsp), %r14          # 8-byte Reload
	adcq	1176(%rsp), %r14
	adcq	1184(%rsp), %r12
	movq	%r12, 64(%rsp)          # 8-byte Spill
	movq	%r13, %r12
	adcq	1192(%rsp), %r12
	movq	88(%rsp), %rcx          # 8-byte Reload
	adcq	1200(%rsp), %rcx
	movq	%rcx, 88(%rsp)          # 8-byte Spill
	adcq	1208(%rsp), %r15
	movq	%r15, %r13
	adcq	1216(%rsp), %rbx
	movq	%rbx, 112(%rsp)         # 8-byte Spill
	adcq	1224(%rsp), %rbp
	movq	72(%rsp), %rcx          # 8-byte Reload
	adcq	1232(%rsp), %rcx
	movq	%rcx, 72(%rsp)          # 8-byte Spill
	sbbq	%r15, %r15
	movq	%rax, %rdx
	movq	%rax, %rbx
	imulq	16(%rsp), %rdx          # 8-byte Folded Reload
	leaq	1080(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	%r15, %rax
	andl	$1, %eax
	addq	1080(%rsp), %rbx
	movq	80(%rsp), %rcx          # 8-byte Reload
	adcq	1088(%rsp), %rcx
	movq	%rcx, 80(%rsp)          # 8-byte Spill
	movq	%r14, %r15
	adcq	1096(%rsp), %r15
	movq	64(%rsp), %r14          # 8-byte Reload
	adcq	1104(%rsp), %r14
	movq	%r12, %rbx
	adcq	1112(%rsp), %rbx
	movq	88(%rsp), %rcx          # 8-byte Reload
	adcq	1120(%rsp), %rcx
	movq	%rcx, 88(%rsp)          # 8-byte Spill
	adcq	1128(%rsp), %r13
	movq	%r13, 104(%rsp)         # 8-byte Spill
	movq	112(%rsp), %r13         # 8-byte Reload
	adcq	1136(%rsp), %r13
	adcq	1144(%rsp), %rbp
	movq	72(%rsp), %r12          # 8-byte Reload
	adcq	1152(%rsp), %r12
	adcq	$0, %rax
	movq	%rax, 96(%rsp)          # 8-byte Spill
	movq	32(%rsp), %rax          # 8-byte Reload
	movq	24(%rax), %rdx
	leaq	1000(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	80(%rsp), %rax          # 8-byte Reload
	addq	1000(%rsp), %rax
	adcq	1008(%rsp), %r15
	movq	%r15, 40(%rsp)          # 8-byte Spill
	adcq	1016(%rsp), %r14
	movq	%r14, %r15
	adcq	1024(%rsp), %rbx
	movq	%rbx, 48(%rsp)          # 8-byte Spill
	movq	88(%rsp), %r14          # 8-byte Reload
	adcq	1032(%rsp), %r14
	movq	104(%rsp), %rcx         # 8-byte Reload
	adcq	1040(%rsp), %rcx
	movq	%rcx, 104(%rsp)         # 8-byte Spill
	adcq	1048(%rsp), %r13
	movq	%r13, 112(%rsp)         # 8-byte Spill
	adcq	1056(%rsp), %rbp
	adcq	1064(%rsp), %r12
	movq	96(%rsp), %rcx          # 8-byte Reload
	adcq	1072(%rsp), %rcx
	movq	%rcx, 96(%rsp)          # 8-byte Spill
	sbbq	%rbx, %rbx
	movq	%rax, %rdx
	movq	%rax, %r13
	imulq	16(%rsp), %rdx          # 8-byte Folded Reload
	leaq	920(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	andl	$1, %ebx
	movq	%rbx, %rax
	addq	920(%rsp), %r13
	movq	40(%rsp), %rcx          # 8-byte Reload
	adcq	928(%rsp), %rcx
	movq	%rcx, 40(%rsp)          # 8-byte Spill
	adcq	936(%rsp), %r15
	movq	%r15, 64(%rsp)          # 8-byte Spill
	movq	48(%rsp), %r15          # 8-byte Reload
	adcq	944(%rsp), %r15
	movq	%r14, %r13
	adcq	952(%rsp), %r13
	movq	104(%rsp), %r14         # 8-byte Reload
	adcq	960(%rsp), %r14
	movq	112(%rsp), %rbx         # 8-byte Reload
	adcq	968(%rsp), %rbx
	adcq	976(%rsp), %rbp
	adcq	984(%rsp), %r12
	movq	96(%rsp), %rcx          # 8-byte Reload
	adcq	992(%rsp), %rcx
	movq	%rcx, 96(%rsp)          # 8-byte Spill
	adcq	$0, %rax
	movq	%rax, 80(%rsp)          # 8-byte Spill
	movq	32(%rsp), %rax          # 8-byte Reload
	movq	32(%rax), %rdx
	leaq	840(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	40(%rsp), %rax          # 8-byte Reload
	addq	840(%rsp), %rax
	movq	64(%rsp), %rcx          # 8-byte Reload
	adcq	848(%rsp), %rcx
	movq	%rcx, 64(%rsp)          # 8-byte Spill
	adcq	856(%rsp), %r15
	adcq	864(%rsp), %r13
	movq	%r13, 88(%rsp)          # 8-byte Spill
	adcq	872(%rsp), %r14
	movq	%r14, 104(%rsp)         # 8-byte Spill
	adcq	880(%rsp), %rbx
	movq	%rbx, 112(%rsp)         # 8-byte Spill
	adcq	888(%rsp), %rbp
	adcq	896(%rsp), %r12
	movq	96(%rsp), %r13          # 8-byte Reload
	adcq	904(%rsp), %r13
	movq	80(%rsp), %rcx          # 8-byte Reload
	adcq	912(%rsp), %rcx
	movq	%rcx, 80(%rsp)          # 8-byte Spill
	sbbq	%rbx, %rbx
	movq	%rax, %rdx
	movq	%rax, %r14
	imulq	16(%rsp), %rdx          # 8-byte Folded Reload
	leaq	760(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	andl	$1, %ebx
	movq	%rbx, %rax
	addq	760(%rsp), %r14
	movq	64(%rsp), %rcx          # 8-byte Reload
	adcq	768(%rsp), %rcx
	movq	%rcx, 64(%rsp)          # 8-byte Spill
	adcq	776(%rsp), %r15
	movq	88(%rsp), %r14          # 8-byte Reload
	adcq	784(%rsp), %r14
	movq	104(%rsp), %rcx         # 8-byte Reload
	adcq	792(%rsp), %rcx
	movq	%rcx, 104(%rsp)         # 8-byte Spill
	movq	112(%rsp), %rcx         # 8-byte Reload
	adcq	800(%rsp), %rcx
	movq	%rcx, 112(%rsp)         # 8-byte Spill
	adcq	808(%rsp), %rbp
	movq	%r12, %rbx
	adcq	816(%rsp), %rbx
	movq	%r13, %r12
	adcq	824(%rsp), %r12
	movq	80(%rsp), %r13          # 8-byte Reload
	adcq	832(%rsp), %r13
	adcq	$0, %rax
	movq	%rax, 40(%rsp)          # 8-byte Spill
	movq	32(%rsp), %rax          # 8-byte Reload
	movq	40(%rax), %rdx
	leaq	680(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	64(%rsp), %rax          # 8-byte Reload
	addq	680(%rsp), %rax
	adcq	688(%rsp), %r15
	movq	%r15, 48(%rsp)          # 8-byte Spill
	adcq	696(%rsp), %r14
	movq	%r14, 88(%rsp)          # 8-byte Spill
	movq	104(%rsp), %rcx         # 8-byte Reload
	adcq	704(%rsp), %rcx
	movq	%rcx, 104(%rsp)         # 8-byte Spill
	movq	112(%rsp), %r15         # 8-byte Reload
	adcq	712(%rsp), %r15
	adcq	720(%rsp), %rbp
	adcq	728(%rsp), %rbx
	movq	%rbx, 72(%rsp)          # 8-byte Spill
	adcq	736(%rsp), %r12
	movq	%r12, 96(%rsp)          # 8-byte Spill
	adcq	744(%rsp), %r13
	movq	%r13, 80(%rsp)          # 8-byte Spill
	movq	40(%rsp), %r13          # 8-byte Reload
	adcq	752(%rsp), %r13
	sbbq	%r14, %r14
	movq	%rax, %rdx
	movq	%rax, %rbx
	imulq	16(%rsp), %rdx          # 8-byte Folded Reload
	leaq	600(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	andl	$1, %r14d
	addq	600(%rsp), %rbx
	movq	48(%rsp), %rax          # 8-byte Reload
	adcq	608(%rsp), %rax
	movq	%rax, 48(%rsp)          # 8-byte Spill
	movq	88(%rsp), %rax          # 8-byte Reload
	adcq	616(%rsp), %rax
	movq	%rax, 88(%rsp)          # 8-byte Spill
	movq	104(%rsp), %rbx         # 8-byte Reload
	adcq	624(%rsp), %rbx
	adcq	632(%rsp), %r15
	movq	%r15, 112(%rsp)         # 8-byte Spill
	adcq	640(%rsp), %rbp
	movq	72(%rsp), %r12          # 8-byte Reload
	adcq	648(%rsp), %r12
	movq	96(%rsp), %rax          # 8-byte Reload
	adcq	656(%rsp), %rax
	movq	%rax, 96(%rsp)          # 8-byte Spill
	movq	80(%rsp), %r15          # 8-byte Reload
	adcq	664(%rsp), %r15
	adcq	672(%rsp), %r13
	adcq	$0, %r14
	movq	%r14, 64(%rsp)          # 8-byte Spill
	movq	32(%rsp), %rax          # 8-byte Reload
	movq	48(%rax), %rdx
	leaq	520(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	48(%rsp), %rax          # 8-byte Reload
	addq	520(%rsp), %rax
	movq	88(%rsp), %r14          # 8-byte Reload
	adcq	528(%rsp), %r14
	adcq	536(%rsp), %rbx
	movq	%rbx, 104(%rsp)         # 8-byte Spill
	movq	112(%rsp), %rcx         # 8-byte Reload
	adcq	544(%rsp), %rcx
	movq	%rcx, 112(%rsp)         # 8-byte Spill
	adcq	552(%rsp), %rbp
	adcq	560(%rsp), %r12
	movq	%r12, 72(%rsp)          # 8-byte Spill
	movq	96(%rsp), %r12          # 8-byte Reload
	adcq	568(%rsp), %r12
	adcq	576(%rsp), %r15
	movq	%r15, 80(%rsp)          # 8-byte Spill
	adcq	584(%rsp), %r13
	movq	%r13, 40(%rsp)          # 8-byte Spill
	movq	64(%rsp), %r15          # 8-byte Reload
	adcq	592(%rsp), %r15
	sbbq	%rbx, %rbx
	movq	%rax, %rdx
	movq	%rax, %r13
	imulq	16(%rsp), %rdx          # 8-byte Folded Reload
	leaq	440(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	andl	$1, %ebx
	movq	%rbx, %rax
	addq	440(%rsp), %r13
	adcq	448(%rsp), %r14
	movq	%r14, 88(%rsp)          # 8-byte Spill
	movq	104(%rsp), %r14         # 8-byte Reload
	adcq	456(%rsp), %r14
	movq	112(%rsp), %rbx         # 8-byte Reload
	adcq	464(%rsp), %rbx
	adcq	472(%rsp), %rbp
	movq	%rbp, 8(%rsp)           # 8-byte Spill
	movq	72(%rsp), %rcx          # 8-byte Reload
	adcq	480(%rsp), %rcx
	movq	%rcx, 72(%rsp)          # 8-byte Spill
	adcq	488(%rsp), %r12
	movq	%r12, 96(%rsp)          # 8-byte Spill
	movq	80(%rsp), %rbp          # 8-byte Reload
	adcq	496(%rsp), %rbp
	movq	40(%rsp), %r12          # 8-byte Reload
	adcq	504(%rsp), %r12
	adcq	512(%rsp), %r15
	movq	%r15, %r13
	adcq	$0, %rax
	movq	%rax, 48(%rsp)          # 8-byte Spill
	movq	32(%rsp), %rax          # 8-byte Reload
	movq	56(%rax), %rdx
	leaq	360(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	88(%rsp), %rax          # 8-byte Reload
	addq	360(%rsp), %rax
	adcq	368(%rsp), %r14
	adcq	376(%rsp), %rbx
	movq	%rbx, 112(%rsp)         # 8-byte Spill
	movq	8(%rsp), %rcx           # 8-byte Reload
	adcq	384(%rsp), %rcx
	movq	%rcx, 8(%rsp)           # 8-byte Spill
	movq	72(%rsp), %rbx          # 8-byte Reload
	adcq	392(%rsp), %rbx
	movq	96(%rsp), %r15          # 8-byte Reload
	adcq	400(%rsp), %r15
	adcq	408(%rsp), %rbp
	movq	%rbp, 80(%rsp)          # 8-byte Spill
	adcq	416(%rsp), %r12
	movq	%r12, %rbp
	adcq	424(%rsp), %r13
	movq	%r13, 64(%rsp)          # 8-byte Spill
	movq	48(%rsp), %rcx          # 8-byte Reload
	adcq	432(%rsp), %rcx
	movq	%rcx, 48(%rsp)          # 8-byte Spill
	sbbq	%r13, %r13
	movq	%rax, %rdx
	movq	%rax, %r12
	imulq	16(%rsp), %rdx          # 8-byte Folded Reload
	leaq	280(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	andl	$1, %r13d
	addq	280(%rsp), %r12
	adcq	288(%rsp), %r14
	movq	%r14, 104(%rsp)         # 8-byte Spill
	movq	112(%rsp), %rax         # 8-byte Reload
	adcq	296(%rsp), %rax
	movq	%rax, 112(%rsp)         # 8-byte Spill
	movq	8(%rsp), %r14           # 8-byte Reload
	adcq	304(%rsp), %r14
	adcq	312(%rsp), %rbx
	movq	%rbx, 72(%rsp)          # 8-byte Spill
	adcq	320(%rsp), %r15
	movq	%r15, 96(%rsp)          # 8-byte Spill
	movq	80(%rsp), %rbx          # 8-byte Reload
	adcq	328(%rsp), %rbx
	adcq	336(%rsp), %rbp
	movq	%rbp, 40(%rsp)          # 8-byte Spill
	movq	64(%rsp), %r12          # 8-byte Reload
	adcq	344(%rsp), %r12
	movq	48(%rsp), %rbp          # 8-byte Reload
	adcq	352(%rsp), %rbp
	adcq	$0, %r13
	movq	32(%rsp), %rax          # 8-byte Reload
	movq	64(%rax), %rdx
	leaq	200(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	104(%rsp), %rax         # 8-byte Reload
	addq	200(%rsp), %rax
	movq	112(%rsp), %r15         # 8-byte Reload
	adcq	208(%rsp), %r15
	adcq	216(%rsp), %r14
	movq	%r14, 8(%rsp)           # 8-byte Spill
	movq	72(%rsp), %r14          # 8-byte Reload
	adcq	224(%rsp), %r14
	movq	96(%rsp), %rcx          # 8-byte Reload
	adcq	232(%rsp), %rcx
	movq	%rcx, 96(%rsp)          # 8-byte Spill
	adcq	240(%rsp), %rbx
	movq	%rbx, 80(%rsp)          # 8-byte Spill
	movq	40(%rsp), %rcx          # 8-byte Reload
	adcq	248(%rsp), %rcx
	movq	%rcx, 40(%rsp)          # 8-byte Spill
	adcq	256(%rsp), %r12
	movq	%r12, 64(%rsp)          # 8-byte Spill
	adcq	264(%rsp), %rbp
	movq	%rbp, 48(%rsp)          # 8-byte Spill
	adcq	272(%rsp), %r13
	sbbq	%rbx, %rbx
	movq	16(%rsp), %rdx          # 8-byte Reload
	imulq	%rax, %rdx
	movq	%rax, %r12
	leaq	120(%rsp), %rdi
	movq	56(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	andl	$1, %ebx
	addq	120(%rsp), %r12
	adcq	128(%rsp), %r15
	movq	8(%rsp), %rbp           # 8-byte Reload
	adcq	136(%rsp), %rbp
	movq	%r14, %rcx
	adcq	144(%rsp), %rcx
	movq	%rcx, 72(%rsp)          # 8-byte Spill
	movq	96(%rsp), %r8           # 8-byte Reload
	adcq	152(%rsp), %r8
	movq	%r8, 96(%rsp)           # 8-byte Spill
	movq	80(%rsp), %r9           # 8-byte Reload
	adcq	160(%rsp), %r9
	movq	%r9, 80(%rsp)           # 8-byte Spill
	movq	40(%rsp), %r10          # 8-byte Reload
	adcq	168(%rsp), %r10
	movq	%r10, 40(%rsp)          # 8-byte Spill
	movq	64(%rsp), %rdi          # 8-byte Reload
	adcq	176(%rsp), %rdi
	movq	%rdi, 64(%rsp)          # 8-byte Spill
	movq	48(%rsp), %r14          # 8-byte Reload
	adcq	184(%rsp), %r14
	adcq	192(%rsp), %r13
	adcq	$0, %rbx
	movq	%r15, %rsi
	movq	%r15, %r12
	movq	56(%rsp), %rdx          # 8-byte Reload
	subq	(%rdx), %rsi
	movq	%rbp, %rax
	movq	%rbp, %r15
	sbbq	8(%rdx), %rax
	movq	%rcx, %rbp
	sbbq	16(%rdx), %rbp
	movq	%r8, %rcx
	sbbq	24(%rdx), %rcx
	movq	%r9, %r8
	sbbq	32(%rdx), %r8
	movq	%r10, %r11
	sbbq	40(%rdx), %r11
	movq	%rdi, %r10
	sbbq	48(%rdx), %r10
	movq	%r14, %rdi
	sbbq	56(%rdx), %rdi
	movq	%r13, %r9
	sbbq	64(%rdx), %r9
	sbbq	$0, %rbx
	andl	$1, %ebx
	cmovneq	%r13, %r9
	testb	%bl, %bl
	cmovneq	%r12, %rsi
	movq	(%rsp), %rbx            # 8-byte Reload
	movq	%rsi, (%rbx)
	cmovneq	%r15, %rax
	movq	%rax, 8(%rbx)
	cmovneq	72(%rsp), %rbp          # 8-byte Folded Reload
	movq	%rbp, 16(%rbx)
	cmovneq	96(%rsp), %rcx          # 8-byte Folded Reload
	movq	%rcx, 24(%rbx)
	cmovneq	80(%rsp), %r8           # 8-byte Folded Reload
	movq	%r8, 32(%rbx)
	cmovneq	40(%rsp), %r11          # 8-byte Folded Reload
	movq	%r11, 40(%rbx)
	cmovneq	64(%rsp), %r10          # 8-byte Folded Reload
	movq	%r10, 48(%rbx)
	cmovneq	%r14, %rdi
	movq	%rdi, 56(%rbx)
	movq	%r9, 64(%rbx)
	addq	$1560, %rsp             # imm = 0x618
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end130:
	.size	mcl_fp_mont9Lbmi2, .Lfunc_end130-mcl_fp_mont9Lbmi2

	.globl	mcl_fp_montNF9Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montNF9Lbmi2,@function
mcl_fp_montNF9Lbmi2:                    # @mcl_fp_montNF9Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	subq	$1560, %rsp             # imm = 0x618
	movq	%rcx, 64(%rsp)          # 8-byte Spill
	movq	%rdx, 16(%rsp)          # 8-byte Spill
	movq	%rsi, 24(%rsp)          # 8-byte Spill
	movq	%rdi, (%rsp)            # 8-byte Spill
	movq	-8(%rcx), %rbx
	movq	%rbx, 32(%rsp)          # 8-byte Spill
	movq	(%rdx), %rdx
	leaq	1480(%rsp), %rdi
	callq	.LmulPv576x64
	movq	1480(%rsp), %r12
	movq	1488(%rsp), %rax
	movq	%rax, 88(%rsp)          # 8-byte Spill
	movq	%r12, %rdx
	imulq	%rbx, %rdx
	movq	1552(%rsp), %rax
	movq	%rax, 112(%rsp)         # 8-byte Spill
	movq	1544(%rsp), %r13
	movq	1536(%rsp), %rax
	movq	%rax, 72(%rsp)          # 8-byte Spill
	movq	1528(%rsp), %rax
	movq	%rax, 48(%rsp)          # 8-byte Spill
	movq	1520(%rsp), %r14
	movq	1512(%rsp), %r15
	movq	1504(%rsp), %rbx
	movq	1496(%rsp), %rbp
	leaq	1400(%rsp), %rdi
	movq	64(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	addq	1400(%rsp), %r12
	movq	88(%rsp), %rax          # 8-byte Reload
	adcq	1408(%rsp), %rax
	movq	%rax, 88(%rsp)          # 8-byte Spill
	adcq	1416(%rsp), %rbp
	movq	%rbp, 8(%rsp)           # 8-byte Spill
	adcq	1424(%rsp), %rbx
	movq	%rbx, 104(%rsp)         # 8-byte Spill
	adcq	1432(%rsp), %r15
	movq	%r15, 56(%rsp)          # 8-byte Spill
	adcq	1440(%rsp), %r14
	movq	%r14, 40(%rsp)          # 8-byte Spill
	movq	48(%rsp), %rbx          # 8-byte Reload
	adcq	1448(%rsp), %rbx
	movq	72(%rsp), %r12          # 8-byte Reload
	adcq	1456(%rsp), %r12
	adcq	1464(%rsp), %r13
	movq	%r13, 96(%rsp)          # 8-byte Spill
	movq	112(%rsp), %rbp         # 8-byte Reload
	adcq	1472(%rsp), %rbp
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	8(%rax), %rdx
	leaq	1320(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	1392(%rsp), %rax
	movq	88(%rsp), %rcx          # 8-byte Reload
	addq	1320(%rsp), %rcx
	movq	8(%rsp), %r15           # 8-byte Reload
	adcq	1328(%rsp), %r15
	movq	104(%rsp), %r14         # 8-byte Reload
	adcq	1336(%rsp), %r14
	movq	56(%rsp), %rdx          # 8-byte Reload
	adcq	1344(%rsp), %rdx
	movq	%rdx, 56(%rsp)          # 8-byte Spill
	movq	40(%rsp), %r13          # 8-byte Reload
	adcq	1352(%rsp), %r13
	adcq	1360(%rsp), %rbx
	movq	%rbx, 48(%rsp)          # 8-byte Spill
	adcq	1368(%rsp), %r12
	movq	%r12, 72(%rsp)          # 8-byte Spill
	movq	96(%rsp), %rdx          # 8-byte Reload
	adcq	1376(%rsp), %rdx
	movq	%rdx, 96(%rsp)          # 8-byte Spill
	adcq	1384(%rsp), %rbp
	movq	%rbp, 112(%rsp)         # 8-byte Spill
	adcq	$0, %rax
	movq	%rax, %rbp
	movq	%rcx, %rdx
	movq	%rcx, %rbx
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	1240(%rsp), %rdi
	movq	64(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	addq	1240(%rsp), %rbx
	adcq	1248(%rsp), %r15
	movq	%r15, 8(%rsp)           # 8-byte Spill
	adcq	1256(%rsp), %r14
	movq	%r14, 104(%rsp)         # 8-byte Spill
	movq	56(%rsp), %r12          # 8-byte Reload
	adcq	1264(%rsp), %r12
	adcq	1272(%rsp), %r13
	movq	%r13, %r14
	movq	48(%rsp), %r13          # 8-byte Reload
	adcq	1280(%rsp), %r13
	movq	72(%rsp), %rbx          # 8-byte Reload
	adcq	1288(%rsp), %rbx
	movq	96(%rsp), %r15          # 8-byte Reload
	adcq	1296(%rsp), %r15
	movq	112(%rsp), %rax         # 8-byte Reload
	adcq	1304(%rsp), %rax
	movq	%rax, 112(%rsp)         # 8-byte Spill
	adcq	1312(%rsp), %rbp
	movq	%rbp, 80(%rsp)          # 8-byte Spill
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	16(%rax), %rdx
	leaq	1160(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	1232(%rsp), %rax
	movq	8(%rsp), %rcx           # 8-byte Reload
	addq	1160(%rsp), %rcx
	movq	104(%rsp), %rbp         # 8-byte Reload
	adcq	1168(%rsp), %rbp
	adcq	1176(%rsp), %r12
	movq	%r12, 56(%rsp)          # 8-byte Spill
	adcq	1184(%rsp), %r14
	adcq	1192(%rsp), %r13
	movq	%r13, %r12
	adcq	1200(%rsp), %rbx
	movq	%rbx, 72(%rsp)          # 8-byte Spill
	adcq	1208(%rsp), %r15
	movq	%r15, 96(%rsp)          # 8-byte Spill
	movq	112(%rsp), %rbx         # 8-byte Reload
	adcq	1216(%rsp), %rbx
	movq	80(%rsp), %rdx          # 8-byte Reload
	adcq	1224(%rsp), %rdx
	movq	%rdx, 80(%rsp)          # 8-byte Spill
	movq	%rax, %r15
	adcq	$0, %r15
	movq	%rcx, %rdx
	movq	%rcx, %r13
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	1080(%rsp), %rdi
	movq	64(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	addq	1080(%rsp), %r13
	adcq	1088(%rsp), %rbp
	movq	%rbp, 104(%rsp)         # 8-byte Spill
	movq	56(%rsp), %r13          # 8-byte Reload
	adcq	1096(%rsp), %r13
	adcq	1104(%rsp), %r14
	adcq	1112(%rsp), %r12
	movq	%r12, 48(%rsp)          # 8-byte Spill
	movq	72(%rsp), %r12          # 8-byte Reload
	adcq	1120(%rsp), %r12
	movq	96(%rsp), %rbp          # 8-byte Reload
	adcq	1128(%rsp), %rbp
	adcq	1136(%rsp), %rbx
	movq	%rbx, 112(%rsp)         # 8-byte Spill
	movq	80(%rsp), %rbx          # 8-byte Reload
	adcq	1144(%rsp), %rbx
	adcq	1152(%rsp), %r15
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	24(%rax), %rdx
	leaq	1000(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	1072(%rsp), %rax
	movq	104(%rsp), %rcx         # 8-byte Reload
	addq	1000(%rsp), %rcx
	adcq	1008(%rsp), %r13
	movq	%r13, 56(%rsp)          # 8-byte Spill
	adcq	1016(%rsp), %r14
	movq	%r14, 40(%rsp)          # 8-byte Spill
	movq	48(%rsp), %r14          # 8-byte Reload
	adcq	1024(%rsp), %r14
	adcq	1032(%rsp), %r12
	adcq	1040(%rsp), %rbp
	movq	%rbp, 96(%rsp)          # 8-byte Spill
	movq	112(%rsp), %r13         # 8-byte Reload
	adcq	1048(%rsp), %r13
	adcq	1056(%rsp), %rbx
	movq	%rbx, 80(%rsp)          # 8-byte Spill
	adcq	1064(%rsp), %r15
	movq	%r15, 88(%rsp)          # 8-byte Spill
	adcq	$0, %rax
	movq	%rax, 104(%rsp)         # 8-byte Spill
	movq	%rcx, %rdx
	movq	%rcx, %rbx
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	920(%rsp), %rdi
	movq	64(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	addq	920(%rsp), %rbx
	movq	56(%rsp), %rax          # 8-byte Reload
	adcq	928(%rsp), %rax
	movq	%rax, 56(%rsp)          # 8-byte Spill
	movq	40(%rsp), %rbp          # 8-byte Reload
	adcq	936(%rsp), %rbp
	movq	%r14, %rbx
	adcq	944(%rsp), %rbx
	adcq	952(%rsp), %r12
	movq	96(%rsp), %rax          # 8-byte Reload
	adcq	960(%rsp), %rax
	movq	%rax, 96(%rsp)          # 8-byte Spill
	adcq	968(%rsp), %r13
	movq	%r13, %r15
	movq	80(%rsp), %r13          # 8-byte Reload
	adcq	976(%rsp), %r13
	movq	88(%rsp), %r14          # 8-byte Reload
	adcq	984(%rsp), %r14
	movq	104(%rsp), %rax         # 8-byte Reload
	adcq	992(%rsp), %rax
	movq	%rax, 104(%rsp)         # 8-byte Spill
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	32(%rax), %rdx
	leaq	840(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	912(%rsp), %rax
	movq	56(%rsp), %rcx          # 8-byte Reload
	addq	840(%rsp), %rcx
	adcq	848(%rsp), %rbp
	movq	%rbp, 40(%rsp)          # 8-byte Spill
	adcq	856(%rsp), %rbx
	movq	%rbx, 48(%rsp)          # 8-byte Spill
	adcq	864(%rsp), %r12
	movq	96(%rsp), %rbp          # 8-byte Reload
	adcq	872(%rsp), %rbp
	adcq	880(%rsp), %r15
	movq	%r15, 112(%rsp)         # 8-byte Spill
	adcq	888(%rsp), %r13
	adcq	896(%rsp), %r14
	movq	%r14, 88(%rsp)          # 8-byte Spill
	movq	104(%rsp), %rdx         # 8-byte Reload
	adcq	904(%rsp), %rdx
	movq	%rdx, 104(%rsp)         # 8-byte Spill
	adcq	$0, %rax
	movq	%rax, %r14
	movq	%rcx, %rdx
	movq	%rcx, %rbx
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	760(%rsp), %rdi
	movq	64(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	addq	760(%rsp), %rbx
	movq	40(%rsp), %rax          # 8-byte Reload
	adcq	768(%rsp), %rax
	movq	%rax, 40(%rsp)          # 8-byte Spill
	movq	48(%rsp), %r15          # 8-byte Reload
	adcq	776(%rsp), %r15
	adcq	784(%rsp), %r12
	movq	%r12, 72(%rsp)          # 8-byte Spill
	movq	%rbp, %rbx
	adcq	792(%rsp), %rbx
	movq	112(%rsp), %rbp         # 8-byte Reload
	adcq	800(%rsp), %rbp
	adcq	808(%rsp), %r13
	movq	88(%rsp), %rax          # 8-byte Reload
	adcq	816(%rsp), %rax
	movq	%rax, 88(%rsp)          # 8-byte Spill
	movq	104(%rsp), %r12         # 8-byte Reload
	adcq	824(%rsp), %r12
	adcq	832(%rsp), %r14
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	40(%rax), %rdx
	leaq	680(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	752(%rsp), %rcx
	movq	40(%rsp), %rax          # 8-byte Reload
	addq	680(%rsp), %rax
	adcq	688(%rsp), %r15
	movq	%r15, 48(%rsp)          # 8-byte Spill
	movq	72(%rsp), %rdx          # 8-byte Reload
	adcq	696(%rsp), %rdx
	movq	%rdx, 72(%rsp)          # 8-byte Spill
	adcq	704(%rsp), %rbx
	movq	%rbx, 96(%rsp)          # 8-byte Spill
	adcq	712(%rsp), %rbp
	movq	%rbp, 112(%rsp)         # 8-byte Spill
	adcq	720(%rsp), %r13
	movq	%r13, %r15
	movq	88(%rsp), %rbx          # 8-byte Reload
	adcq	728(%rsp), %rbx
	adcq	736(%rsp), %r12
	movq	%r12, 104(%rsp)         # 8-byte Spill
	adcq	744(%rsp), %r14
	movq	%r14, 40(%rsp)          # 8-byte Spill
	adcq	$0, %rcx
	movq	%rcx, 56(%rsp)          # 8-byte Spill
	movq	%rax, %rdx
	movq	%rax, %r13
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	600(%rsp), %rdi
	movq	64(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	addq	600(%rsp), %r13
	movq	48(%rsp), %r13          # 8-byte Reload
	adcq	608(%rsp), %r13
	movq	72(%rsp), %r12          # 8-byte Reload
	adcq	616(%rsp), %r12
	movq	96(%rsp), %rbp          # 8-byte Reload
	adcq	624(%rsp), %rbp
	movq	112(%rsp), %rax         # 8-byte Reload
	adcq	632(%rsp), %rax
	movq	%rax, 112(%rsp)         # 8-byte Spill
	adcq	640(%rsp), %r15
	movq	%r15, 80(%rsp)          # 8-byte Spill
	adcq	648(%rsp), %rbx
	movq	%rbx, 88(%rsp)          # 8-byte Spill
	movq	104(%rsp), %r14         # 8-byte Reload
	adcq	656(%rsp), %r14
	movq	40(%rsp), %rbx          # 8-byte Reload
	adcq	664(%rsp), %rbx
	movq	56(%rsp), %r15          # 8-byte Reload
	adcq	672(%rsp), %r15
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	48(%rax), %rdx
	leaq	520(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	592(%rsp), %rcx
	movq	%r13, %rax
	addq	520(%rsp), %rax
	adcq	528(%rsp), %r12
	movq	%r12, 72(%rsp)          # 8-byte Spill
	movq	%rbp, %r12
	adcq	536(%rsp), %r12
	movq	112(%rsp), %rbp         # 8-byte Reload
	adcq	544(%rsp), %rbp
	movq	80(%rsp), %rdx          # 8-byte Reload
	adcq	552(%rsp), %rdx
	movq	%rdx, 80(%rsp)          # 8-byte Spill
	movq	88(%rsp), %rdx          # 8-byte Reload
	adcq	560(%rsp), %rdx
	movq	%rdx, 88(%rsp)          # 8-byte Spill
	adcq	568(%rsp), %r14
	movq	%r14, 104(%rsp)         # 8-byte Spill
	adcq	576(%rsp), %rbx
	movq	%rbx, 40(%rsp)          # 8-byte Spill
	adcq	584(%rsp), %r15
	movq	%r15, 56(%rsp)          # 8-byte Spill
	adcq	$0, %rcx
	movq	%rcx, %r13
	movq	%rax, %rdx
	movq	%rax, %r14
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	440(%rsp), %rdi
	movq	64(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	addq	440(%rsp), %r14
	movq	72(%rsp), %rax          # 8-byte Reload
	adcq	448(%rsp), %rax
	movq	%rax, 72(%rsp)          # 8-byte Spill
	adcq	456(%rsp), %r12
	adcq	464(%rsp), %rbp
	movq	%rbp, 112(%rsp)         # 8-byte Spill
	movq	80(%rsp), %r14          # 8-byte Reload
	adcq	472(%rsp), %r14
	movq	88(%rsp), %r15          # 8-byte Reload
	adcq	480(%rsp), %r15
	movq	104(%rsp), %rbp         # 8-byte Reload
	adcq	488(%rsp), %rbp
	movq	40(%rsp), %rbx          # 8-byte Reload
	adcq	496(%rsp), %rbx
	movq	56(%rsp), %rax          # 8-byte Reload
	adcq	504(%rsp), %rax
	movq	%rax, 56(%rsp)          # 8-byte Spill
	adcq	512(%rsp), %r13
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	56(%rax), %rdx
	leaq	360(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	432(%rsp), %rcx
	movq	72(%rsp), %rax          # 8-byte Reload
	addq	360(%rsp), %rax
	adcq	368(%rsp), %r12
	movq	%r12, 96(%rsp)          # 8-byte Spill
	movq	112(%rsp), %rdx         # 8-byte Reload
	adcq	376(%rsp), %rdx
	movq	%rdx, 112(%rsp)         # 8-byte Spill
	adcq	384(%rsp), %r14
	movq	%r14, 80(%rsp)          # 8-byte Spill
	adcq	392(%rsp), %r15
	movq	%r15, 88(%rsp)          # 8-byte Spill
	adcq	400(%rsp), %rbp
	movq	%rbp, 104(%rsp)         # 8-byte Spill
	adcq	408(%rsp), %rbx
	movq	%rbx, 40(%rsp)          # 8-byte Spill
	movq	56(%rsp), %r14          # 8-byte Reload
	adcq	416(%rsp), %r14
	adcq	424(%rsp), %r13
	movq	%r13, %r15
	adcq	$0, %rcx
	movq	%rcx, 72(%rsp)          # 8-byte Spill
	movq	%rax, %rdx
	movq	%rax, %r12
	imulq	32(%rsp), %rdx          # 8-byte Folded Reload
	leaq	280(%rsp), %rdi
	movq	64(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	addq	280(%rsp), %r12
	movq	96(%rsp), %rax          # 8-byte Reload
	adcq	288(%rsp), %rax
	movq	%rax, 96(%rsp)          # 8-byte Spill
	movq	112(%rsp), %rbp         # 8-byte Reload
	adcq	296(%rsp), %rbp
	movq	80(%rsp), %rax          # 8-byte Reload
	adcq	304(%rsp), %rax
	movq	%rax, 80(%rsp)          # 8-byte Spill
	movq	88(%rsp), %r13          # 8-byte Reload
	adcq	312(%rsp), %r13
	movq	104(%rsp), %r12         # 8-byte Reload
	adcq	320(%rsp), %r12
	movq	40(%rsp), %rbx          # 8-byte Reload
	adcq	328(%rsp), %rbx
	adcq	336(%rsp), %r14
	movq	%r14, 56(%rsp)          # 8-byte Spill
	adcq	344(%rsp), %r15
	movq	%r15, 48(%rsp)          # 8-byte Spill
	movq	72(%rsp), %r14          # 8-byte Reload
	adcq	352(%rsp), %r14
	movq	16(%rsp), %rax          # 8-byte Reload
	movq	64(%rax), %rdx
	leaq	200(%rsp), %rdi
	movq	24(%rsp), %rsi          # 8-byte Reload
	callq	.LmulPv576x64
	movq	272(%rsp), %rcx
	movq	96(%rsp), %rax          # 8-byte Reload
	addq	200(%rsp), %rax
	adcq	208(%rsp), %rbp
	movq	%rbp, 112(%rsp)         # 8-byte Spill
	movq	80(%rsp), %rbp          # 8-byte Reload
	adcq	216(%rsp), %rbp
	adcq	224(%rsp), %r13
	movq	%r13, 88(%rsp)          # 8-byte Spill
	adcq	232(%rsp), %r12
	movq	%r12, 104(%rsp)         # 8-byte Spill
	adcq	240(%rsp), %rbx
	movq	%rbx, 40(%rsp)          # 8-byte Spill
	movq	56(%rsp), %r15          # 8-byte Reload
	adcq	248(%rsp), %r15
	movq	48(%rsp), %r12          # 8-byte Reload
	adcq	256(%rsp), %r12
	adcq	264(%rsp), %r14
	adcq	$0, %rcx
	movq	%rcx, 96(%rsp)          # 8-byte Spill
	movq	32(%rsp), %rdx          # 8-byte Reload
	imulq	%rax, %rdx
	movq	%rax, %rbx
	leaq	120(%rsp), %rdi
	movq	64(%rsp), %r13          # 8-byte Reload
	movq	%r13, %rsi
	callq	.LmulPv576x64
	addq	120(%rsp), %rbx
	movq	112(%rsp), %rcx         # 8-byte Reload
	adcq	128(%rsp), %rcx
	movq	%rbp, %rdx
	adcq	136(%rsp), %rdx
	movq	88(%rsp), %rsi          # 8-byte Reload
	adcq	144(%rsp), %rsi
	movq	%rsi, 88(%rsp)          # 8-byte Spill
	movq	104(%rsp), %rdi         # 8-byte Reload
	adcq	152(%rsp), %rdi
	movq	%rdi, 104(%rsp)         # 8-byte Spill
	movq	40(%rsp), %rbx          # 8-byte Reload
	adcq	160(%rsp), %rbx
	movq	%rbx, 40(%rsp)          # 8-byte Spill
	movq	%r15, %r8
	adcq	168(%rsp), %r8
	movq	%r8, 56(%rsp)           # 8-byte Spill
	movq	%r12, %r15
	adcq	176(%rsp), %r15
	adcq	184(%rsp), %r14
	movq	96(%rsp), %r9           # 8-byte Reload
	adcq	192(%rsp), %r9
	movq	%rcx, %rax
	movq	%rcx, %r11
	movq	%r13, %rbp
	subq	(%rbp), %rax
	movq	%rdx, %rcx
	movq	%rdx, %r12
	sbbq	8(%rbp), %rcx
	movq	%rsi, %rdx
	sbbq	16(%rbp), %rdx
	movq	%rdi, %rsi
	sbbq	24(%rbp), %rsi
	movq	%rbx, %rdi
	sbbq	32(%rbp), %rdi
	movq	%r8, %r10
	sbbq	40(%rbp), %r10
	movq	%r15, %r13
	sbbq	48(%rbp), %r13
	movq	%r14, %r8
	sbbq	56(%rbp), %r8
	movq	%rbp, %rbx
	movq	%r9, %rbp
	sbbq	64(%rbx), %rbp
	movq	%rbp, %rbx
	sarq	$63, %rbx
	cmovsq	%r11, %rax
	movq	(%rsp), %rbx            # 8-byte Reload
	movq	%rax, (%rbx)
	cmovsq	%r12, %rcx
	movq	%rcx, 8(%rbx)
	cmovsq	88(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, 16(%rbx)
	cmovsq	104(%rsp), %rsi         # 8-byte Folded Reload
	movq	%rsi, 24(%rbx)
	cmovsq	40(%rsp), %rdi          # 8-byte Folded Reload
	movq	%rdi, 32(%rbx)
	cmovsq	56(%rsp), %r10          # 8-byte Folded Reload
	movq	%r10, 40(%rbx)
	cmovsq	%r15, %r13
	movq	%r13, 48(%rbx)
	cmovsq	%r14, %r8
	movq	%r8, 56(%rbx)
	cmovsq	%r9, %rbp
	movq	%rbp, 64(%rbx)
	addq	$1560, %rsp             # imm = 0x618
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end131:
	.size	mcl_fp_montNF9Lbmi2, .Lfunc_end131-mcl_fp_montNF9Lbmi2

	.globl	mcl_fp_montRed9Lbmi2
	.align	16, 0x90
	.type	mcl_fp_montRed9Lbmi2,@function
mcl_fp_montRed9Lbmi2:                   # @mcl_fp_montRed9Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	subq	$936, %rsp              # imm = 0x3A8
	movq	%rdx, %rax
	movq	%rax, 128(%rsp)         # 8-byte Spill
	movq	%rdi, 80(%rsp)          # 8-byte Spill
	movq	-8(%rax), %rcx
	movq	%rcx, 120(%rsp)         # 8-byte Spill
	movq	(%rsi), %r14
	movq	8(%rsi), %rdx
	movq	%rdx, 192(%rsp)         # 8-byte Spill
	movq	%r14, %rdx
	imulq	%rcx, %rdx
	movq	136(%rsi), %rcx
	movq	%rcx, 112(%rsp)         # 8-byte Spill
	movq	128(%rsi), %rcx
	movq	%rcx, 152(%rsp)         # 8-byte Spill
	movq	120(%rsi), %rcx
	movq	%rcx, 104(%rsp)         # 8-byte Spill
	movq	112(%rsi), %rcx
	movq	%rcx, 144(%rsp)         # 8-byte Spill
	movq	104(%rsi), %rcx
	movq	%rcx, 184(%rsp)         # 8-byte Spill
	movq	96(%rsi), %rcx
	movq	%rcx, 208(%rsp)         # 8-byte Spill
	movq	88(%rsi), %rcx
	movq	%rcx, 200(%rsp)         # 8-byte Spill
	movq	80(%rsi), %rcx
	movq	%rcx, 160(%rsp)         # 8-byte Spill
	movq	72(%rsi), %r12
	movq	64(%rsi), %rcx
	movq	%rcx, 176(%rsp)         # 8-byte Spill
	movq	56(%rsi), %rcx
	movq	%rcx, 168(%rsp)         # 8-byte Spill
	movq	48(%rsi), %rcx
	movq	%rcx, 136(%rsp)         # 8-byte Spill
	movq	40(%rsi), %rbp
	movq	32(%rsi), %rbx
	movq	24(%rsi), %r13
	movq	16(%rsi), %r15
	movq	%rax, %rcx
	movq	(%rcx), %rax
	movq	%rax, 16(%rsp)          # 8-byte Spill
	movq	64(%rcx), %rax
	movq	%rax, 72(%rsp)          # 8-byte Spill
	movq	56(%rcx), %rax
	movq	%rax, 64(%rsp)          # 8-byte Spill
	movq	48(%rcx), %rax
	movq	%rax, 56(%rsp)          # 8-byte Spill
	movq	40(%rcx), %rax
	movq	%rax, 48(%rsp)          # 8-byte Spill
	movq	32(%rcx), %rax
	movq	%rax, 40(%rsp)          # 8-byte Spill
	movq	24(%rcx), %rax
	movq	%rax, 32(%rsp)          # 8-byte Spill
	movq	16(%rcx), %rax
	movq	%rax, 24(%rsp)          # 8-byte Spill
	movq	8(%rcx), %rax
	movq	%rax, 8(%rsp)           # 8-byte Spill
	movq	%rcx, %rsi
	leaq	856(%rsp), %rdi
	callq	.LmulPv576x64
	addq	856(%rsp), %r14
	movq	192(%rsp), %rcx         # 8-byte Reload
	adcq	864(%rsp), %rcx
	adcq	872(%rsp), %r15
	adcq	880(%rsp), %r13
	adcq	888(%rsp), %rbx
	movq	%rbx, 88(%rsp)          # 8-byte Spill
	adcq	896(%rsp), %rbp
	movq	%rbp, 96(%rsp)          # 8-byte Spill
	movq	136(%rsp), %rax         # 8-byte Reload
	adcq	904(%rsp), %rax
	movq	%rax, 136(%rsp)         # 8-byte Spill
	movq	168(%rsp), %rax         # 8-byte Reload
	adcq	912(%rsp), %rax
	movq	%rax, 168(%rsp)         # 8-byte Spill
	movq	176(%rsp), %rax         # 8-byte Reload
	adcq	920(%rsp), %rax
	movq	%rax, 176(%rsp)         # 8-byte Spill
	adcq	928(%rsp), %r12
	movq	%r12, 192(%rsp)         # 8-byte Spill
	movq	160(%rsp), %rbp         # 8-byte Reload
	adcq	$0, %rbp
	adcq	$0, 200(%rsp)           # 8-byte Folded Spill
	adcq	$0, 208(%rsp)           # 8-byte Folded Spill
	adcq	$0, 184(%rsp)           # 8-byte Folded Spill
	adcq	$0, 144(%rsp)           # 8-byte Folded Spill
	adcq	$0, 104(%rsp)           # 8-byte Folded Spill
	adcq	$0, 152(%rsp)           # 8-byte Folded Spill
	movq	112(%rsp), %r14         # 8-byte Reload
	adcq	$0, %r14
	sbbq	%r12, %r12
	movq	%rcx, %rdx
	movq	%rcx, %rbx
	imulq	120(%rsp), %rdx         # 8-byte Folded Reload
	leaq	776(%rsp), %rdi
	movq	128(%rsp), %rsi         # 8-byte Reload
	callq	.LmulPv576x64
	andl	$1, %r12d
	addq	776(%rsp), %rbx
	adcq	784(%rsp), %r15
	adcq	792(%rsp), %r13
	movq	%r13, (%rsp)            # 8-byte Spill
	movq	88(%rsp), %rax          # 8-byte Reload
	adcq	800(%rsp), %rax
	movq	%rax, 88(%rsp)          # 8-byte Spill
	movq	96(%rsp), %rax          # 8-byte Reload
	adcq	808(%rsp), %rax
	movq	%rax, 96(%rsp)          # 8-byte Spill
	movq	136(%rsp), %rax         # 8-byte Reload
	adcq	816(%rsp), %rax
	movq	%rax, 136(%rsp)         # 8-byte Spill
	movq	168(%rsp), %rax         # 8-byte Reload
	adcq	824(%rsp), %rax
	movq	%rax, 168(%rsp)         # 8-byte Spill
	movq	176(%rsp), %rax         # 8-byte Reload
	adcq	832(%rsp), %rax
	movq	%rax, 176(%rsp)         # 8-byte Spill
	movq	192(%rsp), %rax         # 8-byte Reload
	adcq	840(%rsp), %rax
	movq	%rax, 192(%rsp)         # 8-byte Spill
	adcq	848(%rsp), %rbp
	movq	%rbp, 160(%rsp)         # 8-byte Spill
	movq	200(%rsp), %r13         # 8-byte Reload
	adcq	$0, %r13
	adcq	$0, 208(%rsp)           # 8-byte Folded Spill
	adcq	$0, 184(%rsp)           # 8-byte Folded Spill
	adcq	$0, 144(%rsp)           # 8-byte Folded Spill
	adcq	$0, 104(%rsp)           # 8-byte Folded Spill
	movq	152(%rsp), %rbx         # 8-byte Reload
	adcq	$0, %rbx
	adcq	$0, %r14
	movq	%r14, 112(%rsp)         # 8-byte Spill
	adcq	$0, %r12
	movq	%r15, %rdx
	imulq	120(%rsp), %rdx         # 8-byte Folded Reload
	leaq	696(%rsp), %rdi
	movq	128(%rsp), %rsi         # 8-byte Reload
	callq	.LmulPv576x64
	addq	696(%rsp), %r15
	movq	(%rsp), %rcx            # 8-byte Reload
	adcq	704(%rsp), %rcx
	movq	88(%rsp), %rax          # 8-byte Reload
	adcq	712(%rsp), %rax
	movq	%rax, 88(%rsp)          # 8-byte Spill
	movq	96(%rsp), %rax          # 8-byte Reload
	adcq	720(%rsp), %rax
	movq	%rax, 96(%rsp)          # 8-byte Spill
	movq	136(%rsp), %rbp         # 8-byte Reload
	adcq	728(%rsp), %rbp
	movq	168(%rsp), %r14         # 8-byte Reload
	adcq	736(%rsp), %r14
	movq	176(%rsp), %r15         # 8-byte Reload
	adcq	744(%rsp), %r15
	movq	192(%rsp), %rax         # 8-byte Reload
	adcq	752(%rsp), %rax
	movq	%rax, 192(%rsp)         # 8-byte Spill
	movq	160(%rsp), %rax         # 8-byte Reload
	adcq	760(%rsp), %rax
	movq	%rax, 160(%rsp)         # 8-byte Spill
	adcq	768(%rsp), %r13
	movq	%r13, 200(%rsp)         # 8-byte Spill
	adcq	$0, 208(%rsp)           # 8-byte Folded Spill
	movq	184(%rsp), %r13         # 8-byte Reload
	adcq	$0, %r13
	adcq	$0, 144(%rsp)           # 8-byte Folded Spill
	adcq	$0, 104(%rsp)           # 8-byte Folded Spill
	adcq	$0, %rbx
	movq	%rbx, 152(%rsp)         # 8-byte Spill
	adcq	$0, 112(%rsp)           # 8-byte Folded Spill
	adcq	$0, %r12
	movq	%rcx, %rbx
	movq	%rbx, %rdx
	imulq	120(%rsp), %rdx         # 8-byte Folded Reload
	leaq	616(%rsp), %rdi
	movq	128(%rsp), %rsi         # 8-byte Reload
	callq	.LmulPv576x64
	addq	616(%rsp), %rbx
	movq	88(%rsp), %rax          # 8-byte Reload
	adcq	624(%rsp), %rax
	movq	96(%rsp), %rcx          # 8-byte Reload
	adcq	632(%rsp), %rcx
	movq	%rcx, 96(%rsp)          # 8-byte Spill
	adcq	640(%rsp), %rbp
	movq	%rbp, 136(%rsp)         # 8-byte Spill
	adcq	648(%rsp), %r14
	movq	%r14, 168(%rsp)         # 8-byte Spill
	adcq	656(%rsp), %r15
	movq	192(%rsp), %r14         # 8-byte Reload
	adcq	664(%rsp), %r14
	movq	160(%rsp), %rbp         # 8-byte Reload
	adcq	672(%rsp), %rbp
	movq	200(%rsp), %rcx         # 8-byte Reload
	adcq	680(%rsp), %rcx
	movq	%rcx, 200(%rsp)         # 8-byte Spill
	movq	208(%rsp), %rcx         # 8-byte Reload
	adcq	688(%rsp), %rcx
	movq	%rcx, 208(%rsp)         # 8-byte Spill
	adcq	$0, %r13
	movq	%r13, 184(%rsp)         # 8-byte Spill
	adcq	$0, 144(%rsp)           # 8-byte Folded Spill
	adcq	$0, 104(%rsp)           # 8-byte Folded Spill
	adcq	$0, 152(%rsp)           # 8-byte Folded Spill
	adcq	$0, 112(%rsp)           # 8-byte Folded Spill
	adcq	$0, %r12
	movq	%rax, %rbx
	movq	%rbx, %rdx
	imulq	120(%rsp), %rdx         # 8-byte Folded Reload
	leaq	536(%rsp), %rdi
	movq	128(%rsp), %rsi         # 8-byte Reload
	callq	.LmulPv576x64
	addq	536(%rsp), %rbx
	movq	96(%rsp), %rax          # 8-byte Reload
	adcq	544(%rsp), %rax
	movq	136(%rsp), %rcx         # 8-byte Reload
	adcq	552(%rsp), %rcx
	movq	%rcx, 136(%rsp)         # 8-byte Spill
	movq	168(%rsp), %rcx         # 8-byte Reload
	adcq	560(%rsp), %rcx
	movq	%rcx, 168(%rsp)         # 8-byte Spill
	adcq	568(%rsp), %r15
	movq	%r15, 176(%rsp)         # 8-byte Spill
	adcq	576(%rsp), %r14
	movq	%r14, 192(%rsp)         # 8-byte Spill
	adcq	584(%rsp), %rbp
	movq	%rbp, 160(%rsp)         # 8-byte Spill
	movq	200(%rsp), %r13         # 8-byte Reload
	adcq	592(%rsp), %r13
	movq	208(%rsp), %r15         # 8-byte Reload
	adcq	600(%rsp), %r15
	movq	184(%rsp), %rbp         # 8-byte Reload
	adcq	608(%rsp), %rbp
	movq	144(%rsp), %rbx         # 8-byte Reload
	adcq	$0, %rbx
	adcq	$0, 104(%rsp)           # 8-byte Folded Spill
	adcq	$0, 152(%rsp)           # 8-byte Folded Spill
	adcq	$0, 112(%rsp)           # 8-byte Folded Spill
	adcq	$0, %r12
	movq	%rax, %rdx
	movq	%rax, %r14
	imulq	120(%rsp), %rdx         # 8-byte Folded Reload
	leaq	456(%rsp), %rdi
	movq	128(%rsp), %rsi         # 8-byte Reload
	callq	.LmulPv576x64
	addq	456(%rsp), %r14
	movq	136(%rsp), %rax         # 8-byte Reload
	adcq	464(%rsp), %rax
	movq	168(%rsp), %rcx         # 8-byte Reload
	adcq	472(%rsp), %rcx
	movq	%rcx, 168(%rsp)         # 8-byte Spill
	movq	176(%rsp), %rcx         # 8-byte Reload
	adcq	480(%rsp), %rcx
	movq	%rcx, 176(%rsp)         # 8-byte Spill
	movq	192(%rsp), %rcx         # 8-byte Reload
	adcq	488(%rsp), %rcx
	movq	%rcx, 192(%rsp)         # 8-byte Spill
	movq	160(%rsp), %rcx         # 8-byte Reload
	adcq	496(%rsp), %rcx
	movq	%rcx, 160(%rsp)         # 8-byte Spill
	adcq	504(%rsp), %r13
	movq	%r13, 200(%rsp)         # 8-byte Spill
	adcq	512(%rsp), %r15
	movq	%r15, 208(%rsp)         # 8-byte Spill
	adcq	520(%rsp), %rbp
	movq	%rbp, 184(%rsp)         # 8-byte Spill
	adcq	528(%rsp), %rbx
	movq	%rbx, 144(%rsp)         # 8-byte Spill
	movq	104(%rsp), %r14         # 8-byte Reload
	adcq	$0, %r14
	movq	152(%rsp), %r13         # 8-byte Reload
	adcq	$0, %r13
	movq	112(%rsp), %rbx         # 8-byte Reload
	adcq	$0, %rbx
	adcq	$0, %r12
	movq	%rax, %rdx
	movq	%rax, %r15
	imulq	120(%rsp), %rdx         # 8-byte Folded Reload
	leaq	376(%rsp), %rdi
	movq	128(%rsp), %rsi         # 8-byte Reload
	callq	.LmulPv576x64
	addq	376(%rsp), %r15
	movq	168(%rsp), %rax         # 8-byte Reload
	adcq	384(%rsp), %rax
	movq	176(%rsp), %rcx         # 8-byte Reload
	adcq	392(%rsp), %rcx
	movq	%rcx, 176(%rsp)         # 8-byte Spill
	movq	192(%rsp), %rcx         # 8-byte Reload
	adcq	400(%rsp), %rcx
	movq	%rcx, 192(%rsp)         # 8-byte Spill
	movq	160(%rsp), %rbp         # 8-byte Reload
	adcq	408(%rsp), %rbp
	movq	200(%rsp), %rcx         # 8-byte Reload
	adcq	416(%rsp), %rcx
	movq	%rcx, 200(%rsp)         # 8-byte Spill
	movq	208(%rsp), %rcx         # 8-byte Reload
	adcq	424(%rsp), %rcx
	movq	%rcx, 208(%rsp)         # 8-byte Spill
	movq	184(%rsp), %rcx         # 8-byte Reload
	adcq	432(%rsp), %rcx
	movq	%rcx, 184(%rsp)         # 8-byte Spill
	movq	144(%rsp), %r15         # 8-byte Reload
	adcq	440(%rsp), %r15
	adcq	448(%rsp), %r14
	movq	%r14, 104(%rsp)         # 8-byte Spill
	adcq	$0, %r13
	movq	%r13, %r14
	adcq	$0, %rbx
	movq	%rbx, 112(%rsp)         # 8-byte Spill
	adcq	$0, %r12
	movq	%rax, %rbx
	movq	%rbx, %rdx
	imulq	120(%rsp), %rdx         # 8-byte Folded Reload
	leaq	296(%rsp), %rdi
	movq	128(%rsp), %rsi         # 8-byte Reload
	callq	.LmulPv576x64
	addq	296(%rsp), %rbx
	movq	176(%rsp), %rax         # 8-byte Reload
	adcq	304(%rsp), %rax
	movq	192(%rsp), %r13         # 8-byte Reload
	adcq	312(%rsp), %r13
	adcq	320(%rsp), %rbp
	movq	200(%rsp), %rcx         # 8-byte Reload
	adcq	328(%rsp), %rcx
	movq	%rcx, 200(%rsp)         # 8-byte Spill
	movq	208(%rsp), %rcx         # 8-byte Reload
	adcq	336(%rsp), %rcx
	movq	%rcx, 208(%rsp)         # 8-byte Spill
	movq	184(%rsp), %rcx         # 8-byte Reload
	adcq	344(%rsp), %rcx
	movq	%rcx, 184(%rsp)         # 8-byte Spill
	adcq	352(%rsp), %r15
	movq	%r15, 144(%rsp)         # 8-byte Spill
	movq	104(%rsp), %r15         # 8-byte Reload
	adcq	360(%rsp), %r15
	adcq	368(%rsp), %r14
	movq	%r14, 152(%rsp)         # 8-byte Spill
	movq	112(%rsp), %r14         # 8-byte Reload
	adcq	$0, %r14
	adcq	$0, %r12
	movq	120(%rsp), %rdx         # 8-byte Reload
	imulq	%rax, %rdx
	movq	%rax, %rbx
	leaq	216(%rsp), %rdi
	movq	128(%rsp), %rsi         # 8-byte Reload
	callq	.LmulPv576x64
	addq	216(%rsp), %rbx
	movq	%r13, %rsi
	adcq	224(%rsp), %rsi
	movq	%rsi, 192(%rsp)         # 8-byte Spill
	adcq	232(%rsp), %rbp
	movq	%rbp, 160(%rsp)         # 8-byte Spill
	movq	200(%rsp), %r9          # 8-byte Reload
	adcq	240(%rsp), %r9
	movq	%r9, 200(%rsp)          # 8-byte Spill
	movq	208(%rsp), %r8          # 8-byte Reload
	adcq	248(%rsp), %r8
	movq	%r8, 208(%rsp)          # 8-byte Spill
	movq	184(%rsp), %rbx         # 8-byte Reload
	adcq	256(%rsp), %rbx
	movq	144(%rsp), %rax         # 8-byte Reload
	adcq	264(%rsp), %rax
	movq	%r15, %rcx
	adcq	272(%rsp), %rcx
	movq	152(%rsp), %rdx         # 8-byte Reload
	adcq	280(%rsp), %rdx
	movq	%rdx, 152(%rsp)         # 8-byte Spill
	adcq	288(%rsp), %r14
	movq	%r14, %r11
	adcq	$0, %r12
	subq	16(%rsp), %rsi          # 8-byte Folded Reload
	movq	%rbp, %rdi
	sbbq	8(%rsp), %rdi           # 8-byte Folded Reload
	movq	%r9, %rbp
	sbbq	24(%rsp), %rbp          # 8-byte Folded Reload
	movq	%r8, %r13
	sbbq	32(%rsp), %r13          # 8-byte Folded Reload
	movq	%rbx, %r15
	sbbq	40(%rsp), %r15          # 8-byte Folded Reload
	movq	%rax, %r14
	sbbq	48(%rsp), %r14          # 8-byte Folded Reload
	movq	%rcx, %r10
	sbbq	56(%rsp), %r10          # 8-byte Folded Reload
	movq	%rdx, %r8
	sbbq	64(%rsp), %r8           # 8-byte Folded Reload
	movq	%r11, %r9
	sbbq	72(%rsp), %r9           # 8-byte Folded Reload
	sbbq	$0, %r12
	andl	$1, %r12d
	cmovneq	%r11, %r9
	testb	%r12b, %r12b
	cmovneq	192(%rsp), %rsi         # 8-byte Folded Reload
	movq	80(%rsp), %rdx          # 8-byte Reload
	movq	%rsi, (%rdx)
	cmovneq	160(%rsp), %rdi         # 8-byte Folded Reload
	movq	%rdi, 8(%rdx)
	cmovneq	200(%rsp), %rbp         # 8-byte Folded Reload
	movq	%rbp, 16(%rdx)
	cmovneq	208(%rsp), %r13         # 8-byte Folded Reload
	movq	%r13, 24(%rdx)
	cmovneq	%rbx, %r15
	movq	%r15, 32(%rdx)
	cmovneq	%rax, %r14
	movq	%r14, 40(%rdx)
	cmovneq	%rcx, %r10
	movq	%r10, 48(%rdx)
	cmovneq	152(%rsp), %r8          # 8-byte Folded Reload
	movq	%r8, 56(%rdx)
	movq	%r9, 64(%rdx)
	addq	$936, %rsp              # imm = 0x3A8
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end132:
	.size	mcl_fp_montRed9Lbmi2, .Lfunc_end132-mcl_fp_montRed9Lbmi2

	.globl	mcl_fp_addPre9Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addPre9Lbmi2,@function
mcl_fp_addPre9Lbmi2:                    # @mcl_fp_addPre9Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	64(%rdx), %r8
	movq	64(%rsi), %r15
	movq	56(%rsi), %r9
	movq	48(%rsi), %r10
	movq	40(%rsi), %r11
	movq	24(%rsi), %r12
	movq	32(%rsi), %r14
	movq	(%rdx), %rbx
	movq	8(%rdx), %rcx
	addq	(%rsi), %rbx
	adcq	8(%rsi), %rcx
	movq	16(%rdx), %rax
	adcq	16(%rsi), %rax
	adcq	24(%rdx), %r12
	movq	56(%rdx), %r13
	movq	48(%rdx), %rsi
	movq	40(%rdx), %rbp
	movq	32(%rdx), %rdx
	movq	%rbx, (%rdi)
	movq	%rcx, 8(%rdi)
	movq	%rax, 16(%rdi)
	movq	%r12, 24(%rdi)
	adcq	%r14, %rdx
	movq	%rdx, 32(%rdi)
	adcq	%r11, %rbp
	movq	%rbp, 40(%rdi)
	adcq	%r10, %rsi
	movq	%rsi, 48(%rdi)
	adcq	%r9, %r13
	movq	%r13, 56(%rdi)
	adcq	%r8, %r15
	movq	%r15, 64(%rdi)
	sbbq	%rax, %rax
	andl	$1, %eax
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end133:
	.size	mcl_fp_addPre9Lbmi2, .Lfunc_end133-mcl_fp_addPre9Lbmi2

	.globl	mcl_fp_subPre9Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subPre9Lbmi2,@function
mcl_fp_subPre9Lbmi2:                    # @mcl_fp_subPre9Lbmi2
# BB#0:
	movq	32(%rdx), %r8
	movq	(%rsi), %rcx
	xorl	%eax, %eax
	subq	(%rdx), %rcx
	movq	%rcx, (%rdi)
	movq	8(%rsi), %rcx
	sbbq	8(%rdx), %rcx
	movq	%rcx, 8(%rdi)
	movq	16(%rsi), %rcx
	sbbq	16(%rdx), %rcx
	movq	%rcx, 16(%rdi)
	movq	24(%rsi), %rcx
	sbbq	24(%rdx), %rcx
	movq	%rcx, 24(%rdi)
	movq	32(%rsi), %rcx
	sbbq	%r8, %rcx
	movq	40(%rdx), %r8
	movq	%rcx, 32(%rdi)
	movq	40(%rsi), %rcx
	sbbq	%r8, %rcx
	movq	48(%rdx), %r8
	movq	%rcx, 40(%rdi)
	movq	48(%rsi), %rcx
	sbbq	%r8, %rcx
	movq	56(%rdx), %r8
	movq	%rcx, 48(%rdi)
	movq	56(%rsi), %rcx
	sbbq	%r8, %rcx
	movq	%rcx, 56(%rdi)
	movq	64(%rdx), %rcx
	movq	64(%rsi), %rdx
	sbbq	%rcx, %rdx
	movq	%rdx, 64(%rdi)
	sbbq	$0, %rax
	andl	$1, %eax
	retq
.Lfunc_end134:
	.size	mcl_fp_subPre9Lbmi2, .Lfunc_end134-mcl_fp_subPre9Lbmi2

	.globl	mcl_fp_shr1_9Lbmi2
	.align	16, 0x90
	.type	mcl_fp_shr1_9Lbmi2,@function
mcl_fp_shr1_9Lbmi2:                     # @mcl_fp_shr1_9Lbmi2
# BB#0:
	pushq	%rbx
	movq	64(%rsi), %r8
	movq	56(%rsi), %r9
	movq	48(%rsi), %r10
	movq	40(%rsi), %r11
	movq	32(%rsi), %rcx
	movq	24(%rsi), %rdx
	movq	16(%rsi), %rax
	movq	(%rsi), %rbx
	movq	8(%rsi), %rsi
	shrdq	$1, %rsi, %rbx
	movq	%rbx, (%rdi)
	shrdq	$1, %rax, %rsi
	movq	%rsi, 8(%rdi)
	shrdq	$1, %rdx, %rax
	movq	%rax, 16(%rdi)
	shrdq	$1, %rcx, %rdx
	movq	%rdx, 24(%rdi)
	shrdq	$1, %r11, %rcx
	movq	%rcx, 32(%rdi)
	shrdq	$1, %r10, %r11
	movq	%r11, 40(%rdi)
	shrdq	$1, %r9, %r10
	movq	%r10, 48(%rdi)
	shrdq	$1, %r8, %r9
	movq	%r9, 56(%rdi)
	shrq	%r8
	movq	%r8, 64(%rdi)
	popq	%rbx
	retq
.Lfunc_end135:
	.size	mcl_fp_shr1_9Lbmi2, .Lfunc_end135-mcl_fp_shr1_9Lbmi2

	.globl	mcl_fp_add9Lbmi2
	.align	16, 0x90
	.type	mcl_fp_add9Lbmi2,@function
mcl_fp_add9Lbmi2:                       # @mcl_fp_add9Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	64(%rdx), %r12
	movq	64(%rsi), %r8
	movq	56(%rsi), %r13
	movq	48(%rsi), %r9
	movq	40(%rsi), %r10
	movq	24(%rsi), %r14
	movq	32(%rsi), %r11
	movq	(%rdx), %rbx
	movq	8(%rdx), %r15
	addq	(%rsi), %rbx
	adcq	8(%rsi), %r15
	movq	16(%rdx), %rax
	adcq	16(%rsi), %rax
	adcq	24(%rdx), %r14
	adcq	32(%rdx), %r11
	adcq	40(%rdx), %r10
	movq	56(%rdx), %rsi
	adcq	48(%rdx), %r9
	movq	%rbx, (%rdi)
	movq	%r15, 8(%rdi)
	movq	%rax, 16(%rdi)
	movq	%r14, 24(%rdi)
	movq	%r11, 32(%rdi)
	movq	%r10, 40(%rdi)
	movq	%r9, 48(%rdi)
	adcq	%r13, %rsi
	movq	%rsi, 56(%rdi)
	adcq	%r12, %r8
	movq	%r8, 64(%rdi)
	sbbq	%rdx, %rdx
	andl	$1, %edx
	subq	(%rcx), %rbx
	sbbq	8(%rcx), %r15
	sbbq	16(%rcx), %rax
	sbbq	24(%rcx), %r14
	sbbq	32(%rcx), %r11
	sbbq	40(%rcx), %r10
	sbbq	48(%rcx), %r9
	sbbq	56(%rcx), %rsi
	sbbq	64(%rcx), %r8
	sbbq	$0, %rdx
	testb	$1, %dl
	jne	.LBB136_2
# BB#1:                                 # %nocarry
	movq	%rbx, (%rdi)
	movq	%r15, 8(%rdi)
	movq	%rax, 16(%rdi)
	movq	%r14, 24(%rdi)
	movq	%r11, 32(%rdi)
	movq	%r10, 40(%rdi)
	movq	%r9, 48(%rdi)
	movq	%rsi, 56(%rdi)
	movq	%r8, 64(%rdi)
.LBB136_2:                              # %carry
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	retq
.Lfunc_end136:
	.size	mcl_fp_add9Lbmi2, .Lfunc_end136-mcl_fp_add9Lbmi2

	.globl	mcl_fp_addNF9Lbmi2
	.align	16, 0x90
	.type	mcl_fp_addNF9Lbmi2,@function
mcl_fp_addNF9Lbmi2:                     # @mcl_fp_addNF9Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rdi, %r8
	movq	64(%rdx), %r10
	movq	56(%rdx), %r11
	movq	48(%rdx), %r9
	movq	40(%rdx), %rax
	movq	32(%rdx), %rdi
	movq	24(%rdx), %rbp
	movq	16(%rdx), %r15
	movq	(%rdx), %rbx
	movq	8(%rdx), %r13
	addq	(%rsi), %rbx
	adcq	8(%rsi), %r13
	adcq	16(%rsi), %r15
	adcq	24(%rsi), %rbp
	movq	%rbp, -40(%rsp)         # 8-byte Spill
	adcq	32(%rsi), %rdi
	movq	%rdi, -16(%rsp)         # 8-byte Spill
	adcq	40(%rsi), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	adcq	48(%rsi), %r9
	movq	%r9, -32(%rsp)          # 8-byte Spill
	movq	%r9, %rdi
	adcq	56(%rsi), %r11
	movq	%r11, -24(%rsp)         # 8-byte Spill
	movq	%r11, %rax
	adcq	64(%rsi), %r10
	movq	%r10, %r9
	movq	%rbx, %rsi
	subq	(%rcx), %rsi
	movq	%r13, %rdx
	sbbq	8(%rcx), %rdx
	movq	%r15, %r12
	sbbq	16(%rcx), %r12
	sbbq	24(%rcx), %rbp
	movq	-16(%rsp), %r14         # 8-byte Reload
	sbbq	32(%rcx), %r14
	movq	-8(%rsp), %r11          # 8-byte Reload
	sbbq	40(%rcx), %r11
	movq	%rdi, %r10
	sbbq	48(%rcx), %r10
	movq	%rax, %rdi
	sbbq	56(%rcx), %rdi
	movq	%r9, %rax
	sbbq	64(%rcx), %rax
	movq	%rax, %rcx
	sarq	$63, %rcx
	cmovsq	%rbx, %rsi
	movq	%rsi, (%r8)
	cmovsq	%r13, %rdx
	movq	%rdx, 8(%r8)
	cmovsq	%r15, %r12
	movq	%r12, 16(%r8)
	cmovsq	-40(%rsp), %rbp         # 8-byte Folded Reload
	movq	%rbp, 24(%r8)
	cmovsq	-16(%rsp), %r14         # 8-byte Folded Reload
	movq	%r14, 32(%r8)
	cmovsq	-8(%rsp), %r11          # 8-byte Folded Reload
	movq	%r11, 40(%r8)
	cmovsq	-32(%rsp), %r10         # 8-byte Folded Reload
	movq	%r10, 48(%r8)
	cmovsq	-24(%rsp), %rdi         # 8-byte Folded Reload
	movq	%rdi, 56(%r8)
	cmovsq	%r9, %rax
	movq	%rax, 64(%r8)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end137:
	.size	mcl_fp_addNF9Lbmi2, .Lfunc_end137-mcl_fp_addNF9Lbmi2

	.globl	mcl_fp_sub9Lbmi2
	.align	16, 0x90
	.type	mcl_fp_sub9Lbmi2,@function
mcl_fp_sub9Lbmi2:                       # @mcl_fp_sub9Lbmi2
# BB#0:
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	64(%rdx), %r13
	movq	(%rsi), %rax
	movq	8(%rsi), %r9
	xorl	%ebx, %ebx
	subq	(%rdx), %rax
	sbbq	8(%rdx), %r9
	movq	16(%rsi), %r10
	sbbq	16(%rdx), %r10
	movq	24(%rsi), %r11
	sbbq	24(%rdx), %r11
	movq	32(%rsi), %r12
	sbbq	32(%rdx), %r12
	movq	40(%rsi), %r14
	sbbq	40(%rdx), %r14
	movq	48(%rsi), %r15
	sbbq	48(%rdx), %r15
	movq	64(%rsi), %r8
	movq	56(%rsi), %rsi
	sbbq	56(%rdx), %rsi
	movq	%rax, (%rdi)
	movq	%r9, 8(%rdi)
	movq	%r10, 16(%rdi)
	movq	%r11, 24(%rdi)
	movq	%r12, 32(%rdi)
	movq	%r14, 40(%rdi)
	movq	%r15, 48(%rdi)
	movq	%rsi, 56(%rdi)
	sbbq	%r13, %r8
	movq	%r8, 64(%rdi)
	sbbq	$0, %rbx
	testb	$1, %bl
	je	.LBB138_2
# BB#1:                                 # %carry
	addq	(%rcx), %rax
	movq	%rax, (%rdi)
	movq	8(%rcx), %rax
	adcq	%r9, %rax
	movq	%rax, 8(%rdi)
	movq	16(%rcx), %rax
	adcq	%r10, %rax
	movq	%rax, 16(%rdi)
	movq	24(%rcx), %rax
	adcq	%r11, %rax
	movq	%rax, 24(%rdi)
	movq	32(%rcx), %rax
	adcq	%r12, %rax
	movq	%rax, 32(%rdi)
	movq	40(%rcx), %rax
	adcq	%r14, %rax
	movq	%rax, 40(%rdi)
	movq	48(%rcx), %rax
	adcq	%r15, %rax
	movq	%rax, 48(%rdi)
	movq	56(%rcx), %rax
	adcq	%rsi, %rax
	movq	%rax, 56(%rdi)
	movq	64(%rcx), %rax
	adcq	%r8, %rax
	movq	%rax, 64(%rdi)
.LBB138_2:                              # %nocarry
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	retq
.Lfunc_end138:
	.size	mcl_fp_sub9Lbmi2, .Lfunc_end138-mcl_fp_sub9Lbmi2

	.globl	mcl_fp_subNF9Lbmi2
	.align	16, 0x90
	.type	mcl_fp_subNF9Lbmi2,@function
mcl_fp_subNF9Lbmi2:                     # @mcl_fp_subNF9Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rcx, %r8
	movq	%rdi, %r10
	movq	64(%rsi), %r14
	movq	56(%rsi), %rax
	movq	48(%rsi), %rcx
	movq	40(%rsi), %rdi
	movq	32(%rsi), %rbp
	movq	24(%rsi), %rbx
	movq	16(%rsi), %r15
	movq	(%rsi), %r13
	movq	8(%rsi), %r12
	subq	(%rdx), %r13
	sbbq	8(%rdx), %r12
	sbbq	16(%rdx), %r15
	sbbq	24(%rdx), %rbx
	movq	%rbx, -40(%rsp)         # 8-byte Spill
	sbbq	32(%rdx), %rbp
	movq	%rbp, -32(%rsp)         # 8-byte Spill
	sbbq	40(%rdx), %rdi
	movq	%rdi, -24(%rsp)         # 8-byte Spill
	sbbq	48(%rdx), %rcx
	movq	%rcx, -16(%rsp)         # 8-byte Spill
	sbbq	56(%rdx), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	sbbq	64(%rdx), %r14
	movq	%r14, %rdx
	sarq	$63, %rdx
	movq	%rdx, %rbp
	shldq	$1, %r14, %rbp
	movq	24(%r8), %rbx
	andq	%rbp, %rbx
	movq	8(%r8), %rdi
	andq	%rbp, %rdi
	andq	(%r8), %rbp
	movq	64(%r8), %r11
	andq	%rdx, %r11
	rorxq	$63, %rdx, %rax
	andq	56(%r8), %rdx
	movq	48(%r8), %r9
	andq	%rax, %r9
	movq	40(%r8), %rsi
	andq	%rax, %rsi
	movq	32(%r8), %rcx
	andq	%rax, %rcx
	andq	16(%r8), %rax
	addq	%r13, %rbp
	adcq	%r12, %rdi
	movq	%rbp, (%r10)
	adcq	%r15, %rax
	movq	%rdi, 8(%r10)
	adcq	-40(%rsp), %rbx         # 8-byte Folded Reload
	movq	%rax, 16(%r10)
	movq	%rbx, 24(%r10)
	adcq	-32(%rsp), %rcx         # 8-byte Folded Reload
	movq	%rcx, 32(%r10)
	adcq	-24(%rsp), %rsi         # 8-byte Folded Reload
	movq	%rsi, 40(%r10)
	adcq	-16(%rsp), %r9          # 8-byte Folded Reload
	movq	%r9, 48(%r10)
	adcq	-8(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, 56(%r10)
	adcq	%r14, %r11
	movq	%r11, 64(%r10)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end139:
	.size	mcl_fp_subNF9Lbmi2, .Lfunc_end139-mcl_fp_subNF9Lbmi2

	.globl	mcl_fpDbl_add9Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_add9Lbmi2,@function
mcl_fpDbl_add9Lbmi2:                    # @mcl_fpDbl_add9Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rcx, %r15
	movq	136(%rdx), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	movq	128(%rdx), %rax
	movq	%rax, -40(%rsp)         # 8-byte Spill
	movq	120(%rdx), %r10
	movq	112(%rdx), %r11
	movq	24(%rsi), %rcx
	movq	32(%rsi), %r14
	movq	16(%rdx), %rbp
	movq	(%rdx), %rax
	movq	8(%rdx), %rbx
	addq	(%rsi), %rax
	adcq	8(%rsi), %rbx
	adcq	16(%rsi), %rbp
	adcq	24(%rdx), %rcx
	adcq	32(%rdx), %r14
	movq	104(%rdx), %r9
	movq	96(%rdx), %r13
	movq	%rax, (%rdi)
	movq	88(%rdx), %r8
	movq	%rbx, 8(%rdi)
	movq	80(%rdx), %r12
	movq	%rbp, 16(%rdi)
	movq	40(%rdx), %rax
	movq	%rcx, 24(%rdi)
	movq	40(%rsi), %rbp
	adcq	%rax, %rbp
	movq	48(%rdx), %rcx
	movq	%r14, 32(%rdi)
	movq	48(%rsi), %rax
	adcq	%rcx, %rax
	movq	56(%rdx), %r14
	movq	%rbp, 40(%rdi)
	movq	56(%rsi), %rbp
	adcq	%r14, %rbp
	movq	72(%rdx), %rcx
	movq	64(%rdx), %rdx
	movq	%rax, 48(%rdi)
	movq	64(%rsi), %rax
	adcq	%rdx, %rax
	movq	136(%rsi), %rbx
	movq	%rbp, 56(%rdi)
	movq	72(%rsi), %rbp
	adcq	%rcx, %rbp
	movq	128(%rsi), %rcx
	movq	%rax, 64(%rdi)
	movq	80(%rsi), %rdx
	adcq	%r12, %rdx
	movq	88(%rsi), %r12
	adcq	%r8, %r12
	movq	96(%rsi), %r14
	adcq	%r13, %r14
	movq	%r14, -48(%rsp)         # 8-byte Spill
	movq	104(%rsi), %rax
	adcq	%r9, %rax
	movq	%rax, -32(%rsp)         # 8-byte Spill
	movq	120(%rsi), %rax
	movq	112(%rsi), %rsi
	adcq	%r11, %rsi
	movq	%rsi, -24(%rsp)         # 8-byte Spill
	adcq	%r10, %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	adcq	-40(%rsp), %rcx         # 8-byte Folded Reload
	movq	%rcx, -40(%rsp)         # 8-byte Spill
	adcq	-8(%rsp), %rbx          # 8-byte Folded Reload
	movq	%rbx, -8(%rsp)          # 8-byte Spill
	sbbq	%r9, %r9
	andl	$1, %r9d
	movq	%rbp, %r10
	subq	(%r15), %r10
	movq	%rdx, %r11
	sbbq	8(%r15), %r11
	movq	%r12, %rbx
	sbbq	16(%r15), %rbx
	sbbq	24(%r15), %r14
	movq	-32(%rsp), %r13         # 8-byte Reload
	sbbq	32(%r15), %r13
	movq	-24(%rsp), %rsi         # 8-byte Reload
	sbbq	40(%r15), %rsi
	movq	-16(%rsp), %rax         # 8-byte Reload
	sbbq	48(%r15), %rax
	sbbq	56(%r15), %rcx
	movq	-8(%rsp), %r8           # 8-byte Reload
	sbbq	64(%r15), %r8
	sbbq	$0, %r9
	andl	$1, %r9d
	cmovneq	%rbp, %r10
	movq	%r10, 72(%rdi)
	testb	%r9b, %r9b
	cmovneq	%rdx, %r11
	movq	%r11, 80(%rdi)
	cmovneq	%r12, %rbx
	movq	%rbx, 88(%rdi)
	cmovneq	-48(%rsp), %r14         # 8-byte Folded Reload
	movq	%r14, 96(%rdi)
	cmovneq	-32(%rsp), %r13         # 8-byte Folded Reload
	movq	%r13, 104(%rdi)
	cmovneq	-24(%rsp), %rsi         # 8-byte Folded Reload
	movq	%rsi, 112(%rdi)
	cmovneq	-16(%rsp), %rax         # 8-byte Folded Reload
	movq	%rax, 120(%rdi)
	cmovneq	-40(%rsp), %rcx         # 8-byte Folded Reload
	movq	%rcx, 128(%rdi)
	cmovneq	-8(%rsp), %r8           # 8-byte Folded Reload
	movq	%r8, 136(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end140:
	.size	mcl_fpDbl_add9Lbmi2, .Lfunc_end140-mcl_fpDbl_add9Lbmi2

	.globl	mcl_fpDbl_sub9Lbmi2
	.align	16, 0x90
	.type	mcl_fpDbl_sub9Lbmi2,@function
mcl_fpDbl_sub9Lbmi2:                    # @mcl_fpDbl_sub9Lbmi2
# BB#0:
	pushq	%rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	movq	%rcx, %r14
	movq	136(%rdx), %rax
	movq	%rax, -8(%rsp)          # 8-byte Spill
	movq	128(%rdx), %rax
	movq	%rax, -16(%rsp)         # 8-byte Spill
	movq	120(%rdx), %rax
	movq	%rax, -24(%rsp)         # 8-byte Spill
	movq	16(%rsi), %r11
	movq	(%rsi), %r12
	movq	8(%rsi), %r13
	xorl	%r9d, %r9d
	subq	(%rdx), %r12
	sbbq	8(%rdx), %r13
	sbbq	16(%rdx), %r11
	movq	24(%rsi), %rbx
	sbbq	24(%rdx), %rbx
	movq	32(%rsi), %rbp
	sbbq	32(%rdx), %rbp
	movq	112(%rdx), %r10
	movq	104(%rdx), %rcx
	movq	%r12, (%rdi)
	movq	96(%rdx), %rax
	movq	%r13, 8(%rdi)
	movq	88(%rdx), %r13
	movq	%r11, 16(%rdi)
	movq	40(%rdx), %r11
	movq	%rbx, 24(%rdi)
	movq	40(%rsi), %rbx
	sbbq	%r11, %rbx
	movq	48(%rdx), %r11
	movq	%rbp, 32(%rdi)
	movq	48(%rsi), %rbp
	sbbq	%r11, %rbp
	movq	56(%rdx), %r11
	movq	%rbx, 40(%rdi)
	movq	56(%rsi), %rbx
	sbbq	%r11, %rbx
	movq	64(%rdx), %r11
	movq	%rbp, 48(%rdi)
	movq	64(%rsi), %rbp
	sbbq	%r11, %rbp
	movq	80(%rdx), %r8
	movq	72(%rdx), %r11
	movq	%rbx, 56(%rdi)
	movq	72(%rsi), %r15
	sbbq	%r11, %r15
	movq	136(%rsi), %rdx
	movq	%rbp, 64(%rdi)
	movq	80(%rsi), %rbp
	sbbq	%r8, %rbp
	movq	88(%rsi), %r12
	sbbq	%r13, %r12
	movq	96(%rsi), %r13
	sbbq	%rax, %r13
	movq	104(%rsi), %rax
	sbbq	%rcx, %rax
	movq	%rax, -40(%rsp)         # 8-byte Spill
	movq	112(%rsi), %rax
	sbbq	%r10, %rax
	movq	%rax, -32(%rsp)         # 8-byte Spill
	movq	128(%rsi), %rax
	movq	120(%rsi), %rcx
	sbbq	-24(%rsp), %rcx         # 8-byte Folded Reload
	movq	%rcx, -24(%rsp)         # 8-byte Spill
	sbbq	-16(%rsp), %rax         # 8-byte Folded Reload
	movq	%rax, -16(%rsp)         # 8-byte Spill
	sbbq	-8(%rsp), %rdx          # 8-byte Folded Reload
	movq	%rdx, -8(%rsp)          # 8-byte Spill
	movl	$0, %r8d
	sbbq	$0, %r8
	andl	$1, %r8d
	movq	(%r14), %r10
	cmoveq	%r9, %r10
	testb	%r8b, %r8b
	movq	16(%r14), %r8
	cmoveq	%r9, %r8
	movq	8(%r14), %rdx
	cmoveq	%r9, %rdx
	movq	64(%r14), %rbx
	cmoveq	%r9, %rbx
	movq	56(%r14), %r11
	cmoveq	%r9, %r11
	movq	48(%r14), %rsi
	cmoveq	%r9, %rsi
	movq	40(%r14), %rcx
	cmoveq	%r9, %rcx
	movq	32(%r14), %rax
	cmoveq	%r9, %rax
	cmovneq	24(%r14), %r9
	addq	%r15, %r10
	adcq	%rbp, %rdx
	movq	%r10, 72(%rdi)
	adcq	%r12, %r8
	movq	%rdx, 80(%rdi)
	adcq	%r13, %r9
	movq	%r8, 88(%rdi)
	movq	%r9, 96(%rdi)
	adcq	-40(%rsp), %rax         # 8-byte Folded Reload
	movq	%rax, 104(%rdi)
	adcq	-32(%rsp), %rcx         # 8-byte Folded Reload
	movq	%rcx, 112(%rdi)
	adcq	-24(%rsp), %rsi         # 8-byte Folded Reload
	movq	%rsi, 120(%rdi)
	adcq	-16(%rsp), %r11         # 8-byte Folded Reload
	movq	%r11, 128(%rdi)
	adcq	-8(%rsp), %rbx          # 8-byte Folded Reload
	movq	%rbx, 136(%rdi)
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
.Lfunc_end141:
	.size	mcl_fpDbl_sub9Lbmi2, .Lfunc_end141-mcl_fpDbl_sub9Lbmi2


	.section	".note.GNU-stack","",@progbits
