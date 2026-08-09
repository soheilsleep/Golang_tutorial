package main

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Print("Hello World")
	log.Info().Str("Category", "Search").Msg("Searching for a thing")
	err := errors.New("this is an error")
	log.Error().Err(err).Msg("something went wrong")
	err = func3()
	if err != nil {
		log.Error().Stack().Err(err).Msg("Error occurred (func 3)")
	}

}
func func1() error {
	return errors.New("something went wrong (func 1)")
}
func func2() error {
	err := func1()
	if err != nil {
		return err
	}
	return nil
}
func func3() error {
	err := func2()
	if err != nil {
		return err
	}
	return nil
}

//trace
//debug
//info
//warn
//error
//fatal(critical)
//panic
