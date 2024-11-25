# FastCode: Snippet Management Application

FastCode is a desktop application designed for developers to create, edit, search, and manage code snippets efficiently and effectively. With its simple interface, you can organize your snippets and access them quickly when needed.

---

## 🚀 Features

- **Snippet Management**:
    - Add, edit, and delete snippets easily.
- **Theme Support**:
    - Default **Dark Theme**.
    - Option to switch to **Light Theme**.
- **Search Functionality**:
    - Quickly search through snippet titles.
- **Persistent Storage**:
    - Snippets are stored in a `snippets.json` file for permanence.
- **User-Friendly Interface**:
    - Minimal and intuitive design for quick access.

---

## 📂 Folder Structure

```plaintext
FastCode/
├── main.go          # Main application file
├── snippets.json    # JSON file to store snippets
├── services/        # Folder for application services
│   └── snippet.go   # Service for snippet management
└── README.md        # Project documentation
```

## 📖  Usage Guide

- Make sure Go is installed on your system:
  `$ go version`

- Navigate to the project folder:
  `$ cd fast-paste`

- Run the application:
  `$ go run main.go`

##   Application Features

### Snippet Ekleme

1. Enter a title in the "Subject" field.
2. Write your code snippet in the "Snippet" field.
3. Click the Save button to save your snippet.

###  Searching Snippets

1. Use the search bar at the top to search for a snippet title.
2. Matching results will appear in the list.

### Editing Snippets

1. Click on a snippet from the list to select it.
2. Make the desired changes in the "Title" and "Content" fields.
3. Click the Save button to save the updated snippet.

### Deleting Snippets

1. Select the snippet you want to delete from the list.
2. Click the Delete button to remove the snippet.

### Switching Themes

1. Use the theme selector dropdown to choose between "Dark" and "Light" themes.
2. Light theme adjusts all inputs and text areas to have a gray background and darker text for visibility.


## 🛠️ Technologies Used

- Programming Language: Go (Golang)
- UI Framework: Fyne (https://fyne.io)


## 🤝 Contributing

1. Fork this repository:
   `$ git clone https://github.com/Tahaknd/fast-paste.git`

2. Create a new branch:
   `$ git checkout -b new-feature`

3. Make your changes and commit them:
   `$ git commit -m "New feature added: ..."`

4. Push your changes to the remote repository:
   `$ git push origin new-feature`

5. Submit a Pull Request.


## 🐞  Issues and Feedback

- If you encounter any issues or have suggestions for new features, feel free to create an issue here: https://github.com/Tahaknd/fast-paste/issues

