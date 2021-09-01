package main

import "fmt"

//Abstract factory interface
type iSportsFactory interface{
	makeShoe() iShoe
	makeTshirt() iTshirts
}
//Abstract factory interface for the shoe
type iShoe interface{
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}
// Shoe structure 
type shoe struct{
	logo string 
	size int
}
//Abstract factory interface for the iTshirts
type iTshirts interface{
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type shirt struct{
	logo string
	size int
}

type nikeShirt struct{
	shirt
} 
type nikeShoe struct {
	shoe
}

type adidasShirt struct{
	shirt
} 
type adidasShoe struct {
	shoe
}

func (s *shoe) setLogo(logo string){
	s.logo  = logo
}
func (s *shoe) setSize(size int) {
	s.size = size
}
func (s *shoe) getLogo() string{
	return s.logo
}
func (s *shoe) getSize() int {
	return s.size
}

func (s *shirt) setLogo(logo string){
	s.logo  = logo
}
func (s *shirt) setSize(size int) {
	s.size = size
}
func (s *shirt) getLogo() string{
	return s.logo
}
func (s *shirt) getSize() int {
	return s.size
}
type adidas struct{}
type nike struct{}

func (a *adidas)makeShoe() iShoe{
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size:14,
		},
	}
}

func (a *adidas)makeTshirt() iTshirts{
	return &adidasShirt{
		shirt: shirt{
			logo: "adidas",
			size:44,
		},
	}
}

func (a *nike)makeShoe() iShoe{
	return &adidasShoe{
		shoe: shoe{
			logo: "Nike",
			size:14,
		},
	}
}

func (a *nike)makeTshirt() iTshirts{
	return &adidasShirt{
		shirt: shirt{
			logo: "Nike",
			size:44,
		},
	}
}

func getsportfactory(brand string) (iSportsFactory , error){
	if brand == "adidas"{
		return &adidas{}, nil 
	}
	if brand == "nike"{
		return &nike{},nil
	}
	return nil , fmt.Errorf("wrong brand given")
}
func printShoeDetails(s iShoe) {
    fmt.Printf("Logo: %s", s.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.getSize())
    fmt.Println()
}

func printShirtDetails(s iTshirts) {
    fmt.Printf("Logo: %s", s.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.getSize())
    fmt.Println()
}

func main(){
	adidasfactory ,_ := getsportfactory("adidas")
	nikefactory ,_ := getsportfactory("nike")

	adidasShoe := adidasfactory.makeShoe()
	adidasShirt := adidasfactory.makeTshirt()

	nikeShoe := nikefactory.makeShoe()
	nikeShirt := nikefactory.makeShoe()

	printShoeDetails(nikeShoe)
    printShirtDetails(nikeShirt)

    printShoeDetails(adidasShoe)
    printShirtDetails(adidasShirt)

}