package queue_test

import "testing"
import "coloso-queue/queue"

var regionTesting = "lan"
var queueTesting *queue.Queue

func init() {
  queueTesting = queue.New(regionTesting)
}

func TestNewQueue(t *testing.T) {
  queueNew := queue.New(regionTesting)
  queueNameExpected := "coloso_queue_" + regionTesting

  if queueNameExpected != queueNew.Name {
    t.Fatalf("expected Queue.Name: %s got Queue.Name %s", queueNameExpected, queueNew.Name)
  }
}
