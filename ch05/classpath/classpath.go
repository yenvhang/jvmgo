package classpath

import (
	"path/filepath"
	"os"
)

type Classpath struct {
	bootClassPath Entry
	extClassPath Entry
	userClassPath Entry
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir,"lib","*")
	self.bootClassPath =newWildcardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir,"lib","ext","*")
	self.extClassPath =newWildcardEntry(jreExtPath)
}


func (self *Classpath) parseUserClasspath(cpOption string){
	if cpOption ==""{
		cpOption="."
	}
	self.userClassPath=newEntry(cpOption)
}
func getJreDir(jreOption string) string{
	if jreOption != "" && exists(jreOption){
		return jreOption
	}
	if exists("./jre"){
		return ".jre"
	}
	if jh :=os.Getenv("JAVA_HOME");jh!=""{
		return filepath.Join(jh,"jre")
	}
	panic("Can not find jre folder!")
}

func exists (path string) bool{
	if _,err := os.Stat(path); err!=nil{
		if os.IsNotExist(err){
			return false
		}

	}
	return true
}

func Parse(jreOption,cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}
func (self *Classpath) ReadClass(className string) ([]byte,Entry,error){
	className=className+".class"
	if data,entry,err := self.bootClassPath.readClass(className);err ==nil{
		return data,entry,err
	}
	if data,entry,err :=self.extClassPath.readClass(className) ; err ==nil{
		return data,entry,err
	}
	return self.userClassPath.readClass(className)
}
func (self *Classpath) String() string{
	return self.userClassPath.String()
}
