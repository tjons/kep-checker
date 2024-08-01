package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/yuin/goldmark"
)

var (
	KepUrl, DefaultKepTemplateURL string
)

func init() {
	flag.StringVar(&KepUrl, "kep-url", "", "required: URL to KEP to validate")
	flag.StringVar(&DefaultKepTemplateURL, "kep-template-url", "https://raw.githubusercontent.com/kubernetes/enhancements/master/keps/NNNN-kep-template/README.md", "URL to KEP template")
	flag.Parse()

	if KepUrl == "" {
		fmt.Println("kep-url is required")
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Getting latest KEP template...")

	resp, err := http.Get(DefaultKepTemplateURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Error getting KEP template:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	template, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading KEP template:", err)
		os.Exit(1)
	}

	fmt.Println(string(template))

	fmt.Println("Getting KEP to validate...")
	resp, err = http.Get(KepUrl)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Error getting KEP:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	kep, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading KEP:", err)
		os.Exit(1)
	}

	goldmark.New().Parse(kep)
}
