package piimasker

type Config struct {
	maxPiiStringLength int  `default:"100"`
	scrambleNumbers    bool `default:"false"`
}
