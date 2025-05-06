package main

import (
	"context"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Entry point for running Ginkgo tests
func TestChrome(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chrome Search Suite")
}

var _ = Describe("Chrome", func() {
	var (
		ctx      context.Context
		cancel   context.CancelFunc
		allocCtx context.Context
	)

	BeforeEach(func() {
		// Browser launch options
		opts := append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
			chromedp.Flag("disable-gpu", true),
			chromedp.Flag("no-sandbox", true),
			chromedp.Flag("start-maximized", true),
		)

		// Setup Chrome instance
		allocCtx, cancel = chromedp.NewExecAllocator(context.Background(), opts...)
		ctx, cancel = chromedp.NewContext(allocCtx)
		ctx, cancel = context.WithTimeout(ctx, 20*time.Second)
	})

	AfterEach(func() {
		// Ensure to cancel context after each test
		cancel()
	})

	It("should search for 'Apple device' on Google", func() {
		err := chromedp.Run(ctx,
			chromedp.Navigate("https://www.google.com"),
			chromedp.WaitVisible(`//textarea[@name="q"]`, chromedp.BySearch),              // BySearch due to Xpath for css use -> ByQuer
			chromedp.SendKeys(`//textarea[@name="q"]`, "Apple device", chromedp.BySearch), // BySearch due to Xpath for css use -> ByQuer
			chromedp.Click(`//input[@name="btnK"]`, chromedp.BySearch),                    // BySearch due to Xpath for css use -> ByQuer
			chromedp.Sleep(10*time.Second),
		)
		Expect(err).To(BeNil())
	})
})
