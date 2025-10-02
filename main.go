package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// numbers é a lista global de números gerenciada pela aplicação.
var numbers = []int{15, 80, 46, 35, 71, 13, 22, 98}

// addNum solicita um número ao usuário e o adiciona à lista.
// Retorna um erro se a entrada for inválida ou o número for negativo.
func addNum(reader *bufio.Reader) error {
	fmt.Print("Digite um número: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		return errors.New("Entrada inválida. Por favor, digite um número inteiro.")
	}
	if num < 0 {
		return fmt.Errorf("O número %d é negativo. Não foi adicionado.", num)
	}
	numbers = append(numbers, num)
	fmt.Println("Adicionado:", numbers)
	return nil
}

// listNum exibe todos os números na lista.
func listNum() {
	fmt.Printf("Lista: %v\n", numbers)
}

// removeByInd solicita um índice e remove o número correspondente da lista.
// Retorna um erro se a entrada for inválida ou o índice estiver fora do alcance.
func removeByInd(reader *bufio.Reader) error {
	fmt.Print("Digite um índice: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	index, err := strconv.Atoi(input)
	if err != nil {
		return errors.New("Entrada inválida. Por favor, digite um número inteiro.")
	}
	if index < 0 || index >= len(numbers) {
		return errors.New("O índice está fora do alcance.")
	}
	numbers = append(numbers[:index], numbers[index+1:]...)
	fmt.Println("Removido:", numbers)
	return nil
}

// statistics calcula a média, mínimo e máximo da lista.
// Retorna um erro se a lista estiver vazia.
func statistics() (float64, int, int, error) {
	if len(numbers) == 0 {
		return 0, 0, 0, errors.New("não há números para calcular estatísticas")
	}
	sum := 0
	minNum := numbers[0]
	maxNum := numbers[0]
	for _, num := range numbers {
		sum += num
		if num < minNum {
			minNum = num
		}
		if num > maxNum {
			maxNum = num
		}
	}
	average := float64(sum) / float64(len(numbers))
	return average, minNum, maxNum, nil
}

// safeDivision solicita dois números e realiza uma divisão segura.
// Retorna um erro se a entrada for inválida ou o divisor for zero.
func safeDivision(reader *bufio.Reader) error {
	fmt.Print("Digite o dividendo: ")
	inputDividend, _ := reader.ReadString('\n')
	inputDividend = strings.TrimSpace(inputDividend)
	dividend, err := strconv.ParseFloat(inputDividend, 64)
	if err != nil {
		return errors.New("Entrada inválida para o dividendo.")
	}

	fmt.Print("Digite o divisor: ")
	inputDivisor, _ := reader.ReadString('\n')
	inputDivisor = strings.TrimSpace(inputDivisor)
	divisor, err := strconv.ParseFloat(inputDivisor, 64)
	if err != nil {
		return errors.New("Entrada inválida para o divisor.")
	}

	if divisor == 0 {
		return errors.New("Divisão por zero não é permitida.")
	}
	fmt.Printf("Resultado: %.2f\n", dividend/divisor)
	return nil
}

// clearList esvazia a lista de números.
func clearList() {
	numbers = []int{}
	fmt.Println("Lista limpa.")
}

// sortList ordena a lista em ordem crescente e decrescente.
func sortList() error {
	if len(numbers) == 0 {
		return errors.New("Nenhum número na lista para ordenar")
	}
	sort.Ints(numbers)
	fmt.Printf("Ordem crescente: %v\n", numbers)
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Printf("Ordem decrescente: %v\n", numbers)
	return nil
}

// evenNum exibe apenas os números pares da lista.
func evenNum() error {
	if len(numbers) == 0 {
		return errors.New("A lista está vazia, não há números pares")
	}
	var evenNumbers []int
	for _, num := range numbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	if len(evenNumbers) == 0 {
		return errors.New("A lista não contém números pares")
	}
	fmt.Printf("Números pares: %v\n", evenNumbers)
	return nil
}

// exportToFile exporta a lista para um arquivo de texto.
func exportToFile() error {
	if len(numbers) == 0 {
		return errors.New("A lista está vazia, não há o que exportar")
	}

	file, err := os.Create("numbers.txt")
	if err != nil {
		return fmt.Errorf("Erro ao criar o arquivo: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, num := range numbers {
		_, err := writer.WriteString(fmt.Sprintf("%d\n", num))
		if err != nil {
			return fmt.Errorf("Erro ao escrever no arquivo: %w", err)
		}
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("Erro ao limpar o buffer do arquivo: %w", err)
	}
	return nil
}

// main é a função principal que gerencia o menu e a interação do usuário.
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
			if err := addNum(reader); err != nil {
				fmt.Println("Erro:", err)
			}
		case "2":
			listNum()
		case "3":
			if err := removeByInd(reader); err != nil {
				fmt.Println("Erro:", err)
			}
		case "4":
			average, min, max, err := statistics()
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Printf("Estatísticas: Média: %.2f, Mínimo: %d, Máximo: %d\n", average, min, max)
			}
		case "5":
			if err := safeDivision(reader); err != nil {
				fmt.Println("Erro:", err)
			}
		case "6":
			clearList()
		case "7":
			if err := sortList(); err != nil {
				fmt.Println("Erro:", err)
			}
		case "8":
			if err := evenNum(); err != nil {
				fmt.Println("Erro:", err)
			}
		case "9":
			if err := exportToFile(); err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println("Lista exportada para 'numbers.txt' com sucesso.")
			}
		case "0":
			fmt.Println("Saindo do programa. Obrigado!")
			return
		default:
			fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
		}
	}
}
