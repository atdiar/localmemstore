package localmemstore

import (
	"testing"
	"time"
)

func TestStorage(t *testing.T) {
	// Let's declare s a storage datastructure.
	s := New().DefaultExpiry(9 * time.Millisecond)

	// Let's put some elements in it
	err := s.Put("user1", "key1", []byte("some stuff"), 98789)
	if err != nil {
		t.Log(err)
	}

	err = s.Put("user1", "key2", []byte("some stuff again"), 90789756)
	if err != nil {
		t.Log(err)
	}

	err = s.Put("user2", "key2", []byte("some foo"), 9787)
	if err != nil {
		t.Log(err)
	}

	// Let's see if the first item was correctly put and can be retrieved.
	v, err := s.Get("user1", "key1")
	if err != nil {
		t.Log(err)
	}
	if string(v) != "some stuff" {
		t.Errorf("Expected %v but got %v \n", "some stuff", v)
	}

	// Let's see if the second item was correctly put and can be retrieved.
	v, err = s.Get("user1", "key2")
	if err != nil {
		t.Log(err)
	}
	if string(v) != "some stuff again" {
		t.Errorf("Expected %v but got %v \n", "some stuff again", string(v))
	}

	// Let's delete user1's value stored under key2 and see if it is done correctly
	err = s.Delete("user1", "key1")
	if err != nil {
		t.Log(err)
	}
	v, err = s.Get("user1", "key1")
	if v != nil {
		t.Errorf("Was not expecting any value but got %v", v)
	}
	if err == nil {
		t.Error("Was expecting a 'Not found' error.")
	}

	s.SetExpiry("user2", 0)

	time.Sleep(50 * time.Millisecond)

	// Now if we attempt to retrieve any value, given the settings, we should not
	// be able to since all the values are supposed to be expired.

	v, err = s.Get("user1", "key1")
	if v != nil {
		t.Errorf("Was not expecting any value but got %v", v)
	}
	if err == nil {
		t.Error("Was expecting a 'Not found' error.")
	}

	v, err = s.Get("user1", "key2")
	if v != nil {
		t.Errorf("Was not expecting any value but got %v", v)
	}
	if err == nil {
		t.Error("Was expecting a 'Not found' error.")
	}

	v, err = s.Get("user2", "key2")
	if v != nil {
		t.Errorf("Was not expecting any value but got %v", v)
	}
	if err == nil {
		t.Error("Was expecting a 'Not found' error.")
	}
}

func TestClear(t *testing.T) {
	s := New().DefaultExpiry(9 * time.Millisecond)

	// Let's test the clear() method
	// Let's put some elements in it
	err := s.Put("user1", "key1", []byte("some stuff"), 9877)
	if err != nil {
		t.Log(err)
	}

	err = s.Put("user1", "key2", []byte("some stuff again"), 5768586)
	if err != nil {
		t.Log(err)
	}

	err = s.Put("user2", "key2", []byte("some foo"), 65474568)
	if err != nil {
		t.Log(err)
	}
	// Applying the clear method
	func(st Store) {
		st.Clear()
	}(s)

	// Now if we attempt to retrieve any value, given the settings, we should not
	// be able to since all the values are supposed to have been deleted.

	v, err := s.Get("user1", "key1")
	if v != nil {
		t.Errorf("Was not expecting any value but got %v", v)
	}
	if err == nil {
		t.Error("Was expecting a 'Not found' error.")
	}

	v, err = s.Get("user1", "key2")
	if v != nil {
		t.Errorf("Was not expecting any value but got %v", v)
	}
	if err == nil {
		t.Error("Was expecting a 'Not found' error.")
	}

	v, err = s.Get("user2", "key2")
	if v != nil {
		t.Errorf("Was not expecting any value but got %v", v)
	}
	if err == nil {
		t.Error("Was expecting a 'Not found' error.")
	}
}
