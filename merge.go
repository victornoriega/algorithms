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
    array, err := readFile("randomList1kk.txt")
    if err != nil { panic(err) }

    fmt.Println(len(array))
    start := time.Now()
    mergeSort(array)
    elapsed := time.Since(start)
	  log.Printf("Merge took %dms", elapsed.Nanoseconds()/1000000)
}

func mergeSort(a []int) {
  if len(a) < 2 {
        return
    }
  var s = make([]int, len(a)/2+1) // scratch space for merge step
  mid := len(a) / 2
  mergeSort(a[:mid])
  mergeSort(a[mid:])
  if a[mid-1] <= a[mid] {
      return
  }
  // merge step, with the copy-half optimization
  copy(s, a[:mid])
  l, r := 0, mid
  for i := 0; ; i++ {
      if s[l] <= a[r] {
          a[i] = s[l]
          l++
          if l == mid {
              break
          }
      } else {
          a[i] = a[r]
          r++
          if r == len(a) {
              copy(a[i+1:], s[l:mid])
              break
          }
      }
  }
  return
}
