package main

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

func (m *manager) start() {
	for {
		err := m.cycle()

		if err != nil {
			fmt.Println("ERROR:", err)
			m.printAccidentsInfo()

			fmt.Println("Press ENTER to continue")
			_, _ = fmt.Fscanln(os.Stdin)

			m.resetMechanisms()
		}
	}
}

func (m manager) cycle() error {
	// При текущем алгоритме следующая проверка никогда не вернет ошибку.
	// Такая проверка может потребоваться на случай, когда при инициализации линии, механизм уже будет поврежден.
	if err := m.checkMechanisms(); err != nil {
		return err
	}

	m.moveLine()

	return m.performMechanismOperations()
}

func (m manager) moveLine() {
	fmt.Println("Line is moving")
	m.line.move()
}

func (m manager) performMechanismOperations() error {
	accidents := make(chan int, len(m.mechanisms))
	done := make(chan bool)
	var wg sync.WaitGroup

	for id := range m.mechanisms {
		wg.Add(1)

		mech := m.mechanisms[id]

		go func() {
			defer wg.Done()
			fmt.Printf(" - %s: is working\n", mech.name())
			mech.operate(accidents)
		}()
	}

	go func() {
		wg.Wait()
		close(done)
		close(accidents)
	}()

	timeout := time.After(mechanismTimeout)

	var err error

	select {
	case <-done:
		break
	case <-timeout:
		err = errors.New("timeout")
	case <-accidents:
		err = errors.New("there was an accident(s)")
	}

	return err
}

func (m manager) checkMechanisms() error {
	for _, mech := range m.mechanisms {
		if mech.accident {
			return errors.New("accident found")
		}
	}
	return nil
}

func (m *manager) resetMechanisms() {
	for id := range m.mechanisms {
		if mech := m.mechanisms[id]; mech.accident {
			mech.accident = false
		}
	}
}

func (m manager) printAccidentsInfo() {
	for _, mech := range m.mechanisms {
		if mech.accident {
			fmt.Printf(" - [broken] %s\n", mech.name())
		}
	}
}
