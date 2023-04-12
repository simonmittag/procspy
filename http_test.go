package procspy

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

const concurrency = 16

func TestConnections(t *testing.T) {
	t.Log("\ntestprocspy-starts-open-conns-")

	go func() {
		for i := 0; i < 25; i++ {
			time.Sleep(time.Millisecond * 100)
			//p := fmt.Sprintf("%02d-", spy())
			//t.Log(p)
		}
	}()

	for i := 0; i < concurrency; i++ {
		time.Sleep(time.Millisecond * 100)
		go func() {
			c := initHTTPClient()
			for i := 0; i < 100; i++ {
				res, _ := c.Get("http://jsonplaceholder.typicode.com/todos/1")
				_, _ = ioutil.ReadAll(res.Body)
				res.Body.Close()
			}
		}()
	}

	time.Sleep(time.Second * 2)

	want := concurrency
	got := spy()
	if want != got {
		t.Errorf("did not find expected amount of conns, want %v got %v", want, got)
	} else {
		t.Logf("normal. found %v conns", got)
	}

	t.Log("testprocspy-ends\n")
}

func initHTTPClient() http.Client {
	c := http.Client{
		Transport: &http.Transport{
			MaxConnsPerHost: 1,
			IdleConnTimeout: time.Duration(1 * time.Second),
		},
	}
	return c
}

func spy() int {
	pid := os.Getpid()
	cs, _ := Connections(true)
	d := 0
	for c := cs.Next(); c != nil; c = cs.Next() {
		if c.PID == uint(pid) && c.RemotePort == 80 {
			d++
		}
	}
	return d
}
