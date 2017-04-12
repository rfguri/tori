package crawler

import (
	"reflect"
	"testing"
)

func TestCrawler(t *testing.T) {
	for i, c := range cases {
		got := Get(c.input, 0)
		eq := reflect.DeepEqual(c.expect, got)
		if !eq {
			t.Errorf("Case #%v: Expected: %v, got: %v", i, c.expect, got)
		}
	}
}
