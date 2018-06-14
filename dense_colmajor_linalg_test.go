package tensor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var colMajorTraceTests = []struct {
	data interface{}

	correct interface{}
	err     bool
}{
	{[]int{0, 1, 2, 3, 4, 5}, int(4), false},
	{[]int8{0, 1, 2, 3, 4, 5}, int8(4), false},
	{[]int16{0, 1, 2, 3, 4, 5}, int16(4), false},
	{[]int32{0, 1, 2, 3, 4, 5}, int32(4), false},
	{[]int64{0, 1, 2, 3, 4, 5}, int64(4), false},
	{[]uint{0, 1, 2, 3, 4, 5}, uint(4), false},
	{[]uint8{0, 1, 2, 3, 4, 5}, uint8(4), false},
	{[]uint16{0, 1, 2, 3, 4, 5}, uint16(4), false},
	{[]uint32{0, 1, 2, 3, 4, 5}, uint32(4), false},
	{[]uint64{0, 1, 2, 3, 4, 5}, uint64(4), false},
	{[]float32{0, 1, 2, 3, 4, 5}, float32(4), false},
	{[]float64{0, 1, 2, 3, 4, 5}, float64(4), false},
	{[]complex64{0, 1, 2, 3, 4, 5}, complex64(4), false},
	{[]complex128{0, 1, 2, 3, 4, 5}, complex128(4), false},
	{[]bool{true, false, true, false, true, false}, nil, true},
}

func TestColMajor_Dense_Trace(t *testing.T) {
	assert := assert.New(t)
	for i, tts := range colMajorTraceTests {
		T := New(WithShape(2, 3), AsFortran(tts.data))
		trace, err := T.Trace()

		if checkErr(t, tts.err, err, "Trace", i) {
			continue
		}
		assert.Equal(tts.correct, trace)

		//
		T = New(WithBacking(tts.data))
		_, err = T.Trace()
		if err == nil {
			t.Error("Expected an error when Trace() on non-matrices")
		}
	}
}

var colMajorInnerTests = []struct {
	a, b           interface{}
	shapeA, shapeB Shape

	correct interface{}
	err     bool
}{
	{Range(Float64, 0, 3), Range(Float64, 0, 3), Shape{3}, Shape{3}, float64(5), false},
	{Range(Float64, 0, 3), Range(Float64, 0, 3), Shape{3, 1}, Shape{3}, float64(5), false},
	{Range(Float64, 0, 3), Range(Float64, 0, 3), Shape{1, 3}, Shape{3}, float64(5), false},
	{Range(Float64, 0, 3), Range(Float64, 0, 3), Shape{3, 1}, Shape{3, 1}, float64(5), false},
	{Range(Float64, 0, 3), Range(Float64, 0, 3), Shape{1, 3}, Shape{3, 1}, float64(5), false},
	{Range(Float64, 0, 3), Range(Float64, 0, 3), Shape{3, 1}, Shape{1, 3}, float64(5), false},
	{Range(Float64, 0, 3), Range(Float64, 0, 3), Shape{1, 3}, Shape{1, 3}, float64(5), false},

	{Range(Float32, 0, 3), Range(Float32, 0, 3), Shape{3}, Shape{3}, float32(5), false},
	{Range(Float32, 0, 3), Range(Float32, 0, 3), Shape{3, 1}, Shape{3}, float32(5), false},
	{Range(Float32, 0, 3), Range(Float32, 0, 3), Shape{1, 3}, Shape{3}, float32(5), false},
	{Range(Float32, 0, 3), Range(Float32, 0, 3), Shape{3, 1}, Shape{3, 1}, float32(5), false},
	{Range(Float32, 0, 3), Range(Float32, 0, 3), Shape{1, 3}, Shape{3, 1}, float32(5), false},
	{Range(Float32, 0, 3), Range(Float32, 0, 3), Shape{3, 1}, Shape{1, 3}, float32(5), false},
	{Range(Float32, 0, 3), Range(Float32, 0, 3), Shape{1, 3}, Shape{1, 3}, float32(5), false},

	// stupids: type differences
	{Range(Int, 0, 3), Range(Float64, 0, 3), Shape{3}, Shape{3}, nil, true},
	{Range(Float32, 0, 3), Range(Byte, 0, 3), Shape{3}, Shape{3}, nil, true},
	{Range(Float64, 0, 3), Range(Float32, 0, 3), Shape{3}, Shape{3}, nil, true},
	{Range(Float32, 0, 3), Range(Float64, 0, 3), Shape{3}, Shape{3}, nil, true},

	// differing size
	{Range(Float64, 0, 4), Range(Float64, 0, 3), Shape{4}, Shape{3}, nil, true},

	// A is not a matrix
	{Range(Float64, 0, 4), Range(Float64, 0, 3), Shape{2, 2}, Shape{3}, nil, true},
}

func TestColMajor_Dense_Inner(t *testing.T) {
	for i, its := range colMajorInnerTests {
		a := New(WithShape(its.shapeA...), AsFortran(its.a))
		b := New(WithShape(its.shapeB...), AsFortran(its.b))

		T, err := a.Inner(b)
		if checkErr(t, its.err, err, "Inner", i) {
			continue
		}

		assert.Equal(t, its.correct, T)
	}
}

var colMajorMatVecMulTests = []linalgTest{
	// Float64s
	{Range(Float64, 0, 6), Range(Float64, 0, 3), Shape{2, 3}, Shape{3}, false, false,
		Range(Float64, 52, 54), Range(Float64, 100, 102), Shape{2}, Shape{2},
		[]float64{5, 14}, []float64{105, 115}, []float64{110, 129}, Shape{2}, false, false, false},
	{Range(Float64, 0, 6), Range(Float64, 0, 3), Shape{2, 3}, Shape{3, 1}, false, false,
		Range(Float64, 52, 54), Range(Float64, 100, 102), Shape{2}, Shape{2},
		[]float64{5, 14}, []float64{105, 115}, []float64{110, 129}, Shape{2}, false, false, false},
	{Range(Float64, 0, 6), Range(Float64, 0, 3), Shape{2, 3}, Shape{1, 3}, false, false,
		Range(Float64, 52, 54), Range(Float64, 100, 102), Shape{2}, Shape{2},
		[]float64{5, 14}, []float64{105, 115}, []float64{110, 129}, Shape{2}, false, false, false},

	// float64s with transposed matrix
	{Range(Float64, 0, 6), Range(Float64, 0, 2), Shape{2, 3}, Shape{2}, true, false,
		Range(Float64, 52, 55), Range(Float64, 100, 103), Shape{3}, Shape{3},
		[]float64{3, 4, 5}, []float64{103, 105, 107}, []float64{106, 109, 112}, Shape{3}, false, false, false},

	// Float32s
	{Range(Float32, 0, 6), Range(Float32, 0, 3), Shape{2, 3}, Shape{3}, false, false,
		Range(Float32, 52, 54), Range(Float32, 100, 102), Shape{2}, Shape{2},
		[]float32{5, 14}, []float32{105, 115}, []float32{110, 129}, Shape{2}, false, false, false},
	{Range(Float32, 0, 6), Range(Float32, 0, 3), Shape{2, 3}, Shape{3, 1}, false, false,
		Range(Float32, 52, 54), Range(Float32, 100, 102), Shape{2}, Shape{2},
		[]float32{5, 14}, []float32{105, 115}, []float32{110, 129}, Shape{2}, false, false, false},
	{Range(Float32, 0, 6), Range(Float32, 0, 3), Shape{2, 3}, Shape{1, 3}, false, false,
		Range(Float32, 52, 54), Range(Float32, 100, 102), Shape{2}, Shape{2},
		[]float32{5, 14}, []float32{105, 115}, []float32{110, 129}, Shape{2}, false, false, false},

	// stupids : unpossible shapes (wrong A)
	{Range(Float64, 0, 6), Range(Float64, 0, 3), Shape{6}, Shape{3}, false, false,
		Range(Float64, 52, 54), Range(Float64, 100, 102), Shape{2}, Shape{2},
		[]float64{5, 14}, []float64{105, 115}, []float64{110, 129}, Shape{2}, true, false, false},

	//stupids: bad A shape
	{Range(Float64, 0, 8), Range(Float64, 0, 3), Shape{4, 2}, Shape{3}, false, false,
		Range(Float64, 52, 54), Range(Float64, 100, 102), Shape{2}, Shape{2},
		[]float64{5, 14}, []float64{105, 115}, []float64{110, 129}, Shape{2}, true, false, false},

	//stupids: bad B shape
	{Range(Float64, 0, 6), Range(Float64, 0, 6), Shape{2, 3}, Shape{3, 2}, false, false,
		Range(Float64, 52, 54), Range(Float64, 100, 102), Shape{2}, Shape{2},
		[]float64{5, 14}, []float64{105, 115}, []float64{110, 129}, Shape{2}, true, false, false},

	//stupids: bad reuse
	{Range(Float64, 0, 6), Range(Float64, 0, 3), Shape{2, 3}, Shape{3}, false, false,
		Range(Float64, 52, 55), Range(Float64, 100, 102), Shape{3}, Shape{2},
		[]float64{5, 14}, []float64{105, 115}, []float64{110, 129}, Shape{2}, false, false, true},

	//stupids: bad incr shape
	{Range(Float64, 0, 6), Range(Float64, 0, 3), Shape{2, 3}, Shape{3}, false, false,
		Range(Float64, 52, 54), Range(Float64, 100, 105), Shape{2}, Shape{5},
		[]float64{5, 14}, []float64{105, 115}, []float64{110, 129}, Shape{2}, false, true, false},

	// stupids: type mismatch A and B
	{Range(Float64, 0, 6), Range(Float32, 0, 3), Shape{2, 3}, Shape{3}, false, false,
		Range(Float64, 52, 54), Range(Float64, 100, 103), Shape{2}, Shape{3},
		[]float64{5, 14}, []float64{105, 115}, []float64{110, 129}, Shape{2}, true, false, false},

	// stupids: type mismatch A and B
	{Range(Float32, 0, 6), Range(Float64, 0, 3), Shape{2, 3}, Shape{3}, false, false,
		Range(Float64, 52, 54), Range(Float64, 100, 103), Shape{2}, Shape{3},
		[]float64{5, 14}, []float64{105, 115}, []float64{110, 129}, Shape{2}, true, false, false},

	// stupids: type mismatch A and B
	{Range(Float64, 0, 6), Range(Float32, 0, 3), Shape{2, 3}, Shape{3}, false, false,
		Range(Float64, 52, 54), Range(Float64, 100, 103), Shape{2}, Shape{3},
		[]float64{5, 14}, []float64{105, 115}, []float64{110, 129}, Shape{2}, true, false, false},

	// stupids: type mismatch A and B
	{Range(Float32, 0, 6), Range(Float64, 0, 3), Shape{2, 3}, Shape{3}, false, false,
		Range(Float64, 52, 54), Range(Float64, 100, 103), Shape{2}, Shape{3},
		[]float64{5, 14}, []float64{105, 115}, []float64{110, 129}, Shape{2}, true, false, false},

	// stupids: type mismatch A and B (non-Float)
	{Range(Float64, 0, 6), Range(Int, 0, 3), Shape{2, 3}, Shape{3}, false, false,
		Range(Float64, 52, 54), Range(Float64, 100, 103), Shape{2}, Shape{3},
		[]float64{5, 14}, []float64{105, 115}, []float64{110, 129}, Shape{2}, true, false, false},

	// stupids: type mismatch, reuse
	{Range(Float64, 0, 6), Range(Float64, 0, 3), Shape{2, 3}, Shape{3}, false, false,
		Range(Float32, 52, 54), Range(Float64, 100, 102), Shape{2}, Shape{2},
		[]float64{5, 14}, []float64{105, 115}, []float64{110, 129}, Shape{2}, false, false, true},

	// stupids: type mismatch, incr
	{Range(Float64, 0, 6), Range(Float64, 0, 3), Shape{2, 3}, Shape{3}, false, false,
		Range(Float64, 52, 54), Range(Float32, 100, 103), Shape{2}, Shape{3},
		[]float64{5, 14}, []float64{105, 115}, []float64{110, 129}, Shape{2}, false, true, false},

	// stupids: type mismatch, incr not a Number
	{Range(Float64, 0, 6), Range(Float64, 0, 3), Shape{2, 3}, Shape{3}, false, false,
		Range(Float64, 52, 54), []bool{true, true, true}, Shape{2}, Shape{3},
		[]float64{5, 14}, []float64{105, 115}, []float64{110, 129}, Shape{2}, false, true, false},
}

func TestColMajor_Dense_MatVecMul(t *testing.T) {
	assert := assert.New(t)
	for i, mvmt := range colMajorMatVecMulTests {
		a := New(WithShape(mvmt.shapeA...), AsFortran(mvmt.a))
		b := New(WithShape(mvmt.shapeB...), AsFortran(mvmt.b))

		if mvmt.transA {
			if err := a.T(); err != nil {
				t.Error(err)
				continue
			}
		}

		T, err := a.MatVecMul(b)
		if checkErr(t, mvmt.err, err, "Safe", i) {
			continue
		}

		assert.True(mvmt.correctShape.Eq(T.Shape()))
		assert.Equal(mvmt.correct, T.Data())

		// incr
		incr := New(WithShape(mvmt.shapeI...), AsFortran(mvmt.incr))
		T, err = a.MatVecMul(b, WithIncr(incr))
		if checkErr(t, mvmt.errIncr, err, "WithIncr", i) {
			continue
		}

		assert.True(mvmt.correctShape.Eq(T.Shape()))
		assert.Equal(mvmt.correctIncr, T.Data())

		// reuse
		reuse := New(WithShape(mvmt.shapeR...), AsFortran(mvmt.reuse))
		T, err = a.MatVecMul(b, WithReuse(reuse))
		if checkErr(t, mvmt.errReuse, err, "WithReuse", i) {
			continue
		}

		assert.True(mvmt.correctShape.Eq(T.Shape()))
		assert.Equal(mvmt.correct, T.Data())

		// reuse AND incr
		T, err = a.MatVecMul(b, WithIncr(incr), WithReuse(reuse))
		if checkErr(t, mvmt.err, err, "WithReuse and WithIncr", i) {
			continue
		}
		assert.True(mvmt.correctShape.Eq(T.Shape()))
		assert.Equal(mvmt.correctIncrReuse, T.Data())
	}
}
