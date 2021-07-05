package mdtowiki

import "strings"

// MdToWiki ... md >>> pukiwiki
func MdToWiki(markdown string) string {
	markdownLines := append(strings.Split(strings.NewReplacer("\r\n", "\n", "\r", "\n").Replace(markdown), "\n"), "")

	return markdownLines[0]
}
