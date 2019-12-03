function sendMsg() {
    var chatLog = document.getElementById("chatbox").innerHTML;
    var msg = document.getElementById("usermsg").value;
    if (msg != '') {
        // alert(msg);
        // alert(chatLog);
        // alert(chatLog+"\n"+msg);
        chatLog = msg;
    }
}
    

