package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	// Mundo 2D
	tamañoMundo         = 400
	tamañoCelda         = 40
	anchoMapa, altoMapa = 10, 10

	// Vista 3D
	anchoVista3D = 800
	altoVista3D  = tamañoMundo

	// Pantalla completa
	anchoPantalla = anchoVista3D
	altoPantalla  = altoVista3D

	// Raycasting
	cantidadRayos   = anchoVista3D // un rayo por columna
	anchoColumna    = float64(anchoVista3D) / float64(cantidadRayos)
	campoVision     = math.Pi / 3 // 60°
	largoMaximoRayo = tamañoMundo // no más allá del mundo
	distanciaNiebla = largoMaximoRayo

	// Movimiento y sensibilidad
	velocidadMovimiento = 1.0
	sensibilidadGiro    = 0.002 // mouse-look horizontal

	// Minimap
	tamañoMiniCelda = 10
	margenMiniMapa  = 8
)

var mapaMundo = [altoMapa][anchoMapa]int{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 1, 0, 1, 1, 0, 1, 0, 1},
	{1, 0, 1, 0, 0, 0, 0, 1, 0, 1},
	{1, 0, 1, 1, 1, 1, 0, 1, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 1, 0, 1},
	{1, 0, 1, 1, 1, 1, 0, 1, 0, 1},
	{1, 0, 1, 0, 0, 0, 0, 1, 0, 1},
	{1, 0, 0, 0, 1, 1, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

type Juego struct {
	posX, posY   float64 // posición del jugador en pixeles
	angulo       float64 // orientación horizontal (rad)
	ultimoMouseX int     // para delta X del mouse
}

func NuevoJuego() *Juego {
	return &Juego{
		posX:         tamañoMundo / 2,
		posY:         tamañoMundo / 2,
		angulo:       0,
		ultimoMouseX: anchoPantalla / 2,
	}
}

func (j *Juego) Update() error {
	// Capturar delta X del ratón
	mouseX, _ := ebiten.CursorPosition()
	deltaX := mouseX - j.ultimoMouseX
	j.ultimoMouseX = mouseX

	// Ajustar ángulo con sensibilidad
	j.angulo += float64(deltaX) * sensibilidadGiro

	// Movimiento adelante/atrás
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		nx := j.posX + math.Cos(j.angulo)*velocidadMovimiento
		ny := j.posY + math.Sin(j.angulo)*velocidadMovimiento
		if mapaMundo[int(ny)/tamañoCelda][int(nx)/tamañoCelda] == 0 {
			j.posX, j.posY = nx, ny
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		nx := j.posX - math.Cos(j.angulo)*velocidadMovimiento
		ny := j.posY - math.Sin(j.angulo)*velocidadMovimiento
		if mapaMundo[int(ny)/tamañoCelda][int(nx)/tamañoCelda] == 0 {
			j.posX, j.posY = nx, ny
		}
	}

	// Strafing lateral
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		nx := j.posX + math.Sin(j.angulo)*velocidadMovimiento
		ny := j.posY - math.Cos(j.angulo)*velocidadMovimiento
		if mapaMundo[int(ny)/tamañoCelda][int(nx)/tamañoCelda] == 0 {
			j.posX, j.posY = nx, ny
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		nx := j.posX - math.Sin(j.angulo)*velocidadMovimiento
		ny := j.posY + math.Cos(j.angulo)*velocidadMovimiento
		if mapaMundo[int(ny)/tamañoCelda][int(nx)/tamañoCelda] == 0 {
			j.posX, j.posY = nx, ny
		}
	}

	return nil
}

func (j *Juego) Draw(screen *ebiten.Image) {
	// ── VISTA 3D ────────────────────────────────────────────────────
	// Cielo
	ebitenutil.DrawRect(screen, 0, 0, anchoVista3D, altoVista3D/2,
		color.RGBA{100, 150, 255, 255})
	// Suelo
	ebitenutil.DrawRect(screen, 0, altoVista3D/2, anchoVista3D, altoVista3D/2,
		color.RGBA{50, 50, 50, 255})

	// Raycasting: una columna por cada rayo
	for i := 0; i < cantidadRayos; i++ {
		angRayo := j.angulo - campoVision/2 + campoVision*(float64(i)/float64(cantidadRayos))
		dirX, dirY := math.Cos(angRayo), math.Sin(angRayo)

		// Avanzar hasta colisión
		var distancia float64
		for t := 0.0; t < largoMaximoRayo; t++ {
			px := j.posX + dirX*t
			py := j.posY + dirY*t
			mx, my := int(px)/tamañoCelda, int(py)/tamañoCelda
			if mx < 0 || mx >= anchoMapa || my < 0 || my >= altoMapa ||
				mapaMundo[my][mx] == 1 {
				distancia = t
				break
			}
		}

		// Corrige fish-eye y convierte a “unidades de celda”
		perp := distancia * math.Cos(angRayo-j.angulo) / float64(tamañoCelda)
		if perp <= 0 {
			perp = 0.0001
		}
		altoCol := int(float64(altoVista3D) / perp)
		if altoCol > altoVista3D {
			altoCol = altoVista3D
		}
		inicioY := (altoVista3D - altoCol) / 2

		// Intensidad de niebla
		intensidad := 1 - distancia/distanciaNiebla
		if intensidad < 0 {
			intensidad = 0
		}
		c := uint8(255 * intensidad)

		// Dibuja slice vertical
		ebitenutil.DrawRect(screen,
			float64(i)*anchoColumna, float64(inicioY),
			anchoColumna, float64(altoCol),
			color.RGBA{c, c, c, 255},
		)
	}

	// ── MINIMAPA ────────────────────────────────────────────────────
	anchoMiniMapa := anchoMapa * tamañoMiniCelda
	altoMiniMapa := altoMapa * tamañoMiniCelda
	offX := anchoPantalla - anchoMiniMapa - margenMiniMapa
	offY := altoPantalla - altoMiniMapa - margenMiniMapa

	// Fondo semitransparente
	ebitenutil.DrawRect(screen,
		float64(offX), float64(offY),
		float64(anchoMiniMapa), float64(altoMiniMapa),
		color.RGBA{0, 0, 0, 192},
	)

	// Dibujar paredes en el minimapa
	for my := 0; my < altoMapa; my++ {
		for mx := 0; mx < anchoMapa; mx++ {
			if mapaMundo[my][mx] == 1 {
				ebitenutil.DrawRect(screen,
					float64(offX+mx*tamañoMiniCelda),
					float64(offY+my*tamañoMiniCelda),
					tamañoMiniCelda-1, tamañoMiniCelda-1,
					color.White,
				)
			}
		}
	}

	// Dibujar rayos en el minimapa (cada 40 para no saturar)
	rayColor := color.RGBA{255, 255, 0, 128}
	for i := 0; i < cantidadRayos; i += 40 {
		angRayo := j.angulo - campoVision/2 + campoVision*(float64(i)/float64(cantidadRayos))
		dirX, dirY := math.Cos(angRayo), math.Sin(angRayo)
		var hx, hy float64
		for t := 0.0; t < largoMaximoRayo; t++ {
			px := j.posX + dirX*t
			py := j.posY + dirY*t
			mx, my := int(px)/tamañoCelda, int(py)/tamañoCelda
			if mx < 0 || mx >= anchoMapa || my < 0 || my >= altoMapa ||
				mapaMundo[my][mx] == 1 {
				hx, hy = px, py
				break
			}
		}
		// Convertir a coordenadas de minimapa
		x0 := offX + int(j.posX/float64(tamañoCelda)*float64(tamañoMiniCelda))
		y0 := offY + int(j.posY/float64(tamañoCelda)*float64(tamañoMiniCelda))
		x1 := offX + int(hx/float64(tamañoCelda)*float64(tamañoMiniCelda))
		y1 := offY + int(hy/float64(tamañoCelda)*float64(tamañoMiniCelda))
		ebitenutil.DrawLine(screen, float64(x0), float64(y0), float64(x1), float64(y1), rayColor)
	}

	// Dibujar jugador en el minimapa
	px := offX + int(j.posX/float64(tamañoCelda)*float64(tamañoMiniCelda))
	py := offY + int(j.posY/float64(tamañoCelda)*float64(tamañoMiniCelda))
	ebitenutil.DrawRect(screen,
		float64(px-2), float64(py-2),
		4, 4,
		color.RGBA{255, 0, 0, 255},
	)
}

func (j *Juego) Layout(_, _ int) (int, int) {
	return anchoPantalla, altoPantalla
}

func main() {
	ebiten.SetCursorMode(ebiten.CursorModeCaptured)
	ebiten.SetWindowTitle("Raycasting – Variables en Español con Minimap")
	ebiten.SetWindowSize(anchoPantalla, altoPantalla)
	ebiten.SetWindowResizable(false)
	if err := ebiten.RunGame(NuevoJuego()); err != nil {
		log.Fatal(err)
	}
}
