package testing    
import "testing"


// TEST FUNC always begin with Uppercases  
// Test func must begin with word Test them completed with file being tested 
// Test files always extend with _test        3




// func TestSum(t *testing.T)  {  
// 	// Test condition from sum()---in the sum.go 
// 	  total := sum(5, 5)          

//     //  If test fails            
// 	  if total != 10 {
// 		t.Errorf("Test failed, have %d, was expecting %d", total, 10)         
// 	  }            
	  
// }                    
     
func TestSum(t *testing.T)  {       
// A slice of type struct..the structs elements(a,b,c) of type int.Each struct has 3 elemnts.Expressly initialiaze the values of the structs 
// SLices  only tekes elements of the same  type!!!!
    table := []struct{
     
		a int  
		b int  
		c int
	}{
		{1,2,3}, 
		{2,2,4},
		{25,25,50},
	}  
     // iterating through the table to check the values in each sruct        
	 for _,  item:= range table {

         total := sum(item.a, item.b)     
		 if total != item.c{
         
			t.Errorf("Test failed, have %d , was expecting %d ",  item.c, total)        
		 }
	 }      
     
}        
