package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"./converter"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [markdown] output\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	var (
		css   = flag.String("css", "", "css file")
		title = flag.String("title", "md2pdf", "document title")
	)
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	var r io.ReadCloser
	var out string
	if len(flag.Args()) > 1 {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		r = f
		out = flag.Arg(1)
	} else {
		r = os.Stdin
		out = flag.Arg(0)
	}
	defer func() { r.Close() }()

	text, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	md := converter.NewMarkdown(text)

	opts := []string{"--print-media-type"}
	if *css != "" {
		var cssUrl string
		u, _ := url.Parse(*css)
		if u.Scheme != "http" || u.Scheme != "https" {
			abs, _ := filepath.Abs(*css)
			cssUrl = "file://" + abs
		} else {
			cssUrl = u.String()
		}
		opts = append(opts, "--user-style-sheet", filepath.ToSlash(cssUrl))
	}

	err = md.ToPdf(out, *title, opts...)
	if err != nil {
		log.Fatal(err)
	}
}
