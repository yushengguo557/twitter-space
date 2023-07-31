package upload

import (
	"fmt"
	"testing"
)

func TestSaveImageUrl(t *testing.T) {
	url, err := SaveImageUrl("twitter-user", "https://pbs.twimg.com/profile_images/1645027119481724929/ES3aeKZr_normal.jpg")
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
}
