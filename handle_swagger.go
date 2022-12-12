package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed docs/swagger-ui/*
var swaggerStatic embed.FS

func (s *Server) handleSwagger() http.HandlerFunc {

	swaggerSubDirectoryFs, err := fs.Sub(swaggerStatic, "docs/swagger-ui")
	if err != nil {
		s.log.Panic("failed to load swagger directory", err)
	}

	// TODO figure out why route has to be `/swagger/` and not `/swagger`
	return http.StripPrefix("/swagger/", http.FileServer(http.FS(swaggerSubDirectoryFs))).ServeHTTP
}
