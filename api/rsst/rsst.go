package rsst

// Info ...
type Info struct {
	ID   uint16
	Data []byte
	Ok   bool // handled in Process
}

// Handler ..
type Handler func(*Info)

// Registrator ...
type Registrator interface {
	AddHandler(ID uint16, f Handler)
}

// Packer ...
type Packer interface {
	Pack([]Info) []byte
	Unpack([]byte) []Info
}

// Processor ...
type Processor interface {
	Process([]Info)
}
