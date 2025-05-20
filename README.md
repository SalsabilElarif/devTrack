# devTrack 🛠️  
A simple and efficient CLI task tracker built with Go.

## 📌 Overview  
**devTrack** is a lightweight command-line tool that helps developers and teams manage tasks directly from the terminal. Designed for speed and simplicity, it supports basic task operations such as adding, listing, completing, and removing tasks.

## 🧰 Built With  
- [Go](https://golang.org/) – for performance and portability  
- Standard Library (no external dependencies)

## ✨ Features  
- Add tasks with titles and optional descriptions  
- List tasks in a readable format  
- Mark tasks as complete  
- Delete tasks by ID  
- Data persistence with local file storage (JSON)

## 🚀 Getting Started  

### Prerequisites  
- Go installed (version 1.18+ recommended)  
- Git

### Installation  

```bash
# Clone the repo
git clone https://github.com/SalsabilElarif/devTrack.git

# Navigate to the project directory
cd devTrack

# Build the binary
go build -o devtrack
```

### Usage  

```bash
./devtrack add "Write documentation"
./devtrack list
./devtrack complete 1
./devtrack delete 1 
```

## 🧪 To Do / Planned Features
- Due dates and reminders
- Priority tags
- Filtering/sorting options
- Sync with cloud (e.g., Firebase or Supabase)
- Unit tests

## 🙋‍♀️ About Me
Built by Salsabil Elarif – trainee full-stack developer passionate about learning Go, Python, JavaScript, and React.
