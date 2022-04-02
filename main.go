package main

import (
	`github.com/pangum/pangu`
)

func main() {
	panic(pangu.New(
		pangu.Banner(`Ziyunix Agent`, pangu.BannerTypeAscii),
	).Run(newBootstrap))
}
