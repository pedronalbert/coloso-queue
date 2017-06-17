package main

func main() {
  regions := [3]string{"lan", "las", "na"}
  queues := make(map[string]*Queue)

  for _, region := range regions {
    queues[region] = NewQueue(region)
  }

  queues["lan"].Enqueue(QueueEntry{
    FetchType: "summoner",
    FetchID: "123456",
  })
}
