package mdtowiki

import (
	"regexp"
	"strings"
)

// MdToWiki ... md >>> pukiwiki
func MdToWiki(markdown string) string {
	markdownLines := append(strings.Split(strings.NewReplacer("\r\n", "\n", "\r", "\n").Replace(markdown), "\n"), "")

	// 2行以上の改行を削除
	flg := false
	wikiLines := []string{"[[FrontPage]]", "", "#setlinebreak(true)", ""}
	for i, v := range markdownLines {
		if !(flg && v == "") {
			wikiLines = append(wikiLines, markdownLines[i])
			flg = v == ""
		}
	}

	// 各行を変換
	for i := range wikiLines {
		l := wikiLines[i]
		l = regexp.MustCompile(`^#\s`).ReplaceAllString(l, "* ")
		l = regexp.MustCompile(`^##\s`).ReplaceAllString(l, "** ")
		l = regexp.MustCompile(`^###\s`).ReplaceAllString(l, "*** ")
		l = regexp.MustCompile(`^-\s`).ReplaceAllString(l, "- ")
		l = regexp.MustCompile(`^  -\s`).ReplaceAllString(l, "-- ")
		l = regexp.MustCompile(`^    -\s`).ReplaceAllString(l, "--- ")
		l = regexp.MustCompile(`\*\*(.*)\*\*`).ReplaceAllString(l, "''$1''")
		l = regexp.MustCompile(`~~(.*)~~`).ReplaceAllString(l, "%%$1%%")
		l = regexp.MustCompile(`\[(.*)\]\((.*)\)`).ReplaceAllString(l, "[[$1:$2]]")
		wikiLines[i] = l
	}

	// 結合
	return strings.Join(wikiLines, "\n")
}
