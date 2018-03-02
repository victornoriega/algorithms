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
    array, err := readFile("randomList10k.txt")
    if err != nil { panic(err) }
    fmt.Println(len(array))

    n := len(array)

    start := time.Now()
    for i := 0; i < n; i++ {
        var minIdx = i
        for j := i; j < n; j++ {
            if array[j] < array[minIdx] {
                minIdx = j
            }
        }
        array[i], array[minIdx] = array[minIdx], array[i]
    }
    elapsed := time.Since(start)
	  log.Printf("Selection took %dms", elapsed.Nanoseconds()/1000000)
}
