
$(document).ready(function () {
    document.getElementById("text-empty-error").style.display = 'none';
    document.getElementById("result").style.display = 'none';
    document.getElementById("calculation-error").style.display = 'none';
    $("form").submit(function (event) {
        event.preventDefault();
        a = $("#expr1").val()
        if  (a == ""){
            document.getElementById("text-empty-error").innerHTML="Please fill the text field...";
            document.getElementById("text-empty-error").style.display = 'block';
            return
        }
        var formData = {
            expr: a,
        };
    
        $.ajax({
            type: "POST",
            url: "/calculator",
            data: formData,
            dataType: "json",
            encode: true,
            // }).done(function (data, status) {
                
            // 	document.getElementById("result").innerHTML=data;
            // 	document.getElementById("result").style.display = 'block';
            // 	console.log(status);
            // });
            success: function (response, status, xhr) {
                document.getElementById("result").innerHTML=response;
                document.getElementById("result").style.display = 'block';
            },
            error: function (errMsg) {
                document.getElementById("calculation-error").innerHTML=errMsg.responseText;
                document.getElementById("calculation-error").style.display = 'block';
            },
            complete: function(resp) {
                console.log("*****complete*****");
                console.log(resp);
            }
        });

    });
});