$(document).ready(function () {
    //your code here
    $("#form").submit(function(e) {
        console.log('submitted');

        e.preventDefault(); // avoid to execute the actual submit of the form.
        $.ajax({
            type: "GET",
            url: "http://localhost:4000/hello?name=" + $("#input").val(),
            success: function(data) {
                console.log(data);
            }
        });

        setInterval(getMemory, 1000);


    });

    function getMemory() {
        var url = "http://localhost:4000/memory"; // the script where you handle the form input.
        $.ajax({
            type: "GET",
            url: url,
            data: $("#form").serialize(), // serializes the form's elements.
            success: function(data)
            {
                console.log(data); // show response from the php script.
            }
        });
    }
});



