package main

import (
	"bufio" // operações de leitura e escrita mais eficientes
	"fmt"
	"os"
	"sort"
	"strconv" // conversões de e para representações de string de tipos de dados básicos
	"strings"
)

var numbers = []int{15, 80, 46, 35, 71, 13, 22, 98}

// addNum prompts for an integer and adds it to the slice.
func addNum() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite um número: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Erro: Entrada inválida. Por favor, digite um número inteiro.")
		return
	}
	if num < 0 {
		fmt.Printf("O número %d é negativo. Não foi adicionado.\n", num)
		return
	}
	numbers = append(numbers, num)
	fmt.Println("Adicionado:", numbers)
}

// listNum displays all stored numbers.
func listNum() {
	fmt.Printf("Lista: %v\n", numbers)
}

// removeByInd removes a number from the slice based on a given index.
func removeByInd() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite um index: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	index, err := strconv.Atoi(input) // de string para int
	if err != nil {
		fmt.Println("Erro: Entrada inválida. Por favor, digite um número inteiro.")
		return
	}
	if index < 0 || index >= len(numbers) {
		fmt.Println("Erro: O índice está fora do alcance.")
		return
	}
	numbers = append(numbers[:index], numbers[index+1:]...)
	fmt.Println("Removido:", numbers)
}

// statistics calculates the minimum, maximum, and average of the numbers.
func statistics() {
	if len(numbers) == 0 {
		fmt.Println("Erro: Nenhum número na lista para calcular estatísticas.")
		return
	}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	average := float64(sum) / float64(len(numbers))

	minNum := numbers[0]
	maxNum := numbers[0]
	for _, num := range numbers {
		if num < minNum {
			minNum = num
		}
		if num > maxNum {
			maxNum = num
		}
	}
	fmt.Printf("Estatísticas: Média: %.2f, Mínimo: %d, Máximo: %d\n", average, minNum, maxNum)
}

// safeDivision performs division between two numbers with error handling.
func safeDivision() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o dividendo: ")
	inputDividend, _ := reader.ReadString('\n')
	inputDividend = strings.TrimSpace(inputDividend)
	dividend, err := strconv.ParseFloat(inputDividend, 64)
	if err != nil {
		fmt.Println("Erro: Entrada inválida para o dividendo.")
		return
	}

	fmt.Print("Digite o divisor: ")
	inputDivisor, _ := reader.ReadString('\n')
	inputDivisor = strings.TrimSpace(inputDivisor)
	divisor, err := strconv.ParseFloat(inputDivisor, 64)
	if err != nil {
		fmt.Println("Erro: Entrada inválida para o divisor.")
		return
	}

	if divisor == 0 {
		fmt.Println("Erro: Divisão por zero não é permitida.")
		return
	}

	fmt.Printf("Resultado: %.2f\n", dividend/divisor)
}

// clearList empties the number slice.
func clearList() {
	numbers = []int{}
	fmt.Printf("Lista limpa: %v\n", numbers)
}

// sortList implements ascending and descending list sorting.
func sortList() {
	if len(numbers) == 0 {
		fmt.Println("Erro: Nenhum número para ordenar.")
		return
	}
	sort.Ints(numbers)
	fmt.Printf("Ordem crescente:  %v\n", numbers)
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Printf("Ordem decrescente:  %v\n", numbers)
}

// evenNum displays only even numbers.
func evenNum() {
	var evenNumbers []int
	if len(numbers) == 0 {
		fmt.Println("A lista não contém números pares.")
		return
	}
	for _, num := range numbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	if len(evenNumbers) == 0 {
		fmt.Println("A lista não contém números pares.")
		return
	}
	fmt.Printf("Números pares:  %v\n", evenNumbers)
}

// exportToFile exports the list to a text file.
func exportToFile() {
	file, err := os.Create("numbers.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, num := range numbers {
		_, err := writer.WriteString(fmt.Sprintf("%d\n", num))
		if err != nil {
			fmt.Println("Erro ao escrever no arquivo:", err)
			return
		}
	}
	writer.Flush()
	fmt.Println("Lista exportada para 'numbers.txt' com sucesso.")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Adicionar número")
		fmt.Println("2. Listar números")
		fmt.Println("3. Remover por índice")
		fmt.Println("4. Estatísticas")
		fmt.Println("5. Divisão segura")
		fmt.Println("6. Limpar lista")
		fmt.Println("7. Ordenar lista")
		fmt.Println("8. Exibir apenas números pares")
		fmt.Println("9. Exportar para arquivo de texto")
		fmt.Println("0. Sair")
		fmt.Print("Digite sua escolha: ")

		escolha, _ := reader.ReadString('\n')
		escolha = strings.TrimSpace(escolha)

		switch escolha {
		case "1":
			addNum()
		case "2":
			listNum()
		case "3":
			removeByInd()
		case "4":
			statistics()
		case "5":
			safeDivision()
		case "6":
			clearList()
		case "7":
			sortList()
		case "8":
			evenNum()
		case "9":
			exportToFile()
		case "0":
			fmt.Println("Saindo do programa. Obrigado!")
			return
		default:
			fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
		}
	}
}
