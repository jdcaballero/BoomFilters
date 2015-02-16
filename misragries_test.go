package boom

import (
	"testing"
)

func TestNewMisraGries(t *testing.T) {
	m1, _ := NewMisraGries(3)
	m1.Add("a").Add("a").Add("a")
	if k := m1.Len(); k != 1 {
		t.Errorf("Expected 1, got %d", k)
	}

}
func TestNoDeleteMisraGries(t *testing.T) {
	m1, _ := NewMisraGries(3)
	m1 = m1.Add("s")
	m1 = m1.Add("s")

	m1 = m1.Add("as")
	m1 = m1.Add("as")
	m1 = m1.Add("as")

	m1 = m1.Add("ass")

	m1 = m1.Add("as")
	m1 = m1.Add("as")

	if k := m1.Counts("as"); k != 5 {
		t.Errorf("Expected 5, got %d", k)
	}

}

func TestDeleteMisraGries(t *testing.T) {
	m1, _ := NewMisraGries(3)
	m1 = m1.Add("s")
	m1 = m1.Add("s")

	m1 = m1.Add("as")
	m1 = m1.Add("as")
	m1 = m1.Add("as")

	m1 = m1.Add("ass")

	m1 = m1.Add("as")
	m1 = m1.Add("as")

	m1 = m1.Add("asx")
	m1 = m1.Add("asx")
	m1 = m1.Add("asx")

	if k := m1.Counts("asx"); k != 2 {
		t.Errorf("Expected 2, got %d", k)
	}

}
func TestNotPresentMisraGries(t *testing.T) {
	m1, _ := NewMisraGries(3)
	m1 = m1.Add("s")
	m1 = m1.Add("s")

	m1 = m1.Add("as")
	m1 = m1.Add("as")
	m1 = m1.Add("as")

	m1 = m1.Add("ass")

	m1 = m1.Add("as")
	m1 = m1.Add("as")

	m1 = m1.Add("asx")
	m1 = m1.Add("asx")
	m1 = m1.Add("asx")

	if k := m1.Counts("NOT_PRESENT"); k != 0 {
		t.Errorf("Expected 0, got %d", k)
	}

}
