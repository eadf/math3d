/*
This code is an incomplete port of the C++ algebra library WildMagic5 (geometrictools.com)
Note that this code uses column major matrixes, just like OpenGl
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
http://www.geometrictools.com/License/Boost/LICENSE_1_0.txt
*/

package math3d64
//This code is auto generated from the math3d32 package. Do not edit.
import "math"

import "fmt"


type Vector3 []float64

func MakeVector3(x, y, z float64) Vector3 {
	return Vector3{x, y, z}[:]
}

func MakeVector3V(v []float64) Vector3 {
	return Vector3{v[0], v[1], v[2]}[:]
}

// return v1+v2 (won't modify any of them)
func (v1 Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}[:]
}

// return v1-v2 (won't modify any of them)
func (v1 Vector3) Sub(v2 Vector3) Vector3 {
	return Vector3{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}[:]
}

func (v1 Vector3) Dot(v2 Vector3) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func (v1 Vector3) Cross(v2 Vector3) Vector3 {
	return Vector3{v1[1]*v2[2] - v1[2]*v2[1], v1[2]*v2[0] - v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}[:]
}

func (v Vector3) Length() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

// If two vectors represents points the distance between them can be calculated
func (v0 Vector3) Distance(v1 Vector3) float64 {
	d0 := v0[0] - v1[0]
	d1 := v0[1] - v1[1]
	d2 := v0[2] - v1[2]
	return math.Sqrt(d0*d0 + d1*d1 + d2*d2)
}

// If two vectors represents points the distance between them can be calculated
// Same as Distance(), just omitting the final sqrt() call
func (v0 Vector3) DistanceSqr(v1 Vector3) float64 {
	d0 := v0[0] - v1[0]
	d1 := v0[1] - v1[1]
	d2 := v0[2] - v1[2]
	return d0*d0 + d1*d1 + d2*d2
}

func (m Vector3) Copy() Vector3 {
	return Vector3{m[0], m[1], m[2]}[:]
}

// Normalize will modify this vector
func (v Vector3) Normalize() Vector3 {
	l := 1.0 / v.Length()
	v[0] *= l
	v[1] *= l
	v[2] *= l
	return v
}

func (m1 Vector3) Equal(q Vector3) bool {
	return m1[0] == q[0] && m1[1] == q[1] && m1[2] == q[2]
}

// seems rather pointless to have both tests..  todo
func (m1 Vector3) NotEqual(q Vector3) bool {
	return m1[0] != q[0] || m1[1] != q[1] || m1[2] != q[2]
}

/*
Tests to see if the difference between two matrices,
element-wise, exceeds ε.
*/
func (a Vector3) ApproxEquals(b Vector3, ε float64) bool {
	for i := 0; i < 3; i++ {
		if math.Fabs(a[i]-b[i]) > ε {
			return false
		}
	}
	return true
}

// untested
func (v Vector3) Yaw() float64 {
	return float64(-math.Atan2(float64(v[0]), float64(v[2])))
}

// untested
func (v Vector3) Pitch() float64 {
	return float64(-math.Atan2(float64(v[1]), math.Sqrt(float64(v[0])*float64(v[0])+float64(v[2])*float64(v[2]))))
}

// Multiply v (as a row vector) with the matrix m
func (v Vector3) MultiplyM(m Matrix3) Vector3 {
	return Vector3{v[0]*m[0] + v[1]*m[1] + v[2]*m[2],
		v[0]*m[3] + v[1]*m[4] + v[2]*m[5],
		v[0]*m[6] + v[1]*m[7] + v[2]*m[8]}[:]
}

func (v Vector3) String() string {
	return fmt.Sprintf("[%.5f,%.5f,%.5f]", v[0], v[1], v[2])
}

// p1,p2,p3 represents points
func SurfaceNormal(p1, p2, p3 Vector3) Vector3 {
	u := Vector3{p2[0] - p1[0], p2[1] - p1[2], p2[2] - p1[2]}
	v := Vector3{p3[0] - p1[0], p3[1] - p1[2], p3[2] - p1[2]}
	return Vector3{u[1]*v[2] - u[2]*v[1], u[2]*v[0] - u[0]*v[2], u[0]*v[1] - u[1]*v[0]}
}
