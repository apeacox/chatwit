$(function(){var c;var f=$("#msg");var b=$("#log");function e(i){var h=b[0];var g=h.scrollTop==h.scrollHeight-h.clientHeight;i.appendTo(b);if(g){h.scrollTop=h.scrollHeight-h.clientHeight}}function d(){return(new Date().toTimeString().replace(/.*(\d{2}:\d{2}:\d{2}).*/,"$1"))}function a(g){switch(g.Event){case 0:e($("<div><b>["+d()+"] "+g.Arguments.Nickname+":</b> "+g.Arguments.Body+"</div>"));break;case 1:e($("<div><b>["+d()+"] *** "+g.Arguments.Nickname+" has joined ***</b></div>"));break;case 2:e($("<div><b>["+d()+"] *** "+g.Arguments.Nickname+" has left ***</b></div>"));break}}$("#form").submit(function(){if(!c){return false}if(!f.val()){return false}c.send(f.val());f.val("");return false});if(window.WebSocket){c=new WebSocket($("#form").data("socket"));c.onclose=function(g){e($("<div><b>*** Connection closed. ***</b></div>"))};c.onmessage=function(g){console.log(g.data);o=jQuery.parseJSON(g.data);console.log(o.Event);console.log(o.Event);a(o)}}else{e($("<div><b>*** Your browser does not support WebSockets. ***</b></div>"))}});