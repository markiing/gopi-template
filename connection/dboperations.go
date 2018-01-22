package connection

import (
	m "gopkg.in/mgo.v2"
)

var database = "go";

func openSession()(*m.Session,error){
	s,err := m.Dial("localhost")

	if err != nil {
		panic("Oops. Houve um problema ao estabelecer a conexão com a base de dados")
	}

	return s,err
}

func withCollection(collection string) (*m.Collection, *m.Session){
	c,_ := openSession()
	return c.DB(database).C(collection),c
}

func FindOne(collection string,rules interface{},param interface{})(a *interface{}){
	c, s := withCollection(collection)
	err := c.Find(rules).One(&param)
	if err != nil{
		panic("Oops, houve um problema ao executar o metodo FindOne !");
	}
	defer CloseSession(s)
	return &param
}

func Insert(collection string, param interface{}){
	c, s := withCollection(collection)
	c.Insert(param)
	defer CloseSession(s)
}

func Delete(collection string, param interface{}){
	c,s := withCollection(collection)
	c.Remove(param)
	defer CloseSession(s)
}


func CloseSession(session *m.Session){
	session.Close()
}