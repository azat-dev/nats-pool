package nats_pool

import (
	"sync"
	. "testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"time"
)

func TestPool(t *T) {
	size := 10
	pool, err := New("nats://localhost:4222", size)
	require.Nil(t, err)

	var wg sync.WaitGroup

	//start listeners
	for i := 0; i < size/2; i++ {

		subject := "test_" + strconv.Itoa(i)

		wg.Add(1)
		go func(subject string) {
			defer wg.Done()

			conn, err := pool.Get()
			assert.Nil(t, err)
			defer pool.Put(conn)

			subscription, err := conn.SubscribeSync(subject)
			assert.Nil(t, err)
			defer subscription.Unsubscribe()

			msg, err := subscription.NextMsg(20 * time.Second)
			assert.Nil(t, err, "Error on receive message")
			assert.NotNil(t, msg, "Received message should be not nil")

			if msg != nil {
				assert.Equal(t, string(msg.Data), "hello")
			}

		}(subject)
	}

	time.Sleep(1 * time.Second)

	//start publishers
	for i := 0; i < size/2; i++ {

		subject := "test_" + strconv.Itoa(i)

		wg.Add(1)
		go func(subject string) {

			defer wg.Done()

			conn, err := pool.Get()
			assert.Nil(t, err)
			defer pool.Put(conn)

			err = conn.Publish(subject, []byte("hello"))
			assert.Nil(t, err)

		}(subject)
	}

	wg.Wait()
	assert.Equal(t, size, len(pool.pool))

	pool.Empty()
	assert.Equal(t, 0, len(pool.pool))
}
