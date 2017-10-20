package resources

type Disk interface {
	Create() bool
	Dumps() vm
}

type disk struct {
	ResourceBaseIface
	Type string
	Path string
	Uuid string
}

func NewDISK() ResourceBaseIface {
	return disk{Type:"KVM.DISK"}
}

func (thisDisk disk) Create() bool {
	return true
}

func (thisDisk disk) Dumps() interface{} {
	return thisDisk
}

func (thisDisk disk) GetType() string {
	return thisDisk.Type
}

func (thisDisk disk) GetId() string {
	return thisDisk.Uuid
}