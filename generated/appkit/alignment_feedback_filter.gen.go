
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AlignmentFeedbackFilter] class.
var AlignmentFeedbackFilterClass _AlignmentFeedbackFilterClass

func init() {
	AlignmentFeedbackFilterClass = _AlignmentFeedbackFilterClass{objc.GetClass("NSAlignmentFeedbackFilter")}
}

type _AlignmentFeedbackFilterClass struct {
	objc.Class
}

// An interface definition for the [AlignmentFeedbackFilter] class.
type IAlignmentFeedbackFilter interface {
	ID() objc.ID
}

type AlignmentFeedbackFilter struct {
	id objc.ID
}

func AlignmentFeedbackFilterFrom(ptr unsafe.Pointer) AlignmentFeedbackFilter {
	return AlignmentFeedbackFilter{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ AlignmentFeedbackFilter) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _AlignmentFeedbackFilterClass) Alloc() AlignmentFeedbackFilter {
	rv := objc.Send[AlignmentFeedbackFilter](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _AlignmentFeedbackFilterClass) New() AlignmentFeedbackFilter {
	rv := objc.Send[AlignmentFeedbackFilter](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewAlignmentFeedbackFilter creates and returns a new initialized instance.
func NewAlignmentFeedbackFilter() AlignmentFeedbackFilter {
	return AlignmentFeedbackFilterClass.New()
}

// Init initializes the instance.
func (a_ AlignmentFeedbackFilter) Init() AlignmentFeedbackFilter {
	rv := objc.Send[AlignmentFeedbackFilter](a_.ID(), selInit)
	return rv
}
