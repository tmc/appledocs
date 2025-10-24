// Code generated from Apple documentation for FileProviderUI. DO NOT EDIT.

package fileproviderui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class FPUIActionExtensionContext */


/* debug [class_header]: Header for FPUIActionExtensionContext */
// The class instance for the [FPUIActionExtensionContext] class.
var (
	FPUIActionExtensionContextClass     _FPUIActionExtensionContextClass
	FPUIActionExtensionContextClassOnce sync.Once
)

func getFPUIActionExtensionContextClass() _FPUIActionExtensionContextClass {
	FPUIActionExtensionContextClassOnce.Do(func() {
		FPUIActionExtensionContextClass = _FPUIActionExtensionContextClass{objc.GetClass("FPUIActionExtensionContext")}
	})
	return FPUIActionExtensionContextClass
}

type _FPUIActionExtensionContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FPUIActionExtensionContext */
// An interface definition for the [FPUIActionExtensionContext] class.
type IFPUIActionExtensionContext interface {
	foundation.IExtensionContext
	
/* debug [class_interface_properties]: Properties for FPUIActionExtensionContext */
	// properties:
	DomainIdentifier() FileProviderDomainIdentifier /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FPUIActionExtensionContext */
	// methods:
	CancelRequestWithError(error_ objc.IObject /* cross-framework: Error */)
	CompleteRequest()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FPUIActionExtensionContext */
// Alloc allocates a new instance without initialization.
func (fc _FPUIActionExtensionContextClass) Alloc() FPUIActionExtensionContext {
	rv := objc.Send[FPUIActionExtensionContext](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FPUIActionExtensionContextClass) New() FPUIActionExtensionContext {
	rv := objc.Send[FPUIActionExtensionContext](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FPUIActionExtensionContext) Init() FPUIActionExtensionContext {
	rv := objc.Send[FPUIActionExtensionContext](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FPUIActionExtensionContext) Autorelease() FPUIActionExtensionContext {
	rv := objc.Send[FPUIActionExtensionContext](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFPUIActionExtensionContext creates a new FPUIActionExtensionContext instance.
func NewFPUIActionExtensionContext() FPUIActionExtensionContext {
	return getFPUIActionExtensionContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FPUIActionExtensionContext */
// An extension context provided to File Provider UI extensions.


// An extension context provided to File Provider UI extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProviderUI/FPUIActionExtensionContext
type FPUIActionExtensionContext struct {
	foundation.ExtensionContext
}

// FPUIActionExtensionContextFrom constructs a [FPUIActionExtensionContext] from an unsafe.Pointer.
//
// An extension context provided to File Provider UI extensions.
func FPUIActionExtensionContextFrom(ptr unsafe.Pointer) FPUIActionExtensionContext {
	return FPUIActionExtensionContext{
		ExtensionContext: foundation.ExtensionContextFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FPUIActionExtensionContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FPUIActionExtensionContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FPUIActionExtensionContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FPUIActionExtensionContext */

// Cancels the action and returns the provided error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProviderUI/FPUIActionExtensionContext/cancelRequest(withError:)
func (f_ FPUIActionExtensionContext) CancelRequestWithError(error_ objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("cancelRequestWithError:"), error_)
}/* debug [instance_methods/method]: CancelRequestWithError */


// Marks the action as complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProviderUI/FPUIActionExtensionContext/completeRequest()
func (f_ FPUIActionExtensionContext) CompleteRequest() {
	objc.Send[objc.ID](f_.ID, objc.Sel("completeRequest"))
}/* debug [instance_methods/method]: CompleteRequest */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FPUIActionExtensionContext */

// The identifier for the domain managed by the current file provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProviderUI/FPUIActionExtensionContext/domainIdentifier
func (f_ FPUIActionExtensionContext) DomainIdentifier() FileProviderDomainIdentifier /* not a class type */ {
	rv := objc.Send[FileProviderDomainIdentifier](f_.ID, objc.Sel("domainIdentifier"))
	return rv
}/* debug [instance_properties/getter]: domainIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FPUIActionExtensionContext */






