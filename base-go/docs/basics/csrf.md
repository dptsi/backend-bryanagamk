# Proteksi CSRF

## Pengenalan

[CSRF](https://en.wikipedia.org/wiki/Cross-site_request_forgery) adalah jenis eksploitasi berbahaya di mana perintah tidak sah dilakukan atas nama pengguna yang diautentikasi.

### Penjelasan Vulnerability

Bayangkan aplikasi anda memiliki rute `/user/email` yang menerima `POST` request untuk mengubah email pengguna. Kemunghkinan rute ini menginginkan field berupa `email` yang berisikan alamat email dari pengguna yang ingin digunakan. Tanpa proteksi CSRF, website lain dapat membuat form HTML yang mengarahkan ke rute tersebut dan mengganti email dari pengguna tersebut.

```html
<form action="https://your-application.com/user/email" method="POST">
  <input type="email" value="malicious-email@example.com" />
</form>

<script>
  document.forms[0].submit()
</script>
```

Jika web tersebut otomatis submit form pada saat laman dimuat, pengguna hanya perlu mengunjungi website tersebut dan emailnya akan berubah. Untuk mencegah ini, kita wajib menginpeksi request dengan method selain `GET`, `HEAD`, dan `OPTIONS` dan mencocokkan dengan secret yang tidak dapat diakses oleh website lain.

## Mencegah CSRF

Base project ini otomatis melakukan generate CSRF token untuk setiap session yang ada. Token ini digunakan untuk melakukan verifikasi. Token dapat diakses oleh front-end dengan mengambil nilai cookie `CSRF-TOKEN` yang akan diberikan oleh back-end untuk setiap request. Oleh karena itu, sebelum pengguna melakukan log in, front-end perlu memanggil route `/csrf-cookie`. Route tersebut hanya berfungsi untuk memberikan nilai cookie tersebut. Setelah itu, untuk setiap request dari front-end dengan method selain dari yang disebutkan di atas perlu memberikan nilai CSRF token tersebut melalui header `X-CSRF-TOKEN`.

```js
function getCookie(key) {
  const cookies = document.cookie.split('; ')
  for (let i = 0; i < cookies.length; i++) {
    const [cKey, value] = cookies[i].split('=')
    if (key === cKey) return value
  }
  return null
}

// Memanggil POST request
fetch('/user/email', {
  credentials: 'include',
  method: 'POST',
  body: JSON.stringify({ email: 'dptsi@its.ac.id' }),
  headers: {
    'X-CSRF-TOKEN': getCookie('CSRF-TOKEN'),
  },
})
```
