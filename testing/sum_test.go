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

         total := Sum(item.a, item.b)       

		 if total != item.c{
         
			t.Errorf("Test failed, have %d , was expecting %d ",  item.c, total)            
		 }
	 }      
     
}        
   

func TestGetMax(t *testing.T)  {
	  
	table :=[]struct{

        a int 
		b int   
		c int
        
	}{
         
    {4, 2, 4},
	{6, 3, 6},  
	// Adding a  struct here with last element not matching test makes the other part of getMax()--tested return b  
	// here 3(b) is greater than  a thus returns b unlike previously where a has always been greater.
	// After this code average  will 100% 
    {2, 3, 3},
		
	}
  
    for _, item := range table {  
           
         max := GetMax(item.a, item.b)   

    if max != item.c{
		t.Errorf("Get max incorrect, have %d, was expecting %d ", max, item.c)
	}
	       
	}
}    
   
// FIbonacci Test          

func TestFibo(t *testing.T)  {     

    table :=[]struct{
		n int  
		r int 
	}{
		{1,1} ,  
		{8,21},
		{50,12586269025},      
	}       
   
	for _, item := range table {   

		    fibo :=Fibonacci(item.n)   

			if fibo !=item.r{
				t.Errorf("Fibo n  %d , was expecting %d ", fibo, item.r)
			}	
	}
     
}
          









