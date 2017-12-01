package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}
func (self *NOP) Execute(fram *rtda.Frame){

	//do nothin
}