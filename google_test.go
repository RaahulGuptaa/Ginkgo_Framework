package main

import (
	"context"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// This is the entry point of the test
func TestChrome(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chrome Search Suite")
}

var _ = Describe("Chrome", func() {
	It("should search for 'Apple device' on Google", func() {
		// Launch options
		opts := append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),       // false to show browser
			chromedp.Flag("disable-gpu", true),     // Disable GPU hardware acceleration (useful in headless mode)
			chromedp.Flag("no-sandbox", true),      // Disable the sandbox (required for certain environments like Docker)
			chromedp.Flag("start-maximized", true), // Start the browser maximized
		)

		// Start Chrome instance
		allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
		defer cancel()

		ctx, cancel := chromedp.NewContext(allocCtx)
		defer cancel()

		// // Set timeout
		ctx, cancel = context.WithTimeout(ctx, 20*time.Second)
		defer cancel()

		// Run browser automation
		err := chromedp.Run(ctx,
			// Open Google
			chromedp.Navigate("https://www.google.com"),
			chromedp.WaitVisible(`//textarea[@name="q"]`, chromedp.BySearch), // BySearch due to Xpath for css use -> ByQuery
			// Type into search box and press Enter
			chromedp.SendKeys(`//textarea[@name="q"]`, "Apple device", chromedp.BySearch),
			chromedp.Click(`//input[@name="btnK"]`, chromedp.BySearch),
			chromedp.Sleep(10*time.Second), // wait for results to load
		)

		Expect(err).To(BeNil())
	})
})
