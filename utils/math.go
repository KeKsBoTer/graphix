package utils


func Div(v1,v2 int32) float32{
	return float32(v1)/float32(v2)
}

func Min(v1,v2 int) int{
	if v1<v2{
		return v1
	}else{
		return v2
	}
}
func Max(v1,v2 int) int{
	if v1>v2{
		return v1
	}else{
		return v2
	}
}