package log

// StdLogger ...
type StdLogger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})
}

// Logger ...
type Logger interface {
	Error() StdLogger
	Warn() StdLogger
	Info() StdLogger
	Debug() StdLogger

	SetLevel(Level)
	GetLevel() Level
}
