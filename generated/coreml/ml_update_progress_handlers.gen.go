// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UpdateProgressHandlers] class.
var (
	UpdateProgressHandlersClass     _UpdateProgressHandlersClass
	UpdateProgressHandlersClassOnce sync.Once
)

func getUpdateProgressHandlersClass() _UpdateProgressHandlersClass {
	UpdateProgressHandlersClassOnce.Do(func() {
		UpdateProgressHandlersClass = _UpdateProgressHandlersClass{objc.GetClass("MLUpdateProgressHandlers")}
	})
	return UpdateProgressHandlersClass
}

type _UpdateProgressHandlersClass struct {
	class objc.Class
}

// An interface definition for the [UpdateProgressHandlers] class.
type IUpdateProgressHandlers interface {
	objectivec.IObject
}

// A collection of closures an update task uses to notify your app of its progress.


// A collection of closures an update task uses to notify your app of its progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers

type UpdateProgressHandlers struct {
	objectivec.Object
}

// UpdateProgressHandlersFrom constructs a [UpdateProgressHandlers] from an unsafe.Pointer.
//
// A collection of closures an update task uses to notify your app of its progress.
func UpdateProgressHandlersFrom(ptr unsafe.Pointer) UpdateProgressHandlers {
	return UpdateProgressHandlers{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UpdateProgressHandlersClass) Alloc() UpdateProgressHandlers {
	rv := objc.Send[UpdateProgressHandlers](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UpdateProgressHandlersClass) New() UpdateProgressHandlers {
	rv := objc.Send[UpdateProgressHandlers](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UpdateProgressHandlers) Init() UpdateProgressHandlers {
	rv := objc.Send[UpdateProgressHandlers](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UpdateProgressHandlers) Autorelease() UpdateProgressHandlers {
	rv := objc.Send[UpdateProgressHandlers](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUpdateProgressHandlers creates a new UpdateProgressHandlers instance.
func NewUpdateProgressHandlers() UpdateProgressHandlers {
	return getUpdateProgressHandlersClass().New()
}




// Creates the collection of closures an update task uses to notify your app of its progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers/init(forEvents:progressHandler:completionHandler:)

func NewUpdateProgressHandlersForEventsProgressHandlerCompletionHandler(interestedEvents IUpdateProgressEvent, progressHandler unsafe.Pointer, completionHandler unsafe.Pointer) UpdateProgressHandlers {
	instance := getUpdateProgressHandlersClass().Alloc()
	rv := objc.Send[UpdateProgressHandlers](instance.ID, objc.Sel("initForEvents:progressHandler:completionHandler:"), interestedEvents, progressHandler, completionHandler)
	rv.Autorelease()
	return rv
}



