package robotname

import (
	"errors"
	"math/rand"
)

// Define the Robot type here.
type Robot struct {
	name    string
	hasName bool
}

var namesMap = make(map[string]bool)
var namesGenerated = 0
var maxNamesLimit = 26 * 26 * 10 * 10 * 10

func (r *Robot) Name() (string, error) {
	if r.hasName {
		return r.name, nil
	}

	if namesGenerated >= maxNamesLimit {
		return "", errors.New("Every possible robot name has been allocated")
	} else {
		for {
			name := GetRandomName()
			if _, ok := namesMap[name]; !ok {
				namesMap[name] = true
				namesGenerated++
				r.hasName = true
				r.name = name
				return name, nil
			}
		}
	}
}

func (r *Robot) Reset() {
	r.name = ""
	r.hasName = false
}

func GetRandomName() string {
	letterRunes := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	digitsRunes := []rune("0123456789")

	r1 := letterRunes[rand.Intn(len(letterRunes))]
	r2 := letterRunes[rand.Intn(len(letterRunes))]
	r3 := digitsRunes[rand.Intn(len(digitsRunes))]
	r4 := digitsRunes[rand.Intn(len(digitsRunes))]
	r5 := digitsRunes[rand.Intn(len(digitsRunes))]

	randomName := ""
	randomName = string(r1) + string(r2) + string(r3) + string(r4) + string(r5)

	return randomName
}
