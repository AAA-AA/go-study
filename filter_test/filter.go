package filtertest

type Request interface{}

type Response interface{}

type Filter interface {
	Process(data Request) (Response, error)
}
