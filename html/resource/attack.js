$('#attack_btn').on("click", function(event) {
    jQuery.get("http://razor.iti.lab:8080/rscad/activate")
})

$(document).ready(function() {
    req = $.get("conf/names")
    req.done(function(data) {
        if (console && console.log) {
            console.log(data);
        }
        var bnames = JSON.parse(data)
        for (b in bnames) {
            btn = $("<button type=\"button\" class=\"btn btn-primary\" id=\"" + bnames[b] + "\">" + bnames[b] + "</button><p/>");
            btn.click(clickeroo);
            $("#clickers_here").append(btn);
        }
    });
});

$(".btn").on("click", function(e) {
    rq = $.post("set", JSON.stringify(this.id));
});

function clickeroo(e) {
    rq = $.post("set", JSON.stringify(this.id));
}
