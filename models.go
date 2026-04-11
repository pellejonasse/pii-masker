package piimasker

type PiiMode string

const (
	PiiModeNone      PiiMode = ""      // no tag — inherit parent mode
	PiiModeShow      PiiMode = "show"
	PiiModeMask      PiiMode = "mask"
	PiiModeAnonymize PiiMode = "anonymize"
)

type MaskerConfig struct {
	maxPiiStringLength int `default:"100"`
}

type Number interface {
	int64
	float64
	uint64
}
