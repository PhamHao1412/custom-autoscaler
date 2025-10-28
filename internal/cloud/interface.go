package cloud

type Provider interface {
	AddNode(name string) error
	RemoveNode(name string) error
	ListNodes() ([]string, error)
}
