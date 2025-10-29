/*
Функция or-channel
Необходимо пеализовать функцию объединения done-каналов (описана в задаче 14 уровня 2). Проект хоть и небольшой, но концептуально важен: правильная работа с конкурентностью и каналами.

Готовая функция может быть оформлена как утилита/пакет с примерами использования и тестами.

Результат: пакет or (например, or.go + or_test.go), экспортирующий функцию Or(ch1, ch2, ... chN <-chan interface{}) <-chan interface{}.

Дополнительно: реализовать тесты (or_test.go) и пример использования.
*/
package or

func Or(channels ...<-chan any) <-chan any {

	if len(channels) == 0 {
		return nil
	}
	if len(channels) == 1 {
		return channels[0]
	}

	orDoneChannel := make(chan any)
	go func() {
		defer close(orDoneChannel)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-Or(append(channels[3:], orDoneChannel)...):
			}

		}
	}()

	return orDoneChannel
}
