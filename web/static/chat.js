
// const userCode = localStorage.getItem("userCode");
const userCode = "082a60c4-c494-4415-9d23-b0cdd9c11b07"
const ws = new WebSocket(`wss://127.0.0.1/ws?code=${userCode}`);


ws.onopen = () => {
	console.log("connected to websocket")
};


fetch("/api/messages")
  .then(res => res.json())
  .then(data => {
	const chatbox = document.getElementById("chatbox");
    for (const msg of data.reverse()) {
      const time = new Date(msg.timestamp).toLocaleTimeString();
      chatbox.innerHTML += `<div><strong>${msg.sender}</strong> [${time}]: ${msg.content}</div>`;
    }
  });


ws.onmessage = (event) => {
	const data = JSON.parse(event.data);
	const chatbox = document.getElementById("chatbox");
  
	if (data.type === "online_users") {
	  const list = data.users.map(name => `<li>${name}</li>`).join("");
	  document.getElementById("online-users").innerHTML = `<ul>${list}</ul>`;
	  return;
	}
  
	// Chat message
	const time = new Date(data.timestamp).toLocaleTimeString();
	chatbox.innerHTML += `<div><strong>${data.sender}</strong> [${time}]: ${data.content}</div>`;
  };
  

document.getElementById("input").addEventListener("keydown", (e) => {
	if (e.key === "Enter") {
		const msg = e.target.value;

		ws.send(JSON.stringify(msg));
		// const chatbox = document.getElementById("chatbox");
		// chatbox.innerHTML += `<div>You: ${msg}</div>`;
		e.target.value = "";
	}
});
