package email

type ContentType struct {
	ContentTypeBody string
}

type IContentType interface {
	Body() string
}

func NewContentType() *ContentType {
	return &ContentType{}
}

func (e *ContentType) Body() string {
	return e.ContentTypeBody
}

func ContentTypeBody(body IContentType) string {
	return body.Body()
}
