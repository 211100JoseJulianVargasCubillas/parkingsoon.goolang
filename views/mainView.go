package views

import (
	"estacionamiento/scenes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// Definición de la estructura MainView
type MainView struct{}

// Función NewMainView crea una nueva instancia de MainView
func NewMainView() *MainView {
	return &MainView{}
}

// Función Run inicia la vista principal de la aplicación
func (v *MainView) Run() {
	myApp := app.New() // Crea una nueva instancia de la aplicación Fyne
	window := myApp.NewWindow("Estacionamiento concurrente") // Crea una nueva ventana con un título
	window.CenterOnScreen() // Centra la ventana en la pantalla
	window.SetFixedSize(true) // Establece el tamaño de la ventana como fijo
	window.Resize(fyne.NewSize(700, 400)) // Establece el tamaño de la ventana

	mainScene := scenes.NewMainScene(window) // Crea una instancia de la escena principal
	mainScene.Show() // Muestra la escena en la ventana
	go mainScene.Run() // Inicia la simulación en segundo plano

	window.ShowAndRun() // Muestra la ventana y ejecuta la aplicación
}
