package instructions

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/instructions/constants"
	"fmt"
)

func NewInstruction(opcode byte) base.Instruction{
	switch opcode {
	case 0x00: return nil
	case 0x01 :return &constants.ACONST_NULL{}
	default:
		panic(fmt.Errorf("Unsupported opcode : 0x%x !",opcode))
	}
}
