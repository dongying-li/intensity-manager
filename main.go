package main

import (
	"fmt"
	im "github.com/dongying-li/jerry-take-home/intensitymanager"
)
func main() {
	// Example Usage
	im := &im.IntensityManager{Segments: []im.Segment{}}
	im.Add(10, 20, 1)
	fmt.Println(im.Segments)
	im.Set(15, 17, 2)
	fmt.Println(im.Segments)

}
