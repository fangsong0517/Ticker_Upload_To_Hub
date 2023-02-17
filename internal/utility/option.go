package utility

import "flag"

// Option program parameters.
type Option struct {
	ConfigFile string
}

// NewOption instantiate a option.
func NewOption() *Option {
	return &Option{}
}

// Parse parse program parameters.
func (o *Option) Parse() {
	flag.StringVar(&o.ConfigFile, "c", "../conf/config.yaml", "Specify the path for configuration")
	flag.Parse()
}
