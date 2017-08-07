package log

type Options struct {
	ModuleName     string
	LogMode        string
	LogLevel       string
	LogPrefix      string
	Verbose        bool
	LogPath        string
	LogMaxSize     int64
	LogRotateDaily bool
	LogMaxAge      int
}

func NewOptions() *Options {
	return &Options{
		LogMode:        "normal",
		LogLevel:       "info",
		Verbose:        false,
		LogMaxSize:     100,
		LogRotateDaily: false,
		LogMaxAge:      30,
	}
}
