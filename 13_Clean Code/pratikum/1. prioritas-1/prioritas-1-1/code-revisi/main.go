package main

// menggunakan huruf kapital diawal `User` untuk membedakan mana yang struct / fungsi atau instance struct / parameter
type User struct {
	id       int
	// merubah tipe data menjadi `string` karena tidak ada username yang menggunakan integer/numeric saja
	username string
	// merubah tipe data menjadi `string` karena tidak ada password yang menggunakan integer/numeric saja
	password string
}

// menggunakan CamelCase dan untuk struct menggunakan huruf kapital diawal jadi `UserService` untuk mempermudah membedakan mana yang struct / fungsi atau instance struct / parameter dan juga mudah untuk dieja
type UserService struct {
	// merubah menjadi `users []User` karena merupakan array atau banyak user dari struct `User`. sehingga akan lebih mudah untuk dipahami maksud dari propertinya
	users []User
}

// merubah method ke CamelCase menjadi `getAllUsers` sehingga lebih mudah dieja dan dibaca. 
// untuk isntance parameter dirubah menjadi `userService` yang merupakan instance dari `UserService` sehingga lebih mudah dibaca maupun dipahami.
// untuk returnnya mengikuti perubahan sebelumnya
func (userService UserService) getAllUsers() []User {
	return userService.users
}

// pada penamaan fungsi atau method atau variabel harus konsisten dimana bila CamelCase maka seterusnya akan CamelCase jadi nama method dirubah menjadi `getUserById` sehingga lebih mudah dieja dan dibaca. 
// untuk isntance parameter dirubah menjadi `userService` yang merupakan instance dari `UserService` sehingga lebih mudah dibaca maupun dipahami
func (userService UserService) getUserById(id int) User {
	// merubah menjadi `user` karena akan lebih mudah dipahami maksud dari variabel tersebut sesuai dengan bahasa inggris dimana secara umum penambahan s untuk jamak dan tanpa s untuk tunggal
	// untuk userService.users dan returnnya mengikuti perubahan sebelumnya
	for _, User := range userService.users {
		if id == User.id {
			return User
		}
	}

	return User{}
}