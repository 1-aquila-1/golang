package main

import (
	"bytes"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/1-aquila-1/sdk-go/v1_0/common"
	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
)

func RedimencionarTodas() {
	path := "C:\\Workspace\\Pessoal\\img1.jpg"
	novoArquivo := "C:\\Workspace\\Pessoal\\Projetos\\controle-acesso\\banner5-" + common.Uuid10() + ".png"

	width := uint(640)
	height := uint(640)

	filebytes, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	mimeType := http.DetectContentType(filebytes)

	switch mimeType {
	case "image/jpeg", "image/png":
	default:
		log.Fatal("Formato não permitido")
	}

	tam_max := 200000
	tam_file := len(filebytes)
	if tam_file > tam_max {
		log.Fatal("Somente imagem de 200bk é permitido")
	}

	img, err := imaging.Decode(bytes.NewReader(filebytes))

	if err != nil {
		log.Fatal(err)
	}
	m := resize.Resize(width, height, img, resize.Lanczos3)

	out, err := os.Create(novoArquivo)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)

}
