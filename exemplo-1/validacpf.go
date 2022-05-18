package exemplo1

import (
	"errors"
	"fmt"
	"log"
	"regexp"
)

type Cpf struct {
	Numeros string
}

func CpfValido(cpf Cpf) (bool, error) {
	cpf.CleanCpf()
	fmt.Println(cpf.Numeros)
	var c Cpf
	if !c.Length() {
		return false, errors.New("CPF com tamanho inválido")
	}
	return true, nil
}

func (c *Cpf) CleanCpf() string {
	r, err := regexp.Compile(`/[\.\-]/g`)
	if err != nil {
		log.Fatal(err)
	}
	return r.ReplaceAllString(c.Numeros, " ")
}

func (c *Cpf) Length() bool {
	return len(c.Numeros) == 11
}
