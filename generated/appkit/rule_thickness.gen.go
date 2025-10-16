
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ruleThickness] class.
var ruleThicknessClass _ruleThicknessClass

func init() {
	ruleThicknessClass = _ruleThicknessClass{objc.GetClass("ruleThickness")}
}

type _ruleThicknessClass struct {
	objc.Class
}

// An interface definition for the [ruleThickness] class.
type IruleThickness interface {
	ID() objc.ID
}

type ruleThickness struct {
	id objc.ID
}

func ruleThicknessFrom(ptr unsafe.Pointer) ruleThickness {
	return ruleThickness{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ ruleThickness) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _ruleThicknessClass) Alloc() ruleThickness {
	rv := objc.Send[ruleThickness](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _ruleThicknessClass) New() ruleThickness {
	rv := objc.Send[ruleThickness](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewruleThickness creates and returns a new initialized instance.
func NewruleThickness() ruleThickness {
	return ruleThicknessClass.New()
}

// Init initializes the instance.
func (r_ ruleThickness) Init() ruleThickness {
	rv := objc.Send[ruleThickness](r_.ID(), selInit)
	return rv
}
