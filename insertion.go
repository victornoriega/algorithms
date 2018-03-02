package main

import (
  "fmt"
  "io/ioutil"
  "strconv"
  "strings"
  "time"
  "log"
)

// It would be better for such a function to return error, instead of handling
// it on their own.
func readFile(fname string) (nums []int, err error) {
  b, err := ioutil.ReadFile(fname)
  if err != nil { return nil, err }

  lines := strings.Split(string(b), "\r\n")
  // Assign cap to avoid resize on every append.
  nums = make([]int, 0, len(lines))

  for _,l := range lines {
      // Empty line occurs at the end of the file when we use Split.
      if len(l) == 0 { continue }
      // Atoi better suits the job when we know exactly what we're dealing
      // with. Scanf is the more general option.
      n, err := strconv.Atoi(l)
      if err != nil { return nil, err }
      nums = append(nums, n)
  }

  return nums, nil
}

func main() {
    array, err := readFile("randomList9k.txt")
    if err != nil { panic(err) }
    fmt.Println(len(array))

    lenght := len(array)

    start := time.Now()
    for i := 1; i < lenght; i++ {
      j := i
      // swap the values until it is the smallest
      for j > 0 && array[j - 1] > array[j] {
        array[j - 1], array[j] = array[j], array[j - 1]
        // values swapped so we need to check for the following array element
        j = j - 1
      }
    }
    elapsed := time.Since(start)
	  log.Printf("Insertion took %dms", elapsed.Nanoseconds()/1000000)
}
