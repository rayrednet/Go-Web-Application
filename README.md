# Go Projects

## ⭐ Identity

- **Name:** Rayssa Ravelia  
- **NRP:** 5025211219  
- **Class:** Pemrograman Berbasis Kerangka Kerja (D)  


## 📄 Overview

This Go project is part of my course, **Pemrograman Berbasis Kerangka Kerja (D)**. The project consists of four main sections, each located in a dedicated directory:

1. **data-access**  
   - **Description:** This section demonstrates database access in Go.  
   - **Reference:** [Go Database Access Documentation](https://go.dev/doc/tutorial/database-access)

2. **gorm**  
   - **Description:** This section covers the use of GORM, a powerful ORM library for Go.  
   - **Reference:** [GORM Documentation](https://gorm.io/docs/)

3. **gowiki**  
   - **Description:** This section implements a simple wiki application in Go.  
   - **Reference:** [Go Wiki Tutorial](https://go.dev/doc/articles/wiki/)

4. **web-service-gin**  
   - **Description:** This section demonstrates building a web service using the Gin framework.  
   - **Reference:** [Go Gin Web Service Tutorial](https://go.dev/doc/tutorial/web-service-gin)

Each directory contains a `README.md` file with detailed explanations specific to its respective section.

## 📂 File Structure

```
GO-WEB-APPLICATION/                # Main project directory
|
├── data-access/            # Demonstrates database access in Go
│   ├── main.go
│   ├── go.mod
│   ├── README.md
│   └── ...
│
├── gorm/                   # Covers the use of GORM library
│   ├── main.go
│   ├── go.mod
│   ├── README.md
│   └── ...
│
├── gowiki/                 # Implements a simple wiki application
│   ├── data/               # Stores page data in .txt files
│   │   ├── ANewPage.txt
│   │   ├── FrontPage.txt
│   │   └── NewPage.txt
│   │
│   ├── tmpl/               # HTML templates and CSS for rendering pages
│   │   ├── edit.html
│   │   ├── view.html
│   │   └── style.css
│   │
│   ├── img/                # Screenshots of the application views
│   │   ├── front-page-view.png
│   │   ├── front-page-edit.png
│   │   ├── new-page.png
│   │   ├── new-page-edit.png
│   │   └── new-page-view.png
│   │
│   ├── README.md           # Documentation for the gowiki section
│   ├── go.mod
│   ├── wiki.go             # Main application logic
│   └── TestPage.txt        # Sample page data file
│
├── web-service-gin/        # Demonstrates building a web service using Gin
│   ├── main.go
│   ├── go.mod
│   ├── README.md
│   └── ...

```

## 📚 How to Use

1. Clone the repository:
   ```bash
   git clone git@github.com:rayrednet/Go-Web-Application.git
   ```

2. Navigate to the desired section directory (e.g., `gowiki/`):
   ```bash
   cd gowiki
   ```

3. Follow the instructions in the respective `README.md` to run the code.

---

Feel free to explore each section and learn from the provided examples!

