package math2

import "testing"

type averageTestPair struct {
	values  []float64
	average float64
}

var averageTests = []averageTestPair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
	{[]float64{}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range averageTests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

type maxTestPair struct {
	values []float64
	max    float64
}

var maxTests = []maxTestPair{
	{[]float64{1, 2}, 2},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 1},
}

func TestMax(t *testing.T) {
	for _, pair := range maxTests {
		v := Max(pair.values)
		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}

type minTestPair struct {
	values []float64
	min    float64
}

var minTests = []minTestPair{
	{[]float64{1, 2}, 1},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, -1},
}

func TestMin(t *testing.T) {
	for _, pair := range minTests {
		v := Min(pair.values)
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", v,
			)
		}
	}
}
