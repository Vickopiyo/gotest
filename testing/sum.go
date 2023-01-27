package testing  
    
// CODE COVERAGE--is the way of determining  how much of code is actually covered through the tests.
// go test -cover  - checks amount of code coverage in testfile.in this case only sum() was tested and getmax()-untested thus 25% of statement tested       

// ALL GO TESTS RUN AT THE FOLDER(PACKAGE)  BEING TESTED 

//  go test - testing a package 

// go test -cover   

// go test -coverprofile=coverage.out  - works same as go test -cover only that it creates a file coverage.out     

// go tool cover -func=coverage.out -  shows each func that has been tested with percentage test.sum()--100%..GetMax()--0%..
// 	Total is 33% of sum.go tested 

func sum(  a, b int) int {     
	return a + b 
	      
}     
   

func GetMax(a, b int)  int{  
	  
  if a > b {
	return a 
  }    

  return b 
}