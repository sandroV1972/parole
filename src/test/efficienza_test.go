//go:build perf
// +build perf

package main

import "testing"

func TestEfficienzaInsert(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"mega-in",
		"mega-out",
		verbose,
	)

}

func TestEfficienzaCatena(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"megacatena-in",
		"megacatena-out",
		verbose,
	)

}

func TestEfficienzaGruppo(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"megagruppo-in",
		"megagruppo-out",
		verbose,
	)

}
