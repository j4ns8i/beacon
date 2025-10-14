package release

type WithNotesOption struct {
	Notes string
}

func (o *WithNotesOption) Apply(opts *createOptions) {
	opts.Notes = o.Notes
}

type WithAssetsOption struct {
	Assets []string
}

func (o *WithAssetsOption) Apply(opts *createOptions) {
	opts.Assets = o.Assets
}
