package utils

import (
	"testing"
)

func TestTempToken(t *testing.T) {
	Tests:=[]([]string){
		[]string{"a","a@a.com"},
		[]string{"b","b@b.com"},
		[]string{"c","c@c.com"},
		[]string{"d","d@d.com"},
	}
	for i:=range Tests{
		tempToken,err:=SerialiseTempToken(Tests[i][0],Tests[i][1])
		if err!=nil{
			t.Errorf(err.Error())
		}
		username,gmail,err:=DeserialiseTempToken(tempToken)
		if err!=nil{
			t.Errorf(err.Error())
		}
		if(username!=Tests[i][0] || gmail!=Tests[i][1]){
			t.Errorf("error on TC#%d",i)
		}

	}
}
