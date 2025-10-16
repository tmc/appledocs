
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [showsAdditionalContentIndicators] class.
var showsAdditionalContentIndicatorsClass _showsAdditionalContentIndicatorsClass

func init() {
	showsAdditionalContentIndicatorsClass = _showsAdditionalContentIndicatorsClass{objc.GetClass("showsAdditionalContentIndicators")}
}

type _showsAdditionalContentIndicatorsClass struct {
	objc.Class
}

// An interface definition for the [showsAdditionalContentIndicators] class.
type IshowsAdditionalContentIndicators interface {
	ID() objc.ID
}

type showsAdditionalContentIndicators struct {
	id objc.ID
}

func showsAdditionalContentIndicatorsFrom(ptr unsafe.Pointer) showsAdditionalContentIndicators {
	return showsAdditionalContentIndicators{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ showsAdditionalContentIndicators) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _showsAdditionalContentIndicatorsClass) Alloc() showsAdditionalContentIndicators {
	rv := objc.Send[showsAdditionalContentIndicators](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _showsAdditionalContentIndicatorsClass) New() showsAdditionalContentIndicators {
	rv := objc.Send[showsAdditionalContentIndicators](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewshowsAdditionalContentIndicators creates and returns a new initialized instance.
func NewshowsAdditionalContentIndicators() showsAdditionalContentIndicators {
	return showsAdditionalContentIndicatorsClass.New()
}

// Init initializes the instance.
func (s_ showsAdditionalContentIndicators) Init() showsAdditionalContentIndicators {
	rv := objc.Send[showsAdditionalContentIndicators](s_.ID(), selInit)
	return rv
}
