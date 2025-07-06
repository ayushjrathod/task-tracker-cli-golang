package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateID() string {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(90) + 10
	return strconv.Itoa(id)
}
