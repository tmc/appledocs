// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PViewControllerPresentationAnimator is the NSViewControllerPresentationAnimator protocol interface.
//
// A set of methods that let you define animations to play when transitioning between two view controllers.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSViewControllerPresentationAnimator
type PViewControllerPresentationAnimator interface {
	// Required methods
	AnimateDismissalOfViewControllerFromViewController(viewController IViewController, fromViewController IViewController)/* debug [protocol_interface/required_method]: AnimateDismissalOfViewControllerFromViewController */
	AnimatePresentationOfViewControllerFromViewController(viewController IViewController, fromViewController IViewController)/* debug [protocol_interface/required_method]: AnimatePresentationOfViewControllerFromViewController */
}
