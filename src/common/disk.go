package common


type disk struct {
	Type string
	Path string
	Size string
}

type Disk interface{}


func NewDisk() Disk {
	return disk{Type:"KVM.DISK"}
}