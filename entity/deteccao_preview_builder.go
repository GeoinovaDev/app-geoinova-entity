package entity

type DeteccaoPreviewBuilder struct {
	preview *DeteccaoPreview
}

func NewDeteccaoPreviewBuilder() *DeteccaoPreviewBuilder {
	return &DeteccaoPreviewBuilder{
		preview: NewDeteccaoPreview(),
	}
}

func (b *DeteccaoPreviewBuilder) WithNome(nome string) *DeteccaoPreviewBuilder {
	b.preview.Nome = nome
	return b
}

func (b *DeteccaoPreviewBuilder) WithTipo(tipo string) *DeteccaoPreviewBuilder {
	b.preview.Tipo = tipo
	return b
}

func (b *DeteccaoPreviewBuilder) WithBase64(base64 string) *DeteccaoPreviewBuilder {
	b.preview.Base64 = base64
	return b
}

func (b *DeteccaoPreviewBuilder) WithPath(path string) *DeteccaoPreviewBuilder {
	b.preview.Path = path
	return b
}

func (b *DeteccaoPreviewBuilder) WithBucket(bucket string) *DeteccaoPreviewBuilder {
	b.preview.Bucket = bucket
	return b
}

func (b *DeteccaoPreviewBuilder) Build() *DeteccaoPreview {
	return b.preview
}
