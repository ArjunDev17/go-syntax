package main

import "testing"

func TestCreateUser(t *testing.T) {
	store := &InmemoryStore{
		Users: make(map[string]User),
	}

}
