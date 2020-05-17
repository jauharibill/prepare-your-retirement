package main

import "fmt"

func main() {

const revenue float64 = 0.073
const expense float64 = 0.0048
var tabungan float64 = 2000000
var profit float64

const year int = 20
var helper float64

for i:=0; i<year; i++ {
  
   for n:=0; n<12; n++ {
    
    helper += tabungan	
    
   }

   profit = (helper * revenue)
   operationalcost := (helper * expense)
   helper += profit
   helper -= operationalcost
 }
 
 fmt.Printf("Total Nilai investasinya : %f \n", helper)
 fmt.Printf("Total Profit : %f", profit)

}
