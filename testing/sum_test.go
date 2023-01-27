package testing    
import "testing"


func TestSum(t *testing.T)  {  


	// Test condition from sum()---in the sum.go 
	  total := sum(5, 5)          

    //  If test fails 
	  if total != 10 {
		t.Errorf("Test failed, have %d, was expecting %d", total, 10)
    
	  }     


}                    

