package controller

import (
	"github.com/camptocamp/varnish-operator/pkg/controller/varnish"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, varnish.Add)
}
