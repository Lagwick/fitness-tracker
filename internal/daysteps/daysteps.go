package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
	parsedString := strings.Split(data, ",")
	if len(parsedString) != 2 {
		return 0, 0, errors.New("incorrect string format")
	}
	steps, err := strconv.Atoi(parsedString[0])
	if err != nil {
		return 0, 0, fmt.Errorf("convertation error: %w", err)
	}
	if steps <= 0 {
		return 0, 0, errors.New("steps must be more than zero")
	}

	duration, err := time.ParseDuration(parsedString[1])
	if err != nil {
		return 0, 0, fmt.Errorf("parsing error: %w", err)
	}
	if duration <= 0 {
		return 0, 0, errors.New("duration must be more than zero")
	}
	return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Errorf("Parsing error: %w", err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distance := float64(steps) * StepLength / 1000
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила: %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distance, calories,
	)
}
