// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SObject] class.
var (
	SObjectClass     _SObjectClass
	SObjectClassOnce sync.Once
)

func getSObjectClass() _SObjectClass {
	SObjectClassOnce.Do(func() {
		SObjectClass = _SObjectClass{objc.GetClass("CLSObject")}
	})
	return SObjectClass
}

type _SObjectClass struct {
	class objc.Class
}

// An interface definition for the [SObject] class.
type ISObject interface {
	objectivec.IObject
	// properties:
	DateCreated() foundation.objc.IObject /* cross-framework: Date */
	SetDateCreated(value foundation.objc.IObject /* cross-framework: Date */)
	DateLastModified() foundation.objc.IObject /* cross-framework: Date */
	SetDateLastModified(value foundation.objc.IObject /* cross-framework: Date */)
	// methods:
}

// The abstract base class for objects managed by ClassKit.


// The abstract base class for objects managed by ClassKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSObject
type SObject struct {
	objectivec.Object
}

// SObjectFrom constructs a [SObject] from an unsafe.Pointer.
//
// The abstract base class for objects managed by ClassKit.
func SObjectFrom(ptr unsafe.Pointer) SObject {
	return SObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SObjectClass) Alloc() SObject {
	rv := objc.Send[SObject](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SObjectClass) New() SObject {
	rv := objc.Send[SObject](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SObject) Init() SObject {
	rv := objc.Send[SObject](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SObject) Autorelease() SObject {
	rv := objc.Send[SObject](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSObject creates a new SObject instance.
func NewSObject() SObject {
	return getSObjectClass().New()
}



// The date on which the object was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clsobject/datecreated
func (s_ SObject) DateCreated() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](s_.ID, objc.Sel("dateCreated"))
	return rv
}


// The date on which the object was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clsobject/datecreated
func (s_ SObject) SetDateCreated(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDateCreated:"), value)
}


// The date on which the object was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clsobject/datelastmodified
func (s_ SObject) DateLastModified() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](s_.ID, objc.Sel("dateLastModified"))
	return rv
}


// The date on which the object was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clsobject/datelastmodified
func (s_ SObject) SetDateLastModified(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDateLastModified:"), value)
}



