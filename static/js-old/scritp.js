function changeMessage() {
    const messageElement = document.getElementById("message");
    const currentMessage = messageElement.textContent;

    const alternateMessages = [
        "Hola de nuevo",
        "Bienvenido de vuelta",
        "Que tenga un buen día"
    ];

    let newMessage;

    do {
        const randomIndex = Math.floor(Math.random() * alternateMessages.length);
        newMessage = alternateMessages[randomIndex];
    } while (newMessage === currentMessage);

    messageElement.textContent = newMessage;
}
