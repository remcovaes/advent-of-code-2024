package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleA(t *testing.T) {
	input := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`
	result := TotalResultA(input)
	assert.Equal(t, 14, result)
}
func TestExampleA2(t *testing.T) {
	input := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
..........xx`
	result := TotalResultA(input)
	assert.Equal(t, 15, result)
}
func TestExampleA3(t *testing.T) {
	input := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
..zz.....A..
..zz........
..........xx`
	result := TotalResultA(input)
	assert.Equal(t, 27, result)
}
func TestExampleB(t *testing.T) {
	input := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

	result := TotalResultB(input)
	assert.Equal(t, 34, result)
}
func TestExampleB2(t *testing.T) {
	input := `T.........
...T......
.T........
..........
..........
..........
..........
..........
..........
..........`

	result := TotalResultB(input)
	assert.Equal(t, 9, result)
}
func TestFullA(t *testing.T) {
	input := GetDataDay()
	result := TotalResultA(input)
	assert.Equal(t, 336, result)
}

func TestFullB(t *testing.T) {
	input := GetDataDay()
	result := TotalResultB(input)
	assert.Equal(t, 1131, result)
}

func TestHelper1(t *testing.T) {
	input := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

	parsed := SplitIntoStringLines(input)

	result := FindLetters(parsed)
	assert.Equal(t, 4, len(result["0"]))
}
func TestHelper5(t *testing.T) {
	input := `......#.....
............
.....0......
............
....0.......
............
...#........
............
............
............
............
............`

	result := TotalResultA(input)
	assert.Equal(t, 2, result)
}

func TestHelper2(t *testing.T) {
	result := GetNewCoords(Coor{x: 4, y: 3}, Coor{x: 5, y: 5})
	first := result[0]
	second := result[1]
	assert.Equal(t, Coor{x: 3, y: 1}, first)
	assert.Equal(t, Coor{x: 6, y: 7}, second)
}

func TestHelper3(t *testing.T) {
	input := ``

	result := Helper3(input)
	assert.Equal(t, 0, result)
}