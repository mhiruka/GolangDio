package main

import (
	"fmt"
	"time"
)

func ping(pingCh chan struct{}, pongCh chan struct{}) {
	for {
		<-pingCh               // espera sua vez
		fmt.Println("ping")
		time.Sleep(500 * time.Millisecond)
		pongCh <- struct{}{}   // libera o pong
	}
}

func pong(pongCh chan struct{}, pingCh chan struct{}) {
	for {
		<-pongCh              // espera sua vez
		fmt.Println("pong")
		time.Sleep(500 * time.Millisecond)
		pingCh <- struct{}{}  // libera o ping
	}
}

func main() {
	pingCh := make(chan struct{})
	pongCh := make(chan struct{})

	// inicia goroutines
	go ping(pingCh, pongCh)
	go pong(pongCh, pingCh)

	// começa pelo ping
	pingCh <- struct{}{}

	// mantém o programa rodando
	select {}
}