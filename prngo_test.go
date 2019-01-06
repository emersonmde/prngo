package prngo

import (
	"testing"
)

func TestMsws_Rand(t *testing.T) {
	a := new(Msws)

	a.Seed(0xb5ad4eceda1ce2a9)
	a.Rand()
}

func BenchmarkMsws_Rand(b *testing.B) {
	a := new(Msws)
	a.Seed(0xb5ad4eceda1ce2a9)

	for i := 0; i < b.N; i++ {
		a.Rand()
	}
}

func TestLcg_Rand(t *testing.T) {
	a := new(Lcg)

	a.Seed(238483277383)
	a.Rand()
}

func BenchmarkLcg_Rand(b *testing.B) {
	a := new(Lcg)
	a.Seed(238483277383)

	for i := 0; i < b.N; i++ {
		a.Rand()
	}
}