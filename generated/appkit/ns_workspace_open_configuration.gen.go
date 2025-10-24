// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSWorkspaceOpenConfiguration */


/* debug [class_header]: Header for NSWorkspaceOpenConfiguration */
// The class instance for the [WorkspaceOpenConfiguration] class.
var (
	WorkspaceOpenConfigurationClass     _WorkspaceOpenConfigurationClass
	WorkspaceOpenConfigurationClassOnce sync.Once
)

func getWorkspaceOpenConfigurationClass() _WorkspaceOpenConfigurationClass {
	WorkspaceOpenConfigurationClassOnce.Do(func() {
		WorkspaceOpenConfigurationClass = _WorkspaceOpenConfigurationClass{objc.GetClass("NSWorkspaceOpenConfiguration")}
	})
	return WorkspaceOpenConfigurationClass
}

type _WorkspaceOpenConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WorkspaceOpenConfiguration */
// An interface definition for the [WorkspaceOpenConfiguration] class.
type IWorkspaceOpenConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WorkspaceOpenConfiguration */
	// properties:
	Activates() bool
	SetActivates(value bool)
	AddsToRecentItems() bool
	SetAddsToRecentItems(value bool)
	AllowsRunningApplicationSubstitution() bool
	SetAllowsRunningApplicationSubstitution(value bool)
	AppleEvent() foundation.AppleEventDescriptor
	SetAppleEvent(value foundation.AppleEventDescriptor)
	Architecture() objectivec.IObject
	SetArchitecture(value objectivec.IObject)
	Arguments() []string
	SetArguments(value []string)
	CreatesNewApplicationInstance() bool
	SetCreatesNewApplicationInstance(value bool)
	Environment() foundation.IDictionary
	SetEnvironment(value foundation.IDictionary)
	Hides() bool
	SetHides(value bool)
	HidesOthers() bool
	SetHidesOthers(value bool)
	ForPrinting() bool
	SetForPrinting(value bool)
	PromptsUserIfNeeded() bool
	SetPromptsUserIfNeeded(value bool)
	RequiresUniversalLinks() bool
	SetRequiresUniversalLinks(value bool)
	IsForPrinting() bool
	SetIsForPrinting(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WorkspaceOpenConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WorkspaceOpenConfiguration */
// Alloc allocates a new instance without initialization.
func (wc _WorkspaceOpenConfigurationClass) Alloc() WorkspaceOpenConfiguration {
	rv := objc.Send[WorkspaceOpenConfiguration](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WorkspaceOpenConfigurationClass) New() WorkspaceOpenConfiguration {
	rv := objc.Send[WorkspaceOpenConfiguration](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WorkspaceOpenConfiguration) Init() WorkspaceOpenConfiguration {
	rv := objc.Send[WorkspaceOpenConfiguration](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WorkspaceOpenConfiguration) Autorelease() WorkspaceOpenConfiguration {
	rv := objc.Send[WorkspaceOpenConfiguration](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWorkspaceOpenConfiguration creates a new WorkspaceOpenConfiguration instance.
func NewWorkspaceOpenConfiguration() WorkspaceOpenConfiguration {
	return getWorkspaceOpenConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WorkspaceOpenConfiguration */
// The configuration options for opening URLs or launching apps.
//
// Create an object before launching an app or opening a URL using the shared object. Use the properties of this object to customize the behavior of the launched app or the handling of the URLs. For example, you might tell the app to hide itself immediately after launch.


// The configuration options for opening URLs or launching apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration
type WorkspaceOpenConfiguration struct {
	objectivec.Object
}

// WorkspaceOpenConfigurationFrom constructs a [WorkspaceOpenConfiguration] from an unsafe.Pointer.
//
// The configuration options for opening URLs or launching apps.
func WorkspaceOpenConfigurationFrom(ptr unsafe.Pointer) WorkspaceOpenConfiguration {
	return WorkspaceOpenConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WorkspaceOpenConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WorkspaceOpenConfiguration */

// Creates and returns a new workspace configuration object containing default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspaceOpenConfiguration/configuration
func (wc _WorkspaceOpenConfigurationClass) Configuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(wc.class), objc.Sel("configuration"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Configuration) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WorkspaceOpenConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WorkspaceOpenConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WorkspaceOpenConfiguration */

// A Boolean value indicating whether the system activates the app and brings it to the foreground.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/activates
func (w_ WorkspaceOpenConfiguration) Activates() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("activates"))
	return rv
}/* debug [instance_properties/getter]: activates */


// A Boolean value indicating whether the system activates the app and brings it to the foreground.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/activates
func (w_ WorkspaceOpenConfiguration) SetActivates(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setActivates:"), value)
}/* debug [instance_properties/setter]: activates */


// A Boolean value indicating whether to add the app or documents to the Recent Items menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/addsToRecentItems
func (w_ WorkspaceOpenConfiguration) AddsToRecentItems() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("addsToRecentItems"))
	return rv
}/* debug [instance_properties/getter]: addsToRecentItems */


// A Boolean value indicating whether to add the app or documents to the Recent Items menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/addsToRecentItems
func (w_ WorkspaceOpenConfiguration) SetAddsToRecentItems(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAddsToRecentItems:"), value)
}/* debug [instance_properties/setter]: addsToRecentItems */


// A Boolean value that indicates whether to use a running instance of an application even if it’s at a different URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/allowsRunningApplicationSubstitution
func (w_ WorkspaceOpenConfiguration) AllowsRunningApplicationSubstitution() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsRunningApplicationSubstitution"))
	return rv
}/* debug [instance_properties/getter]: allowsRunningApplicationSubstitution */


// A Boolean value that indicates whether to use a running instance of an application even if it’s at a different URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/allowsRunningApplicationSubstitution
func (w_ WorkspaceOpenConfiguration) SetAllowsRunningApplicationSubstitution(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsRunningApplicationSubstitution:"), value)
}/* debug [instance_properties/setter]: allowsRunningApplicationSubstitution */


// The first Apple event to send to the new app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/appleEvent
func (w_ WorkspaceOpenConfiguration) AppleEvent() foundation.AppleEventDescriptor {
	rv := objc.Send[foundation.AppleEventDescriptor](w_.ID, objc.Sel("appleEvent"))
	return rv
}/* debug [instance_properties/getter]: appleEvent */


// The first Apple event to send to the new app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/appleEvent
func (w_ WorkspaceOpenConfiguration) SetAppleEvent(value foundation.AppleEventDescriptor) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAppleEvent:"), value)
}/* debug [instance_properties/setter]: appleEvent */


// The architecture version of the app to launch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/architecture
func (w_ WorkspaceOpenConfiguration) Architecture() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](w_.ID, objc.Sel("architecture"))
	return rv
}/* debug [instance_properties/getter]: architecture */


// The architecture version of the app to launch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/architecture
func (w_ WorkspaceOpenConfiguration) SetArchitecture(value objectivec.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setArchitecture:"), value)
}/* debug [instance_properties/setter]: architecture */


// The set of command-line arguments to pass to a new app instance at launch time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/arguments
func (w_ WorkspaceOpenConfiguration) Arguments() []string {
	rv := objc.Send[[]string](w_.ID, objc.Sel("arguments"))
	return rv
}/* debug [instance_properties/getter]: arguments */


// The set of command-line arguments to pass to a new app instance at launch time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/arguments
func (w_ WorkspaceOpenConfiguration) SetArguments(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](w_.ID, objc.Sel("setArguments:"), nsArray)
}/* debug [instance_properties/setter]: arguments */


// A Boolean value indicating whether you want the system to launch a new instance of the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/createsNewApplicationInstance
func (w_ WorkspaceOpenConfiguration) CreatesNewApplicationInstance() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("createsNewApplicationInstance"))
	return rv
}/* debug [instance_properties/getter]: createsNewApplicationInstance */


// A Boolean value indicating whether you want the system to launch a new instance of the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/createsNewApplicationInstance
func (w_ WorkspaceOpenConfiguration) SetCreatesNewApplicationInstance(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCreatesNewApplicationInstance:"), value)
}/* debug [instance_properties/setter]: createsNewApplicationInstance */


// The set of environment variables to set in a new app instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/environment
func (w_ WorkspaceOpenConfiguration) Environment() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](w_.ID, objc.Sel("environment"))
	return rv
}/* debug [instance_properties/getter]: environment */


// The set of environment variables to set in a new app instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/environment
func (w_ WorkspaceOpenConfiguration) SetEnvironment(value foundation.IDictionary) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setEnvironment:"), value)
}/* debug [instance_properties/setter]: environment */


// A Boolean value indicating whether you want the app to hide itself after it launches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/hides
func (w_ WorkspaceOpenConfiguration) Hides() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hides"))
	return rv
}/* debug [instance_properties/getter]: hides */


// A Boolean value indicating whether you want the app to hide itself after it launches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/hides
func (w_ WorkspaceOpenConfiguration) SetHides(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHides:"), value)
}/* debug [instance_properties/setter]: hides */


// A Boolean value indicating whether you want to hide all apps except the one that launched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/hidesOthers
func (w_ WorkspaceOpenConfiguration) HidesOthers() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hidesOthers"))
	return rv
}/* debug [instance_properties/getter]: hidesOthers */


// A Boolean value indicating whether you want to hide all apps except the one that launched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/hidesOthers
func (w_ WorkspaceOpenConfiguration) SetHidesOthers(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHidesOthers:"), value)
}/* debug [instance_properties/setter]: hidesOthers */


// A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/isForPrinting
func (w_ WorkspaceOpenConfiguration) ForPrinting() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("forPrinting"))
	return rv
}/* debug [instance_properties/getter]: forPrinting */


// A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/isForPrinting
func (w_ WorkspaceOpenConfiguration) SetForPrinting(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setForPrinting:"), value)
}/* debug [instance_properties/setter]: forPrinting */


// A Boolean value indicating whether to display errors, authentication requests, or other UI elements to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/promptsUserIfNeeded
func (w_ WorkspaceOpenConfiguration) PromptsUserIfNeeded() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("promptsUserIfNeeded"))
	return rv
}/* debug [instance_properties/getter]: promptsUserIfNeeded */


// A Boolean value indicating whether to display errors, authentication requests, or other UI elements to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/promptsUserIfNeeded
func (w_ WorkspaceOpenConfiguration) SetPromptsUserIfNeeded(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPromptsUserIfNeeded:"), value)
}/* debug [instance_properties/setter]: promptsUserIfNeeded */


// A Boolean value indicating whether you require the URL to have an associated universal link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/requiresUniversalLinks
func (w_ WorkspaceOpenConfiguration) RequiresUniversalLinks() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("requiresUniversalLinks"))
	return rv
}/* debug [instance_properties/getter]: requiresUniversalLinks */


// A Boolean value indicating whether you require the URL to have an associated universal link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/requiresUniversalLinks
func (w_ WorkspaceOpenConfiguration) SetRequiresUniversalLinks(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setRequiresUniversalLinks:"), value)
}/* debug [instance_properties/setter]: requiresUniversalLinks */


// A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/openconfiguration/isforprinting
func (w_ WorkspaceOpenConfiguration) IsForPrinting() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isForPrinting"))
	return rv
}/* debug [instance_properties/getter]: isForPrinting */


// A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/openconfiguration/isforprinting
func (w_ WorkspaceOpenConfiguration) SetIsForPrinting(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsForPrinting:"), value)
}/* debug [instance_properties/setter]: isForPrinting */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSWorkspaceOpenConfiguration */



