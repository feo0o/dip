package ipv4

import "testing"

func TestIPNext(t *testing.T) {
	s := "10.10.255.255"
	i, err := ParseIPv4(s)
	if err != nil {
		t.Error(err)
	}
	t.Log(i)
	t.Log(i.Next())
}

func TestIPPrev(t *testing.T) {
	s := "0.0.0.0"
	i, err := ParseIPv4(s)
	if err != nil {
		t.Error(err)
	}
	t.Log(i)
	t.Log(i.Prev())
}
