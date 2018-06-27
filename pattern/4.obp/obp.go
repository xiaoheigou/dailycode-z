package obp

type Object *string

type Pool chan *Object

func New(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- new(Object)
	}

	return &p
}

//usage
// p := pool.New(2)

// select {
// case obj := <-p:
// 	obj.Do( /*...*/ )

// 	p <- obj
// default:
// 	// No more objects left — retry later or fail
// 	return
// }
