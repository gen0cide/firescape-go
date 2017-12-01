package main

import (
	"archive/zip"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/davecgh/go-spew/spew"
)

type Section struct {
	X        int
	Y        int
	Level    int
	SectionX int
	SectionY int
	Tiles    []Tile
	Data     []byte
}

type Tile struct {
	ID              int
	GroundElevation byte
	GroundTexture   byte
	RoofTexture     byte
	HorizontalWall  byte
	VerticalWall    byte
	GroundOverlay   byte
	DiagnalWall     int
}

var SectorIndex = 2304

func main() {
	// data, err := ioutil.ReadFile(os.Args[1])
	// if err != nil {
	// 	panic(err)
	// }

	sections := []Section{}

	r, err := zip.OpenReader(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer r.Close()

	for lvl := 0; lvl < 4; lvl++ {
		wildX := 2304
		wildY := 1776 - (lvl * 944)
		for sx := 0; sx < 1000; sx += 48 {
			for sy := 0; sy < 1000; sy += 48 {
				x := (sx + wildX) / 48
				y := (sy + (lvl * 944) + wildY) / 48
				newS := Section{
					Data:     LoadSection(x, y, lvl, sx, sy, r),
					X:        x,
					Y:        y,
					Level:    lvl,
					SectionX: sx,
					SectionY: sy,
					Tiles:    []Tile{},
				}
				newS.LoadTiles()
				sections = append(sections, newS)
				//fmt.Printf("SECTION: x=%d y=%d lvl=%d section_x=%d section_y=%d\n", x, y, lvl, sx, sy)

			}
		}
	}

	spew.Dump(sections[0])

}

func (s *Section) LoadTiles() {
	sectorBuffer := bytes.NewBuffer(s.Data)
	for i := 0; i < SectorIndex; i++ {
		tile := Tile{ID: i}
		tileData := sectorBuffer.Next(8)
		tile.GroundElevation = tileData[0]
		tile.GroundTexture = tileData[1]
		tile.GroundOverlay = tileData[2]
		tile.RoofTexture = tileData[3]
		tile.HorizontalWall = tileData[4]
		tile.VerticalWall = tileData[5]
		tile.DiagnalWall = int(binary.BigEndian.Uint16(tileData[6:]))
		s.Tiles = append(s.Tiles, tile)
	}
}

func LoadSection(x, y, h, sx, sy int, z *zip.ReadCloser) []byte {
	filename := fmt.Sprintf("h%dx%dy%d", h, x, y)
	for _, f := range z.File {
		if f.Name != filename {
			continue
		}
		buf := new(bytes.Buffer)
		stream, err := f.Open()
		if err != nil {
			panic(err)
		}
		defer stream.Close()
		io.Copy(buf, stream)
		return buf.Bytes()
	}
	return []byte{}
}
