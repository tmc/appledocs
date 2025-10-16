
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [prefersCompactControlSizeMetrics] class.
var prefersCompactControlSizeMetricsClass _prefersCompactControlSizeMetricsClass

func init() {
	prefersCompactControlSizeMetricsClass = _prefersCompactControlSizeMetricsClass{objc.GetClass("prefersCompactControlSizeMetrics")}
}

type _prefersCompactControlSizeMetricsClass struct {
	objc.Class
}

// An interface definition for the [prefersCompactControlSizeMetrics] class.
type IprefersCompactControlSizeMetrics interface {
	ID() objc.ID
}

type prefersCompactControlSizeMetrics struct {
	id objc.ID
}

func prefersCompactControlSizeMetricsFrom(ptr unsafe.Pointer) prefersCompactControlSizeMetrics {
	return prefersCompactControlSizeMetrics{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ prefersCompactControlSizeMetrics) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _prefersCompactControlSizeMetricsClass) Alloc() prefersCompactControlSizeMetrics {
	rv := objc.Send[prefersCompactControlSizeMetrics](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _prefersCompactControlSizeMetricsClass) New() prefersCompactControlSizeMetrics {
	rv := objc.Send[prefersCompactControlSizeMetrics](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewprefersCompactControlSizeMetrics creates and returns a new initialized instance.
func NewprefersCompactControlSizeMetrics() prefersCompactControlSizeMetrics {
	return prefersCompactControlSizeMetricsClass.New()
}

// Init initializes the instance.
func (p_ prefersCompactControlSizeMetrics) Init() prefersCompactControlSizeMetrics {
	rv := objc.Send[prefersCompactControlSizeMetrics](p_.ID(), selInit)
	return rv
}
