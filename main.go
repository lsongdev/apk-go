package main

import (
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/lsongdev/apk-go/apk"
)

func main() {
	apk, err := apk.Open("testdata/helloworld.apk")
	if err != nil {
		log.Fatal(err)
	}
	defer apk.Close()

	icon, err := apk.Icon(nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(icon)

	label, err := apk.Label(nil)
	if err != nil {
		log.Fatal(err)
	}
	pkgName := apk.PackageName()
	log.Println(label, pkgName)

	mainActivity, err := apk.MainActivity()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(mainActivity)

	manifest := apk.Manifest()
	targetSdk := manifest.SDK.Target.MustInt32()
	log.Println(targetSdk)
}
