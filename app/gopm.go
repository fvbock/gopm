package gopm

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fvbock/trie"
)

type GoPM struct {
	tui         *TUI
	prefixIndex *trie.Trie
	fullIndex   *trie.Trie
	titleIndex  map[string]*Entry
}

func NewGoPMApp() (gopm *GoPM) {
	gopm = &GoPM{
		tui:         NewTUI(),
		prefixIndex: trie.NewTrie(),
		fullIndex:   trie.NewTrie(),
		titleIndex:  make(map[string]*Entry),
	}

	gopm.tui.inputField.SetChangedFunc(gopm.FindEntries)

	return
}

func (gpm *GoPM) Run() {
	if err := gpm.tui.app.SetRoot(gpm.tui.layout, true).Run(); err != nil {
		panic(err)
	}
}

func (gpm *GoPM) ScanFile(fname string) (entries []*Entry) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	newEntry := true
	var entry *Entry
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		if strings.HasPrefix(line, DELIMITER) {
			entries = append(entries, entry)
			newEntry = true
		}
		if newEntry {
			entry = &Entry{}
			newEntry = false
			continue
		}
		if len(entry.Title) == 0 {
			entry.Title = line
		} else {
			entry.Text += fmt.Sprintf("%s\n", line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func (gpm *GoPM) IndexEntries(entries []*Entry) {
	for _, entry := range entries {
		parts := strings.Fields(entry.Title)
		for _, part := range parts {
			gpm.titleIndex[strings.ToLower(part)] = entry
		}
	}
}

func (gpm *GoPM) FindEntries(term string) {
	results := []*Entry{}
	var buffer bytes.Buffer
	for key, entry := range gpm.titleIndex {
		// buffer.WriteString(fmt.Sprintf("%s <> %s\n", key, term))
		if strings.Contains(key, strings.ToLower(term)) {
			results = append(results, entry)
		}
	}
	buffer.WriteString(fmt.Sprintf("%v results\n", len(results)))
	gpm.tui.statusView.SetText(buffer.String())

	gpm.ShowEntries(results)
}

func (gpm *GoPM) ShowEntries(entries []*Entry) {
	var buffer bytes.Buffer
	for _, entry := range entries {
		// buffer seems to get full
		// fmt.Fprintf(gpm.tui.textView, fmt.Sprintf("[#ff0000]%s[white]\n%s\n", entry.Title, entry.Text))
		buffer.WriteString(fmt.Sprintf("[#ff0000]%s[white]\n%s\n", entry.Title, entry.Text))
	}
	gpm.tui.textView.SetText(buffer.String())
}

// func bufferedWrite()
