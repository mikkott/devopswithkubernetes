package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/labstack/echo"
)

func generateRandomNumber() int {
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)
	return r.Intn(1000) + 1
}

func saveImageFromURL(url string, filename string) error {
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	response, err := client.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func isFileOlderThanDuration(filename string, duration time.Duration) (bool, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false, err
	}
	currentTime := time.Now()
	thresholdTime := currentTime.Add(-time.Second * duration)
	if fileInfo.ModTime().Before(thresholdTime) {
		return true, nil
	}
	return false, nil
}

func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func listFilesInDirectory(directory string) ([]string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames, nil
}

type Config struct {
	filepath      string
	cacheDuration time.Duration
}

var config = Config{
	filepath:      "/app/files/image.jpg",
	cacheDuration: time.Duration(time.Second * 60 * 60),
}

func getFileStatAsString(filepath string) (string, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return "", err
	}
	stat := fileInfo.Sys().(*syscall.Stat_t)
	statString := fmt.Sprintf("File Name: %s\n", fileInfo.Name())
	statString += fmt.Sprintf("Size: %d bytes\n", fileInfo.Size())
	statString += fmt.Sprintf("Permissions: %s\n", fileInfo.Mode().String())
	statString += fmt.Sprintf("Owner UID: %d\n", stat.Uid)
	statString += fmt.Sprintf("Owner GID: %d\n", stat.Gid)
	statString += fmt.Sprintf("Last Access Time: %s\n", fileInfo.ModTime().Format(time.RFC3339))
	return statString, nil
}

func main() {
	// Save image from URL if it doesn't exist
	if !fileExists(config.filepath) {
		err := saveImageFromURL(fmt.Sprintf("https://source.unsplash.com/%v", generateRandomNumber()), config.filepath)
		if err != nil {
			panic(err)
		}
	}

	// Debugging
	files, _ := listFilesInDirectory("/app")
	log.Printf("Files in /app: %v\n", files)
	files, _ = listFilesInDirectory("/app/files")
	log.Printf("Files in /app/files: %v\n", files)
	files, _ = listFilesInDirectory("/app/static")
	log.Printf("Files in /app/static: %v\n", files)

	// Echo instance
	e := echo.New()
	e.Static("/static", "/app/static")
	e.GET("/", indexHandler)
	e.GET("/files/image.jpg", imageHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if PORT environment variable is not set
	}
	e.Start(":" + port)
}

func indexHandler(c echo.Context) error {
	log.Println("Index handler called")
	// Open the HTML file
	htmlFile, err := os.Open("/app/index.html")
	if err != nil {
		log.Println(err)
		return err
	}
	defer htmlFile.Close()

	// Set the response header for HTML
	c.Response().Header().Set("Content-Type", "text/html")

	// Copy the HTML file to the response writer
	_, err = io.Copy(c.Response().Writer, htmlFile)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func imageHandler(c echo.Context) error {
	fileOlderThanDuration, err := isFileOlderThanDuration(config.filepath, config.cacheDuration)
	if err != nil {
		log.Println(err)
		return err
	}

	if fileOlderThanDuration {
		err := saveImageFromURL(fmt.Sprintf("https://source.unsplash.com/%v", generateRandomNumber()), config.filepath)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	fileStat, err := getFileStatAsString(config.filepath)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(fileStat)

	// Open the image file
	imageFile, err := os.Open("/app/files/image.jpg")
	if err != nil {
		log.Println(err)
		return err
	}
	defer imageFile.Close()

	// Set the response header for image
	c.Response().Header().Set("Content-Type", "image/jpg")

	// Copy the image file to the response writer
	_, err = io.Copy(c.Response().Writer, imageFile)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
