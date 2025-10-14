package release

type CreateOptions interface {
	Apply(opts *createOptions)
}

type createOptions struct {
	Notes  string
	Assets []string
}

func (m *manager) Create(name string, tag string, opts ...CreateOptions) error {
	createOpts := &createOptions{}
	for _, opt := range opts {
		opt.Apply(createOpts)
	}

	return nil
}
