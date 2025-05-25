package groupsVals

import (
	"fmt"
	"math/rand"
	"time"
)

type Secret struct {
	Val string
}

func NewSecret() Secret {
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	return Secret{Val: code}
}

func ExistingSecret(val string) Secret {

	return Secret{Val: val}
}
