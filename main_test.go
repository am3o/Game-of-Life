package main

import (
	"image"
	"testing"
)

func TestGrid_countNeighbours(t *testing.T) {
	tt := []struct {
		name               string
		grid               Grid
		point              image.Point
		expectedNeighbours uint
	}{
		{
			name:               "empty area",
			grid:               map[image.Point]interface{}{},
			point:              image.Point{X: 0, Y: 0},
			expectedNeighbours: 0,
		},
		{
			name:               "only one point exists in the area",
			grid:               map[image.Point]interface{}{image.Point{}: 0},
			point:              image.Point{},
			expectedNeighbours: 0,
		},
		{
			name:               "only two points exists in the area and expect one neighbour",
			grid:               map[image.Point]interface{}{image.Point{}: 0, image.Point{X: 1}: 0},
			point:              image.Point{},
			expectedNeighbours: 1,
		},
		{
			name:               "three points exists in the area and expect one neighbour",
			grid:               map[image.Point]interface{}{image.Point{}: 0, image.Point{X: 1}: 0, image.Point{Y: 1}: 0},
			point:              image.Point{},
			expectedNeighbours: 2,
		},
		{
			name:               "three points exists in the area and expect one neighbour",
			grid:               map[image.Point]interface{}{image.Point{}: 0, image.Point{X: 1}: 0, image.Point{X: 5, Y: 1}: 0},
			point:              image.Point{},
			expectedNeighbours: 1,
		},
		{
			name: "eight points exists in the area and all of them are neighbours",
			grid: map[image.Point]interface{}{
				image.Point{X: 0, Y: 0}: 0, image.Point{X: 1, Y: 0}: 0, image.Point{X: 2, Y: 0}: 0,
				image.Point{X: 0, Y: 1}: 0, image.Point{X: 1, Y: 1}: 0, image.Point{X: 2, Y: 1}: 0,
				image.Point{X: 0, Y: 2}: 0, image.Point{X: 1, Y: 2}: 0, image.Point{X: 2, Y: 2}: 0,
			},
			point:              image.Point{X: 1, Y: 1},
			expectedNeighbours: 8,
		},
		{
			name: "many points exists in the area and only eight of them are neighbours",
			grid: map[image.Point]interface{}{
				image.Point{X: 0, Y: 0}: 0, image.Point{X: 1, Y: 0}: 0, image.Point{X: 2, Y: 0}: 0,
				image.Point{X: 0, Y: 1}: 0, image.Point{X: 1, Y: 1}: 0, image.Point{X: 2, Y: 1}: 0,
				image.Point{X: 0, Y: 2}: 0, image.Point{X: 1, Y: 2}: 0, image.Point{X: 2, Y: 2}: 0,
				image.Point{X: 0, Y: 3}: 0, image.Point{X: 1, Y: 3}: 0, image.Point{X: 2, Y: 3}: 0,
			},
			point:              image.Point{X: 1, Y: 1},
			expectedNeighbours: 8,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actualNeighbours := tc.grid.countNeighbours(tc.point)
			if actualNeighbours != tc.expectedNeighbours {
				t.Fail()
			}
		})
	}
}
