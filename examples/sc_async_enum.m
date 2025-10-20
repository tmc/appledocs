#import <Foundation/Foundation.h>
#import <ScreenCaptureKit/ScreenCaptureKit.h>

int main(int argc, const char * argv[]) {
    @autoreleasepool {
        NSLog(@"=== SCShareableContent Async Enumeration Test ===\n");

        // Create semaphore to wait for async completion
        dispatch_semaphore_t sem = dispatch_semaphore_create(0);
        __block SCShareableContent *shareableContent = nil;
        __block NSError *error = nil;

        NSLog(@"Requesting shareable content asynchronously...");

        // Call the async API
        [SCShareableContent getShareableContentWithCompletionHandler:^(SCShareableContent *content, NSError *err) {
            NSLog(@"✓ Completion handler called!");

            if (err) {
                error = [err retain];
                NSLog(@"  Error: %@", err.localizedDescription);
            } else {
                shareableContent = [content retain];
                NSLog(@"  Success!");
            }

            dispatch_semaphore_signal(sem);
        }];

        NSLog(@"Waiting for completion...");
        dispatch_semaphore_wait(sem, DISPATCH_TIME_FOREVER);

        if (error) {
            NSLog(@"\n❌ Error: %@", error.localizedDescription);
            return 1;
        }

        if (!shareableContent) {
            NSLog(@"\n❌ No shareable content returned");
            return 1;
        }

        // Display results
        NSLog(@"\n=== Results ===");
        NSLog(@"✓ Displays: %lu", (unsigned long)shareableContent.displays.count);
        for (SCDisplay *display in shareableContent.displays) {
            NSLog(@"  - Display %u: %lu x %lu",
                  display.displayID,
                  (unsigned long)display.width,
                  (unsigned long)display.height);
        }

        NSLog(@"\n✓ Windows: %lu", (unsigned long)shareableContent.windows.count);
        NSLog(@"  (showing first 5)");
        NSUInteger count = 0;
        for (SCWindow *window in shareableContent.windows) {
            if (count++ >= 5) break;
            NSLog(@"  - Window %u: %@ (app: %@)",
                  window.windowID,
                  window.title ?: @"(no title)",
                  window.owningApplication.applicationName);
        }

        NSLog(@"\n✓ Running Applications: %lu", (unsigned long)shareableContent.applications.count);
        NSLog(@"  (showing first 5)");
        count = 0;
        for (SCRunningApplication *app in shareableContent.applications) {
            if (count++ >= 5) break;
            NSLog(@"  - %@ (pid: %d)",
                  app.applicationName,
                  app.processID);
        }

        [shareableContent release];

        NSLog(@"\n=== ✅ Async Enumeration Test PASSED ===");
    }
    return 0;
}
