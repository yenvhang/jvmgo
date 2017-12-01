package ch05

import (
	"jvmgo/ch05/classfile"
	"jvmgo/ch05/rtda"
	"fmt"
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/instructions/constants"
	"jvmgo/ch05/instructions"
)

func interpret(methodInfo *classfile.MemberInfo){
	codeAttr :=methodInfo.CodeAttribute()
	maxLocals :=codeAttr.MaxLocals()
	maxStack :=codeAttr.MaxStack()
	byteCode :=codeAttr.Code()

	thread :=rtda.NewThread()
	frame :=thread.NewFrame(maxLocals,maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread,byteCode)
}
func catchErr(frame *rtda.Frame){
	if r:=recover();r!=nil{
		fmt.Printf("LocalVars:%v\n","frame.LocalVars")
		fmt.Printf("OperandStack:%v\n",frame.OperandStack())
		panic(r)
	}
}
func loop(thread rtda.Thread,bytecode []byte){
	frame :=thread.PopFrame()
	reader :=&base.BytecodeReader{}
	for{
		pc := frame.NextPC()


		reader.Reset(bytecode,pc)
		opcode :=reader.ReadUint8()
		inst :=instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.setNextPC(reader.PC())
		fmt.Printf("pc:%2d inst :%T %V\n",pc,inst,inst)
		inst.Excute(frame)

	}
}
func newInstruction(opcode byte) base.Instruction{
	switch opcode {
	case 0x00: return &NOP{}
	case 0x01 :return &constants.ACONST_NULL{}
	default:
		panic(fmt.Errorf("Unsupported opcode : 0x%x !",opcode))
	}
}
