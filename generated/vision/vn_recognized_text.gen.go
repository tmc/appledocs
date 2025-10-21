// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RecognizedText] class.
var (
	RecognizedTextClass     _RecognizedTextClass
	RecognizedTextClassOnce sync.Once
)

func getRecognizedTextClass() _RecognizedTextClass {
	RecognizedTextClassOnce.Do(func() {
		RecognizedTextClass = _RecognizedTextClass{objc.GetClass("VNRecognizedText")}
	})
	return RecognizedTextClass
}

type _RecognizedTextClass struct {
	class objc.Class
}

// An interface definition for the [RecognizedText] class.
type IRecognizedText interface {
	objectivec.IObject
}

// Text recognized in an image through a text recognition request.
//
// A single can contain multiple recognized text objects—one for each candidate.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedText
type RecognizedText struct {
	objectivec.Object
}

// RecognizedTextFrom constructs a [RecognizedText] from an unsafe.Pointer.
//
// Text recognized in an image through a text recognition request.
func RecognizedTextFrom(ptr unsafe.Pointer) RecognizedText {
	return RecognizedText{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RecognizedTextClass) Alloc() RecognizedText {
	rv := objc.Send[RecognizedText](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RecognizedTextClass) New() RecognizedText {
	rv := objc.Send[RecognizedText](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecognizedText) Init() RecognizedText {
	rv := objc.Send[RecognizedText](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecognizedText) Autorelease() RecognizedText {
	rv := objc.Send[RecognizedText](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecognizedText creates a new RecognizedText instance.
func NewRecognizedText() RecognizedText {
	return getRecognizedTextClass().New()
}




