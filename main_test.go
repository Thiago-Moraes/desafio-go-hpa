package main

import "testing"

func TestMain(t *testing.T) {

	x := executando(16.0)

	if x != 4 {
		t.Error("Erro!")
	}
}
