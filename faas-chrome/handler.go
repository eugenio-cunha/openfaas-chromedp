package function

import (
	"fmt"
	"log"
	"context"

	"github.com/chromedp/chromedp"
)

func Handle(req []byte) string {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
		chromedp.NoFirstRun,
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(string(req)),
		chromedp.Evaluate(`document.documentElement.innerHTML;`, &res),
	)

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%v", res)
}
