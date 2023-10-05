<a name="readme-top"></a>

<!-- [![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url] -->

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/zydhanlinnar11/base-go">
    <img src="https://go.dev/images/go-logo-white.svg" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Base Project Golang</h3>

  <p align="center">
    Base project untuk pengembangan back-end menggunakan Go
    <br />
    <a href="https://github.com/zydhanlinnar11/base-go"><strong>Eksplor Dokumentasi »</strong></a>
    <br />
    <br />
    <a href="https://github.com/zydhanlinnar11/base-go">Lihat Demo</a>
    ·
    <a href="https://github.com/zydhanlinnar11/base-go/issues">Laporkan Bug</a>
    ·
    <a href="https://github.com/zydhanlinnar11/base-go/issues">Request Fitur</a>
  </p>
</div>

<!-- TABLE OF CONTENTS
<details>
  <summary>Daftar Isi</summary>
  <ol>
    <li>
      <a href="#tentang-project">Tentang Project</a>
      <ul>
        <li><a href="#dibangun-dengan">Dibangun Dengan</a></li>
      </ul>
    </li>
    <li>
      <a href="#memulai">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details> -->

<!-- ABOUT THE PROJECT -->

## Tentang Project

<!-- [![Product Name Screen Shot][product-screenshot]](https://example.com)

TBD

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->

### Dibangun Dengan

- [![Go][Go]][Go-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Memulai

### Prasyarat

- Go ([Link download dan install Go](https://go.dev/doc/install))

### Menjalankan Aplikasi

1. Clone repositori

   ```bash
   // Menggunakan SSH
   git clone git@github.com:zydhanlinnar11/base-go.git

   // Menggunakan HTTPS
   git clone https://github.com/zydhanlinnar11/base-go.git
   ```

2. Copy .env.example ke .env

   ```bash
   cp .env .env.example
   ```

3. Set up environment variable untuk OpenID Connect (Bisa didapatkan dari Secman)

   ```bash
   OIDC_PROVIDER=
   OIDC_CLIENT_ID=
   OIDC_CLIENT_SECRET=
   OIDC_REDIRECT_URL=
   OIDC_SCOPES=
   OIDC_END_SESSION_ENDPOINT=
   OIDC_POST_LOGOUT_REDIRECT_URI=
   ```

4. Jalankan server pada port 8080
   ```bash
   go run main.go
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Penggunaan

> Pastikan server sudah berjalan saat mengakses dokumentasi!

_Penjelasan lebih lengkap terdapat pada [/doc/project](http://localhost:8080/doc/project) (Dokumentasi project) dan [/doc/api](http://localhost:8080/doc/api) (Dokumentasi OpenAPI 3.1)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->

## Roadmap

- [x] Dependency Injection
- [x] Routing menggunakan Gin
- [x] Delegasi config setiap modul
- [x] Session
- [x] Autentikasi menggunakan session
- [x] Autentikasi menggunakan Basic Auth header
- [x] CSRF Protection
- [x] Common error & response
- [x] OpenID Connect
- [x] Script make controller & module
- [x] Dokumentasi project
- [x] Dokumentasi OpenAPI 3.1

<!-- See the [open issues](https://github.com/zydhanlinnar11/base-go/issues) for a full list of proposed features (and known issues). -->

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->

<!-- ## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->

<!-- LICENSE -->

<!-- ## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->

<!-- CONTACT -->

## Contact

Zydhan Linnar Putra - zyd@its.ac.id

Project Link: [https://github.com/zydhanlinnar11/base-go](https://github.com/zydhanlinnar11/base-go)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

<!-- ## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

- [Choose an Open Source License](https://choosealicense.com)
- [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
- [Malven's Flexbox Cheatsheet](https://flexbox.malven.co/)
- [Malven's Grid Cheatsheet](https://grid.malven.co/)
- [Img Shields](https://shields.io)
- [GitHub Pages](https://pages.github.com)
- [Font Awesome](https://fontawesome.com)
- [React Icons](https://react-icons.github.io/react-icons/search)

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[Go]: https://img.shields.io/badge/Go-007d9c?style=for-the-badge&logo=go&logoColor=FFFFFF
[Go-url]: https://go.dev/
