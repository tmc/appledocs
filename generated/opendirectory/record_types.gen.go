// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [recordTypes] class.
var (
	RecordTypesClass     _recordTypesClass
	RecordTypesClassOnce sync.Once
)

func getrecordTypesClass() _recordTypesClass {
	RecordTypesClassOnce.Do(func() {
		RecordTypesClass = _recordTypesClass{objc.GetClass("recordTypes")}
	})
	return RecordTypesClass
}

type _recordTypesClass struct {
	class objc.Class
}

// An interface definition for the [recordTypes] class.
type IrecordTypes interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/recordTypes-c.ivar
type recordTypes struct {
	objectivec.Object
}

// recordTypesFrom constructs a [recordTypes] from an unsafe.Pointer.
func recordTypesFrom(ptr unsafe.Pointer) recordTypes {
	return recordTypes{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _recordTypesClass) Alloc() recordTypes {
	rv := objc.Send[recordTypes](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _recordTypesClass) New() recordTypes {
	rv := objc.Send[recordTypes](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ recordTypes) Init() recordTypes {
	rv := objc.Send[recordTypes](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ recordTypes) Autorelease() recordTypes {
	rv := objc.Send[recordTypes](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewrecordTypes creates a new recordTypes instance.
func NewrecordTypes() recordTypes {
	return getrecordTypesClass().New()
}




