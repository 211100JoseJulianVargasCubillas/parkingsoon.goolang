package main

import "estacionamiento/views"

func main() {
    mainView := views.NewMainView() // Crea una instancia de la vista principal
    mainView.Run() // Ejecuta la vista principal
}
