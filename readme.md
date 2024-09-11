# **Snippetbox** - Effortlessly craft your snippet in a flash

This is a web application similar to Ubuntu Pastebin or GitHub Gist, designed for users to create and save text snippets. Registered users can create their own snippets, set an expiration time for each snippet, and view saved snippets.

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
- Features such as search or tagging are implemented yet. But I am working on it.

---

## **Demonstration**



#### 1. Homepage
![Homepage](./ui/static/img/screenshoot/homepage_snippet.png)
The homepage provides an overview of the application and options to log in or register. Unregistered users can only see the snipeets but to create a snippet and save it users have to register to create an account or to login if already have an account.

#### 2. Registration Page
![Registration Page](./ui/static/img/screenshoot/signup.png) 
![No Duplicate Users Allowed](./ui/static/img/screenshoot/signupReq.png)
Users can register by filling in the registration form. If more then one users enter same email then error showing user already exists.

#### 3. Login Page
![Login Page](./ui/static/img/screenshoot/login.png)
Users can log in to create their snippets.

#### 4. Create Snippet
![Create Snippet](./ui/static/img/screenshoot/createSnippet.png)
Users can create and save new snippets along with the setting of expiration time.

#### 5. See Snippet
![See Snippet](./ui/static/img/screenshoot/seesnippet.png)
Users can view their saved snippets.

---
