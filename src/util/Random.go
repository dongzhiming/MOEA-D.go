package util

import (
	"math/rand"
	"time"
)

type Random struct {
	r *rand.Rand
}

func (this *Random) NewRand() {
	this.r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func (this *Random) Float64() float64 {
	return this.r.Float64()
}

func (this *Random) Perm(n int) []int {
	return this.r.Perm(n)
}

func (this *Random) Intn(n int) int {
	return this.r.Intn(n)
}
