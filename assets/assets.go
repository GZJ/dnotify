package assets

import (
	"embed"
	"fmt"
	"io"
	"os"
)

//go:embed alert.png info.png
var content embed.FS

// GetIconPath returns the path to the icon file.
// If iconPath is empty, it extracts the default icon to a temp file.
func GetIconPath(iconPath, defaultIconName string) (string, error) {
	if iconPath != "" {
		return iconPath, nil
	}

	defaultIcon, err := content.Open(defaultIconName)
	if err != nil {
		return "", fmt.Errorf("error opening default icon: %w", err)
	}
	defer defaultIcon.Close()

	tempIcon, err := os.CreateTemp("", "default_icon_*.png")
	if err != nil {
		return "", fmt.Errorf("error creating temporary file for default icon: %w", err)
	}
	defer tempIcon.Close()

	_, err = io.Copy(tempIcon, defaultIcon)
	if err != nil {
		return "", fmt.Errorf("error copying default icon to temporary file: %w", err)
	}

	return tempIcon.Name(), nil
}
