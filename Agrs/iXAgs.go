package agrs

type iXargs interface {
	Commit()
}
type AgrsAgent struct {
	Ip string

	ProceMax int
}
type xargsImp struct {
}

func (a xargsImp) Cmmd() {

}
