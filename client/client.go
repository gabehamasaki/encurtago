package client

import (
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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

	// Create file systems for static files and index.html
	distFS := http.FS(dist)
	indexFS := http.FS(indexHTML)

	// Serve static files

	r.StaticFS("/", distFS)

	// Serve index.html for the root path

	r.GET("/", func(c *gin.Context) {
		file, err := indexFS.Open("dist/index.html")
		if err != nil {

			c.Status(http.StatusInternalServerError)

			return

		}

		defer file.Close()

		http.ServeContent(c.Writer, c.Request, "index.html", time.Now(), file.(io.ReadSeeker))
	})

	// Handle SPA routes - serve index.html for non-API routes and redirects routes

	r.NoRoute(func(c *gin.Context) {
		// Skip API routes

		if strings.HasPrefix(c.Request.URL.Path, "/api") || strings.HasPrefix(c.Request.URL.Path, "/r") {
			c.Next()

			return

		}

		file, err := indexFS.Open("dist/index.html")
		if err != nil {
			c.Status(http.StatusInternalServerError)

			return

		}

		defer file.Close()

		http.ServeContent(c.Writer, c.Request, "index.html", time.Now(), file.(io.ReadSeeker))
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
