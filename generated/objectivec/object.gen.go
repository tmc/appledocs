// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Object] class.
var objectClass = _ObjectClass{objc.GetClass("NSObject")}

type _ObjectClass struct {
	class objc.Class
}

// An interface definition for the [Object] class.
type IObject interface {
	objc.IObject
	URLResourceDataDidBecomeAvailable(sender unsafe.Pointer, newBytes unsafe.Pointer)
	URLResourceDidFailLoadingWithReason(sender unsafe.Pointer, reason string)
	URLResourceDidCancelLoading(sender unsafe.Pointer)
	URLResourceDidFinishLoading(sender unsafe.Pointer)
	AcceptsPreviewPanelControl(panel unsafe.Pointer) bool
	AccessibilityActivate() bool
	AccessibilityAssistiveTechnologyFocusedIdentifiers() unsafe.Pointer
	AccessibilityDecrement()
	AccessibilityElementAtIndex(index int) objc.ID
	AccessibilityElementCount() int
	AccessibilityElementIsFocused() bool
	AccessibilityHitTest(point unsafe.Pointer) objc.ID
	AccessibilityHitTestWithEvent(point unsafe.Pointer, event unsafe.Pointer) objc.ID
	AccessibilityIncrement()
	AccessibilityLineEndPositionFromCurrentSelection() int
	AccessibilityLineRangeForPosition(position int) unsafe.Pointer
	AccessibilityLineStartPositionFromCurrentSelection() int
	AccessibilityZoomInAtPoint(point unsafe.Pointer) bool
	AccessibilityZoomOutAtPoint(point unsafe.Pointer) bool
	ActionProperty() unsafe.Pointer
	ApplicationDelegateHandlesKey(sender unsafe.Pointer, key string) bool
	AttemptRecoveryFromErrorOptionIndex(error unsafe.Pointer, recoveryOptionIndex uint) bool
	AttemptRecoveryFromErrorOptionIndexDelegateDidRecoverSelectorContextInfo(error unsafe.Pointer, recoveryOptionIndex uint, delegate objc.ID, didRecoverSelector objc.SEL, contextInfo unsafe.Pointer)
	AttributedStringForIdentityPropertiesWithNamesInRecordComparisonRecordsFirstLineAttributesSecondLineAttributes(propertyNames unsafe.Pointer, record unsafe.Pointer, comparisonRecords unsafe.Pointer, firstLineAttributes unsafe.Pointer, secondLineAttributes unsafe.Pointer) unsafe.Pointer
	AttributedStringForPropertiesWithNamesInRecordComparisonRecordsDefaultAttributes(propertyNames unsafe.Pointer, record unsafe.Pointer, comparisonRecords unsafe.Pointer, defaultAttributes unsafe.Pointer) unsafe.Pointer
	AuthorizationViewCreatedAuthorization(view unsafe.Pointer)
	AuthorizationViewDidAuthorize(view unsafe.Pointer)
	AuthorizationViewDidDeauthorize(view unsafe.Pointer)
	AuthorizationViewDidHide(view unsafe.Pointer)
	AuthorizationViewReleasedAuthorization(view unsafe.Pointer)
	AuthorizationViewShouldDeauthorize(view unsafe.Pointer) bool
	AwakeFromNib()
	BeginPreviewPanelControl(panel unsafe.Pointer)
	BrowserAccessibilityAttributedValueInRange(range_ unsafe.Pointer) unsafe.Pointer
	BrowserAccessibilityDeleteTextAtCursor(numberOfCharacters int)
	BrowserAccessibilityInsertTextAtCursor(text string)
	BrowserAccessibilitySelectedTextRange() unsafe.Pointer
	BrowserAccessibilitySetSelectedTextRange(range_ unsafe.Pointer)
	BrowserAccessibilityValueInRange(range_ unsafe.Pointer) unsafe.Pointer
	BurnProgressPanelBurnDidFinish(theBurnPanel unsafe.Pointer, burn unsafe.Pointer) bool
	BurnProgressPanelDidFinish(aNotification unsafe.Pointer)
	BurnProgressPanelWillBegin(aNotification unsafe.Pointer)
	Candidates(sender objc.ID) unsafe.Pointer
	CertificatePanelShowHelp(sender unsafe.Pointer) bool
	ChangeColor(sender objc.ID)
	ChangeFont(sender objc.ID)
	ChooseIdentityPanelShowHelp(sender unsafe.Pointer) bool
	CommitComposition(sender objc.ID)
	CommitEditing() bool
	CommitEditingAndReturnError(error unsafe.Pointer) bool
	CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.ID, didCommitSelector objc.SEL, contextInfo unsafe.Pointer)
	ComposedString(sender objc.ID) objc.ID
	CompositionParameterViewDidChangeParameterWithKey(parameterView unsafe.Pointer, portKey string)
	CompositionParameterViewShouldDisplayParameterWithKeyAttributes(parameterView unsafe.Pointer, portKey string, portAttributes unsafe.Pointer) bool
	CompositionPickerViewDidSelectComposition(pickerView unsafe.Pointer, composition unsafe.Pointer)
	CompositionPickerViewDidStartAnimating(pickerView unsafe.Pointer)
	CompositionPickerViewWillStopAnimating(pickerView unsafe.Pointer)
	ControlTextDidBeginEditing(obj unsafe.Pointer)
	ControlTextDidChange(obj unsafe.Pointer)
	ControlTextDidEndEditing(obj unsafe.Pointer)
	CopyScriptingValueForKeyWithProperties(value objc.ID, key string, properties unsafe.Pointer) objc.ID
	Dealloc()
	DidCommandBySelectorClient(aSelector objc.SEL, sender objc.ID) bool
	DiscardEditing()
	DoesContain(object objc.ID) bool
	DraggedImageBeganAt(image unsafe.Pointer, screenPoint unsafe.Pointer)
	DraggedImageEndedAtDeposited(image unsafe.Pointer, screenPoint unsafe.Pointer, flag bool)
	DraggedImageEndedAtOperation(image unsafe.Pointer, screenPoint unsafe.Pointer, operation unsafe.Pointer)
	DraggedImageMovedTo(image unsafe.Pointer, screenPoint unsafe.Pointer)
	DraggingSourceOperationMaskForLocal(flag bool) unsafe.Pointer
	EndPreviewPanelControl(panel unsafe.Pointer)
	EraseProgressPanelEraseDidFinish(theErasePanel unsafe.Pointer, erase unsafe.Pointer) bool
	EraseProgressPanelDidFinish(aNotification unsafe.Pointer)
	EraseProgressPanelWillBegin(aNotification unsafe.Pointer)
	ExceptionHandlerShouldHandleExceptionMask(sender unsafe.Pointer, exception unsafe.Pointer, aMask uint) bool
	ExceptionHandlerShouldLogExceptionMask(sender unsafe.Pointer, exception unsafe.Pointer, aMask uint) bool
	FileManagerShouldProceedAfterError(fm unsafe.Pointer, errorInfo unsafe.Pointer) bool
	FileManagerWillProcessPath(fm unsafe.Pointer, path string)
	FileTransferServicesAbortCompleteError(inServices unsafe.Pointer, inError unsafe.Pointer)
	FileTransferServicesConnectionCompleteError(inServices unsafe.Pointer, inError unsafe.Pointer)
	FileTransferServicesCopyRemoteFileCompleteError(inServices unsafe.Pointer, inError unsafe.Pointer)
	FileTransferServicesCopyRemoteFileProgressTransferProgress(inServices unsafe.Pointer, inProgressDescription unsafe.Pointer)
	FileTransferServicesCreateFolderCompleteErrorFolder(inServices unsafe.Pointer, inError unsafe.Pointer, inFolderName string)
	FileTransferServicesDisconnectionCompleteError(inServices unsafe.Pointer, inError unsafe.Pointer)
	FileTransferServicesFilePreparationCompleteError(inServices unsafe.Pointer, inError unsafe.Pointer)
	FileTransferServicesPathChangeCompleteErrorFinalPath(inServices unsafe.Pointer, inError unsafe.Pointer, inPath string)
	FileTransferServicesRemoveItemCompleteErrorRemovedItem(inServices unsafe.Pointer, inError unsafe.Pointer, inItemName string)
	FileTransferServicesRetrieveFolderListingCompleteErrorListing(inServices unsafe.Pointer, inError unsafe.Pointer, inListing unsafe.Pointer)
	FileTransferServicesSendFileCompleteError(inServices unsafe.Pointer, inError unsafe.Pointer)
	FileTransferServicesSendFileProgressTransferProgress(inServices unsafe.Pointer, inProgressDescription unsafe.Pointer)
	GetL2CAPChannelRef() unsafe.Pointer
	GetOpenGLBufferContextPixelFormat(contextOut unsafe.Pointer, pixelFormatOut unsafe.Pointer)
	GetPixelBufferPixelFormat(pixelFormatOut unsafe.Pointer)
	HandleEventClient(event unsafe.Pointer, sender objc.ID) bool
	IgnoreModifierKeysWhileDragging() bool
	ImageBrowserBackgroundWasRightClickedWithEvent(aBrowser unsafe.Pointer, event unsafe.Pointer)
	ImageBrowserCellWasDoubleClickedAtIndex(aBrowser unsafe.Pointer, index uint)
	ImageBrowserCellWasRightClickedAtIndexWithEvent(aBrowser unsafe.Pointer, index uint, event unsafe.Pointer)
	ImageBrowserGroupAtIndex(aBrowser unsafe.Pointer, index uint) unsafe.Pointer
	ImageBrowserItemAtIndex(aBrowser unsafe.Pointer, index uint) objc.ID
	ImageBrowserMoveItemsAtIndexesToIndex(aBrowser unsafe.Pointer, indexes unsafe.Pointer, destinationIndex uint) bool
	ImageBrowserRemoveItemsAtIndexes(aBrowser unsafe.Pointer, indexes unsafe.Pointer)
	ImageBrowserWriteItemsAtIndexesToPasteboard(aBrowser unsafe.Pointer, itemIndexes unsafe.Pointer, pasteboard unsafe.Pointer) uint
	ImageBrowserSelectionDidChange(aBrowser unsafe.Pointer)
	ImageRepresentation() objc.ID
	ImageRepresentationType() unsafe.Pointer
	ImageSubtitle() unsafe.Pointer
	ImageTitle() unsafe.Pointer
	ImageUID() unsafe.Pointer
	ImageVersion() uint
	IndexOfAccessibilityElement(element objc.ID) int
	IndicesOfObjectsByEvaluatingObjectSpecifier(specifier unsafe.Pointer) unsafe.Pointer
	InputTextClient(string string, sender objc.ID) bool
	InputTextKeyModifiersClient(string string, keyCode int, flags uint, sender objc.ID) bool
	InverseForRelationshipKey(relationshipKey string) unsafe.Pointer
	IsCaseInsensitiveLike(object string) bool
	IsEqualTo(object objc.ID) bool
	IsGreaterThan(object objc.ID) bool
	IsGreaterThanOrEqualTo(object objc.ID) bool
	IsLessThan(object objc.ID) bool
	IsLessThanOrEqualTo(object objc.ID) bool
	IsLike(object string) bool
	IsNotEqualTo(object objc.ID) bool
	LayerShouldInheritContentsScaleFromWindow(layer unsafe.Pointer, newScale float64, window unsafe.Pointer) bool
	MethodForSelector(aSelector objc.SEL) unsafe.Pointer
	NewScriptingObjectOfClassForValueForKeyWithContentsValueProperties(objectClass objc.Class, key string, contentsValue objc.ID, properties unsafe.Pointer) objc.ID
	NumberOfGroupsInImageBrowser(aBrowser unsafe.Pointer) uint
	NumberOfItemsInImageBrowser(aBrowser unsafe.Pointer) uint
	ObjectDidBeginEditing(editor unsafe.Pointer)
	ObjectDidEndEditing(editor unsafe.Pointer)
	OriginalString(sender objc.ID) unsafe.Pointer
	PanelCompareFilenameWithCaseSensitive(sender objc.ID, name1 string, name2 string, caseSensitive bool) unsafe.Pointer
	PanelDirectoryDidChange(sender objc.ID, path string)
	PanelIsValidFilename(sender objc.ID, filename string) bool
	PanelShouldShowFilename(sender objc.ID, filename string) bool
	PasteboardProvideDataForType(sender unsafe.Pointer, type_ unsafe.Pointer)
	PasteboardChangedOwner(sender unsafe.Pointer)
	PerformSelectorOnThreadWithObjectWaitUntilDoneModes(aSelector objc.SEL, thr unsafe.Pointer, arg objc.ID, wait bool, array unsafe.Pointer)
	PerformActionForPersonIdentifier(person unsafe.Pointer, identifier string)
	PerformSelectorOnMainThreadWithObjectWaitUntilDoneModes(aSelector objc.SEL, arg objc.ID, wait bool, array unsafe.Pointer)
	PrepareForInterfaceBuilder()
	ProvideImageToMTLTextureCommandBufferOriginxOriginyWidthHeightUserInfo(texture unsafe.Pointer, commandBuffer unsafe.Pointer, originx uintptr, originy uintptr, width uintptr, height uintptr, info objc.ID)
	ProvideImageDataBytesPerRowOriginSizeUserInfo(data unsafe.Pointer, rowbytes uintptr, originx uintptr, originy uintptr, width uintptr, height uintptr, info objc.ID)
	QuartzFilterManagerDidAddFilter(sender unsafe.Pointer, filter unsafe.Pointer)
	QuartzFilterManagerDidModifyFilter(sender unsafe.Pointer, filter unsafe.Pointer)
	QuartzFilterManagerDidRemoveFilter(sender unsafe.Pointer, filter unsafe.Pointer)
	QuartzFilterManagerDidSelectFilter(sender unsafe.Pointer, filter unsafe.Pointer)
	ReadLinkQualityForDeviceCompleteDeviceInfoError(controller objc.ID, device unsafe.Pointer, info unsafe.Pointer, error unsafe.Pointer)
	ReadRSSIForDeviceCompleteDeviceInfoError(controller objc.ID, device unsafe.Pointer, info unsafe.Pointer, error unsafe.Pointer)
	RegisterIncomingDataListenerRefCon(listener unsafe.Pointer, refCon unsafe.Pointer) unsafe.Pointer
	RenderIntoOpenGLBufferOnScreenForTime(buffer unsafe.Pointer, screenInOut unsafe.Pointer, timeStamp unsafe.Pointer) bool
	RenderIntoPixelBufferForTime(buffer unsafe.Pointer, timeStamp unsafe.Pointer) bool
	ReplacementObjectForKeyedArchiver(archiver unsafe.Pointer) objc.ID
	ReplacementObjectForArchiver(archiver unsafe.Pointer) objc.ID
	SaveOptionsShouldShowUTType(saveOptions unsafe.Pointer, utType string) bool
	ScriptingValueForSpecifier(objectSpecifier unsafe.Pointer) objc.ID
	SessionDriverDidNegotiateAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool
	SessionDriverDidPullAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool
	SessionDriverDidPushAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool
	SessionDriverDidReceiveSyncAlertAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool
	SessionDriverDidRegisterClientAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool
	SessionDriverWillFinishSessionAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool
	SessionDriverWillNegotiateAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool
	SessionDriverWillPullAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool
	SessionDriverWillPushAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool
	SessionDriverDidCancelSession(sender unsafe.Pointer)
	SessionDriverDidFinishSession(sender unsafe.Pointer)
	SessionDriverWillCancelSession(sender unsafe.Pointer)
	SetSharedObservers(sharedObservers unsafe.Pointer)
	SetupPanelDetermineBestDeviceOfAOrB(aPanel unsafe.Pointer, deviceA unsafe.Pointer, device unsafe.Pointer) unsafe.Pointer
	SetupPanelDeviceContainsSuitableMediaPromptString(aPanel unsafe.Pointer, device unsafe.Pointer, prompt string) bool
	SetupPanelDeviceCouldBeTarget(aPanel unsafe.Pointer, device unsafe.Pointer) bool
	SetupPanelDeviceSelectionChanged(aNotification unsafe.Pointer)
	SetupPanelShouldHandleMediaReservations(aPanel unsafe.Pointer) bool
	ShouldEnableActionForPersonIdentifier(person unsafe.Pointer, identifier string) bool
	TableViewWriteRowsToPasteboard(tableView unsafe.Pointer, rows unsafe.Pointer, pboard unsafe.Pointer) bool
	TitleForPersonIdentifier(person unsafe.Pointer, identifier string) unsafe.Pointer
	ValidModesForFontPanel(fontPanel unsafe.Pointer) unsafe.Pointer
	ValidateMenuItem(menuItem unsafe.Pointer) bool
	ValidateToolbarItem(item unsafe.Pointer) bool
	ViewStringForToolTipPointUserData(view unsafe.Pointer, tag unsafe.Pointer, point unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer
	WorkflowControllerDidError(controller unsafe.Pointer, error unsafe.Pointer)
	WorkflowControllerDidRunAction(controller unsafe.Pointer, action unsafe.Pointer)
	WorkflowControllerWillRunAction(controller unsafe.Pointer, action unsafe.Pointer)
	WorkflowControllerDidRun(controller unsafe.Pointer)
	WorkflowControllerDidStop(controller unsafe.Pointer)
	WorkflowControllerWillRun(controller unsafe.Pointer)
	WorkflowControllerWillStop(controller unsafe.Pointer)
	WriteLength(data unsafe.Pointer, length unsafe.Pointer) unsafe.Pointer
}

// The root class of most Objective-C class hierarchies, from which subclasses inherit a basic interface to the runtime system and the ability to behave as Objective-C objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class
type Object struct {
	objc.ID
}

// ObjectFrom constructs a [Object] from an unsafe.Pointer.
//
// The root class of most Objective-C class hierarchies, from which subclasses inherit a basic interface to the runtime system and the ability to behave as Objective-C objects.
func ObjectFrom(ptr unsafe.Pointer) Object {
	return Object{objc.ID(ptr)}
}
// Alloc allocates a new instance without initialization.
func (oc _ObjectClass) Alloc() Object {
	rv := objc.Send[Object](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (oc _ObjectClass) New() Object {
	rv := objc.Send[Object](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ Object) Init() Object {
	rv := objc.Send[Object](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ Object) Autorelease() Object {
	rv := objc.Send[Object](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewObject creates a new Object instance.
func NewObject() Object {
	return objectClass.New()
}


// Returns a Boolean value that indicates whether the observed object supports automatic key-value observation for the given key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/automaticallyNotifiesObservers(forKey:)
func (oc _ObjectClass) AutomaticallyNotifiesObserversForKey(key string) bool {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("automaticallyNotifiesObserversForKey:"), key)
	return rv
}
// Cancels perform requests previously registered with the instance method. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/cancelPreviousPerformRequests(withTarget:)
func (oc _ObjectClass) CancelPreviousPerformRequestsWithTarget(aTarget objc.ID) {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("cancelPreviousPerformRequestsWithTarget:"), aTarget)
}
// Cancels perform requests previously registered with . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/cancelPreviousPerformRequests(withTarget:selector:object:)
func (oc _ObjectClass) CancelPreviousPerformRequestsWithTargetSelectorObject(aTarget objc.ID, aSelector objc.SEL, anArgument objc.ID) {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("cancelPreviousPerformRequestsWithTarget:selector:object:"), aTarget, aSelector, anArgument)
}
// Returns the class object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/class
func (oc _ObjectClass) Class() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(oc.class), objc.Sel("class"))
	return rv
}
// Overridden to return the names of classes that can be used to decode objects if their class is unavailable. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/classFallbacksForKeyedArchiver()
func (oc _ObjectClass) ClassFallbacksForKeyedArchiver() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("classFallbacksForKeyedArchiver"))
	return rv
}
// Overridden by subclasses to substitute a new class during keyed unarchiving. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/classForKeyedUnarchiver()
func (oc _ObjectClass) ClassForKeyedUnarchiver() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(oc.class), objc.Sel("classForKeyedUnarchiver"))
	return rv
}
// Returns a Boolean value that indicates whether the target conforms to a given protocol. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/conforms(to:)
func (oc _ObjectClass) ConformsToProtocol(protocol unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("conformsToProtocol:"), protocol)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/debugDescription()
func (oc _ObjectClass) DebugDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("debugDescription"))
	return rv
}
// Returns a string that represents the contents of the receiving class. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/description()
func (oc _ObjectClass) Description() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("description"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/hash()
func (oc _ObjectClass) Hash() uint {
	rv := objc.Send[uint](objc.ID(oc.class), objc.Sel("hash"))
	return rv
}
// Initializes the class before it receives its first message. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/initialize()
func (oc _ObjectClass) Initialize() {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("initialize"))
}
// Locates and returns the address of the implementation of the instance method identified by a given selector. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/instanceMethod(for:)
func (oc _ObjectClass) InstanceMethodForSelector(aSelector objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("instanceMethodForSelector:"), aSelector)
	return rv
}
// Returns an object that contains a description of the instance method identified by a given selector. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/instanceMethodSignatureForSelector:
func (oc _ObjectClass) InstanceMethodSignatureForSelector(aSelector objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("instanceMethodSignatureForSelector:"), aSelector)
	return rv
}
// Returns a Boolean value that indicates whether instances of the receiver are capable of responding to a given selector. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/instancesRespond(to:)
func (oc _ObjectClass) InstancesRespondToSelector(aSelector objc.SEL) bool {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("instancesRespondToSelector:"), aSelector)
	return rv
}
// Returns a Boolean value that indicates whether the receiving class is a subclass of, or identical to, a given class. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isSubclass(of:)
func (oc _ObjectClass) IsSubclassOfClass(aClass objc.Class) bool {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("isSubclassOfClass:"), aClass)
	return rv
}
// Invoked whenever a class or category is added to the Objective-C runtime; implement this method to perform class-specific behavior upon loading. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/load()
func (oc _ObjectClass) Load() {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("load"))
}
// Dynamically provides an implementation for a given selector for a class method. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/resolveClassMethod(_:)
func (oc _ObjectClass) ResolveClassMethod(sel objc.SEL) bool {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("resolveClassMethod:"), sel)
	return rv
}
// Dynamically provides an implementation for a given selector for an instance method. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/resolveInstanceMethod(_:)
func (oc _ObjectClass) ResolveInstanceMethod(sel objc.SEL) bool {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("resolveInstanceMethod:"), sel)
	return rv
}
// Sets the receiver’s version number. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setVersion(_:)
func (oc _ObjectClass) SetVersion(aVersion int) {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("setVersion:"), aVersion)
}
// Returns the class object for the receiver’s superclass. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/superclass()
func (oc _ObjectClass) Superclass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(oc.class), objc.Sel("superclass"))
	return rv
}
// Returns the version number assigned to the class. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/version()
func (oc _ObjectClass) Version() int {
	rv := objc.Send[int](objc.ID(oc.class), objc.Sel("version"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/withL2CAPChannelRef:
func (oc _ObjectClass) WithL2CAPChannelRef(l2capChannelRef unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("withL2CAPChannelRef:"), l2capChannelRef)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/URL:resourceDataDidBecomeAvailable:
func (o_ Object) URLResourceDataDidBecomeAvailable(sender unsafe.Pointer, newBytes unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("URL:resourceDataDidBecomeAvailable:"), sender, newBytes)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/URL:resourceDidFailLoadingWithReason:
func (o_ Object) URLResourceDidFailLoadingWithReason(sender unsafe.Pointer, reason string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("URL:resourceDidFailLoadingWithReason:"), sender, reason)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/URLResourceDidCancelLoading:
func (o_ Object) URLResourceDidCancelLoading(sender unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("URLResourceDidCancelLoading:"), sender)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/URLResourceDidFinishLoading:
func (o_ Object) URLResourceDidFinishLoading(sender unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("URLResourceDidFinishLoading:"), sender)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/acceptsPreviewPanelControl(_:)
func (o_ Object) AcceptsPreviewPanelControl(panel unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("acceptsPreviewPanelControl:"), panel)
	return rv
}
// Tells the element to activate itself and report the success or failure of the operation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityActivate()
func (o_ Object) AccessibilityActivate() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityActivate"))
	return rv
}
// Returns a set of identifier keys indicating which assistive app has focus on the accessibility element. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityAssistiveTechnologyFocusedIdentifiers()
func (o_ Object) AccessibilityAssistiveTechnologyFocusedIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityAssistiveTechnologyFocusedIdentifiers"))
	return rv
}
// Tells the accessibility element to decrement the value of its content. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityDecrement()
func (o_ Object) AccessibilityDecrement() {
	objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityDecrement"))
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElement(at:)
func (o_ Object) AccessibilityElementAtIndex(index int) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityElementAtIndex:"), index)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElementCount()
func (o_ Object) AccessibilityElementCount() int {
	rv := objc.Send[int](o_.ID, objc.Sel("accessibilityElementCount"))
	return rv
}
// Returns a Boolean value indicating whether an assistive technology is focused on the accessibility element. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElementIsFocused()
func (o_ Object) AccessibilityElementIsFocused() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityElementIsFocused"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityHitTest(_:)
func (o_ Object) AccessibilityHitTest(point unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityHitTest:"), point)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityHitTest(_:event:)
func (o_ Object) AccessibilityHitTestWithEvent(point unsafe.Pointer, event unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityHitTest:withEvent:"), point, event)
	return rv
}
// Tells the accessibility element to increment the value of its content. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityIncrement()
func (o_ Object) AccessibilityIncrement() {
	objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityIncrement"))
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityLineEndPositionFromCurrentSelection()
func (o_ Object) AccessibilityLineEndPositionFromCurrentSelection() int {
	rv := objc.Send[int](o_.ID, objc.Sel("accessibilityLineEndPositionFromCurrentSelection"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityLineRange(forPosition:)
func (o_ Object) AccessibilityLineRangeForPosition(position int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityLineRangeForPosition:"), position)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityLineStartPositionFromCurrentSelection()
func (o_ Object) AccessibilityLineStartPositionFromCurrentSelection() int {
	rv := objc.Send[int](o_.ID, objc.Sel("accessibilityLineStartPositionFromCurrentSelection"))
	return rv
}
// Zooms in on the content at the specified point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityZoomIn(at:)
func (o_ Object) AccessibilityZoomInAtPoint(point unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityZoomInAtPoint:"), point)
	return rv
}
// Zooms out from the content at the specified point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityZoomOut(at:)
func (o_ Object) AccessibilityZoomOutAtPoint(point unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityZoomOutAtPoint:"), point)
	return rv
}
// Sent to the delegate to request the property the action applies to. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/actionProperty()
func (o_ Object) ActionProperty() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("actionProperty"))
	return rv
}
// Sent by Cocoa’s built-in scripting support during execution of or script commands to find out if the delegate can handle operations on the specified key-value key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/application:delegateHandlesKey:
func (o_ Object) ApplicationDelegateHandlesKey(sender unsafe.Pointer, key string) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("application:delegateHandlesKey:"), sender, key)
	return rv
}
// Implemented to attempt a recovery from an error noted in an application-modal dialog. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/attemptRecovery(fromError:optionIndex:)
func (o_ Object) AttemptRecoveryFromErrorOptionIndex(error unsafe.Pointer, recoveryOptionIndex uint) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("attemptRecoveryFromError:optionIndex:"), error, recoveryOptionIndex)
	return rv
}
// Implemented to attempt a recovery from an error noted in a document-modal sheet. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/attemptRecovery(fromError:optionIndex:delegate:didRecoverSelector:contextInfo:)
func (o_ Object) AttemptRecoveryFromErrorOptionIndexDelegateDidRecoverSelectorContextInfo(error unsafe.Pointer, recoveryOptionIndex uint, delegate objc.ID, didRecoverSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("attemptRecoveryFromError:optionIndex:delegate:didRecoverSelector:contextInfo:"), error, recoveryOptionIndex, delegate, didRecoverSelector, contextInfo)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/attributedStringForIdentityPropertiesWithNames:inRecord:comparisonRecords:firstLineAttributes:secondLineAttributes:
func (o_ Object) AttributedStringForIdentityPropertiesWithNamesInRecordComparisonRecordsFirstLineAttributesSecondLineAttributes(propertyNames unsafe.Pointer, record unsafe.Pointer, comparisonRecords unsafe.Pointer, firstLineAttributes unsafe.Pointer, secondLineAttributes unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("attributedStringForIdentityPropertiesWithNames:inRecord:comparisonRecords:firstLineAttributes:secondLineAttributes:"), propertyNames, record, comparisonRecords, firstLineAttributes, secondLineAttributes)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/attributedStringForPropertiesWithNames:inRecord:comparisonRecords:defaultAttributes:
func (o_ Object) AttributedStringForPropertiesWithNamesInRecordComparisonRecordsDefaultAttributes(propertyNames unsafe.Pointer, record unsafe.Pointer, comparisonRecords unsafe.Pointer, defaultAttributes unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("attributedStringForPropertiesWithNames:inRecord:comparisonRecords:defaultAttributes:"), propertyNames, record, comparisonRecords, defaultAttributes)
	return rv
}
// Sent to the delegate to indicate the authorization object has been created or changed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/authorizationViewCreatedAuthorization(_:)
func (o_ Object) AuthorizationViewCreatedAuthorization(view unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("authorizationViewCreatedAuthorization:"), view)
}
// Sent to the delegate to indicate the user was authorized and the authorization view was changed to unlocked. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/authorizationViewDidAuthorize(_:)
func (o_ Object) AuthorizationViewDidAuthorize(view unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("authorizationViewDidAuthorize:"), view)
}
// Sent to the delegate to indicate the user was deauthorized and the authorization view was changed to locked. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/authorizationViewDidDeauthorize(_:)
func (o_ Object) AuthorizationViewDidDeauthorize(view unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("authorizationViewDidDeauthorize:"), view)
}
// Sent to the delegate to indicate that the view’s visibility has changed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/authorizationViewDidHide(_:)
func (o_ Object) AuthorizationViewDidHide(view unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("authorizationViewDidHide:"), view)
}
// Sent to the delegate to indicate that deauthorization is about to occur. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/authorizationViewReleasedAuthorization(_:)
func (o_ Object) AuthorizationViewReleasedAuthorization(view unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("authorizationViewReleasedAuthorization:"), view)
}
// Sent to the delegate when a user clicks the open lock icon. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/authorizationViewShouldDeauthorize(_:)
func (o_ Object) AuthorizationViewShouldDeauthorize(view unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("authorizationViewShouldDeauthorize:"), view)
	return rv
}
// Overridden by subclasses to substitute another object in place of the object that was decoded and subsequently received this message. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/awakeAfter(using:)
func (o_ Object) AwakeAfterUsingCoder(coder unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("awakeAfterUsingCoder:"), coder)
	return rv
}
// Prepares the receiver for service after it has been loaded from an Interface Builder archive, or nib file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/awakeFromNib()
func (o_ Object) AwakeFromNib() {
	objc.Send[objc.ID](o_.ID, objc.Sel("awakeFromNib"))
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/beginPreviewPanelControl(_:)
func (o_ Object) BeginPreviewPanelControl(panel unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("beginPreviewPanelControl:"), panel)
}
// Returns the value for this element within the given range, as an attributed string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityAttributedValue(in:)
func (o_ Object) BrowserAccessibilityAttributedValueInRange(range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("browserAccessibilityAttributedValueInRange:"), range_)
	return rv
}
// Deletes text from the element at the current cursor position. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityDeleteTextAtCursor(numberOfCharacters:)
func (o_ Object) BrowserAccessibilityDeleteTextAtCursor(numberOfCharacters int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("browserAccessibilityDeleteTextAtCursor:"), numberOfCharacters)
}
// Inserts text into the element at the current cursor position. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityInsertTextAtCursor(text:)
func (o_ Object) BrowserAccessibilityInsertTextAtCursor(text string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("browserAccessibilityInsertTextAtCursor:"), text)
}
// Returns the range of selected text in the element. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilitySelectedTextRange()
func (o_ Object) BrowserAccessibilitySelectedTextRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("browserAccessibilitySelectedTextRange"))
	return rv
}
// Updates the element’s selected text. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilitySetSelectedTextRange(_:)
func (o_ Object) BrowserAccessibilitySetSelectedTextRange(range_ unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("browserAccessibilitySetSelectedTextRange:"), range_)
}
// Returns this element’s value in the given range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityValue(in:)
func (o_ Object) BrowserAccessibilityValueInRange(range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("browserAccessibilityValueInRange:"), range_)
	return rv
}
// Allows the delegate to handle the end-of-burn feedback. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/burnProgressPanel(_:burnDidFinish:)
func (o_ Object) BurnProgressPanelBurnDidFinish(theBurnPanel unsafe.Pointer, burn unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("burnProgressPanel:burnDidFinish:"), theBurnPanel, burn)
	return rv
}
// Notification sent by the panel after ordering out. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/burnProgressPanelDidFinish(_:)
func (o_ Object) BurnProgressPanelDidFinish(aNotification unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("burnProgressPanelDidFinish:"), aNotification)
}
// Notification sent by the panel before display. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/burnProgressPanelWillBegin(_:)
func (o_ Object) BurnProgressPanelWillBegin(aNotification unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("burnProgressPanelWillBegin:"), aNotification)
}
// Returns an array of candidates. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/candidates(_:)
func (o_ Object) Candidates(sender objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("candidates:"), sender)
	return rv
}
// Implements custom help behavior for the modal panel. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/certificatePanelShowHelp(_:)
func (o_ Object) CertificatePanelShowHelp(sender unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("certificatePanelShowHelp:"), sender)
	return rv
}
// Sent to the first responder when the user selects a color in an object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/changeColor:
func (o_ Object) ChangeColor(sender objc.ID) {
	objc.Send[objc.ID](o_.ID, objc.Sel("changeColor:"), sender)
}
// Informs responders of a font change. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/changeFont:
func (o_ Object) ChangeFont(sender objc.ID) {
	objc.Send[objc.ID](o_.ID, objc.Sel("changeFont:"), sender)
}
// Implements custom help behavior for the modal panel. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/chooseIdentityPanelShowHelp(_:)
func (o_ Object) ChooseIdentityPanelShowHelp(sender unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("chooseIdentityPanelShowHelp:"), sender)
	return rv
}
// Informs the controller that the composition should be committed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/commitComposition(_:)
func (o_ Object) CommitComposition(sender objc.ID) {
	objc.Send[objc.ID](o_.ID, objc.Sel("commitComposition:"), sender)
}
// Returns whether the receiver was able to commit any pending edits. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/commitEditing
func (o_ Object) CommitEditing() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("commitEditing"))
	return rv
}
// Attempt to commit pending edits, returning an error in the case of failure. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/commitEditingAndReturnError:
func (o_ Object) CommitEditingAndReturnError(error unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("commitEditingAndReturnError:"), error)
	return rv
}
// Attempt to commit any currently edited results of the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/commitEditingWithDelegate:didCommitSelector:contextInfo:
func (o_ Object) CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.ID, didCommitSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("commitEditingWithDelegate:didCommitSelector:contextInfo:"), delegate, didCommitSelector, contextInfo)
}
// Return the current composed string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/composedString(_:)
func (o_ Object) ComposedString(sender objc.ID) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("composedString:"), sender)
	return rv
}
// Called after an input parameter in the composition parameter view has been edited. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/compositionParameterView(_:didChangeParameterWithKey:)
func (o_ Object) CompositionParameterViewDidChangeParameterWithKey(parameterView unsafe.Pointer, portKey string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("compositionParameterView:didChangeParameterWithKey:"), parameterView, portKey)
}
// Allows you to define which composition parameters are visible in the user interface when the composition parameter view refreshes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/compositionParameterView(_:shouldDisplayParameterWithKey:attributes:)
func (o_ Object) CompositionParameterViewShouldDisplayParameterWithKeyAttributes(parameterView unsafe.Pointer, portKey string, portAttributes unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("compositionParameterView:shouldDisplayParameterWithKey:attributes:"), parameterView, portKey, portAttributes)
	return rv
}
// Performs custom tasks when the selected composition in the composition picker view changes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/compositionPickerView(_:didSelect:)
func (o_ Object) CompositionPickerViewDidSelectComposition(pickerView unsafe.Pointer, composition unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("compositionPickerView:didSelectComposition:"), pickerView, composition)
}
// Performs custom tasks when the composition picker view starts animating a composition. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/compositionPickerViewDidStartAnimating(_:)
func (o_ Object) CompositionPickerViewDidStartAnimating(pickerView unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("compositionPickerViewDidStartAnimating:"), pickerView)
}
// Performs custom tasks when the composition picker view stops animating a composition. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/compositionPickerViewWillStopAnimating(_:)
func (o_ Object) CompositionPickerViewWillStopAnimating(pickerView unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("compositionPickerViewWillStopAnimating:"), pickerView)
}
// Sent when a control with editable text begins an editing session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/controlTextDidBeginEditing:
func (o_ Object) ControlTextDidBeginEditing(obj unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("controlTextDidBeginEditing:"), obj)
}
// Sent when the text in the receiving control changes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/controlTextDidChange:
func (o_ Object) ControlTextDidChange(obj unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("controlTextDidChange:"), obj)
}
// Sent when a control with editable text ends an editing session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/controlTextDidEndEditing:
func (o_ Object) ControlTextDidEndEditing(obj unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("controlTextDidEndEditing:"), obj)
}
// Creates and returns one or more scripting objects to be inserted into the specified relationship by copying the passed-in value and setting the properties in the copied object or objects. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/copyScriptingValue(_:forKey:withProperties:)
func (o_ Object) CopyScriptingValueForKeyWithProperties(value objc.ID, key string, properties unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("copyScriptingValue:forKey:withProperties:"), value, key, properties)
	return rv
}
// Deallocates the memory occupied by the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/dealloc
func (o_ Object) Dealloc() {
	objc.Send[objc.ID](o_.ID, objc.Sel("dealloc"))
}
// Processes a command generated by user action such as typing certain keys or pressing the mouse button. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/didCommand(by:client:)
func (o_ Object) DidCommandBySelectorClient(aSelector objc.SEL, sender objc.ID) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("didCommandBySelector:client:"), aSelector, sender)
	return rv
}
// Causes the receiver to discard any changes, restoring the previous values. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/discardEditing
func (o_ Object) DiscardEditing() {
	objc.Send[objc.ID](o_.ID, objc.Sel("discardEditing"))
}
// Returns a Boolean value that indicates whether the receiver contains a given object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/doesContain(_:)
func (o_ Object) DoesContain(object objc.ID) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("doesContain:"), object)
	return rv
}
// Handles messages the receiver doesn’t recognize. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/doesNotRecognizeSelector(_:)
func (o_ Object) DoesNotRecognizeSelector(aSelector objc.SEL) {
	objc.Send[objc.ID](o_.ID, objc.Sel("doesNotRecognizeSelector:"), aSelector)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/draggedImage:beganAt:
func (o_ Object) DraggedImageBeganAt(image unsafe.Pointer, screenPoint unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("draggedImage:beganAt:"), image, screenPoint)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/draggedImage:endedAt:deposited:
func (o_ Object) DraggedImageEndedAtDeposited(image unsafe.Pointer, screenPoint unsafe.Pointer, flag bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("draggedImage:endedAt:deposited:"), image, screenPoint, flag)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/draggedImage:endedAt:operation:
func (o_ Object) DraggedImageEndedAtOperation(image unsafe.Pointer, screenPoint unsafe.Pointer, operation unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("draggedImage:endedAt:operation:"), image, screenPoint, operation)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/draggedImage:movedTo:
func (o_ Object) DraggedImageMovedTo(image unsafe.Pointer, screenPoint unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("draggedImage:movedTo:"), image, screenPoint)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/draggingSourceOperationMaskForLocal:
func (o_ Object) DraggingSourceOperationMaskForLocal(flag bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("draggingSourceOperationMaskForLocal:"), flag)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/endPreviewPanelControl(_:)
func (o_ Object) EndPreviewPanelControl(panel unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("endPreviewPanelControl:"), panel)
}
// Notification sent by the panel before display. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/eraseProgressPanel(_:eraseDidFinish:)
func (o_ Object) EraseProgressPanelEraseDidFinish(theErasePanel unsafe.Pointer, erase unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("eraseProgressPanel:eraseDidFinish:"), theErasePanel, erase)
	return rv
}
// Notification sent by the panel after ordering out. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/eraseProgressPanelDidFinish(_:)
func (o_ Object) EraseProgressPanelDidFinish(aNotification unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("eraseProgressPanelDidFinish:"), aNotification)
}
// Notification sent by the panel before display. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/eraseProgressPanelWillBegin(_:)
func (o_ Object) EraseProgressPanelWillBegin(aNotification unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("eraseProgressPanelWillBegin:"), aNotification)
}
// Implemented by the delegate to evaluate whether the delegating exception handler should handle a given exception. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/exceptionHandler(_:shouldHandle:mask:)
func (o_ Object) ExceptionHandlerShouldHandleExceptionMask(sender unsafe.Pointer, exception unsafe.Pointer, aMask uint) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("exceptionHandler:shouldHandleException:mask:"), sender, exception, aMask)
	return rv
}
// Implemented by the delegate to evaluate whether the delegating exception hangler should log a given exception. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/exceptionHandler(_:shouldLogException:mask:)
func (o_ Object) ExceptionHandlerShouldLogExceptionMask(sender unsafe.Pointer, exception unsafe.Pointer, aMask uint) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("exceptionHandler:shouldLogException:mask:"), sender, exception, aMask)
	return rv
}
// An object sends this message to its handler for each error it encounters when copying, moving, removing, or linking files or directories. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileManager(_:shouldProceedAfterError:)
func (o_ Object) FileManagerShouldProceedAfterError(fm unsafe.Pointer, errorInfo unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("fileManager:shouldProceedAfterError:"), fm, errorInfo)
	return rv
}
// An object sends this message to a handler immediately before attempting to move, copy, rename, or delete, or before attempting to link to a given path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileManager(_:willProcessPath:)
func (o_ Object) FileManagerWillProcessPath(fm unsafe.Pointer, path string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileManager:willProcessPath:"), fm, path)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesAbortComplete(_:error:)
func (o_ Object) FileTransferServicesAbortCompleteError(inServices unsafe.Pointer, inError unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesAbortComplete:error:"), inServices, inError)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesConnectionComplete(_:error:)
func (o_ Object) FileTransferServicesConnectionCompleteError(inServices unsafe.Pointer, inError unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesConnectionComplete:error:"), inServices, inError)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesCopyRemoteFileComplete(_:error:)
func (o_ Object) FileTransferServicesCopyRemoteFileCompleteError(inServices unsafe.Pointer, inError unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesCopyRemoteFileComplete:error:"), inServices, inError)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesCopyRemoteFileProgress(_:transferProgress:)
func (o_ Object) FileTransferServicesCopyRemoteFileProgressTransferProgress(inServices unsafe.Pointer, inProgressDescription unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesCopyRemoteFileProgress:transferProgress:"), inServices, inProgressDescription)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesCreateFolderComplete(_:error:folder:)
func (o_ Object) FileTransferServicesCreateFolderCompleteErrorFolder(inServices unsafe.Pointer, inError unsafe.Pointer, inFolderName string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesCreateFolderComplete:error:folder:"), inServices, inError, inFolderName)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesDisconnectionComplete(_:error:)
func (o_ Object) FileTransferServicesDisconnectionCompleteError(inServices unsafe.Pointer, inError unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesDisconnectionComplete:error:"), inServices, inError)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesFilePreparationComplete(_:error:)
func (o_ Object) FileTransferServicesFilePreparationCompleteError(inServices unsafe.Pointer, inError unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesFilePreparationComplete:error:"), inServices, inError)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesPathChangeComplete(_:error:finalPath:)
func (o_ Object) FileTransferServicesPathChangeCompleteErrorFinalPath(inServices unsafe.Pointer, inError unsafe.Pointer, inPath string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesPathChangeComplete:error:finalPath:"), inServices, inError, inPath)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesRemoveItemComplete(_:error:removedItem:)
func (o_ Object) FileTransferServicesRemoveItemCompleteErrorRemovedItem(inServices unsafe.Pointer, inError unsafe.Pointer, inItemName string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesRemoveItemComplete:error:removedItem:"), inServices, inError, inItemName)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesRetrieveFolderListingComplete(_:error:listing:)
func (o_ Object) FileTransferServicesRetrieveFolderListingCompleteErrorListing(inServices unsafe.Pointer, inError unsafe.Pointer, inListing unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesRetrieveFolderListingComplete:error:listing:"), inServices, inError, inListing)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesSendFileComplete(_:error:)
func (o_ Object) FileTransferServicesSendFileCompleteError(inServices unsafe.Pointer, inError unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesSendFileComplete:error:"), inServices, inError)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesSendFileProgress(_:transferProgress:)
func (o_ Object) FileTransferServicesSendFileProgressTransferProgress(inServices unsafe.Pointer, inProgressDescription unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesSendFileProgress:transferProgress:"), inServices, inProgressDescription)
}
// Overridden by subclasses to forward messages to other objects. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/forwardInvocation:
func (o_ Object) ForwardInvocation(anInvocation unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("forwardInvocation:"), anInvocation)
}
// Returns the object to which unrecognized messages should first be directed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/forwardingTarget(for:)
func (o_ Object) ForwardingTargetForSelector(aSelector objc.SEL) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("forwardingTargetForSelector:"), aSelector)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/getL2CAPChannelRef
func (o_ Object) GetL2CAPChannelRef() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("getL2CAPChannelRef"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/getOpenGLBufferContext:pixelFormat:
func (o_ Object) GetOpenGLBufferContextPixelFormat(contextOut unsafe.Pointer, pixelFormatOut unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("getOpenGLBufferContext:pixelFormat:"), contextOut, pixelFormatOut)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/getPixelBufferPixelFormat:
func (o_ Object) GetPixelBufferPixelFormat(pixelFormatOut unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("getPixelBufferPixelFormat:"), pixelFormatOut)
}
// Handles key down and mouse events. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/handle(_:client:)
func (o_ Object) HandleEventClient(event unsafe.Pointer, sender objc.ID) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("handleEvent:client:"), event, sender)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/ignoreModifierKeysWhileDragging
func (o_ Object) IgnoreModifierKeysWhileDragging() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("ignoreModifierKeysWhileDragging"))
	return rv
}
// Performs custom tasks when the user right-clicks the image browser view background. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:backgroundWasRightClickedWith:)
func (o_ Object) ImageBrowserBackgroundWasRightClickedWithEvent(aBrowser unsafe.Pointer, event unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("imageBrowser:backgroundWasRightClickedWithEvent:"), aBrowser, event)
}
// Performs custom tasks when the user double-clicks an item in the image browser view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:cellWasDoubleClickedAt:)
func (o_ Object) ImageBrowserCellWasDoubleClickedAtIndex(aBrowser unsafe.Pointer, index uint) {
	objc.Send[objc.ID](o_.ID, objc.Sel("imageBrowser:cellWasDoubleClickedAtIndex:"), aBrowser, index)
}
// Performs custom tasks when the user right-clicks an item in the image browser view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:cellWasRightClickedAt:with:)
func (o_ Object) ImageBrowserCellWasRightClickedAtIndexWithEvent(aBrowser unsafe.Pointer, index uint, event unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("imageBrowser:cellWasRightClickedAtIndex:withEvent:"), aBrowser, index, event)
}
// Returns the group at the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:groupAt:)
func (o_ Object) ImageBrowserGroupAtIndex(aBrowser unsafe.Pointer, index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("imageBrowser:groupAtIndex:"), aBrowser, index)
	return rv
}
// Returns an object for the item in an image browser view that corresponds to the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:itemAt:)
func (o_ Object) ImageBrowserItemAtIndex(aBrowser unsafe.Pointer, index uint) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("imageBrowser:itemAtIndex:"), aBrowser, index)
	return rv
}
// Signals that the specified items should be moved to the specified destination. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:moveItemsAt:to:)
func (o_ Object) ImageBrowserMoveItemsAtIndexesToIndex(aBrowser unsafe.Pointer, indexes unsafe.Pointer, destinationIndex uint) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("imageBrowser:moveItemsAtIndexes:toIndex:"), aBrowser, indexes, destinationIndex)
	return rv
}
// Signals that a remove operation should be applied to the specified items. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:removeItemsAt:)
func (o_ Object) ImageBrowserRemoveItemsAtIndexes(aBrowser unsafe.Pointer, indexes unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("imageBrowser:removeItemsAtIndexes:"), aBrowser, indexes)
}
// Signals that a drag should begin. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:writeItemsAt:to:)
func (o_ Object) ImageBrowserWriteItemsAtIndexesToPasteboard(aBrowser unsafe.Pointer, itemIndexes unsafe.Pointer, pasteboard unsafe.Pointer) uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("imageBrowser:writeItemsAtIndexes:toPasteboard:"), aBrowser, itemIndexes, pasteboard)
	return rv
}
// Performs custom tasks when the selection changes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowserSelectionDidChange(_:)
func (o_ Object) ImageBrowserSelectionDidChange(aBrowser unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("imageBrowserSelectionDidChange:"), aBrowser)
}
// Returns the image to display. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageRepresentation()
func (o_ Object) ImageRepresentation() objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("imageRepresentation"))
	return rv
}
// Returns the representation type of the image to display. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageRepresentationType()
func (o_ Object) ImageRepresentationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("imageRepresentationType"))
	return rv
}
// Returns the display subtitle of the image. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageSubtitle()
func (o_ Object) ImageSubtitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("imageSubtitle"))
	return rv
}
// Returns the display title of the image. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageTitle()
func (o_ Object) ImageTitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("imageTitle"))
	return rv
}
// Returns a unique string that identifies the data source item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageUID()
func (o_ Object) ImageUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("imageUID"))
	return rv
}
// Returns the version of the item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageVersion()
func (o_ Object) ImageVersion() uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("imageVersion"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/index(ofAccessibilityElement:)
func (o_ Object) IndexOfAccessibilityElement(element objc.ID) int {
	rv := objc.Send[int](o_.ID, objc.Sel("indexOfAccessibilityElement:"), element)
	return rv
}
// Returns the indices of the specified container objects. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/indicesOfObjects(byEvaluatingObjectSpecifier:)
func (o_ Object) IndicesOfObjectsByEvaluatingObjectSpecifier(specifier unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("indicesOfObjectsByEvaluatingObjectSpecifier:"), specifier)
	return rv
}
// Handles key down events that do not map to an action method. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/inputText(_:client:)
func (o_ Object) InputTextClient(string string, sender objc.ID) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("inputText:client:"), string, sender)
	return rv
}
// Receives Unicode, the key code that generated it, and any modifier flags. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/inputText(_:key:modifiers:client:)
func (o_ Object) InputTextKeyModifiersClient(string string, keyCode int, flags uint, sender objc.ID) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("inputText:key:modifiers:client:"), string, keyCode, flags, sender)
	return rv
}
// For a given key that defines the name of the relationship from the receiver’s class to another class, returns the name of the relationship from the other class to the receiver’s class. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/inverse(forRelationshipKey:)
func (o_ Object) InverseForRelationshipKey(relationshipKey string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("inverseForRelationshipKey:"), relationshipKey)
	return rv
}
// Returns a Boolean value that indicates whether receiver is considered to be “like” a given string when the case of characters in the receiver is ignored. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isCaseInsensitiveLike(_:)
func (o_ Object) IsCaseInsensitiveLike(object string) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isCaseInsensitiveLike:"), object)
	return rv
}
// Returns a Boolean value that indicates whether the receiver is equal to another given object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isEqual(to:)
func (o_ Object) IsEqualTo(object objc.ID) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isEqualTo:"), object)
	return rv
}
// Returns a Boolean value that indicates whether the receiver is greater than another given object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isGreaterThan(_:)
func (o_ Object) IsGreaterThan(object objc.ID) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isGreaterThan:"), object)
	return rv
}
// Returns a Boolean value that indicates whether the receiver is greater than or equal to another given object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isGreaterThanOrEqual(to:)
func (o_ Object) IsGreaterThanOrEqualTo(object objc.ID) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isGreaterThanOrEqualTo:"), object)
	return rv
}
// Returns a Boolean value that indicates whether the receiver is less than another given object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isLessThan(_:)
func (o_ Object) IsLessThan(object objc.ID) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isLessThan:"), object)
	return rv
}
// Returns a Boolean value that indicates whether the receiver is less than or equal to another given object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isLessThanOrEqual(to:)
func (o_ Object) IsLessThanOrEqualTo(object objc.ID) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isLessThanOrEqualTo:"), object)
	return rv
}
// Returns a Boolean value that indicates whether the receiver is “like” another given object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isLike(_:)
func (o_ Object) IsLike(object string) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isLike:"), object)
	return rv
}
// Returns a Boolean value that indicates whether the receiver is not equal to another given object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isNotEqual(to:)
func (o_ Object) IsNotEqualTo(object objc.ID) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isNotEqualTo:"), object)
	return rv
}
// Invoked when a resolution changes occurs for the window that hosts the layer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/layer:shouldInheritContentsScale:fromWindow:
func (o_ Object) LayerShouldInheritContentsScaleFromWindow(layer unsafe.Pointer, newScale float64, window unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("layer:shouldInheritContentsScale:fromWindow:"), layer, newScale, window)
	return rv
}
// Locates and returns the address of the receiver’s implementation of a method so it can be called as a function. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/method(for:)
func (o_ Object) MethodForSelector(aSelector objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("methodForSelector:"), aSelector)
	return rv
}
// Returns an object that contains a description of the method identified by a given selector. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/methodSignatureForSelector:
func (o_ Object) MethodSignatureForSelector(aSelector objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("methodSignatureForSelector:"), aSelector)
	return rv
}
// Creates and returns an instance of a scriptable class, setting its contents and properties, for insertion into the relationship identified by the key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/newScriptingObject(of:forValueForKey:withContentsValue:properties:)
func (o_ Object) NewScriptingObjectOfClassForValueForKeyWithContentsValueProperties(objectClass objc.Class, key string, contentsValue objc.ID, properties unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("newScriptingObjectOfClass:forValueForKey:withContentsValue:properties:"), objectClass, key, contentsValue, properties)
	return rv
}
// Returns the number of groups in an image browser view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/numberOfGroups(inImageBrowser:)
func (o_ Object) NumberOfGroupsInImageBrowser(aBrowser unsafe.Pointer) uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("numberOfGroupsInImageBrowser:"), aBrowser)
	return rv
}
// Returns the number of records managed by the data source object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/numberOfItems(inImageBrowser:)
func (o_ Object) NumberOfItemsInImageBrowser(aBrowser unsafe.Pointer) uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("numberOfItemsInImageBrowser:"), aBrowser)
	return rv
}
// This message should be sent to the receiver when has uncommitted changes that can affect the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/objectDidBeginEditing:
func (o_ Object) ObjectDidBeginEditing(editor unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("objectDidBeginEditing:"), editor)
}
// This message should be sent to the receiver when has finished editing a property belonging to the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/objectDidEndEditing:
func (o_ Object) ObjectDidEndEditing(editor unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("objectDidEndEditing:"), editor)
}
// Return the string that consists of the precomposed Unicode characters. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/originalString(_:)
func (o_ Object) OriginalString(sender objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("originalString:"), sender)
	return rv
}
// Controls the ordering of files presented by the object specified. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/panel:compareFilename:with:caseSensitive:
func (o_ Object) PanelCompareFilenameWithCaseSensitive(sender objc.ID, name1 string, name2 string, caseSensitive bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("panel:compareFilename:with:caseSensitive:"), sender, name1, name2, caseSensitive)
	return rv
}
// Tells the delegate that the user has changed the selected directory in the object specified. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/panel:directoryDidChange:
func (o_ Object) PanelDirectoryDidChange(sender objc.ID, path string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("panel:directoryDidChange:"), sender, path)
}
// Gives the delegate the opportunity to validate selected items. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/panel:isValidFilename:
func (o_ Object) PanelIsValidFilename(sender objc.ID, filename string) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("panel:isValidFilename:"), sender, filename)
	return rv
}
// Gives the delegate the opportunity to filter items that it doesn’t want the user to choose. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/panel:shouldShowFilename:
func (o_ Object) PanelShouldShowFilename(sender objc.ID, filename string) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("panel:shouldShowFilename:"), sender, filename)
	return rv
}
// Implemented by an owner object to provide promised data. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/pasteboard:provideDataForType:
func (o_ Object) PasteboardProvideDataForType(sender unsafe.Pointer, type_ unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("pasteboard:provideDataForType:"), sender, type_)
}
// Notifies a prior owner of the specified pasteboard (and owners of representations on the pasteboard) that the pasteboard has changed owners. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/pasteboardChangedOwner:
func (o_ Object) PasteboardChangedOwner(sender unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("pasteboardChangedOwner:"), sender)
}
// Invokes a method of the receiver on the specified thread using the default mode. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/perform(_:on:with:waitUntilDone:)
func (o_ Object) PerformSelectorOnThreadWithObjectWaitUntilDone(aSelector objc.SEL, thr unsafe.Pointer, arg objc.ID, wait bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelector:onThread:withObject:waitUntilDone:"), aSelector, thr, arg, wait)
}
// Invokes a method of the receiver on the specified thread using the specified modes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/perform(_:on:with:waitUntilDone:modes:)
func (o_ Object) PerformSelectorOnThreadWithObjectWaitUntilDoneModes(aSelector objc.SEL, thr unsafe.Pointer, arg objc.ID, wait bool, array unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelector:onThread:withObject:waitUntilDone:modes:"), aSelector, thr, arg, wait, array)
}
// Invokes a method of the receiver on the current thread using the default mode after a delay. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/perform(_:with:afterDelay:)
func (o_ Object) PerformSelectorWithObjectAfterDelay(aSelector objc.SEL, anArgument objc.ID, delay TimeInterval) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelector:withObject:afterDelay:"), aSelector, anArgument, delay)
}
// Invokes a method of the receiver on the current thread using the specified modes after a delay. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/perform(_:with:afterDelay:inModes:)
func (o_ Object) PerformSelectorWithObjectAfterDelayInModes(aSelector objc.SEL, anArgument objc.ID, delay TimeInterval, modes unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelector:withObject:afterDelay:inModes:"), aSelector, anArgument, delay, modes)
}
// Sent to the delegate to perform the action. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/performAction(for:identifier:)
func (o_ Object) PerformActionForPersonIdentifier(person unsafe.Pointer, identifier string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performActionForPerson:identifier:"), person, identifier)
}
// Invokes a method of the receiver on a new background thread. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/performSelector(inBackground:with:)
func (o_ Object) PerformSelectorInBackgroundWithObject(aSelector objc.SEL, arg objc.ID) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelectorInBackground:withObject:"), aSelector, arg)
}
// Invokes a method of the receiver on the main thread using the default mode. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/performSelector(onMainThread:with:waitUntilDone:)
func (o_ Object) PerformSelectorOnMainThreadWithObjectWaitUntilDone(aSelector objc.SEL, arg objc.ID, wait bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelectorOnMainThread:withObject:waitUntilDone:"), aSelector, arg, wait)
}
// Invokes a method of the receiver on the main thread using the specified modes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/performSelector(onMainThread:with:waitUntilDone:modes:)
func (o_ Object) PerformSelectorOnMainThreadWithObjectWaitUntilDoneModes(aSelector objc.SEL, arg objc.ID, wait bool, array unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelectorOnMainThread:withObject:waitUntilDone:modes:"), aSelector, arg, wait, array)
}
// Called when a designable object is created in Interface Builder. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/prepareForInterfaceBuilder()
func (o_ Object) PrepareForInterfaceBuilder() {
	objc.Send[objc.ID](o_.ID, objc.Sel("prepareForInterfaceBuilder"))
}
// An optional method that an image provider object way implement. With this method, the provider object can use the Metal API to provide pixel data into a MTLTexture when the image object is rendered. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/provideImage(to:commandBuffer:originx:originy:width:height:userInfo:)
func (o_ Object) ProvideImageToMTLTextureCommandBufferOriginxOriginyWidthHeightUserInfo(texture unsafe.Pointer, commandBuffer unsafe.Pointer, originx uintptr, originy uintptr, width uintptr, height uintptr, info objc.ID) {
	objc.Send[objc.ID](o_.ID, objc.Sel("provideImageToMTLTexture:commandBuffer:originx:originy:width:height:userInfo:"), texture, commandBuffer, originx, originy, width, height, info)
}
// Supplies data to a object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/provideImageData(_:bytesPerRow:origin:_:size:_:userInfo:)
func (o_ Object) ProvideImageDataBytesPerRowOriginSizeUserInfo(data unsafe.Pointer, rowbytes uintptr, originx uintptr, originy uintptr, width uintptr, height uintptr, info objc.ID) {
	objc.Send[objc.ID](o_.ID, objc.Sel("provideImageData:bytesPerRow:origin::size::userInfo:"), data, rowbytes, originx, originy, width, height, info)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/quartzFilterManager(_:didAdd:)
func (o_ Object) QuartzFilterManagerDidAddFilter(sender unsafe.Pointer, filter unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("quartzFilterManager:didAddFilter:"), sender, filter)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/quartzFilterManager(_:didModifyFilter:)
func (o_ Object) QuartzFilterManagerDidModifyFilter(sender unsafe.Pointer, filter unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("quartzFilterManager:didModifyFilter:"), sender, filter)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/quartzFilterManager(_:didRemove:)
func (o_ Object) QuartzFilterManagerDidRemoveFilter(sender unsafe.Pointer, filter unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("quartzFilterManager:didRemoveFilter:"), sender, filter)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/quartzFilterManager(_:didSelect:)
func (o_ Object) QuartzFilterManagerDidSelectFilter(sender unsafe.Pointer, filter unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("quartzFilterManager:didSelectFilter:"), sender, filter)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/readLinkQuality(forDeviceComplete:device:info:error:)
func (o_ Object) ReadLinkQualityForDeviceCompleteDeviceInfoError(controller objc.ID, device unsafe.Pointer, info unsafe.Pointer, error unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("readLinkQualityForDeviceComplete:device:info:error:"), controller, device, info, error)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/readRSSI(forDeviceComplete:device:info:error:)
func (o_ Object) ReadRSSIForDeviceCompleteDeviceInfoError(controller objc.ID, device unsafe.Pointer, info unsafe.Pointer, error unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("readRSSIForDeviceComplete:device:info:error:"), controller, device, info, error)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/registerIncomingDataListener:refCon:
func (o_ Object) RegisterIncomingDataListenerRefCon(listener unsafe.Pointer, refCon unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("registerIncomingDataListener:refCon:"), listener, refCon)
	return rv
}
// Called for each frame to be sent to Messages. This method will not be called on the main thread. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/renderIntoOpenGLBuffer:onScreen:forTime:
func (o_ Object) RenderIntoOpenGLBufferOnScreenForTime(buffer unsafe.Pointer, screenInOut unsafe.Pointer, timeStamp unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("renderIntoOpenGLBuffer:onScreen:forTime:"), buffer, screenInOut, timeStamp)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/renderIntoPixelBuffer:forTime:
func (o_ Object) RenderIntoPixelBufferForTime(buffer unsafe.Pointer, timeStamp unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("renderIntoPixelBuffer:forTime:"), buffer, timeStamp)
	return rv
}
// Overridden by subclasses to substitute another object for itself during encoding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/replacementObject(for:)-2l8ox
func (o_ Object) ReplacementObjectForCoder(coder unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("replacementObjectForCoder:"), coder)
	return rv
}
// Overridden by subclasses to substitute another object for itself during keyed archiving. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/replacementObject(for:)-60vwc
func (o_ Object) ReplacementObjectForKeyedArchiver(archiver unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("replacementObjectForKeyedArchiver:"), archiver)
	return rv
}
// Overridden by subclasses to substitute another object for itself during archiving. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/replacementObject(for:)-8ih2x
func (o_ Object) ReplacementObjectForArchiver(archiver unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("replacementObjectForArchiver:"), archiver)
	return rv
}
// Called to determine if the specified uniform type identifier should be shown in the save panel. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/saveOptions(_:shouldShowUTType:)
func (o_ Object) SaveOptionsShouldShowUTType(saveOptions unsafe.Pointer, utType string) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("saveOptions:shouldShowUTType:"), saveOptions, utType)
	return rv
}
// Given an object specifier, returns the specified object or objects in the receiving container. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingValue(for:)
func (o_ Object) ScriptingValueForSpecifier(objectSpecifier unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("scriptingValueForSpecifier:"), objectSpecifier)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/sessionDriver:didNegotiateAndReturnError:
func (o_ Object) SessionDriverDidNegotiateAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("sessionDriver:didNegotiateAndReturnError:"), sender, outError)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/sessionDriver:didPullAndReturnError:
func (o_ Object) SessionDriverDidPullAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("sessionDriver:didPullAndReturnError:"), sender, outError)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/sessionDriver:didPushAndReturnError:
func (o_ Object) SessionDriverDidPushAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("sessionDriver:didPushAndReturnError:"), sender, outError)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/sessionDriver:didReceiveSyncAlertAndReturnError:
func (o_ Object) SessionDriverDidReceiveSyncAlertAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("sessionDriver:didReceiveSyncAlertAndReturnError:"), sender, outError)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/sessionDriver:didRegisterClientAndReturnError:
func (o_ Object) SessionDriverDidRegisterClientAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("sessionDriver:didRegisterClientAndReturnError:"), sender, outError)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/sessionDriver:willFinishSessionAndReturnError:
func (o_ Object) SessionDriverWillFinishSessionAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("sessionDriver:willFinishSessionAndReturnError:"), sender, outError)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/sessionDriver:willNegotiateAndReturnError:
func (o_ Object) SessionDriverWillNegotiateAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("sessionDriver:willNegotiateAndReturnError:"), sender, outError)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/sessionDriver:willPullAndReturnError:
func (o_ Object) SessionDriverWillPullAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("sessionDriver:willPullAndReturnError:"), sender, outError)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/sessionDriver:willPushAndReturnError:
func (o_ Object) SessionDriverWillPushAndReturnError(sender unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("sessionDriver:willPushAndReturnError:"), sender, outError)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/sessionDriverDidCancelSession:
func (o_ Object) SessionDriverDidCancelSession(sender unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("sessionDriverDidCancelSession:"), sender)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/sessionDriverDidFinishSession:
func (o_ Object) SessionDriverDidFinishSession(sender unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("sessionDriverDidFinishSession:"), sender)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/sessionDriverWillCancelSession:
func (o_ Object) SessionDriverWillCancelSession(sender unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("sessionDriverWillCancelSession:"), sender)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setSharedObservers(_:)
func (o_ Object) SetSharedObservers(sharedObservers unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSharedObservers:"), sharedObservers)
}
// Sets the property of the receiver specified by a given key to a given value. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setValue(_:forKey:)
func (o_ Object) SetValueForKey(value objc.ID, key string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setValue:forKey:"), value, key)
}
// Sets the value for the property identified by a given key path to a given value. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setValue(_:forKeyPath:)
func (o_ Object) SetValueForKeyPath(value objc.ID, keyPath string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setValue:forKeyPath:"), value, keyPath)
}
// Sets properties of the receiver with values from a given dictionary, using its keys to identify the properties. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setValuesForKeys(_:)
func (o_ Object) SetValuesForKeysWithDictionary(keyedValues unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setValuesForKeysWithDictionary:"), keyedValues)
}
// Allows the delegate to specify which device is its preferred. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setupPanel(_:determineBestDeviceOfA:orB:)
func (o_ Object) SetupPanelDetermineBestDeviceOfAOrB(aPanel unsafe.Pointer, deviceA unsafe.Pointer, device unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("setupPanel:determineBestDeviceOfA:orB:"), aPanel, deviceA, device)
	return rv
}
// This delegate method allows the delegate to determine if the media inserted in the device is suitable for whatever operation is to be performed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setupPanel(_:deviceContainsSuitableMedia:promptString:)
func (o_ Object) SetupPanelDeviceContainsSuitableMediaPromptString(aPanel unsafe.Pointer, device unsafe.Pointer, prompt string) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setupPanel:deviceContainsSuitableMedia:promptString:"), aPanel, device, prompt)
	return rv
}
// Allows the delegate to determine if device can be used as a target. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setupPanel(_:deviceCouldBeTarget:)
func (o_ Object) SetupPanelDeviceCouldBeTarget(aPanel unsafe.Pointer, device unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setupPanel:deviceCouldBeTarget:"), aPanel, device)
	return rv
}
// Sent by the default notification center when the device selection in the panel has changed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setupPanelDeviceSelectionChanged(_:)
func (o_ Object) SetupPanelDeviceSelectionChanged(aNotification unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setupPanelDeviceSelectionChanged:"), aNotification)
}
// This delegate method allows the delegate to control how media reservations are handled. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setupPanelShouldHandleMediaReservations(_:)
func (o_ Object) SetupPanelShouldHandleMediaReservations(aPanel unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setupPanelShouldHandleMediaReservations:"), aPanel)
	return rv
}
// Sent to the delegate to determine whether the action should be enabled. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/shouldEnableAction(for:identifier:)
func (o_ Object) ShouldEnableActionForPersonIdentifier(person unsafe.Pointer, identifier string) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("shouldEnableActionForPerson:identifier:"), person, identifier)
	return rv
}
// Writes the specified rows to the specified pasteboard. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/tableView:writeRows:toPasteboard:
func (o_ Object) TableViewWriteRowsToPasteboard(tableView unsafe.Pointer, rows unsafe.Pointer, pboard unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("tableView:writeRows:toPasteboard:"), tableView, rows, pboard)
	return rv
}
// Sent to the delegate to request the title of the menu item for the action. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/title(for:identifier:)
func (o_ Object) TitleForPersonIdentifier(person unsafe.Pointer, identifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("titleForPerson:identifier:"), person, identifier)
	return rv
}
// Returns the mode mask corresponding to the expected font panel mode. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/validModesForFontPanel:
func (o_ Object) ValidModesForFontPanel(fontPanel unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("validModesForFontPanel:"), fontPanel)
	return rv
}
// Implemented to override the default action of enabling or disabling a specific menu item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/validateMenuItem:
func (o_ Object) ValidateMenuItem(menuItem unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("validateMenuItem:"), menuItem)
	return rv
}
// If this method is implemented and returns , NSToolbar will disable ; returning causes to be enabled. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/validateToolbarItem:
func (o_ Object) ValidateToolbarItem(item unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("validateToolbarItem:"), item)
	return rv
}
// Returns the value for the property identified by a given key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(forKey:)
func (o_ Object) ValueForKey(key string) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("valueForKey:"), key)
	return rv
}
// Returns the value for the derived property identified by a given key path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(forKeyPath:)
func (o_ Object) ValueForKeyPath(keyPath string) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("valueForKeyPath:"), keyPath)
	return rv
}
// Returns the tool tip string to be displayed due to the cursor pausing at location within the tool tip rectangle identified by in the view . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/view:stringForToolTip:point:userData:
func (o_ Object) ViewStringForToolTipPointUserData(view unsafe.Pointer, tag unsafe.Pointer, point unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("view:stringForToolTip:point:userData:"), view, tag, point, data)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/workflowController(_:didError:)
func (o_ Object) WorkflowControllerDidError(controller unsafe.Pointer, error unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("workflowController:didError:"), controller, error)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/workflowController(_:didRun:)
func (o_ Object) WorkflowControllerDidRunAction(controller unsafe.Pointer, action unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("workflowController:didRunAction:"), controller, action)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/workflowController(_:willRun:)
func (o_ Object) WorkflowControllerWillRunAction(controller unsafe.Pointer, action unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("workflowController:willRunAction:"), controller, action)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/workflowControllerDidRun(_:)
func (o_ Object) WorkflowControllerDidRun(controller unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("workflowControllerDidRun:"), controller)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/workflowControllerDidStop(_:)
func (o_ Object) WorkflowControllerDidStop(controller unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("workflowControllerDidStop:"), controller)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/workflowControllerWillRun(_:)
func (o_ Object) WorkflowControllerWillRun(controller unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("workflowControllerWillRun:"), controller)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/workflowControllerWillStop(_:)
func (o_ Object) WorkflowControllerWillStop(controller unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("workflowControllerWillStop:"), controller)
}
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/write:length:
func (o_ Object) WriteLength(data unsafe.Pointer, length unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("write:length:"), data, length)
	return rv
}

