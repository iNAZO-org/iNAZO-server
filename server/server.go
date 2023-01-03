package server

import "karintou8710/iNAZO-server/config"

// Init initialize server
func Init() error {
	c := config.GetConfig()
	r, err := NewRouter()
	if err != nil {
		return err
	}
	r.Logger.Fatal(r.Start(":" + c.GetString("app.port")))
	return nil
}
