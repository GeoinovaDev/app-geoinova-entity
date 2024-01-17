package entity

type AtivoBucketBuilder struct {
	ativoBucket *AtivoBucket
}

func NewAtivoBucketBuilder(id uint) *AtivoBucketBuilder {
	return &AtivoBucketBuilder{
		ativoBucket: NewAtivoBucket(id),
	}
}

func (b *AtivoBucketBuilder) WithAtivo(ativo *Ativo) *AtivoBucketBuilder {
	b.ativoBucket.Ativo = ativo
	return b
}

func (b *AtivoBucketBuilder) WithBucket(bucket string) *AtivoBucketBuilder {
	b.ativoBucket.Bucket = bucket
	return b
}

func (b *AtivoBucketBuilder) WithIsDefault(isDefault bool) *AtivoBucketBuilder {
	b.ativoBucket.IsDefault = isDefault
	return b
}

func (b *AtivoBucketBuilder) Build() *AtivoBucket {
	return b.ativoBucket
}
