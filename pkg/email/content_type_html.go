package email

type ContentTypeHtml struct {
	ContentType
}

type IContentTypeHtml interface {
	Body() string
}

func NewContentTypeHtml() *ContentTypeHtml {
	return &ContentTypeHtml{}
}

func (h *ContentTypeHtml) Body() string {
	return "text/html"
}
