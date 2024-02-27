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

func NewArquivo(nome string, path string) *Arquivo {
	return &Arquivo{
		Nome: nome,
		Path: path,
	}
}

func NewArquivoFromFilename(filename string) *Arquivo {
	nomeDoArquivo := filepath.Base(filename)
	caminho := filepath.Dir(filename)

	return NewArquivo(nomeDoArquivo, caminho)
}
