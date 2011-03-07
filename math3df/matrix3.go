/*
This code is an incomplete port of the C++ algebra library WildMagic5 (geometrictools.com)
Note that this code uses column major matrixes, just like OpenGl
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
http://www.geometrictools.com/License/Boost/LICENSE_1_0.txt
*/

package math3df

import "fmt"

type Matrix3 []float32

func MakeMatrix3V(v[]float32,rowMajor bool) Matrix3{
	if rowMajor {
	    // transform the data to OpenGl format
		return Matrix3{v[0],v[3],v[6],v[1],v[4],v[7],v[2],v[5],v[8]}[:]
	}
	return Matrix3{v[0],v[1],v[2],v[3],v[4],v[5],v[6],v[7],v[8]}[:]
}

func MakeMatrix3() Matrix3{
	return Matrix3{.0,.0,.0,.0,.0,.0,.0,.0,.0}[:]
}

func (m Matrix3) Copy() Matrix3 {
	return Matrix3{m[0],m[1],m[2],m[3],m[4],m[5],m[6],m[7],m[8]}[:]
}

func (m Matrix3) MakeZero() Matrix3 {
	m[ 0]=.0;m[ 1]=.0;m[ 2]=.0;m[ 3]=.0;m[ 4]=.0;m[ 5]=.0;m[ 6]=.0;m[ 7]=.0;m[ 8]=.0
	return m
}

func (m Matrix3) MakeIdentity() Matrix3 {
	m[ 0]=1.;m[ 4]=1.;m[8]=1.
	m[ 1]=.0;m[ 2]=.0;m[ 3]=.0;m[ 5]=.0;m[ 6]=.0;m[ 7]=.0
	return m
}

func (m Matrix3) Determinant() float32{
	return m[0]*(m[4]*m[8]-m[5]*m[7])-m[1]*(m[3]*m[8]-m[5]*m[6])+m[2]*(m[3]*m[7]-m[4]*m[6]);
}

func (m Matrix3) MulS(scalar float32) Matrix3 {
	s := scalar
	return Matrix3{m[ 0]*s,m[ 1]*s,m[ 2]*s,m[ 3]*s,m[ 4]*s,m[ 5]*s,m[ 6]*s,m[ 7]*s,m[ 8]*s}[:]
}

func (m Matrix3) Inverse() Matrix3 {
	r := MakeMatrix3()
	d := 1.0/m.Determinant();
	r[0]= d*(m[4]*m[8]-m[5]*m[7]);
	r[1]=-d*(m[1]*m[8]-m[2]*m[7]);
	r[2]= d*(m[1]*m[5]-m[2]*m[4]);
	r[3]=-d*(m[3]*m[8]-m[5]*m[6]);
	r[4]= d*(m[0]*m[8]-m[2]*m[6]);
	r[5]=-d*(m[0]*m[5]-m[2]*m[3]);
	r[6]= d*(m[3]*m[7]-m[4]*m[6]);
	r[7]=-d*(m[0]*m[7]-m[1]*m[6]);
	r[8]= d*(m[0]*m[4]-m[1]*m[3]);
	return r;
}

func (m Matrix3) Cofactor() Matrix3 {
	r := MakeMatrix3()
	r[0]= (m[4]*m[8]-m[5]*m[7]);r[1]=-(m[3]*m[8]-m[5]*m[6]);r[2]= (m[3]*m[7]-m[4]*m[6]);
	r[3]=-(m[1]*m[8]-m[2]*m[7]);r[4]= (m[0]*m[8]-m[2]*m[6]);r[5]=-(m[0]*m[7]-m[1]*m[6]);
	r[6]= (m[1]*m[5]-m[2]*m[4]);r[7]=-(m[0]*m[5]-m[2]*m[3]);r[8]= (m[0]*m[4]-m[1]*m[3]);
	return r;
}

func (m Matrix3) Equal(q Matrix3) bool {
	return m[0]==q[0]&&m[3]==q[3]&&m[6]==q[6]&&m[1]==q[1]&&m[4]==q[4]&&m[7]==q[7]&&m[2]==q[2]&&m[5]==q[5]&&m[8]==q[8]
}

func (m Matrix3) NotEqual(q Matrix3) bool {
	return m[0]!=q[0]||m[3]!=q[3]||m[6]!=q[6]||m[1]!=q[1]||m[4]!=q[4]||m[7]!=q[7]||m[2]!=q[2]||m[5]!=q[5]||m[8]!=q[8]
}

// Mutiply this matrix with a column vector v, resulting in another column vector
func (m Matrix3) MultiplyV(v Vector3) Vector3{
	return Vector3{ m[0]*v[0]+m[1]*v[1]+m[2]*v[2],
					m[3]*v[0]+m[4]*v[1]+m[5]*v[2],
					m[6]*v[0]+m[7]*v[1]+m[8]*v[2] };
}

func (m Matrix3) MultiplyM(q Matrix3)Matrix3{
	r := MakeMatrix3()
	r[0]=q[0]*m[0]+q[1]*m[3]+q[2]*m[6]
	r[1]=q[0]*m[1]+q[1]*m[4]+q[2]*m[7]
	r[2]=q[0]*m[2]+q[1]*m[5]+q[2]*m[8]
	r[3]=q[3]*m[0]+q[4]*m[3]+q[5]*m[6]
	r[4]=q[3]*m[1]+q[4]*m[4]+q[5]*m[7]
	r[5]=q[3]*m[2]+q[4]*m[5]+q[5]*m[8]
	r[6]=q[6]*m[0]+q[7]*m[3]+q[8]*m[6]
	r[7]=q[6]*m[1]+q[7]*m[4]+q[8]*m[7]
	r[8]=q[6]*m[2]+q[7]*m[5]+q[8]*m[8]
	return r;
}

// Transposed will *not* modify m
func (m Matrix3) Transposed() Matrix3 {
	return Matrix3{m[0],m[3],m[6],m[1],m[4],m[7],m[2],m[5],m[8]}[:]
}

// Transpose will modify m
func (m Matrix3) Transpose() Matrix3 {
	m[1],m[3]=m[3],m[1]
	m[2],m[6]=m[6],m[2]
	m[5],m[7]=m[7],m[5]
	return m
}

/*
// Orthogonalize will modify this matrix
func (m Matrix3) Orthogonalize(){
	i := MakeVector3(m[0],m[1],m[2])
	j := MakeVector3(m[3],m[4],m[5]) 
	k := MakeVector3(m[6],m[7],m[8]).Normalize();
	i = j.Cross(k).Normalize()
	j=k.Cross(i);
	m[0]=i[0]; m[3]=j[0]; m[6]=k[0]
	m[1]=i[3]; m[4]=j[3]; m[7]=k[3]
	m[2]=i[6]; m[5]=j[6]; m[8]=k[6]
}

func (m1 Matrix3) Orthogonalized() Matrix3{
	m := m1.Copy()
	m.Orthogonalize();
	return m;
}
*/

/*
Tests to see if the difference between two matrices,
element-wise, exceeds ε.
*/
func (a Matrix3) ApproxEquals(b Matrix3, ε float32) bool {
	for i := 0; i < 9; i++ {
		if Fabsf(a[i]-b[i]) > ε {
			return false
		}
	}
	return true
}

func (m Matrix3)String() string {
	// output in octave format for easy testing
	return fmt.Sprintf("[%.5f,%.5f,%.5f;%.5f,%.5f,%.5f;%.5f,%.5f,%.5f]",m[0],m[3],m[6],m[1],m[4],m[7],m[2],m[5],m[8])
}
