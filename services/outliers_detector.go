package services

import (
	"machineagevalidator/models"
	"machineagevalidator/utils"
	"sort"
)

func DetectOutliers(machines []models.Machine) ([]models.Machine, error) {
	ages := make([]int, 0, len(machines))
	for _, m := range machines {
		age, err := utils.ParseToDays(m.Age)
		if err != nil {
			return nil, err
		}
		ages = append(ages, age)
	}

	sort.Ints(ages)

	q1, q3 := calculateQuartiles(ages)
	iqr := q3 - q1

	lowerBound := q1 - int(1.5*float64(iqr))
	upperBound := q3 + int(1.5*float64(iqr))

	var outliers []models.Machine
	for _, m := range machines {
		age, _ := utils.ParseToDays(m.Age)
		if age < lowerBound || age > upperBound {
			outliers = append(outliers, m)
		}
	}

	return outliers, nil
}

func calculateQuartiles(data []int) (int, int) {
	n := len(data)
	mid := n / 2

	var q1, q3 int
	if n%2 == 0 {
		q1 = median(data[:mid])
		q3 = median(data[mid:])
	} else {
		q1 = median(data[:mid])
		q3 = median(data[mid+1:])
	}

	return q1, q3
}

func median(data []int) int {
	n := len(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2
	}
	return data[n/2]
}
