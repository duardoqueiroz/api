package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Filter struct {
	Years    []string
	Months   []string
	Agencies []string
	Category string
	Types    string
}

func NewFilter(yearsQp, monthsQp, agenciesQp, categoriesQp, typesQp string) (*Filter, error) {
	var years []string
	var months []string
	var agencies []string
	var types string

	if yearsQp == "" && monthsQp == "" && agenciesQp == "" && categoriesQp == "" && typesQp == "" {
		return nil, nil
	}
	if yearsQp != "" {
		years = strings.Split(yearsQp, ",")
		for _, y := range years {
			if _, err := strconv.Atoi(y); err != nil {
				return nil, fmt.Errorf("parâmetro ano '%s' é inválido!", y)
			}
		}
	}
	if monthsQp != "" {
		months = strings.Split(monthsQp, ",")
		for _, m := range months {
			if _, err := strconv.Atoi(m); err != nil {
				return nil, fmt.Errorf("parâmetro mês '%s' é inválido!", m)
			}
		}
	}
	if agenciesQp != "" {
		agencies = strings.Split(agenciesQp, ",")
	}

	return &Filter{
		Years:    years,
		Months:   months,
		Agencies: agencies,
		Category: categoriesQp,
		Types:    types,
	}, nil
}
