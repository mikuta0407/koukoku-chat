package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {

	/* TELNETS接続周り */
	config := tls.Config{Certificates: []tls.Certificate{}, InsecureSkipVerify: false}
	conn, err := tls.Dial("tcp", "koukoku.shadan.open.ad.jp:992", &config)
	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}
	defer conn.Close()
	log.Println("client: connected to: ", conn.RemoteAddr())

	// nobody送ってチャットのみにする
	fmt.Fprintln(conn, "nobody")

	// 白黒モード
	cmode := true
	if len(os.Args) > 1 && os.Args[1] == "mono" {
		cmode = false
	}
	// 白黒モード時の削除用
	prefixRe := regexp.MustCompile("[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))")

	/* UI周り */
	app := tview.NewApplication()

	textView := tview.NewTextView()
	textView.SetTitle("公告チャット")
	textView.SetBorder(true)

	inputField := tview.NewInputField()
	inputField.SetLabel("Chat: ")
	inputField.SetTitle("inputField").
		SetBorder(true)

	/* チャット送信 */
	inputField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			fmt.Fprintln(conn, inputField.GetText()+"\r\n")
			inputField.SetText("")

			return nil
		}
		return event

	})

	textView.SetChangedFunc(func() {
		textView.ScrollToEnd()
		app.Draw()
	})

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		AddItem(textView, 0, 1, false).
		AddItem(inputField, 3, 0, true)

	/* チャット受取 */
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			line := scanner.Text()
			if cmode {
				fmt.Fprintln(textView, line)
			} else {
				fmt.Fprintln(textView, prefixRe.ReplaceAllString(line, ""))
			}
		}
	}()

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}

	wg.Wait()

}
