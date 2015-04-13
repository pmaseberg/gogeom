package wkb

import (
	"encoding/binary"
	"github.com/pmaseberg/gogeom/geom"
	"io"
)

func polygonReader(r io.Reader, byteOrder binary.ByteOrder) (geom.T, error) {
	var numRings uint32
	if err := binary.Read(r, byteOrder, &numRings); err != nil {
		return nil, err
	}
	rings := make([][]geom.Point, numRings)
	for i := uint32(0); i < numRings; i++ {
		if points, err := readPoints(r, byteOrder); err != nil {
			return nil, err
		} else {
			rings[i] = points
		}
	}
	return geom.Polygon{Rings: rings}, nil
}

func writePolygon(w io.Writer, byteOrder binary.ByteOrder, polygon geom.Polygon) error {
	return writePointss(w, byteOrder, polygon.Rings)
}

func polygonZReader(r io.Reader, byteOrder binary.ByteOrder) (geom.T, error) {
	var numRings uint32
	if err := binary.Read(r, byteOrder, &numRings); err != nil {
		return nil, err
	}
	ringZs := make([][]geom.PointZ, numRings)
	for i := uint32(0); i < numRings; i++ {
		if pointZs, err := readPointZs(r, byteOrder); err != nil {
			return nil, err
		} else {
			ringZs[i] = pointZs
		}
	}
	return geom.PolygonZ{Rings: ringZs}, nil
}

func writePolygonZ(w io.Writer, byteOrder binary.ByteOrder, polygonZ geom.PolygonZ) error {
	return writePointZss(w, byteOrder, polygonZ.Rings)
}

func polygonMReader(r io.Reader, byteOrder binary.ByteOrder) (geom.T, error) {
	var numRings uint32
	if err := binary.Read(r, byteOrder, &numRings); err != nil {
		return nil, err
	}
	ringMs := make([][]geom.PointM, numRings)
	for i := uint32(0); i < numRings; i++ {
		if pointMs, err := readPointMs(r, byteOrder); err != nil {
			return nil, err
		} else {
			ringMs[i] = pointMs
		}
	}
	return geom.PolygonM{Rings: ringMs}, nil
}

func writePolygonM(w io.Writer, byteOrder binary.ByteOrder, polygonM geom.PolygonM) error {
	return writePointMss(w, byteOrder, polygonM.Rings)
}

func polygonZMReader(r io.Reader, byteOrder binary.ByteOrder) (geom.T, error) {
	var numRings uint32
	if err := binary.Read(r, byteOrder, &numRings); err != nil {
		return nil, err
	}
	ringZMs := make([][]geom.PointZM, numRings)
	for i := uint32(0); i < numRings; i++ {
		if pointZMs, err := readPointZMs(r, byteOrder); err != nil {
			return nil, err
		} else {
			ringZMs[i] = pointZMs
		}
	}
	return geom.PolygonZM{Rings: ringZMs}, nil
}

func writePolygonZM(w io.Writer, byteOrder binary.ByteOrder, polygonZM geom.PolygonZM) error {
	return writePointZMss(w, byteOrder, polygonZM.Rings)
}
