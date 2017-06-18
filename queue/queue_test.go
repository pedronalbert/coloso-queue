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
var fetchIDTesting string

func setup() {
  queueTesting = queue.New(regionTesting)
  fetchIDTesting = strconv.Itoa(rand.Intn(100))
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
  pos, err := queueTesting.Enqueue(queue.Entry{
    FetchType: "summoner",
    FetchID: fetchIDTesting,
  })

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

  if entry.FetchID != fetchIDTesting {
    t.Fatalf("Entry in redis not match\nexpected FetchID: %s got FetchID %s", fetchIDTesting, entry.FetchID)
  }

  //Check can't allow same entry
  pos, err = queueTesting.Enqueue(queue.Entry{
    FetchType: "summoner",
    FetchID: fetchIDTesting,
  })

  if err == nil {
    t.Fatalf("Queue is allowing duplicated entries")
  }
}
