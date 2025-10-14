package release

type Manager interface {
	Create(name string, tag string, opts ...CreateOptions) error
}

func NewManager() Manager {
	return &manager{}
}

type manager struct {
}
