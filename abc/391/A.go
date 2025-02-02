package _391

import "fmt"

func A() {
	D := readString(r)

	newsMap := map[string]string{
		"N":  "S",
		"NE": "SW",
		"E":  "W",
		"SE": "NW",
		"S":  "N",
		"SW": "NE",
		"W":  "E",
		"NW": "SE",
	}

	fmt.Fprintln(w, newsMap[D])
}
