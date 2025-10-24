// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [SContext] class.
type ISContext interface {
	ISObject
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
	NavigationChildContexts() []ISContext
	Parent() ICLSContext
	ProgressReportingCapabilities() unsafe.Pointer
	SuggestedAge() objc.IObject /* cross-framework: Range */
	SetSuggestedAge(value objc.IObject /* cross-framework: Range */)
	SuggestedCompletionTime() objc.IObject /* cross-framework: Range */
	SetSuggestedCompletionTime(value objc.IObject /* cross-framework: Range */)
	Summary() objc.IObject /* cross-framework: NSString */
	SetSummary(value objc.IObject /* cross-framework: NSString */)
	Thumbnail() ImageRef /* not a class type */
	SetThumbnail(value ImageRef /* not a class type */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	Topic() objc.IObject /* cross-framework: SContextTopic */
	SetTopic(value objc.IObject /* cross-framework: SContextTopic */)
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
}

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

// Alloc allocates a new instance without initialization.
func (sc _SContextClass) Alloc() SContext {
	rv := objc.Send[SContext](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Initializes a new context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/init(type:identifier:title:)
func NewSContextWithTypeIdentifierTitle(type_ SContextType, identifier objc.IObject /* cross-framework: NSString */, title objc.IObject /* cross-framework: NSString */) SContext {
	instance := getSContextClass().Alloc()
	rv := objc.Send[SContext](instance.ID, objc.Sel("initWithType:identifier:title:"), type_, identifier, title)
	rv.Autorelease()
	return rv
}



// Adds the specifed context as a child of the context receiving the method call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/addChildContext(_:)
func (s_ SContext) AddChildContext(child ICLSContext) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addChildContext:"), child)
}


// Adds a child context that users can navigate to from this context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/addNavigationChildContext(_:)
func (s_ SContext) AddNavigationChildContext(child ICLSContext) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addNavigationChildContext:"), child)
}


// Adds a progress reporting capability to the set of capabilities for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/addProgressReportingCapabilities(_:)
func (s_ SContext) AddProgressReportingCapabilities(capabilities unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addProgressReportingCapabilities:"), capabilities)
}


// Tells a context to become the active context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/becomeActive()
func (s_ SContext) BecomeActive() {
	objc.Send[objc.ID](s_.ID, objc.Sel("becomeActive"))
}


// Creates and returns a new activity instance for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/createNewActivity()
func (s_ SContext) CreateNewActivity() ISActivity {
	rv := objc.Send[SActivity](s_.ID, objc.Sel("createNewActivity"))
	return rv
}


// Finds the context with the given identifier path relative to this context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/descendant(matchingIdentifierPath:completion:)
func (s_ SContext) DescendantMatchingIdentifierPathCompletion(identifierPath []string, completion unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("descendantMatchingIdentifierPath:completion:"), identifierPath, completion)
}


// Removes the context from its parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/removeFromParent()
func (s_ SContext) RemoveFromParent() {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeFromParent"))
}


// Removes the specified context as a presentable child of this context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/removeNavigationChildContext(_:)
func (s_ SContext) RemoveNavigationChildContext(child ICLSContext) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeNavigationChildContext:"), child)
}


// Resets the set of capabilities for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/resetProgressReportingCapabilities()
func (s_ SContext) ResetProgressReportingCapabilities() {
	objc.Send[objc.ID](s_.ID, objc.Sel("resetProgressReportingCapabilities"))
}


// Tells a context to stop being the active context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/resignActive()
func (s_ SContext) ResignActive() {
	objc.Send[objc.ID](s_.ID, objc.Sel("resignActive"))
}


// The activity available for recording progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/currentActivity
func (s_ SContext) CurrentActivity() ICLSActivity {
	rv := objc.Send[SActivity](s_.ID, objc.Sel("currentActivity"))
	return rv
}


// An optional name that the system presents to the user if you choose the custom context type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/customTypeName
func (s_ SContext) CustomTypeName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("customTypeName"))
	return rv
}


// An optional name that the system presents to the user if you choose the custom context type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/customTypeName
func (s_ SContext) SetCustomTypeName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomTypeName:"), value)
}


// The position of a context relative to its siblings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/displayOrder
func (s_ SContext) DisplayOrder() int {
	rv := objc.Send[int](s_.ID, objc.Sel("displayOrder"))
	return rv
}


// The position of a context relative to its siblings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/displayOrder
func (s_ SContext) SetDisplayOrder(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDisplayOrder:"), value)
}


// A string that uniquely identifies a context among its siblings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/identifier
func (s_ SContext) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("identifier"))
	return rv
}


// The identifier path that locates the context within the data store’s context hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/identifierPath
func (s_ SContext) IdentifierPath() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("identifierPath"))
	return rv
}


// A Boolean indicating whether the context is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/isActive
func (s_ SContext) Active() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("active"))
	return rv
}


// A Boolean that indicates whether teachers can assign the context as a task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/isAssignable
func (s_ SContext) Assignable() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("assignable"))
	return rv
}


// A Boolean that indicates whether teachers can assign the context as a task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/isAssignable
func (s_ SContext) SetAssignable(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAssignable:"), value)
}


// The child contexts that a user can navigate to from this context in the Schoolwork app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/navigationChildContexts
func (s_ SContext) NavigationChildContexts() []ISContext {
	rv := objc.Send[[]SContext](s_.ID, objc.Sel("navigationChildContexts"))
	return rv
}


// The direct ancestor of this context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/parent
func (s_ SContext) Parent() ICLSContext {
	rv := objc.Send[SContext](s_.ID, objc.Sel("parent"))
	return rv
}


// The kinds of progress reporting that the context can perform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/progressReportingCapabilities
func (s_ SContext) ProgressReportingCapabilities() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("progressReportingCapabilities"))
	return rv
}


// The range of ages, measured in years, for which you deem a context’s content suitable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/suggestedAge
func (s_ SContext) SuggestedAge() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[corefoundation.Range](s_.ID, objc.Sel("suggestedAge"))
	return rv
}


// The range of ages, measured in years, for which you deem a context’s content suitable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/suggestedAge
func (s_ SContext) SetSuggestedAge(value objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSuggestedAge:"), value)
}


// A suggested time range to complete a task, measured in minutes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/suggestedCompletionTime
func (s_ SContext) SuggestedCompletionTime() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[corefoundation.Range](s_.ID, objc.Sel("suggestedCompletionTime"))
	return rv
}


// A suggested time range to complete a task, measured in minutes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/suggestedCompletionTime
func (s_ SContext) SetSuggestedCompletionTime(value objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSuggestedCompletionTime:"), value)
}


// An optional, user-visible description of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/summary
func (s_ SContext) Summary() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("summary"))
	return rv
}


// An optional, user-visible description of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/summary
func (s_ SContext) SetSummary(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSummary:"), value)
}


// An optional thumbnail image associated with the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/thumbnail
func (s_ SContext) Thumbnail() ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](s_.ID, objc.Sel("thumbnail"))
	return rv
}


// An optional thumbnail image associated with the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/thumbnail
func (s_ SContext) SetThumbnail(value ImageRef /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setThumbnail:"), value)
}


// The name of the context as it appears to users.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/title
func (s_ SContext) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("title"))
	return rv
}


// The name of the context as it appears to users.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/title
func (s_ SContext) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitle:"), value)
}


// The area of study to which a context relates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/topic
func (s_ SContext) Topic() objc.IObject /* cross-framework: SContextTopic */ {
	rv := objc.Send[SContextTopic](s_.ID, objc.Sel("topic"))
	return rv
}


// The area of study to which a context relates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/topic
func (s_ SContext) SetTopic(value objc.IObject /* cross-framework: SContextTopic */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTopic:"), value)
}


// The kind of content a context represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/type
func (s_ SContext) Type() SContextType {
	rv := objc.Send[SContextType](s_.ID, objc.Sel("type"))
	return rv
}


// A URL that leads to the content in your app associated with the current context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/universalLinkURL
func (s_ SContext) UniversalLinkURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("universalLinkURL"))
	return rv
}


// A URL that leads to the content in your app associated with the current context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/universalLinkURL
func (s_ SContext) SetUniversalLinkURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUniversalLinkURL:"), value)
}


// A Boolean indicating whether the context is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clscontext/isactive
func (s_ SContext) IsActive() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isActive"))
	return rv
}


// A Boolean indicating whether the context is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clscontext/isactive
func (s_ SContext) SetIsActive(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsActive:"), value)
}


// A Boolean that indicates whether teachers can assign the context as a task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clscontext/isassignable
func (s_ SContext) IsAssignable() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isAssignable"))
	return rv
}


// A Boolean that indicates whether teachers can assign the context as a task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clscontext/isassignable
func (s_ SContext) SetIsAssignable(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsAssignable:"), value)
}


// The identifier path associated with a user activity generated by an app that adopts ClassKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/contextIdentifierPath
func (s_ SContext) ContextIdentifierPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("contextIdentifierPath"))
	return rv
}


// The identifier path associated with a user activity generated by an app that adopts ClassKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/contextIdentifierPath
func (s_ SContext) SetContextIdentifierPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContextIdentifierPath:"), value)
}


// A Boolean value that indicates whether a user activity represents a ClassKit context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/isClassKitDeepLink
func (s_ SContext) IsClassKitDeepLink() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isClassKitDeepLink"))
	return rv
}


// A Boolean value that indicates whether a user activity represents a ClassKit context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/isClassKitDeepLink
func (s_ SContext) SetIsClassKitDeepLink(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsClassKitDeepLink:"), value)
}


