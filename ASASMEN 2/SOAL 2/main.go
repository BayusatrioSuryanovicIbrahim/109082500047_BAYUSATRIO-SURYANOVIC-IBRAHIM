package main

import "fmt"

const nMax = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func cariNilaiPertama(T arrayMahasiswa, N int, nimTarget string) int {
	for i := 0; i < N; i++ {
		if T[i].NIM == nimTarget {
			return T[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(T arrayMahasiswa, N int, nimTarget string) int {
	maxNilai := -1
	for i := 0; i < N; i++ {
		if T[i].NIM == nimTarget {
			if T[i].nilai > maxNilai {
				maxNilai = T[i].nilai
			}
		}
	}
	return maxNilai
}

func main() {
	var T arrayMahasiswa
	var N int
	var nimTarget string

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&N)

	if N > nMax {
		N = nMax
	}

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimTarget)

	nilaiPertama := cariNilaiPertama(T, N, nimTarget)
	nilaiTerbesar := cariNilaiTerbesar(T, N, nimTarget)

	if nilaiPertama != -1 {
		fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", nimTarget, nilaiPertama)
		fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", nimTarget, nilaiTerbesar)
	} else {
		fmt.Printf("Data mahasiswa dengan NIM %s tidak ditemukan.\n", nimTarget)
	}
}
