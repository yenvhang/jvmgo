package classfile
/**
 location: 类文件
 usage   : 记录源文件名称
 */
type SourceFileAttribute struct {
	cp ConstantPool
	sourceFileIndex uint16
}
func (self *SourceFileAttribute) readInfo(reader *ClassReader){
	self.sourceFileIndex =reader.readUint16()
}
func (self *SourceFileAttribute) FileName() string  {
	return self.cp.getUtf8(self.sourceFileIndex)
}
