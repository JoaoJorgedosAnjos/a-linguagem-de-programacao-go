# Exercício 1.3

Experimente medir a diferença de tempo de execução entre nossas versões potencialmente ineficentes e a versão queusa string.Join.(A seção 1.6 mostra parte do pacote time, e a seção 11.4 mostra como escrever testes comparativos para uma avaliação sistemática de desempenho)

// eficiente.go
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    // strings.Join faz tudo em uma única operação otimizada
    fmt.Println(strings.Join(os.Args[1:], " "))
}

x

// ineficiente.go
package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", "" // s começa vazia, sep também
    for _, arg := range os.Args[1:] {
        s += sep + arg // Concatena a string a cada passo
        sep = " "      // A partir do segundo argumento, o separador será um espaço
    }
    fmt.Println(s)
}

## Exercício 1.3: Medindo a Eficiência de Concatenação de Strings

Este exercício tem como objetivo demonstrar experimentalmente a diferença de desempenho entre duas maneiras de concatenar (juntar) uma grande quantidade de strings em Go.

1.  **Método Ineficiente:** Usando o operador `+=` dentro de um loop.
2.  **Método Eficiente:** Usando a função otimizada `strings.Join`.

---

### O Experimento e os Resultados

Para medir o tempo de execução de forma visível, utilizamos o comando `time` do terminal, que mede o tempo total de execução de um processo. Para gerar uma carga de trabalho massiva, passamos 50.000 argumentos para cada programa.

#### Entrada (Comando de Teste)

#### 1. Testando a versão ineficiente:**
```bash
time go run ineficiente.go $(seq 1 50000)
```

#### 2. Testando a versão eficiente:
Bash
```
time go run eficiente.go $(seq 1 50000)
```

Saída (Resultados Observados)

Os resultados mostram uma diferença de performance drástica.

1. Saída da versão ineficiente:
O programa demorou um tempo considerável para ser concluído. O tempo de usuário (user) foi de aproximadamente 2.16 segundos.

### Exemplo de saída do comando 'time' para a versão ineficiente
### (os valores exatos podem variar)
```
real    0m2.253s
user    0m2.165s
sys     0m0.082s
```
2. Saída da versão eficiente:
O programa foi executado quase que instantaneamente. O tempo de usuário (user) foi próximo de zero.

### Exemplo de saída do comando 'time' para a versão eficiente
```
real    0m0.061s
user    0m0.034s
sys     0m0.027s
```
Conclusão: O método com strings.Join é ordens de magnitude mais rápido. Isso ocorre porque ele aloca a memória necessária uma única vez, enquanto o método com += realiza milhares de alocações e cópias de memória, um processo extremamente custoso.

Explicação do Comando de Teste: $(seq 1 50000)

Esta é uma técnica poderosa do shell (o terminal, como o Bash), e não do Go. Ela nos permite gerar uma grande quantidade de argumentos de forma automática. Vamos dividi-la em duas partes:

#### 1. O Comando seq

seq é a abreviação de "sequence" (sequência). Sua função é gerar uma sequência de números. O comando seq 1 50000 simplesmente imprime todos os números de 1 a 50000, cada um em uma nova linha.

#### 2. A Sintaxe $(...)

Isso é chamado de Substituição de Comando. Ela instrui o terminal a fazer o seguinte:

    Executar o comando que está dentro dos parênteses (seq 1 50000).

    Capturar toda a saída de texto gerada por esse comando.

    Substituir o trecho $(...) por essa saída de texto, tratando cada palavra ou número separado por espaço/quebra de linha como um argumento individual.

Como Tudo Funciona Junto

Quando você executa go run eficiente.go $(seq 1 50000), o terminal primeiro resolve o $(seq 1 50000), que se transforma em uma lista de 50.000 números. Em seguida, ele executa o comando principal, passando essa lista gigante como argumentos. Para o programa Go, é como se ele tivesse sido chamado assim:
```
Bash
go run eficiente.go 1 2 3 4 5 ... 49999 50000
```
Isso nos permite testar o desempenho do nosso código sob uma carga de trabalho pesada de forma prática e rápida.
