package larchive

type extractor struct {
	DryRun bool

	// Allow symlink extraction
	AllowSymlink bool
}

type ExtractOption func(*extractor)

func WithDryRun() ExtractOption {
	return func(e *extractor) {
		e.DryRun = true
	}
}

func WithAllowSymlink() ExtractOption {
	return func(e *extractor) {
		e.AllowSymlink = true
	}
}
