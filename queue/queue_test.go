package queue_test

import "testing"
import "math/rand"
import "os"
import "strconv"
import "encoding/json"
import "coloso-queue/queue"
import "coloso-queue/clients/redis"

var regionTesting = "test"
var queueTesting *queue.Queue
var entryTesting queue.Entry

func setup() {
  queueTesting = queue.New(regionTesting)
  entryTesting = queue.Entry{
    FetchID: strconv.Itoa(rand.Intn(100)),
    FetchType: "summoner",
  }
}

func clearRedis() {
  redis.Client.Del(queueTesting.Name).Result()
}

func shutdown() {
  clearRedis()
}

func TestMain(m *testing.M) {
  setup()
  code := m.Run()
  shutdown()
  os.Exit(code)
}

func TestNewQueue(t *testing.T) {
  queueNew := queue.New(regionTesting)
  queueNameExpected := "coloso_queue_" + regionTesting

  if queueNameExpected != queueNew.Name {
    t.Fatalf("expected Queue.Name: %s got Queue.Name %s", queueNameExpected, queueNew.Name)
  }
}

func TestEnqueue(t *testing.T) {
  var entry queue.Entry
  pos, err := queueTesting.Enqueue(entryTesting)

  if err != nil {
    t.Fatalf("Can't enqueue the entry\n%s", err)
  }

  // Check in redis
  redisIndex := pos - 1
  redisEntry, err := redis.Client.LIndex(queueTesting.Name, redisIndex).Result()

  if err != nil {
    t.Fatalf("Can't get entry from redis, index: %d\nerror: %s",redisIndex, err)
  }

  // Parse entry to JSON
  err = json.Unmarshal([]byte(redisEntry), &entry)

  if err != nil {
    t.Fatalf("Can't decode redis entry to JSON\n%s", err)
  }

  if entry != entryTesting {
    t.Fatalf("Entry in redis doesn't match\nexpected Entry: %+v \ngot Entry: %+v", entryTesting, entry)
  }

  //Check can't allow same entry
  pos, err = queueTesting.Enqueue(entryTesting)

  if err == nil {
    t.Fatalf("Queue is allowing duplicated entries")
  }
}

func TestGetAllEntries(t *testing.T) {
  entries, err := queueTesting.GetAllEntries()

  if err != nil {
    t.Fatalf("Can't get all entries from queue\n error: %s", err)
  }

  if len(entries) == 0 {
    t.Fatalf("Entries slice is empty")
  }
}
