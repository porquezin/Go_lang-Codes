package models

import "ex2/db"

type Produto struct {
	Nome, Desc string
	Preco      float64
	Quant, ID  int
}

func BuscaAll() []Produto {
	db := db.ConectDB()
	selectAll, err := db.Query("select * from produtos;")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectAll.Next() {
		var id, quant int
		var nome, desc string
		var preco float64

		err = selectAll.Scan(&id, &nome, &desc, &preco, &quant)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Desc = desc
		p.Preco = preco
		p.Quant = quant

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}
