package main

import (
	"fmt"
	"github.com/iauglov/goinsta/v2"
	"log"
	"os"
)

func fetchTag(insta *goinsta.Instagram, tag string) error {
	feedTag, err := insta.Feed.Tags(tag)
	if err != nil {
		return err
	}
	for i := range feedTag.Images {
		comments := feedTag.Images[i].Comments
		comments.Sync()
		for comments.Next() {
			for _, item := range comments.Items {
				fmt.Println(item.Text)
			}
		}
	}
	for i := range feedTag.RankedItems {
		comments := feedTag.RankedItems[i].Comments
		comments.Sync()
		for comments.Next() {
			for _, item := range comments.Items {
				fmt.Println(item.Text)
			}
		}
	}
	fmt.Println(feedTag)
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

	for _, tag := range []string{
		"москва",
	} {
		if err := fetchTag(insta, tag); err != nil {
			log.Println(tag, err)
		}
	}
}
