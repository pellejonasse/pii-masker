package piimasker

type piiMode string

const (
	piiModeNone      piiMode = "" // no tag — inherit parent mode, which is nice for external objects
	piiModeShow      piiMode = "show"
	piiModeMask      piiMode = "mask"
	piiModeAnonymize piiMode = "anonymize"
)

type MaskerConfig struct {
	MaxPiiStringLength int // defaults to 100
}

// don't think I want to use this
type number interface {
	int64 | float64 | uint64
}
