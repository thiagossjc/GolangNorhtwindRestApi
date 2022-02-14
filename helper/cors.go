package helper

import "github.com/go-chi/cors"

func GetCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                                                       //todos con acesso
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                 //cuales metodos son permitidos
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, //cuales headers son permitidos
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	})
}
