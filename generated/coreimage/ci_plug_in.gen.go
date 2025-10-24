// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CIPlugIn */


/* debug [class_header]: Header for CIPlugIn */
// The class instance for the [PlugIn] class.
var (
	PlugInClass     _PlugInClass
	PlugInClassOnce sync.Once
)

func getPlugInClass() _PlugInClass {
	PlugInClassOnce.Do(func() {
		PlugInClass = _PlugInClass{objc.GetClass("CIPlugIn")}
	})
	return PlugInClass
}

type _PlugInClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlugIn */
// An interface definition for the [PlugIn] class.
type IPlugIn interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PlugIn */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlugIn */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlugIn */
// Alloc allocates a new instance without initialization.
func (pc _PlugInClass) Alloc() PlugIn {
	rv := objc.Send[PlugIn](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlugInClass) New() PlugIn {
	rv := objc.Send[PlugIn](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlugIn) Init() PlugIn {
	rv := objc.Send[PlugIn](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlugIn) Autorelease() PlugIn {
	rv := objc.Send[PlugIn](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlugIn creates a new PlugIn instance.
func NewPlugIn() PlugIn {
	return getPlugInClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlugIn */
// The mechanism for loading image units in macOS.
//
// An image unit is an image processing bundle that contains one or more Core Image filters. Th extension indicates one or more filters packaged as an image unit.


// The mechanism for loading image units in macOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn
type PlugIn struct {
	objectivec.Object
}

// PlugInFrom constructs a [PlugIn] from an unsafe.Pointer.
//
// The mechanism for loading image units in macOS.
func PlugInFrom(ptr unsafe.Pointer) PlugIn {
	return PlugIn{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlugIn *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlugIn */

// Loads filters from an image unit that have the appropriate executable status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn/load(_:allowExecutableCode:)
func (pc _PlugInClass) LoadPlugInAllowExecutableCode(url objc.IObject /* cross-framework: NSURL */, allowExecutableCode bool) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("loadPlugIn:allowExecutableCode:"), url, allowExecutableCode)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadPlugInAllowExecutableCode) */


// Scans directories for files that have the extension and then loads the image units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn/loadAllPlugIns()
func (pc _PlugInClass) LoadAllPlugIns() {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("loadAllPlugIns"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadAllPlugIns) */


// Loads a non-executable plug-in specified by its URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn/loadNonExecutablePlugIn(_:)
func (pc _PlugInClass) LoadNonExecutablePlugIn(url objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("loadNonExecutablePlugIn:"), url)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadNonExecutablePlugIn) */


// Scans directories for plugins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn/loadNonExecutablePlugIns()
func (pc _PlugInClass) LoadNonExecutablePlugIns() {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("loadNonExecutablePlugIns"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadNonExecutablePlugIns) */


// Loads filters from an image unit that have the appropriate executable status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn/loadPlugIn:allowNonExecutable:
func (pc _PlugInClass) LoadPlugInAllowNonExecutable(url objc.IObject /* cross-framework: NSURL */, allowNonExecutable bool) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("loadPlugIn:allowNonExecutable:"), url, allowNonExecutable)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadPlugInAllowNonExecutable) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlugIn */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlugIn */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlugIn */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CIPlugIn */



