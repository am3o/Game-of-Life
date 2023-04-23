package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math/rand"
	"os"
	"time"
)

var matrix = []image.Point{
	{X: -1, Y: 1},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: 1, Y: 0},
	{X: 1, Y: -1},
	{X: 0, Y: -1},
	{X: -1, Y: -1},
	{X: -1, Y: 0},
}

type Grid map[image.Point]interface{}

func (g Grid) evolve() Grid {
	grid := make(map[image.Point]interface{})
	for origin := range g {
		points := []image.Point{origin}
		for _, neighbour := range matrix {
			points = append(points, image.Point{
				X: origin.X + neighbour.X,
				Y: origin.Y + neighbour.Y,
			})
		}

		for _, point := range points {
			switch g.countNeighbours(point) {
			case 2:
				if _, exists := g[point]; exists {
					// - Any live cell with two or three live neighbours survives.
					grid[point] = nil
				}
			case 3:
				grid[point] = nil
			default:
				// - Any dead cell with three live neighbours becomes a live cell.
				// - All other live cells die in the next generation. Similarly, all other dead cells stay dead.
			}
		}
	}
	return grid
}

func (g Grid) countNeighbours(point image.Point) uint {
	var neighbours uint
	for _, neighbour := range matrix {
		if _, exists := g[image.Point{X: point.X + neighbour.X, Y: point.Y + neighbour.Y}]; exists {
			neighbours += 1
		}
	}

	return neighbours
}

func createImage(width, height int, grid Grid) *image.Paletted {
	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0xff, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0xff, 0xff},
		color.RGBA{0xff, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0xff},
	}

	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

	for position := range grid {
		img.Set(position.X, position.Y, color.White)
	}

	return img
}

func main() {
	var (
		width, height = 1_600, 1_200
	)

	start := time.Now()
	defer fmt.Printf("time used to calculate 1.000 iterations: %v", time.Since(start).Milliseconds())

	example := Grid{}
	for i := 0; i < 1_500_000; i++ {
		example[image.Point{
			X: int(rand.Uint32() % uint32(width)),
			Y: int(rand.Uint32() % uint32(height)),
		}] = nil
	}

	series := []Grid{example}
	for i := 0; i < 1_000; i++ {
		series = append(series, series[len(series)-1].evolve())
	}

	var (
		images []*image.Paletted
		delays []int
	)

	for _, serie := range series {
		images = append(images, createImage(width, height, serie))
		delays = append(delays, 1)
	}

	f, err := os.OpenFile("game_of_life.gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	if err := gif.EncodeAll(f,
		&gif.GIF{
			Image: images,
			Delay: delays,
		},
	); err != nil {
		panic(err)
	}
}
