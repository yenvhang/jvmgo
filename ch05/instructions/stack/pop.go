package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type POP struct {
	base.NoOperandsInstruction
}
type POP2 struct {
	base.NoOperandsInstruction
}
func (self *POP) Execute(fram *rtda.Frame){
	stack :=fram.OpoperandStack()
	stack.PopSlot()
}
func (self *POP2) Execute(frame *rtda.Frame){
	stack :=frame.OpoperandStack()
	stack.PopSlot()
	stack.PopSlot()
}