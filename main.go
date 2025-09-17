package main

import (
	"context"
	"dockhome/types"
	"dockhome/web"
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("DEV_MODE") == "true" {
			os.Setenv("DOMAIN", "localhost")

			frontendUrl, _ := url.Parse("http://localhost:5173")
			httputil.NewSingleHostReverseProxy(frontendUrl).ServeHTTP(w, r)
		} else {
			path := r.URL.Path[1:]
			if path == "" {
				path = "index"
			}

			ext := filepath.Ext(r.URL.Path)
			if ext == "" {
				ext = ".html"
				path += ext
			}

			var fileContent []byte
			fileContent, err := web.BuildDir.ReadFile(filepath.Join("build", path))
			if err != nil {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			}

			w.Header().Set("Content-Type", mime.TypeByExtension(ext))
			w.Write(fileContent)
		}
	})

	r.HandleFunc("/api/domain", func(w http.ResponseWriter, r *http.Request) {
		response := types.NeededData{
			Domain: os.Getenv("DOMAIN"),
		}

		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	r.HandleFunc("/api/containers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

		for {
			containers, _ := cli.ContainerList(context.Background(), container.ListOptions{
				All: true,
			})

			var response []types.ContainerData

			for _, container := range containers {
				response = append(response, types.ContainerData{
					Id:    container.ID,
					Name:  strings.Split(container.Names[0], "/")[1],
					Icon:  container.Labels["Home.icon"],
					Alias: container.Labels["Dockport.alias"],
				})
			}

			sort.SliceStable(response, func(i, j int) bool {
				return response[i].Name < response[j].Name
			})

			jsonResponse, _ := json.Marshal(response)

			fmt.Fprintf(w, "event: message\n")
			fmt.Fprintf(w, "data: %s\n\n", jsonResponse)

			w.(http.Flusher).Flush()
			time.Sleep(1 * time.Second)
		}
	})

	fmt.Println("DockHome is ready on port :80")
	http.ListenAndServe(":80", r)
}
