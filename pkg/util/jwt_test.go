package util

import (
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	l := make([]string, 0)
	m := make(map[string]string)
	l = append(l, "test")
	for _, v := range l {
		s, err := Encode("test")
		if err != nil {
			t.Fatalf("Encode err : %v", err)
		}
		m[v] = s
	}

	for k, v := range m {
		name, err := Decode(v)
		if err != nil {
			t.Fatalf("Decode err : %v", err)
		}
		if !strings.EqualFold(k, name.Name) {
			t.Fatalf("Decode err : %v", err)
		}
		t.Logf("Key : %v , token : %s", k, v)
	}
}
