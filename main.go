package main

import (
  "coloso-queue/queue"
)
func main() {
  newEntry := queue.JSONEntry{
    FetchType: "probando",
    FetchID: "1234233452",
  }

  queue.AddEntry(newEntry)
}
