package services

import (
	"main/models"
	"main/utils"
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
	q1 := ages[len(ages)/4]
	q3 := ages[3*len(ages)/4]
	iqr := q3 - q1

	iqrMultiplier := 2.0
	upperBound := q3 + int(iqrMultiplier*float64(iqr))

	var outliers []models.Machine
	for _, m := range machines {
		age, _ := utils.ParseToDays(m.Age)
		if age > upperBound {
			outliers = append(outliers, m)
		}
	}
	return outliers, nil
}
