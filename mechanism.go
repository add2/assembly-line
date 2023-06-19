package main

import (
	"fmt"
	"math/rand"
	"time"
)

type mechanism struct {
	id       int
	accident bool
}

func newMechanism(id int) *mechanism {
	return &mechanism{id: id}
}

func (m mechanism) name() string {
	return fmt.Sprintf("Mechanism_#%d", m.id)
}

func (m *mechanism) operate(accident chan<- int) {
	time.Sleep(mechanismOperateTime) // working time

	if rand.Float32() < accidentProbability {
		m.accident = true
		accident <- m.id
	}
}
