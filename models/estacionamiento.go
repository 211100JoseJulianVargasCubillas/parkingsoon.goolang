package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"sync"
)



// Definición de la estructura Estacionamiento
type Estacionamiento struct {
	cajones      chan int      // Un canal para gestionar los espacios del estacionamiento
	entrada        *sync.Mutex   // Un mutex (cerrojo) para gestionar el acceso a la puerta del estacionamiento
	cajonesArray [20]bool     // Un array de 20 booleanos para representar la disponibilidad de espacios en el estacionamiento
}

// Función NewEstacionamiento crea una nueva instancia de Estacionamiento
func NewEstacionamiento(cajones chan int, entradaMu *sync.Mutex) *Estacionamiento {
	return &Estacionamiento{
		cajones:      cajones,
		entrada:        entradaMu,
		cajonesArray: [20]bool{},
	}
}

// Función GetEspacios devuelve el canal para gestionar los espacios del estacionamiento
func (p *Estacionamiento) GetCajones() chan int {
	return p.cajones
}

// Función GetPuertaMu devuelve el mutex (cerrojo) para gestionar la puerta del estacionamiento
func (p *Estacionamiento) GetEntradaMu() *sync.Mutex {
	return p.entrada
}

// Función GetEspaciosArray devuelve el array que representa la disponibilidad de espacios en el estacionamiento
func (p *Estacionamiento) GetCajonesArray() [20]bool {
	return p.cajonesArray
}

// Función SetEspaciosArray establece el array que representa la disponibilidad de espacios en el estacionamiento
func (p *Estacionamiento) SetCajonesArray(cajonesArray [20]bool) {
	p.cajonesArray = cajonesArray
}

// Función ColaSalida agrega una imagen al contenedor y refresca la interfaz gráfica para representar un automóvil en cola de salida
func (p *Estacionamiento) ColaSalida(contenedor *fyne.Container, imagen *canvas.Image) {
	imagen.Move(fyne.NewPos(80, 280))
	contenedor.Add(imagen)
	contenedor.Refresh()
}
