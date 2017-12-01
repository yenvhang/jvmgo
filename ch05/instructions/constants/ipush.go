package constants

import (

	"jvmgo/ch05/rtda"
	"jvmgo/ch05/instructions/base"
)

type BIPUSH struct {

	val int8

}
type SIPUSH struct {

	val int16
}
func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val=reader.ReadInt8()
}
func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val=reader.ReadInt16()
}
func (self *BIPUSH) Execute(frame *rtda.Frame){
	frame.OpoperandStack().PushInt(int32(self.val))
}
func (self *SIPUSH) Execute(frame *rtda.Frame){
	frame.OpoperandStack().PushInt(int32(self.val))
}