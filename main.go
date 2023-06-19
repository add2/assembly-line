package main

import "time"

/**
Требуется написать программу управления сборочным конвейером.
Конвейер состоит из движущейся ленты, на которой установлены детали, и исполнительных механизмов
(далее ИМ для краткости).

Конвейер разбит на участки, каждый из которых обслуживает отдельный ИМ. Технологический цикл работы конвейера
включает в себя сдвиг ленты и обработку деталей.

Сдвиг ленты возможен, только если все ИМ сообщили об успешном выполнении операции. Если ИМ сообщает об аварии
или не отвечает в течение заданного времени, то конвейер останавливается и регистрируется авария, после чего возврат
в автоматический режим работы возможен только по команде оператора.

После сдвига ленты, ИМ по команде управляющей программы выполняет одну технологическую операцию над деталью. После
того как все ИМ успешно отработали операцию, технологический цикл повторяется.

Программу можно написать на любом знакомом языке или псевдокоде
*/

const numberOfMechanism = 5
const accidentProbability = 0.1
const lineMoveTime = 500 * time.Millisecond
const mechanismOperateTime = 500 * time.Millisecond
const mechanismTimeout = 750 * time.Millisecond

func main() {
	l := newLine()
	m := newManager(l)

	for i := 0; i < numberOfMechanism; i++ {
		mech := newMechanism(i + 1)
		m.append(mech)
	}

	m.start()
}
