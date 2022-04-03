package Http

type HttpReq struct {
	IsConnection bool //是否是长链接
	Url          string
	Body         []byte
}
