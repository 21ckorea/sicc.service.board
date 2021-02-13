package main

import (
  "fmt"
  "time"
  )

func main() {
  for {
    fmt.Println("현재 시각 : ", time.Now().Format("2006-01-02 15:04:05"))
  }
}
