// Package objc provides cached Objective-C runtime helpers.
//
// This package wraps purego/objc to provide selector caching for better performance.
package objc

import (
	"sync"

	purego "github.com/ebitengine/purego/objc"
)

// Type aliases for convenience
type (
	ID    = purego.ID
	SEL   = purego.SEL
	Class = purego.Class
)

var (
	selCache sync.Map // map[string]purego.SEL
)

// Sel returns a cached selector for the given name.
// This avoids the global lock in purego.RegisterName on repeated calls.
func Sel(name string) SEL {
	if sel, ok := selCache.Load(name); ok {
		return sel.(SEL)
	}
	sel := purego.RegisterName(name)
	selCache.Store(name, sel)
	return sel
}

// Send calls purego.Send with the given arguments.
func Send[T any](id ID, sel SEL, args ...any) T {
	return purego.Send[T](id, sel, args...)
}

// GetClass returns the class with the given name.
func GetClass(name string) Class {
	return purego.GetClass(name)
}

