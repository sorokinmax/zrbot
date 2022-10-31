package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/cdproto/page"

	"github.com/chromedp/chromedp"
)

func dailyReport() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", true),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("window-size", "1024,768"),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	var res string
	err := chromedp.Run(ctx, submit(cfg.Zrbot.ZabbixRootURL, `//input[@name="name"]`, cfg.Zrbot.Login, `//input[@name="password"]`, cfg.Zrbot.Password, &res))
	if err != nil {
		log.Fatal(err)
	}

	// capture pdf
	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(cfg.Reportlinks.Daily, &buf)); err != nil {
		log.Fatal(err)
	}
	chromedp.Sleep(5 * time.Second)
	if err := ioutil.WriteFile("dailyReport.pdf", buf, 0o644); err != nil {
		log.Fatal(err)
	}

	SendMail(cfg.SMTP.Host, cfg.SMTP.Port, cfg.SMTP.User, cfg.SMTP.Pass, cfg.SMTP.From, cfg.Zrbot.SendReportsTo, "", "Zabbix daily report", "", "./dailyReport.pdf")
}

func weeklyReport() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", true),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("window-size", "1024,768"),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	var res string
	err := chromedp.Run(ctx, submit(cfg.Zrbot.ZabbixRootURL, `//input[@name="name"]`, cfg.Zrbot.Login, `//input[@name="password"]`, cfg.Zrbot.Password, &res))
	if err != nil {
		log.Fatal(err)
	}

	// capture pdf
	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(cfg.Reportlinks.Weekly, &buf)); err != nil {
		log.Fatal(err)
	}
	chromedp.Sleep(5 * time.Second)
	if err := ioutil.WriteFile("weeklyReport.pdf", buf, 0o644); err != nil {
		log.Fatal(err)
	}

	SendMail(cfg.SMTP.Host, cfg.SMTP.Port, cfg.SMTP.User, cfg.SMTP.Pass, cfg.SMTP.From, cfg.Zrbot.SendReportsTo, "", "Zabbix weekly report", "", "./weeklyReport.pdf")
}

// print a specific pdf page.
func printToPDF(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Sleep(5 * time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				WithPreferCSSPageSize(true).
				WithPaperWidth(pixels2inches(1024)).
				WithPaperHeight(pixels2inches(768)).
				Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}

func submit(urlstr, uInput, user string, pInput, password string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(uInput),
		chromedp.SendKeys(uInput, user),
		chromedp.SendKeys(pInput, password),
		chromedp.Click(`#enter`, chromedp.NodeVisible),
		chromedp.Sleep(5 * time.Second),
	}
}

func pixels2inches(value int64) float64 {
	return float64(value) * 0.0104166667
}
