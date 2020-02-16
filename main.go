package main

import (
  "fmt"
  "flag"
  "os/exec"
  "io/ioutil"

  "compiler/frontend"
  "compiler/backend"
)

func main() {
  filePtr := flag.String("file", "main.fa", "")

  flag.Parse()

  fmt.Printf("Starting... \n")

  data, err := ioutil.ReadFile(*filePtr)
  if err != nil {
    fmt.Printf("[Error] %v \n", err)
    return
  }

  root, worked := frontend.GenerateSytnaxTree(data)
  if !worked {
    return
  }

  code := backend.OptimizeAndGenerateCode(root)
  fmt.Printf("Code: '%s' \n", code)

  err = ioutil.WriteFile("tmp.c", []byte(code), 0644)
  if err != nil {
    fmt.Printf("Error writing to tmp file \n")
    return
  }

  cmd := exec.Command("gcc", "-O3", "-o", "tmp.out", "tmp.c")
  err = cmd.Run()
  if err != nil {
    fmt.Printf("Error running gcc \n")
    fmt.Printf("%v \n", err)
    return
  }
}
