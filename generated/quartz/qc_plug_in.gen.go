// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [QCPlugIn] class.
var (
	QCPlugInClass     _QCPlugInClass
	QCPlugInClassOnce sync.Once
)

func getQCPlugInClass() _QCPlugInClass {
	QCPlugInClassOnce.Do(func() {
		QCPlugInClass = _QCPlugInClass{objc.GetClass("QCPlugIn")}
	})
	return QCPlugInClass
}

type _QCPlugInClass struct {
	class objc.Class
}

// An interface definition for the [QCPlugIn] class.
type IQCPlugIn interface {
	objectivec.IObject
	ExecuteAtTimeWithArguments(context objectivec.IObject, time foundation.ITimeInterval, arguments objectivec.IObject) bool
}

// A base class to subclass for writing custom patches.
//
// The class provides the base class to subclass for writing custom Quartz Composer patches. You implement a custom patch by subclassing , overriding the appropriate methods, packaging the code as an object, and installing the bundle in the appropriate location. A bundle can contain more than one subclass of , allowing you to provide a suite of custom patches in one bundle. provides detailed instructions on how to create and package a custom patch. supplements the information in the programming guide. The methods related to the executing the custom patch (called when the Quartz Composer engine is rendering) are passed an opaque object that conforms to the protocol. This object represents the execution context of the object. You should not retain the execution context or use it outside of the scope of the execution method that it is passed to.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugIn
type QCPlugIn struct {
	objectivec.Object
}

// QCPlugInFrom constructs a [QCPlugIn] from an unsafe.Pointer.
//
// A base class to subclass for writing custom patches.
func QCPlugInFrom(ptr unsafe.Pointer) QCPlugIn {
	return QCPlugIn{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (qc _QCPlugInClass) Alloc() QCPlugIn {
	rv := objc.Send[QCPlugIn](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QCPlugInClass) New() QCPlugIn {
	rv := objc.Send[QCPlugIn](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QCPlugIn) Init() QCPlugIn {
	rv := objc.Send[QCPlugIn](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QCPlugIn) Autorelease() QCPlugIn {
	rv := objc.Send[QCPlugIn](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQCPlugIn creates a new QCPlugIn instance.
func NewQCPlugIn() QCPlugIn {
	return getQCPlugInClass().New()
}


// Performs the processing or rendering tasks appropriate for the custom patch.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugIn/execute(_:atTime:withArguments:)
func (q_ QCPlugIn) ExecuteAtTimeWithArguments(context objectivec.IObject, time foundation.ITimeInterval, arguments objectivec.IObject) bool {
	rv := objc.Send[bool](q_.ID, objc.Sel("execute:atTime:withArguments:"), context, time, arguments)
	return rv
}



