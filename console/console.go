package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/tarantool/go-tarantool"
)

func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack //nolint: reassign
	log.Logger = log.With().Stack().Logger()

	if err := doMain(); err != nil {
		log.Fatal().Err(err).Send()
	}
}

func doMain() error {
	// actually tarantella doesn't check password, and username used for directory, where data files are stored
	opts := tarantool.Opts{User: "user", Pass: "DSoXbver3p4bbMK6dGhUfo"}
	conn, err := tarantool.Connect("127.0.0.1:3302", opts)
	if err != nil {
		return err
	}
	log.Info().Any("greeting", conn.Greeting).Msg("Greeting received")
	r, err := conn.Ping()
	log.Error().Err(err).Uint32("res", r.Code).Msg("Ping successfully")

	resp, err := conn.Insert("tester", []any{"assd", "ABBA", 1972})
	// resp, err := conn.Insert("tester", []any{5, "Lord Huron", 2020})
	if err != nil {
		log.Error().Err(err).Msg("Insert failed")
	}
	log.Info().Any("response", resp).Msg("Insert")

	resp, err = conn.Select("tester", "secondary", 0, 1, tarantool.IterEq, []any{"ABBA"})
	if err != nil {
		log.Error().Err(err).Msg("Select failed")
	}
	log.Info().Any("response", resp).Msg("Select")

	return nil
}