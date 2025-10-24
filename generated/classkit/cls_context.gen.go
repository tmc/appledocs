// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CLSContext */


/* debug [class_header]: Header for CLSContext */
// The class instance for the [SContext] class.
var (
	SContextClass     _SContextClass
	SContextClassOnce sync.Once
)

func getSContextClass() _SContextClass {
	SContextClassOnce.Do(func() {
		SContextClass = _SContextClass{objc.GetClass("CLSContext")}
	})
	return SContextClass
}

type _SContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SContext */
// An interface definition for the [SContext] class.
type ISContext interface {
	ISObject
	
/* debug [class_interface_properties]: Properties for SContext */
	// properties:
	CurrentActivity() ICLSActivity
	CustomTypeName() objc.IObject /* cross-framework: NSString */
	SetCustomTypeName(value objc.IObject /* cross-framework: NSString */)
	DisplayOrder() int
	SetDisplayOrder(value int)
	Identifier() objc.IObject /* cross-framework: NSString */
	IdentifierPath() []string
	Active() bool
	Assignable() bool
	SetAssignable(value bool)
	NavigationChildContexts() []SContext
	Parent() ICLSContext
	ProgressReportingCapabilities() unsafe.Pointer
	SuggestedAge() corefoundation.Range
	SetSuggestedAge(value corefoundation.Range)
	SuggestedCompletionTime() corefoundation.Range
	SetSuggestedCompletionTime(value corefoundation.Range)
	Summary() objc.IObject /* cross-framework: NSString */
	SetSummary(value objc.IObject /* cross-framework: NSString */)
	Thumbnail() ImageRef /* not a class type */
	SetThumbnail(value ImageRef /* not a class type */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	Topic() SContextTopic /* typedef */
	SetTopic(value SContextTopic /* typedef */)
	Type() SContextType
	UniversalLinkURL() objc.IObject /* cross-framework: NSURL */
	SetUniversalLinkURL(value objc.IObject /* cross-framework: NSURL */)
	IsActive() bool
	SetIsActive(value bool)
	IsAssignable() bool
	SetIsAssignable(value bool)
	ContextIdentifierPath() objc.IObject /* cross-framework: NSString */
	SetContextIdentifierPath(value objc.IObject /* cross-framework: NSString */)
	IsClassKitDeepLink() bool
	SetIsClassKitDeepLink(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SContext */
	// methods:
	AddChildContext(child ICLSContext)
	AddNavigationChildContext(child ICLSContext)
	AddProgressReportingCapabilities(capabilities unsafe.Pointer)
	BecomeActive()
	CreateNewActivity() ISActivity
	DescendantMatchingIdentifierPathCompletion(identifierPath []string, completion unsafe.Pointer)
	RemoveFromParent()
	RemoveNavigationChildContext(child ICLSContext)
	ResetProgressReportingCapabilities()
	ResignActive()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SContext */
// Alloc allocates a new instance without initialization.
func (sc _SContextClass) Alloc() SContext {
	rv := objc.Send[SContext](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SContextClass) New() SContext {
	rv := objc.Send[SContext](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SContext) Init() SContext {
	rv := objc.Send[SContext](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SContext) Autorelease() SContext {
	rv := objc.Send[SContext](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSContext creates a new SContext instance.
func NewSContext() SContext {
	return getSContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SContext */
// An area of your app that represents an assignable task, like a quiz or a chapter.
//
// Make it easy for teachers to understand the app content a context represents by configuring it with information like a clear, concise title localized for the regions that your app supports. A context can contain groups of other contexts, like a book that contains chapters or a chapter that contains sections. You can assemble contexts into a hierarchy of up to eight levels that acts as a table of contents for teachers who want to assign your app content. See for more details.


// An area of your app that represents an assignable task, like a quiz or a chapter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext
type SContext struct {
	SObject
}

// SContextFrom constructs a [SContext] from an unsafe.Pointer.
//
// An area of your app that represents an assignable task, like a quiz or a chapter.
func SContextFrom(ptr unsafe.Pointer) SContext {
	return SContext{
		SObject: SObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SContext */

// Initializes a new context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/init(type:identifier:title:)
func NewSContextWithTypeIdentifierTitle(type_ SContextType, identifier objc.IObject /* cross-framework: NSString */, title objc.IObject /* cross-framework: NSString */) SContext {
	instance := getSContextClass().Alloc()
	rv := objc.Send[SContext](instance.ID, objc.Sel("initWithType:identifier:title:"), type_, identifier, title)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSContextWithTypeIdentifierTitle */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SContext */

// Adds the specifed context as a child of the context receiving the method call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/addChildContext(_:)
func (s_ SContext) AddChildContext(child ICLSContext) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addChildContext:"), child)
}/* debug [instance_methods/method]: AddChildContext */


// Adds a child context that users can navigate to from this context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/addNavigationChildContext(_:)
func (s_ SContext) AddNavigationChildContext(child ICLSContext) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addNavigationChildContext:"), child)
}/* debug [instance_methods/method]: AddNavigationChildContext */


// Adds a progress reporting capability to the set of capabilities for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/addProgressReportingCapabilities(_:)
func (s_ SContext) AddProgressReportingCapabilities(capabilities unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addProgressReportingCapabilities:"), capabilities)
}/* debug [instance_methods/method]: AddProgressReportingCapabilities */


// Tells a context to become the active context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/becomeActive()
func (s_ SContext) BecomeActive() {
	objc.Send[objc.ID](s_.ID, objc.Sel("becomeActive"))
}/* debug [instance_methods/method]: BecomeActive */


// Creates and returns a new activity instance for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/createNewActivity()
func (s_ SContext) CreateNewActivity() ISActivity {
	rv := objc.Send[SActivity](s_.ID, objc.Sel("createNewActivity"))
	return rv
}/* debug [instance_methods/method]: CreateNewActivity */


// Finds the context with the given identifier path relative to this context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/descendant(matchingIdentifierPath:completion:)
func (s_ SContext) DescendantMatchingIdentifierPathCompletion(identifierPath []string, completion unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("descendantMatchingIdentifierPath:completion:"), identifierPath, completion)
}/* debug [instance_methods/method]: DescendantMatchingIdentifierPathCompletion */


// Removes the context from its parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/removeFromParent()
func (s_ SContext) RemoveFromParent() {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeFromParent"))
}/* debug [instance_methods/method]: RemoveFromParent */


// Removes the specified context as a presentable child of this context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/removeNavigationChildContext(_:)
func (s_ SContext) RemoveNavigationChildContext(child ICLSContext) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeNavigationChildContext:"), child)
}/* debug [instance_methods/method]: RemoveNavigationChildContext */


// Resets the set of capabilities for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/resetProgressReportingCapabilities()
func (s_ SContext) ResetProgressReportingCapabilities() {
	objc.Send[objc.ID](s_.ID, objc.Sel("resetProgressReportingCapabilities"))
}/* debug [instance_methods/method]: ResetProgressReportingCapabilities */


// Tells a context to stop being the active context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/resignActive()
func (s_ SContext) ResignActive() {
	objc.Send[objc.ID](s_.ID, objc.Sel("resignActive"))
}/* debug [instance_methods/method]: ResignActive */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SContext */

// The activity available for recording progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/currentActivity
func (s_ SContext) CurrentActivity() ICLSActivity {
	rv := objc.Send[SActivity](s_.ID, objc.Sel("currentActivity"))
	return rv
}/* debug [instance_properties/getter]: currentActivity */


// An optional name that the system presents to the user if you choose the custom context type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/customTypeName
func (s_ SContext) CustomTypeName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("customTypeName"))
	return rv
}/* debug [instance_properties/getter]: customTypeName */


// An optional name that the system presents to the user if you choose the custom context type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/customTypeName
func (s_ SContext) SetCustomTypeName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomTypeName:"), value)
}/* debug [instance_properties/setter]: customTypeName */


// The position of a context relative to its siblings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/displayOrder
func (s_ SContext) DisplayOrder() int {
	rv := objc.Send[int](s_.ID, objc.Sel("displayOrder"))
	return rv
}/* debug [instance_properties/getter]: displayOrder */


// The position of a context relative to its siblings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/displayOrder
func (s_ SContext) SetDisplayOrder(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDisplayOrder:"), value)
}/* debug [instance_properties/setter]: displayOrder */


// A string that uniquely identifies a context among its siblings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/identifier
func (s_ SContext) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The identifier path that locates the context within the data store’s context hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/identifierPath
func (s_ SContext) IdentifierPath() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("identifierPath"))
	return rv
}/* debug [instance_properties/getter]: identifierPath */


// A Boolean indicating whether the context is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/isActive
func (s_ SContext) Active() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("active"))
	return rv
}/* debug [instance_properties/getter]: active */


// A Boolean that indicates whether teachers can assign the context as a task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/isAssignable
func (s_ SContext) Assignable() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("assignable"))
	return rv
}/* debug [instance_properties/getter]: assignable */


// A Boolean that indicates whether teachers can assign the context as a task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/isAssignable
func (s_ SContext) SetAssignable(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAssignable:"), value)
}/* debug [instance_properties/setter]: assignable */


// The child contexts that a user can navigate to from this context in the Schoolwork app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/navigationChildContexts
func (s_ SContext) NavigationChildContexts() []SContext {
	rv := objc.Send[[]SContext](s_.ID, objc.Sel("navigationChildContexts"))
	return rv
}/* debug [instance_properties/getter]: navigationChildContexts */


// The direct ancestor of this context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/parent
func (s_ SContext) Parent() ICLSContext {
	rv := objc.Send[SContext](s_.ID, objc.Sel("parent"))
	return rv
}/* debug [instance_properties/getter]: parent */


// The kinds of progress reporting that the context can perform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/progressReportingCapabilities
func (s_ SContext) ProgressReportingCapabilities() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("progressReportingCapabilities"))
	return rv
}/* debug [instance_properties/getter]: progressReportingCapabilities */


// The range of ages, measured in years, for which you deem a context’s content suitable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/suggestedAge
func (s_ SContext) SuggestedAge() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](s_.ID, objc.Sel("suggestedAge"))
	return rv
}/* debug [instance_properties/getter]: suggestedAge */


// The range of ages, measured in years, for which you deem a context’s content suitable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/suggestedAge
func (s_ SContext) SetSuggestedAge(value corefoundation.Range) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSuggestedAge:"), value)
}/* debug [instance_properties/setter]: suggestedAge */


// A suggested time range to complete a task, measured in minutes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/suggestedCompletionTime
func (s_ SContext) SuggestedCompletionTime() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](s_.ID, objc.Sel("suggestedCompletionTime"))
	return rv
}/* debug [instance_properties/getter]: suggestedCompletionTime */


// A suggested time range to complete a task, measured in minutes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/suggestedCompletionTime
func (s_ SContext) SetSuggestedCompletionTime(value corefoundation.Range) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSuggestedCompletionTime:"), value)
}/* debug [instance_properties/setter]: suggestedCompletionTime */


// An optional, user-visible description of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/summary
func (s_ SContext) Summary() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("summary"))
	return rv
}/* debug [instance_properties/getter]: summary */


// An optional, user-visible description of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/summary
func (s_ SContext) SetSummary(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSummary:"), value)
}/* debug [instance_properties/setter]: summary */


// An optional thumbnail image associated with the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/thumbnail
func (s_ SContext) Thumbnail() ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](s_.ID, objc.Sel("thumbnail"))
	return rv
}/* debug [instance_properties/getter]: thumbnail */


// An optional thumbnail image associated with the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/thumbnail
func (s_ SContext) SetThumbnail(value ImageRef /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setThumbnail:"), value)
}/* debug [instance_properties/setter]: thumbnail */


// The name of the context as it appears to users.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/title
func (s_ SContext) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The name of the context as it appears to users.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/title
func (s_ SContext) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// The area of study to which a context relates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/topic
func (s_ SContext) Topic() SContextTopic /* typedef */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("topic"))
	return rv
}/* debug [instance_properties/getter]: topic */


// The area of study to which a context relates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/topic
func (s_ SContext) SetTopic(value SContextTopic /* typedef */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTopic:"), value)
}/* debug [instance_properties/setter]: topic */


// The kind of content a context represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/type
func (s_ SContext) Type() SContextType {
	rv := objc.Send[SContextType](s_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// A URL that leads to the content in your app associated with the current context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/universalLinkURL
func (s_ SContext) UniversalLinkURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("universalLinkURL"))
	return rv
}/* debug [instance_properties/getter]: universalLinkURL */


// A URL that leads to the content in your app associated with the current context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/universalLinkURL
func (s_ SContext) SetUniversalLinkURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUniversalLinkURL:"), value)
}/* debug [instance_properties/setter]: universalLinkURL */


// A Boolean indicating whether the context is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clscontext/isactive
func (s_ SContext) IsActive() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isActive"))
	return rv
}/* debug [instance_properties/getter]: isActive */


// A Boolean indicating whether the context is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clscontext/isactive
func (s_ SContext) SetIsActive(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsActive:"), value)
}/* debug [instance_properties/setter]: isActive */


// A Boolean that indicates whether teachers can assign the context as a task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clscontext/isassignable
func (s_ SContext) IsAssignable() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isAssignable"))
	return rv
}/* debug [instance_properties/getter]: isAssignable */


// A Boolean that indicates whether teachers can assign the context as a task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clscontext/isassignable
func (s_ SContext) SetIsAssignable(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsAssignable:"), value)
}/* debug [instance_properties/setter]: isAssignable */


// The identifier path associated with a user activity generated by an app that adopts ClassKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/contextIdentifierPath
func (s_ SContext) ContextIdentifierPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("contextIdentifierPath"))
	return rv
}/* debug [instance_properties/getter]: contextIdentifierPath */


// The identifier path associated with a user activity generated by an app that adopts ClassKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/contextIdentifierPath
func (s_ SContext) SetContextIdentifierPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContextIdentifierPath:"), value)
}/* debug [instance_properties/setter]: contextIdentifierPath */


// A Boolean value that indicates whether a user activity represents a ClassKit context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/isClassKitDeepLink
func (s_ SContext) IsClassKitDeepLink() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isClassKitDeepLink"))
	return rv
}/* debug [instance_properties/getter]: isClassKitDeepLink */


// A Boolean value that indicates whether a user activity represents a ClassKit context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/isClassKitDeepLink
func (s_ SContext) SetIsClassKitDeepLink(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsClassKitDeepLink:"), value)
}/* debug [instance_properties/setter]: isClassKitDeepLink */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CLSContext */


