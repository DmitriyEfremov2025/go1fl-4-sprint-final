package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
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
	// Алгоритм реализации функции:
	// 1. Получить данные о количестве шагов и продолжительности прогулки с помощью функции parsePackage(). В случае возникновения ошибки вывести её на экран и вернуть пустую строку.
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println("error: ", err)
		return ""
	}

	// 2. Проверить, чтобы количество шагов было больше 0. В противном случае вернуть пустую строку.

	if steps <= 0 {
		log.Println("error: steps must be greater than zero")
		return ""
	}

	// 3. Вычислить дистанцию в метрах. Дистанция равна произведению количества шагов на длину шага. Константа stepLength (длина шага) уже определена в коде.

	distanceMeter := float64(steps) * stepLength

	// 4. Перевести дистанцию в километры, разделив её на число метров в километре (константа mInKm, определена в пакете).
	distanceKm := distanceMeter / mInKm

	// 5. Вычислить количество калорий, потраченных на прогулке. Функция для вычисления калорий WalkingSpentCalories() будет определена в пакете spentcalories, которую вы тоже реализуете.

	if weight <= 0 || height <= 0 {
		log.Println("error: weight and height must be greater than zero")
		return ""
	}

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println("error: ", err)
		return ""
	}

	// 6. Сформировать строку, которую будете возвращать, пример которой был представлен выше.

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, calories)
}
