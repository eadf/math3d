/*
This code is an incomplete port of the C++ algebra library WildMagic5 (geometrictools.com)
Note that this code uses column major matrixes, just like OpenGl
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
http://www.geometrictools.com/License/Boost/LICENSE_1_0.txt
*/

package math3d32

import "math"

const internalε float32 = 0.000001
const internalεε float32 = internalε * internalε

const Rad2Deg float32 = float32(180.0 / math.Pi)
const Deg2Rad float32 = float32(math.Pi / 180.0)

// some ready converted float32 values
const Pi float32 = float32(math.Pi)
const TwoPi float32 = float32(math.Pi * 2.)
const PiHalf float32 = float32(math.Pi * .5)
const Epsilon float32 = 0.000001

// these functions only exists so that we don't have to 
// use ugly float32() and float64() convertions all over the math3d32 code 
func Sinf(a float32) float32 {
	return float32(math.Sin(float64(a)))
}

func Asinf(a float32) float32 {
	return float32(math.Asin(float64(a)))
}

func Cosf(a float32) float32 {
	return float32(math.Cos(float64(a)))
}

func Acosf(a float32) float32 {
	return float32(math.Acos(float64(a)))
}

func Fabsf(a float32) float32 {
	return float32(math.Fabs(float64(a)))
}

// Signbit returns true if x is negative or negative zero.
func Signbit(a float32) bool {
	return math.Signbit(float64(a))
}

func Sqrtf(a float32) float32 {
	return float32(math.Sqrt(float64(a)))
}

func Min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func AbsMin(a, b float32) float32 {
	if Fabsf(a) < Fabsf(b) {
		return a
	}
	return b
}

func AbsMin3(a, b, c float32) float32 {
	fabsa := Fabsf(a)
	fabsb := Fabsf(b)
	fabsc := Fabsf(c)

	if fabsa < fabsb && fabsa < fabsc {
		return a
	}
	if fabsb < fabsa && fabsb < fabsc {
		return b
	}
	return c
}

func AbsMax(a, b float32) float32 {
	if Fabsf(a) > Fabsf(b) {
		return a
	}
	return b
}

// return the smallest angle between two radians
// if any of the angles are larger than -+2*Pi it won't work
func MinAngleBetween(a1, a2 float32) float32 {
	diff1 := a1 - a2
	diff2 := a1 - a2 + TwoPi
	diff3 := a1 - a2 - TwoPi

	return AbsMin3(diff1, diff2, diff3)
}

/*
func MinAngleBetweenVersion2(a1,a2 float32) float32 {
	// this solution does not care about the sign  
	var crossDiff, directDiff float32
	if a1 > a2 {
		crossDiff = TwoPi - a1 + a2
		directDiff = a1 - a2
	} else {
		crossDiff = TwoPi - a2 + a1
		directDiff = a2 - a1
	}
	if crossDiff < directDiff {
		return crossDiff
	}
	return directDiff
}
*/

/*
Tests to see if the difference between two floats exceeds ε.
*/
func ApproxEquals(f1, f2 float32, ε float32) bool {
	if Fabsf(f1-f2) > ε {
		//print ("diff is ", Fabsf(f1-f2))
		return false
	}
	return true
}
