//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for MusicPlayerApplicationController


// Changes the contents of the media items in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerApplicationController/perform(queueTransaction:completionHandler:)
func (m_ MusicPlayerApplicationController) PerformQueueTransactionCompletionHandler(queueTransaction unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("performQueueTransaction:completionHandler:"), queueTransaction, completionHandler)
}

// iOS-only properties





