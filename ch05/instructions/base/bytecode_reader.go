package base
type BytecodeReader struct {
	code []byte
	pc int
}
func (self *BytecodeReader) Reset(code []byte,pc int){
	self.code=code
	self.pc=pc
}
func (self *BytecodeReader) ReadUint8() uint8{
	i:=self.code[self.pc]
	self.pc++
	return i
}
func (self *BytecodeReader) ReadInt8() int8{
	return int8(self.ReadUint8())
}
func (self *BytecodeReader) ReadUint16() uint16{
	byte1 :=uint16(self.ReadUint8())
	byte2 :=uint16(self.ReadUint8())
	return byte1<<8 |byte2
}
func (self *BytecodeReader) ReadInt16() int16{
	return int16(self.ReadUint16())
}
func (self *BytecodeReader) ReaderInt32() int32 {
	byte1 := int32(self.ReadUint16())
	byte2 := int32(self.ReadUint16())
	return byte1<<byte1 | byte2
}
func (self *BytecodeReader) ReaderInt32s(n int32) []int32{
	ints :=make([]int32,n)
	for i:=range ints{
		ints[i] =self.ReaderInt32()
	}
	return ints
}
func (self *BytecodeReader) SkipPadding(){
	for self.pc%4!=0{
		self.ReadUint8()
	}
}
func (self *BytecodeReader) PC() int{
	return self.PC()
}
