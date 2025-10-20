#import <Foundation/Foundation.h>
#import <ScreenCaptureKit/ScreenCaptureKit.h>
#import <CoreGraphics/CoreGraphics.h>

int main(int argc, const char * argv[]) {
    @autoreleasepool {
        NSLog(@"=== SCScreenshotManager Test ===\n");

        // First, get shareable content to get a display
        dispatch_semaphore_t contentSem = dispatch_semaphore_create(0);
        __block SCShareableContent *shareableContent = nil;
        __block NSError *contentError = nil;

        NSLog(@"Getting shareable content...");
        [SCShareableContent getShareableContentWithCompletionHandler:^(SCShareableContent *content, NSError *error) {
            if (error) {
                contentError = error;
            } else {
                shareableContent = [content retain];
            }
            dispatch_semaphore_signal(contentSem);
        }];

        // Wait for content
        dispatch_semaphore_wait(contentSem, DISPATCH_TIME_FOREVER);

        if (contentError) {
            NSLog(@"Error getting content: %@", contentError.localizedDescription);
            return 1;
        }

        if (!shareableContent || shareableContent.displays.count == 0) {
            NSLog(@"No displays found");
            return 1;
        }

        NSLog(@"Found %lu displays", (unsigned long)shareableContent.displays.count);

        // Get first display
        SCDisplay *display = shareableContent.displays.firstObject;
        NSLog(@"Using display: %u (%lu x %lu)", display.displayID, (unsigned long)display.width, (unsigned long)display.height);

        // Create content filter
        NSArray *emptyArray = @[];
        SCContentFilter *filter = [[SCContentFilter alloc] initWithDisplay:display
                                                       excludingApplications:emptyArray
                                                          exceptingWindows:emptyArray];

        NSLog(@"Created content filter");

        // Capture screenshot
        dispatch_semaphore_t captureSem = dispatch_semaphore_create(0);
        __block CGImageRef capturedImage = NULL;
        __block NSError *captureError = nil;

        NSLog(@"Requesting screenshot...");
        [SCScreenshotManager captureImageWithContentFilter:filter
                                              configuration:nil
                                          completionHandler:^(CGImageRef sampleBuffer, NSError *error) {
            NSLog(@"Completion handler called!");
            NSLog(@"  Image: %p", sampleBuffer);
            NSLog(@"  Error: %@", error);

            if (error) {
                captureError = [error retain];
            } else if (sampleBuffer) {
                capturedImage = CGImageRetain(sampleBuffer);
            }
            dispatch_semaphore_signal(captureSem);
        }];

        NSLog(@"Waiting for screenshot...");
        long result = dispatch_semaphore_wait(captureSem, dispatch_time(DISPATCH_TIME_NOW, 10 * NSEC_PER_SEC));

        if (result != 0) {
            NSLog(@"Timeout waiting for screenshot");
            return 1;
        }

        if (captureError) {
            NSLog(@"Capture error: %@", captureError.localizedDescription);
            return 1;
        }

        if (!capturedImage) {
            NSLog(@"No image returned");
            return 1;
        }

        size_t width = CGImageGetWidth(capturedImage);
        size_t height = CGImageGetHeight(capturedImage);

        NSLog(@"✓ Screenshot captured: %zu x %zu pixels", width, height);

        CGImageRelease(capturedImage);
        [filter release];
        [shareableContent release];

        NSLog(@"\n=== Test PASSED ===");
    }
    return 0;
}
