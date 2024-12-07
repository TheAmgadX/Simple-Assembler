# Assembler Document!  
         The Project is a very Simple Assembler that translate the assembly code to machine code for the Basic Computer Instruction Set Architecture as per M.Mano's book "Computer System Architecture"

         The Basic Computer has a 16-bit instruction divided into 12-bit address, 3-bit opcode and 1-bit for addressing mode.

##   Instructions set:
    - Memory Reference Instructions:

            direct - indirect [Addressing]
        and 000    - 1000
        add 001    - 1001
        lda 010    - 1010
        sta 011    - 1011
        bun 100    - 1100
        bsa 101    - 1101
        isz 110    - 1110

    - Register reference Instructions

        cla 0111100000000000
        cle 0111010000000000
        cma 0111001000000000
        cme 0111000100000000
        cir 0111000010000000 
        cil 0111000001000000
        inc 0111000000100000
        spa 0111000000010000
        sna 0111000000001000
        sza 0111000000000100 
        sze 0111000000000010
        hlt 0111000000000001
    
    - Input Output Instructions

        inp 1111100000000000
        out 1111010000000000
        ski 1111001000000000
        sko 1111000100000000
        ion 1111000010000000 
        iof 1111000001000000


There are only 4 pseudo-instructions supported by this assembler: `ORG`, `END`, `HEX` and `DEC`. These instructions do not have a direct binary mapping, but are instructions to the assembler to behave in a certain way during the first and second passes.

# First Test Case

        ORG     100     / Origin of program is location 0x100
        LDA     SUB     / Load subtrahend to AC
        CMA             / Complement AC
        INC             / Increment AC
        ADD     MIN     / Add minuend to AC
        STA     DIF     / Store difference
        HLT             / Halt computer

MIN,    DEC     83      / Minuend
SUB,    DEC     -23     / Subtrahend
DIF,    HEX     0       / Difference stored here
        END             / End of symbolic program

# Output of first test case

    -------------------- Symbol Table -----------------------
    MIN  :  106
    SUB  :  107
    DIF  :  108
    ---------------------------------------------------------
    -------------------- Machine Code -----------------------
    100  :  0010000100000111
    101  :  0111001000000000
    102  :  0111000000100000
    103  :  0001000100000110
    104  :  0011000100001000
    105  :  0111000000000001
    106  :  0000000001010011
    107  :  1111111111101001
    108  :  0000000000000000
    ---------------------------------------------------------
# Second Test Case 

ORG 0      /Origin of program is location 0
LDA A     /Load operand from location A
ADD B     /Add operand from location B
STA C     /Store sum in location C
HLT       /Halt computer

A, DEC 83     /Decimal operand
B, DEC -23    /Decimal operand
C, DEC 0      /Sum stored in location C

END      /End of symbolic program

# Output of second test case

    -------------------- Symbol Table -----------------------
    C  :  6
    A  :  4
    B  :  5
    ---------------------------------------------------------
    -------------------- Machine Code -----------------------
    0  :  0010000000000100
    1  :  0001000000000101
    2  :  0011000000000110
    3  :  0111000000000001
    4  :  0000000001010011
    5  :  1111111111101001
    6  :  0000000000000000
    ---------------------------------------------------------

# Third Test Case 

ORG 0      /Origin of program is location 0
LDA A I    /idirect addressing from Location A
ADD B     /Add operand from location B
STA C     /Store sum in location C
HLT       /Halt computer

A, DEC 83     /Decimal operand
B, DEC -23    /Decimal operand
C, DEC 0      /Sum stored in location C

END      /End of symbolic program

# Output of third test case

    -------------------- Symbol Table -----------------------
    A  :  4
    B  :  5
    C  :  6
    ---------------------------------------------------------
    -------------------- Machine Code -----------------------
    0  :  1010000001010011
    1  :  0001000000000101
    2  :  0011000000000110
    3  :  0111000000000001
    4  :  0000000001010011
    5  :  1111111111101001
    6  :  0000000000000000
    ---------------------------------------------------------
