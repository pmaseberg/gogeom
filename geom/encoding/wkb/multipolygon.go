package wkb

import (
	"encoding/binary"
	"github.com/pmaseberg/gogeom/geom"
	"io"
)

func multiPolygonReader(r io.Reader, byteOrder binary.ByteOrder) (geom.T, error) {
	var numPolygons uint32
	if err := binary.Read(r, byteOrder, &numPolygons); err != nil {
		return nil, err
	}
	polygons := make([]geom.Polygon, numPolygons)
	for i := uint32(0); i < numPolygons; i++ {
		if g, err := Read(r); err == nil {
			var ok bool
			polygons[i], ok = g.(geom.Polygon)
			if !ok {
				return nil, &UnexpectedGeometryError{g}
			}
		} else {
			return nil, err
		}
	}
	return geom.MultiPolygon{Polygons: polygons}, nil
}

func writeMultiPolygon(w io.Writer, byteOrder binary.ByteOrder, multiPolygon geom.MultiPolygon) error {
	if err := binary.Write(w, byteOrder, uint32(len(multiPolygon.Polygons))); err != nil {
		return err
	}
	for _, polygon := range multiPolygon.Polygons {
		if err := Write(w, byteOrder, polygon); err != nil {
			return err
		}
	}
	return nil
}

func multiPolygonZReader(r io.Reader, byteOrder binary.ByteOrder) (geom.T, error) {
	var numPolygons uint32
	if err := binary.Read(r, byteOrder, &numPolygons); err != nil {
		return nil, err
	}
	polygonZs := make([]geom.PolygonZ, numPolygons)
	for i := uint32(0); i < numPolygons; i++ {
		if g, err := Read(r); err == nil {
			var ok bool
			polygonZs[i], ok = g.(geom.PolygonZ)
			if !ok {
				return nil, &UnexpectedGeometryError{g}
			}
		} else {
			return nil, err
		}
	}
	return geom.MultiPolygonZ{Polygons: polygonZs}, nil
}

func writeMultiPolygonZ(w io.Writer, byteOrder binary.ByteOrder, multiPolygonZ geom.MultiPolygonZ) error {
	if err := binary.Write(w, byteOrder, uint32(len(multiPolygonZ.Polygons))); err != nil {
		return err
	}
	for _, polygonZ := range multiPolygonZ.Polygons {
		if err := Write(w, byteOrder, polygonZ); err != nil {
			return err
		}
	}
	return nil
}

func multiPolygonMReader(r io.Reader, byteOrder binary.ByteOrder) (geom.T, error) {
	var numPolygons uint32
	if err := binary.Read(r, byteOrder, &numPolygons); err != nil {
		return nil, err
	}
	polygonMs := make([]geom.PolygonM, numPolygons)
	for i := uint32(0); i < numPolygons; i++ {
		if g, err := Read(r); err == nil {
			var ok bool
			polygonMs[i], ok = g.(geom.PolygonM)
			if !ok {
				return nil, &UnexpectedGeometryError{g}
			}
		} else {
			return nil, err
		}
	}
	return geom.MultiPolygonM{Polygons: polygonMs}, nil
}

func writeMultiPolygonM(w io.Writer, byteOrder binary.ByteOrder, multiPolygonM geom.MultiPolygonM) error {
	if err := binary.Write(w, byteOrder, uint32(len(multiPolygonM.Polygons))); err != nil {
		return err
	}
	for _, polygonM := range multiPolygonM.Polygons {
		if err := Write(w, byteOrder, polygonM); err != nil {
			return err
		}
	}
	return nil
}

func multiPolygonZMReader(r io.Reader, byteOrder binary.ByteOrder) (geom.T, error) {
	var numPolygons uint32
	if err := binary.Read(r, byteOrder, &numPolygons); err != nil {
		return nil, err
	}
	polygonZMs := make([]geom.PolygonZM, numPolygons)
	for i := uint32(0); i < numPolygons; i++ {
		if g, err := Read(r); err == nil {
			var ok bool
			polygonZMs[i], ok = g.(geom.PolygonZM)
			if !ok {
				return nil, &UnexpectedGeometryError{g}
			}
		} else {
			return nil, err
		}
	}
	return geom.MultiPolygonZM{Polygons: polygonZMs}, nil
}

func writeMultiPolygonZM(w io.Writer, byteOrder binary.ByteOrder, multiPolygonZM geom.MultiPolygonZM) error {
	if err := binary.Write(w, byteOrder, uint32(len(multiPolygonZM.Polygons))); err != nil {
		return err
	}
	for _, polygonZM := range multiPolygonZM.Polygons {
		if err := Write(w, byteOrder, polygonZM); err != nil {
			return err
		}
	}
	return nil
}
