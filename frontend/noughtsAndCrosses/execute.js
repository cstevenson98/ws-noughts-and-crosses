let conn = new WebSocket("ws://"+ self.location.href.slice(5) + "/connect");

conn.onclose = function (evt) {
    console.log("Connection closed.");
};

// Messages will always be the game state.
conn.onmessage = function (evt) {
    const message = evt.data.toString()
    console.log(message)

    let boxes = document.querySelectorAll(".box");
    Array.from(boxes, function(box) {
        let i = box.classList[1][1];
        let j = box.classList[1][0];

        let intI = parseInt(i);
        let intJ = parseInt(j);

        let XorO = message[3*intI + intJ]
        if (XorO == "X") {
            box.style.background = "red"
        } else if (XorO == "0") {
            box.style.background = "green"
        } else {
            box.style.background = "gray"
        }
        box.textContent = XorO
    });
};

// Waits for boxes to be clicked and sends turn info to backend.
window.addEventListener("DOMContentLoaded", function() {
    let boxes = document.querySelectorAll(".box");
    Array.from(boxes, function(box) {
        box.addEventListener("click", function() {
            conn.send("[" + this.classList[1][1] + ", " + this.classList[1][0] + "]")
        });
    });
});