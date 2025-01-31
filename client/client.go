package client

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (

	//go:embed dist/*

	dist embed.FS

	//go:embed dist/index.html

	indexHTML embed.FS

	// Vite commonly uses ports 5173-5183

	vitePorts = []int{5173, 5174, 5175, 5176, 5177, 5178, 5179, 5180, 5181, 5182, 5183}
)

func findActiveViteServer() (*url.URL, error) {
	client := http.Client{
		Timeout: 100 * time.Millisecond,
	}

	// Wait for the Vite dev server to start
	time.Sleep(1 * time.Second)

	for _, port := range vitePorts {

		targetURL := fmt.Sprintf("http://localhost:%d", port)

		_, err := client.Get(targetURL)

		if err == nil {

			log.Printf("Found Vite dev server at port %d", port)

			return url.Parse(targetURL)

		}

	}

	return nil, fmt.Errorf("no active Vite dev server found on ports %v", vitePorts)
}

func RegisterHandlers(r *gin.Engine, env string) {
	if env == "dev" {
		log.Println("Running in dev mode")
		setupDevProxy(r)
		return

	}

	log.Println("Running in prod mode")
	// Strip the "dist" prefix from the embedded filesystem

	distFS, err := fs.Sub(dist, "dist")
	if err != nil {
		log.Fatalf("Failed to create sub-filesystem: %v", err)
	}

	// Serve index.html for the root path and handle SPA routes

	serveIndexHTML := func(c *gin.Context) {
		indexFile, err := distFS.Open("index.html")
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		defer indexFile.Close()

		http.ServeContent(c.Writer, c.Request, "index.html", time.Now(), indexFile.(io.ReadSeeker))
	}

	// Handle all routes

	r.NoRoute(func(c *gin.Context) {
		// Try to serve static files

		filePath := strings.TrimPrefix(c.Request.URL.Path, "/")

		if filePath == "" {
			filePath = "index.html"
		}

		file, err := distFS.Open(filePath)
		if err != nil {
			// If the file doesn't exist, it might be a frontend route
			if strings.HasPrefix(c.Request.URL.Path, "/api") || strings.HasPrefix(c.Request.URL.Path, "/r") {
				// For API or redirect routes, let Gin handle it
				c.Next()
			} else {
				// For frontend routes, serve index.html
				serveIndexHTML(c)
			}
			return
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		if stat.IsDir() {
			// If it's a directory, try to serve index.html from that directory
			indexPath := path.Join(filePath, "index.html")

			indexFile, err := distFS.Open(indexPath)
			if err != nil {
				serveIndexHTML(c)
				return
			}

			defer indexFile.Close()

			http.ServeContent(c.Writer, c.Request, "index.html", stat.ModTime(), indexFile.(io.ReadSeeker))

		} else {
			http.ServeContent(c.Writer, c.Request, stat.Name(), stat.ModTime(), file.(io.ReadSeeker))
		}
	})
}

func setupDevProxy(r *gin.Engine) {
	targetURL, err := findActiveViteServer()
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	r.Use(func(c *gin.Context) {
		// Skip the proxy if the prefix is /api or /r

		if strings.HasPrefix(c.Request.URL.Path, "/api") || strings.HasPrefix(c.Request.URL.Path, "/r") {

			c.Next()

			return

		}

		proxy.ServeHTTP(c.Writer, c.Request)
	})
}
