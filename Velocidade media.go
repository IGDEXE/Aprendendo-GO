package main
import "fmt"
func main(){
    var d,ti,tf,v int
     fmt.Print("Velocidade media")
     fmt.Print("Qual eh a distancia??\n")
     fmt.Scanf("%d",&d)
     fmt.Print("Qual eh o tempo inicial??\n")
     fmt.Scanf("%d",&ti)
     fmt.Print("Qual eh o tempo final??\n")
     fmt.Scanf("%d",&tf)
     v = d/(tf - ti)
     fmt.Print("Velocidade media = ", v )
}
