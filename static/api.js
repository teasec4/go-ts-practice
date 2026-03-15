export const api = {
  getUsers() {
    return fetch("http://localhost:8080/users")
      .then(r => r.json())
  },
  createUser(name) {
    return fetch("http://localhost:8080/users", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name }),
    })
      .then(r => r.json())
  },
}