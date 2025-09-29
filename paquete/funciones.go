package PAQUETE

import "strings"


type Persona struct {
	NombreCompleto string
}
var Personas []string

func validacionHM(nombreCompleto string) bool {
	primerNombre := strings.Fields(nombreCompleto)[0]
	if len(primerNombre) == 0 {
		return false
		
	}
	ultima := primerNombre[len(primerNombre)-1]
	return !(ultima == 'a' || ultima == 'A')
}

// Obtiene los apellidos de una persona
func obtenerApellidos(nombreCompleto string) (string, string) {
	partes := strings.Fields(nombreCompleto)
	if len(partes) < 3 {
		return "", ""
	}
	return partes[len(partes)-2], partes[len(partes)-1]
}

//  busca una persona por nombre exacto
func BuscarPersonaPorNombre(nombre string) *Persona {
	for _, p := range Personas {
		if strings.EqualFold(p, nombre) {
			return &Persona{NombreCompleto: p}
		}
	}
	return nil
}

//  busca los padres
func ObtenerPadres(persona *Persona) (string, string) {
	apellidoPadre, apellidoMadre := obtenerApellidos(persona.NombreCompleto)
	var padre, madre string

	for _, p := range Personas {
		if p == persona.NombreCompleto {
			continue
		}
		a1, _ := obtenerApellidos(p)
		
		if padre == "" && validacionHM(p) && a1 == apellidoPadre {
			padre = p
		}
		if madre == "" && !validacionHM(p) && a1 == apellidoMadre {
			madre = p
		}
		if padre != "" && madre != "" {
			break
		}
	}
	return padre, madre
}

// busca los hermanos
func ObtenerHermanos(persona *Persona) []string {
	apellido1, apellido2 := obtenerApellidos(persona.NombreCompleto)
	var hermanos []string

	for _, p := range Personas {
		if p == persona.NombreCompleto {
			continue
		}
		a1, a2 := obtenerApellidos(p)
		if a1 == apellido1 && a2 == apellido2 {
		
			hermanos = append(hermanos, p)
		}
	}
	return hermanos
}
//mostrar hermanos 
func MostrarHermanos(hermanos []string) {

	if len(hermanos) == 0 {
		println("No se encontraron hermanos.")
		return
	}
	println("Hermanos:")
	for _, h := range hermanos {
		println("- " + h)
	}
}