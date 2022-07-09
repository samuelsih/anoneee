package faker

import "math/rand"

func randInt(min, max int) int {
	return min + rand.Intn(max - min)
}

func randBool() bool {
	t := []bool{true, false}
	return t[rand.Intn(2)]
}