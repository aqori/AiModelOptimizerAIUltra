// cmd/aimodeloptimizeraiultra/main.go
package main

import (
"flag"
"log"
"os"

"aimodeloptimizeraiultra/internal/aimodeloptimizeraiultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := aimodeloptimizeraiultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
