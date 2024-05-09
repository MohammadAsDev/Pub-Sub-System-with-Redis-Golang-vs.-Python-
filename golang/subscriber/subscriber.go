package subscriber

type Subscriber interface {
	Subscribe() (chan string, chan error)
	Id() int
}
