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

	r.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

		containers, _ := cli.ContainerList(context.Background(), container.ListOptions{
			All: true,
		})

		response := types.NeededData{
			Domain: os.Getenv("DOMAIN"),
		}

		for _, container := range containers {
			containerData := types.ContainerData{
				Id:    container.ID,
				Name:  strings.Split(container.Names[0], "/")[1],
				Icon:  container.Labels["Home.icon"],
				Alias: container.Labels["Dockport.alias"],
			}

			response.Containers = append(response.Containers, containerData)

			sort.SliceStable(response.Containers, func(i, j int) bool {
				return response.Containers[i].Name < response.Containers[j].Name
			})
		}

		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("DockHome is ready on port :80")
	http.ListenAndServe(":80", r)
}
