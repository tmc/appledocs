// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSafariApplication */


/* debug [class_header]: Header for SFSafariApplication */
// The class instance for the [SFSafariApplication] class.
var (
	SFSafariApplicationClass     _SFSafariApplicationClass
	SFSafariApplicationClassOnce sync.Once
)

func getSFSafariApplicationClass() _SFSafariApplicationClass {
	SFSafariApplicationClassOnce.Do(func() {
		SFSafariApplicationClass = _SFSafariApplicationClass{objc.GetClass("SFSafariApplication")}
	})
	return SFSafariApplicationClass
}

type _SFSafariApplicationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSafariApplication */
// An interface definition for the [SFSafariApplication] class.
type ISFSafariApplication interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFSafariApplication */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSafariApplication */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSafariApplication */
// Alloc allocates a new instance without initialization.
func (sc _SFSafariApplicationClass) Alloc() SFSafariApplication {
	rv := objc.Send[SFSafariApplication](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSafariApplicationClass) New() SFSafariApplication {
	rv := objc.Send[SFSafariApplication](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariApplication) Init() SFSafariApplication {
	rv := objc.Send[SFSafariApplication](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariApplication) Autorelease() SFSafariApplication {
	rv := objc.Send[SFSafariApplication](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariApplication creates a new SFSafariApplication instance.
func NewSFSafariApplication() SFSafariApplication {
	return getSFSafariApplicationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSafariApplication */
// A proxy for the Safari app.
//
// The class is used by a Safari app extension to access the active Safari window, open a new window, and update the toolbar items on a window. An application that acts as a host container for a Safari app extension can use this class to send messages to the app extension. There is no object instance for this class.


// A proxy for the Safari app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariApplication
type SFSafariApplication struct {
	objectivec.Object
}

// SFSafariApplicationFrom constructs a [SFSafariApplication] from an unsafe.Pointer.
//
// A proxy for the Safari app.
func SFSafariApplicationFrom(ptr unsafe.Pointer) SFSafariApplication {
	return SFSafariApplication{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSafariApplication *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSafariApplication */

// Sends a message to a Safari app extension, launching Safari if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariApplication/dispatchMessage(withName:toExtensionWithIdentifier:userInfo:completionHandler:)
func (sc _SFSafariApplicationClass) DispatchMessageWithNameToExtensionWithIdentifierUserInfoCompletionHandler(messageName objc.IObject /* cross-framework: NSString */, identifier objc.IObject /* cross-framework: NSString */, userInfo foundation.IDictionary, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("dispatchMessageWithName:toExtensionWithIdentifier:userInfo:completionHandler:"), messageName, identifier, userInfo, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DispatchMessageWithNameToExtensionWithIdentifierUserInfoCompletionHandler) */


// Calls the completion handler with the active browser window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariApplication/getActiveWindow(completionHandler:)
func (sc _SFSafariApplicationClass) GetActiveWindowWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getActiveWindowWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GetActiveWindowWithCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariApplication/getAllWindows(completionHandler:)
func (sc _SFSafariApplicationClass) GetAllWindowsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getAllWindowsWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GetAllWindowsWithCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariApplication/getHostApplication(completionHandler:)
func (sc _SFSafariApplicationClass) GetHostApplicationWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getHostApplicationWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GetHostApplicationWithCompletionHandler) */


// Opens a new window with the desired webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariApplication/openWindow(with:completionHandler:)
func (sc _SFSafariApplicationClass) OpenWindowWithURLCompletionHandler(url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("openWindowWithURL:completionHandler:"), url, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OpenWindowWithURLCompletionHandler) */


// Updates the enabled states and badges of toolbar items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariApplication/setToolbarItemsNeedUpdate()
func (sc _SFSafariApplicationClass) SetToolbarItemsNeedUpdate() {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("setToolbarItemsNeedUpdate"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetToolbarItemsNeedUpdate) */


// Launches Safari and opens the preferences panel for a Safari app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariApplication/showPreferencesForExtension(withIdentifier:completionHandler:)
func (sc _SFSafariApplicationClass) ShowPreferencesForExtensionWithIdentifierCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("showPreferencesForExtensionWithIdentifier:completionHandler:"), identifier, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ShowPreferencesForExtensionWithIdentifierCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSafariApplication */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSafariApplication */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSafariApplication */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSafariApplication */



