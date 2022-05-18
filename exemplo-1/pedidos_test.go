package exemplo1

import "testing"

func TestPedidoCPFInvalido(t *testing.T) {
	var c Cpf
	c.Numeros = "10231234541"
	validaCPF, err := CpfValido(c)
	if err != nil {
		t.Errorf("Erro no CPF ")
	}
	if !validaCPF {
		t.Errorf("CPF Inválido")
	}

}
