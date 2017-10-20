package resources

type InitFunc func() ResourceBaseIface

const (
	A string = "VM"
	B string = "DISK"
)

var initFuncs = map[string]InitFunc{
	A: NewVM,
	B: NewDISK,
}