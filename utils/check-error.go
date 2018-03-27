package utils

// Check errors in order handle them
func Check(e error){
	if e != nil{
		Log.Println(e)
		panic(e)
	}
}