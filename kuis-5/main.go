package main

import "fmt"

func main() {
	// nilai:= []int {80, 65, 90, 70}
	// max,total:= 0,0
	// lulus := []int{}
	// remed := []int{}
	
	// for _, n:= range nilai {
	// 	if n > max {
	// 		max = n
	// 	}
	// 	total += n
	// 	if n >= 70 {
	// 		lulus = append(lulus, n)
	// 	} else {
	// 		remed = append(remed, n)
	// 	}
		
	// }
	// a:= float64(total)/ float64(len(nilai))
	// // average := float64(total) / float64(len(nilai))
	// fmt.Println("Nilai Tertinggi:", max)
	// fmt.Println("Rata-rata Nilai:", a)
	// fmt.Println("Lulus : ", lulus)
	// fmt.Println("Remedial: ",remed)


	harga := []int{15000, 8000, 12000}
	totHarga := 0
	for _, h := range harga {
		totHarga += h
	}
	fmt.Println("Total harga : %d", totHarga)




	nilaiSiswa := map[string]int{
		"Ali": 80,
		"Budi": 65,
		"Citra": 90,
	}
	for nama, nilai := range nilaiSiswa {
		fmt.Println(nama, "Mendapat Nilai:", nilai)
	}



	stokBarang := map[string]int{
		"Pensil": 10,
		"Pulpen" : 20,
		"Buku": 5,
	}
	totalBarang, jenisBarang := 0, 0

	for _, stok := range stokBarang {
		jenisBarang++
		totalBarang += stok
	}

	fmt.Println("Jenis Barang:", jenisBarang)
	fmt.Println("Total Barang:", totalBarang)

	

	stokBarang2 := map[string]int{
		"Kopi":16000,
		"Roti": 13000,
		"Teh":9000,
	}

	hargaBaru := 0
	for _, harga := range stokBarang2 {
		hargaBaru += harga
	}
	fmt.Println (stokBarang2)


}