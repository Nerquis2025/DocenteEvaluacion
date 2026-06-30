package docente

import (
	"testing"
)

func TestValidarEmailPrivado(t *testing.T) {
	d1 := &Docente{Email: "nerquis@uide.edu.ec"}
	if !d1.validarEmail() {
		t.Errorf("Error: Se esperaba que '%s' sea válido", d1.Email)
	}

	d2 := &Docente{Email: "nerquis@uide"}
	if d2.validarEmail() {
		t.Errorf("Error: Se esperaba que '%s' sea inválido por falta de punto", d2.Email)
	}
}

func TestNormalizarNombrePrivado(t *testing.T) {
	d := &Docente{Nombre: "nerquis carrera"}
	resultado := d.normalizarNombre()
	esperado := "Nerquis Carrera"

	if resultado != esperado {
		t.Errorf("Obtenido: '%s'; Se esperaba: '%s'", resultado, esperado)
	}
}

func TestAgregarEvaluacionInterna(t *testing.T) {
	d := NuevoDocente("1", "nerquis carrera", "nerquis@uide.edu.ec", "Sistemas", "Docente")
	d.agregarEvaluacionInterna("EVAL-2026-T1")

	if len(d.evaluaciones) != 1 {
		t.Fatalf("Error: La lista debería tener 1 evaluación, pero tiene %d", len(d.evaluaciones))
	}

	// Llamamos a GetEvaluaciones para completar la cobertura
	lista := d.GetEvaluaciones()
	if len(lista) != 1 || lista[0] != "EVAL-2026-T1" {
		t.Errorf("Error en GetEvaluaciones: la lista no coincide")
	}
}
