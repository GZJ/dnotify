package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gen2brain/beeep"
	"github.com/gzj/dnotify/assets"
)

func main() {
	title := flag.String("title", "Default Title", "Notification title")
	body := flag.String("body", "Default Message Body", "Notification message body")
	iconPath := flag.String("icon", "", "Path to custom icon (optional)")

	flag.Parse()

	icon, err := assets.GetIconPath(*iconPath, "info.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = beeep.Notify(*title, *body, icon)
	if err != nil {
		fmt.Println("Error sending notification:", err)
		os.Exit(1)
	}

	fmt.Println("Notification sent successfully!")
}
