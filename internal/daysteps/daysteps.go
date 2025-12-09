package daysteps

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	//Алгоритм реализации функции:
	//1. Разделить строку на слайс строк.

	parts := strings.Split(data, ",")

	//2. Проверить, чтобы длина слайса была равна 2, так как в строке данных у нас количество шагов и продолжительность.

	if len(parts) != 2 {
		return 0, 0, errors.New("incorrect data format")
	}

	//3. Преобразовать первый элемент слайса (количество шагов) в тип int. Обработать возможные ошибки. При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	//4. Проверить: количество шагов должно быть больше 0. Если это не так, вернуть нули и ошибку.

	if steps <= 0 {
		return 0, 0, errors.New("the number of steps must be greater than zero")
	}

	//5. Преобразовать второй элемент слайса в time.Duration. В пакете time есть метод для парсинга строки в time.Duration. Обработать возможные ошибки. При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, err
	}

	if duration <= 0 {
		return 0, 0, errors.New("the duration must be greater than zero")
	}

	//6. Если всё прошло без ошибок, верните количество шагов, продолжительность и nil (для ошибки).

	return steps, duration, nil

}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
}
