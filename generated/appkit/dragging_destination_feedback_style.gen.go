
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [draggingDestinationFeedbackStyle] class.
var draggingDestinationFeedbackStyleClass _draggingDestinationFeedbackStyleClass

func init() {
	draggingDestinationFeedbackStyleClass = _draggingDestinationFeedbackStyleClass{objc.GetClass("draggingDestinationFeedbackStyle")}
}

type _draggingDestinationFeedbackStyleClass struct {
	objc.Class
}

// An interface definition for the [draggingDestinationFeedbackStyle] class.
type IdraggingDestinationFeedbackStyle interface {
	ID() objc.ID
}

type draggingDestinationFeedbackStyle struct {
	id objc.ID
}

func draggingDestinationFeedbackStyleFrom(ptr unsafe.Pointer) draggingDestinationFeedbackStyle {
	return draggingDestinationFeedbackStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ draggingDestinationFeedbackStyle) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _draggingDestinationFeedbackStyleClass) Alloc() draggingDestinationFeedbackStyle {
	rv := objc.Send[draggingDestinationFeedbackStyle](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _draggingDestinationFeedbackStyleClass) New() draggingDestinationFeedbackStyle {
	rv := objc.Send[draggingDestinationFeedbackStyle](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdraggingDestinationFeedbackStyle creates and returns a new initialized instance.
func NewdraggingDestinationFeedbackStyle() draggingDestinationFeedbackStyle {
	return draggingDestinationFeedbackStyleClass.New()
}

// Init initializes the instance.
func (d_ draggingDestinationFeedbackStyle) Init() draggingDestinationFeedbackStyle {
	rv := objc.Send[draggingDestinationFeedbackStyle](d_.ID(), selInit)
	return rv
}
