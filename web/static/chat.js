
// const userCode = localStorage.getItem("userCode");
const userCode = "082a60c4-c494-4415-9d23-b0cdd9c11b07"
const ws = new WebSocket(`wss://127.0.0.1/ws?code=${userCode}`);


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
