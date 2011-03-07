package main
/*
This is a quick and dirty hack to copy the code from the float32 package (math3df) into the float64 package (math3dd)
*/
import "os"
import "bufio"
import "fmt"
import "strings"

func findandreplace(line string) string {
	line = strings.Replace(line,"import \"math\"","",-1) // import "math" is added in the next replace
	line = strings.Replace(line,"package math3df","package math3dd\n//This code is auto generated from the math3df package. Do not edit.\nimport \"math\"",-1)
	line = strings.Replace(line,"Vector4f","Vector4d",-1)
	line = strings.Replace(line,"math3d/math3df","math3d/math3dd",-1)
	line = strings.Replace(line,"Vector3f","Vector3d",-1)
	line = strings.Replace(line,"Matrix3f","Matrix3d",-1)
	line = strings.Replace(line,"Cosf(","math.Cos(",-1)
	line = strings.Replace(line,"func math.Cos","func Cos",-1)
	line = strings.Replace(line,"Acosf(","math.Acos(",-1)
	line = strings.Replace(line,"func math.Acos","func Acos",-1)
	line = strings.Replace(line,"Sinf(","math.Sin(",-1)
	line = strings.Replace(line,"func math.Sin","func Sin",-1)
	line = strings.Replace(line,"Fabsf(","math.Fabs(",-1)
	line = strings.Replace(line,"func math.Fabs","func Fabs",-1)
	line = strings.Replace(line,"Sqrtf(","math.Sqrt(",-1)
	line = strings.Replace(line,"func math.Sqrt","func Sqrt",-1)
	line = strings.Replace(line,"float32","float64",-1)
	line = strings.Replace(line,"float64(math.Pi)","math.Pi",-1)
	// With float32 we can up the precision a bit 
	line = strings.Replace(line,"const ε3d = 0.0001","const ε3d = 0.000001",-1)
	//line = strings.Replace(line,"const ε4d = 0.00001","const ε4d = 0.00001",-1)
	
	return line
}

func convertFile(srcFilename, destFilename string) (err os.Error) {

	sourceFile,e := os.Open(srcFilename,os.O_RDONLY,0)
	if e != nil {
		err = e
		return
	}
	defer sourceFile.Close()
	
	destFile,e := os.Open(destFilename,os.O_WRONLY|os.O_CREAT,0)
	if e != nil {
		err = e
		return
	}
	defer destFile.Close()
	r := bufio.NewReader(sourceFile)
	
	var line string
	
	for {
		line,e = r.ReadString(13)
		if len(line) == 0 {
		    break
		}
		line = findandreplace(line)
		destFile.WriteString(line)
	}
	fmt.Println("Converted and saved ", srcFilename , " to ", destFilename)
	return 
}

func main() {
	files := []string{"vector3.go","vector4.go","matrix3.go","matrix4.go","quaternion.go",
					   "math3d_test.go","math4d_test.go","util.go","Makefile"}
	for i:=0;i<len(files);i++{
		e := convertFile(fmt.Sprintf("math3df/%s",files[i]),fmt.Sprintf("math3dd/%s",files[i]))
		if e != nil {
			fmt.Println(e)
			break
		}
	}
}
