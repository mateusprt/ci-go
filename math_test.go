package main

import "testing"

func TestSoma(t *testing.T) {
	valor_esperado := 30
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado %d", total, valor_esperado)
	}
}
