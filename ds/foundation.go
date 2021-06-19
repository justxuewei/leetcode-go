package ds

import "errors"

var NoSpace = errors.New("the length of the queue exceeds the capacity")
var NoElement = errors.New("there is no elements")
