package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

/*
Reference

- xml parser : https://tutorialedge.net/golang/parsing-xml-with-golang/
- file exists : https://golangcode.com/check-if-a-file-exists/
- resolve relative path : https://stackoverflow.com/questions/47261719/how-can-i-resolve-a-relative-path-to-absolute-path-in-golang
*/

type GameList struct {
	XMLName xml.Name `xml:"gameList"`
	GameList []Game `xml:"game"`
}
type Game struct {
	XMLName xml.Name `xml:"game"`
	Path string `xml:"path"`
	Name string `xml:"name"`
	Favorite bool `xml:"favorite"`
}

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(src, dst string) (err error) {
    sfi, err := os.Stat(src)
    if err != nil {
        return
    }
    if !sfi.Mode().IsRegular() {
        // cannot copy non-regular files (e.g., directories,
        // symlinks, devices, etc.)
        return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
    }
    dfi, err := os.Stat(dst)
    if err != nil {
        if !os.IsNotExist(err) {
            return
        }
    } else {
        if !(dfi.Mode().IsRegular()) {
            return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
        }
        if os.SameFile(sfi, dfi) {
            return
        }
    }
    if err = os.Link(src, dst); err == nil {
        return
    }
    err = copyFileContents(src, dst)
    return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
    in, err := os.Open(src)
    if err != nil {
        return
    }
    defer in.Close()
    out, err := os.Create(dst)
    if err != nil {
        return
    }
    defer func() {
        cerr := out.Close()
        if err == nil {
            err = cerr
        }
    }()
    if _, err = io.Copy(out, in); err != nil {
        return
    }
    err = out.Sync()
    return
}


func main() {
    xmlFile, err := os.Open("/Users/hmc/workspaces/zerocool/gamelist.xml")    
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened /Users/hmc/workspaces/zerocool/gamelist.xml")    
    defer xmlFile.Close()

	bytes, _ := ioutil.ReadAll(xmlFile)
	fmt.Println("bytes size:", len(bytes))

	var gamelist GameList
	xml.Unmarshal(bytes, &gamelist)

	var favTotal int = 0
	var romNotFound int = 0
	var baseRomPath = []string{
		"/Users/hmc/game/raspberry/roms/MAME32v0120/roms",
		"/Users/hmc/game/raspberry/roms/favorites",
	}

	for i:=0; i< len(gamelist.GameList); i++ {
		var fav string
		if gamelist.GameList[i].Favorite==true {
			fav = "O"
		} else { 
			fav = "X"
		}
		if (gamelist.GameList[i].Favorite==true){			
			fmt.Println("[favorite:", fav, "] name:", gamelist.GameList[i].Name, ", path:", gamelist.GameList[i].Path)

			romFound := false
			for j, basePath := range baseRomPath {
				romFile := path.Join(basePath, gamelist.GameList[i].Path)
				romFileExists := fileExists(romFile)
				fmt.Println("\t [",j,"] romFile:", romFile, " -> ", romFileExists)

				if (romFileExists) {
					romFound = true
					CopyFile(romFile, path.Join("/Users/hmc/game/raspberry/roms/favorites-210318", gamelist.GameList[i].Path))
					break
				}
			}
			favTotal++
			if (!romFound) {
				romNotFound++
			}
		}
	}
	fmt.Println("Total ", favTotal, " fovorite roms found. (romNotFound:", romNotFound, ")")
}