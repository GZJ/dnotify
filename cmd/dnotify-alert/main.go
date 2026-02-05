package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gen2brain/beeep"
	"github.com/gzj/dnotify/assets"
)

func main() {
	title := flag.String("title", "Default Title", "Notification title")
	body := flag.String("body", "", "Notification message body (optional, can use args instead)")
	iconPath := flag.String("icon", "", "Path to custom icon (optional)")

	flag.Parse()

	// Use args as body if no -body flag provided
	message := *body
	if message == "" && len(flag.Args()) > 0 {
		message = strings.Join(flag.Args(), " ")
	}
	if message == "" {
		message = "Default Message Body"
	}

	icon, err := assets.GetIconPath(*iconPath, "alert.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = beeep.Notify(*title, message, icon)
	if err != nil {
		fmt.Println("Error sending notification:", err)
		os.Exit(1)
	}

	fmt.Println("Notification sent successfully!")
}
