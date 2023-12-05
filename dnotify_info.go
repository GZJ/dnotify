package main

import (
	"embed"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/gen2brain/beeep"
)

//go:embed assets/info.png
var content embed.FS

func main() {
	title := flag.String("title", "Default Title", "Notification title")
	body := flag.String("body", "Default Message Body", "Notification message body")
	iconPath := flag.String("icon", "", "Path to custom icon (optional)")

	flag.Parse()

	if *iconPath == "" {
		defaultIcon, err := content.Open("assets/info.png")
		if err != nil {
			fmt.Println("Error opening default icon:", err)
			os.Exit(1)
		}
		defer defaultIcon.Close()

		tempIcon, err := os.CreateTemp("", "default_icon_*.png")
		if err != nil {
			fmt.Println("Error creating temporary file for default icon:", err)
			os.Exit(1)
		}
		defer tempIcon.Close()

		_, err = io.Copy(tempIcon, defaultIcon)
		if err != nil {
			fmt.Println("Error copying default icon to temporary file:", err)
			os.Exit(1)
		}

		*iconPath = tempIcon.Name()
	}

	err := beeep.Notify(*title, *body, *iconPath)
	if err != nil {
		fmt.Println("Error sending notification:", err)
		os.Exit(1)
	}

	fmt.Println("Notification sent successfully!")
}
