package contas

import "ex1/clientes"

type ContaPoupanca struct {
	Titular                     clientes.Titular
	Agencia, NumConta, Operacao int
	saldo                       float64
}

func (c *ContaPoupanca) Sacar(saque float64) (string, float64) {
	if saque > 0 && saque < c.saldo {
		c.saldo -= saque
		return "foi", c.saldo
	} else {
		return "nao foi", c.saldo
	}
}

func (c *ContaPoupanca) Depositar(deposito float64) (string, float64) {
	if deposito > 0 {
		c.saldo += deposito
		return "foi", c.saldo
	} else {
		return "nao foi", c.saldo
	}
}

func (c *ContaPoupanca) Transf(
	transf float64,
	conta *ContaCorrente,
) bool {
	if transf <= c.saldo && transf > 0 {
		c.saldo -= transf
		conta.Depositar(transf)
		return true
	} else {
		return false
	}
}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}
