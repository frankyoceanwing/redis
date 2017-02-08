package main

import (
	"log"
	"time"

	"github.com/chasex/redis-go-cluster"
)

func main() {

	cluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes:   []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002"},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})
	if err != nil {
		log.Fatal(err)
	}

	cluster.Do("SET", "foo", "bar")
	cluster.Do("INCR", "mycount")
	cluster.Do("LPUSH", "mylist", "foo", "bar")
	cluster.Do("HMSET", "myhash", "f1", "foo", "f2", "bar")

	mycount, err := redis.Int(cluster.Do("INCR", "mycount"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("INCR mycount\n(integer) %d\n", mycount)

	foo, err := redis.String(cluster.Do("GET", "foo"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GET foo\n%s\n", foo)

	mylist, err := redis.Strings(cluster.Do("LRANGE", "mylist", 0, -1))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("LRANGE mylist\n")
	if len(mylist) == 0 {
		log.Printf("(empty list or set)\n")
	} else {
		for i, v := range mylist {
			log.Printf("%d) \"%s\"\n", i+1, v)
		}
	}

	myhash, err := redis.StringMap(cluster.Do("HGETALL", "myhash"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("HGETALL myhash\n")
	if len(myhash) == 0 {
		log.Printf("(empty list or set)\n")
	} else {
		i := 1
		for k, v := range myhash {
			log.Printf("%d) \"%s\"\n", i, k)
			i++
			log.Printf("%d) \"%s\"\n", i, v)
			i++
		}
	}
}
