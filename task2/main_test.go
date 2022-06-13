package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaradal/flaconi-challenge/nesting"
)

func TestCreateLeavesFromLoop(t *testing.T) {
	tt := []struct {
		description string
		// Inputs
		args  []string
		items []map[string]interface{}
		// Outputs
		wantOutput map[string]interface{}
		wantErr    bool
	}{
		{
			description: "empty arguments",
			args:        []string{},
			wantErr:     true,
		},
		{
			description: "one argument that matches a non-string value",
			args:        []string{"field1"},
			items: []map[string]interface{}{
				{"field1": 1},
				{"field2": "foobar"},
			},
			wantErr: true,
		},
		{
			description: "one argument that matches a string value",
			args:        []string{"field1"},
			items: []map[string]interface{}{
				{"field1": "foo"},
				{"field1": "bar"},
			},
			wantOutput: map[string]interface{}{
				"foo": []interface{}{map[string]interface{}{}},
				"bar": []interface{}{map[string]interface{}{}},
			},
		},
		{
			description: "one argument with two fields",
			args:        []string{"field1"},
			items: []map[string]interface{}{
				{"field1": "foo", "field2": "abc"},
				{"field1": "bar", "field2": "def"},
			},
			wantOutput: map[string]interface{}{
				"foo": []interface{}{map[string]interface{}{"field2": "abc"}},
				"bar": []interface{}{map[string]interface{}{"field2": "def"}},
			},
		},
		{
			description: "one argument that has nil value",
			args:        []string{"field1"},
			items: []map[string]interface{}{
				{"field2": "abc"},
				{"field1": "bar", "field2": "def"},
			},
			wantOutput: map[string]interface{}{
				"":    []interface{}{map[string]interface{}{"field2": "abc"}},
				"bar": []interface{}{map[string]interface{}{"field2": "def"}},
			},
		},
		{
			description: "two arguments with three fields",
			args:        []string{"field1", "field2"},
			items: []map[string]interface{}{
				{"field1": "foo", "field2": "abc", "field3": "xyz"},
				{"field1": "bar", "field2": "def", "field3": "klmn"},
			},
			wantOutput: map[string]interface{}{
				"foo": map[string]interface{}{
					"abc": []interface{}{map[string]interface{}{"field3": "xyz"}},
				},
				"bar": map[string]interface{}{
					"def": []interface{}{map[string]interface{}{"field3": "klmn"}},
				}},
		},
	}
	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			svc := nesting.NewService()
			got, err := svc.CreateLeavesFromLoop(tc.args, tc.items)
			if (err != nil) != tc.wantErr {
				t.Fatalf("expected err: %v, got: %v", tc.wantErr, err)
			}
			assert.Equal(t, tc.wantOutput, got)
		})
	}
}
