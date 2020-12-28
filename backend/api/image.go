package api

import (
	"image"
	"io/ioutil"
	"os"

	"github.com/lafin/http"
	"gitlab.com/opennota/screengen"
)

// GetImage - get image
func GetImage(url string) (image.Image, error) {
	resp, err := http.Get(url, nil)
	if err != nil {
		return nil, err
	}
	tmpfile, err := ioutil.TempFile("", "stream")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpfile.Name())
	if _, err = tmpfile.Write(resp); err != nil {
		return nil, err
	}
	err = tmpfile.Close()
	if err != nil {
		return nil, err
	}
	gen, err := screengen.NewGenerator(tmpfile.Name())
	if err != nil {
		return nil, err
	}
	defer gen.Close()
	return gen.Image(0)
}
