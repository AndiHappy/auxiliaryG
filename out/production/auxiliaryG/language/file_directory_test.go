package gotutorial

import (
	"archive/zip"
	"bufio"
	"fmt"
	"github.com/AndiHappy/auxiliaryG/util"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

var path = "./TT/test"

func TestMakeDirectory(t *testing.T) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(path, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}

func TestRenameFile(t *testing.T) {
	oldName := path + "/" + "test.txt"
	util.CreateFile(oldName)

	newName := path + "/" + "testing.txt"
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}

func TestCopyFile(t *testing.T) {
	sourceFile, err := os.Open("/var/www/html/src/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create("/var/www/html/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	//copy 的
	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesCopied)
}

func TestMoveFile(t *testing.T) {
	oldLocation := "./TT/test/testing.txt"
	// newLocation 的路径下的文件夹必须存在
	newLocation := "./TT/test1/testing.txt"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}

func TestShowMeteDateFileInfo(t *testing.T) {
	newLocation := "./TT/test1/testing.txt"
	fileStat, err := os.Stat(newLocation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name:", fileStat.Name())        // Base name of the file
	fmt.Println("Size:", fileStat.Size())             // Length in bytes for regular files
	fmt.Println("Permissions:", fileStat.Mode())      // File mode bits
	fmt.Println("Last Modified:", fileStat.ModTime()) // Last modification time
	fmt.Println("Is Directory: ", fileStat.IsDir())   // Abbreviation for Mode().IsDir()
}

func TestDeleteFile(t *testing.T) {
	err := os.Remove("/var/www/html/test.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func TestReadCharacterFromFile(t *testing.T) {
	newLocation := "./TT/test1/testing.txt"
	fileBuffer, err := os.ReadFile(newLocation)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputData := string(fileBuffer)
	data := bufio.NewScanner(strings.NewReader(inputData))
	data.Split(bufio.ScanRunes)
	for data.Scan() {
		fmt.Print(data.Text())
	}
}

func TestOsTruncate(t *testing.T) {
	newLocation := "./TT/test1/testing.txt"
	configFile, e := os.OpenFile(newLocation, os.O_RDWR, 0666)
	defer configFile.Close()
	util.CheckErr(e)
	//some actions happen here
	configFile.Truncate(0)
	configFile.Write(util.S2ByteArray("ssssssssssssssssss"))
	configFile.Sync()
}

func TestModifyFileInfo(t *testing.T) {
	// Test File existence.
	_, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File exist.")

	// Change permissions Linux.
	err = os.Chmod("test.txt", 0777)
	if err != nil {
		log.Println(err)
	}

	// Change file ownership.
	err = os.Chown("test.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}

	// Change file timestamps.
	addOneDayFromNow := time.Now().Add(24 * time.Hour)
	lastAccessTime := addOneDayFromNow
	lastModifyTime := addOneDayFromNow
	err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}

}

func TestZipFile(t *testing.T) {
	zipReader, _ := zip.OpenReader("test.zip")
	for _, file := range zipReader.Reader.File {

		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		if file.FileInfo().IsDir() {
			log.Println("Directory Created:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			log.Println("File extracted:", file.Name)

			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
