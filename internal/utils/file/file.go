package file

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/nfnt/resize"
	tele "gopkg.in/telebot.v3"
)

type Photo struct {
	File *os.File
	// Parameters for resizing the image to the specified dimensions.
	Width  int
	Height int
}
// DownloadPhoto downloads a photo from Telegram and saves it to the local file system.
// If the message contains a document, it will be downloaded as a document.
// Otherwise, it will be downloaded as a photo.
func (p *Photo) DownloadPhoto(msg *tele.Message, b *tele.Bot) (photo *tele.Photo , err error){
	// Path to the file.
	path := preparePath(msg.Chat.ID)
	if msg.Document != nil{
		// Download document.
		err = b.Download(&msg.Document.File, path)
		if err != nil {
			return nil, err
		}
	}else {
		// Download file.png.
		err = b.Download(&msg.Photo.File, path)
		if err != nil {
			return nil, err
		}
	}
	fmt.Printf("Download")
	// Create Photo. Reading the photo from file.
	photo = &tele.Photo{File: tele.FromDisk(path)}
	if photo.Height != p.Height || photo.Width != p.Width {
		if err = p.resizePhoto(path); err != nil {
			return nil, err
		}
		photo = &tele.Photo{File: tele.FromDisk(path)}
    }

    return photo, nil
}

func (p *Photo) resizePhoto(path string) error {
	// open file.
	file, err := os.Open(path)
	if err!= nil {
		return err
	}
	// read bytes from file.
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	// detection image and resize.
	if DetectImage(fileBytes) == "image/png" {
		// resize image PNG.
		if err := p.resizePNG(file, &path); err != nil {
			return err
        }
		return nil
	}
    // resize image JPEG.
	if err := p.resizeJPEG(file, &path); err != nil {
		return err
	}
	return nil
}
// resizePNG resizes the image to the given size.
func (p *Photo) resizePNG(file *os.File, path *string) error{
	// decode png into image.Image
	img, err := png.Decode(file)
	if err != nil {
		return err
	}
	file.Close()
	// resize to 512x512 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(512, 512, img, resize.Lanczos3)
	out, err := os.Create(*path)
	if err != nil {
		return err
	}
	defer out.Close()
	// write new image to file
	png.Encode(out, m)
	return nil
}
// resizeJPEG resizes the image to the given size.
func (p *Photo) resizeJPEG(file *os.File, path *string) error{
	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}
	file.Close()
	// resize to 512x512 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(512, 512, img, resize.Lanczos3)
	out, err := os.Create(*path)
	if err != nil {
		return err
	}
	defer out.Close()
	// write new image to file
	jpeg.Encode(out, m, nil)
	return nil
}
// DetectImage returns the mime type of an image file from its first few
// bytes or the empty string if the file does not look like a known file type
func DetectImage(incipit []byte) string{
	// image formats and magic numbers
		var magicTable = map[string]string{
			"\xff\xd8\xff":      "image/jpeg",
			"\x89PNG\r\n\x1a\n": "image/png",
		}
	
		incipitStr := string(incipit) 
		for magic, mime := range magicTable {
			if strings.HasPrefix(incipitStr, magic) {
				fmt.Printf("%s\n", mime)
				return mime
			}
		}
		return ""
	}
	

func preparePath(chatID int64) string{
	return fmt.Sprintf("photos/%d", chatID)
}