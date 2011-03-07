/*
This code is an incomplete port of the C++ algebra library WildMagic5 (geometrictools.com)
Note that this code uses column major matrixes, just like OpenGl
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
http://www.geometrictools.com/License/Boost/LICENSE_1_0.txt
*/

package math3dd
//This code is auto generated from the math3df package. Do not edit.
import "math"



const internalε float64 = 0.000001
const internalεε float64 = internalε*internalε

const Rad2Deg float64 = float64(180.0/math.Pi)
const Deg2Rad float64 = float64(math.Pi/180.0)

// some ready converted float64 values
const Pi float64 = math.Pi
const TwoPi float64 = float64(math.Pi*2.)
const PiHalf float64 = float64(math.Pi*.5)
const Epsilon float64 = 0.000001

// these functions only exists so that we don't have to 
// use ugly float64() and float64() convertions all over the math3df code 
func Sin(a float64) float64 {
	return float64(math.Sin(float64(a)))
}

func Asinf(a float64) float64 {
	return float64(math.Asin(float64(a)))
}

func Cos(a float64) float64 {
	return float64(math.Cos(float64(a)))
}

func Acos(a float64) float64 {
	return float64(math.Acos(float64(a)))
}

func Fabs(a float64) float64 {
	return float64(math.Fabs(float64(a)))
}

// Signbit returns true if x is negative or negative zero.
func Signbit(a float64) bool {
	return math.Signbit(float64(a))
}

func Sqrt(a float64) float64 {
	return float64(math.Sqrt(float64(a)))
}

func Min(a,b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func Max(a,b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func AbsMin(a,b float64) float64 {
	if math.Fabs(a) < math.Fabs(b) {
		return a
	}
	return b
}

func AbsMin3(a,b,c float64) float64 {
	fabsa := math.Fabs(a)
	fabsb := math.Fabs(b)
	fabsc := math.Fabs(c)
	
	if fabsa < fabsb && fabsa < fabsc {
		return a
	}
	if fabsb < fabsa && fabsb < fabsc {
		return b
	}
	return c
}

func AbsMax(a,b float64) float64 {
	if math.Fabs(a) > math.Fabs(b) {
		return a
	}
	return b
}

// return the smallest angle between two radians
// if any of the angles are larger than -+2*Pi it won't work
func MinAngleBetween(a1,a2 float64) float64 {
	diff1 := a1-a2
	diff2 := a1-a2 + TwoPi	
	diff3 := a1-a2 - TwoPi
	
	return AbsMin3(diff1,diff2,diff3)
}

/*
func MinAngleBetweenVersion2(a1,a2 float64) float64 {
	// this solution does not care about the sign  
	var crossDiff, directDiff float64
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
func ApproxEquals(f1, f2 float64,ε float64) bool {
	if math.Fabs(f1-f2) > ε {
		//print ("diff is ", math.Fabs(f1-f2))
		return false
	}
	return true
}
