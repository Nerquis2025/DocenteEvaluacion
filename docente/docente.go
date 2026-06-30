package docente

import (
	"strings"
)

// Estructura principal del docente
type Docente struct {
	ID           string
	Nombre       string
	Email        string
	Departamento string
	Cargo        string
	evaluaciones []string // Es privado por empezar con minúscula
}

// Constructor público para crear un nuevo docente
func NuevoDocente(id, nombre, email, depto, cargo string) *Docente {
	d := &Docente{
		ID:           id,
		Nombre:       nombre,
		Email:        email,
		Departamento: depto,
		Cargo:        cargo,
		evaluaciones: []string{}, // Inicializa la lista vacía
	}

	// Corregimos el formato del nombre automáticamente al crear al docente
	d.Nombre = d.normalizarNombre()
	return d
}

// === MÉTODOS PÚBLICOS ===
func (d *Docente) GetID() string {
	return d.ID
}

func (d *Docente) GetNombre() string {
	return d.Nombre
}

func (d *Docente) GetEvaluaciones() []string {
	return d.evaluaciones
}

// === MÉTODOS PRIVADOS / ENCAPSULADOS ===

// validarEmail revisa que la dirección tenga un arroba y un punto
func (d *Docente) validarEmail() bool {
	tieneArroba := strings.Contains(d.Email, "@")
	tienePunto := strings.Contains(d.Email, ".")
	return tieneArroba && tienePunto
}

// normalizarNombre arregla el texto: todo a minúsculas y luego iniciales mayúsculas
func (d *Docente) normalizarNombre() string {
	nombreMinuscula := strings.ToLower(d.Nombre)
	nombreArreglado := strings.Title(nombreMinuscula)
	return nombreArreglado
}

// agregarEvaluacionInterna añade un ID de examen o evaluación a la lista oculta
func (d *Docente) agregarEvaluacionInterna(idEvaluacion string) {
	d.evaluaciones = append(d.evaluaciones, idEvaluacion)
}
// cambio final 
