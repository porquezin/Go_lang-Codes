package models

import "ex2/db"

type Produto struct {
	Nome, Desc string
	Preco      float64
	Quant, ID  int
}

func BuscaAll() []Produto {
	db := db.ConectDB()
	selectAll, err := db.Query("select * from produtos order by id asc")

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

		p.ID = id
		p.Nome = nome
		p.Desc = desc
		p.Preco = preco
		p.Quant = quant

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

// func buscaProd() {

// }

func AddProd(nome, desc string, preco float64, quant int) {
	db := db.ConectDB()
	query, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}
	query.Exec(nome, desc, preco, quant)
	defer db.Close()
}

func Delete(id string) {
	db := db.ConectDB()

	del, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}
	del.Exec(id)
	defer db.Close()
}

func Edit(id string) Produto {
	db := db.ConectDB()
	query, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	prod := Produto{}

	for query.Next() {
		var id, quant int
		var nome, desc string
		var preco float64

		err = query.Scan(&id, &nome, &desc, &preco, &quant)
		if err != nil {
			panic(err.Error())
		}
		prod.ID = id
		prod.Nome = nome
		prod.Desc = desc
		prod.Preco = preco
		prod.Quant = quant
	}
	defer db.Close()
	return prod
}

func Update(id int, nome, desc string, preco float64, quant int) {
	db := db.ConectDB()

	up, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	up.Exec(nome, desc, preco, quant, id)
	defer db.Close()
}
