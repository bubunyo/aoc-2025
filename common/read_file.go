package common

import (
	"bufio"
	"iter"
	"log"
	"os"
	"path/filepath"
)

func IterateFileContent(fp string) iter.Seq[string] {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(filepath.Join(path, fp))
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	return func(yield func(string) bool) {
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				break
			}
		}

		defer file.Close()

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
