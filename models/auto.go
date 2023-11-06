package models

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/storage"
    "math/rand"
    "sync"
    "time"
)

// Definición de la estructura Auto
type Auto struct {
    id              int
    tiempoLim       time.Duration
    espacioAsignado int
    imagenEntrada   *canvas.Image
    imagenEspera    *canvas.Image
    imagenSalida    *canvas.Image
}

// Función NewAuto crea una nueva instancia de Auto
func NewAuto(id int) *Auto {
    // Carga imágenes desde archivos y establece campos iniciales
    imagenEntrada := canvas.NewImageFromURI(storage.NewFileURI("./assets/auto_entrada.png"))
    imagenEspera := canvas.NewImageFromURI(storage.NewFileURI("./assets/auto_espera.png"))
    imagenSalida := canvas.NewImageFromURI(storage.NewFileURI("./assets/auto_salida.png"))
    return &Auto{
        id:              id,
        tiempoLim:       time.Duration(rand.Intn(50)+50) * time.Second,
        espacioAsignado: 0,
        imagenEntrada:   imagenEntrada,
        imagenEspera:    imagenEspera,
        imagenSalida:    imagenSalida,
    }
}

// Función Entrar permite que el automóvil entre al estacionamiento
func (a *Auto) Entrar(p *Estacionamiento, contenedor *fyne.Container) {
    p.GetEspacios() <- a.GetId() // Intenta adquirir un espacio en el estacionamiento
    p.GetPuertaMu().Lock() // Bloquea la puerta de entrada

    espacios := p.GetEspaciosArray()
    const (
        columnasPorGrupo = 5
        espacioEntreGrupos = 22
        espacioHorizontal = 60
        espacioVertical = 100
    )

    // Encuentra un espacio vacío y asigna el automóvil a ese espacio
    for i := 0; i < len(espacios); i++ {
        if !espacios[i] {
            espacios[i] = true
            a.espacioAsignado = i

            fila := i / (columnasPorGrupo * 2)
            columna := i % (columnasPorGrupo * 2)

            if columna >= columnasPorGrupo {
                columna += 1
            }

            x := float32(35 + columna*espacioHorizontal)
            if columna >= columnasPorGrupo {
                x += espacioEntreGrupos
            }
            y := float32(250 + fila*espacioVertical)

            a.imagenEntrada.Move(fyne.NewPos(x, y))
            break
        }
    }

    p.SetEspaciosArray(espacios)
    p.GetPuertaMu().Unlock() // Desbloquea la puerta
    contenedor.Refresh() // Actualiza el contenedor de la interfaz gráfica
}

// Función Salir permite que el automóvil salga del estacionamiento
func (a *Auto) Salir(p *Estacionamiento, contenedor *fyne.Container) {
    <-p.GetEspacios() // Libera el espacio en el estacionamiento
    p.GetPuertaMu().Lock() // Bloquea la puerta

    spacesArray := p.GetEspaciosArray()
    spacesArray[a.espacioAsignado] = false
    p.SetEspaciosArray(spacesArray)

    p.GetPuertaMu().Unlock() // Desbloquea la puerta

    contenedor.Remove(a.imagenEspera)
    a.imagenSalida.Resize(fyne.NewSize(30, 50))
    a.imagenSalida.Move(fyne.NewPos(300, 290))

    contenedor.Add(a.imagenSalida)
    contenedor.Refresh()

    // Realiza una animación de salida
    for i := 0; i < 10; i++ {
        a.imagenSalida.Move(fyne.NewPos(a.imagenSalida.Position().X, a.imagenSalida.Position().Y-30))
        time.Sleep(time.Millisecond * 200)
    }

    contenedor.Remove(a.imagenSalida)
    contenedor.Refresh()
}

// Función Iniciar inicia el proceso del automóvil en el estacionamiento
func (a *Auto) Iniciar(p *Estacionamiento, contenedor *fyne.Container, wg *sync.WaitGroup) {
    a.Avanzar(9) // Realiza una animación de avance

    a.Entrar(p, contenedor) // El automóvil entra en el estacionamiento

    time.Sleep(a.tiempoLim) // Espera el tiempo límite

    contenedor.Remove(a.imagenEntrada)
    a.imagenEspera.Resize(fyne.NewSize(50, 30))
    p.ColaSalida(contenedor, a.imagenEspera) // El automóvil se prepara para salir
    a.Salir(p, contenedor) // El automóvil sale del estacionamiento

    wg.Done() // Indica que el automóvil ha terminado su proceso
}

// Función Avanzar realiza una animación de avance del automóvil
func (a *Auto) Avanzar(pasos int) {
    for i := 0; i < pasos; i++ {
        a.imagenEntrada.Move(fyne.NewPos(a.imagenEntrada.Position().X, a.imagenEntrada.Position().Y+20))
        time.Sleep(time.Millisecond * 200)
    }
}

// Función GetId devuelve el identificador del automóvil
func (a *Auto) GetId() int {
    return a.id
}

// Función GetImagenEntrada devuelve la imagen de entrada del automóvil
func (a *Auto) GetImagenEntrada() *canvas.Image {
    return a.imagenEntrada
}
