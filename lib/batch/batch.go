package batch

import (
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	res = make([]user, 0, n)
	var i int64
	var mut sync.Mutex
	var group errgroup.Group

	group.SetLimit(int(pool))

	for ; i < n; i++ {

		userId := i

		group.Go(func() error {
			user := getOne(userId)

			mut.Lock()
			res = append(res, user)
			mut.Unlock()

			return nil
		})
	}

	group.Wait()

	return res
}
