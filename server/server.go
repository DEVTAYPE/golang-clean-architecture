package server

import (
	"fmt"
	"net/http"
	"strings"
)

type App struct {
	mux            *http.ServeMux //definimos un multiplexer(un multiplexor es un componente que dirige las solicitudes HTTP entrantes a los controladores adecuados) de rutas
	handlerCounter int            //contador de handlers registrados
}

func NewApp() *App {
	return &App{
		mux:            http.NewServeMux(), //inicializamos el multiplexer
		handlerCounter: 0,
	}
}

func textToCentered(text string, width int) string {
	if len(text) >= width {
		return text[:width]
	}

	padding := (width - len(text)) / 2

	return strings.Repeat(" ", padding) + text + strings.Repeat(" ", padding)
}

func (app *App) PrintBanner(port string) {
	urlBase := fmt.Sprintf("http://localhost:%s", port)
	handlerCount := fmt.Sprintf("Tienes registrados %d handlers", app.handlerCounter)

	fmt.Println("===============================")
	fmt.Printf("%s", textToCentered(urlBase, 30)+"\n")
	fmt.Printf("%s", textToCentered(handlerCount, 30)+"\n")
	fmt.Println("===============================")

}

func (app *App) RunServer(port string) error {
	server := &http.Server{
		Addr:    ":" + port,
		Handler: app.mux, //asignamos el multiplexer como el manejador de solicitudes del servidor
	}

	app.PrintBanner(port)

	return server.ListenAndServe()
}
