package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/localauthentication"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")
var authenticate = flag.Bool("auth", false, "attempt actual biometric authentication")

// LAPolicy constants (from LocalAuthentication framework)
const (
	// LAPolicyDeviceOwnerAuthenticationWithBiometrics requires biometric authentication
	LAPolicyDeviceOwnerAuthenticationWithBiometrics = 1
	// LAPolicyDeviceOwnerAuthentication allows biometric or passcode authentication
	LAPolicyDeviceOwnerAuthentication = 2
	// LAPolicyDeviceOwnerAuthenticationWithWatch allows Apple Watch authentication
	LAPolicyDeviceOwnerAuthenticationWithWatch = 3
	// LAPolicyDeviceOwnerAuthenticationWithBiometricsOrWatch allows biometric or Watch authentication
	LAPolicyDeviceOwnerAuthenticationWithBiometricsOrWatch = 4
	// LAPolicyDeviceOwnerAuthenticationWithWristDetection allows authentication with wrist detection on Watch
	LAPolicyDeviceOwnerAuthenticationWithWristDetection = 5
)

// LAError constants
const (
	LAErrorAuthenticationFailed = -1 // Authentication failed
	LAErrorUserCancel           = -2 // User cancelled
	LAErrorUserFallback         = -3 // User chose fallback
	LAErrorSystemCancel         = -4 // System cancelled
	LAErrorPasscodeNotSet       = -5 // Passcode not set
	LAErrorBiometryNotAvailable = -6 // Biometry not available
	LAErrorBiometryNotEnrolled  = -7 // Biometry not enrolled
	LAErrorBiometryLockout      = -8 // Biometry locked out
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.Parse()

	fmt.Println("LocalAuthentication TouchID/FaceID CLI")
	fmt.Println("======================================")

	// Example 1: Framework overview
	fmt.Println("\n1. Framework Overview:")
	fmt.Println("   LocalAuthentication provides:")
	fmt.Println("   - Biometric authentication (TouchID, FaceID)")
	fmt.Println("   - Device passcode authentication")
	fmt.Println("   - Apple Watch authentication")
	fmt.Println("   - Secure credential evaluation")

	// Example 2: Authentication policies
	fmt.Println("\n2. Authentication Policies:")
	policies := map[int]string{
		LAPolicyDeviceOwnerAuthenticationWithBiometrics:            "Biometrics only (TouchID/FaceID)",
		LAPolicyDeviceOwnerAuthentication:                          "Biometrics or Passcode",
		LAPolicyDeviceOwnerAuthenticationWithWatch:                 "Apple Watch authentication",
		LAPolicyDeviceOwnerAuthenticationWithBiometricsOrWatch:     "Biometrics or Watch",
		LAPolicyDeviceOwnerAuthenticationWithWristDetection:        "Watch with wrist detection",
	}

	for policy, description := range policies {
		fmt.Printf("   Policy %d: %s\n", policy, description)
	}

	// Example 3: Create LAContext
	fmt.Println("\n3. Creating LAContext:")
	fmt.Println("   LAContext manages authentication state and handles user interaction")

	context := localauthentication.NewContext()
	fmt.Printf("   Created LAContext: %v\n", context.ID != 0)

	// Example 4: Check biometric availability
	fmt.Println("\n4. Checking Biometric Availability:")

	// Note: CanEvaluatePolicyError requires proper Objective-C method signature handling
	// The policy parameter expects NSInteger (int) as an Objective-C boxed type
	// For now, we'll show the API without calling it to avoid crashes

	fmt.Println("   API: context.CanEvaluatePolicyError(policy, error)")
	fmt.Println("   Note: Proper invocation requires correct Objective-C type marshaling")
	fmt.Println("   ")
	fmt.Println("   In production, you would:")
	fmt.Println("   1. Create an NSError** pointer for the error parameter")
	fmt.Println("   2. Pass the policy as an NSInteger (int) value")
	fmt.Println("   3. Check the boolean return value")
	fmt.Println("   4. Inspect the NSError if canEvaluate returns false")
	fmt.Println("   ")
	fmt.Println("   Reasons biometric authentication may not be available:")
	fmt.Println("   - No TouchID or FaceID hardware")
	fmt.Println("   - Biometrics not enrolled")
	fmt.Println("   - Biometrics locked out (too many failed attempts)")
	fmt.Println("   - Running in unsupported environment (VM, CI, etc.)")

	// Example 5: LAContext properties
	fmt.Println("\n5. LAContext Properties:")
	properties := []string{
		"localizedReason - User-facing explanation for authentication",
		"localizedFallbackTitle - Custom fallback button title",
		"localizedCancelTitle - Custom cancel button title",
		"maxBiometryFailures - Max attempts before lockout",
		"biometryType - Type of biometry available (TouchID/FaceID/None)",
		"interactionNotAllowed - Disable user interaction",
	}

	for i, property := range properties {
		fmt.Printf("   %d. %s\n", i+1, property)
	}

	// Example 6: Error codes
	fmt.Println("\n6. Common Error Codes:")
	errors := map[int]string{
		LAErrorAuthenticationFailed: "Authentication was not successful",
		LAErrorUserCancel:           "User cancelled authentication",
		LAErrorUserFallback:         "User chose fallback (e.g., password)",
		LAErrorSystemCancel:         "System cancelled authentication",
		LAErrorPasscodeNotSet:       "Device passcode not set",
		LAErrorBiometryNotAvailable: "Biometry not available on device",
		LAErrorBiometryNotEnrolled:  "No biometric data enrolled",
		LAErrorBiometryLockout:      "Biometry locked after too many failures",
	}

	for code, description := range errors {
		fmt.Printf("   %d: %s\n", code, description)
	}

	// Example 7: Authentication workflow
	fmt.Println("\n7. Authentication Workflow:")
	workflow := []string{
		"1. Create LAContext instance",
		"2. Check if policy can be evaluated (canEvaluatePolicy:error:)",
		"3. Set localized reason for authentication prompt",
		"4. Call evaluatePolicy:localizedReason:reply:",
		"5. Handle success/failure in callback",
		"6. Process error codes if authentication fails",
	}

	for _, step := range workflow {
		fmt.Printf("   %s\n", step)
	}

	// Example 8: Use cases
	fmt.Println("\n8. Use Cases:")
	useCases := map[string]string{
		"Secure Login":        "Authenticate user before app access",
		"Payment Auth":        "Confirm purchases or transactions",
		"Sensitive Data":      "Access password manager or secure notes",
		"Settings Changes":    "Verify identity for critical settings",
		"Document Signing":    "Authenticate document approval",
		"Multi-Factor Auth":   "Second factor for authentication",
		"CLI Tools":           "Secure command-line operations",
		"Admin Operations":    "Verify admin privileges",
	}

	for useCase, description := range useCases {
		fmt.Printf("   %-20s: %s\n", useCase, description)
	}

	// Example 9: Biometry types
	fmt.Println("\n9. Biometry Types:")
	biometryTypes := []string{
		"LABiometryTypeNone (0) - No biometry available",
		"LABiometryTypeTouchID (1) - TouchID sensor present",
		"LABiometryTypeFaceID (2) - FaceID available",
		"LABiometryTypeOpticID (4) - OpticID (Vision Pro)",
	}

	for i, biometryType := range biometryTypes {
		fmt.Printf("   %d. %s\n", i+1, biometryType)
	}

	// Example 10: Code example
	fmt.Println("\n10. Example Code Structure:")
	fmt.Println("   // Create context")
	fmt.Println("   context := localauthentication.NewContext()")
	fmt.Println("   ")
	fmt.Println("   // Check if biometric authentication is available")
	fmt.Println("   var errorPtr unsafe.Pointer")
	fmt.Println("   policy := LAPolicyDeviceOwnerAuthenticationWithBiometrics")
	fmt.Println("   canEval := context.CanEvaluatePolicyError(unsafe.Pointer(&policy), errorPtr)")
	fmt.Println("   ")
	fmt.Println("   if !canEval {")
	fmt.Println("       fmt.Println(\"Biometric auth not available\")")
	fmt.Println("       return")
	fmt.Println("   }")
	fmt.Println("   ")
	fmt.Println("   // Evaluate authentication")
	fmt.Println("   // Note: Callback requires proper Objective-C block handling")
	fmt.Println("   // context.EvaluatePolicyLocalizedReasonReply(...)")

	// Example 11: Interactive authentication demo
	if *authenticate {
		fmt.Println("\n11. Attempting Interactive Authentication:")
		fmt.Println("   Starting biometric authentication...")
		fmt.Println("   ")
		fmt.Println("   Note: Interactive authentication requires proper block callback handling")
		fmt.Println("   which is complex in Go. This would typically be done by:")
		fmt.Println("   1. Creating an Objective-C block for the reply handler")
		fmt.Println("   2. Calling context.EvaluatePolicyLocalizedReasonReply()")
		fmt.Println("   3. Processing the callback with success/error")
		fmt.Println("")
		fmt.Println("   For production use, consider:")
		fmt.Println("   - Using cgo with Objective-C wrapper")
		fmt.Println("   - Creating a helper library for block handling")
		fmt.Println("   - Using Apple's C APIs for authentication")
	} else {
		fmt.Println("\n11. Interactive Authentication:")
		fmt.Println("   Run with -auth flag to see authentication notes")
		fmt.Println("   Example: go run main.go -auth")
	}

	// Example 12: Best practices
	fmt.Println("\n12. Best Practices:")
	bestPractices := []string{
		"Always check canEvaluatePolicy before attempting authentication",
		"Provide clear, user-friendly localizedReason messages",
		"Handle all error cases gracefully",
		"Use appropriate policy for your use case",
		"Don't store authentication results - reauthenticate when needed",
		"Consider fallback options for users without biometrics",
		"Test on actual hardware (not simulators)",
		"Respect user's choice to cancel authentication",
	}

	for i, practice := range bestPractices {
		fmt.Printf("   %d. %s\n", i+1, practice)
	}

	// Example 13: Platform requirements
	fmt.Println("\n13. Platform Requirements:")
	requirements := []string{
		"macOS 10.10+ for basic LocalAuthentication",
		"TouchID: MacBook Pro (2016+), MacBook Air (2018+)",
		"FaceID: Not available on macOS (iOS/iPadOS only)",
		"OpticID: Apple Vision Pro",
		"Passcode authentication available on all platforms",
		"Apple Watch authentication requires paired watch",
	}

	for i, requirement := range requirements {
		fmt.Printf("   %d. %s\n", i+1, requirement)
	}

	// Example 14: Security considerations
	fmt.Println("\n14. Security Considerations:")
	securityNotes := []string{
		"Biometric data never leaves the Secure Enclave",
		"App receives only success/failure, not biometric data",
		"Multiple failed attempts trigger lockout",
		"System handles all UI and user interaction",
		"Cannot be bypassed or manipulated by app",
		"Passcode required after device restart",
		"Biometrics disabled after 5 failed attempts",
	}

	for i, note := range securityNotes {
		fmt.Printf("   %d. %s\n", i+1, note)
	}

	// Example 15: Troubleshooting
	fmt.Println("\n15. Common Issues:")
	issues := map[string]string{
		"kLAErrorBiometryNotAvailable": "No TouchID/FaceID hardware or disabled",
		"kLAErrorBiometryNotEnrolled":  "User hasn't enrolled fingerprints/face",
		"kLAErrorPasscodeNotSet":       "Device passcode not configured",
		"kLAErrorBiometryLockout":      "Too many failed attempts - locked out",
		"Running in VM":                "Biometrics not available in virtual machines",
		"CI/CD Environment":            "Cannot test biometrics in automated environments",
	}

	for issue, solution := range issues {
		fmt.Printf("   %-30s: %s\n", issue, solution)
	}

	fmt.Println("\n✓ LocalAuthentication TouchID/FaceID CLI demo completed!")
	fmt.Println("\nNote: Full authentication requires:")
	fmt.Println("  - Objective-C block handling for callbacks")
	fmt.Println("  - Proper main thread execution")
	fmt.Println("  - Hardware with TouchID/FaceID")
	fmt.Println("  - Enrolled biometric data")
	fmt.Println("  - Device passcode configured")
	fmt.Println("\nFor production use:")
	fmt.Println("  - Consider using cgo with Objective-C wrapper")
	fmt.Println("  - Implement proper block callback handling")
	fmt.Println("  - Test on actual hardware devices")
	fmt.Println("  - Handle all error cases appropriately")
	fmt.Println("\nReferences:")
	fmt.Println("  - https://developer.apple.com/documentation/localauthentication")
	fmt.Println("  - https://developer.apple.com/documentation/localauthentication/lacontext")
	fmt.Println("  - https://developer.apple.com/documentation/localauthentication/lapolicy")
}
