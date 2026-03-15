import readline from 'readline';

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

function question(prompt: string): Promise<string> {
  return new Promise((resolve) => {
    rl.question(prompt, (answer) => {
      resolve(answer);
    });
  });
}

async function main() {
  console.log("Go backend CLI tester (SQLite users) - Bun version")
  
  while (true) {
    const cmd = (await question("> ")).trim();
    
    if (cmd === "exit") break;
    
    if (cmd.startsWith("create ")) {
      const name = cmd.slice(7).trim();
      if (!name) {
        console.log("Введите имя после 'create'");
        continue;
      }

      const res = await fetch("http://localhost:8080/users", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name }),
      });
      const data = await res.json();
      console.log("Created:", data);
    } else if (cmd === "list") {
      const res = await fetch("http://localhost:8080/users");
      const data = await res.json();
      console.log(data);
    } else if (cmd.startsWith("get ")) {
      const id = cmd.slice(4).trim();
      const res = await fetch(`http://localhost:8080/user?id=${id}`);
      const data = await res.json();
      console.log(data);
    } else {
      console.log("Команды:");
      console.log("  create <name>  - создать пользователя");
      console.log("  list           - показать всех пользователей");
      console.log("  get <uuid>     - получить пользователя по UUID");
      console.log("  exit           - выйти");
    }
  }
  
  rl.close()
}

main().catch(console.error)