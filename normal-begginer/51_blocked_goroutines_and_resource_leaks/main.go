package main


//from Rob Pike
func First(query string, replicas ...Search) (Result <-chan) {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

//first variant - channel with LEN = len(replicas)
func First(query string, replicas ...Search) Result {
	c := make(chan Result,len(replicas))
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

//second - default block
func First(query string, replicas ...Search) Result {
	c := make(chan Result,1)
	searchReplica := func(i int) {
		select {
		case c <- replicas[i](query):
		default:
		}
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

//third - special cancellation channel
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	done := make(chan struct{})
	defer close(done)
	searchReplica := func(i int) {
		select {
		case c <- replicas[i](query):
		case <- done:
		}
	}
	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}


func main() {
	
}
