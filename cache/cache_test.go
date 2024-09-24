package cache

import (
	"testing"
)

func TestCache(t *testing.T) {
	Init()
	Tests := [][]string{
		[]string{"a", "1"},
		[]string{"b", "2"},
		[]string{"c", "3"},
		[]string{"d", "4"},
		[]string{"e", "5"},
	}
	for i := range Tests {
		Add(Tests[i][0],Tests[i][1])
	}
	for i:= range Tests {
		v,found:=Get(Tests[i][0])
		if !found || v!=Tests[i][1]{
			t.Errorf("error on TC#%d",i)
		}
	}
}
