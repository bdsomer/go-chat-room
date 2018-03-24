const socket = new WebSocket('ws://' + window.location.host + '/chatSocket')

document.addEventListener("DOMContentLoaded", function() {
	const usernameInput = document.getElementById('usernameInput')
	const messageInput = document.getElementById('messageInput')
	function send() {
		socket.send(usernameInput.value + ',' + messageInput.value)
		messageInput.value = ''
	}
	document.getElementById('sendButton').addEventListener('click', send)
	messageInput.onkeyup = (e) => {
		if (e.code === 'Enter') {
			send()
		}
	}

	const messageContainer = document.getElementById('messageContainer')
	socket.addEventListener('message', ( { data }) => {
		const msg = /(.+?),(.+)/.exec(data)
		const usernameDiv = document.createElement('div')
		usernameDiv.textContent = msg[1]
		usernameDiv.className = 'usernameDiv'
		const msgDiv = document.createElement('div')
		msgDiv.textContent = msg[2]
		msgDiv.className = 'msgDiv'
		const br = document.createElement('br')
		messageContainer.appendChild(usernameDiv)
		messageContainer.appendChild(msgDiv)
		messageContainer.appendChild(br)
	})
})