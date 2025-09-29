package main

import (
	"fmt"
)

// Estructura de Persona
type Persona struct {
	NombreCompleto string
	Padres         []string
	Hermanos       []string
}

// Base de datos simulada
var personas = []string{
	"Adriana Pena Jimenez",
	"Hernan Rojas Perez",
	"Carolina Sanchez Mejia",
	"Pedro Jose Perez Lopez",
	"Laura Castillo Martinez",
	"Camilo Castro Molina",
	"Paula Herrera Suarez",
	"Javier Molina Torres",
	"Beatriz Arias Lozano",
	"Juliana Vargas Ortiz",
	"Julio Diaz Romero",
	"Tatiana Castano Arias",
	"Carlos Andres Herrera Suarez",
	"Claudia Diaz Romero",
	"Raul Pena Jimenez",
	"Cristian Castano Arias",
	"Daniela Gomez Castillo",
	"Luisa Morales Peña",
	"Beatriz Romero Gil",
	"Camilo Diaz Romero",
	"Paola Martinez Sanchez",
	"Carolina Martinez Sanchez",
	"Manuel Mejia Ruiz",
	"Maria Fernanda Torres Herrera",
	"Sandra Vargas Torres",
	"Martin Gil Vargas",
	"Sebastian Lopez Diaz",
	"Ricardo Molina Torres",
	"Andres Gil Vargas",
	"Clara Ruiz Rojas",
	"Daniel Ruiz Beltran",
	"Angela Diaz Morales",
	"Patricia Sanchez Gomez",
	"Pablo Pena Jimenez",
	"Rodrigo Sanchez Mejia",
	"Hernan Pena Jimenez",
	"Nicolas Vargas Ortiz",
	"Laura Rojas Perez",
	"Marcela Beltran Gomez",
	"David Martinez Sanchez",
	"Oscar Torres Herrera",
	"Carolina Gil Vargas",
	"Gloria Jimenez Rojas",
	"Mateo Lopez Diaz",
	"Esteban Mejia Ruiz",
	"Gloria Naranjo Perez",
	"Tatiana Castano Arias",
	"Monica Herrera Suarez",
	"Victor Morales Peña",
	"Maria Lopez Ramirez",
	"Diana Ortiz Lopez",
	"Catalina Rojas Perez",
	"Andres Vargas Ortiz",
	"Orlando Ruiz Beltran",
	"Valeria Castro Molina",
	"Angela Diaz Morales",
	"Silvia Molina Rios",
	"Santiago Ruiz Beltran",
	"Mauricio Vargas Ortiz",
	"Pablo Pena Jimenez",
	"Oscar Morales Peña",
	"Patricia Torres Sanchez",
	"Valentina Molina Torres",
	"Luis Fernando Perez Lopez",
	"Javier Castro Molina",
	"Felipe Castro Molina",
	"Hector Ramirez Naranjo",
	"Alberto Martinez Sanchez",
	"Ricardo Lopez Diaz",
	"Rosa Suarez Diaz",
	"Ana Sofia Perez Lopez",
	"Miguel Rodriguez Torres",
	"Jorge Gomez Castillo",
	"Carolina Ramirez Naranjo",
	"Enrique Gil Vargas",
	"Teresa Peña Suarez",
	"Juan Pablo Rodriguez Torres",
	"Camila Gomez Castillo",
	"Julian Ramirez Naranjo",
	"Carlos Perez Lopez",
	"Luz Perez Marin",
	"Paola Martinez Sanchez",
	"Marcela Beltran Gomez",
	"Natalia Ramirez Naranjo",
	"Valentina Rodriguez Torres",
	"Andres Felipe Gomez Castillo",
	"Daniel Morales Peña",
	"Sofia Sanchez Mejia",
	"Santiago Rodriguez Torres",
	"Manuel Mejia Ruiz",
	"Claudia Torres Ruiz",
	"Alfonso Castano Arias",
	"Francisco Diaz Romero",
	"Sofia Torres Herrera",
	"Ricardo Molina Torres",
	"Liliana Mejia Ruiz",
	"German Herrera Suarez",
	"Fernando Mejia Ruiz",
	"Jorge Molina Torres",
}

func main() {
	var nombre string
	fmt.Print("Ingrese un nombre: ")
	fmt.Scanln(&nombre)

	persona := buscarPersonaPorNombre(nombre)
	if persona == nil {
		fmt.Println("No se encontró ninguna persona con ese nombre.")
		return
	}

	fmt.Printf("\nPersona encontrada: %s\n", persona.NombreCompleto)
	fmt.Printf("Padres: %v\n", obtenerPadres(persona))
	fmt.Printf("Cantidad de hermanos: %d\n", obtenerCantidadHermanos(persona))
}
