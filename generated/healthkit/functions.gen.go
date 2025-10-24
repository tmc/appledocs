// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

/* debug [functions.gen.go]: Generating 11 functions for HealthKit */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// HealthKit Functions (11 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_HKAppleSleepingBreathingDisturbancesMinimumQuantityForClassification func(HKAppleSleepingBreathingDisturbancesClassification) unsafe.Pointer
	_HKAppleSleepingBreathingDisturbancesClassificationForQuantity func(unsafe.Pointer) unsafe.Pointer
	_HKAppleWalkingSteadinessClassificationForQuantity func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_HKAppleWalkingSteadinessMaximumQuantityForClassification func(HKAppleWalkingSteadinessClassification) unsafe.Pointer
	_HKAppleWalkingSteadinessMinimumQuantityForClassification func(HKAppleWalkingSteadinessClassification) unsafe.Pointer
	_HKCategoryValueSleepAnalysisAsleepValues func() unsafe.Pointer
	_HKMaximumScoreForGAD7AssessmentRisk func(HKGAD7AssessmentRisk) int
	_HKMaximumScoreForPHQ9AssessmentRisk func(HKPHQ9AssessmentRisk) int
	_HKMinimumScoreForGAD7AssessmentRisk func(HKGAD7AssessmentRisk) int
	_HKMinimumScoreForPHQ9AssessmentRisk func(HKPHQ9AssessmentRisk) int
	_HKStateOfMindValenceClassificationForValence func(float64) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_HKAppleSleepingBreathingDisturbancesMinimumQuantityForClassification, lib, "HKAppleSleepingBreathingDisturbancesMinimumQuantityForClassification")
	tryRegister(&_HKAppleSleepingBreathingDisturbancesClassificationForQuantity, lib, "HKAppleSleepingBreathingDisturbancesClassificationForQuantity")
	tryRegister(&_HKAppleWalkingSteadinessClassificationForQuantity, lib, "HKAppleWalkingSteadinessClassificationForQuantity")
	tryRegister(&_HKAppleWalkingSteadinessMaximumQuantityForClassification, lib, "HKAppleWalkingSteadinessMaximumQuantityForClassification")
	tryRegister(&_HKAppleWalkingSteadinessMinimumQuantityForClassification, lib, "HKAppleWalkingSteadinessMinimumQuantityForClassification")
	tryRegister(&_HKCategoryValueSleepAnalysisAsleepValues, lib, "HKCategoryValueSleepAnalysisAsleepValues")
	tryRegister(&_HKMaximumScoreForGAD7AssessmentRisk, lib, "HKMaximumScoreForGAD7AssessmentRisk")
	tryRegister(&_HKMaximumScoreForPHQ9AssessmentRisk, lib, "HKMaximumScoreForPHQ9AssessmentRisk")
	tryRegister(&_HKMinimumScoreForGAD7AssessmentRisk, lib, "HKMinimumScoreForGAD7AssessmentRisk")
	tryRegister(&_HKMinimumScoreForPHQ9AssessmentRisk, lib, "HKMinimumScoreForPHQ9AssessmentRisk")
	tryRegister(&_HKStateOfMindValenceClassificationForValence, lib, "HKStateOfMindValenceClassificationForValence")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// HKAppleSleepingBreathingDisturbancesMinimumQuantityForClassification is a HealthKit function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleSleepingBreathingDisturbancesClassification/minimum
func HKAppleSleepingBreathingDisturbancesMinimumQuantityForClassification(classification HKAppleSleepingBreathingDisturbancesClassification) unsafe.Pointer {
	return _HKAppleSleepingBreathingDisturbancesMinimumQuantityForClassification(classification)
}/* debug [functions.gen.go/function]: HKAppleSleepingBreathingDisturbancesMinimumQuantityForClassification */

// HKAppleSleepingBreathingDisturbancesClassificationForQuantity is a HealthKit function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleSleepingBreathingDisturbancesClassificationForQuantity
func HKAppleSleepingBreathingDisturbancesClassificationForQuantity(value unsafe.Pointer) unsafe.Pointer {
	return _HKAppleSleepingBreathingDisturbancesClassificationForQuantity(value)
}/* debug [functions.gen.go/function]: HKAppleSleepingBreathingDisturbancesClassificationForQuantity */

// Provides a classification for a score that measures the steadiness of the user’s gait.
//
// Added in macOS 13.0.
// Provides a classification for a score that measures the steadiness of the user’s gait.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleWalkingSteadinessClassificationForQuantity
func HKAppleWalkingSteadinessClassificationForQuantity(value unsafe.Pointer, classificationOut unsafe.Pointer, errorOut unsafe.Pointer) bool {
	return _HKAppleWalkingSteadinessClassificationForQuantity(value, classificationOut, errorOut)
}/* debug [functions.gen.go/function]: HKAppleWalkingSteadinessClassificationForQuantity */

// Returns the maximum score for the steadiness of the user’s gait based on the provided classification.
//
// Added in macOS 13.0.
// Returns the maximum score for the steadiness of the user’s gait based on the provided classification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleWalkingSteadinessMaximumQuantityForClassification
func HKAppleWalkingSteadinessMaximumQuantityForClassification(classification HKAppleWalkingSteadinessClassification) unsafe.Pointer {
	return _HKAppleWalkingSteadinessMaximumQuantityForClassification(classification)
}/* debug [functions.gen.go/function]: HKAppleWalkingSteadinessMaximumQuantityForClassification */

// Returns the minimum score for the steadiness of the user’s gait based on the provided classification.
//
// Added in macOS 13.0.
// Returns the minimum score for the steadiness of the user’s gait based on the provided classification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleWalkingSteadinessMinimumQuantityForClassification
func HKAppleWalkingSteadinessMinimumQuantityForClassification(classification HKAppleWalkingSteadinessClassification) unsafe.Pointer {
	return _HKAppleWalkingSteadinessMinimumQuantityForClassification(classification)
}/* debug [functions.gen.go/function]: HKAppleWalkingSteadinessMinimumQuantityForClassification */

// HKCategoryValueSleepAnalysisAsleepValues is a HealthKit function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSleepAnalysisAsleepValues
func HKCategoryValueSleepAnalysisAsleepValues() unsafe.Pointer {
	return _HKCategoryValueSleepAnalysisAsleepValues()
}/* debug [functions.gen.go/function]: HKCategoryValueSleepAnalysisAsleepValues */

// Returns the upper bound of the score range for the given GAD-7 risk classification.

// Returns the upper bound of the score range for the given GAD-7 risk classification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMaximumScoreForGAD7AssessmentRisk
func HKMaximumScoreForGAD7AssessmentRisk(risk HKGAD7AssessmentRisk) int {
	return _HKMaximumScoreForGAD7AssessmentRisk(risk)
}/* debug [functions.gen.go/function]: HKMaximumScoreForGAD7AssessmentRisk */

// Returns the upper bound of the score range for the given PHQ-9 risk classification.

// Returns the upper bound of the score range for the given PHQ-9 risk classification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMaximumScoreForPHQ9AssessmentRisk
func HKMaximumScoreForPHQ9AssessmentRisk(risk HKPHQ9AssessmentRisk) int {
	return _HKMaximumScoreForPHQ9AssessmentRisk(risk)
}/* debug [functions.gen.go/function]: HKMaximumScoreForPHQ9AssessmentRisk */

// Returns the lower bound of the score range for the given GAD-7 risk classification.

// Returns the lower bound of the score range for the given GAD-7 risk classification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMinimumScoreForGAD7AssessmentRisk
func HKMinimumScoreForGAD7AssessmentRisk(risk HKGAD7AssessmentRisk) int {
	return _HKMinimumScoreForGAD7AssessmentRisk(risk)
}/* debug [functions.gen.go/function]: HKMinimumScoreForGAD7AssessmentRisk */

// Returns the lower bound of the score range for the given PHQ-9 risk classification.

// Returns the lower bound of the score range for the given PHQ-9 risk classification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMinimumScoreForPHQ9AssessmentRisk
func HKMinimumScoreForPHQ9AssessmentRisk(risk HKPHQ9AssessmentRisk) int {
	return _HKMinimumScoreForPHQ9AssessmentRisk(risk)
}/* debug [functions.gen.go/function]: HKMinimumScoreForPHQ9AssessmentRisk */

// HKStateOfMindValenceClassificationForValence is a HealthKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMindValenceClassificationForValence
func HKStateOfMindValenceClassificationForValence(valence float64) unsafe.Pointer {
	return _HKStateOfMindValenceClassificationForValence(valence)
}/* debug [functions.gen.go/function]: HKStateOfMindValenceClassificationForValence */




