package api

type Context struct {
	UserId       int
	ClientId     []int
	PartnerIds   []int
	TraceId      string
	InnerRequest string
	Vars         map[string]string
	Body         []byte
	With         []string
	Permissions  []string
	Roles        []string
}
