package main

import (
	"errors"
	"fmt"
)

func main() {
	/*Declarar um int sem especificação irá utilizar conforme os bits da arquitetura do seu computador*/
	var numero int = 15
	/*Declaração implicita e também irá utilizar conforme os bits da arquitetura do seu computador*/
	numero2 := 18889997777
	/*tipos de inteiros de acordo com os bits, quanto maior os bits maior a capacidade. Ex.:int8, int16, int32,int64*/
	var learnInt int64 = 1000000000000
	/*uint é o int sem sinal (unsygned int), Ex.:int8, int16, int32,int64*/
	var numero3 uint32 = 120
	/*No go quando não é inicializada a variável é sempre inicializado a 0
	Em inteiros retorna 0*/
	var numero4 int16

	//alias
	/*INT32 = RUNE */
	var numero5 rune = 123456
	/*BYTE = UINT8*/
	var numero6 byte = 123

	/*float32, float64 > sempre tem de por o ponto porque virgula nao funciona*/
	var numeroReal1 float32 = 123.45
	var numeroReal2 float64 = 1230000000000000000000.45

	fmt.Println(numero, numero2, learnInt, numero3, numero4, numero5, numero6, numeroReal1, numeroReal2)

	//Fim dos números reais

	/*Não tem char, o caracteres char são convertidos para numero
	String sempre precisa de colocar aspas duplas
	*/
	var str string = "Text"
	str2 := "Text2"
	fmt.Println(str, str2)

	/*	Aspas simples considera char
		char no go nao existe, sempre que declarar um caractere com aspas simples vai guardar a posição do caractere na tabela asci*/
	exChar := 'A'
	fmt.Println(exChar)
	/*No go quando não é inicializada a variável é sempre inicializado a 0
	Em string retorna string em branco*/
	var textZero string
	fmt.Println(textZero)

	var booleano1 bool
	fmt.Println(booleano1)
	/*retorna <nil> que significa 0*/
	var erro error
	fmt.Println(erro)
	/*Para criar um erro é preciso ter o pacote dos erros*/
	var erro2 error = errors.New("Erro Interno")
	fmt.Println(erro2)

}
