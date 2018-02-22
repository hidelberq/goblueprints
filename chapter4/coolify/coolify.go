package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	duplicated = true
	remove     = false
)

func randBool() bool {
	return rand.Intn(2) == 0
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := []byte(s.Text())
		if randBool() {
			var vI int = -1
			for i, char := range text {
				switch char {
				case 'a', 'e', 'i', 'u', 'o', 'A', 'E', 'I', 'U', 'O':
					if randBool() {
						vI = i
					}
				}
			}

			if vI >= 0 {
				switch randBool() {
				case duplicated:
					text = append(text[:vI+1], text[vI:]...)
				case remove:
					text = append(text[:vI], text[vI+1:]...)
				}
			}
		}

		fmt.Println(string(text))
	}
}
