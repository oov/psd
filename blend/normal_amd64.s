#include "textflag.h"

TEXT	·drawNormalNRGBAToNRGBAFast(SB),0,$0-104
   MOVL alpha+48(FP), AX
   IMULL $0x00008081, AX
   MOVL AX, alpha+48(FP)

   MOVQ src+24(FP), SI
   MOVQ dest+0(FP), DI
   MOVQ y+56(FP), CX

   MOVQ dDelta+80(FP), R10
   MOVQ sDelta+88(FP), R11
   MOVQ xDelta+96(FP), R12

   MOVQ $0x0000808180818081, AX
   MOVQ AX, X14

   PXOR X15, X15

LOOPY:
   MOVQ xMin+64(FP), R15
   MOVQ SI, R13
   MOVQ DI, R14

   LOOPX:
      MOVL 0(R13), AX
      SHRL $24, AX
      CMPB AX, $0
      JZ NEXTX

      ANDQ $0xff, AX
      MULL alpha+48(FP)
      SHRQ $23, AX
      CMPB AX, $0
      JZ NEXTX

      MOVL $0xff, DX
      MOVQ DX, X4

      MOVL AX, X0
      MOVL 0(R14), X1
      PSRLDQ $3, X1

      PUNPCKLBW X0, X0
      PUNPCKLWL X0, X0 // sa sa sa sa
      PXOR X4, X0      // sa sa sa 255-sa

      PUNPCKLBW X1, X1
      PUNPCKLWL X1, X1 // da da da da
      PSRLDQ $1, X4
      PXOR X4, X1      // 00 da 255-da da

      PUNPCKLBW X15, X0
      PUNPCKLBW X15, X1
      PMULLW X1, X0
      PMULHUW X14, X0
      PSRLW $7, X0 // 0000 sa*da/255 sa*(255-da)/255 (255-sa)*da/255

      MOVOA X0, X1
      MOVOA X0, X5
      PSRLDQ $2, X5
      PADDW X5, X1
      PSRLDQ $2, X5
      PADDW X5, X1 // a

      MOVL 0(R13), X2 // sa sr sg sb
      MOVL 0(R14), X3 // da dr dg db
      MOVOA X2, X4 // 00  r  g  b

      PUNPCKLBW  X2, X3 // sa da sr dr sg dg sb db
      PUNPCKLBW X15, X4 // 00 00 00  r 00  g 00  b
      PUNPCKLWL  X4, X3 // 00 00 sa da 00  r sr dr 00  g sg dg 00  b sb db

      MOVOA X3, X2
      PSRLDQ $8, X2
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
      SHLL $2, DX
      LEAQ ·divTable<>(SB), AX
      ADDL DX, AX
      MOVL (AX), X0

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

      MOVOA X3, X4
      PSRLDQ $8, X3

      PUNPCKLBW  X3, X4 // g b
      PUNPCKLBW  X1, X2 // a r
      PUNPCKLWL  X2, X4 // a r g b
      MOVL X4, 0(R14)

      NEXTX:
      ADDQ R12, R13
      ADDQ R12, R14
      ADDQ R12, R15
      CMPL R15, xMax+72(FP)
      JNE LOOPX

   ADDQ R10, DI
   ADDQ R11, SI
   DECQ CX
   JNZ LOOPY
   RET

TEXT	·drawNormalNRGBAToNRGBAProtectAlphaFast(SB),0,$0-104
   MOVL alpha+48(FP), AX
   IMULL $0x00008081, AX
   MOVL AX, alpha+48(FP)

   MOVQ src+24(FP), SI
   MOVQ dest+0(FP), DI
   MOVQ y+56(FP), CX

   MOVQ dDelta+80(FP), R10
   MOVQ sDelta+88(FP), R11
   MOVQ xDelta+96(FP), R12

   PXOR X15, X15

LOOPY:
   MOVQ xMin+64(FP), R15
   MOVQ SI, R13
   MOVQ DI, R14

   LOOPX:
      MOVL 0(R13), AX
      SHRL $24, AX
      CMPB AX, $0
      JZ NEXTX

      ANDQ $0xff, AX
      MULL alpha+48(FP)
      SHRQ $23, AX
      CMPB AX, $0
      JZ NEXTX

      MOVL 0(R14), DX
      SHRL $24, DX
      CMPB DX, $0
      JZ NEXTX

      MOVL AX, X0 // sa
      MOVL DX, X1 // da

      MOVL $0xff, DX
      MOVQ DX, X4

      PUNPCKLBW X0, X0 // sa sa
      PXOR X4, X0 // sa 255-sa
      PUNPCKLBW X15, X0 // 00 sa 00 255-sa

      MOVL 0(R14), X2 // da dr dg db
      MOVL 0(R13), X3 //  a  r  g  b

      PUNPCKLBW  X3, X2 //  a da  r dr  g dg  b db
      PUNPCKLBW X15, X2 // 00  a 00 da 00  r 00 dr 00  g 00 dg 00  b 00 db

      PUNPCKLLQ X0, X0 // 00 a1 00 a3 00 a1 00 a3
      PUNPCKLQDQ X0, X0 // 00 a1 00 a3 00 a1 00 a3 00 a1 00 a3 00 a1 00 a3
      PMULLW X0, X2

      MOVOA X2, X0
      PSRLDQ $2, X0
      PADDW X0, X2      // 00 00 00  a 00 xx 00  r 00 xx 00  g 00 xx 00  b
      MOVOA X2, X3
      PUNPCKLLQ X15, X2 // 00 00 00 00 00 xx 00  g 00 00 00 00 00 xx 00  b
      PSRLDQ $8, X3     // 00 00 00 00 00 00 00 00 00 00 00  a 00 xx 00  r

      MOVQ $0xffff, AX
      MOVQ AX, X5
      INCL AX
      MOVQ AX, X4
      PAND X5, X1
      MOVL X1, DX
      SHLL $2, DX
      LEAQ ·divTable<>(SB), AX
      ADDL DX, AX
      MOVL (AX), X0

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

      MOVOA X2, X4
      PSRLDQ $8, X2

      PUNPCKLBW  X2, X4 // g b
      PUNPCKLBW  X1, X3 // a r
      PUNPCKLWL  X3, X4 // a r g b
      MOVL X4, 0(R14)

      NEXTX:
      ADDQ R12, R13
      ADDQ R12, R14
      ADDQ R12, R15
      CMPL R15, xMax+72(FP)
      JNE LOOPX

   ADDQ R10, DI
   ADDQ R11, SI
   DECQ CX
   JNZ LOOPY
   RET


// func main() {
// 	for i := 1; i < 256; i++ {
// 		fmt.Printf("0x%08x, // %d\n", 0xffffffff/i, i)
// 	}
// }
DATA ·divTable<>+0x0000(SB)/4, $0x00000000 // 0
DATA ·divTable<>+0x0004(SB)/4, $0xffffffff // 1
DATA ·divTable<>+0x0008(SB)/4, $0x7fffffff // 2
DATA ·divTable<>+0x000c(SB)/4, $0x55555555 // 3
DATA ·divTable<>+0x0010(SB)/4, $0x3fffffff // 4
DATA ·divTable<>+0x0014(SB)/4, $0x33333333 // 5
DATA ·divTable<>+0x0018(SB)/4, $0x2aaaaaaa // 6
DATA ·divTable<>+0x001c(SB)/4, $0x24924924 // 7
DATA ·divTable<>+0x0020(SB)/4, $0x1fffffff // 8
DATA ·divTable<>+0x0024(SB)/4, $0x1c71c71c // 9
DATA ·divTable<>+0x0028(SB)/4, $0x19999999 // 10
DATA ·divTable<>+0x002c(SB)/4, $0x1745d174 // 11
DATA ·divTable<>+0x0030(SB)/4, $0x15555555 // 12
DATA ·divTable<>+0x0034(SB)/4, $0x13b13b13 // 13
DATA ·divTable<>+0x0038(SB)/4, $0x12492492 // 14
DATA ·divTable<>+0x003c(SB)/4, $0x11111111 // 15
DATA ·divTable<>+0x0040(SB)/4, $0x0fffffff // 16
DATA ·divTable<>+0x0044(SB)/4, $0x0f0f0f0f // 17
DATA ·divTable<>+0x0048(SB)/4, $0x0e38e38e // 18
DATA ·divTable<>+0x004c(SB)/4, $0x0d79435e // 19
DATA ·divTable<>+0x0050(SB)/4, $0x0ccccccc // 20
DATA ·divTable<>+0x0054(SB)/4, $0x0c30c30c // 21
DATA ·divTable<>+0x0058(SB)/4, $0x0ba2e8ba // 22
DATA ·divTable<>+0x005c(SB)/4, $0x0b21642c // 23
DATA ·divTable<>+0x0060(SB)/4, $0x0aaaaaaa // 24
DATA ·divTable<>+0x0064(SB)/4, $0x0a3d70a3 // 25
DATA ·divTable<>+0x0068(SB)/4, $0x09d89d89 // 26
DATA ·divTable<>+0x006c(SB)/4, $0x097b425e // 27
DATA ·divTable<>+0x0070(SB)/4, $0x09249249 // 28
DATA ·divTable<>+0x0074(SB)/4, $0x08d3dcb0 // 29
DATA ·divTable<>+0x0078(SB)/4, $0x08888888 // 30
DATA ·divTable<>+0x007c(SB)/4, $0x08421084 // 31
DATA ·divTable<>+0x0080(SB)/4, $0x07ffffff // 32
DATA ·divTable<>+0x0084(SB)/4, $0x07c1f07c // 33
DATA ·divTable<>+0x0088(SB)/4, $0x07878787 // 34
DATA ·divTable<>+0x008c(SB)/4, $0x07507507 // 35
DATA ·divTable<>+0x0090(SB)/4, $0x071c71c7 // 36
DATA ·divTable<>+0x0094(SB)/4, $0x06eb3e45 // 37
DATA ·divTable<>+0x0098(SB)/4, $0x06bca1af // 38
DATA ·divTable<>+0x009c(SB)/4, $0x06906906 // 39
DATA ·divTable<>+0x00a0(SB)/4, $0x06666666 // 40
DATA ·divTable<>+0x00a4(SB)/4, $0x063e7063 // 41
DATA ·divTable<>+0x00a8(SB)/4, $0x06186186 // 42
DATA ·divTable<>+0x00ac(SB)/4, $0x05f417d0 // 43
DATA ·divTable<>+0x00b0(SB)/4, $0x05d1745d // 44
DATA ·divTable<>+0x00b4(SB)/4, $0x05b05b05 // 45
DATA ·divTable<>+0x00b8(SB)/4, $0x0590b216 // 46
DATA ·divTable<>+0x00bc(SB)/4, $0x0572620a // 47
DATA ·divTable<>+0x00c0(SB)/4, $0x05555555 // 48
DATA ·divTable<>+0x00c4(SB)/4, $0x05397829 // 49
DATA ·divTable<>+0x00c8(SB)/4, $0x051eb851 // 50
DATA ·divTable<>+0x00cc(SB)/4, $0x05050505 // 51
DATA ·divTable<>+0x00d0(SB)/4, $0x04ec4ec4 // 52
DATA ·divTable<>+0x00d4(SB)/4, $0x04d4873e // 53
DATA ·divTable<>+0x00d8(SB)/4, $0x04bda12f // 54
DATA ·divTable<>+0x00dc(SB)/4, $0x04a7904a // 55
DATA ·divTable<>+0x00e0(SB)/4, $0x04924924 // 56
DATA ·divTable<>+0x00e4(SB)/4, $0x047dc11f // 57
DATA ·divTable<>+0x00e8(SB)/4, $0x0469ee58 // 58
DATA ·divTable<>+0x00ec(SB)/4, $0x0456c797 // 59
DATA ·divTable<>+0x00f0(SB)/4, $0x04444444 // 60
DATA ·divTable<>+0x00f4(SB)/4, $0x04325c53 // 61
DATA ·divTable<>+0x00f8(SB)/4, $0x04210842 // 62
DATA ·divTable<>+0x00fc(SB)/4, $0x04104104 // 63
DATA ·divTable<>+0x0100(SB)/4, $0x03ffffff // 64
DATA ·divTable<>+0x0104(SB)/4, $0x03f03f03 // 65
DATA ·divTable<>+0x0108(SB)/4, $0x03e0f83e // 66
DATA ·divTable<>+0x010c(SB)/4, $0x03d22635 // 67
DATA ·divTable<>+0x0110(SB)/4, $0x03c3c3c3 // 68
DATA ·divTable<>+0x0114(SB)/4, $0x03b5cc0e // 69
DATA ·divTable<>+0x0118(SB)/4, $0x03a83a83 // 70
DATA ·divTable<>+0x011c(SB)/4, $0x039b0ad1 // 71
DATA ·divTable<>+0x0120(SB)/4, $0x038e38e3 // 72
DATA ·divTable<>+0x0124(SB)/4, $0x0381c0e0 // 73
DATA ·divTable<>+0x0128(SB)/4, $0x03759f22 // 74
DATA ·divTable<>+0x012c(SB)/4, $0x0369d036 // 75
DATA ·divTable<>+0x0130(SB)/4, $0x035e50d7 // 76
DATA ·divTable<>+0x0134(SB)/4, $0x03531dec // 77
DATA ·divTable<>+0x0138(SB)/4, $0x03483483 // 78
DATA ·divTable<>+0x013c(SB)/4, $0x033d91d2 // 79
DATA ·divTable<>+0x0140(SB)/4, $0x03333333 // 80
DATA ·divTable<>+0x0144(SB)/4, $0x0329161f // 81
DATA ·divTable<>+0x0148(SB)/4, $0x031f3831 // 82
DATA ·divTable<>+0x014c(SB)/4, $0x03159721 // 83
DATA ·divTable<>+0x0150(SB)/4, $0x030c30c3 // 84
DATA ·divTable<>+0x0154(SB)/4, $0x03030303 // 85
DATA ·divTable<>+0x0158(SB)/4, $0x02fa0be8 // 86
DATA ·divTable<>+0x015c(SB)/4, $0x02f14990 // 87
DATA ·divTable<>+0x0160(SB)/4, $0x02e8ba2e // 88
DATA ·divTable<>+0x0164(SB)/4, $0x02e05c0b // 89
DATA ·divTable<>+0x0168(SB)/4, $0x02d82d82 // 90
DATA ·divTable<>+0x016c(SB)/4, $0x02d02d02 // 91
DATA ·divTable<>+0x0170(SB)/4, $0x02c8590b // 92
DATA ·divTable<>+0x0174(SB)/4, $0x02c0b02c // 93
DATA ·divTable<>+0x0178(SB)/4, $0x02b93105 // 94
DATA ·divTable<>+0x017c(SB)/4, $0x02b1da46 // 95
DATA ·divTable<>+0x0180(SB)/4, $0x02aaaaaa // 96
DATA ·divTable<>+0x0184(SB)/4, $0x02a3a0fd // 97
DATA ·divTable<>+0x0188(SB)/4, $0x029cbc14 // 98
DATA ·divTable<>+0x018c(SB)/4, $0x0295fad4 // 99
DATA ·divTable<>+0x0190(SB)/4, $0x028f5c28 // 100
DATA ·divTable<>+0x0194(SB)/4, $0x0288df0c // 101
DATA ·divTable<>+0x0198(SB)/4, $0x02828282 // 102
DATA ·divTable<>+0x019c(SB)/4, $0x027c4597 // 103
DATA ·divTable<>+0x01a0(SB)/4, $0x02762762 // 104
DATA ·divTable<>+0x01a4(SB)/4, $0x02702702 // 105
DATA ·divTable<>+0x01a8(SB)/4, $0x026a439f // 106
DATA ·divTable<>+0x01ac(SB)/4, $0x02647c69 // 107
DATA ·divTable<>+0x01b0(SB)/4, $0x025ed097 // 108
DATA ·divTable<>+0x01b4(SB)/4, $0x02593f69 // 109
DATA ·divTable<>+0x01b8(SB)/4, $0x0253c825 // 110
DATA ·divTable<>+0x01bc(SB)/4, $0x024e6a17 // 111
DATA ·divTable<>+0x01c0(SB)/4, $0x02492492 // 112
DATA ·divTable<>+0x01c4(SB)/4, $0x0243f6f0 // 113
DATA ·divTable<>+0x01c8(SB)/4, $0x023ee08f // 114
DATA ·divTable<>+0x01cc(SB)/4, $0x0239e0d5 // 115
DATA ·divTable<>+0x01d0(SB)/4, $0x0234f72c // 116
DATA ·divTable<>+0x01d4(SB)/4, $0x02302302 // 117
DATA ·divTable<>+0x01d8(SB)/4, $0x022b63cb // 118
DATA ·divTable<>+0x01dc(SB)/4, $0x0226b902 // 119
DATA ·divTable<>+0x01e0(SB)/4, $0x02222222 // 120
DATA ·divTable<>+0x01e4(SB)/4, $0x021d9ead // 121
DATA ·divTable<>+0x01e8(SB)/4, $0x02192e29 // 122
DATA ·divTable<>+0x01ec(SB)/4, $0x0214d021 // 123
DATA ·divTable<>+0x01f0(SB)/4, $0x02108421 // 124
DATA ·divTable<>+0x01f4(SB)/4, $0x020c49ba // 125
DATA ·divTable<>+0x01f8(SB)/4, $0x02082082 // 126
DATA ·divTable<>+0x01fc(SB)/4, $0x02040810 // 127
DATA ·divTable<>+0x0200(SB)/4, $0x01ffffff // 128
DATA ·divTable<>+0x0204(SB)/4, $0x01fc07f0 // 129
DATA ·divTable<>+0x0208(SB)/4, $0x01f81f81 // 130
DATA ·divTable<>+0x020c(SB)/4, $0x01f44659 // 131
DATA ·divTable<>+0x0210(SB)/4, $0x01f07c1f // 132
DATA ·divTable<>+0x0214(SB)/4, $0x01ecc07b // 133
DATA ·divTable<>+0x0218(SB)/4, $0x01e9131a // 134
DATA ·divTable<>+0x021c(SB)/4, $0x01e573ac // 135
DATA ·divTable<>+0x0220(SB)/4, $0x01e1e1e1 // 136
DATA ·divTable<>+0x0224(SB)/4, $0x01de5d6e // 137
DATA ·divTable<>+0x0228(SB)/4, $0x01dae607 // 138
DATA ·divTable<>+0x022c(SB)/4, $0x01d77b65 // 139
DATA ·divTable<>+0x0230(SB)/4, $0x01d41d41 // 140
DATA ·divTable<>+0x0234(SB)/4, $0x01d0cb58 // 141
DATA ·divTable<>+0x0238(SB)/4, $0x01cd8568 // 142
DATA ·divTable<>+0x023c(SB)/4, $0x01ca4b30 // 143
DATA ·divTable<>+0x0240(SB)/4, $0x01c71c71 // 144
DATA ·divTable<>+0x0244(SB)/4, $0x01c3f8f0 // 145
DATA ·divTable<>+0x0248(SB)/4, $0x01c0e070 // 146
DATA ·divTable<>+0x024c(SB)/4, $0x01bdd2b8 // 147
DATA ·divTable<>+0x0250(SB)/4, $0x01bacf91 // 148
DATA ·divTable<>+0x0254(SB)/4, $0x01b7d6c3 // 149
DATA ·divTable<>+0x0258(SB)/4, $0x01b4e81b // 150
DATA ·divTable<>+0x025c(SB)/4, $0x01b20364 // 151
DATA ·divTable<>+0x0260(SB)/4, $0x01af286b // 152
DATA ·divTable<>+0x0264(SB)/4, $0x01ac5701 // 153
DATA ·divTable<>+0x0268(SB)/4, $0x01a98ef6 // 154
DATA ·divTable<>+0x026c(SB)/4, $0x01a6d01a // 155
DATA ·divTable<>+0x0270(SB)/4, $0x01a41a41 // 156
DATA ·divTable<>+0x0274(SB)/4, $0x01a16d3f // 157
DATA ·divTable<>+0x0278(SB)/4, $0x019ec8e9 // 158
DATA ·divTable<>+0x027c(SB)/4, $0x019c2d14 // 159
DATA ·divTable<>+0x0280(SB)/4, $0x01999999 // 160
DATA ·divTable<>+0x0284(SB)/4, $0x01970e4f // 161
DATA ·divTable<>+0x0288(SB)/4, $0x01948b0f // 162
DATA ·divTable<>+0x028c(SB)/4, $0x01920fb4 // 163
DATA ·divTable<>+0x0290(SB)/4, $0x018f9c18 // 164
DATA ·divTable<>+0x0294(SB)/4, $0x018d3018 // 165
DATA ·divTable<>+0x0298(SB)/4, $0x018acb90 // 166
DATA ·divTable<>+0x029c(SB)/4, $0x01886e5f // 167
DATA ·divTable<>+0x02a0(SB)/4, $0x01861861 // 168
DATA ·divTable<>+0x02a4(SB)/4, $0x0183c977 // 169
DATA ·divTable<>+0x02a8(SB)/4, $0x01818181 // 170
DATA ·divTable<>+0x02ac(SB)/4, $0x017f405f // 171
DATA ·divTable<>+0x02b0(SB)/4, $0x017d05f4 // 172
DATA ·divTable<>+0x02b4(SB)/4, $0x017ad220 // 173
DATA ·divTable<>+0x02b8(SB)/4, $0x0178a4c8 // 174
DATA ·divTable<>+0x02bc(SB)/4, $0x01767dce // 175
DATA ·divTable<>+0x02c0(SB)/4, $0x01745d17 // 176
DATA ·divTable<>+0x02c4(SB)/4, $0x01724287 // 177
DATA ·divTable<>+0x02c8(SB)/4, $0x01702e05 // 178
DATA ·divTable<>+0x02cc(SB)/4, $0x016e1f76 // 179
DATA ·divTable<>+0x02d0(SB)/4, $0x016c16c1 // 180
DATA ·divTable<>+0x02d4(SB)/4, $0x016a13cd // 181
DATA ·divTable<>+0x02d8(SB)/4, $0x01681681 // 182
DATA ·divTable<>+0x02dc(SB)/4, $0x01661ec6 // 183
DATA ·divTable<>+0x02e0(SB)/4, $0x01642c85 // 184
DATA ·divTable<>+0x02e4(SB)/4, $0x01623fa7 // 185
DATA ·divTable<>+0x02e8(SB)/4, $0x01605816 // 186
DATA ·divTable<>+0x02ec(SB)/4, $0x015e75bb // 187
DATA ·divTable<>+0x02f0(SB)/4, $0x015c9882 // 188
DATA ·divTable<>+0x02f4(SB)/4, $0x015ac056 // 189
DATA ·divTable<>+0x02f8(SB)/4, $0x0158ed23 // 190
DATA ·divTable<>+0x02fc(SB)/4, $0x01571ed3 // 191
DATA ·divTable<>+0x0300(SB)/4, $0x01555555 // 192
DATA ·divTable<>+0x0304(SB)/4, $0x01539094 // 193
DATA ·divTable<>+0x0308(SB)/4, $0x0151d07e // 194
DATA ·divTable<>+0x030c(SB)/4, $0x01501501 // 195
DATA ·divTable<>+0x0310(SB)/4, $0x014e5e0a // 196
DATA ·divTable<>+0x0314(SB)/4, $0x014cab88 // 197
DATA ·divTable<>+0x0318(SB)/4, $0x014afd6a // 198
DATA ·divTable<>+0x031c(SB)/4, $0x0149539e // 199
DATA ·divTable<>+0x0320(SB)/4, $0x0147ae14 // 200
DATA ·divTable<>+0x0324(SB)/4, $0x01460cbc // 201
DATA ·divTable<>+0x0328(SB)/4, $0x01446f86 // 202
DATA ·divTable<>+0x032c(SB)/4, $0x0142d662 // 203
DATA ·divTable<>+0x0330(SB)/4, $0x01414141 // 204
DATA ·divTable<>+0x0334(SB)/4, $0x013fb013 // 205
DATA ·divTable<>+0x0338(SB)/4, $0x013e22cb // 206
DATA ·divTable<>+0x033c(SB)/4, $0x013c995a // 207
DATA ·divTable<>+0x0340(SB)/4, $0x013b13b1 // 208
DATA ·divTable<>+0x0344(SB)/4, $0x013991c2 // 209
DATA ·divTable<>+0x0348(SB)/4, $0x01381381 // 210
DATA ·divTable<>+0x034c(SB)/4, $0x013698df // 211
DATA ·divTable<>+0x0350(SB)/4, $0x013521cf // 212
DATA ·divTable<>+0x0354(SB)/4, $0x0133ae45 // 213
DATA ·divTable<>+0x0358(SB)/4, $0x01323e34 // 214
DATA ·divTable<>+0x035c(SB)/4, $0x0130d190 // 215
DATA ·divTable<>+0x0360(SB)/4, $0x012f684b // 216
DATA ·divTable<>+0x0364(SB)/4, $0x012e025c // 217
DATA ·divTable<>+0x0368(SB)/4, $0x012c9fb4 // 218
DATA ·divTable<>+0x036c(SB)/4, $0x012b404a // 219
DATA ·divTable<>+0x0370(SB)/4, $0x0129e412 // 220
DATA ·divTable<>+0x0374(SB)/4, $0x01288b01 // 221
DATA ·divTable<>+0x0378(SB)/4, $0x0127350b // 222
DATA ·divTable<>+0x037c(SB)/4, $0x0125e227 // 223
DATA ·divTable<>+0x0380(SB)/4, $0x01249249 // 224
DATA ·divTable<>+0x0384(SB)/4, $0x01234567 // 225
DATA ·divTable<>+0x0388(SB)/4, $0x0121fb78 // 226
DATA ·divTable<>+0x038c(SB)/4, $0x0120b470 // 227
DATA ·divTable<>+0x0390(SB)/4, $0x011f7047 // 228
DATA ·divTable<>+0x0394(SB)/4, $0x011e2ef3 // 229
DATA ·divTable<>+0x0398(SB)/4, $0x011cf06a // 230
DATA ·divTable<>+0x039c(SB)/4, $0x011bb4a4 // 231
DATA ·divTable<>+0x03a0(SB)/4, $0x011a7b96 // 232
DATA ·divTable<>+0x03a4(SB)/4, $0x01194538 // 233
DATA ·divTable<>+0x03a8(SB)/4, $0x01181181 // 234
DATA ·divTable<>+0x03ac(SB)/4, $0x0116e068 // 235
DATA ·divTable<>+0x03b0(SB)/4, $0x0115b1e5 // 236
DATA ·divTable<>+0x03b4(SB)/4, $0x011485f0 // 237
DATA ·divTable<>+0x03b8(SB)/4, $0x01135c81 // 238
DATA ·divTable<>+0x03bc(SB)/4, $0x0112358e // 239
DATA ·divTable<>+0x03c0(SB)/4, $0x01111111 // 240
DATA ·divTable<>+0x03c4(SB)/4, $0x010fef01 // 241
DATA ·divTable<>+0x03c8(SB)/4, $0x010ecf56 // 242
DATA ·divTable<>+0x03cc(SB)/4, $0x010db20a // 243
DATA ·divTable<>+0x03d0(SB)/4, $0x010c9714 // 244
DATA ·divTable<>+0x03d4(SB)/4, $0x010b7e6e // 245
DATA ·divTable<>+0x03d8(SB)/4, $0x010a6810 // 246
DATA ·divTable<>+0x03dc(SB)/4, $0x010953f3 // 247
DATA ·divTable<>+0x03e0(SB)/4, $0x01084210 // 248
DATA ·divTable<>+0x03e4(SB)/4, $0x01073260 // 249
DATA ·divTable<>+0x03e8(SB)/4, $0x010624dd // 250
DATA ·divTable<>+0x03ec(SB)/4, $0x0105197f // 251
DATA ·divTable<>+0x03f0(SB)/4, $0x01041041 // 252
DATA ·divTable<>+0x03f4(SB)/4, $0x0103091b // 253
DATA ·divTable<>+0x03f8(SB)/4, $0x01020408 // 254
DATA ·divTable<>+0x03fc(SB)/4, $0x01010101 // 255
GLOBL ·divTable<>(SB),NOPTR+RODATA,$0x400
