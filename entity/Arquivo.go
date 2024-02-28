package entity

import "path/filepath"

type Arquivo struct {
	Nome   string
	Path   string
	Tipo   string
	Size   uint
	Source string
	Bucket string
	Url    string
}

func NewArquivo() *Arquivo {
	return &Arquivo{}
}

func NewArquivoFromFilename(filename string) *Arquivo {
	nomeDoArquivo := filepath.Base(filename)
	caminho := filepath.Dir(filename)

	return &Arquivo{
		Nome: nomeDoArquivo,
		Path: caminho,
	}
}
