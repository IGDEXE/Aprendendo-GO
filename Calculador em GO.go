package main
import "fmt"
func main(){
    fmt.Print("Bem vindo!!\n")
    var op,a,b,c int
    fmt.Print("1-Soma\n")
    fmt.Print("2-Subtracao\n")
    fmt.Print("3-Divisao\n")
    fmt.Print("4-Multiplicacao\n")
    fmt.Scanf("%d",&op)
    if op == 1{
        fmt.Print("Digite o primeiro numero\n")
        fmt.Scanf("%d",&a)
        fmt.Print("Digite o segundo numero\n")
        fmt.Scanf("%d",&b)
        c = a+b
        fmt.Print("O resultado eh : ",c)
    }
    if op == 2{
        fmt.Print("Digite o primeiro numero\n")
        fmt.Scanf("%d",&a)
        fmt.Print("Digite o segundo numero\n")
        fmt.Scanf("%d",&b)
        c = a-b
        fmt.Print("O resultado eh : ",c)
    }
    if op == 3{
        fmt.Print("Digite o primeiro numero\n")
        fmt.Scanf("%d",&a)
        fmt.Print("Digite o segundo numero\n")
        fmt.Scanf("%d",&b)
        c = a/b
        fmt.Print("O resultado eh : ",c)
    }
    if op == 4{
        fmt.Print("Digite o primeiro numero\n")
        fmt.Scanf("%d",&a)
        fmt.Print("Digite o segundo numero\n")
        fmt.Scanf("%d",&b)
        c = a*b
        fmt.Print("O resultado eh : ",c)
    }
    if (op < 1 && op > 4){
        fmt.Print("Opcao invalida")
    }
    
    
}
