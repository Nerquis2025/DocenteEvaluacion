package docente_test

import (
	"DocenteEvaluacion/docente"
	"testing"
)

func TestAccesoPublicoDocente(t *testing.T) {
	profesor := docente.NuevoDocente("2", "gonzalo tierra", "gonzalo@uide.edu.ec", "Sistemas", "Docente")

	if profesor.GetNombre() != "Gonzalo Tierra" {
		t.Errorf("El constructor no normalizó el nombre correctamente. Salió: %s", profesor.GetNombre())
	}

	if profesor.GetID() != "2" {
		t.Errorf("Error en GetID: Se esperaba '2' pero salió '%s'", profesor.GetID())
	}

	// NOTA DE DEFENSA: Si descomentas esto, el programa no compila porque son privados:
	// profesor.validarEmail()
	// profesor.normalizarNombre()
}
