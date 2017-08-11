package main

import (
	"bufio"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	// "os"
)

func main() {
	res, err := http.Get(
		"http://www-cs-faculty.stanford.edu/~knuth/sgb-words.txt")

	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	buckets := make([]int, 200)

	for scanner.Scan() {
		n := hashBucket(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets[:12])
	// words := make(map[string]string)

	// sc := bufio.NewScanner(res.Body)
	// sc.Split(bufio.ScanWords)

	// for sc.Scan() {
	// 	words[sc.Text()] = ""
	// }
	// if err := sc.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "reading input:", err)
	// }

	// i := 0
	// for k := range words {
	// 	fmt.Println(k)
	// 	if i == 200 {
	// 		break
	// 	}
	// 	i++
	// }

	// bs, _ := ioutil.ReadAll(res.Body)
	// str := string(bs)
	// fmt.Println(str)
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
