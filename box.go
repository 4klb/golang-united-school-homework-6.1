package golang_united_school_homework

import (
	"errors"
	"reflect"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity != 0 {
		b.shapes = append(b.shapes, shape)
		b.shapesCapacity--
	} else {
		return errors.New("It goes out of the shapesCapacity range")
	}
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	var shape Shape

	for indx, value := range b.shapes {
		if indx == i {
			return value, nil
		}
	}
	return shape, errors.New("Shape by index doesn't exist or index went out of the range")
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	var err error
	var deletedShape Shape
	var exist bool

	for indx, value := range b.shapes {
		if indx == i {
			deletedShape = value
			value = shape
			exist = true
			// return deletedShape, nil
		}
	}
	if !exist {
		err = errors.New("Shape by index doesn't exist or index went out of the range")
	} else {
		err = nil
	}
	return deletedShape, err
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, value := range b.shapes {
		per := value.CalcPerimeter()
		sum += per
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, value := range b.shapes {
		per := value.CalcArea()
		sum += per
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var err error
	var circle Circle
	var exist bool

	for indx, _ := range b.shapes {
		if reflect.TypeOf(b.shapes[indx]) == reflect.TypeOf(circle) {
			copy(b.shapes[indx:], b.shapes[indx+1:])
			b.shapes = b.shapes[:len(b.shapes)-1]
			exist = true
			indx--
		}
	}

	if !exist {
		err = errors.New("Circles are not exist in the list")
	} else {
		err = nil
	}
	return err
}

//ExtractByIndex allows getting shape by index and removes this shape from the list.
//whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	var err error
	var exist bool
	var shape Shape

	for indx, value := range b.shapes {
		if indx == i {
			copy(b.shapes[indx:], b.shapes[indx+1:])
			b.shapes = b.shapes[:len(b.shapes)-1]
			shape = value
			exist = true
			indx--
		}
	}

	if !exist {
		err = errors.New("Shape by index doesn't exist or index went out of the range")
	} else {
		err = nil
	}
	return shape, err
}
