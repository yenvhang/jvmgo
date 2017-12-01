package classfile

import "fmt"




type ClassFile struct {
	/*
	魔数 确定class文件能否被虚拟接收 值为：0xCAFEBABE
	 */
	magic uint32
	/*
	次版本号
	 */
	minorVersion uint16
	/*
	主版本号
	 */
	majorVersion uint16
	/*
	常量池
	 */
	constantPool ConstantPool
	/*
	访问标志
	 */
	accessFlags uint16
	/*
	类索引
	 */
	thisClass uint16
	/*
	父类索引
	 */
	superClass uint16
	/*
	接口索引集合
	 */
	interfaces []uint16
	/*
	字段表集合
	 */
	fields [] *MemberInfo
	/*
	方法表集合
	 */
	methods [] *MemberInfo
	/*
	属性表集合
	 */
	attributes []AttributeInfo
}

func  Parse(classData []byte)(cf *ClassFile,err error){
	defer func(){
		if r :=recover(); r!=nil{
			var ok bool
			err,ok =r.(error)
			if !ok{
				err =fmt.Errorf("%v",r)
			}
		}

	}()

	cr :=&ClassReader{classData}
	cf =&ClassFile{}
	cf.read(cr)
	return

}


func (self *ClassFile) read(reader *ClassReader){
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool=readConstantPool(reader)
	self.accessFlags=reader.readUint16()
	self.thisClass=reader.readUint16()
	self.superClass=reader.readUint16()
	self.interfaces=reader.readUint16s()
	self.fields=readMembers(reader,self.constantPool)
	self.methods=readMembers(reader,self.constantPool)
	self.attributes=readAttributes(reader,self.constantPool)
}
func (self *ClassFile) readAndCheckMagic(reader *ClassReader){
	magic :=reader.readUint32()
	if magic !=0xCAFEBABE{
		panic("java.lang.classFormatError: magic!")
	}
}
func (self  *ClassFile) readAndCheckVersion(reader *ClassReader){
	self.minorVersion =reader.readUint16()
	self.majorVersion =reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46,47,48,49,50,51,52:
		if self.minorVersion ==0{
			return
		}
	}
	panic ("java.lang.UnsupportedClassVersion")
}

func (self *ClassFile) MinorVersion() uint16{
	return self.minorVersion
}
func (self *ClassFile) MajorVersion() uint16{
	return self.majorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool{
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16{
	return self.accessFlags
}
func (self *ClassFile) Fields() []*MemberInfo{
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo{
	return self.methods
}
func (self *ClassFile) ClassName() string {
 return self.constantPool.getClassName(self.thisClass)
}
func (self *ClassFile) SuperClassName() string{
	if self.superClass >0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}
func (self *ClassFile) InterfaceNames() []string{
	interfaceNames := make([]string,len(self.interfaces))
	for i,cpIndex := range self.interfaces{
		interfaceNames[i] =self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}


