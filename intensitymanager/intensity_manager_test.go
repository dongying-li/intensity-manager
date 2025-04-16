package intensitymanager

import (
	"fmt"
	"reflect"
	"testing"
)

type action string

const (
	add = action("add")
	set = action("set")
)

type input struct {
	name   action
	from   int
	to     int
	amount int
}

func TestIntensityManager(t *testing.T) {

	tests := []struct {
		name     string
		inputs   []input
		expected [][]Segment
	}{
		{
			name: "Add1",
			inputs: []input{
				{add, 10, 30, 1},
				{add, 20, 40, 1},
				{add, 10, 40, -2}},
			expected: [][]Segment{
				{{10, 1}, {30, 0}},
				{{10, 1}, {20, 2}, {30, 1}, {40, 0}},
				{{10, -1}, {20, 0}, {30, -1}, {40, 0}},
			},
		},
		{
			name: "Add2",
			inputs: []input{
				{add, 10, 30, 1},
				{add, 20, 40, 1},
				{add, 10, 40, -1},
				{add, 10, 40, -1}},
			expected: [][]Segment{
				{{10, 1}, {30, 0}},
				{{10, 1}, {20, 2}, {30, 1}, {40, 0}},
				{{20, 1}, {30, 0}},
				{{10, -1}, {20, 0}, {30, -1}, {40, 0}},
			},
		},
		{
			name: "Add Empty",
			inputs: []input{
				{add, 0, 1, 0},
			},
			expected: [][]Segment{
				{},
			},
		},
		{
			name: "Set Empty",
			inputs: []input{
				{set, 0, 1, 0},
			},
			expected: [][]Segment{
				{},
			},
		},
		{
			name: "Set 1",
			inputs: []input{
				{set, 10, 30, 1},
				{set, 20, 40, 2},
				{set, 10, 40, -1},
				{set, 0, 9, 2},
				{set, 40, 50, 1},
				{set, 3, 12, -3},
			},
			expected: [][]Segment{
				{{10, 1}, {30, 0}},
				{{10, 1}, {20, 2}, {40, 0}},
				{{10, -1}, {40, 0}},
				{{0, 2}, {9, 0}, {10, -1}, {40, 0}},
				{{0, 2}, {9, 0}, {10, -1}, {40, 1}, {50, 0}},
				{{0, 2}, {3, -3}, {12, -1}, {40, 1}, {50, 0}},
			},
		},
		{
			name: "Mixed",
			inputs: []input{
				{add, 10, 20, 2},
				{set, 10, 30, 1},
				{add, 5, 12, -1},
				{set, 20, 40, 2},
				{set, 10, 40, -1},
				{set, 0, 9, 2},
				{set, 40, 50, 1},
				{set, 3, 12, -3},
			},
			expected: [][]Segment{
				{{10, 2}, {20, 0}},
				{{10, 1}, {30, 0}},
				{{5, -1}, {10, 0}, {12, 1}, {30, 0}},
				{{5, -1}, {10, 0}, {12, 1}, {20, 2}, {40, 0}},
				{{5, -1}, {40, 0}},
				{{0, 2}, {9, -1}, {40, 0}},
				{{0, 2}, {9, -1}, {40, 1}, {50, 0}},
				{{0, 2}, {3, -3}, {12, -1}, {40, 1}, {50, 0}},
			},
		},
		{
			name: "Mixed 2",
			inputs: []input{
				{add, 10, 20, 1},
				{set, 15, 17, 2},
			},
			expected: [][]Segment{
				{{10, 1}, {20, 0}},
				{{10, 1}, {15, 2}, {17, 1}, {20, 0}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			im := &IntensityManager{[]Segment{}}
			for i, input := range test.inputs {
				assertEqualValue(t, input, im, test.expected[i])
			}
		})
	}
}

func ExampleIntensityManager_Add() {
	im := &IntensityManager{Segments: []Segment{}}
	im.Add(10, 20, 1)
	fmt.Println(im.Segments)
	// Output: [{10 1} {20 0}]
}

func ExampleIntensityManager_Set() {
	im := &IntensityManager{Segments: []Segment{}}
	im.Add(10, 20, 1)
	im.Set(15, 17, 2)
	fmt.Println(im.Segments)
	// Output: [{10 1} {15 2} {17 1} {20 0}]
}

func assertEqualValue(t *testing.T, it input, im *IntensityManager, expected []Segment) {
	t.Helper()
	switch it.name {
	case add:
		err := im.Add(it.from, it.to, it.amount)
		if err != nil {
			t.Errorf("Add Error - %s\n", err)
		}
		if !reflect.DeepEqual(im.Segments, expected) {
			t.Errorf("\nexpect: %v\nactual: %v\n", expected, im.Segments)
		}
	case set:
		err := im.Set(it.from, it.to, it.amount)
		if err != nil {
			t.Errorf("Set Error - %s\n", err)
		}
		fmt.Println(im.Segments)
		if !reflect.DeepEqual(im.Segments, expected) {
			t.Errorf("\nexpect: %v\nactual: %v\n", expected, im.Segments)
		}
	default:
		t.Errorf("assertEqualValue Error - Unknown action")
	}
}
