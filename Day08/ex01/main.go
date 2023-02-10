package main

import (
	"fmt"
	"reflect"
	"strings"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func describePlant(plant any) {
	plantV := reflect.ValueOf(plant)

	for i := 0; i < plantV.NumField(); i++ {
		str := string(reflect.TypeOf(plant).Field(i).Tag)
		str = strings.Replace(str, ":", "=", 1)
		str = strings.Replace(str, "\"", "", -1)
		if str != "" {
			str = "(" + str + ")"
		}
		str = string(reflect.TypeOf(plant).Field(i).Name) + str
		fmt.Printf("%s:%v\n", str, plantV.Field(i).Interface())
	}
}

func main() {
	describePlant(AnotherUnknownPlant{10, "lanceolate", 15})

	fmt.Println()
	describePlant(UnknownPlant{"amazing blue-green shade", "shaped like a claw", 25})

	fmt.Println()
	describePlant(AnotherUnknownPlant{40, "shape of an elegant shoe", 60})

	fmt.Println()
	describePlant(UnknownPlant{"amorphophallus titanic", "foxtail", 85})

	fmt.Println()
	describePlant(AnotherUnknownPlant{115, "beak of a parrot", 150})

	fmt.Println()
	describePlant(UnknownPlant{"rafflesia arnoldi", "oval", 190})
}
