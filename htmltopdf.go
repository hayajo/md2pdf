package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings")

const wkhtmltopdfCmd = "wkhtmltopdf"

func HtmlToPdf(html []byte, outfile string, opts ...string) error {
	wkhtmltopdf := os.Getenv("MD2PDF_WKHTMLTOPDF")
	if wkhtmltopdf == "" {
		wkhtmltopdf = wkhtmltopdfCmd
	}

	tmp, err := ioutil.TempFile("", "")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		os.Remove(tmp.Name())
	}()
	tmp.Write(html)
	tmp.Close()

	pagefile := tmp.Name() + ".html"
	err = os.Rename(tmp.Name(), pagefile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		os.Remove(pagefile)
	}()

	opts = append(opts, pagefile, outfile)
	cmd := exec.Command(wkhtmltopdf, opts...)

	log.Println(strings.Join(cmd.Args, " "))

	output, err := cmd.CombinedOutput()
	if err != nil {
		logger := log.New(os.Stderr, log.Prefix(), log.Flags())
		logger.Println(cmd.Args)
		logger.Println(err)
		return errors.New(string(output))
	}

	return nil
}
