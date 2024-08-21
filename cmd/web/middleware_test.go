package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myHand myHandler
	h := NoSurf(&myHand)

	switch v := h.(type) {

	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.handler, but is of type %T", v))

	}

}

func TestSessionLoad(t *testing.T) {
	var myHand myHandler
	h := SessionLoad(&myHand)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.handler, but is of type %T", v))

	}

}
