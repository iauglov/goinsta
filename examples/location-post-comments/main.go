package main

import (
	"fmt"
	"github.com/iauglov/goinsta/v2"
	"log"
	"os"
)

func fetchLocation(insta *goinsta.Instagram, locationID int64) error {
	feeds, err := insta.Locations.Feeds(locationID)
	if err != nil {
		return err
	}
	for i := range feeds.Sections {
		medias := feeds.Sections[i].LayoutContent.Medias
		for i2 := range medias {
			comments := medias[i2].Media.Comments
			comments.Sync()
			for comments.Next() {
				for _, item := range comments.Items {
					fmt.Println(item.Text)
				}
			}
		}
	}
	return nil
}

func main() {
	insta := goinsta.New(
		os.Getenv("INSTAGRAM_USERNAME"),
		os.Getenv("INSTAGRAM_PASSWORD"),
	)
	if err := insta.Login(); err != nil {
		log.Println(err)
		return
	}
	defer insta.Logout()

	for _, tag := range []int64{
		17142327,
	} {
		if err := fetchLocation(insta, tag); err != nil {
			log.Println(tag, err)
		}
	}
}
