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


// This is a 4 element vector float64
type Quaternion []float64

func MakeQuaternion(s,x,y,z float64) Quaternion{
	return Quaternion{s,x,y,z}[:]
}

func MakeQuaternionV(v[]float64) Quaternion{
	return Quaternion{v[0],v[1],v[2],v[3]}[:]
}

func MakeQuaternionM(rotationMatrix Matrix4) Quaternion{
	return rotationMatrix.ToQuaternion()
}

func MakeQFromAxisAngle(axis Vector3, angle float64) Quaternion {
    // assert:  axis[] is unit length
    //
    // The quaternion representing the rotation is
    //   q = cos(A/2)+sin(A/2)*(x*i+y*j+z*k)

    halfAngle := 0.5*angle
    sn := math.Sin(halfAngle)
    return Quaternion{math.Cos(halfAngle),sn*axis[0],sn*axis[1],sn*axis[2]}[:]
}

func (q Quaternion) Copy() (Quaternion){
	return Quaternion{q[0],q[1],q[2],q[3]}[:]
}

func MakeQuaternionCopy(q Quaternion) (Quaternion){
	return Quaternion{q[0],q[1],q[2],q[3]}[:]
}

// Will copy the values of p into q
func (q Quaternion) CopyFrom(p Quaternion) (Quaternion){
	return Quaternion{q[0],q[1],q[2],q[3]}[:]
}

func (q Quaternion) W() (*float64){
	return &q[0]
}

func (q Quaternion) X() (*float64){
	return &q[1]
}

func (q Quaternion) Y() (*float64){
	return &q[2]
}

func (q Quaternion) Z() (*float64){
	return &q[3]
}

func (m Quaternion) Equal(q Quaternion) bool {
	return m[0]==q[0]&&m[1]==q[1]&&m[2]==q[2]&&m[3]==q[3]
}

func (m Quaternion) NotEqual(q Quaternion) bool {
	return m[0]!=q[0]||m[1]!=q[1]||m[2]!=q[2]||m[3]==q[3]
}

func (m Quaternion) AddQ(q Quaternion) Quaternion {
	return Quaternion{m[0]+q[0],m[1]+q[1],m[2]+q[2],m[3]+q[3]}[:]
}

func (m Quaternion) SubtractQ(q Quaternion) Quaternion {
	return Quaternion{m[0]-q[0],m[1]-q[1],m[2]-q[2],m[3]-q[3]}[:]
}

func (m Quaternion) MultiplyQ(q Quaternion) Quaternion {
	return Quaternion{
        m[0]*q[0]-m[1]*q[1]-m[2]*q[2]-m[3]*q[3],
        m[0]*q[1]+m[1]*q[0]+m[2]*q[3]-m[3]*q[2],
        m[0]*q[2]+m[2]*q[0]+m[3]*q[1]-m[1]*q[3],
        m[0]*q[3]+m[3]*q[0]+m[1]*q[2]-m[2]*q[1]}[:]
}

func (m Quaternion) MultiplyS(scalar float64) Quaternion {
	return Quaternion{m[0]*scalar,m[1]*scalar,m[2]*scalar,m[3]*scalar}[:]
}

func (m Quaternion) DivS(scalar float64) Quaternion {
	if scalar !=0 {
		return Quaternion{m[0]/scalar,m[1]/scalar,m[2]/scalar,m[3]/scalar}[:]
	} 
	return Quaternion{math.MaxFloat32,math.MaxFloat32,math.MaxFloat32,math.MaxFloat32}[:]
}

func (m Quaternion) Conjugate() Quaternion {
    return Quaternion{m[0],-m[1],-m[2],-m[3]}[:]
}

func (m Quaternion) Magnitude() float64 {
    return math.Sqrt(m[0]*m[0]+m[1]*m[1]+m[2]*m[2]+m[3]*m[3])
}

func (m Quaternion) ToRotationMatrix() Matrix4 {
 
    twoX  := 2.*m[1]
    twoY  := 2.*m[2]
    twoZ  := 2.*m[3]
    twoWX := twoX*m[0]
    twoWY := twoY*m[0]
    twoWZ := twoZ*m[0]
    twoXX := twoX*m[1]
    twoXY := twoY*m[1]
    twoXZ := twoZ*m[1]
    twoYY := twoY*m[2]
    twoYZ := twoZ*m[2]
    twoZZ := twoZ*m[3]
	return Matrix4{
	    1.-(twoYY + twoZZ),	twoXY+twoWZ,		twoXZ - twoWY,	0.,
    	twoXY-twoWZ,		1.-(twoXX+twoZZ),	twoYZ + twoWX,	0.,
    	twoXZ+twoWY,		twoYZ-twoWX,		1.-(twoXX + twoYY),	0.,
    	0.,	0.,	0.,	1.}[:]
}

func (m Quaternion) FromAxisAngle(axis Vector3, angle float64) Quaternion {
    // assert:  axis[] is unit length
    //
    // The quaternion representing the rotation is
    //   q = cos(A/2)+sin(A/2)*(x*i+y*j+z*k)

    halfAngle := 0.5*angle
    sn := math.Sin(halfAngle)
    
    m[0]=math.Cos(halfAngle)
    m[1]=sn*axis[0]
    m[2]=sn*axis[1]
    m[3]=sn*axis[2]
    return m
}

func (m Quaternion) ToAxisAngle() (axis Vector3, angle float64) {
    // The quaternion representing the rotation is
    //   q = cos(A/2)+sin(A/2)*(x*i+y*j+z*k)

    sqrLength := m[1]*m[1] + m[2]*m[2] + m[3]*m[3]

    if sqrLength > internalεε {
        angle = 2.*math.Acos(m[0])
        //invLength = math.InvSqrt(sqrLength);
        invLength := 1./math.Sqrt(sqrLength);
        axis = Vector3{m[1]*invLength,m[2]*invLength,m[3]*invLength}[:]
    } else {
        // Angle is 0 (mod 2*pi), so any axis will do.
        angle = 0.
        axis = Vector3{1,0,0}[:]
    }
    return axis,angle
}

func (q Quaternion) Length() float64 {
    return math.Sqrt(q[0]*q[0]+q[1]*q[1]+q[2]*q[2]+q[3]*q[3])
}

func (q Quaternion) SquaredLength() float64 {
    return q[0]*q[0]+q[1]*q[1]+q[2]*q[2]+q[3]*q[3]
}

func (q Quaternion) Normalize(ε float64) (Quaternion,bool) {

	length := q[0]*q[0]+q[1]*q[1]+q[2]*q[2]+q[3]*q[3]
    if length-1.0 <= ε {
    	// already normalized, nothing to do
    	return q,true
    }
    length = math.Sqrt(length)
    if (length > ε){
        invLength := 1./length;
        q[0] *= invLength;
        q[1] *= invLength;
        q[2] *= invLength;
        q[3] *= invLength;
    	return q,true;
    }
    q[0] = 0; q[1] = 0; q[2] = 0; q[3] = 0
    return q,false;
}

func (q1 Quaternion) Dot(q2 Quaternion) float64 {
    return q1[0]*q2[0]+q1[1]*q2[1]+q1[2]*q2[2]+q1[3]*q2[3]
}

// Spherical linear interpolation.
// t is the interpolation value from 0. to 1.
// p and q are 'const'. m is *not*
func (m Quaternion) Slerp(t float64,p,q Quaternion) Quaternion {

    cs := p.Dot(q)
    angle := math.Acos(cs)

    if (math.Fabs(angle) >= internalε){
        sn := math.Sin(angle)
        invSn := 1./sn
        tAngle := t*angle
        coeff0 := math.Sin(angle-tAngle)*invSn
        coeff1 := math.Sin(tAngle)*invSn;

        m[0] = float64(coeff0*p[0] + coeff1*q[0])
        m[1] = float64(coeff0*p[1] + coeff1*q[1])
        m[2] = float64(coeff0*p[2] + coeff1*q[2])
        m[3] = float64(coeff0*p[3] + coeff1*q[3])
    } else {
        m[0] = p[0]
        m[1] = p[1]
        m[2] = p[2]
        m[3] = p[3]
    }

    return m
}

// ------------------------------------
// linearly interpolate each component, then normalize the Quaternion
// Unlike spherical interpolation, this does not rotate at a constant velocity,
// although that's not necessarily a bad thing
// t is the interpolation value from 0. to 1.
// ------------------------------------
func (m Quaternion) NLerp( t float64, a, b Quaternion ) Quaternion {
	w1 := 1.0 - t
	
	// m = a*w1 + b*t
	m[0] = a[0]*w1 + b[0]*t
	m[1] = a[1]*w1 + b[1]*t
	m[2] = a[2]*w1 + b[2]*t
	m[3] = a[3]*w1 + b[3]*t
	
	_,_ = m.Normalize(internalε)
	return m
}


func (m Quaternion) String() string {
	return fmt.Sprintf("[%.5f,%.5f,%.5f,%.5f]",m[0],m[1],m[2],m[3])
}

/*
 Tests to see if the difference between two quaternions, element-wise, exceeds ε.
*/
func (a Quaternion) ApproxEquals(b Quaternion, ε float64) bool {
	for i := 0; i < 4; i++ {
		delta := math.Fabs(a[i]-b[i])
		if delta > ε {
			//fmt.Printf("delta between %f and %f is %f. ε=%f\n",a[i],b[i],delta,ε)
			return false
		}
	}
	return true
}