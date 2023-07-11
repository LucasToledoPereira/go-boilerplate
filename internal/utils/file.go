package utils

import (
	"path/filepath"
	"strings"
)

func IsImage(filename string) bool {
	extension := strings.ToLower(filepath.Ext(filename))
	switch extension {
	case ".png", ".webp", ".gif", ".jpeg", ".jpg":
		return true
	default:
		return false
	}
}
