package password

import (
	"errors"
	"testing"
)

func TestHash(t *testing.T) {
	hash, err := Hash("admin123")
	if err != nil {
		t.Error(err)
	}
	comparator := NewComparator()
	ok, err := comparator.Compare("admin123", hash)
	if err != nil {
		t.Error(err)
	}
	if ok == false {
		t.Error(errors.New("same passwords verified as not equal"))
	}
	// negative case
	hash, err= Hash("not_admin123")
	if err != nil {
		t.Error(err)
	}
	ok, err = comparator.Compare("admin123", hash)
	if ok == true {
		t.Error(errors.New("different passwords verified as equal"))
	}
}
