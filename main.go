package main

import "fmt"

type celcius struct {
	suhu float64
}

type farenheit struct {
	suhu float64
}

type kelvin struct {
	suhu float64
}

func (c celcius) toCelcius() float64 {
	return c.suhu
}

func (c celcius) toFarenheit() float64 {
	return ((9.0 / 5.0) * c.suhu) + 32
}

func (c celcius) toKelvin() float64 {
	return c.suhu + 273.15
}

func (f farenheit) toCelcius() float64 {
	return ((5.0 / 9.0) * (f.suhu - 32))
}

func (f farenheit) toFarenheit() float64 {
	return f.suhu
}

func (f farenheit) toKelvin() float64 {
	return ((5.0 / 9.0) * (f.suhu - 32)) + 273.15
}

func (k kelvin) toCelcius() float64 {
	return k.suhu - 273.15
}

func (k kelvin) toFarenheit() float64 {
	return ((9.0 / 5.0) * (k.suhu - 273.15)) + 32
}

func (k kelvin) toKelvin() float64 {
	return k.suhu
}

type converterSuhu interface {
	toCelcius() float64
	toFarenheit() float64
	toKelvin() float64
}

func main() {
	var input int
	var tujuan int
	var suhu float64

	fmt.Println(`
		Program Konversi Suhu 
		1. Celcius 
		2. Farenheit 
		3. Kelvin 
	`)

	fmt.Print(" MASUKAN PILIHAN : ")
	fmt.Scanf("%d", &input)

	if input < 1 || input > 3 {
		fmt.Println(" PILIHAN TIDAK VALID ")
		return
	}

	fmt.Print(" MASUKAN SUHU : ")
	fmt.Scanf("%f", &suhu)

	fmt.Print(" Pilih Suhu Tujuan : ")
	fmt.Scanf("%d", &tujuan)

	match := map[int]converterSuhu{
		1: celcius{suhu},
		2: farenheit{suhu},
		3: kelvin{suhu},
	}

	mCalculate := map[int]float64{
		1: match[input].toCelcius(),
		2: match[input].toFarenheit(),
		3: match[input].toKelvin(),
	}

	fmt.Printf(" Suhu Awal : %.2f , Suhu Tujuan : %.2f", suhu, mCalculate[tujuan])
}
