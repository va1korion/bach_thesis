package main

import (
	"bach_thesis/api"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/credentials"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"
)

type resp map[string]interface{}

var client api.ManagerClient

func main() {
	// todo TLS certificate from env
	creds, err := credentials.NewClientTLSFromFile("./openssl/root.pem", "")
	if err != nil {
		log.Fatal(err.Error())
	}
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err.Error())
	}

	client = api.NewManagerClient(conn)

	http.Handle("/", handlers(client))
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handlers(client api.ManagerClient) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexPage).Methods("GET")
	router.HandleFunc("/media/{mID:[0-9]+}/stream/", streamHandler).Methods("GET")
	router.HandleFunc("/media/{mID:[0-9]+}/stream/{segName:index[0-9]+.ts}", streamHandler).Methods("GET")
	return router
}

// localhost streaming between this and browser
func streamHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	segName, ok := vars["segName"]
	if !ok {
		serveHlsM3u8(response, request, "../media", "index.m3u8")
	} else {
		serveHlsTs(response, request, "../media", segName)
	}
}

func serveHlsM3u8(w http.ResponseWriter, r *http.Request, mediaBase, m3u8Name string) {
	mediaFile := fmt.Sprintf("%s/%s", mediaBase, m3u8Name)
	req := &api.Request{Filename: mediaFile}

	if response, err := client.GetFile(context.Background(), req); err == nil {
		fo, err := os.Create(mediaFile)
		if err != nil {
			log.Fatal("Failed to Create File")
		}
		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()
		if _, err := fo.Write(response.GetContent()); err != nil {
			log.Fatal("Error in writing file: ", err.Error())
		}
		http.ServeFile(w, r, mediaFile)
		w.Header().Set("Content-Type", "application/x-mpegURL")
	} else {
		fmt.Println("Error getting response: ", err.Error())
	}
}

func serveHlsTs(w http.ResponseWriter, r *http.Request, mediaBase, segName string) {
	mediaFile := fmt.Sprintf("%s/%s", mediaBase, segName)
	req := &api.Request{Filename: mediaFile}

	if response, err := client.GetFile(context.Background(), req); err == nil {
		fo, err := os.Create(mediaFile)
		if err != nil {
			log.Fatal("Failed to Create File")
		}
		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()
		if _, err := fo.Write(response.GetContent()); err != nil {
			log.Fatal("Error in writing file: ", err.Error())
		}
		http.ServeFile(w, r, mediaFile)
		w.Header().Set("Content-Type", "video/MP2T")
	} else {
		fmt.Println("Error getting response: ", err.Error())
	}
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
