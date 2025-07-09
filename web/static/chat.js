
const ws = new WebSocket("wss://127.0.0.1/ws");

ws.onopen = () => {
	console.log("connected to websocket")
};

ws.onmessage = (event) => {
	const chatbox = document.getElementById("chatbox");
	chatbox.innerHTML += `<div>${event.data}</div>`;
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
