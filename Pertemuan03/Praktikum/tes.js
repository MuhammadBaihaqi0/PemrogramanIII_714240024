const jsonFromServer = '{"nama":"bai","umur":19,"kota":"Bandung"}';

console.log(jsonFromServer);

const userObject = JSON.parse(jsonFromServer);
console.log("Nama Mahasiswa:", userObject.nama);
