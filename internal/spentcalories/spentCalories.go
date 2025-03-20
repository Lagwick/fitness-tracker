package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep = 0.65 // средняя длина шага.
	mInKm   = 1000 // количество метров в километре.
	minInH  = 60   // количество минут в часе.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// ваш код ниже
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, errors.New("Incorrect string format")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("convertation error: %v", err)
	}
	activity := parts[1]
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("parsing error: %v", err)
	}
	return steps, activity, duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	return float64(steps) * lenStep / float64(mInKm)

}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	// ваш код ниже
	if duration.Seconds() <= 0 {
		return 0
	}
	distanceCovered := distance(steps)
	hours := duration.Hours()
	speed := distanceCovered / hours
	return speed
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	// ваш код ниже
	stepsDone, trainingType, timeOftraining, err := parseTraining(data)
	if err != nil {
		return fmt.Sprintf("Parsing error: %v", err)
	}
	if stepsDone <= 0 {
		return fmt.Sprintf("not enough steps: %v", err)
	}
	distanceCovered := distance(stepsDone)
	avSpeed := meanSpeed(stepsDone, timeOftraining)
	var cals float64
	switch trainingType {
	case "Бег":
		cals = RunningSpentCalories(stepsDone, weight, timeOftraining)
	case "Ходьба":
		cals = WalkingSpentCalories(stepsDone, weight, height, timeOftraining)
	default:
		return "Unknown type of activity"
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч\nДистанция: %.2f км\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		trainingType, timeOftraining.Hours(), distanceCovered, avSpeed, cals)
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	runSpeed := meanSpeed(steps, duration)
	calsBurnedofRun := ((runningCaloriesMeanSpeedMultiplier * runSpeed) - runningCaloriesMeanSpeedShift) * weight
	return calsBurnedofRun
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	walkSpeed := meanSpeed(steps, duration)
	hours := duration.Hours()
	calsBurnedOfWalk := ((walkingCaloriesWeightMultiplier * weight) + (walkSpeed*walkSpeed/height)*walkingSpeedHeightMultiplier) * hours * minInH

	return calsBurnedOfWalk
}
