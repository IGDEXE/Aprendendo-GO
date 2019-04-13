package main
import "fmt"
func main(){
    //fmt.Print() == Imprime na tela, como a função printf()
    //fmt.Scanf() == Recebe um dado e armazena na variavel, como a função scanf()
    //If variavel operador logico teste { codigo }
    fmt.Print("Par ou impar\n")
    var teste int
    fmt.Print("Que numero deseja testar???\n")
    fmt.Scanf("%d",&teste)
    if teste % 2 == 0{
        fmt.Print("O numeor eh par !!")  
    }
    if teste % 2 != 0{
        fmt.Print("O numero eh impar !!!")   
    }
}
