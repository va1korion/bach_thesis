package main

import (
	"bach_thesis/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"google.golang.org/grpc"
)

type resp map[string]interface{}

var client api.ManagerClient

func main() {

	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}

	client = api.NewManagerClient(conn)

	http.Handle("/", handlers(client))
	http.ListenAndServe(":8000", nil)
}

func handlers(client proto.StreamServiceClient) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexPage).Methods("GET")
	router.HandleFunc("/media/{mID:[0-9]+}/stream/", streamHandler).Methods("GET")
	router.HandleFunc("/media/{mID:[0-9]+}/stream/{segName:index[0-9]+.ts}", streamHandler).Methods("GET")
	return router
}

func streamHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	// TODO: To be done later for multiple media files
	// mID, err := strconv.Atoi(vars["mID"])
	// if err != nil {
	// 	fmt.Println("ERROR FOUND")
	// 	response.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	// fmt.Println("mID: ", mID)

	segName, ok := vars["segName"]
	if !ok {
		serveHlsM3u8(response, request, "../media", "index.m3u8")
	} else {
		serveHlsTs(response, request, "../media", segName)
	}
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
