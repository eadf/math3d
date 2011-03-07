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

import "fmt"

type Vector4 []float64

func MakeVector4V(v[]float64)(Vector4){
	return Vector4{v[0],v[1],v[2],v[3]}[:]
}

func MakeVector4(x,y,z,o float64)(Vector4){
	v:=Vector4{x,y,z,o} 
	return v
}

// return v1+v2 (won't modify any of them)
func (v1 Vector4) Add(v2 Vector4) Vector4 {
	return Vector4{v1[0]+v2[0],v1[1]+v2[1],v1[2]+v2[2],v1[3]+v2[3]}
}

// return v1-v2 (won't modify any of them)
func (v1 Vector4) Sub(v2 Vector4) Vector4 {
	return Vector4{v1[0]-v2[0],v1[1]-v2[1],v1[2]-v2[2],v1[3]-v2[3]}
}

func (v1 Vector4) Dot(v2 Vector4) float64 {
	return v1[0]*v2[0]+v1[1]*v2[1]+v1[2]*v2[2]
}

func (v1 Vector4) Cross(v2 Vector4) Vector4 {
	return Vector4{v1[1]*v2[2]-v1[2]*v2[1],v1[2]*v2[0]-v1[0]*v2[2],v1[0]*v2[1]-v1[1]*v2[0]} 
}

/*
// For those cases when the 4d vector represents just a 3d vector. 4:t axis is ignored
func (v1 Vector4) Dot3d(v2 Vector4) float64 {
	return v1[0]*v2[0]+v1[1]*v2[1]+v1[2]*v2[2]
}

// For those cases when the 4d vector represents just a 3d vector. 4:t axis is ignored
func (v1 Vector4) Cross3d(v2 Vector4) Vector4 {
	return Vector4{v1[1]*v2[2]-v1[2]*v2[1],v1[2]*v2[0]-v1[0]*v2[2],v1[0]*v2[1]-v1[1]*v2[0]} 
}
*/

// If two vectors represents points the distance between them can be calculated
// Forth value is ignored
func (v0 Vector4) Distance3d(v1 Vector4) float64 {
	d0 := v0[0]-v1[0]
	d1 := v0[1]-v1[1]
	d2 := v0[2]-v1[2]
	return math.Sqrt(d0*d0+d1*d1+d2*d2)
}

func (v Vector4) Length() float64 {
	return math.Sqrt(v[0]*v[0]+v[1]*v[1]+v[2]*v[2])
}

// Normalize will modify this vector
func (v Vector4) Normalize() Vector4{
	l:=v.Length();v[0]/=l;v[1]/=l;v[2]/=l;
	return v
}

func (m Vector4) Equal(q Vector4) bool {
	return m[0]==q[0]&&m[1]==q[1]&&m[2]==q[2]&&m[3]==q[3]
}

func (m Vector4) NotEqual(q Vector4) bool {
	return m[0]!=q[0]||m[1]!=q[1]||m[2]!=q[2]||m[3]!=q[3]
}

func (m Vector4) Copy() Vector4 {
	return Vector4{m[0],m[1],m[2],m[3]}[:]
}

func (v Vector4) String() string {
	return fmt.Sprintf("[%.5f,%.5f,%.5f,%.5f]",v[0],v[1],v[2],v[3])
}

func (v Vector4) MultiplyM(m Matrix4) Vector4 {
    return Vector4{	
    	v[ 0]*m[ 0]+v[ 1]*m[ 1]+v[ 2]*m[ 2]+v[ 3]*m[ 3],
    	v[ 0]*m[ 4]+v[ 1]*m[ 5]+v[ 2]*m[ 6]+v[ 3]*m[ 7],
    	v[ 0]*m[ 8]+v[ 1]*m[ 9]+v[ 2]*m[10]+v[ 3]*m[11],
    	v[ 0]*m[12]+v[ 1]*m[13]+v[ 2]*m[14]+v[ 3]*m[15]}[:]
}

/*
Tests to see if the difference between two matrices,
element-wise, exceeds ε.
*/
func (a Vector4) ApproxEquals(b Vector4, ε float64) bool {
	for i := 0; i < 4; i++ {
		if math.Fabs(a[i]-b[i]) > ε {
			return false
		}
	}
	return true
}
