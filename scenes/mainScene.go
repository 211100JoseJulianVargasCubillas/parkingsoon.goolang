package scenes

import (
    "estacionamiento/models"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "gonum.org/v1/gonum/stat/distuv"
    "sync"
    "time"
)

// Definición de la estructura MainScene
type MainScene struct {
    window fyne.Window
}

// Función NewMainScene crea una nueva instancia de MainScene
func NewMainScene(window fyne.Window) *MainScene {
    return &MainScene{
        window: window,
    }
}

var contenedor = container.NewWithoutLayout()

// Función Show muestra la escena principal
func (s *MainScene) Show() {
    // Reemplazar el rectángulo con una imagen
    imagenContorno := canvas.NewImageFromFile("./assets/parking2.jpg") // Asegúrate de cambiar la ruta al archivo de imagen correcto
    imagenContorno.FillMode = canvas.ImageFillContain // Puede cambiar a ImageFillOriginal si es necesario
    imagenContorno.Resize(fyne.NewSize(700, 550))
    imagenContorno.Move(fyne.NewPos(0, -80))

    contenedor.Add(imagenContorno) // Agregar la imagen al contenedor
    s.window.SetContent(contenedor)
}

// Función Run inicia la simulación del estacionamiento
func (s *MainScene) Run() {
    p := models.NewEstacionamiento(make(chan int, 20), &sync.Mutex{})

    var wg sync.WaitGroup

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(id int) {
            auto := models.NewAuto(id)
            imagen := auto.GetImagenEntrada()
            imagen.Resize(fyne.NewSize(30, 50))
            imagen.Move(fyne.NewPos(290, -10))

            contenedor.Add(imagen)
            contenedor.Refresh()

            auto.Iniciar(p, contenedor, &wg)
        }(i)

        var poisson = generarPoisson(float64(2)) // Genera un valor Poisson con una tasa de llegada de 2
        time.Sleep(time.Second * time.Duration(poisson))
    }

    wg.Wait()
}

// Función generarPoisson genera una variable aleatoria Poisson con una tasa dada
func generarPoisson(lambda float64) float64 {
    poisson := distuv.Poisson{Lambda: lambda, Src: nil}
    return poisson.Rand()
}
