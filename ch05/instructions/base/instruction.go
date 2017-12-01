package base

import "jvmgo/ch05/rtda"

type Instruction interface {
	/**
	获取操作数
	 */
	FetchOperands(reader *BytecodeReader)
	Excute(frame *rtda.Frame)
}
type NoOperandsInstruction struct {


}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}
/**
跳转指令
 */
type BranchInstruction struct {
	Offset int
}
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset=int(reader.ReadInt16())
}

/**
访问局部变量
 */
type Index8Instruction struct {
	Index uint
}
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index=uint(reader.ReadUint8())
}
/**
访问常量池
 */
type Index16Instruction struct {
	Index uint
}
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index=uint(reader.ReadUint16())
}
