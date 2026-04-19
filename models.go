package piimasker

type piiMode string

const (
	piiModeNone      piiMode = "" // no tag — inherit parent mode, which is nice for external objects
	piiModeShow      piiMode = "show"
	piiModeMask      piiMode = "mask"
	piiModeAnonymize piiMode = "anonymize"
)

type maskerConfig struct {
	MaxPiiStringLength int    // defaults to 100
	TagField           string // struct tag key to read, defaults to "Pii"
}

// Option is a functional option for configuring a PiiMasker.
type Option func(*maskerConfig)

// WithMaxPiiStringLength sets the maximum length of a masked string.
func WithMaxPiiStringLength(n int) Option {
	return func(c *maskerConfig) {
		c.MaxPiiStringLength = n
	}
}

// WithTagField sets the struct tag key used to identify PII fields (default "Pii").
func WithTagField(tag string) Option {
	return func(c *maskerConfig) {
		c.TagField = tag
	}
}

// don't think I want to use this
type number interface {
	int64 | float64 | uint64
}
