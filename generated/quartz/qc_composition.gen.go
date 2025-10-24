// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [QCComposition] class.
var (
	QCCompositionClass     _QCCompositionClass
	QCCompositionClassOnce sync.Once
)

func getQCCompositionClass() _QCCompositionClass {
	QCCompositionClassOnce.Do(func() {
		QCCompositionClass = _QCCompositionClass{objc.GetClass("QCComposition")}
	})
	return QCCompositionClass
}

type _QCCompositionClass struct {
	class objc.Class
}

// An interface definition for the [QCComposition] class.
type IQCComposition interface {
	objectivec.IObject
	// properties:
	// methods:
}

// The class represents a Quartz Composer composition that either:
//
// comes from the system-wide composition repository ( and ) where it can be accessed by any application through the methods of the class is created from an arbitrary source (typically a file on disk) using one of its methods This class cannot be subclassed. A object has the following information associated with it and that you can obtain by using the appropriate method of the class: Attributes include the name and description of the composition, copyright information, and whether or not its provided by macOS (built-in). The protocols that the composition conforms to. A defines a set of required and optional input parameters and output results. Many methods of the , , and classes take a object as a parameter.


// The class represents a Quartz Composer composition that either:
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCComposition
type QCComposition struct {
	objectivec.Object
}

// QCCompositionFrom constructs a [QCComposition] from an unsafe.Pointer.
//
// The class represents a Quartz Composer composition that either:
func QCCompositionFrom(ptr unsafe.Pointer) QCComposition {
	return QCComposition{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (qc _QCCompositionClass) Alloc() QCComposition {
	rv := objc.Send[QCComposition](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QCCompositionClass) New() QCComposition {
	rv := objc.Send[QCComposition](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QCComposition) Init() QCComposition {
	rv := objc.Send[QCComposition](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QCComposition) Autorelease() QCComposition {
	rv := objc.Send[QCComposition](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQCComposition creates a new QCComposition instance.
func NewQCComposition() QCComposition {
	return getQCCompositionClass().New()
}




