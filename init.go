package main

import "github.com/gsora/garbagetextbot/meme"

func init() {
	apiKey = meme.GetVariableOrFatal("TG_APIKEY")
}
