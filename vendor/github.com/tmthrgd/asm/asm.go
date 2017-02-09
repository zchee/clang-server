// Copyright 2014 Benoît Amiaux. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package asm

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var Seperator = " " // or \t

type Opcodes struct {
	a *Asm
}

type Asm struct {
	Opcodes

	w      *bufio.Writer
	errors []string

	// per function
	name  string
	args  int
	stack int
	split bool
}

func NewAsm(w io.Writer) *Asm {
	a := &Asm{
		w: bufio.NewWriter(w),
	}
	a.Opcodes.a = a

	a.write("\n#include \"textflag.h\"")
	return a
}

func (a *Asm) NewFunction(name string) {
	a.name = name
	a.args = 0
	a.stack = 0
	a.split = true
}

func (a *Asm) NoSplit() {
	a.split = false
}

func isZeroSlice(s []byte) bool {
	for _, b := range s {
		if b != 0 {
			return false
		}
	}

	return true
}

type Data string

func (d Data) String() string {
	return fmt.Sprintf("%v(SB)", string(d))
}

func (Data) Gas() string {
	panic("referencing GLOBL directives in unsupported opcodes is forbidden")
}

func (d Data) Offset(i int) Data {
	if i == 0 {
		return d
	}

	return Data(fmt.Sprintf("%v+0x%02x", string(d), i))
}

func (d Data) Address() Data {
	return Data(fmt.Sprintf("$%v", string(d)))
}

func (a *Asm) Data(name string, data []byte) Data {
	name = fmt.Sprintf("%v<>", name)

	a.write("")

	i := 0
	for ; i < len(data); i += 8 {
		if isZeroSlice(data[i : i+8]) {
			continue
		}

		a.write(fmt.Sprintf("DATA%v%v+0x%02x(SB)/8, $0x%016x", Seperator, name, i, data[i:i+8]))
	}

	for ; i < len(data); i++ {
		if data[i] == 0 {
			continue
		}

		a.write(fmt.Sprintf("DATA%v%v+0x%02x(SB)/1, $0x%02x", Seperator, name, i, data[i]))
	}

	a.write(fmt.Sprintf("GLOBL%v%v(SB),RODATA,$%v", Seperator, name, len(data)))
	return Data(name)
}

func (a *Asm) Data16(name string, data []uint16) Data {
	name = fmt.Sprintf("%v<>", name)

	a.write("")

	for i := 0; i < len(data); i++ {
		if data[i] == 0 {
			continue
		}

		a.write(fmt.Sprintf("DATA%v%v+0x%02x(SB)/2, $0x%04x", Seperator, name, 2*i, data[i]))
	}

	a.write(fmt.Sprintf("GLOBL%v%v(SB),RODATA,$%v", Seperator, name, 2*len(data)))
	return Data(name)
}

func (a *Asm) Data32(name string, data []uint32) Data {
	name = fmt.Sprintf("%v<>", name)

	a.write("")

	for i := 0; i < len(data); i++ {
		if data[i] == 0 {
			continue
		}

		a.write(fmt.Sprintf("DATA%v%v+0x%02x(SB)/4, $0x%08x", Seperator, name, 4*i, data[i]))
	}

	a.write(fmt.Sprintf("GLOBL%v%v(SB),RODATA,$%v", Seperator, name, 4*len(data)))
	return Data(name)
}

func (a *Asm) Data64(name string, data []uint64) Data {
	name = fmt.Sprintf("%v<>", name)

	a.write("")

	for i := 0; i < len(data); i++ {
		if data[i] == 0 {
			continue
		}

		a.write(fmt.Sprintf("DATA%v%v+0x%02x(SB)/8, $0x%016x", Seperator, name, 8*i, data[i]))
	}

	a.write(fmt.Sprintf("GLOBL%v%v(SB),RODATA,$%v", Seperator, name, 8*len(data)))
	return Data(name)
}

func (a *Asm) DataString(name string, data string) Data {
	name = fmt.Sprintf("%v<>", name)

	a.write("")

	i := 0
	for ; i < len(data); i += 8 {
		if isZeroSlice([]byte(data[i : i+8])) {
			continue
		}

		a.write(fmt.Sprintf("DATA%v%v+0x%02x(SB)/8, $%q", Seperator, name, i, data[i:i+8]))
	}

	for ; i < len(data); i++ {
		if data[i] == 0 {
			continue
		}

		a.write(fmt.Sprintf("DATA%v%v+0x%02x(SB)/1, $%q", Seperator, name, i, data[i]))
	}

	a.write(fmt.Sprintf("GLOBL%v%v(SB),RODATA,$%v", Seperator, name, len(data)))
	return Data(name)
}

type invalid struct{}

func (invalid) String() string {
	panic("invalid operand")
}

func (invalid) Gas() string {
	panic("invalid operand")
}

var Invalid Operand = invalid{}

type Argument struct {
	name   string
	offset int
}

func (s *Argument) String() string {
	return fmt.Sprintf("%v+%v(FP)", s.name, s.offset)
}

func (*Argument) Gas() string {
	panic("referencing arguments in unsupported opcodes is forbidden")
}

func (a *Asm) Argument(name string, size int) Operand {
	a.args += size
	return &Argument{
		name:   name,
		offset: a.args - size,
	}
}

func (a *Asm) SliceArgument(name string) []Operand {
	var rpy []Operand

	for i := 0; i < 3; i++ {
		rpy = append(rpy, a.Argument(name, 8))
	}

	return rpy
}

type StackOperand struct {
	name   string
	offset int
}

func (s *StackOperand) String() string {
	return fmt.Sprintf("%v+-%v(SP)", s.name, s.offset)
}

func (*StackOperand) Gas() string {
	panic("referencing stack variables in unsupported opcodes is forbidden")
}

func (a *Asm) PushStack(name string, size int) Operand {
	a.stack += size
	return &StackOperand{
		name:   name,
		offset: a.stack,
	}
}

func (a *Asm) Start() {
	if a.split {
		a.write(fmt.Sprintf("\nTEXT ·%v(SB),0,$%v-%v", a.name, a.stack, a.args))
	} else {
		a.write(fmt.Sprintf("\nTEXT ·%v(SB),NOSPLIT,$%v", a.name, a.stack))
	}
}

func (a *Asm) Flush() error {
	err := a.w.Flush()
	a.save(err)
	return a.getErrors()
}

func (a *Asm) save(err error) {
	if err == nil {
		return
	}

	a.errors = append(a.errors, err.Error())
}

func (a *Asm) getErrors() error {
	if len(a.errors) == 0 {
		return nil
	}

	return fmt.Errorf("%s", strings.Join(a.errors, "\n"))
}

func (a *Asm) write(msg string) {
	_, err := a.w.WriteString(msg + "\n")
	a.save(err)
}

type Operand interface {
	String() string
	Gas() string
}

type constant string

func (cons constant) String() string {
	return string(cons)
}

func (cons constant) Gas() string {
	return string(cons)
}

func Constant(value interface{}) Operand {
	return constant(fmt.Sprintf("$%v", value))
}

type Register struct {
	literal string
	gas     string
}

func (r Register) String() string {
	return r.literal
}

func (r Register) Gas() string {
	return r.gas
}

type Scale uint

const (
	SX0 Scale = 0
	SX1 Scale = 1 << (iota - 1)
	SX2
	SX4
	SX8
)

type addressOperand struct {
	lit string
	gas string
}

func (a addressOperand) String() string {
	return a.lit
}

func (a addressOperand) Gas() string {
	return a.gas
}

func address(base Register) Operand {
	return addressOperand{
		fmt.Sprintf("(%v)", base.String()),
		fmt.Sprintf("(%v)", base.Gas()),
	}
}

func displaceaddress(base Register, index int) Operand {
	if index == 0 {
		return address(base)
	}

	return addressOperand{
		fmt.Sprintf("%v(%v)", index, base.String()),
		fmt.Sprintf("%v(%v)", index, base.Gas()),
	}
}

func scaledindex(index Register, scale Scale) string {
	if scale == SX0 {
		return ""
	}

	return fmt.Sprintf("(%v*%v)", index.String(), scale)
}

func indexaddress(base Register, index Register, scale Scale) Operand {
	return addressOperand{
		fmt.Sprintf("(%v)%v", base.String(), scaledindex(index, scale)),
		fmt.Sprintf("(%v, %v, %v)", base.Gas(), index.Gas(), scale),
	}
}

func fulladdress(base Register, index Register, scale Scale, displacement int) Operand {
	d := ""

	if displacement != 0 {
		d = fmt.Sprintf("%v", displacement)
	}

	return addressOperand{
		fmt.Sprintf("%v(%v)%v", d, base.String(), scaledindex(index, scale)),
		fmt.Sprintf("%v(%v, %v, %v)", d, base.Gas(), index.Gas(), scale),
	}
}

func Address(base Register, offsets ...interface{}) Operand {
	// happily panics if not given expected input
	switch len(offsets) {
	case 0:
		return address(base)
	case 1:
		switch t := offsets[0].(type) {
		case int:
			return displaceaddress(base, t)
		case uint:
			return displaceaddress(base, int(t))
		case Register:
			return indexaddress(base, t, SX1)
		case Scale:
			return addressOperand{
				scaledindex(base, t),
				fmt.Sprintf("(, %v, %v)", base.String(), t),
			}
		}
	case 2:
		index, ok := offsets[0].(Register)
		if !ok {
			break
		}

		switch t := offsets[1].(type) {
		case int:
			return fulladdress(base, index, SX1, t)
		case uint:
			return fulladdress(base, index, SX1, int(t))
		case Scale:
			return indexaddress(base, index, t)
		}
	case 3:
		index, ok := offsets[0].(Register)
		if !ok {
			break
		}

		scale, ok := offsets[1].(Scale)
		if !ok {
			break
		}

		switch t := offsets[2].(type) {
		case int:
			return fulladdress(base, index, scale, t)
		case uint:
			return fulladdress(base, index, scale, int(t))
		}
	}

	panic("unexpected input")
}

type SimdRegister struct {
	literal string
	gas     string
}

func (r SimdRegister) String() string {
	return r.literal
}

func (r SimdRegister) Gas() string {
	return r.gas
}

var (
	SP  = Register{literal: "SP", gas: "%rsp"}
	AX  = Register{literal: "AX", gas: "%rax"}
	AH  = Register{literal: "AH", gas: "%ah"}
	AL  = Register{literal: "AL", gas: "%al"}
	BX  = Register{literal: "BX", gas: "%rbx"}
	BH  = Register{literal: "BH", gas: "%bh"}
	BL  = Register{literal: "BL", gas: "%bl"}
	CX  = Register{literal: "CX", gas: "%rcx"}
	CH  = Register{literal: "CH", gas: "%ch"}
	CL  = Register{literal: "CL", gas: "%cl"}
	DX  = Register{literal: "DX", gas: "%rdx"}
	DH  = Register{literal: "DH", gas: "%dh"}
	DL  = Register{literal: "DL", gas: "%dl"}
	BP  = Register{literal: "BP", gas: "%rbp"}
	DI  = Register{literal: "DI", gas: "%rdi"}
	SI  = Register{literal: "SI", gas: "%rsi"}
	R8  = Register{literal: "R8", gas: "%r8"}
	R9  = Register{literal: "R9", gas: "%r9"}
	R10 = Register{literal: "R10", gas: "%r10"}
	R11 = Register{literal: "R11", gas: "%r11"}
	R12 = Register{literal: "R12", gas: "%r12"}
	R13 = Register{literal: "R13", gas: "%r13"}
	R14 = Register{literal: "R14", gas: "%r14"}
	R15 = Register{literal: "R15", gas: "%r15"}

	X0  = SimdRegister{literal: "X0", gas: "%xmm0"}
	X1  = SimdRegister{literal: "X1", gas: "%xmm1"}
	X2  = SimdRegister{literal: "X2", gas: "%xmm2"}
	X3  = SimdRegister{literal: "X3", gas: "%xmm3"}
	X4  = SimdRegister{literal: "X4", gas: "%xmm4"}
	X5  = SimdRegister{literal: "X5", gas: "%xmm5"}
	X6  = SimdRegister{literal: "X6", gas: "%xmm6"}
	X7  = SimdRegister{literal: "X7", gas: "%xmm7"}
	X8  = SimdRegister{literal: "X8", gas: "%xmm8"}
	X9  = SimdRegister{literal: "X9", gas: "%xmm9"}
	X10 = SimdRegister{literal: "X10", gas: "%xmm10"}
	X11 = SimdRegister{literal: "X11", gas: "%xmm11"}
	X12 = SimdRegister{literal: "X12", gas: "%xmm12"}
	X13 = SimdRegister{literal: "X13", gas: "%xmm13"}
	X14 = SimdRegister{literal: "X14", gas: "%xmm14"}
	X15 = SimdRegister{literal: "X15", gas: "%xmm15"}

	Y0  = SimdRegister{literal: "Y0", gas: "%ymm0"}
	Y1  = SimdRegister{literal: "Y1", gas: "%ymm1"}
	Y2  = SimdRegister{literal: "Y2", gas: "%ymm2"}
	Y3  = SimdRegister{literal: "Y3", gas: "%ymm3"}
	Y4  = SimdRegister{literal: "Y4", gas: "%ymm4"}
	Y5  = SimdRegister{literal: "Y5", gas: "%ymm5"}
	Y6  = SimdRegister{literal: "Y6", gas: "%ymm6"}
	Y7  = SimdRegister{literal: "Y7", gas: "%ymm7"}
	Y8  = SimdRegister{literal: "Y8", gas: "%ymm8"}
	Y9  = SimdRegister{literal: "Y9", gas: "%ymm9"}
	Y10 = SimdRegister{literal: "Y10", gas: "%ymm10"}
	Y11 = SimdRegister{literal: "Y11", gas: "%ymm11"}
	Y12 = SimdRegister{literal: "Y12", gas: "%ymm12"}
	Y13 = SimdRegister{literal: "Y13", gas: "%ymm13"}
	Y14 = SimdRegister{literal: "Y14", gas: "%ymm14"}
	Y15 = SimdRegister{literal: "Y15", gas: "%ymm15"}
)

type Label struct{ name string }

func (a *Asm) NewLabel(name string) Label {
	return Label{name}
}

func (l Label) String() string {
	return l.name
}

func (Label) Gas() string {
	panic("referencing labels in unsupported opcodes is forbidden")
}

func (l Label) Suffix(suffix string) Label {
	return Label{fmt.Sprintf("%s_%s", l.name, suffix)}
}

func (a *Asm) op(instruction string, ops ...Operand) {
	if len(ops) == 0 {
		a.write("\t" + instruction)
		return
	}

	var sOps []string

	for i := len(ops) - 1; i >= 0; i-- {
		sOps = append(sOps, ops[i].String())
	}

	a.write(fmt.Sprintf("\t%v%v%v", instruction, Seperator, strings.Join(sOps, ", ")))
}

var objdumpRegex = regexp.MustCompile(`^\s+\d:\s+((?:[0-9a-fA-F]{2} )+)`)

func (a *Asm) unsupOp(instruction string, ops ...Operand) {
	tmp, err := ioutil.TempFile("", "")
	if err != nil {
		panic(err)
	}

	defer os.Remove(tmp.Name())
	tmp.Close()

	var gOps []string

	for i := len(ops) - 1; i >= 0; i-- {
		gOps = append(gOps, ops[i].Gas())
	}

	cmd := exec.Command("as", "-o", tmp.Name(), "-")
	cmd.Stdin = strings.NewReader(fmt.Sprintf("%v\t%s\n", instruction, strings.Join(gOps, ", ")))
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	cmd = exec.Command("objdump", "-d", tmp.Name())
	cmd.Stderr = os.Stderr

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	if err = cmd.Start(); err != nil {
		panic(err)
	}

	gOps = gOps[:0]

	for i := len(ops) - 1; i >= 0; i-- {
		gOps = append(gOps, ops[i].String())
	}

	a.write(fmt.Sprintf("\t// %v%v%s", instruction, Seperator, strings.Join(gOps, ", ")))

	scan2 := bufio.NewScanner(stdout)

	for scan2.Scan() {
		m := objdumpRegex.FindStringSubmatch(scan2.Text())
		if m == nil {
			continue
		}

		a.write(fmt.Sprintf("\tBYTE $0x%s", strings.Join(strings.Split(strings.TrimSpace(m[1]), " "), "; BYTE $0x")))
	}

	if err = cmd.Wait(); err != nil {
		panic(err)
	}
}

func (a *Asm) Label(name Label) {
	a.write(name.String() + ":")
}

func Do(file, header string, run func(*Asm)) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := io.WriteString(f, header); err != nil {
		return err
	}

	a := NewAsm(f)
	run(a)
	return a.Flush()
}

//go:generate go run ./opcode_gen.go -i $GOROOT/src/cmd/internal/obj/x86/a.out.go -o opcode.go -p asm
//go:generate gofmt -w opcode.go

// general opcodes

func (o Opcodes) Nop(ops ...Operand)  { o.a.op("NOP", ops...) }
func (o Opcodes) NOP(ops ...Operand)  { o.a.op("NOP", ops...) }
func (o Opcodes) Ret(ops ...Operand)  { o.a.op("RET", ops...) }
func (o Opcodes) RET(ops ...Operand)  { o.a.op("RET", ops...) }
func (o Opcodes) Call(ops ...Operand) { o.a.op("CALL", ops...) }
func (o Opcodes) CALL(ops ...Operand) { o.a.op("CALL", ops...) }
func (o Opcodes) Jmp(ops ...Operand)  { o.a.op("JMP", ops...) }
func (o Opcodes) JMP(ops ...Operand)  { o.a.op("JMP", ops...) }

// other jumps

func (o Opcodes) Je(ops ...Operand)  { o.a.op("JE", ops...) }
func (o Opcodes) JE(ops ...Operand)  { o.a.op("JE", ops...) }
func (o Opcodes) Jb(ops ...Operand)  { o.a.op("JB", ops...) }
func (o Opcodes) JB(ops ...Operand)  { o.a.op("JB", ops...) }
func (o Opcodes) Jae(ops ...Operand) { o.a.op("JAE", ops...) }
func (o Opcodes) JAE(ops ...Operand) { o.a.op("JAE", ops...) }
func (o Opcodes) Jz(ops ...Operand)  { o.a.op("JZ", ops...) }
func (o Opcodes) JZ(ops ...Operand)  { o.a.op("JZ", ops...) }
func (o Opcodes) Jnz(ops ...Operand) { o.a.op("JNZ", ops...) }
func (o Opcodes) JNZ(ops ...Operand) { o.a.op("JNZ", ops...) }
func (o Opcodes) Jc(ops ...Operand)  { o.a.op("JC", ops...) }
func (o Opcodes) JC(ops ...Operand)  { o.a.op("JC", ops...) }

// faulty/incomplete opcodes

func (o Opcodes) Pextrw(ops ...Operand) { o.a.unsupOp("PEXTRW", ops...) }
func (o Opcodes) PEXTRW(ops ...Operand) { o.a.unsupOp("PEXTRW", ops...) }

// unsupported opcodes

func (o Opcodes) Ptest(ops ...Operand)        { o.a.unsupOp("PTEST", ops...) }
func (o Opcodes) PTEST(ops ...Operand)        { o.a.unsupOp("PTEST", ops...) }
func (o Opcodes) Vpunpckhbw(ops ...Operand)   { o.a.unsupOp("VPUNPCKHBW", ops...) }
func (o Opcodes) VPUNPCKHBW(ops ...Operand)   { o.a.unsupOp("VPUNPCKHBW", ops...) }
func (o Opcodes) Vpunpcklqdq(ops ...Operand)  { o.a.unsupOp("VPUNPCKLQDQ", ops...) }
func (o Opcodes) VPUNPCKLQDQ(ops ...Operand)  { o.a.unsupOp("VPUNPCKLQDQ", ops...) }
func (o Opcodes) Vpunpckhqdq(ops ...Operand)  { o.a.unsupOp("VPUNPCKHQDQ", ops...) }
func (o Opcodes) VPUNPCKHQDQ(ops ...Operand)  { o.a.unsupOp("VPUNPCKHQDQ", ops...) }
func (o Opcodes) Vpshufb(ops ...Operand)      { o.a.unsupOp("VPSHUFB", ops...) }
func (o Opcodes) VPSHUFB(ops ...Operand)      { o.a.unsupOp("VPSHUFB", ops...) }
func (o Opcodes) Vpor(ops ...Operand)         { o.a.unsupOp("VPOR", ops...) }
func (o Opcodes) VPOR(ops ...Operand)         { o.a.unsupOp("VPOR", ops...) }
func (o Opcodes) Vpcmpgtb(ops ...Operand)     { o.a.unsupOp("VPCMPGTB", ops...) }
func (o Opcodes) VPCMPGTB(ops ...Operand)     { o.a.unsupOp("VPCMPGTB", ops...) }
func (o Opcodes) Vpcmpeqq(ops ...Operand)     { o.a.unsupOp("VPCMPEQQ", ops...) }
func (o Opcodes) VPCMPEQQ(ops ...Operand)     { o.a.unsupOp("VPCMPEQQ", ops...) }
func (o Opcodes) Pblendvb(ops ...Operand)     { o.a.unsupOp("PBLENDVB", ops...) }
func (o Opcodes) PBLENDVB(ops ...Operand)     { o.a.unsupOp("PBLENDVB", ops...) }
func (o Opcodes) Vinsertf128(ops ...Operand)  { o.a.unsupOp("VINSERTF128", ops...) }
func (o Opcodes) VINSERTF128(ops ...Operand)  { o.a.unsupOp("VINSERTF128", ops...) }
func (o Opcodes) Vpblendvb(ops ...Operand)    { o.a.unsupOp("VPBLENDVB", ops...) }
func (o Opcodes) VPBLENDVB(ops ...Operand)    { o.a.unsupOp("VPBLENDVB", ops...) }
func (o Opcodes) Vpsrldq(ops ...Operand)      { o.a.unsupOp("VPSRLDQ", ops...) }
func (o Opcodes) VPSRLDQ(ops ...Operand)      { o.a.unsupOp("VPSRLDQ", ops...) }
func (o Opcodes) Vpsrad(ops ...Operand)       { o.a.unsupOp("VPSRAD", ops...) }
func (o Opcodes) VPSRAD(ops ...Operand)       { o.a.unsupOp("VPSRAD", ops...) }
func (o Opcodes) Vpsrld(ops ...Operand)       { o.a.unsupOp("VPSRLD", ops...) }
func (o Opcodes) VPSRLD(ops ...Operand)       { o.a.unsupOp("VPSRLD", ops...) }
func (o Opcodes) Vpslld(ops ...Operand)       { o.a.unsupOp("VPSLLD", ops...) }
func (o Opcodes) VPSLLD(ops ...Operand)       { o.a.unsupOp("VPSLLD", ops...) }
func (o Opcodes) Pmaddubsw(ops ...Operand)    { o.a.unsupOp("PMADDUBSW", ops...) }
func (o Opcodes) PMADDUBSW(ops ...Operand)    { o.a.unsupOp("PMADDUBSW", ops...) }
func (o Opcodes) Vpsubb(ops ...Operand)       { o.a.unsupOp("VPSUBB", ops...) }
func (o Opcodes) VPSUBB(ops ...Operand)       { o.a.unsupOp("VPSUBB", ops...) }
func (o Opcodes) Vbroadcastss(ops ...Operand) { o.a.unsupOp("VBROADCASTSS", ops...) }
func (o Opcodes) VBROADCASTSS(ops ...Operand) { o.a.unsupOp("VBROADCASTSS", ops...) }
func (o Opcodes) Pcmpeqq(ops ...Operand)      { o.a.unsupOp("PCMPEQQ", ops...) }
func (o Opcodes) PCMPEQQ(ops ...Operand)      { o.a.unsupOp("PCMPEQQ", ops...) }
