package funciones

import (
	"strings"
)

// Buscar persona por nombre parcial (ejemplo: "Pedro")
func buscarPersonaPorNombre(nombre string) *Persona {
	for _, p := range personas {
		if strings.Contains(p.NombreCompleto, nombre) {
			return &p
		}
	}
	return nil
}

// Obtener padres
func obtenerPadres(p *Persona) []string {
	return p.Padres
}

// Obtener cantidad de hermanos
func obtenerCantidadHermanos(p *Persona) int {
	return len(p.Hermanos)
}
