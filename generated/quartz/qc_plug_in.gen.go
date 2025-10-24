// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class QCPlugIn */


/* debug [class_header]: Header for QCPlugIn */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QCPlugIn */
// An interface definition for the [QCPlugIn] class.
type IQCPlugIn interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for QCPlugIn */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QCPlugIn */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QCPlugIn */
// Alloc allocates a new instance without initialization.
func (qc _QCPlugInClass) Alloc() QCPlugIn {
	rv := objc.Send[QCPlugIn](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QCPlugIn */
// A base class to subclass for writing custom patches.
//
// The class provides the base class to subclass for writing custom Quartz Composer patches. You implement a custom patch by subclassing , overriding the appropriate methods, packaging the code as an object, and installing the bundle in the appropriate location. A bundle can contain more than one subclass of , allowing you to provide a suite of custom patches in one bundle. provides detailed instructions on how to create and package a custom patch. supplements the information in the programming guide. The methods related to the executing the custom patch (called when the Quartz Composer engine is rendering) are passed an opaque object that conforms to the protocol. This object represents the execution context of the object. You should not retain the execution context or use it outside of the scope of the execution method that it is passed to.


// A base class to subclass for writing custom patches.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QCPlugIn *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QCPlugIn */

// Returns a dictionary that contains strings for the user interface that describe the custom patch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugIn/attributes()
func (qc _QCPlugInClass) Attributes() foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](objc.ID(qc.class), objc.Sel("attributes"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Attributes) */


// Returns a dictionary that contains strings for the user interface that describe the optional attributes for ports created from properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugIn/attributesForPropertyPort(withKey:)
func (qc _QCPlugInClass) AttributesForPropertyPortWithKey(key objc.IObject /* cross-framework: NSString */) foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](objc.ID(qc.class), objc.Sel("attributesForPropertyPortWithKey:"), key)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AttributesForPropertyPortWithKey) */


// Returns the execution mode of the custom patch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugIn/executionMode()
func (qc _QCPlugInClass) ExecutionMode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(qc.class), objc.Sel("executionMode"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExecutionMode) */


// Loads a Quartz Composer plug-in bundle from the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugIn/load(atPath:)
func (qc _QCPlugInClass) LoadPlugInAtPath(path objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](objc.ID(qc.class), objc.Sel("loadPlugInAtPath:"), path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadPlugInAtPath) */


// Returns the keys for the internal settings of a custom patch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugIn/plugInKeys()
func (qc _QCPlugInClass) PlugInKeys() foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(qc.class), objc.Sel("plugInKeys"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PlugInKeys) */


// Registers a subclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugIn/registerClass(_:)
func (qc _QCPlugInClass) RegisterPlugInClass(aClass objc.Class) {
	objc.Send[objc.ID](objc.ID(qc.class), objc.Sel("registerPlugInClass:"), aClass)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RegisterPlugInClass) */


// Returns and array of property port keys in the order you want them to appear in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugIn/sortedPropertyPortKeys()
func (qc _QCPlugInClass) SortedPropertyPortKeys() foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(qc.class), objc.Sel("sortedPropertyPortKeys"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SortedPropertyPortKeys) */


// Returns the time mode for the custom patch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugIn/timeMode()
func (qc _QCPlugInClass) TimeMode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(qc.class), objc.Sel("timeMode"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TimeMode) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QCPlugIn */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QCPlugIn */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QCPlugIn */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QCPlugIn */



