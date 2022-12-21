package day18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	scan := parseMany(sample)
	enclosed := enclosedCells(scan)
	a.Len(enclosed, 1)
	a.ElementsMatch(enclosed, []cube{{2, 2, 5}})
	faces := externalFaces(scan, enclosed)
	a.Equal(58, faces)
}

func TestPart2(t *testing.T) {
	scan := parseMany(input)
	enclosed := enclosedCells(scan)
	faces := externalFaces(scan, enclosed)
	t.Log(faces)
}

func externalFaces(scan, enclosed []cube) int {
	// external faces are those in the scan minus the faces shared with another
	// scan cell, minus the faces shared with an enclosed cell. faces shared
	// between scna and enclosed cells are shared faces for that combined set
	// minus shared faces just within the enclosed set.
	scanShared := countSharedFaces(scan)
	scan2 := append(scan, enclosed...)
	totalShared := countSharedFaces(scan2)
	enclosedShared := countSharedFaces(enclosed)
	// faces shared between scan & enclosed cells = total shared among the
	// combined set minus how many are with just the scan or enclosed sets
	scanEnclosedShared := totalShared - scanShared - enclosedShared
	return 6*len(scan) - // total possible scan faces
		2*scanShared - // ... minus interior faces between scan cells
		scanEnclosedShared // ... minus faces shared between a scan & enclosed cell
}
