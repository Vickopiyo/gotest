package testing  
    
// CODE COVERAGE--is the way of determining  how much of code is actually covered through the tests.
// go test -cover  - checks amount of code coverage in testfile.in this case only sum() was tested and getmax()-untested thus 25% of statement tested       

// ALL GO TESTS RUN AT THE FOLDER(PACKAGE)  BEING TESTED 

//  go test - testing a package but reeurns in PASS or FAIL aspect only 

// go test -cover  testing package but outputs in percentage--  code coverage  

// go test -coverprofile=coverage.out  - works same as go test -cover only that it creates a file coverage.out     

// go tool cover -func=coverage.out -  shows each func that has been tested with percentage test.sum()--100%..GetMax()--0(before writing GO tests  for GetMax(())   

// 	Total is 33% of sum.go tested     

// go tool cover -html=coverage.out -   run after running go tool cover -func=coverage.out to open web browser to display covered and untracked funcs        
       
// go test -cpuprofile=cou.out -  used checkout cpu activty while executing programs--ususally inidicates time of excution     

// go tool pprof cou.out - prints info about cpu  execution of program in terms of file , time , duration..it also enters interactive mode..Once it enters interactive mode we can use command like (top-breaks down the top nodes --func which consumes the largest time ..)  ..another one is (list--  shows the functons stil)   ,,there is (web)  too  (quit--leaving the interactive mode )  
        
// After installing graphviz , visualization tool ...run go tool pprof cou.out-  to enter interactive mode  .. enter top and then enter list Fibonnaci,  then enter web to redirect you to a webpage..enter pdf in the interactive mode dowload a pdf to project about  cpu activity and func executed   

   
func Sum( a, b int) int {     
	return a + b 
	      
}     

func GetMax(a, b int)  int{  
	  
  if a > b {
	return a 
  }    

  return b 
}                         
 
//PROFILING - analyzing the complexity and costs of a go program such as its memory usage and frequently called funcs to identify the expensive sections of a go program   

func Fibonacci(n int)  int{       
	
	if n <= 1{
	  return n 
	}
   
	return Fibonacci(n -1 ) + Fibonacci(n -2 )
}     