/* first example*/
package main
import ( "fmt )
func main(){
var a [5]int
a[2] = 7
fmt.println(a)
}

/* second example*/
package main
import ( "fmt )
func main(){
 a:=[5]int{5,4,3,2,1}
fmt.println(a)
}

/*third example(scelices)*/

package main
import ( "fmt )
func main(){
 a:=[5]int{5,4,3,2,1}
 a = append(a,13)
fmt.println(a)
}

/*fourth* typedefination()*/

package main
import{
"fmt"
}
func main(){
vertices := make(map[string]int)
 vertices["triangle"] =2
 vertices["square"]=3
 vertices["dodecagon"]=12
 delete(vertices,"square")
 fmt.println(vertices)
}

/*fifth is loops*/
package main
import( "fmt")
func main(){
 for i := 0; i<5; i++{
 fmt.println(i)
 }
 }

/*sixth while loop*/
package main
import{"fmt")
func main()
{
 i:=0
 while(i<5)
  fmt.println(i)
   i++
   }
   }
   
   /*seventh using scelising  by range*/
   package main
   import ("fmt")
   func main()
   {
   arr:=[]string{"a","b","c"}
    for index,value := range arr{
    fmt.println("index",index,"value",value)
    }
    }

/*eigth we can do the same thing with make also instead of index=range is replaced*/

 package main
   import ("fmt")
   func main()
   {

m := make(map[string]string)
m["a"] ="alpha"
["b"] = "beta"
    for key,value := range m{
    fmt.println("key",key,"value",value)
    }
    }


/*ninth*/
package main
import("fmt")
func main(){
}
func sum (x int,y int)int {
return x+y
}

/* tenth(go lang can return n number of return types )*/
package main
import("fmt"
)
func main(){
result :=sum(2,3)
fmt.println(result)
}
func sum (x int,y int)int {
return x+y
}

/*eleventh*/
package main
import("fmt" "error"
)
func main(){
 result , err:=squrt(-16)
if err!= nil{
fmt.println(err)
}
else{
fmt.println(result)
}
}
func squrt(x float64)(float64,error){
if x<0{
return 0,errors.new("undefined for negative numbers")
}
return math.sqrt(x),nil
}

/*12th structure*/
package main
import("fmt" "error"
)
type person struct {
name string
age int
}
func main()
{
p := person{name:"jake".age:34}
fmt.println(p.name)
}
