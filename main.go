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
  // Get the Command Line Argument for which file to compile
  filePtr := flag.String("file", "main.fa", "")
  flag.Parse()

  fmt.Printf("Starting... \n")
  fmt.Printf("Loading File '%s' \n", *filePtr)

  // Reads in the Data from the given File
  data, err := ioutil.ReadFile(*filePtr)
  if err != nil {
    fmt.Printf("[Error] %v \n", err)
    return
  }

  // Uses the Data to create an AST using the Frontend
  root, worked := frontend.GenerateSytnaxTree(data)
  if !worked {
    return
  }

  // Use the backend to optimze and then generate the Code
  code := backend.OptimizeAndGenerateCode(root)
  fmt.Printf("Code: '%s' \n", code)

  // Writes the Code to a temporary C file, which is needed for compilation
  err = ioutil.WriteFile("tmp.c", []byte(code), 0644)
  if err != nil {
    fmt.Printf("Error writing to tmp file \n")
    return
  }

  // Starts gcc with the fitting compiler options to turn that temporary C file into a program
  cmd := exec.Command("gcc", "-O3", "-o", "tmp.out", "tmp.c")
  err = cmd.Run()
  if err != nil {
    fmt.Printf("Error running gcc \n")
    fmt.Printf("%v \n", err)
    return
  }
}
