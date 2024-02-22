export function appendMessage(msg, box) {
  const d = document.createElement('div');
  d.textContent = msg;
  box.appendChild(d);
  box.scroll(0, box.scrollHeight);
}

export function joinRoom(ws) {
  ws = new WebSocket('ws://' +
        window.location.host +
        '/add?username=' +
        document.getElementById('username').value +
        '&roomid=' +
        document.getElementById('roomid').value);

  return ws;
}

export function createRoom(ws) {
  ws = new WebSocket('ws://' +
        window.location.host +
        '/create?username=' +
        document.getElementById('username').value);

  return ws;
}
