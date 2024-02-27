package entity

import "path/filepath"

type Arquivo struct {
	Nome string
	Path string
}

func NewArquivo(nome string, path string) *Arquivo {
	return &Arquivo{nome, path}
}

func NewArquivoFromFilename(filename string) *Arquivo {
	nomeDoArquivo := filepath.Base(filename)
	caminho := filepath.Dir(filename)

	return NewArquivo(nomeDoArquivo, caminho)
}
