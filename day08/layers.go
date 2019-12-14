package day08

import (
	"errors"
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
	"math"
	"regexp"
)

const TRANSPARENT = 2

var renderRepresentation = map[int]string{0: " ", 1: "#"}
var digitsRegex = regexp.MustCompile(`^[0-9]+$`)

type Image struct {
	layers []ImageLayer
}

type ImageLayer struct {
	data [][]int
}

func (i *Image) Render() [][]string {
	dim := i.getDimensions()
	ret := make([][]string, dim[1])
	
	for x := 0; x < dim[1]; x++ {
		ret[x] = make([]string, dim[0])
		for y := 0; y < dim[0]; y++ {
			pixel, _ := i.getDominantPixel(x, y)
			representation, ok := renderRepresentation[pixel]; if ! ok {
				representation = " "
			}
			ret[x][y] = representation
		}
	}
	return ret
}

func (i *Image) getDominantPixel(x, y int) (int, error) {
	for _, layer := range i.layers {
		pixel := layer.data[x][y]
		if pixel != TRANSPARENT {
			return pixel, nil
		}
	}

	return TRANSPARENT, nil
}

func (i *Image) GetLayerWithFewestDigits(digit int) *ImageLayer {
	if i.layers == nil {
		return nil
	}

	var minLayer ImageLayer
	min := math.MaxInt32
	for _, layer := range i.layers {
		digitsPerLayer := layer.Count(digit)
		if digitsPerLayer < min {
			minLayer = layer
			min = digitsPerLayer
		}
	}
	
	return &minLayer
}

func (i *ImageLayer) Count(digit int) int {
	cnt := 0
	for row := 0; row < len(i.data); row++ {
		for _, x := range i.data[row] {
			if x == digit {
				cnt++
			}
		}
	}
	return cnt
}

func Answer08() {
	data := util.ReadStringLinesFromFile("resources/day08/input.txt")
	img, _ := ParseImage(data[0],   25, 6)
	layer := img.GetLayerWithFewestDigits(0)
	a := layer.Count(1)
	b := layer.Count(2)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	
	rendered := img.Render()
	for i := 0; i < len(rendered); i++ {
		if len(rendered) > 0 {
			fmt.Println(rendered[i])
		}
	}
}

func ParseImage(data string, width, height int) (*Image, error) {
	if len(data) % (width * height) != 0 {
		return nil, fmt.Errorf("invalid dimensions")
	}

	ret := make([]ImageLayer, len(data) / (width*height))
	i := 0
	for layerIndex := 0; layerIndex < len(data) - 1; layerIndex += (width * height) {
		layer := make([][]int, height)
		j := 0
		for layerRow := layerIndex; layerRow < layerIndex + (width*height); layerRow += width {
			da := data[layerRow : layerRow+width]
			array, _ := getArray(da)
			layer[j] = array
			j++
		}
		
		ret[i] = ImageLayer{layer}
		i++
	}

	return &Image{layers: ret}, nil
}

func getArray(s string) ([]int, error) {
	if ! digitsRegex.MatchString(s) {
		return nil, errors.New("not a number")
	}

	ret := make([]int, len(s))
	for i, value := range s {
		ret[i] = int(value - '0')
	}
	return ret, nil
}

// getDimensions returns an array of the image's dimensions.
// First coordinate is the width, 2nd coordinate is height.
func (i *Image) getDimensions() []int {
	if nil == i.layers || len(i.layers) == 0 {
		return []int{0,0}
	}

	x := len(i.layers[0].data[0])
	y := len(i.layers[0].data)
	return []int{x,y}
}