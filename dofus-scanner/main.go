package main

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

// Fetch a URL using a real browser (bypasses bot protection)
func fetchWithBrowser(ctx context.Context, url string) (string, error) {
	var body string

	// Navigate to the page and wait for it to load
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(2*time.Second), // Wait for challenge to complete
		chromedp.OuterHTML("html", &body),
	)

	return body, err
}

func main() {
	// Parse command-line flags
	startID := flag.Int("start", 1, "Starting ID")
	endID := flag.Int("end", 1000, "Ending ID")
	delay := flag.Int("delay", 2, "Delay between requests in milliseconds")
	baseURL := flag.String("url", "https://dofocus.fr/api/items", "Base API URL")
	headless := flag.Bool("headless", true, "Run browser in headless mode")
	flag.Parse()

	fmt.Printf("Scanning IDs %d to %d\n", *startID, *endID)
	fmt.Printf("Delay: %dms between requests\n", *delay)
	fmt.Printf("Headless: %v\n\n", *headless)

	// Setup Chrome browser context
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", *headless),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create context without debug logging to hide harmless errors
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Loop through all IDs
	for id := *startID; id <= *endID; id++ {
		url := fmt.Sprintf("%s/%d?lang=fr", *baseURL, id)

		// Fetch the page
		body, err := fetchWithBrowser(ctx, url)
		if err != nil {
			fmt.Printf("ID %d: Error - %v\n", id, err)
			time.Sleep(time.Duration(*delay) * time.Millisecond)
			continue
		}

		// Display the result
		if strings.Contains(body, "Item non trouvé") {
			fmt.Printf("ID %d: Not found\n", id)
		} else if strings.Contains(body, "TigerProtect") || strings.Contains(body, "challenge") {
			fmt.Printf("ID %d: Challenge detected (may need more wait time)\n", id)
		} else {
			// Try to extract just the relevant content
			fmt.Printf("ID %d: Found - %s\n", id, body[:min(len(body), 200)])
		}

		// Wait before next request
		time.Sleep(time.Duration(*delay) * time.Millisecond)
	}

	fmt.Printf("\nDone!\n")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
