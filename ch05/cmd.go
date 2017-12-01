package ch05

import "flag"
import "fmt"
import "os"
import (
	"jvmgo/ch05/classpath"
	"strings"
	"jvmgo/ch05/classfile"
)
type Cmd struct {
	helpFlag bool
	versionFlag bool
	cpOption string
	XjreOption string
	class string
	args []string
}

func parseCmd() *Cmd{
	cmd :=&Cmd{}


	flag.Usage =printUsage
	flag.BoolVar(&cmd.helpFlag,"help",false,"print help message")
	flag.BoolVar(&cmd.helpFlag,"?",false,"print help message")
	flag.BoolVar(&cmd.versionFlag,"version",false,"print version and exit")
	flag.StringVar(&cmd.cpOption,"classpath","","classpath")
	flag.StringVar(&cmd.cpOption,"cp","","classpath")
	flag.StringVar(&cmd.XjreOption,"Xjre","","path to jre")
	flag.Parse()

	args :=flag.Args()
	if len(args)>0{
		cmd.class=args[0]
		cmd.args=args[1:]
	}
	return cmd
}

func printUsage(){
	fmt.Printf("Usage: %s [-option] class [args...]\n",os.Args[0])
	flag.PrintDefaults()

}
func startJVM(cmd *Cmd){
	cp := classpath.Parse(cmd.XjreOption,cmd.cpOption)
	fmt.Printf("classpath:%v class : %v args : %v \n",
		cp,cmd.class,cmd.args)
	className := strings.Replace(cmd.class,".","/",-1)
	cf := loadClass(className,cp)
	mainMethod :=getMainMethod(cf)
	if mainMethod !=nil{
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n",cmd.class)
	}

	fmt.Println(cmd.class)
	printClassInfo(cf)
	classData,_,err := cp.ReadClass(className)
	if err !=nil{
		fmt.Printf("Could not find or load main class %s \n",cmd.class)
		return
	}
	fmt.Printf("class data :%v \n",classData)

}
func loadClass(className string,cp *classpath.Classpath) *classfile.ClassFile{
	classData,_,err :=cp.ReadClass(className)
	if err !=nil{
		panic(err)
	}
	cf,err := classfile.Parse(classData)
	if err !=nil {
		panic(err)
	}
	return cf

}
func printClassInfo(cf *classfile.ClassFile)  {
	fmt.Printf("version: %v.%v\n",cf.MajorVersion(),cf.MinorVersion())
	fmt.Printf("constant count: %v\n",len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n",cf.AccessFlags())
	fmt.Printf("this class: %v\n",cf.ClassName())
	fmt.Printf("super class: %v\n",cf.SuperClassName())
	fmt.Printf("interfaces: %v\n",cf.InterfaceNames())
	fmt.Printf("fields count: %v\n",len(cf.Fields()))
	for _,m := range cf.Fields(){
		fmt.Printf("%s\n",m.Name())
	}
	fmt.Printf("methods count: %v\n",len(cf.Methods()))
	for _,m := range cf.Methods(){
		fmt.Printf("%s\n",m.Name())
	}
}
func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo{
	for _,m:=range cf.Methods(){
		if m.Name() =="main" && m.Descriptor() =="([LJava/lang/String;)V"{
			return m
		}
	}
	return nil
}

