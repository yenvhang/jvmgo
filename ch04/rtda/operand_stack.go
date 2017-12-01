package rtda

import "math"
type OperandStack struct {
	size uint
	slots []Slot
}
func newOpreandStack(maxStack uint) *OperandStack{
	if maxStack >0{
		return &OperandStack{
			slots:make([]Slot,maxStack),
		}
	}
	return nil
}
func (self *OperandStack) PushInt(val int32){
	self.slots[self.size].num=val
	self.size++
}
func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}
func (self *OperandStack) PushFloat(val float32)  {
	bits:=math.Float32bits(val)
	self.slots[self.size].num =int32(bits)
	self.size++
}
func (self * OperandStack) PopFloat() float32 {
	self.size--
	bits:=uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}
func (self *OperandStack) PushRef(ref *Object){
	self.slots[self.size].ref=ref
	self.size++
}
func (self *OperandStack) PopRef() *Object {
	self.size--
	ref:= self.slots[self.size].ref
	self.slots[self.size].ref=nil
	return ref
}