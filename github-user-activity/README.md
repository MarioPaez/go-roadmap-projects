# GitHub Activity CLI

This is a simple command-line interface (CLI) application that fetches and displays the recent activity of a GitHub user. It helps you practice working with REST APIs, handling JSON, and building CLI tools without external libraries.

## 🛠️ How It Works

Run the application with:

```bash
go run main.go github-activity <username>
```

Example:

```bash
go run main.go github-activity kamranahmedse
```

The application will fetch and display the most recent public activity for the specified GitHub user using the GitHub Events API:

```
https://api.github.com/users/<username>/events
```

## 🧾 Example Output

```
- Pushed 3 commits to kamranahmedse/developer-roadmap
- Opened a new issue in kamranahmedse/developer-roadmap
- Starred kamranahmedse/developer-roadmap
```

## ⚠️ Error Handling

- Invalid usernames or network/API errors are handled gracefully.
- No external libraries are used.
