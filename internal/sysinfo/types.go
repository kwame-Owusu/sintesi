package sysinfo

type Hardware struct {
	*Cpu
	*Ram
}

type Cpu struct {
	Model string
	MHZ   float64
}

type Ram struct {
	Total     uint64 // kB
	Available uint64 // kB
}
