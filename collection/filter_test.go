package collection

import "testing"

func TestFilter(t *testing.T) {
	input := []int{-7, 6, 2, -4, 7, 3}
	expected := []int{6, 2, 7, 3}
	testFilter(t, input, expected)
}

func TestFilterEmpty(t *testing.T) {
	input := make([]int, 0)
	expected := make([]int, 0)
	testFilter(t, input, expected)
}

func testFilter(t *testing.T, input []int, expected []int) {
	output := Filter(input, func(_ int, elem int) bool { return elem >= 2 })

	if len(output) != len(expected) {
		t.Fatalf("expected %+v but got %+v", expected, output)
	}

	for i, elem := range output {
		if expected[i] != elem {
			t.Fatalf("expected %+v but got %+v", expected, output)
		}
	}
}

func TestFilterTo(t *testing.T) {
	input := []int{-7, 6, 2, -4, 7, 3}
	output := make([]int, len(input)+5)
	expected := []int{6, 2, 7, 3}
	testFilterTo(t, input, &output, expected)
}

func TestFilterToEmpty(t *testing.T) {
	input := make([]int, 0)
	output := make([]int, 5)
	expected := make([]int, 0)
	testFilterTo(t, input, &output, expected)
}

func testFilterTo(t *testing.T, input []int, output *[]int, expected []int) {
	FilterTo(input, output, func(_ int, elem int) bool { return elem >= 2 })

	if len(*output) != len(expected) {
		t.Fatalf("expected %+v but got %+v", expected, *output)
	}

	for i, elem := range *output {
		if expected[i] != elem {
			t.Fatalf("expected %+v but got %+v", expected, *output)
		}
	}
}

func TestFilterMut(t *testing.T) {
	input := []int{-7, 6, 2, -4, 7, 3}
	expected := []int{6, 2, 7, 3}
	testFilterMut(t, &input, expected)
}

func TestFilterMutEmpty(t *testing.T) {
	input := make([]int, 0)
	expected := make([]int, 0)
	testFilterMut(t, &input, expected)
}

func testFilterMut(t *testing.T, input *[]int, expected []int) {
	FilterMut(input, func(_ int, elem int) bool { return elem >= 2 })

	if len(*input) != len(expected) {
		t.Fatalf("expected %+v but got %+v", expected, *input)
	}

	for i, elem := range *input {
		if expected[i] != elem {
			t.Fatalf("expected %+v but got %+v", expected, *input)
		}
	}
}
