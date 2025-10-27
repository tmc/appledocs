// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PWritingToolsCoordinatorDelegate is the NSWritingToolsCoordinatorDelegate protocol interface.
//
// An interface that you use to manage interactions between Writing Tools   and your custom text view.
//
// Availability:
//   - macOS 15.2+
//
// See: doc://com.apple.appkit/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol
type PWritingToolsCoordinatorDelegate interface {
	// Required methods
	WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, textAnimation WritingToolsCoordinatorTextAnimation, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, textAnimation WritingToolsCoordinatorTextAnimation, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, replacementText foundation.foundation.INSAttributedString, reason WritingToolsCoordinatorTextReplacementReason, animationParameters IWritingToolsCoordinatorAnimationParameters, completion unsafe.Pointer)
	WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	WritingToolsCoordinatorRequestsContextsForScopeCompletion(writingToolsCoordinator IWritingToolsCoordinator, scope WritingToolsCoordinatorContextScope, completion unsafe.Pointer)
	WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, rect corefoundation.CGRect, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, textAnimation WritingToolsCoordinatorTextAnimation, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	WritingToolsCoordinatorSelectRangesInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, ranges []foundation.Value, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	// Optional methods
	WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	HasWritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion() bool
	WritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion(writingToolsCoordinator IWritingToolsCoordinator, point corefoundation.CGPoint, completion unsafe.Pointer)
	HasWritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion() bool
	WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	HasWritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion() bool
	WritingToolsCoordinatorWillChangeToStateCompletion(writingToolsCoordinator IWritingToolsCoordinator, newState WritingToolsCoordinatorState, completion unsafe.Pointer)
	HasWritingToolsCoordinatorWillChangeToStateCompletion() bool
}

// WritingToolsCoordinatorDelegate is a delegate implementation builder for the PWritingToolsCoordinatorDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type WritingToolsCoordinatorDelegate struct {
	_WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion func(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	_WritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion func(writingToolsCoordinator IWritingToolsCoordinator, point corefoundation.CGPoint, completion unsafe.Pointer)
	_WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion func(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	_WritingToolsCoordinatorWillChangeToStateCompletion func(writingToolsCoordinator IWritingToolsCoordinator, newState WritingToolsCoordinatorState, completion unsafe.Pointer)
	_WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion func(writingToolsCoordinator IWritingToolsCoordinator, textAnimation WritingToolsCoordinatorTextAnimation, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	_WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion func(writingToolsCoordinator IWritingToolsCoordinator, textAnimation WritingToolsCoordinatorTextAnimation, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	_WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion func(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, replacementText foundation.foundation.INSAttributedString, reason WritingToolsCoordinatorTextReplacementReason, animationParameters IWritingToolsCoordinatorAnimationParameters, completion unsafe.Pointer)
	_WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion func(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	_WritingToolsCoordinatorRequestsContextsForScopeCompletion func(writingToolsCoordinator IWritingToolsCoordinator, scope WritingToolsCoordinatorContextScope, completion unsafe.Pointer)
	_WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion func(writingToolsCoordinator IWritingToolsCoordinator, rect corefoundation.CGRect, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	_WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion func(writingToolsCoordinator IWritingToolsCoordinator, textAnimation WritingToolsCoordinatorTextAnimation, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	_WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion func(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
	_WritingToolsCoordinatorSelectRangesInContextCompletion func(writingToolsCoordinator IWritingToolsCoordinator, ranges []foundation.Value, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)
}

// SetWritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion sets the handler for the WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion delegate method.
//
// Asks the delegate to provide a decoration view for the specified range of text.
func (d *WritingToolsCoordinatorDelegate) SetWritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion(f func(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)) {
	d._WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion = f
}

// SetWritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion sets the handler for the WritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion delegate method.
//
// Asks the delegate to provide the location of the character at the   specified point in your view’s coordinate system.
func (d *WritingToolsCoordinatorDelegate) SetWritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion(f func(writingToolsCoordinator IWritingToolsCoordinator, point corefoundation.CGPoint, completion unsafe.Pointer)) {
	d._WritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion = f
}

// SetWritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion sets the handler for the WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion delegate method.
//
// Asks the delegate to divide the specified range of text into the separate   containers that render that text.
func (d *WritingToolsCoordinatorDelegate) SetWritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion(f func(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)) {
	d._WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion = f
}

// SetWritingToolsCoordinatorWillChangeToStateCompletion sets the handler for the WritingToolsCoordinatorWillChangeToStateCompletion delegate method.
//
// Notifies your delegate of relevant state changes when Writing Tools   is running in your view.
func (d *WritingToolsCoordinatorDelegate) SetWritingToolsCoordinatorWillChangeToStateCompletion(f func(writingToolsCoordinator IWritingToolsCoordinator, newState WritingToolsCoordinatorState, completion unsafe.Pointer)) {
	d._WritingToolsCoordinatorWillChangeToStateCompletion = f
}

// SetWritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion sets the handler for the WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion delegate method.
//
// Asks the delegate to clean up any state related to the specified   Writing Tools animation.
func (d *WritingToolsCoordinatorDelegate) SetWritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion(f func(writingToolsCoordinator IWritingToolsCoordinator, textAnimation WritingToolsCoordinatorTextAnimation, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)) {
	d._WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion = f
}

// SetWritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion sets the handler for the WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion delegate method.
//
// Prepare for animations for the content that Writing Tools is evaluating.
func (d *WritingToolsCoordinatorDelegate) SetWritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion(f func(writingToolsCoordinator IWritingToolsCoordinator, textAnimation WritingToolsCoordinatorTextAnimation, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)) {
	d._WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion = f
}

// SetWritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion sets the handler for the WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion delegate method.
//
// Tells the delegate that there are text changes to incorporate into the view.
func (d *WritingToolsCoordinatorDelegate) SetWritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion(f func(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, replacementText foundation.foundation.INSAttributedString, reason WritingToolsCoordinatorTextReplacementReason, animationParameters IWritingToolsCoordinatorAnimationParameters, completion unsafe.Pointer)) {
	d._WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion = f
}

// SetWritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion sets the handler for the WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion delegate method.
//
// Asks the delegate to provide the bounding paths for the specified   text in your view.
func (d *WritingToolsCoordinatorDelegate) SetWritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion(f func(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)) {
	d._WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion = f
}

// SetWritingToolsCoordinatorRequestsContextsForScopeCompletion sets the handler for the WritingToolsCoordinatorRequestsContextsForScopeCompletion delegate method.
//
// Asks your delegate to provide the text to evaluate during the Writing Tools   operation.
func (d *WritingToolsCoordinatorDelegate) SetWritingToolsCoordinatorRequestsContextsForScopeCompletion(f func(writingToolsCoordinator IWritingToolsCoordinator, scope WritingToolsCoordinatorContextScope, completion unsafe.Pointer)) {
	d._WritingToolsCoordinatorRequestsContextsForScopeCompletion = f
}

// SetWritingToolsCoordinatorRequestsPreviewForRectInContextCompletion sets the handler for the WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion delegate method.
//
// Asks the delegate for a preview image and layout information for the   specified text.
func (d *WritingToolsCoordinatorDelegate) SetWritingToolsCoordinatorRequestsPreviewForRectInContextCompletion(f func(writingToolsCoordinator IWritingToolsCoordinator, rect corefoundation.CGRect, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)) {
	d._WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion = f
}

// SetWritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion sets the handler for the WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion delegate method.
//
// Asks the delegate for a preview image and layout information for the   specified text.
func (d *WritingToolsCoordinatorDelegate) SetWritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion(f func(writingToolsCoordinator IWritingToolsCoordinator, textAnimation WritingToolsCoordinatorTextAnimation, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)) {
	d._WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion = f
}

// SetWritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion sets the handler for the WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion delegate method.
//
// Asks the delegate to provide an underline shape for the specified text   during a proofreading session.
func (d *WritingToolsCoordinatorDelegate) SetWritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion(f func(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)) {
	d._WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion = f
}

// SetWritingToolsCoordinatorSelectRangesInContextCompletion sets the handler for the WritingToolsCoordinatorSelectRangesInContextCompletion delegate method.
//
// Asks the delegate to update your view’s current text selection.
func (d *WritingToolsCoordinatorDelegate) SetWritingToolsCoordinatorSelectRangesInContextCompletion(f func(writingToolsCoordinator IWritingToolsCoordinator, ranges []foundation.Value, context IWritingToolsCoordinatorContext, completion unsafe.Pointer)) {
	d._WritingToolsCoordinatorSelectRangesInContextCompletion = f
}

// WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
func (d *WritingToolsCoordinatorDelegate) WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	if d._WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion != nil {
		d._WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion(writingToolsCoordinator, range_, context, completion)
	}
}

// HasWritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion returns true if a handler for WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion has been set.
func (d *WritingToolsCoordinatorDelegate) HasWritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion() bool {
	return d._WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion != nil
}

// WritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion implements the PWritingToolsCoordinatorDelegate interface.
func (d *WritingToolsCoordinatorDelegate) WritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion(writingToolsCoordinator IWritingToolsCoordinator, point corefoundation.CGPoint, completion unsafe.Pointer) {
	if d._WritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion != nil {
		d._WritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion(writingToolsCoordinator, point, completion)
	}
}

// HasWritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion returns true if a handler for WritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion has been set.
func (d *WritingToolsCoordinatorDelegate) HasWritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion() bool {
	return d._WritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion != nil
}

// WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
func (d *WritingToolsCoordinatorDelegate) WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	if d._WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion != nil {
		d._WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion(writingToolsCoordinator, range_, context, completion)
	}
}

// HasWritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion returns true if a handler for WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion has been set.
func (d *WritingToolsCoordinatorDelegate) HasWritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion() bool {
	return d._WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion != nil
}

// WritingToolsCoordinatorWillChangeToStateCompletion implements the PWritingToolsCoordinatorDelegate interface.
func (d *WritingToolsCoordinatorDelegate) WritingToolsCoordinatorWillChangeToStateCompletion(writingToolsCoordinator IWritingToolsCoordinator, newState WritingToolsCoordinatorState, completion unsafe.Pointer) {
	if d._WritingToolsCoordinatorWillChangeToStateCompletion != nil {
		d._WritingToolsCoordinatorWillChangeToStateCompletion(writingToolsCoordinator, newState, completion)
	}
}

// HasWritingToolsCoordinatorWillChangeToStateCompletion returns true if a handler for WritingToolsCoordinatorWillChangeToStateCompletion has been set.
func (d *WritingToolsCoordinatorDelegate) HasWritingToolsCoordinatorWillChangeToStateCompletion() bool {
	return d._WritingToolsCoordinatorWillChangeToStateCompletion != nil
}

// WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
func (d *WritingToolsCoordinatorDelegate) WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, textAnimation WritingToolsCoordinatorTextAnimation, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	if d._WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion != nil {
		d._WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion(writingToolsCoordinator, textAnimation, range_, context, completion)
	}
}

// HasWritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion returns true if a handler for WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion has been set.
func (d *WritingToolsCoordinatorDelegate) HasWritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion() bool {
	return d._WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion != nil
}

// WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
func (d *WritingToolsCoordinatorDelegate) WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, textAnimation WritingToolsCoordinatorTextAnimation, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	if d._WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion != nil {
		d._WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion(writingToolsCoordinator, textAnimation, range_, context, completion)
	}
}

// HasWritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion returns true if a handler for WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion has been set.
func (d *WritingToolsCoordinatorDelegate) HasWritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion() bool {
	return d._WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion != nil
}

// WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion implements the PWritingToolsCoordinatorDelegate interface.
func (d *WritingToolsCoordinatorDelegate) WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, replacementText foundation.foundation.INSAttributedString, reason WritingToolsCoordinatorTextReplacementReason, animationParameters IWritingToolsCoordinatorAnimationParameters, completion unsafe.Pointer) {
	if d._WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion != nil {
		d._WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion(writingToolsCoordinator, range_, context, replacementText, reason, animationParameters, completion)
	}
}

// HasWritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion returns true if a handler for WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion has been set.
func (d *WritingToolsCoordinatorDelegate) HasWritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion() bool {
	return d._WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion != nil
}

// WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
func (d *WritingToolsCoordinatorDelegate) WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	if d._WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion != nil {
		d._WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion(writingToolsCoordinator, range_, context, completion)
	}
}

// HasWritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion returns true if a handler for WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion has been set.
func (d *WritingToolsCoordinatorDelegate) HasWritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion() bool {
	return d._WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion != nil
}

// WritingToolsCoordinatorRequestsContextsForScopeCompletion implements the PWritingToolsCoordinatorDelegate interface.
func (d *WritingToolsCoordinatorDelegate) WritingToolsCoordinatorRequestsContextsForScopeCompletion(writingToolsCoordinator IWritingToolsCoordinator, scope WritingToolsCoordinatorContextScope, completion unsafe.Pointer) {
	if d._WritingToolsCoordinatorRequestsContextsForScopeCompletion != nil {
		d._WritingToolsCoordinatorRequestsContextsForScopeCompletion(writingToolsCoordinator, scope, completion)
	}
}

// HasWritingToolsCoordinatorRequestsContextsForScopeCompletion returns true if a handler for WritingToolsCoordinatorRequestsContextsForScopeCompletion has been set.
func (d *WritingToolsCoordinatorDelegate) HasWritingToolsCoordinatorRequestsContextsForScopeCompletion() bool {
	return d._WritingToolsCoordinatorRequestsContextsForScopeCompletion != nil
}

// WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
func (d *WritingToolsCoordinatorDelegate) WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, rect corefoundation.CGRect, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	if d._WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion != nil {
		d._WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion(writingToolsCoordinator, rect, context, completion)
	}
}

// HasWritingToolsCoordinatorRequestsPreviewForRectInContextCompletion returns true if a handler for WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion has been set.
func (d *WritingToolsCoordinatorDelegate) HasWritingToolsCoordinatorRequestsPreviewForRectInContextCompletion() bool {
	return d._WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion != nil
}

// WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
func (d *WritingToolsCoordinatorDelegate) WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, textAnimation WritingToolsCoordinatorTextAnimation, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	if d._WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion != nil {
		d._WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion(writingToolsCoordinator, textAnimation, range_, context, completion)
	}
}

// HasWritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion returns true if a handler for WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion has been set.
func (d *WritingToolsCoordinatorDelegate) HasWritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion() bool {
	return d._WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion != nil
}

// WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
func (d *WritingToolsCoordinatorDelegate) WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	if d._WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion != nil {
		d._WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion(writingToolsCoordinator, range_, context, completion)
	}
}

// HasWritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion returns true if a handler for WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion has been set.
func (d *WritingToolsCoordinatorDelegate) HasWritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion() bool {
	return d._WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion != nil
}

// WritingToolsCoordinatorSelectRangesInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
func (d *WritingToolsCoordinatorDelegate) WritingToolsCoordinatorSelectRangesInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, ranges []foundation.Value, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	if d._WritingToolsCoordinatorSelectRangesInContextCompletion != nil {
		d._WritingToolsCoordinatorSelectRangesInContextCompletion(writingToolsCoordinator, ranges, context, completion)
	}
}

// HasWritingToolsCoordinatorSelectRangesInContextCompletion returns true if a handler for WritingToolsCoordinatorSelectRangesInContextCompletion has been set.
func (d *WritingToolsCoordinatorDelegate) HasWritingToolsCoordinatorSelectRangesInContextCompletion() bool {
	return d._WritingToolsCoordinatorSelectRangesInContextCompletion != nil
}

// WritingToolsCoordinatorDelegateObject wraps an existing Objective-C object that conforms to the PWritingToolsCoordinatorDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type WritingToolsCoordinatorDelegateObject struct {
	objectivec.Object
}

// NewWritingToolsCoordinatorDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSWritingToolsCoordinatorDelegate protocol.
func NewWritingToolsCoordinatorDelegateObject(obj objectivec.Object) *WritingToolsCoordinatorDelegateObject {
	return &WritingToolsCoordinatorDelegateObject{obj}
}

// Make sure WritingToolsCoordinatorDelegateObject implements PWritingToolsCoordinatorDelegate.
var _ PWritingToolsCoordinatorDelegate = (*WritingToolsCoordinatorDelegateObject)(nil)

// WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
// This required method is always available on objects conforming to WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion.
func (o *WritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, textAnimation WritingToolsCoordinatorTextAnimation, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("writingToolsCoordinator:finishTextAnimation:forRange:inContext:completion:"), writingToolsCoordinator, textAnimation, range_, context, completion)
}

// WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
// This required method is always available on objects conforming to WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion.
func (o *WritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, textAnimation WritingToolsCoordinatorTextAnimation, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("writingToolsCoordinator:prepareForTextAnimation:forRange:inContext:completion:"), writingToolsCoordinator, textAnimation, range_, context, completion)
}

// WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion implements the PWritingToolsCoordinatorDelegate interface.
// This required method is always available on objects conforming to WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion.
func (o *WritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, replacementText foundation.foundation.INSAttributedString, reason WritingToolsCoordinatorTextReplacementReason, animationParameters IWritingToolsCoordinatorAnimationParameters, completion unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("writingToolsCoordinator:replaceRange:inContext:proposedText:reason:animationParameters:completion:"), writingToolsCoordinator, range_, context, replacementText, reason, animationParameters, completion)
}

// WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
// This required method is always available on objects conforming to WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion.
func (o *WritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("writingToolsCoordinator:requestsBoundingBezierPathsForRange:inContext:completion:"), writingToolsCoordinator, range_, context, completion)
}

// WritingToolsCoordinatorRequestsContextsForScopeCompletion implements the PWritingToolsCoordinatorDelegate interface.
// This required method is always available on objects conforming to WritingToolsCoordinatorRequestsContextsForScopeCompletion.
func (o *WritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorRequestsContextsForScopeCompletion(writingToolsCoordinator IWritingToolsCoordinator, scope WritingToolsCoordinatorContextScope, completion unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("writingToolsCoordinator:requestsContextsForScope:completion:"), writingToolsCoordinator, scope, completion)
}

// WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
// This required method is always available on objects conforming to WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion.
func (o *WritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, rect corefoundation.CGRect, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("writingToolsCoordinator:requestsPreviewForRect:inContext:completion:"), writingToolsCoordinator, rect, context, completion)
}

// WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
// This required method is always available on objects conforming to WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion.
func (o *WritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, textAnimation WritingToolsCoordinatorTextAnimation, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("writingToolsCoordinator:requestsPreviewForTextAnimation:ofRange:inContext:completion:"), writingToolsCoordinator, textAnimation, range_, context, completion)
}

// WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
// This required method is always available on objects conforming to WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion.
func (o *WritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("writingToolsCoordinator:requestsUnderlinePathsForRange:inContext:completion:"), writingToolsCoordinator, range_, context, completion)
}

// WritingToolsCoordinatorSelectRangesInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
// This required method is always available on objects conforming to WritingToolsCoordinatorSelectRangesInContextCompletion.
func (o *WritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorSelectRangesInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, ranges []foundation.Value, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("writingToolsCoordinator:selectRanges:inContext:completion:"), writingToolsCoordinator, ranges, context, completion)
}

// WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *WritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("writingToolsCoordinator:requestsDecorationContainerViewForRange:inContext:completion:"), writingToolsCoordinator, range_, context, completion)
}

// HasWritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion returns true; this is a placeholder for optional method checks.
func (o *WritingToolsCoordinatorDelegateObject) HasWritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// WritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion implements the PWritingToolsCoordinatorDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *WritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion(writingToolsCoordinator IWritingToolsCoordinator, point corefoundation.CGPoint, completion unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("writingToolsCoordinator:requestsRangeInContextWithIdentifierForPoint:completion:"), writingToolsCoordinator, point, completion)
}

// HasWritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion returns true; this is a placeholder for optional method checks.
func (o *WritingToolsCoordinatorDelegateObject) HasWritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion implements the PWritingToolsCoordinatorDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *WritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion(writingToolsCoordinator IWritingToolsCoordinator, range_ foundation.Range, context IWritingToolsCoordinatorContext, completion unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("writingToolsCoordinator:requestsSingleContainerSubrangesOfRange:inContext:completion:"), writingToolsCoordinator, range_, context, completion)
}

// HasWritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion returns true; this is a placeholder for optional method checks.
func (o *WritingToolsCoordinatorDelegateObject) HasWritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// WritingToolsCoordinatorWillChangeToStateCompletion implements the PWritingToolsCoordinatorDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *WritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorWillChangeToStateCompletion(writingToolsCoordinator IWritingToolsCoordinator, newState WritingToolsCoordinatorState, completion unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("writingToolsCoordinator:willChangeToState:completion:"), writingToolsCoordinator, newState, completion)
}

// HasWritingToolsCoordinatorWillChangeToStateCompletion returns true; this is a placeholder for optional method checks.
func (o *WritingToolsCoordinatorDelegateObject) HasWritingToolsCoordinatorWillChangeToStateCompletion() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
