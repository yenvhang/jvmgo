package classfile
/**
 location: 字段表
 usage   : final 关键词定义的常量值
 */
type ConstantValueAttribute struct {
	constantValueIndex uint16
}
func (self * ConstantValueAttribute) readInfo (reader *ClassReader){
	self.constantValueIndex=reader.readUint16()
}
func (self * ConstantValueAttribute) ConstantValueIndex() uint16{
	return self.constantValueIndex
}
