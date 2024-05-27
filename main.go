package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func downloadImage(url string, filename string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer o pedido GET:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Falha ao baixar a imagem. Código de status:", response.StatusCode)
		return
	}

	imgData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erro ao ler os dados da imagem:", err)
		return
	}

	err = os.MkdirAll("img", os.ModePerm)
	if err != nil {
		fmt.Println("Erro ao criar o diretório 'img':", err)
		return
	}

	filepath := filepath.Join("img", filename)
	err = ioutil.WriteFile(filepath, imgData, 0644)
	if err != nil {
		fmt.Println("Erro ao escrever os dados da imagem:", err)
		return
	}

	fmt.Println("Download concluído! Imagem salva em:", filepath)
}

func main() {
	urlPtr := flag.String("url", "", "Link da imagem a ser baixada")
	linksPtr := flag.String("links", "", "Links das imagens a serem baixadas, separados por espaço")
	flag.Parse()

	if *urlPtr != "" {
		ext := filepath.Ext(*urlPtr)
		filename := "ref_pixel_art" + ext
		downloadImage(*urlPtr, filename)
	}

	if *linksPtr != "" {
		links := flag.Args()
		counter := 1
		for _, url := range links {
			ext := filepath.Ext(url)
			filename := fmt.Sprintf("ref_pixel_art_%d%s", counter, ext)
			downloadImage(url, filename)
			counter++
		}
	}
}
