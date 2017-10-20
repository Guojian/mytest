package resources
import "fmt"

type VM interface {
	Start() bool
	Define() bool
	Stop() bool
	Dumps() vm
}

type vm struct {
	ResourceBaseIface
	Type string
	Uuid string

	// arguments
	CPU string
	Memory string
}

func NewVM() ResourceBaseIface {
	return vm{Type:"KVM.VM"}
}

func (myvm vm) Create() bool {
	return true
}

func (myvm vm) Start() bool {
	return true
}

func (myvm vm) Stop() bool {
	return true
}

func (myvm vm) Dumps() interface{} {
	return myvm
}

func (myvm vm) GetType() string {
	return myvm.Type
}

func (myvm vm) GetId() string {
	return myvm.Uuid
}

func (myvm vm) Init(args... string) ResourceBaseIface {
	fmt.Println(args[0])
	
	myvm.CPU = args[0]
	myvm.Memory = args[1]
	return myvm
}