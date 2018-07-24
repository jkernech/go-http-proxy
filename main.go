// Package proxy implements a proxy that forward HTTP requests.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	version = "master"
)

// GetPathMapping returns a key/value map of the PATH_MAPPING env var.
func GetPathMapping() map[string]string {
	pathMappingValue := os.Getenv("PATH_MAPPING")

	var pathMapping map[string]string
	json.Unmarshal([]byte(pathMappingValue), &pathMapping)
	return pathMapping
}

// GetURLPathPrefix returns the first directory if it exists for a given URL.
func GetURLPathPrefix(requestedURL string) string {
	urlParts := strings.Split(requestedURL, "/")
	if len(urlParts) > 0 {
		return urlParts[0]
	}
	return ""
}

// NormalizeURL adds a scheme to a given URL without scheme.
func NormalizeURL(rawURL string) string {
	if !strings.Contains(rawURL, "http") {
		return "http://" + rawURL
	}
	return rawURL
}

// IsValidURL validates URL format and ensure host contains a dot.
func IsValidURL(rawURL string) bool {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	if !strings.Contains(parsedURL.Host, ".") {
		return false
	}
	return true
}

// Index handler process all the incoming HTTP requests.
func Index(c *gin.Context) {
	// Forward to Status page when homepage is requested
	if c.Request.URL.Path == "/" {
		Status(c)
		return
	}

	// Get the requested URL path and trim the / prefix
	requestedURL := strings.TrimPrefix(c.Request.URL.String(), "/")

	urlPathPrefix := GetURLPathPrefix(requestedURL)
	pathMapping := GetPathMapping()

	if host, ok := pathMapping[urlPathPrefix]; ok {
		requestedURL = strings.Replace(requestedURL, urlPathPrefix+"/", "", 1)

		// Forward to 404 page when URL path prefix matches the requested URL
		if urlPathPrefix == requestedURL {
			Error(c, http.StatusNotFound)
			return
		}

		requestedURL = host + "/" + requestedURL
	}

	// Add default HTTP scheme if not provided
	requestedURL = NormalizeURL(requestedURL)

	// Ensure we forward valid URLs
	if !IsValidURL(requestedURL) {
		Error(c, http.StatusNotFound)
		return
	}

	resp, err := http.Get(requestedURL)

	if err != nil {
		Error(c, http.StatusNotFound)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		Error(c, http.StatusNotFound, err)
		return
	}

	c.Header("Content-Type", resp.Header.Get("Content-Type"))
	c.String(resp.StatusCode, string(body))
}

// Status returns the current time in a JSON object
func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"time": time.Now().String(),
	})
}

// Error wrapper centralizes the error responses logic.
func Error(c *gin.Context, statusCode int, err ...error) {
	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": http.StatusText(statusCode),
		"url":     c.Request.URL.String(),
	})

	if err != nil {
		log.Print(err)
	}
}

func main() {
	fmt.Printf("%v", version)

	godotenv.Load()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/*path", Index)
	router.Run()
}
