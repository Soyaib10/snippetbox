# **Snippetbox** - Effortlessly Craft Your Snippet In a Flash

Snippetbox is a web application similar to Ubuntu Pastebin or GitHub Gist, designed for users to create and save text snippets. Registered users can create their own snippets, set an expiration time for each snippet, and view saved snippets.

---

## **Key Features**

### For Users:
- **Create Account/Login**: Users can create an account or log in to access their snippets.
- **Create Snippets**: Users can create and save their snippets.
- **Set Lifetime**: Each snippet can have an expiration time set by the user.
- **View Snippets**: Users can view their created snippets.

---

## **Tech Stack**

### Frontend:
- HTML, CSS, Bootstrap, JavaScript

### Backend:
- Go (Golang)

### Database:
- MySQL

### Packages:
- Pat Router
- Alice Middleware Manager
- Golangcollege Session Manager
- Justin NoSurf Package for CSRF Protection

---

## **Limitations**
- Features such as search or tagging are not yet implemented. However, I am working on it.

---

## **Demonstration**

#### 1. Homepage
![Homepage](./ui/static/img/screenshoot/homepage_snippet.png)
The homepage provides an overview of the application and options to log in or register. Unregistered users can only view the snippets, but to create and save snippets, users need to register or log in if they already have an account.

#### 2. Registration Page
![Registration Page](./ui/static/img/screenshoot/signup.png) 
![No Duplicate Users Allowed](./ui/static/img/screenshoot/signupReq.png)
Users can register by filling in the registration form. If more than one user enters the same email, an error message will be displayed indicating that the user already exists.

#### 3. Login Page
![Login Page](./ui/static/img/screenshoot/login.png)
Users can log in to create their snippets.

#### 4. Create Snippet
![Create Snippet](./ui/static/img/screenshoot/createSnippet.png)
Users can create and save new snippets and set expiration times.

#### 5. See Snippet
![See Snippet](./ui/static/img/screenshoot/seesnippet.png)
Users can view their saved snippets.

---

## **Download Instructions**
Download the source code for "<b>Snippetbox</b>":

- **For Linux Only**
- FIRST, download the source code.
- Extract the file and copy the "Snippetbox" folder.
- Open this folder in your favorite editor (e.g., <b>VS Code</b>).
- Install XAMPP and run the following command to start MySQL:

    ```bash
    sudo /opt/lampp/lampp start
    ```

- Open PHPMyAdmin (http://localhost/phpmyadmin)
- Create a database named <b>snippetbox</b>.
- Import the `snippetbox.sql` file into the newly created database.
- Open a terminal, navigate to your project directory, and run these two commands:

    ```bash
    go mod tidy
    ```

    ```bash
    go run cmd/web/*
    ```

- Open your browser and go to http://localhost:8080/

**LOGIN DETAILS**

You can register a new user and log in to create new snippets.

Hopefully, you can run this project successfully. If you encounter any issues, feel free to contact me using the email link provided in the last section.

---

## 🌟 You've Made It This Far!

Wow, that’s amazing! Thank you so much for your interest and support.

If you’ve enjoyed working with this project, a ⭐️ would be greatly appreciated. Your support motivates me to keep improving and working on more projects.

Feel free to share any advice or thoughts you have—I’d love to hear from you!

---

**🔗 [Give a Star ⭐️](https://github.com/Soyaib10/snippetbox)**

**💬 [Share Your Thoughts](mailto:soyaibzihad10@gmail.com)**

Once again, thank you for your encouragement! 😊
