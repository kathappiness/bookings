package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var mh myHandler
	h := NoSurf(&mh)
	//Checking if the received object is of the type http.Handler
	switch v := h.(type) {
	case http.Handler: // do nothing:)
	default:
		t.Error(fmt.Printf("type is not http.Handler, it's a type %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var mh myHandler
	h := SessionLoad(&mh)
	//Checking if the received object is of the type http.Handler
	switch v := h.(type) {
	case http.Handler: // do nothing:)
	default:
		t.Error(fmt.Printf("type is not http.Handler, it's a type %T", v))
	}
}
