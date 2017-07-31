#include "textflag.h"

TEXT	·drawNormalNRGBAToNRGBAFast(SB),0,$0-128
   MOVL alpha+48(FP), AX
   IMULL $0x00008081, AX
   MOVL AX, alpha+48(FP)

   MOVQ src+24(FP), SI
   ADDQ s0+64(FP), SI
   MOVQ dest+0(FP), DI
   ADDQ d0+56(FP), DI
   MOVQ y+72(FP), BX

   MOVQ dyDelta+136(FP), R12
   MOVQ syDelta+104(FP), R13
   MOVQ sxDelta+96(FP), R14

   MOVQ $0x0000808180818081, AX
   MOVQ AX, X14

   PXOR X15, X15

LOOPY:
   MOVQ sx0+80(FP), R15

   LOOPX:
      MOVL (SI)(R15*1), AX
      SHRL $24, AX
      CMPB AX, $0
      JZ NEXTX

      ANDQ $0xff, AX
      MULL alpha+48(FP)
      SHRQ $23, AX
      CMPB AX, $0
      JZ NEXTX

      MOVL AX, X0
      MOVL (DI)(R15*1), X1
      PSRLDQ $3, X1
      MOVQ $0xff, AX
      MOVQ AX, X2

      PUNPCKLBW X0, X0
      PUNPCKLWL X0, X0 // sa sa sa sa
      PXOR X2, X0 // sa sa sa 255-sa

      PUNPCKLBW X1, X1
      PUNPCKLWL X1, X1 // da da da da
      PSLLDQ $1, X2
      PXOR X2, X1 // 00 da 255-da da

      PUNPCKLBW X15, X0
      PUNPCKLBW X15, X1
      PMULLW X1, X0

      MOVL (SI)(R15*1), X2 // sa sr sg sb
      MOVL (DI)(R15*1), X3 // da dr dg db
      MOVOA X2, X4 // 00  r  g  b

      PMULHUW X14, X0

      PUNPCKLBW  X2, X3 // sa da sr dr sg dg sb db
      PUNPCKLBW X15, X4 // 00 00 00  r 00  g 00  b
      PUNPCKLWL  X4, X3 // 00 00 sa da 00  r sr dr 00  g sg dg 00  b sb db
      MOVOA X3, X2
      PSRLDQ $8, X2

      PSRLW $7, X0 // 0000 sa*da/255 sa*(255-da)/255 (255-sa)*da/255

      MOVOA X0, X1
      MOVOA X0, X5
      PSRLDQ $2, X5
      PADDW X5, X1
      PSRLDQ $2, X5
      PADDW X5, X1 // a

      PUNPCKLBW X15, X2 // 00 00 00 00 00 sa 00 da 00 00 00  r 00 sr 00 dr
      PUNPCKLBW X15, X3 // 00 00 00  g 00 sg 00 dg 00 00 00  b 00 sb 00 db
      PUNPCKLQDQ X0, X0
      PMULLW X0, X2
      PMULLW X0, X3

      MOVOA X2, X0
      PSRLDQ $2, X0
      PADDW X0, X2
      PSRLDQ $2, X0
      PADDW X0, X2 // 00 00 00 00 00 sa 00 da 00 00 00  r

      MOVOA X3, X0
      PSRLDQ $2, X0
      PADDW X0, X3
      PSRLDQ $2, X0
      PADDW X0, X3 // 00 00 00  g 00 sg 00 dg 00 00 00  b

      MOVQ $0xffff, AX
      MOVQ AX, X5
      INCL AX
      MOVQ AX, X4
      PAND X5, X1
      MOVL X1, DX
      CMPL DX, $2
      JB DIVEND
      LEAQ ·divTable(SB), AX
      MOVL (AX)(DX*4), X0

      PUNPCKLQDQ X0, X0
      PUNPCKLQDQ X4, X4
      PUNPCKLQDQ X5, X5

      PAND X5, X2
      PMULDQ X0, X2
      PADDQ X4, X2
      PSRLDQ $4, X2

      PAND X5, X3
      PMULDQ X0, X3
      PADDQ X4, X3
      PSRLDQ $4, X3

      DIVEND:
      MOVOA X3, X4
      PSRLDQ $8, X3

      PUNPCKLBW  X3, X4 // g b
      PUNPCKLBW  X1, X2 // a r
      PUNPCKLWL  X2, X4 // a r g b
      MOVL X4, (DI)(R15*1)

      NEXTX:
      ADDQ R14, R15
      CMPQ R15, sx1+88(FP)
      JNE LOOPX

   ADDQ R12, DI
   ADDQ R13, SI
   DECQ BX
   JNZ LOOPY
   RET
