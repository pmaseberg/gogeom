package wkb

import (
	"encoding/binary"
	"github.com/pmaseberg/gogeom/geom"
	"io"
)

func multiLineStringReader(r io.Reader, byteOrder binary.ByteOrder) (geom.T, error) {
	var numLineStrings uint32
	if err := binary.Read(r, byteOrder, &numLineStrings); err != nil {
		return nil, err
	}
	lineStrings := make([]geom.LineString, numLineStrings)
	for i := uint32(0); i < numLineStrings; i++ {
		if g, err := Read(r); err == nil {
			var ok bool
			lineStrings[i], ok = g.(geom.LineString)
			if !ok {
				return nil, &UnexpectedGeometryError{g}
			}
		} else {
			return nil, err
		}
	}
	return geom.MultiLineString{LineStrings: lineStrings}, nil
}

func writeMultiLineString(w io.Writer, byteOrder binary.ByteOrder, multiLineString geom.MultiLineString) error {
	if err := binary.Write(w, byteOrder, uint32(len(multiLineString.LineStrings))); err != nil {
		return err
	}
	for _, lineString := range multiLineString.LineStrings {
		if err := Write(w, byteOrder, lineString); err != nil {
			return err
		}
	}
	return nil
}

func multiLineStringZReader(r io.Reader, byteOrder binary.ByteOrder) (geom.T, error) {
	var numLineStrings uint32
	if err := binary.Read(r, byteOrder, &numLineStrings); err != nil {
		return nil, err
	}
	lineStringZs := make([]geom.LineStringZ, numLineStrings)
	for i := uint32(0); i < numLineStrings; i++ {
		if g, err := Read(r); err == nil {
			var ok bool
			lineStringZs[i], ok = g.(geom.LineStringZ)
			if !ok {
				return nil, &UnexpectedGeometryError{g}
			}
		} else {
			return nil, err
		}
	}
	return geom.MultiLineStringZ{LineStrings: lineStringZs}, nil
}

func writeMultiLineStringZ(w io.Writer, byteOrder binary.ByteOrder, multiLineStringZ geom.MultiLineStringZ) error {
	if err := binary.Write(w, byteOrder, uint32(len(multiLineStringZ.LineStrings))); err != nil {
		return err
	}
	for _, lineStringZ := range multiLineStringZ.LineStrings {
		if err := Write(w, byteOrder, lineStringZ); err != nil {
			return err
		}
	}
	return nil
}

func multiLineStringMReader(r io.Reader, byteOrder binary.ByteOrder) (geom.T, error) {
	var numLineStrings uint32
	if err := binary.Read(r, byteOrder, &numLineStrings); err != nil {
		return nil, err
	}
	lineStringMs := make([]geom.LineStringM, numLineStrings)
	for i := uint32(0); i < numLineStrings; i++ {
		if g, err := Read(r); err == nil {
			var ok bool
			lineStringMs[i], ok = g.(geom.LineStringM)
			if !ok {
				return nil, &UnexpectedGeometryError{g}
			}
		} else {
			return nil, err
		}
	}
	return geom.MultiLineStringM{LineStrings: lineStringMs}, nil
}

func writeMultiLineStringM(w io.Writer, byteOrder binary.ByteOrder, multiLineStringM geom.MultiLineStringM) error {
	if err := binary.Write(w, byteOrder, uint32(len(multiLineStringM.LineStrings))); err != nil {
		return err
	}
	for _, lineStringM := range multiLineStringM.LineStrings {
		if err := Write(w, byteOrder, lineStringM); err != nil {
			return err
		}
	}
	return nil
}

func multiLineStringZMReader(r io.Reader, byteOrder binary.ByteOrder) (geom.T, error) {
	var numLineStrings uint32
	if err := binary.Read(r, byteOrder, &numLineStrings); err != nil {
		return nil, err
	}
	lineStringZMs := make([]geom.LineStringZM, numLineStrings)
	for i := uint32(0); i < numLineStrings; i++ {
		if g, err := Read(r); err == nil {
			var ok bool
			lineStringZMs[i], ok = g.(geom.LineStringZM)
			if !ok {
				return nil, &UnexpectedGeometryError{g}
			}
		} else {
			return nil, err
		}
	}
	return geom.MultiLineStringZM{LineStrings: lineStringZMs}, nil
}

func writeMultiLineStringZM(w io.Writer, byteOrder binary.ByteOrder, multiLineStringZM geom.MultiLineStringZM) error {
	if err := binary.Write(w, byteOrder, uint32(len(multiLineStringZM.LineStrings))); err != nil {
		return err
	}
	for _, lineStringZM := range multiLineStringZM.LineStrings {
		if err := Write(w, byteOrder, lineStringZM); err != nil {
			return err
		}
	}
	return nil
}
