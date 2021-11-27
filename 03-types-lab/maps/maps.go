package main

import (
	"fmt"
	"log"
	"math"
)

type Vertex struct{ X, Y float64 }

func (v Vertex) Distance(o Vertex) float64 {
	return math.Sqrt((v.X-o.X)*(v.X-o.X) + (v.Y-o.Y)*(v.Y-o.Y))
}

func CalculateDistance(places map[string]Vertex, from string, to string) (float64, error) {
	if v1, ok := places[from]; ok {
		if v2, ok2 := places[to]; ok2 {
			return v1.Distance(v2), nil
		}
	}
	return 0.0, fmt.Errorf("Invalid place name(s) - from: %s, to: %s\n", from, to)
}

func main() {
	v1 := Vertex{3, 4}
	v2 := Vertex{7, 7}
	fmt.Println(v1.Distance(v2))

	places := map[string]Vertex{
		"Bell Labs": Vertex{40.68433, -74.39967},
		"Microsoft": Vertex{60.68433, -84.39967},
	}

	dist, err := CalculateDistance(places, "Bell Labs", "Microsoft")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Distance between %s and %s is %f\n", "Bell Labs", "Microsoft", dist)
}
