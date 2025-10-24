// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEngineStateSerialization */


/* debug [class_header]: Header for CKSyncEngineStateSerialization */
// The class instance for the [CKSyncEngineStateSerialization] class.
var (
	CKSyncEngineStateSerializationClass     _CKSyncEngineStateSerializationClass
	CKSyncEngineStateSerializationClassOnce sync.Once
)

func getCKSyncEngineStateSerializationClass() _CKSyncEngineStateSerializationClass {
	CKSyncEngineStateSerializationClassOnce.Do(func() {
		CKSyncEngineStateSerializationClass = _CKSyncEngineStateSerializationClass{objc.GetClass("CKSyncEngineStateSerialization")}
	})
	return CKSyncEngineStateSerializationClass
}

type _CKSyncEngineStateSerializationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineStateSerialization */
// An interface definition for the [CKSyncEngineStateSerialization] class.
type ICKSyncEngineStateSerialization interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEngineStateSerialization */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineStateSerialization */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineStateSerialization */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineStateSerializationClass) Alloc() CKSyncEngineStateSerialization {
	rv := objc.Send[CKSyncEngineStateSerialization](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineStateSerializationClass) New() CKSyncEngineStateSerialization {
	rv := objc.Send[CKSyncEngineStateSerialization](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineStateSerialization) Init() CKSyncEngineStateSerialization {
	rv := objc.Send[CKSyncEngineStateSerialization](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineStateSerialization) Autorelease() CKSyncEngineStateSerialization {
	rv := objc.Send[CKSyncEngineStateSerialization](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineStateSerialization creates a new CKSyncEngineStateSerialization instance.
func NewCKSyncEngineStateSerialization() CKSyncEngineStateSerialization {
	return getCKSyncEngineStateSerializationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineStateSerialization */
// An opaque object that contains the serialized representation of a sync engine’s current state.


// An opaque object that contains the serialized representation of a sync engine’s current state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineStateSerialization
type CKSyncEngineStateSerialization struct {
	objectivec.Object
}

// CKSyncEngineStateSerializationFrom constructs a [CKSyncEngineStateSerialization] from an unsafe.Pointer.
//
// An opaque object that contains the serialized representation of a sync engine’s current state.
func CKSyncEngineStateSerializationFrom(ptr unsafe.Pointer) CKSyncEngineStateSerialization {
	return CKSyncEngineStateSerialization{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineStateSerialization *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineStateSerialization */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineStateSerialization */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineStateSerialization */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineStateSerialization */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineStateSerialization */



