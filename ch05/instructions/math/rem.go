package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
	"math"
)

type DREM struct {
	base.NoOperandsInstruction
}
type FREM struct {
	base.NoOperandsInstruction
}
type LREM struct {
	base.NoOperandsInstruction
}
type IREM struct {
	base.NoOperandsInstruction
}

func (self *IREM) Execute(frame *rtda.Frame){
	stack:=frame.OpoperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopInt()
	if v2==0{
		panic("java.lang.ArithmeticException: /by zero")
	}
	stack.PushInt(v1%v2)
}
func (self *DREM) Execute(frame *rtda.Frame){
	stack:=frame.OpoperandStack()
	v2:=stack.PopDouble()
	v1:=stack.PopDouble()
	if v2==0{
		panic("java.lang.ArithmeticException: /by zero")
	}
	stack.PushDouble(math.Mod(v1,v2))
}


