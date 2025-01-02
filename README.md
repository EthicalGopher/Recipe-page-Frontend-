# Simple Omelette Recipe Page

This is a simple and responsive web page designed to display a recipe for a classic omelette. The page adapts its layout for desktop and mobile devices, making it user-friendly and visually appealing.

## Features
- Fully responsive layout with separate image handling for desktop and mobile.
- Styled using Bootstrap 4.5.2 and custom CSS.
- Includes preparation time, ingredients, instructions, and nutritional values.
- Easily deployable with a Fiber server for serving static files.

## Technologies Used
- **HTML5**: For semantic structure.
- **CSS3**: Custom styling, including media queries for responsiveness.
- **Bootstrap 4.5.2**: For responsive layout and utility classes.
- **Google Fonts**: Typography using `Young Serif` and `Outfit`.
- **Go with Fiber**: Lightweight server to serve static files.

## Getting Started

### Prerequisites
Ensure you have:
1. A modern web browser installed (e.g., Google Chrome, Mozilla Firefox).
2. Go programming language installed ([Download Go](https://go.dev/)).

### Installation
1. Clone the repository:
   ```bash
   git clone <repository-url>
   ```
2. Navigate to the project directory:
   ```bash
   cd recipe-page
   ```
3. Install the required Go package:
   ```bash
   go get github.com/gofiber/fiber/v2
   ```
4. Run the server:
   ```bash
   go run main.go
   ```
5. Open your browser and navigate to `http://localhost:5430` to view the page.

## File Structure
```
recipe-page/
├── Frontend/
│   ├── assets/
│   │   └── images/
│   │       ├── image-omelette.jpeg
│   │       ├── preview.png
│   ├── index.html
│   ├── style.css
│   └── script.js
├── main.go
├── go.mod
├── go.sum
└── README.md
```

## Customization
You can update the recipe or modify the design by editing the `index.html` and `<style>` block within the `<head>` tag. For instance:

- Change the recipe content under `<h1 class="heading">Simple Omelette Recipe</h1>`.
- Replace `image-omelette.jpeg` with your own image in the `/assets/images/` folder.
- Modify `main.go` to change the server settings or directory paths.

## Preview
To view the page:
1. Run the Go server as described above.
2. Open `http://localhost:5430` in your browser.

## Screenshot
![Screenshot of Recipe Page](Frontend/assets/images/preview.png)

## Credits
- **Bootstrap**: [https://getbootstrap.com/](https://getbootstrap.com/)
- **Google Fonts**: [https://fonts.google.com/](https://fonts.google.com/)
- **Fiber**: [https://gofiber.io/](https://gofiber.io/)

## License
This project is licensed under the MIT License. Feel free to use, modify, and distribute this code as you like.

---

Feel free to contribute and improve this project. Happy coding!
# Recipe-page-Frontend-
