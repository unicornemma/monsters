package dice

import (
	"math/rand"
	"time"
)

func Roll() int64 {
	rando := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rando.Int63n(20)
}
