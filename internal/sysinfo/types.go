package sysinfo

type Hardware struct {
	*Cpu
	Ram string // formatted ram int into string
}

type Cpu struct {
	Model string
	MHZ   string
}
