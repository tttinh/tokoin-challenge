package main

import (
	"github.com/tttinh/tokoin-challenge/config"
	"github.com/tttinh/tokoin-challenge/repo"
	"github.com/tttinh/tokoin-challenge/utils"
)

func main() {
	cfg := config.New()
	if err := cfg.Init("./config", "default"); err != nil {
		utils.WriteLine(err)
		return
	}

	rp := repo.New()
	if err := rp.Init(cfg); err != nil {
		utils.WriteLine(err)
		return
	}

	app := NewApp(rp)
	app.Run()
}
