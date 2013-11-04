
$(document).ready(function() {
    req = $.get("conf/names")
    req.done(function(data) {
        if (console && console.log) {
            console.log(data);
        }
        var bnames = JSON.parse(data)
        for (b in bnames) {
            btn = $("<tr><td><button type=\"button\" class=\"btn btn-lg btn-primary btn-block\" id=\"" + bnames[b] + "\"><br/>" + bnames[b] + "<br/><br/></button></td></tr>");
            $("#clickers_here").append(btn);
            $("#" + bnames[b]).click(clickeroo);
        }
    });
});

$(".btn").on("click", function(e) {
    rq = $.post("set", JSON.stringify(this.id));
});

function clickeroo(e) {
	console.log(this);
	console.log(this.id);
    rq = $.post("set", JSON.stringify(this.id));
}
