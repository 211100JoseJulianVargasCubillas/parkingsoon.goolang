package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"sync"
)

// Definición de la estructura Estacionamiento
type Estacionamiento struct {
	espacios      chan int      // Un canal para gestionar los espacios del estacionamiento
	puerta        *sync.Mutex   // Un mutex (cerrojo) para gestionar el acceso a la puerta del estacionamiento
	espaciosArray [20]bool     // Un array de 20 booleanos para representar la disponibilidad de espacios en el estacionamiento
}

// Función NewEstacionamiento crea una nueva instancia de Estacionamiento
func NewEstacionamiento(espacios chan int, puertaMu *sync.Mutex) *Estacionamiento {
	return &Estacionamiento{
		espacios:      espacios,
		puerta:        puertaMu,
		espaciosArray: [20]bool{},
	}
}

// Función GetEspacios devuelve el canal para gestionar los espacios del estacionamiento
func (p *Estacionamiento) GetEspacios() chan int {
	return p.espacios
}

// Función GetPuertaMu devuelve el mutex (cerrojo) para gestionar la puerta del estacionamiento
func (p *Estacionamiento) GetPuertaMu() *sync.Mutex {
	return p.puerta
}

// Función GetEspaciosArray devuelve el array que representa la disponibilidad de espacios en el estacionamiento
func (p *Estacionamiento) GetEspaciosArray() [20]bool {
	return p.espaciosArray
}

// Función SetEspaciosArray establece el array que representa la disponibilidad de espacios en el estacionamiento
func (p *Estacionamiento) SetEspaciosArray(espaciosArray [20]bool) {
	p.espaciosArray = espaciosArray
}

// Función ColaSalida agrega una imagen al contenedor y refresca la interfaz gráfica para representar un automóvil en cola de salida
func (p *Estacionamiento) ColaSalida(contenedor *fyne.Container, imagen *canvas.Image) {
	imagen.Move(fyne.NewPos(80, 280))
	contenedor.Add(imagen)
	contenedor.Refresh()
}
