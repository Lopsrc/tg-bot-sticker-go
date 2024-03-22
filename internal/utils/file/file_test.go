package file

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"testing"

	"github.com/nfnt/resize"
	"github.com/stretchr/testify/assert"
)

func Test_DownloadPhoto(t *testing.T){

	filePh := "test.jpeg"
	filePhNew := "complete.jpeg"
	path := preparePathTst(filePh)

	file, err := os.Open(path)
    assert.NoError(t, err)

	img, err := jpeg.Decode(file)
	assert.NoError(t, err)
	file.Close()

	m := resize.Resize(512, 512, img, resize.Lanczos3)

	path = preparePathTst(filePhNew)
	out, err := os.Create(path)
	assert.NoError(t, err)
	defer out.Close()

	jpeg.Encode(out, m, nil)

	filePh = "test.png"
	filePhNew = "complete.png"
	path = preparePathTst(filePh)
    file, err = os.Open(path)
    assert.NoError(t, err)

	img, err = png.Decode(file)
	assert.NoError(t, err)
	file.Close()

	m = resize.Resize(512, 512, img, resize.Lanczos3)

	path = preparePathTst(filePhNew)
	out, err = os.Create(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	defer out.Close()

	png.Encode(out, m)
}

func preparePathTst(photo string) string{
	return fmt.Sprintf("../../../photos/%s", photo)
}