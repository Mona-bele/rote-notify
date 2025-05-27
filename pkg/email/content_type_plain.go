package email

type ContentTypePlain struct {
	ContentType
}

type IContentTypePlain interface {
	Body() string
}

func NewContentTypePlain() *ContentTypePlain {
	return &ContentTypePlain{}
}

func (p *ContentTypePlain) Body() string {
	return "text/plain"
}
