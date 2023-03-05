package router

import "github.com/gorilla/mux"

func Router() *mux.router {

	myRouter := mux.NewRouter()

	myRouter.HandleFunc('/api/stock/{id}', middleware.GetStock)
	myRouter.HandleFunc('/api/stock', middleware.GetAllStock)
	myRouter.HandleFunc('/api/newstock', middleware.CreateStock)
	myRouter.HandleFunc('/stock/{id}', middleware.UpdateStock)
	myRouter.HandleFunc('/api/delete/{id}', middleware.DeleteStock)

}
