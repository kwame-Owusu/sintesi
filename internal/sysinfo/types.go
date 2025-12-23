package sysinfo

type Hardware struct {
	*Cpu
	*Ram
}

type Cpu struct {
	Model string
	MHZ   string
}

type Ram struct {
	Total     string
	Available string
}
