package main 

import (
  "shell/shellwords"
)

func main() {
  args, err := shellparse.Parse("./foo --bar=baz")
}
