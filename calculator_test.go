package calc_test

import (
	"fmt"
	"github.com/nickwallen/quick-calc"
	"github.com/stretchr/testify/assert"
	"testing"
)

var expressions = map[string]string{
	"2 oz":                                "2.00 oz",
	"45 lbs":                              "45.00 lbs",
	"2 kilograms in kg":                   "2.00 kg",
	"2 kg + 2000g":                        "4.00 kg",
	"2 kilograms + 2 kilograms":           "4.00 kilograms",
	"2 pounds + 2 kilograms":              "6.41 pounds",
	"2 feet - 2 feet":                     "0.00 feet",
	"2 meters - 2 feet":                   "1.39 meters",
	"2 pounds in ounces":                  "32.00 ounces",
	"2 pounds + 2 kilograms in kilograms": "2.91 kilograms",
	"2 meters - 2 feet in feet":           "4.56 feet",
	"2kg + 34g in grams":                  "2034.00 grams",
	"2 miles + 2 meters in feet":          "10566.56 feet",
	"12 years in days":                    "4383.00 days",
	"12 mmmH2O + 12 mmmH2O":               "24.00 mmmH2O",
	"2 oz + 3 oz + 4 oz + 5 oz":           "14.00 oz",
}

func TestCalculate(t *testing.T) {
	for expression, expected := range expressions {
		actual, err := calc.Calculate(expression)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	}
}

func TestCalculateAmount(t *testing.T) {
	for expression, expected := range expressions {
		actual, err := calc.CalculateAmount(expression)
		actualStr := fmt.Sprintf("%.2f %s", actual.Value, actual.Units)
		assert.Nil(t, err)
		assert.Equal(t, expected, actualStr)
	}
}

var badExpressions = map[string]string{
	"2 miles + 3 pounds": "failed to resolve conversion: no path found",
	"32 googles":         "at position 4, unit \"google\" not found",
	"2 miles / 500 feet": "at position 9, unexpected input '/'",
}

func TestCalculateBadExpr(t *testing.T) {
	for expr, expectedErr := range badExpressions {
		_, err := calc.Calculate(expr)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, fmt.Sprintf("%s", err))
	}
}
