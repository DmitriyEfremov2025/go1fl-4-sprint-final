package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	// Алгоритм реализации функции:
	// 1. Разделить строку на слайс строк.
	parts := strings.Split(data, ",")

	// 2. Проверить, чтобы длина слайса была равна 3, так как в строке данных у нас количество шагов, вид активности и продолжительность.
	if len(parts) != 3 {
		return 0, "", 0, errors.New("incorrect data format")
	}
	// 3. Преобразовать первый элемент слайса (количество шагов) в тип int. Обработать возможные ошибки. При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, errors.New("invalid step values")
	}

	if steps <= 0 {
		return 0, "", 0, errors.New("the number of steps must be greater than zero")
	}

	// 4. Преобразовать третий элемент слайса в time.Duration. В пакете time есть метод для парсинга строки в time.Duration. Обработать возможные ошибки. При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, errors.New("incorrect duration format")
	}
	if duration <= 0 {
		return 0, "", 0, errors.New("the duration must be greater than zero")
	}
	// 5. Если всё прошло без ошибок, верните количество шагов, вид активности, продолжительность и nil (для ошибки).
	return steps, parts[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	// рассчитайте длину шага. Для этого умножьте высоту пользователя на коэффициент длины шага stepLengthCoefficient. Соответствующая константа уже определена в пакете.
	// умножьте пройденное количество шагов на длину шага.
	// разделите полученное значение на число метров в километре (mInKm, константа определена в пакете).
	// Обратите внимание, что целочисленную переменную steps необходимо будет привести к другому числовому типу.

	if steps <= 0 || height <= 0 {
		return 0.0
	}
	stepLength := height * stepLengthCoefficient
	distanceMeter := float64(steps) * stepLength
	distanceKm := distanceMeter / mInKm
	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	// Алгоритм реализации функции:
	// 1. Проверить, что продолжительность duration больше 0. Если это не так, вернуть 0.
	if duration <= 0 {
		return 0
	}
	// 2. Вычислить дистанцию с помощью distance().
	dist := distance(steps, height)

	// 3. Вычислить и вернуть среднюю скорость. Для этого разделите дистанцию на продолжительность в часах. Чтобы перевести продолжительность в часы, воспользуйтесь функцией из пакета time.
	durationInHours := duration.Hours()
	speed := dist / durationInHours
	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	// Алгоритм реализации функции:
	// 1. Получить значения из строки данных с помощью функции parseTraining(), обработать возможные ошибки и вывести их в лог с помощью log.Println(err).
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println("error: ", err)
		return "", err
	}
	// 2. Проверить, какой вид тренировки был передан в строке, которую парсили (лучше использовать switch). Для каждого из видов тренировки рассчитать дистанцию, среднюю скорость и калории.
	// 3. Для каждого вида тренировки сформировать и вернуть строку, образец которой был представлен выше.
	// 4. Если был передан неизвестный тип тренировки, вернуть ошибку с текстом неизвестный тип тренировки.

	var dist, speed, calories float64
	switch activity {
	case "Ходьба":
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	case "Бег":
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	durationInHours := duration.Hours()

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity, durationInHours, dist, speed, calories,
	)

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// 1. Проверить входные параметры на корректность. Если параметры некорректны, вернуть 0 калорий и соответствующую ошибку.

	if steps <= 0 {
		return 0, errors.New("steps must be more than zero")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be more than zero")
	}
	if height <= 0 {
		return 0, errors.New("height must be more than zero")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be more than zero")
	}

	// 2. Рассчитать среднюю скорость с помощью meanSpeed().
	speed := meanSpeed(steps, height, duration)
	if speed <= 0 {
		return 0, errors.New("speed must be more than zero")
	}

	// 3. Рассчитать и вернуть количество калорий. Для этого:
	// - Переведите продолжительность в минуты с помощью функции из пакета time.
	// - Умножьте вес пользователя на среднюю скорость и продолжительность в минутах.
	// - Разделите результат на число минут в часе для получения количества потраченных калорий.

	durationInMinutes := duration.Minutes()
	calories := (weight * speed * durationInMinutes) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// 1. Проверить входные параметры на корректность. Если параметры некорректны, вернуть 0 калорий и соответствующую ошибку.
	if steps <= 0 {
		return 0, errors.New("steps must be more than zero")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be more than zero")
	}
	if height <= 0 {
		return 0, errors.New("height must be more than zero")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be more than zero")
	}

	// 2. Рассчитать среднюю скорость с помощью meanSpeed().
	speed := meanSpeed(steps, height, duration)
	if speed <= 0 {
		return 0, errors.New("speed must be more than zero")
	}
	// 3. Рассчитать количество калорий. Для этого:
	// - Переведите продолжительность в минуты с помощью функции из пакета time.
	// - Умножьте вес пользователя на среднюю скорость и продолжительность в минутах.
	// - Разделите результат на число минут в часе для получения количества потраченных калорий.
	// Умножить полученное число калорий на корректирующий коэффициент walkingCaloriesCoefficient. Соответствующая константа объявлена в пакете. Вернуть полученное значение.

	durationInMinutes := duration.Minutes()
	calories := (weight * speed * durationInMinutes) / minInH * walkingCaloriesCoefficient

	return calories, nil
}
