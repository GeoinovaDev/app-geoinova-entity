package entity

type DeteccaoPreviewsBuilder struct {
	preview *DeteccaoPreviews
}

func NewDeteccaoPreviewsBuilder() *DeteccaoPreviewsBuilder {
	return &DeteccaoPreviewsBuilder{
		preview: NewDeteccaoPreviews(),
	}
}

func (b *DeteccaoPreviewsBuilder) WithAntes(antes *DeteccaoPreview) *DeteccaoPreviewsBuilder {
	b.preview.Antes = antes
	return b
}

func (b *DeteccaoPreviewsBuilder) WithDepois(depois *DeteccaoPreview) *DeteccaoPreviewsBuilder {
	b.preview.Depois = depois
	return b
}

func (b *DeteccaoPreviewsBuilder) WithRegiao(regiao *DeteccaoPreview) *DeteccaoPreviewsBuilder {
	b.preview.Regiao = regiao
	return b
}

func (b *DeteccaoPreviewsBuilder) Build() *DeteccaoPreviews {
	return b.preview
}
