package prngo

type Msws struct {
	seed	uint64
	state	uint64
	seq		uint64
}

func (p *Msws) Seed(s uint64) {
	p.seed = s
}

func (p *Msws) Rand() (i uint32){
	p.state *= p.state
	p.seq += p.seed
	p.state += p.seq

	i = uint32(p.state >> 32) | uint32(p.state << 32)
	p.state = uint64(i)
	return i
}

type Lcg struct {
	state	uint64
	m, a, c	uint64
}

func (l *Lcg) Seed(s uint64) {
	l.state = s
	l.c = 1013904223
	l.a = 1664525
	l.m = (1 << 64) - 1
}

func (l *Lcg) SetParams(m uint64, a uint64, c uint64) {
	l.m = m
	l.a = a
	l.c = c
}

func (l *Lcg) Rand() uint64 {
	l.state = (l.a * l.state + l.c) & l.m
	return l.state
}

