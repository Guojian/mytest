package resources


type ResourceBaseIface interface {
	GetType() string
	GetId() string
	Start() bool
	Define() bool
	Create() bool
	Stop() bool
	Dumps() interface{}
	Init(args... string) ResourceBaseIface
}

type ResourceBase struct {
	Type string
}