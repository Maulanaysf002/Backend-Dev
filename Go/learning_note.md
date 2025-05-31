- dalam 1 folder golang hanya boleh ada 1 package.
- jika ingin mengeksport func dari package lain gunakan kapital didepannya, misal Hello(){}.
- go mod init "nama module" adalah cara inisiasi modul yang akan digunakan untuk import package.

# Basic Sintax

```go
package main

import "fmt"

func main(){
  fmt.Println("Hello Duniaku!")
}
```

- sintax diatas adalah deklarasi main function yang wajib ada dalam go.
- golang akan otomatis menjalankan main funtion sebagai funtion utamanya.
- import fmt merupakan import library fmt yang artinya format untuk memanggil function Println yang bisa mencetak string.

```bash
go run main.go

go build -o bin main.go
```

- dua perintah diatas digunakan untuk menjalankan program golang yang menggunakan compiler.
- perintah 1 untuk menjalankan tanpa mengcompilenya
- perintah 2 code go dicompile menjadi app dan baru dijalankan.

go build
→ Perintah untuk meng-compile file Go.

-o bin
→ Opsi -o (output): memberi nama dan lokasi file output hasil compile.
Dalam hal ini, hasil compile akan disimpan sebagai file bernama bin.

main.go
→ File sumber Go yang akan di-compile.

# chekpoint

https://dasarpemrogramangolang.novalagung.com/A-operator.html (operator) 🚩
