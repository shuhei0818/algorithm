package list_test

import (
	"testing"

	"github.com/shuhei0818/algorithm/list"
)

func TestSuccess(t *testing.T) {

	cases := map[string]struct {
		in  []string
		out []string
	}{
		"string": {in: []string{"1", "2", "3"}, out: []string{"3", "2", "1"}},
	}

	stack := list.NewStack[string](10)
	for _, c := range cases {
		for _, v := range c.in {
			stack.Add(v)
		}
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			for _, v := range c.out {
				r, err := stack.Get()

				if err != nil {
					t.Error("unexpected error")
				}

				if v != r {
					t.Errorf("not equal, expected = %v, given = %v", v, r)
				}
			}
		})
	}
}
