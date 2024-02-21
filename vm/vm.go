package vm

import (
  "monkey/code"
  "monkey/compiler"
  "monkey/object"
)

const StackSize = 2048

type VM struct {
  constants     []object.Object
  instructions  code.Instructions

  stack []object.Object
  sp    int
}

func New(bytecode *compiler.Bytecode) *VM {
  return &VM{
    instructions: bytecode.Instructions,
    constants:    bytecode.Constants,

    stack: make([]object.Object, StackSize),
    sp:    0,
  }
}
