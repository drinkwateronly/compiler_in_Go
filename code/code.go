package code

import "fmt"

type Instructions []byte

type Opcode byte

const (
	OpConstant Opcode = iota
)

type Definition struct {
	Name          string // 操作码可读性
	OperandWidths []int  // 操作码对应操作数占用的字节数
}

var definitions = map[Opcode]*Definition{ // 键 Opcode 值 *Definition
	OpConstant: {"OpConstant", []int{2}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}
	return def, nil
}
