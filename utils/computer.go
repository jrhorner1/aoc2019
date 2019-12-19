package utils

type Computer struct {
    memory          *[]int
    pointer         int
    input           chan int
    output          chan int
    instruction     int
    relativeBase    int
}

const (
    memoryMultiplier = 50
)

func NewComputer(program *[]int) Computer {
    memory := make([]int, len(*program) * memoryMultiplier)
    copy(memory, *program)

    return Computer{
        memory: &memory, 
        pointer: 0, 
        input: make(chan int, 1), 
        output: make(chan int), 
        instruction: 0, 
        relativeBase: 0}
}

func (c *Computer) Run() {
loop:
    for {
        c.instruction = c.Read(1)
        opcode := c.instruction % 100
        switch opcode {
        case 1: c.Add()
        case 2: c.Multiply()
        case 3: c.PutInput()
        case 4: c.PutOutput()
        case 5: c.JumpIfTrue()
        case 6: c.JumpIfFalse()
        case 7: c.LessThan()
        case 8: c.Equals()
        case 9: c.RelativeBaseOffset()
        case 99: break loop
        }
    }
    close(c.output)
    return
}

func (c *Computer) Parameter(mode int) int {
    var address int
    switch mode {
    case 0:
        address = (*c.memory)[c.pointer]
    case 1:
        address = c.pointer
    case 2:
        address = (*c.memory)[c.pointer] + c.relativeBase
    }
    c.pointer++
    return address
}

func (c *Computer) Read(mode int) int {
    address := c.Parameter(mode) 
    return (*c.memory)[address]
}

func (c *Computer) Write(value int, mode int) {
    address := c.Parameter(mode)
    (*c.memory)[address] = value
    return
}

func (c *Computer) Add() {
    a := c.Read((c.instruction / 100) % 10)
    b := c.Read((c.instruction / 1000) % 10)

    c.Write(a + b, (c.instruction / 10000) % 10)
}

func (c *Computer) Multiply() {
    a := c.Read((c.instruction / 100) % 10)
    b := c.Read((c.instruction / 1000) % 10)

    c.Write(a * b, (c.instruction / 10000) % 10)
}

func (c *Computer) PutInput() {
    in := <- c.input
    c.Write(in, (c.instruction / 100) % 10)
}

func (c *Computer) GetInput() chan int {
    return c.input
}

func (c *Computer) PutOutput() {
    out := c.Read((c.instruction / 100) % 10)
    c.output <- out
}

func (c *Computer) GetOutput() chan int {
    return c.output
}

func (c *Computer) JumpIfTrue() {
    test := c.Read((c.instruction / 100) % 10)
    newPointer := c.Read((c.instruction / 1000) % 10)
    if test != 0 {
        c.pointer = newPointer
    }
}

func (c *Computer) JumpIfFalse() {
    test := c.Read((c.instruction / 100) % 10)
    newPointer := c.Read((c.instruction / 1000) % 10)
    if test == 0 {
        c.pointer = newPointer
    }
}

func (c *Computer) LessThan() {
    a := c.Read((c.instruction / 100) % 10)
    b := c.Read((c.instruction / 1000) % 10)
    if a < b {
        c.Write(1, (c.instruction / 10000) % 10)
    } else {
        c.Write(0, (c.instruction / 10000) % 10)
    }
}

func (c *Computer) Equals() {
    a := c.Read((c.instruction / 100) % 10)
    b := c.Read((c.instruction / 1000) % 10)
    if a == b {
        c.Write(1, (c.instruction / 10000) % 10)
    } else {
        c.Write(0, (c.instruction / 10000) % 10)
    }
}

func (c *Computer) RelativeBaseOffset() {
    offset := c.Read((c.instruction / 100) % 10)
    c.relativeBase += offset
}



