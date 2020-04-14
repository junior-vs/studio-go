package contas

import "github.com/junior-vs/studio-go/banco/clientes"

type ContaCorrente struct {
	Titular   clientes.Titular
	NuAgencia int16
	NuConta   int64
	Saldo     float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "saque realizado com sucesso", c.Saldo
	} else {
		return "saldo insuficiente", c.Saldo
	}
}
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.Saldo += valorDoDeposito
		return "Valor deposistado. Saldo", c.Saldo
	} else {
		return "Valor insuficiente para deposito", c.Saldo
	}
}
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}

}
