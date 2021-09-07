package controllers

import (
	"fmt"
	"net/http"
)

func (c Controller) Protected(db) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("success.")
	}
}
